package structure

import "strings"

var weeksMap = map[string]int{
	"ПОНЕДЕЛЬНИК": 0,
	"ВТОРНИК":     1,
	"СРЕДА":       2,
	"ЧЕТВЕРГ":     3,
	"ПЯТНИЦА":     4,
	"СУББОТА":     5,
}

type Lesson struct {
	Subject      string //название предмета
	TypeOfLesson string //тип занятия
	TeacherName  string //фио преподавателя
	Cabinet      string //кабинет
	NumberLesson int    //номер пары
	DayOfWeek    string //день недели
	//occurrenceLesson []int//номера недель в которых присутствует эта пара
	OccurrenceLesson []bool //номера недель в которых присутствует эта пара
	Exists           bool   //для пустых пар??
	//SubGroup         int    // номер подгруппы

}

type Day struct {
	lessons []Lesson
}

type Week struct {
	days []Day
}

type Group struct {
	weeks    []Week
	name     string
	subGroup int // номер подгруппы
}

func NewGroup() Group {
	var g Group
	g.weeks = make([]Week, 17)
	for i := range g.weeks {
		g.weeks[i] = NewWeek()
	}
	return g
}

func NewWeek() Week {
	var w Week
	w.days = make([]Day, 6)
	for i := range w.days {
		w.days[i] = NewDay()
	}
	return w
}

func NewDay() Day {
	var d Day
	d.lessons = make([]Lesson, 8)
	for i := range d.lessons {
		d.lessons[i] = NewLesson()
	}
	return d
}

func NewLesson() Lesson {
	var l Lesson
	l.OccurrenceLesson = make([]bool, 17)
	return l
}

func (g Group) AddLesson(lessons []Lesson) {
	for _, lesson := range lessons {
		for i2, b := range lesson.OccurrenceLesson {
			if b {
				g.weeks[i2].days[weeksMap[lesson.DayOfWeek]].lessons[lesson.NumberLesson] = lesson
			}
		}
	}
}
func (l Lesson) FillInWeeks(flag bool, week string) {
	if flag && strings.Contains(week, "II") {
		for i := 1; i < len(l.OccurrenceLesson)-1; i += 2 {
			l.OccurrenceLesson[i] = true
		}
	} else if flag && strings.Contains(week, "I") {
		for i := 0; i < len(l.OccurrenceLesson)-1; i += 2 {
			l.OccurrenceLesson[i] = true
		}
	} else if !flag && strings.Contains(week, "II") {
		for i := 1; i < len(l.OccurrenceLesson)-1; i += 2 {
			l.OccurrenceLesson[i] = false
		}
	} else if !flag && strings.Contains(week, "I") {
		for i := 0; i < len(l.OccurrenceLesson)-1; i += 2 {
			l.OccurrenceLesson[i] = false
		}
	}
}
