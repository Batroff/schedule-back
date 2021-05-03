package models

type ResponseGroupList struct {
	ErrorMsg  string     `json:"ErrorMsg,omitempty"`
	GroupList *GroupList `json:"GroupList"`
}

type ResponseGroup struct {
	ErrorMsg string `json:"ErrorMsg,omitempty"`
	Group    *Group `json:"Group,omitempty"`
}
