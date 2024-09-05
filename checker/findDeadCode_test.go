package checker

import (
	"testing"
)

func TestCheckIfImplementationPackageIsCorrect(t *testing.T) {
	result := checkIfImplementationPackageIsCorrect("uci.ListenLoop()", "uci", "main", "ListenLoop()")
	if !result {
		t.Error("Test failed")
	}
	result = checkIfImplementationPackageIsCorrect("UCICommand()", "uci", "uci", "UCICommand()")
	if !result {
		t.Error("Test failed")
	}
}
