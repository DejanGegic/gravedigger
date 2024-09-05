package checker

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type GoChecker struct {
	Path   string
	Hash   []byte
	IsTest bool
}

var fileList []GoChecker

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
		fmt.Println(err)
		return err
	}
	// Skip hidden dirs starting with "."
	if strings.HasPrefix(info.Name(), ".") && info.IsDir() {
		return filepath.SkipDir
	}
	if strings.HasPrefix(info.Name(), ".") {
		return nil
	}
	if !strings.HasSuffix(info.Name(), ".go") && !info.IsDir() { // Skip non-go checkers
		return nil
	}
	// Write depending on if it's a test
	if strings.HasSuffix(info.Name(), "_test.go") {
		fileList = append(fileList, GoChecker{Path: path, IsTest: true})
	} else {
		fileList = append(fileList, GoChecker{Path: path, IsTest: false})
	}

	return nil
}
