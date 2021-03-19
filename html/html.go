package html

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"schedule/download"
	"strings"
)

/* Get map of links
0 - Бакалавр/Специалитет;
1 - Магистр;
2 - Аспирант;
3 - Колледж;
4 - Экстерн;
*/
func Parse() map[int][]string {
	const filepath string = "schedule.html"

	downloadErr := download.GetFile(filepath, "https://www.mirea.ru/schedule/")
	if downloadErr != nil {
		log.Panicf("Download error: %v", downloadErr)
	}

	dat, readFileErr := ioutil.ReadFile(filepath)
	if readFileErr != nil {
		log.Panicf("Reading file error: %v", readFileErr)
	}

	document, htmlParseErr := html.Parse(strings.NewReader(string(dat)))
	if htmlParseErr != nil {
		log.Panicf("HTML parse error: %v", htmlParseErr)
	}

	linksContainer := ulSwitcher(document)
	links := getLinkNodes(linksContainer)

	removeFileErr := removeFile(filepath)
	if removeFileErr != nil {
		log.Panicf("Removing file error: %v", removeFileErr)
	}

	return links
}

func removeFile(filepath string) error {
	return os.Remove(filepath)
}

/* Check if node attribute(key) contains substring(value) */
func attrValueContains(node *html.Node, key string, value string) bool {
	for _, attr := range node.Attr {
		if attr.Key == key && strings.Contains(attr.Val, value) {
			return true
		}
	}

	return false
}

/* Returns node attribute(key) value */
func getAttrValue(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}

/* Get root div with links
<div id="tabs">...</div>
*/
func tabs(document *html.Node) *html.Node {
	var find func(node *html.Node)
	var div *html.Node

	find = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "div" && len(node.Attr) != 0 {

			if attrValueContains(node, "id", "tabs") {
				div = node
				return
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			find(child)
		}
	}
	find(document)

	return div
}

/* Get ul container with links
<ul id="tab-content">...</ul>
*/
func ulSwitcher(node *html.Node) *html.Node {
	var find func(node *html.Node)
	var ul *html.Node

	find = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "ul" && len(node.Attr) != 0 {

			if attrValueContains(node, "id", "tab-content") {
				ul = node
				return
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			find(child)
		}
	}
	find(node)

	return ul
}

/*  */
func getLinkNodes(ul *html.Node) map[int][]string {
	/* Find link names */
	var instituteLinks []string
	var find func(node *html.Node)
	find = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" && len(node.Attr) != 0 {
			if attrValueContains(node, "href", "https://webservices.mirea.ru/upload/") {
				instituteLinks = append(instituteLinks, getAttrValue(node, "href"))
			}

			return
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			find(child)
		}
	}

	allLinks := make(map[int][]string, 5)
	/* Find in each <li> tag */
	for index, child := 0, ul.FirstChild; child != nil; child = child.NextSibling {
		instituteLinks = []string{}
		if child.Type == html.ElementNode && child.Data == "li" {
			find(child)
			allLinks[index] = instituteLinks
			index++
		}
	}

	return allLinks
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
