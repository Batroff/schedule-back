package download

import (
	"io"
	"log"
	"net/http"
	"os"
)

func GetFile(filepath, url string) error {
	// Get the data
	log.Printf("Getting url: %s...", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Create the file
	log.Printf("Create file: %s...", filepath)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}

	// Write the body to file
	log.Printf("Copying to file...")
	_, err = io.Copy(out, resp.Body)

	defer func() {
		respErr := resp.Body.Close()
		if respErr != nil {
			log.Panicf("Error occured while response closing, %v", respErr)
		}

		fileErr := out.Close()
		if fileErr != nil {
			log.Panicf("Error occured while file closing, %v", fileErr)
		}
	}()

	return err
}
