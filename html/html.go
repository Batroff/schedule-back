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

// TODO: Create methods & encapsulate struct
type NodeParams struct {
	NodeName   string
	NodeType   html.NodeType
	Attributes []html.Attribute
}

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

	ulParams := &NodeParams{
		NodeName:   "ul",
		NodeType:   html.ElementNode,
		Attributes: []html.Attribute{{Key: "id", Val: "tab-content"}},
	}
	linksContainer := findNode(document, ulParams, attrEquals)

	links := getLinkNodes(linksContainer)
	removeFileErr := os.Remove(filepath)
	if removeFileErr != nil {
		log.Panicf("Removing file error: %v", removeFileErr)
	}

	return links
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

/* Compares *html.Node Attributes array and input Attribute array.
>> compareFunc should take 2 arguments of html.Attribute
*/
func attrCheck(node *html.Node, attributes []html.Attribute, compareFunc func(attr1 html.Attribute, attr2 html.Attribute) bool) bool {
	attributesMatch := make([]bool, len(attributes))

	for index, attr := range attributes {
		for _, nodeAttr := range node.Attr {
			if attributesMatch[index] {
				break
			}
			attributesMatch[index] = compareFunc(nodeAttr, attr)
		}
	}

	match := true
	for _, val := range attributesMatch {
		if val == false {
			match = false
		}
	}

	return match
}

/* Returns true if attr1 equals attr2 */
func attrEquals(attr1, attr2 html.Attribute) bool {
	return attr1 == attr2
}

/* Returns true if attr1 values contain attr2 values */
func attrContains(attr1, attr2 html.Attribute) bool {
	if attr1.Key != attr2.Key || !strings.Contains(attr1.Val, attr2.Val) {
		return false
	}

	return true
}

/*  */
func findNode(root *html.Node, params *NodeParams, compareFunc func(attr1 html.Attribute, attr2 html.Attribute) bool) *html.Node {
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

/* TODO: Write selector for multiple nodes
func findNodes(root *html.Node, params *NodeParams, compareFunc func(attr1 html.Attribute, attr2 html.Attribute) bool) []*html.Node {

}
*/

/* TODO: Rewrite function */
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
