package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func isPDF(filename string) bool {
	return strings.HasSuffix(filename, ".pdf")
}

func main() {
	files := []string{"http://www.africau.edu/images/default/sample.pdf", "https://www.clickdimensions.com/links/TestPDFfile.pdf"}

	var pdfFiles []string
	for idx, file := range files {
		filePath := filepath.Join("files", "file"+strconv.Itoa(idx)+".pdf")
		err := DownloadFile(filePath, file)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + file)

		// pdf, err := filePath.Open()
		pdfFiles = append(pdfFiles, filePath)
	}

	pdfcpu.MergeCreateFile(pdfFiles, "merged.pdf", nil)
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	println(url)

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
