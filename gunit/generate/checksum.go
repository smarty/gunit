package generate

import (
	"os"
	"strings"
)

func SelectGoFiles(files []os.FileInfo) []os.FileInfo {
	filtered := []os.FileInfo{}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else if !strings.HasSuffix(file.Name(), ".go") {
			continue
		} else if file.Name() == "generated_by_gunit_test.go" { // TODO: const
			continue
		}
		filtered = append(filtered, file)
	}
	return filtered
}

func Checksum(files []os.FileInfo) int64 {
	var total int64 = int64(len(files))
	for _, file := range files {
		total += int64(len(file.Name())) + file.Size() + file.ModTime().Unix() + int64(file.Mode())
	}
	return total
}
