package models

type ResponseGroupList struct {
	ErrorMsg  string          `json:"errorMsg,omitempty"`
	GroupList map[string]bool `json:"groupList"`
}

type ResponseGroup struct {
	ErrorMsg string `json:"errorMsg,omitempty"`
	Group    *Group `json:"group"`
}

type ResponseHash struct {
	ErrorMsg string   `json:"errorMsg,omitempty"`
	Hash     []string `json:"hash"`
}
