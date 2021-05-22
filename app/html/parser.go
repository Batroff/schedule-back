package html

import (
	"github.com/batroff/schedule-back/app"
	"github.com/batroff/schedule-back/models"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

/* Get map of links
0 - Бакалавр/Специалитет;
1 - Магистр;
2 - Аспирант;
3 - Колледж;
4 - Экстерн;
*/

// TODO: add catching exceptions
func GetExcelLinks() (map[int][]string, error) {
	const filepath string = "schedule.html"

	downloadErr := app.GetFile(filepath, "https://www.mirea.ru/schedule/")
	if downloadErr != nil {
		return nil, downloadErr
	}

	dat, readFileErr := ioutil.ReadFile(filepath)
	if readFileErr != nil {
		log.Printf("Reading file error: %s\n", readFileErr)
		return nil, readFileErr
	}

	document, htmlParseErr := html.Parse(strings.NewReader(string(dat)))
	if htmlParseErr != nil {
		log.Printf("HTML excel error: %s\n", htmlParseErr)
		return nil, htmlParseErr
	}

	ulParams := &models.NodeParams{
		NodeName:   "ul",
		NodeType:   html.ElementNode,
		Attributes: []html.Attribute{{Key: "id", Val: "tab-content"}},
	}
	linksContainer := findNode(document, ulParams, attrEquals)

	links := getLinkNodes(linksContainer)
	removeFileErr := os.Remove(filepath)
	if removeFileErr != nil {
		log.Printf("Removing file error: %s\n", removeFileErr)
		return nil, removeFileErr
	}

	return links, nil
}

/*  */
func findNode(root *html.Node, params *models.NodeParams, compareFunc func(attr1 html.Attribute, attr2 html.Attribute) bool) *html.Node {
	var find func(node *html.Node)
	var resultNode *html.Node

	isFound := false
	find = func(node *html.Node) {
		if isFound {
			return
		}

		if node.Type != params.NodeType || node.Data != params.NodeName || len(node.Attr) < len(params.Attributes) {
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				find(child)
			}
		} else if node.Type == params.NodeType && node.Data == params.NodeName {
			if len(params.Attributes) == 0 {
				resultNode = node
				isFound = true
				return
			} else if attrCheck(node, params.Attributes, compareFunc) {
				resultNode = node
				isFound = true
				return
			}
		}
	}

	find(root)

	return resultNode
}

func getLinkNodes(ul *html.Node) map[int][]string {
	/* Find link names */
	var instituteLinks []string
	var find func(node *html.Node)
	find = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" && len(node.Attr) != 0 {
			if attrValueContains(node, "href", regexp.MustCompile("зач|экз|сессия|гиа")) {
				return
			}

			if attrValueContains(node, "href", regexp.MustCompile("https://webservices.mirea.ru/upload/")) {
				instituteLinks = append(instituteLinks, getAttrValue(node, "href"))
			}

			return
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			find(child)
		}
	}

	allLinks := make(map[int][]string, 5)
	/* Find in each <li> tag
	* Бакалавриат, магистратура, ...
	 */
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
