package generate

import (
	"os"
	"testing"
	"time"
)

func TestSelectGoFiles(t *testing.T) {
	files := []os.FileInfo{
		NewFakeFile("yes_test.go", 5, 0644, time.Now(), false),
		NewFakeFile(GeneratedFilename, 123, 0644, time.Now(), false),
		NewFakeFile("no.go", 6, 0644, time.Now(), false),
		NewFakeFile("no.txt", 6, 0644, time.Now(), false),
		NewFakeFile("no", 1, 0644, time.Now(), true),
	}

	actual := SelectGoTestFiles(files)
	if len(actual) != 1 {
		t.Errorf("Expected only one file, got: %d", len(actual))
		t.FailNow()
	}
	if name := actual[0].Name(); name != "yes_test.go" {
		t.Errorf("Expected 'yes.go' as the only file. Got '%s'", name)
	}
}

type FakeFile struct {
	name     string
	size     int64
	mode     os.FileMode
	modified time.Time
	isDir    bool
}

func NewFakeFile(name string, size int64, mode os.FileMode, modified time.Time, isDir bool) os.FileInfo {
	return &FakeFile{
		name:     name,
		size:     size,
		mode:     mode,
		modified: modified,
		isDir:    isDir,
	}
}

func (this *FakeFile) Name() string       { return this.name }
func (this *FakeFile) Size() int64        { return this.size }
func (this *FakeFile) ModTime() time.Time { return this.modified }
func (this *FakeFile) Mode() os.FileMode  { return this.mode }
func (this *FakeFile) IsDir() bool        { return this.isDir }
func (this *FakeFile) Sys() interface{}   { return this }
