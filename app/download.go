package app

import (
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

func GetFile(filepath, url string) error {
	// Get the data
	resp, downloadErr := http.Get(url)
	if downloadErr != nil {
		return errors.Wrap(downloadErr, "Download error.")
	}
	log.Printf("File successfuly downloaded from %s", url)

	// Create the file
	out, createErr := os.Create(filepath)
	if createErr != nil {
		return errors.Wrap(createErr, "Creating file error.")
	}
	log.Printf("File %s created.", filepath)

	// Write the body to file
	_, copyErr := io.Copy(out, resp.Body)
	if copyErr != nil {
		return errors.Wrap(copyErr, "Copying file error.")
	}

	respErr := resp.Body.Close()
	if respErr != nil {
		return errors.Wrap(respErr, "Response closing error")
	}

	fileErr := out.Close()
	if fileErr != nil {
		return errors.Wrap(fileErr, "File closing error")
	}

	return nil
}

func GetScheduleXlsx(dirPath string, urls []string) ([]string, error) {
	scheduleFiles := make([]string, len(urls))

	for i, url := range urls {
		currPath, pathErr := os.Getwd()
		if pathErr != nil {
			return nil, pathErr
		}

		workDir := currPath + dirPath
		if _, err := os.Stat(workDir); os.IsNotExist(err) {
			if createDirErr := os.Mkdir(workDir, os.ModeDir); createDirErr != nil {
				return nil, createDirErr
			}
		}

		excelPath := path.Join(workDir, "schedule_"+strconv.Itoa(i)+".xlsx")

		err := GetFile(excelPath, url)
		if err != nil {
			return nil, err
		}

		scheduleFiles[i] = excelPath
	}

	return scheduleFiles, nil
}
