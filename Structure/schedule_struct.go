package Structure

var weeksMap = map[string]int{
	"ПОНЕДЕЛЬНИК": 0,
	"ВТОРНИК":     1,
	"СРЕДА":       2,
	"ЧЕТВЕРГ":     3,
	"ПЯТНИЦА":     4,
	"СУББОТА":     5,
}

type Lesson struct {
	subject      string //название предмета
	typeOfLesson string //тип занятия
	teacherName  string //фио преподавателя
	cabinet      string //кабинет
	numberLesson int    //номер пары
	dayOfWeek    string //день недели
	//occurrenceLesson []int//номера недель в которых присутствует эта пара
	occurrenceLesson []bool //номера недель в которых присутствует эта пара
	exists           bool   //для пустых пар??
	subGroup         int    // номер подгруппы
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
	subGroup bool // номер подгруппы
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
	l.occurrenceLesson = make([]bool, 17)
	return l
}

func (g Group) AddLesson(lessons []Lesson) {
	for _, lesson := range lessons {
		for i2, b := range lesson.occurrenceLesson {
			if b {
				g.weeks[i2].days[weeksMap[lesson.dayOfWeek]].lessons[lesson.numberLesson] = lesson
			}
		}
	}
}
