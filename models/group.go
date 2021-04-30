package models

type Group struct {
	Days     map[string][]Lesson `json:"days" bson:"days,omitempty"`
	Name     string              `json:"name" bson:"name"`
	SubGroup int                 `json:"subgroup,omitempty" bson:"subgroup,omitempty"` // номер подгруппы
}
