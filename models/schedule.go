package models

import (
	"strings"
)

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
}

func RemoveElementLesson(a *[]Lesson, i int) {
	//*a = append((*a)[:i], (*a)[i+1:]...)
	(*a)[i] = (*a)[len(*a)-1]
	(*a)[len(*a)-1] = Lesson{}
	*a = (*a)[:len(*a)-1]
}

var GroupMap = make(map[string]bool)

type GroupList struct {
	Map map[string]bool `json:"map" bson:"map"`
}

func CreateGroupList() GroupList {
	var result GroupList
	result.Map = GroupMap
	return result
}
