package generate

import (
	"os"
	"testing"
	"time"
)

func TestSelectGoFiles(t *testing.T) {
	files := []os.FileInfo{
		NewFakeFile("yes.go", 5, 0644, time.Now(), false),
		NewFakeFile("generated_by_gunit_test.go", 123, 0644, time.Now(), false),
		NewFakeFile("no.txt", 6, 0644, time.Now(), false),
		NewFakeFile("no", 1, 0644, time.Now(), true),
	}

	actual := SelectGoFiles(files)
	if len(actual) != 1 {
		t.Errorf("Expected only one file, got: %d", len(actual))
		t.FailNow()
	}
	if name := actual[0].Name(); name != "yes.go" {
		t.Errorf("Expected 'yes.go' as the only file. Got '%s'", name)
	}
}

func TestChecksumCalculation(t *testing.T) {
	now := time.Now()
	files1 := []os.FileInfo{
		NewFakeFile("1.go", 1, 1, now, false),
		NewFakeFile("2.go", 2, 2, now, false),
	}
	files2 := []os.FileInfo{ // only difference are the bits of the first file (by 9):
		NewFakeFile("1.go", 1, 10, now, false),
		NewFakeFile("2.go", 2, 2, now, false),
	}

	checksum1 := Checksum(files1)
	checksum2 := Checksum(files2)

	expectedDiff := int64(files2[0].Mode() - files1[0].Mode())
	actualDiff := checksum2 - checksum1

	if actualDiff != expectedDiff {
		t.Errorf("The only difference were the permission bits, which should have been different by %d. The difference was %d instead.", expectedDiff, actualDiff)
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

func (self *FakeFile) Name() string       { return self.name }
func (self *FakeFile) Size() int64        { return self.size }
func (self *FakeFile) ModTime() time.Time { return self.modified }
func (self *FakeFile) Mode() os.FileMode  { return self.mode }
func (self *FakeFile) IsDir() bool        { return self.isDir }
func (self *FakeFile) Sys() interface{}   { return self }
