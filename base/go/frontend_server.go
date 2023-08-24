// Copyright 2023 Peridot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package base

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type FrontendInfo struct {
	// Title to add to the HTML page
	Title string

	// Self is the URL to the frontend server
	Self string

	// NoAuth is a flag to disable authentication
	NoAuth bool

	// AllowUnauthenticated is a flag to allow unauthenticated users
	AllowUnauthenticated bool

	// OIDCIssuer is the issuer to use for authentication
	OIDCIssuer string

	// OIDCClientID is the client ID to use for authentication
	OIDCClientID string

	// OIDCClientSecret is the client secret to use for authentication
	OIDCClientSecret string

	// OIDCGroup is the group to check for authentication
	OIDCGroup string

	// OIDCUserInfoOverride is a flag to override the userinfo endpoint
	OIDCUserInfoOverride string

	// AdditionalContent is a map of paths to content to serve
	AdditionalContent map[string][]byte
}

var frontendHtmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, viewport-fit=cover"
    />
    <title>{{.Title}}</title>

    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link
      rel="stylesheet"
      href="https://fonts.googleapis.com/css2?family=Roboto:wght@300;400;500;600;700&display=swap"
    />
</head>
<body>
    <div id="app"></div>

    <script src="{{.BundleJS}}"></script>
</body>
</html>
`

func readDir(embedfs *embed.FS, root string) ([]string, error) {
	var paths []string
	entries, err := embedfs.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			subPaths, err := readDir(embedfs, filepath.Join(root, entry.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, subPaths...)
		} else {
			paths = append(paths, filepath.Join(root, entry.Name()))
		}
	}
	return paths, nil
}

func FrontendServer(info *FrontendInfo, embedfs *embed.FS) {
	port := 9111
	if info == nil {
		info = &FrontendInfo{}
	}

	newTemplate := frontendHtmlTemplate

	// Set the title
	if info.Title == "" {
		info.Title = "Peridot"
	}
	newTemplate = strings.ReplaceAll(newTemplate, "{{.Title}}", info.Title)

	pathToContent := map[string][]byte{}

	// Read the files from the embedfs
	paths, err := readDir(embedfs, "bundle")
	if err != nil {
		LogFatalf("failed to read embedfs: %v", err)
	}

	for _, path := range paths {
		content, err := embedfs.ReadFile(path)
		if err != nil {
			LogFatalf("failed to read embedfs: %v", err)
		}

		// Sha256 hash of the content to add to name
		hash := sha256.New()
		hash.Write(content)
		hashSum := hex.EncodeToString(hash.Sum(nil))

		ext := filepath.Ext(path)
		noExtName := path[:len(path)-len(ext)]
		newPath := fmt.Sprintf("/_ga/%s.%s%s", noExtName, hashSum[:8], ext)

		pathToContent[newPath] = content

		// If name is entrypoint.js, replace the template
		if filepath.Base(path) == "entrypoint.js" {
			newTemplate = strings.ReplaceAll(newTemplate, "{{.BundleJS}}", newPath)
		}
	}

	// Add additional content
	if info.AdditionalContent != nil {
		for path, content := range info.AdditionalContent {
			pathToContent[path] = content
		}
	}

	// Log the paths
	LogInfof("frontend server paths:")
	for path := range pathToContent {
		LogInfof("  %s", path)
	}

	// Serve the content
	http.HandleFunc("/_ga/", func(w http.ResponseWriter, r *http.Request) {
		mimeType := mime.TypeByExtension(filepath.Ext(r.URL.Path))
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}
		w.Header().Set("Content-Type", mimeType)
		_, _ = w.Write([]byte(pathToContent[r.URL.Path]))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(newTemplate))
	})

	// Handle other _ga meta routes
	http.HandleFunc("/_ga/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("ok"))
	})

	// Handle auth routes
	http.HandleFunc("/auth/oidc/login", func(w http.ResponseWriter, r *http.Request) {

	})

	LogInfof("starting frontend server on port %d", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		LogFatalf("failed to start frontend server: %v", err)
	}
}
