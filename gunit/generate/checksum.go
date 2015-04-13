package generate

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Checksum(directory string) (string, error) {
	listing, err := ioutil.ReadDir(directory)
	if err != nil {
		return "", err
	}
	goContents, err := ReadFiles(directory, SelectGoFiles(listing))
	if err != nil {
		return "", err
	}
	hash := md5.Sum(goContents)
	buffer := make([]byte, len(hash))
	copy(buffer, hash[:])
	return hex.EncodeToString(buffer), nil
}

func SelectGoFiles(files []os.FileInfo) []os.FileInfo {
	filtered := []os.FileInfo{}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else if !strings.HasSuffix(file.Name(), ".go") {
			continue
		} else if file.Name() == GeneratedFilename {
			continue
		}
		filtered = append(filtered, file)
	}
	return filtered
}

func ReadFiles(directory string, files []os.FileInfo) ([]byte, error) {
	all := &bytes.Buffer{}
	for _, file := range files {
		content, err := ioutil.ReadFile(filepath.Join(directory, file.Name()))
		if err != nil {
			return nil, err
		}
		_, err = all.Write(content)
		if err != nil {
			return nil, err
		}

	}
	return all.Bytes(), nil
}
