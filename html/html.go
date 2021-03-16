package html

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"schedule/download"
)

func GetLinks(filepath string) []string {
	const (
		scheduleURL string = "https://www.mirea.ru/schedule/"
	)

	var (
		links []string
	)

	err := download.GetFile(filepath, scheduleURL)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	urlPath := regexp.MustCompile("https://.*xlsx")
	hrefElement := regexp.MustCompile("<a.*class=\"uk-link-toggle\".*href=\"https://.*..xlsx\".*>")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if hrefElement.MatchString(line) {
			links = append(links, urlPath.FindString(line))
		}
	}

	fmt.Println(links)
	return links
}
