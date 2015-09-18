package generate

import (
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
)

func CodeListing(directory string) (map[string]string, error) {
	code := make(map[string]string)

	listing, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}
	for _, file := range SelectGoTestFiles(listing) {
		content, err := ioutil.ReadFile(filepath.Join(directory, file.Name()))
		if err != nil {
			return nil, err
		}
		code[file.Name()] = hex.EncodeToString(content)
	}

	return code, nil
}
