package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractFile(t *testing.T) {

	startupPath, _ := os.Getwd()
	startupPath = filepath.Dir(startupPath)
	inFile := fmt.Sprintf("%s/in.pdf", startupPath)
	outFile := fmt.Sprintf("%s/in_page_1.pdf", startupPath)

	file, err := ExtractPageFromFile(inFile, "1")

	if file != outFile {
		t.Fail()
	}

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	//os.Remove(file)
}

func TestAddStamp(t *testing.T) {

	startupPath, _ := os.Getwd()
	startupPath = filepath.Dir(startupPath)
	inFile := fmt.Sprintf("%s/in.pdf", startupPath)
	outFile := fmt.Sprintf("%s/out.pdf", startupPath)
	stampTemplate := fmt.Sprintf("%s/stamp_template.png", startupPath)

	err := AddStamp(inFile, outFile, "1-", stampTemplate, "bl")

	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
