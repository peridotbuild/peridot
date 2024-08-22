package repack_v1

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
	"text/template"

	 packager "go.resf.org/peridot/tools/kernelmanager/kernel_repack"
)

//go:embed data/*
var Data embed.FS

type ChangelogEntry struct {
	Date     string   `json:"-"`
	Name     string   `json:"-"`
	Version  string   `json:"-"`
	BuildID  string   `json:"-"`
	Messages []string `json:"messages"`

	// Subject covers the first line of the changelog entry
	Subject string `json:"subject"`
}

type KernelConfig struct {
	Arch string
	Map  map[string]string
}

type Input struct {
	Version                string
	BuildID                string
	KernelPackage          string
	Changelog              []*ChangelogEntry
	AdditionalKernelConfig []*KernelConfig
	Tarball                []byte
}

// compareVersion takes in two kernel version strings either in the form of "X.Y" or "X.Y.Z" and returns:
// 1 if a > b
// 0 if a == b
// -1 if a < b
func compareVersion(a, b string) int {
	// Split version strings into parts
	aParts := strings.Split(a, ".")
	bParts := strings.Split(b, ".")

	// Compare major version
	if aParts[0] != bParts[0] {
		if aParts[0] > bParts[0] {
			return 1
		}
		return -1
	}

	// Compare minor version
	if aParts[1] != bParts[1] {
		if aParts[1] > bParts[1] {
			return 1
		}
		return -1
	}

	// If there's a third part, compare it
	if len(aParts) == 3 && len(bParts) == 3 {
		if aParts[2] != bParts[2] {
			if aParts[2] > bParts[2] {
				return 1
			}
			return -1
		}
	}

	return 0
}

func kernel(kernelType string, in *Input) (*packager.Output, error) {
	var spec *packager.File
	var files []*packager.File

	dir, err := Data.ReadDir("data")
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".spec") {
			if file.Name() != kernelType+".spec" {
				continue
			}

			// Read spec file
			f, err := Data.Open("data/" + file.Name())
			if err != nil {
				return nil, err
			}
			defer f.Close()

			specBytes, err := io.ReadAll(f)
			if err != nil {
				return nil, err
			}
			specName := fmt.Sprintf("%s.spec", in.KernelPackage)
			spec = &packager.File{
				Name: specName,
				Data: specBytes,
			}

			continue
		}

		// Read other files
		f, err := Data.Open("data/" + file.Name())
		if err != nil {
			return nil, err
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}

		// If the file starts with "config-", then it's a kernel config file.
		// Merge additional kernel config
		if strings.HasPrefix(file.Name(), "config-") {
			parsedConfig, comments, err := ParseConfigFile(data)
			if err != nil {
				return nil, err
			}

			for _, conf := range in.AdditionalKernelConfig {
				if conf.Arch == strings.TrimPrefix(file.Name(), "config-") || conf.Arch == "all" {
					for k, v := range conf.Map {
						parsedConfig[k] = v
					}
				}
			}

			// Write the config back to the data slice
			data = WriteConfigFile(parsedConfig, comments)
		}

		stat, err := f.Stat()
		if err != nil {
			return nil, err
		}
		mode := stat.Mode()

		// If file name ends with ".sh", set executable bit
		if strings.HasSuffix(file.Name(), ".sh") {
			mode |= 0111
		}

		// Read first line of file. If it starts with "#!", set executable bit
		if len(data) > 2 && data[0] == 0x23 && data[1] == 0x21 {
			mode |= 0111
		}

		files = append(files, &packager.File{
			Name:        file.Name(),
			Data:        data,
			Permissions: mode,
		})
	}

	// Get sha256sum of tarball
	hash := sha256.New()
	_, err = hash.Write(in.Tarball)
	if err != nil {
		return nil, err
	}
	sum := hex.EncodeToString(hash.Sum(nil))

	// Create .metadata file
	suffix := "xz"
	if in.Tarball[0] == 0x1f && in.Tarball[1] == 0x8b {
		suffix = "gz"
	}
	tarballName := fmt.Sprintf("linux-%s.tar.%s", in.Version, suffix)
	metadata := fmt.Sprintf("%s SOURCES/%s", tarballName)

	// Create .[kernelpackage].metadata
	metadataName := fmt.Sprintf(".%s.metadata", in.KernelPackage)

	// Replace placeholders in spec file
	var buf bytes.Buffer
	txtTemplate, err := template.
		New("spec").
		Funcs(template.FuncMap{
			"compareVersion": compareVersion,
		}).
		Parse(string(spec.Data))
	if err != nil {
		return nil, err
	}
	err = txtTemplate.Execute(&buf, in)
	if err != nil {
		return nil, err
	}
	spec.Data = buf.Bytes()

	output := &packager.Output{
		Spec:          spec,
		Tarball:       in.Tarball,
		TarballName:   tarballName,
		TarballSha256: sum,
		Metadata: &packager.File{
			Name: metadataName,
			Data: []byte(metadata),
		},
		OtherFiles: files,
	}

	return output, nil
}

func ParseConfigFile(data []byte) (map[string]string, []string, error) {
	// Split into lines so we can parse it
	lines := strings.Split(string(data), "\n")

	var comments []string
	var config []string
	// Loop, and add all lines that start with "#" to the comment slice
	// and all other lines to the config slice
	for _, line := range lines {
		// Ignore empty lines
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, "#") {
			comments = append(comments, line)
		} else {
			config = append(config, line)
		}
	}

	parsedConfig := make(map[string]string)
	for _, line := range config {
		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid kernel config line: %s", line)
		}
		parsedConfig[parts[0]] = parts[1]
	}

	return parsedConfig, comments, nil
}

func WriteConfigFile(config map[string]string, comments []string) []byte {
	var newConfig []string
	for _, comment := range comments {
		newConfig = append(newConfig, comment)
	}
	for k, v := range config {
		newConfig = append(newConfig, fmt.Sprintf("%s=%s", k, v))
	}

	data := []byte(strings.Join(newConfig, "\n"))

	return data
}

// ML creates a new kernel package for the ML kernel
// Returns spec and SOURCE files
func ML(in *Input) (*packager.Output, error) {
	return kernel("ml", in)
}
