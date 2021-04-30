package models

// Lesson TODO: add annotations
type Lesson struct {
	Subject          string `json:"subject" bson:"subject"`                   //название предмета
	TypeOfLesson     string `json:"typeOfLesson" bson:"typeOfLesson"`         //тип занятия
	TeacherName      string `json:"teacherName" bson:"teacherName"`           //фио преподавателя
	Cabinet          string `json:"cabinet" bson:"cabinet"`                   //кабинет
	NumberLesson     int    `json:"numberLesson" bson:"numberLesson"`         //номер пары
	DayOfWeek        string `json:"dayOfWeek" bson:"dayOfWeek,omitempty"`     //день недели
	OccurrenceLesson []bool `json:"occurrenceLesson" bson:"occurrenceLesson"` //номера недель в которых присутствует эта пара
	Exists           bool   `json:"exists,omitempty" bson:"exists,omitempty"` //для пустых пар??
	SubGroup         int    `json:"subGroup" bson:"subGroup,omitempty"`       // номер подгруппы
}
