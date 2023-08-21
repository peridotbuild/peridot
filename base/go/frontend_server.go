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
<body style="background-color:#f1f5f9;">
    <div id="app"></div>

    <script src="{{.BundleJS}}"></script>
</body>
</html>
`

func FrontendServer(info *FrontendInfo, kv ...string) {
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

	pathToContent := map[string]string{}
	// KV is a list of key-value pairs, where the key is the alias and the value
	// is the actual file content.
	for i := 0; i < len(kv); i += 2 {
		content := kv[i+1]

		// Sha256 hash of the content to add to name
		hash := sha256.New()
		hash.Write([]byte(content))
		hashSum := hex.EncodeToString(hash.Sum(nil))

		ext := filepath.Ext(kv[i])
		noExtName := kv[i][:len(kv[i])-len(ext)]
		path := fmt.Sprintf("/_ga/%s.%s%s", noExtName, hashSum[:8], ext)

		pathToContent[path] = content

		// If name is bundle.js, replace the template
		if kv[i] == "bundle.js" {
			newTemplate = strings.ReplaceAll(newTemplate, "{{.BundleJS}}", path)
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

	LogInfof("starting frontend server on port %d", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		LogFatalf("failed to start frontend server: %v", err)
	}
}
