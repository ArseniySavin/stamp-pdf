package pkg

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func Stamp(inFile, outFile, stampPath, page, stampPosition string) error {

	params := fmt.Sprintf("pos:%s, rot:1", stampPosition)

	err := api.AddImageWatermarksFile(inFile, outFile, []string{page}, true, stampPath, params, nil)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func ExtractPageFromFile(inFile, page string) (string, error) {
	if page == "1-" {
		return "", nil
	}

	if filepath.Ext(inFile) != ".pdf" {
		return "", fmt.Errorf("%s %s", inFile, "file unsupported")
	}

	curDir := filepath.Dir(inFile)
	_, file := filepath.Split(inFile)
	//defer os.Remove(file)

	file = strings.Replace(file, ".pdf", "", 1)
	file = fmt.Sprintf("%s/%s_page_%s.pdf", curDir, file, page)

	err := api.ExtractPagesFile(inFile, curDir, []string{page}, nil)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return file, nil
}

func AddStamp(inFile, outFile, page, stampTemplate, pos string) error {

	commands := fmt.Sprintf("pos:%s, rot:1", pos)

	err := api.AddImageWatermarksFile(inFile, outFile, []string{page}, true, stampTemplate, commands, nil)
	if err != nil {
		return err
	}

	return nil
}
