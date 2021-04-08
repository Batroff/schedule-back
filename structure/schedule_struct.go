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

// TODO: add annotations
type Lesson struct {
	Subject      string //название предмета
	TypeOfLesson string //тип занятия
	TeacherName  string //фио преподавателя
	Cabinet      string //кабинет
	NumberLesson int    //номер пары
	DayOfWeek    string //день недели
	//occurrenceLesson []int//номера недель в которых присутствует эта пара
	OccurrenceLesson []bool //номера недель в которых присутствует эта пара
	Exists           bool   `json:"exists,omitempty" bson:"exists,omitempty"` //для пустых пар??
	SubGroup         int    // номер подгруппы
}

type Day struct {
	Lessons []Lesson `json:"lessons" bson:"lessons"`
}

type Week struct {
	Days []Day `json:"days" bson:"days"`
}

type Group struct {
	Weeks    []Week `json:"weeks" bson:"weeks"`
	Name     string `json:"name" bson:"name"`
	SubGroup int    `json:"subgroup,omitempty" bson:"subgroup,omitempty"` // номер подгруппы
}

func NewGroup() Group {
	var g Group
	g.Weeks = make([]Week, 17)
	for i := range g.Weeks {
		g.Weeks[i] = NewWeek()
	}
	return g
}

func NewWeek() Week {
	var w Week
	w.Days = make([]Day, 6)
	for i := range w.Days {
		w.Days[i] = NewDay()
	}
	return w
}

func NewDay() Day {
	var d Day
	d.Lessons = make([]Lesson, 9)
	for i := range d.Lessons {
		d.Lessons[i] = NewLesson()
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
				g.Weeks[i2].Days[weeksMap[lesson.DayOfWeek]].Lessons[lesson.NumberLesson] = lesson
			}
		}
	}
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
