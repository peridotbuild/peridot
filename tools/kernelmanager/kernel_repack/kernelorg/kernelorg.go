package kernelorg

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"github.com/xi2/xz"
	"golang.org/x/crypto/openpgp"
	"io"
	"net/http"
)

//go:embed gregkh.asc
var gregKHPublicKey []byte

//go:embed torvalds.asc
var torvaldsPublicKey []byte

func getKeyring() (openpgp.EntityList, error) {
	var entityList openpgp.EntityList
	keys := [][]byte{gregKHPublicKey, torvaldsPublicKey}
	for _, key := range keys {
		keyRing, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(key))
		if err != nil {
			return nil, err
		}

		entityList = append(entityList, keyRing...)
	}

	return entityList, nil
}

func verifyTarball(tarball []byte, gz bool, signature []byte) (*openpgp.Entity, error) {
	// unpack tarball
	var unpackedTarball []byte
	if gz {
		gzipRead, err := gzip.NewReader(bytes.NewReader(tarball))
		if err != nil {
			return nil, err
		}

		unpackedTarball, err = io.ReadAll(gzipRead)
		if err != nil {
			return nil, err
		}
	} else {
		xzRead, err := xz.NewReader(bytes.NewReader(tarball), 0)
		if err != nil {
			return nil, err
		}

		unpackedTarball, err = io.ReadAll(xzRead)
		if err != nil {
			return nil, err
		}
	}

	keyRing, err := getKeyring()
	if err != nil {
		return nil, err
	}

	entity, err := openpgp.CheckArmoredDetachedSignature(keyRing, bytes.NewReader(unpackedTarball), bytes.NewReader(signature))
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func download(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func downloadKernel(version string) ([]byte, *openpgp.Entity, error) {
	firstDigit := version[0:1]

	downloadURL := fmt.Sprintf("https://cdn.kernel.org/pub/linux/kernel/v%s.x/linux-%s.tar.xz", firstDigit, version)
	tarball, err := download(downloadURL)
	if err != nil {
		return nil, nil, err
	}

	signatureURL := fmt.Sprintf("https://cdn.kernel.org/pub/linux/kernel/v%s.x/linux-%s.tar.sign", firstDigit, version)
	signature, err := download(signatureURL)
	if err != nil {
		return nil, nil, err
	}

	entity, err := verifyTarball(tarball, false, signature)
	if err != nil {
		return nil, nil, err
	}

	return tarball, entity, nil
}

func downloadLT(majorVersion string) (string, []byte, *openpgp.Entity, error) {
	latestVersion, err := GetLTVersion(majorVersion)
	if err != nil {
		return "", nil, nil, err
	}

	tarball, entity, err := downloadKernel(latestVersion)
	if err != nil {
		return "", nil, nil, err
	}

	return latestVersion, tarball, entity, nil
}

func GetLatestML() (string, []byte, *openpgp.Entity, error) {
	latestVersion, err := GetLatestMLVersion()
	if err != nil {
		return "", nil, nil, err
	}

	downloadURL := fmt.Sprintf("https://git.kernel.org/torvalds/t/linux-%s.tar.gz", latestVersion)
	tarball, err := download(downloadURL)
	if err != nil {
		return "", nil, nil, err
	}

	// ML RC doesn't contain signature, so we're relying on TLS

	return latestVersion, tarball, nil, nil
}

func GetLatestLT(prefix string) (string, []byte, *openpgp.Entity, error) {
	return downloadLT(prefix)
}

func GetLatestStable() (string, []byte, *openpgp.Entity, error) {
	latestVersion, err := GetLatestStableVersion()
	if err != nil {
		return "", nil, nil, err
	}

	tarball, entity, err := downloadKernel(latestVersion)
	if err != nil {
		return "", nil, nil, err
	}

	return latestVersion, tarball, entity, nil
}
