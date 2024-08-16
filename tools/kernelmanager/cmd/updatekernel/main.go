package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.resf.org/peridot/tools/kernelmanager/packager/kernelorg"
	"go.resf.org/peridot/tools/kernelmanager/packager"
	repack_v1 "go.resf.org/peridot/tools/kernelmanager/packager/v1"

	"github.com/go-git/go-billy/v5/osfs"
)

// readCustomConfigs reads custom kernel configs from a directory
// There are two options for declaring configs
//   - A file with the name "variant-X.Y.config" where X.Y is the version. This is common arch-independent config
//   - A file with the name "variant-X.Y-ARCH.config" where X.Y is the version and arch is the architecture. This is arch-specific config
//   - Stable does not have a version, so the file name will be "stable.config" or "stable-ARCH.config"
//
// Both approaches can be used for the same variant and version. Both can be present at the same time and merges into the slice
func readCustomConfigs(directory string, variant string, version string) ([]*repack_v1.KernelConfig, error) {
	version = strings.TrimSuffix(version, ".")
	var configs []*repack_v1.KernelConfig

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".config") {
			continue
		}

		filePath := filepath.Join(directory, file.Name())

		// Stable variant
		if strings.HasPrefix(file.Name(), "stable") && variant == "stable" {

			configBytes, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}

			configMap, _, err := repack_v1.ParseConfigFile(configBytes)
			if err != nil {
				return nil, err
			}

			// Check if arch-specific
			if strings.Contains(file.Name(), "-") {
				arch := strings.TrimSuffix(strings.Split(file.Name(), "-")[1], ".config")
				log.Printf("applying arch-specific config: %s", arch)

				configs = append(configs, &repack_v1.KernelConfig{
					Arch: arch,
					Map:  configMap,
				})
				continue
			} else {
				log.Println("applying arch-independent config")
			}

			configs = append(configs, &repack_v1.KernelConfig{
				Arch: "all",
				Map:  configMap,
			})
			continue
		}

		// Other variants
		if strings.HasPrefix(file.Name(), variant+"-"+version) {
			log.Printf("reading config: %s", file.Name())

			configBytes, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}

			configMap, _, err := repack_v1.ParseConfigFile(configBytes)
			if err != nil {
				return nil, err
			}

			// Check if arch-specific
			newName := strings.TrimPrefix(file.Name(), variant+"-"+version+"-")
			if newName != ".config" {
				arch := strings.TrimSuffix(newName, ".config")
				log.Printf("applying arch-specific config: %s", arch)

				configs = append(configs, &repack_v1.KernelConfig{
					Arch: arch,
					Map:  configMap,
				})
				continue
			} else {
				log.Println("applying arch-independent config")
			}
		}
	}

	return configs, nil
}

func main() {
	// variant can only be one of the following: "stable", "longterm"
	allowedVariants := []string{"stable", "longterm"}
	variant := flag.String("variant", "", "variant to use")
	version := flag.String("version", "", "version to use")
	changelogFile := flag.String("changelog", "/tmp/changelog.json", "changelog file to use")
	kernelDirectory := flag.String("kernel-directory", "", "directory to store kernel")
	configDirectory := flag.String("config-directory", "configs", "directory to store configs")
	sourceOutputDirectory := flag.String("source-output-directory", "/tmp/sources", "directory to store source output")

	flag.Parse()

	// Validation
	isAllowedVariant := false
	for _, v := range allowedVariants {
		if *variant == v {
			isAllowedVariant = true
			break
		}
	}
	if !isAllowedVariant {
		log.Fatalf("variant %s is not allowed", *variant)
	}
	if *variant != "stable" && len(*version) == 0 {
		log.Fatalf("version is required")
	}
	if *variant == "stable" && len(*version) > 0 {
		log.Fatalf("version is not allowed for stable variant")
	}
	if len(*version) > 0 && !strings.HasSuffix(*version, ".") {
		*version = *version + "."
	}
	if len(*kernelDirectory) == 0 {
		log.Fatalf("kernel directory is required")
	}

	// Read changelog file
	changelog, err := os.ReadFile(*changelogFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("failed to read changelog file: %v", err)
	}
	if err != nil && os.IsNotExist(err) {
		log.Printf("changelog file %s does not exist, continuing without previous changelog", *changelogFile)
	}

	var changelogEntries []*repack_v1.ChangelogEntry
	if changelog != nil && len(changelog) > 0 {
		err = json.Unmarshal(changelog, &changelogEntries)
		if err != nil {
			log.Fatalf("failed to unmarshal changelog: %v", err)
		}
	}

	// Get kernel tarball and version
	var tarball []byte
	var output *packager.Output
	var kVersion string

	// BuildID should be YYYYMMDDHHMM
	buildID := time.Now().Format("200601021504")

	// Check if there are any custom config files
	customConfigs, err := readCustomConfigs(*configDirectory, *variant, *version)
	if err != nil {
		log.Fatalf("failed to read custom configs: %v", err)
	}

	switch *variant {
	case "stable":
		kVersion, tarball, _, err = kernelorg.GetLatestStable()
	case "longterm":
		kVersion, tarball, _, err = kernelorg.GetLatestLT(*version)
	}

	changelogEntry := &repack_v1.ChangelogEntry{
		Date:    time.Now().Format("Mon Jan 02 2006"),
		Name:    "RESF AutoPackager",
		Version: kVersion,
		BuildID: buildID,
		Messages: []string{
			fmt.Sprintf("Rebase to %s", kVersion),
		},
	}

	// Prepend the new changelog entry
	changelogEntries = append([]*repack_v1.ChangelogEntry{changelogEntry}, changelogEntries...)

	input := &repack_v1.Input{
		Version:                kVersion,
		BuildID:                buildID,
		KernelPackage:          "kernel",
		Changelog:              changelogEntries,
		AdditionalKernelConfig: customConfigs,
		Tarball:                tarball,
	}

	if err != nil {
		log.Fatalf("failed to get kernel: %v", err)
	}

	output, err = repack_v1.ML(input)
	if err != nil {
		log.Fatalf("failed to repack: %v", err)
	}

	// Write files
	_ = os.MkdirAll(*kernelDirectory, 0755)
	fs := osfs.New(*kernelDirectory)
	err = output.ToFS(fs)
	if err != nil {
		log.Fatalf("failed to write files: %v", err)
	}

	_ = os.MkdirAll(*sourceOutputDirectory, 0755)
	f, err := os.Create(filepath.Join(*sourceOutputDirectory, output.TarballName))
	if err != nil {
		log.Fatalf("failed to create tarball: %v", err)
	}

	_, err = f.Write(output.Tarball)
	if err != nil {
		log.Fatalf("failed to write tarball: %v", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("failed to close tarball: %v", err)
	}

	for _, otherFile := range output.OtherFiles {
		f, err := os.OpenFile(filepath.Join(*sourceOutputDirectory, otherFile.Name), os.O_CREATE|os.O_RDWR|os.O_TRUNC, otherFile.Permissions)
		if err != nil {
			log.Fatalf("failed to create file: %v", err)
		}

		_, err = f.Write(otherFile.Data)
		if err != nil {
			log.Fatalf("failed to write file: %v", err)
		}

		err = f.Close()
		if err != nil {
			log.Fatalf("failed to close file: %v", err)
		}
	}

	log.Println("repack successful :)")
}
