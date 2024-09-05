package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindDeadCode() {
	for _, file := range FileList {
		// skip tests
		if file.IsTest {
			continue
		}
		deadCodeSingleFile(file.Path)
	}
}

func deadCodeSingleFile(path string) {

	// scan file line by line
	fileToScan, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileToScan.Close()
	scanner := bufio.NewScanner(fileToScan)

	lineNumber := uint32(0)
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		// go through all keys of map FunctionsList
		for key, value := range FunctionsList {
			// TODO: Include a check for packages. There might be 2 "delete" functions in separate packages
			if strings.Contains(line, value.Name) && !strings.HasPrefix(line, "func ") {
				FunctionInstance := FunctionInstance{
					Path: path,
					Line: lineNumber,
				}
				// add to value.Instances
				value.Instances = append(value.Instances, FunctionInstance)
				FunctionsList[key] = value
			}
		}
	}
}

func ShowDeadCode() {
	for key, value := range FunctionsList {

		if len(value.Instances) != 0 {
			continue
		}
		fileName := strings.Split(key, "|")[0]

		fmt.Printf("%s:%d %s\n", fileName, value.Line, value.Name)

	}
}
