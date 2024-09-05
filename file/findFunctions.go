package file

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
	Instances []FunctionInstance
}
type FunctionInstance struct {
	Path string
	Line uint32
}

var FunctionsList = make(map[string]FunctionData, 0)

func FindAllFunctionDeclarations() {
	for _, file := range FileList {
		if file.IsTest {
			continue
		}
		FindAllFunctionsInAFile(file.Path)
	}
}

func FindAllFunctionsInAFile(path string) {
	// scan file line by line
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

			filePathSplit := strings.Split(path, "/")
			currentPackage := filePathSplit[len(filePathSplit)-2]

			var functionName string
			// depending on if its an interface implementation or not
			if strings.HasPrefix(line, "func (") {
				functionName = strings.Split(line, ")")[1]
				functionName = strings.Split(functionName, "(")[0]
				functionName = strings.Split(functionName, " ")[1]

			} else {
				functionName = strings.Split(line, " ")[1]
				functionName = strings.Split(functionName, "(")[0]
			}

			FunctionsList[path+"|"+functionName] = FunctionData{
				Package: currentPackage,
				Name:    functionName,
				Line:    lineNumber,
			}
		}
	}
}
