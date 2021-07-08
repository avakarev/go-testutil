package testutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func isRelPath(p string) bool {
	sep := string(filepath.Separator)
	return strings.HasPrefix(p, "."+sep) || strings.HasPrefix(p, ".."+sep)
}

// FixturePath returns absolute path to the given fixture
func FixturePath(name string, args ...string) string {
	ext := ""
	if len(args) > 0 {
		ext = args[0]
	}
	name = name + ext
	if isRelPath(name) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return filepath.Join(wd, name)
	}
	return filepath.Join("test", "fixtures", name)
}

// Fixture returns content of the given fixture
func Fixture(t *testing.T, name string, args ...string) []byte {
	path := FixturePath(name, args...)
	bytes, err := ioutil.ReadFile(path) // #nosec
	if err != nil {
		t.Errorf("Failed to read %q fixture: %s", name, err.Error())
	}
	return bytes
}
