package html

import (
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

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

/* Check if node attribute(key) contains substring(value) */
func attrValueContains(node *html.Node, key string, regexp *regexp.Regexp) bool {
	for _, attr := range node.Attr {
		if attr.Key == key && regexp.MatchString(strings.ToLower(attr.Val)) {
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
