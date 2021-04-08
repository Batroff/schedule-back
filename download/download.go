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
	defer resp.Body.Close()

	// Create the file
	log.Printf("Create file: %s...", filepath)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	log.Printf("Copying to file...")
	_, err = io.Copy(out, resp.Body)
	return err
}
