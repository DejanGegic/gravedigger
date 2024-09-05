package checker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FunctionData struct {
	Package   string
	Name      string
	Line      uint32
	IsAMethod bool
	Instances []FunctionInstance
}
type FunctionInstance struct {
	Path    string
	Line    uint32
	Package string
}

var FunctionsList = make(map[string]FunctionData, 0)

func FindAllFunctionDeclarations() {
	for _, checker := range fileList {
		if checker.IsTest {
			continue
		}
		FindAllFunctionsInAFile(checker.Path)
	}
}

func FindAllFunctionsInAFile(path string) {
	// scan checker line by line
	fileToScan, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileToScan.Close()
	scanner := bufio.NewScanner(fileToScan)

	var lineNumber uint32 = 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if strings.HasPrefix(line, "func ") {

			currentPackage := getCurrentPackage(path)

			var functionName string
			var isAMethod bool = false
			// depending on if its an interface implementation or not
			if strings.HasPrefix(line, "func (") {
				functionName = strings.Split(line, ")")[1]
				functionName = strings.Split(functionName, "(")[0]

				functionName = strings.TrimSpace(functionName)

				isAMethod = true
			} else {
				functionName = strings.Split(line, " ")[1]
				functionName = strings.Split(functionName, "(")[0]
			}

			FunctionsList[path+"|"+functionName] = FunctionData{
				Package:   currentPackage,
				Name:      functionName,
				Line:      lineNumber,
				IsAMethod: isAMethod,
			}
		}
	}
}

func getCurrentPackage(path string) string {
	checkerPathSplit := strings.Split(path, "/")
	currentPackage := checkerPathSplit[len(checkerPathSplit)-2]
	return currentPackage
}
