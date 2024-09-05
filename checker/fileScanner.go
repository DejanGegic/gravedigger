package checker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Gochecker struct {
	Path   string
	Hash   []byte
	IsTest bool
}

var checkerList []Gochecker

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
	if !strings.HasSuffix(info.Name(), ".go") { // Skip non-go checkers
		return nil
	}
	// Write depending on if it's a test
	if strings.HasSuffix(info.Name(), "_test.go") {
		checkerList = append(checkerList, Gochecker{Path: path, IsTest: true})
	} else {
		checkerList = append(checkerList, Gochecker{Path: path, IsTest: false})
	}

	return nil
}
