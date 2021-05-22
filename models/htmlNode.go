package models

import "golang.org/x/net/html"

// TODO: Create methods & encapsulate struct
type NodeParams struct {
	NodeName   string
	NodeType   html.NodeType
	Attributes []html.Attribute
}
