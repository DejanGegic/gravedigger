package checker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindDeadCode() {
	for _, goFile := range fileList {
		// skip tests
		if goFile.IsTest {
			continue
		}
		deadCodeSingleChecker(goFile.Path)
	}
}

func deadCodeSingleChecker(path string) {

	// scan checker line by line
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
			// ! Slowest part of the whole codebase, taking over 85% of CPU TIME
			// TODO: Parallelize using a pool
			if strings.Contains(line, value.Name) && !strings.HasPrefix(line, "func ") {

				if !value.IsAMethod {
					if !checkIfImplementationPackageIsCorrect(line, value.Package, getCurrentPackage(path), value.Name) {
						continue
					}
				}

				currentPackage := getCurrentPackage(path)
				FunctionInstance := FunctionInstance{
					Path:    path,
					Line:    lineNumber,
					Package: currentPackage,
				}
				// add to value.Instances
				value.Instances = append(value.Instances, FunctionInstance)
				FunctionsList[key] = value
			}
		}
	}
}
func checkIfImplementationPackageIsCorrect(codeLine string, packageOfSourceFunction string, packageOfCurrentImplementation string, functionName string) bool {

	if packageOfCurrentImplementation == packageOfSourceFunction {
		return true
	}

	// Check if the function is from the package we're looking for
	toCheck := packageOfSourceFunction + "." + functionName
	if strings.Contains(codeLine, toCheck) {
		return true
	}

	// Check if it belongs to another package
	if strings.Contains(codeLine, "."+functionName) {
		return false
	}

	return false
}

func ShowDeadCode() {
	for key, value := range FunctionsList {

		if len(value.Instances) != 0 {
			continue
		}
		checkerName := strings.Split(key, "|")[0]

		fmt.Printf("%s:%d %s\n", checkerName, value.Line, value.Name)

	}
}
