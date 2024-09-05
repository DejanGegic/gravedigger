package checker

import "os"

func RunAll(path string) {
	if path == "" || path == "." || path == "./" {
		path, _ = os.Getwd()
	}
	ScanAllSubDirs(path)
	FindAllFunctionDeclarations()
	FindDeadCode()
	ShowDeadCode()
}
