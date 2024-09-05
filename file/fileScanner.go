package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type GoFile struct {
	Path   string
	Hash   []byte
	IsTest bool
}

var FileList []GoFile

func ScanAllSubDirs(path string) error {
	err := filepath.Walk(path, WalkFunction)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func WalkFunction(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	// Skip hidden dirs starting with "."
	if strings.HasPrefix(info.Name(), ".") && info.IsDir() {
		return filepath.SkipDir
	}
	if !strings.HasSuffix(info.Name(), ".go") { // Skip non-go files
		return nil
	}
	// Write depending on if it's a test
	if strings.HasSuffix(info.Name(), "_test.go") {
		FileList = append(FileList, GoFile{Path: path, IsTest: true})
	} else {
		FileList = append(FileList, GoFile{Path: path, IsTest: false})
	}

	return nil
}
