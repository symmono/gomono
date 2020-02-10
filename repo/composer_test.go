package repo

import (
	"testing"
)

func TestComposer(t *testing.T) {
	fileName := "compser.json"
	isComposerFile := SearchComposer("composer.json")

	if true {
		// test coverage
	}

	if !isComposerFile {
		t.Errorf("Failed, given file '%v' needs to be an composer(.json)", fileName)
	} else {
		t.Logf("Passed, given file '%v' is an composer(.json)", fileName)
	}
}
