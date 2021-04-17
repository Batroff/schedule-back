package structure

import (
	"strings"
)

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

type Group struct {
	Days     map[string][]Lesson `json:"days" bson:"days,omitempty"`
	Name     string              `json:"name" bson:"name"`
	SubGroup int                 `json:"subgroup,omitempty" bson:"subgroup,omitempty"` // номер подгруппы
}

func NewGroup() (g Group) {
	g.SubGroup = 0
	g.Name = ""
	day := []Lesson{NewLesson()}
	g.Days = map[string][]Lesson{
		"ПОНЕДЕЛЬНИК": day,
		"ВТОРНИК":     day,
		"СРЕДА":       day,
		"ЧЕТВЕРГ":     day,
		"ПЯТНИЦА":     day,
		"СУББОТА":     day,
	}
	return g
}

func NewLesson() Lesson {
	var l Lesson
	l.OccurrenceLesson = make([]bool, 17)
	return l
}

func (l Lesson) FillInWeeks(week string) {
	if strings.Contains(week, "II") {
		for i := 1; i < len(l.OccurrenceLesson)-1; i += 2 {
			l.OccurrenceLesson[i] = true
		}
	} else if strings.Contains(week, "I") {
		for i := 0; i < len(l.OccurrenceLesson)-1; i += 2 {
			l.OccurrenceLesson[i] = true
		}
	}
}

func (g Group) Clear() {
	for s := range g.Days {
		for i := 0; i < len(g.Days[s]); i++ {
			if !(g.Days[s])[i].Exists {
				temp := g.Days[s]
				RemoveElementLesson(&temp, i)
				g.Days[s] = temp
				i--
			}
		}
	}
	for s, lessons := range g.Days {
		for i := range lessons {
			g.Days[s][i].Exists = false
			g.Days[s][i].DayOfWeek = ""
			g.Days[s][i].SubGroup = 0
		}
	}
	//объединение уроков на нечётной и чёткной недели в 1 урок
	for s := range g.Days {
		for i := 0; i < len(g.Days[s])-1; i++ {
			for j := i + 1; j < len(g.Days[s]); j++ {
				if Combined(g.Days[s][i], g.Days[s][j]) {
					for i2 := range g.Days[s][i].OccurrenceLesson {
						if g.Days[s][i].OccurrenceLesson[i2] || g.Days[s][j].OccurrenceLesson[i2] {
							g.Days[s][i].OccurrenceLesson[i2] = true
						}
					}
					g.Days[s][j].SubGroup = -1
				}
			}
		}
		for i := 0; i < len(g.Days[s]); i++ {
			if g.Days[s][i].SubGroup == -1 {
				temp := g.Days[s]
				RemoveElementLesson(&temp, i)
				g.Days[s] = temp
				i--
			}
		}
	}
}

func RemoveElementLesson(a *[]Lesson, i int) {
	*a = append((*a)[:i], (*a)[i+1:]...)
	//(*a)[i] = (*a)[len(*a)-1]
	//(*a)[len(*a)-1] = Lesson{}
	//*a = (*a)[:len(*a)-1]
}

func Combined(lesson1, lesson2 Lesson) bool {
	if lesson1.Subject == lesson2.Subject {
		if lesson1.TeacherName == lesson2.TeacherName {
			if lesson1.TypeOfLesson == lesson2.TypeOfLesson {
				if lesson1.Cabinet == lesson2.Cabinet {
					if lesson1.NumberLesson == lesson2.NumberLesson {
						return true
					}
				}
			}
		}
	}
	return false
}
