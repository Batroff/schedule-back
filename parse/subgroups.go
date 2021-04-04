package parse

import (
	"fmt"
	//"fmt"
	"regexp"
	. "schedule/structure"
	"strconv"
	"strings"
)

var Space = regexp.MustCompile(` {2,}`)

var Initials = regexp.MustCompile(`[А-Яа-я]* ([А-Я]\.){2}`)

var GroupRegexp = regexp.MustCompile(`\d *гр\.?`)

var Digit = regexp.MustCompile(`\d`)

//недели или нет предмет подгруппа
var SubgroupRegexp1 = regexp.MustCompile(`^(кр\.? *)?(( *\d{1,2},?\w?)*( *н?\.?,? *-* *))([А-Яа-я]+ *-*,*\.* *)+(\(?\d\s*п/гр?\)?|\d ?гр ?$|\(?\d ?подгр\.? ?\)?$)`)
var SubgroupRegexp1Subgroup = regexp.MustCompile(`[,. ]*\(?\d *(п/г|гр|подгр)\)?\.?`)

//гр. недели предмет
var SubgroupRegexp2 = regexp.MustCompile(`^\dгр\. ?(\d{1,2},?)+ ?н\.? *`)
var SubgroupRegexp2Subgroup = regexp.MustCompile(`\dгр\.? *`)

var SubgroupRegexp3 = regexp.MustCompile(`((\d{1,2},?\w?)*( *н?\.? *-* *))([А-Яа-я]+ *-*,*\.* *)+(\(?\d{1,2}-\d{1,2} *нед\.? ?\)? *)?(\(*(подгруппа|подгр) ?.? ?\d\)* ?$)`)
var SubgroupRegexp3Subgroup = regexp.MustCompile(`\(?подгр(\.? *|уппа) *\d\)?`)

var SubgroupRegexp4 = regexp.MustCompile(`(\dгр\.? *=? *(\d{1,2},?\.?)*н?\.?;?,? *)+=?[А-Яа-я ,-]*`)

var CrutchRegexp1 = regexp.MustCompile(`,? *\d ?гр ?/ ?\d ?гр`)

var CrutchRegexp2 = regexp.MustCompile(`(([А-Яа-я] ?)*\(\d ?подгр\.?\)/?){2}`)

var CrutchRegexp3 = regexp.MustCompile(`[А-Яа-я,?. *]*\((\d,? ?-?)*нед\./(\d,? ?-?)*нед\. *- *подгр\.?\d\)`)
var CrutchRegexp3Lite = regexp.MustCompile(`[А-Яа-я,?. *]*\((\d,? ?-?)*нед\. *- *подгр\.?\d\)`)
var CrutchRegexp3Mini = regexp.MustCompile(` ?-+ ?подгр.?\d`)

var CrutchRegexp4 = regexp.MustCompile(`(\dгр\. *=? *(\d{1,2},?\.?)*н?\.?;?,? *)+[А-Яа-я ,-]*; *(\d{1,2},?)* *н\.?[А-Яа-я ,.]*`)
var CrutchRegexp4Normal = regexp.MustCompile(`(\d{1,2},?)+ *н\.?[А-Яа-я ,.]*$`) // нормальная часть 4-ого костыльного случая
var CrutchRegexp4Subgroup = regexp.MustCompile(`^(\dгр\. *=? *(\d{1,2},?\.?)*н?\.?;?,? *)+[А-Яа-я ,-]*`)

var CrutchRegexp5 = regexp.MustCompile(`(((\dгр\.? ?= ?(\d{1,2},?)*н\.);?,? ?){2}=?[А-Яа-я ]*;? ?){2}`)

var CrutchRegexp6 = regexp.MustCompile(`^(\dгр\.? *=? *(\d{1,2},?\.?)*н?\.?;?,? *){2}=?[А-Яа-я ,-]*`)
var CrutchRegexp6Subgroup = regexp.MustCompile(`\dгр\.? *=? *(\d{1,2},? ?)*н?[.= ,;]*`)
var CrutchRegexp6Mini = regexp.MustCompile(`\dгр\.? *=? *`)

//недели - гр1 недели - гр2 предмет
var CrutchRegexp7 = regexp.MustCompile(`((\d{1,2},? ?)н ?- ?\d ?гр,? ?)+ *([А-Яа-я]* *)*`)
var CrutchRegexp7Subgroup = regexp.MustCompile(`(\d{1,2},?)+ *н? *- *\d *гр,? *`)
var CrutchRegexp7Mini = regexp.MustCompile(`\d *гр`)

var SubgroupNumber = 0

var GlobalWeek string
var GlobalDayOfWeek string
var GlobalNumberLesson string

//несколько уроков в 1 дне надо раскидать по строкам и если одинаковые предметы почему они раскинуты(тип работы/преподы)
func SubGroupParse(subject, typeOfLesson, teacherName, cabinet, dayOfWeek, numberLesson, week string) (resultLessons []Lesson) {
	var lessons []Lesson
	GlobalWeek = week
	GlobalDayOfWeek = dayOfWeek
	GlobalNumberLesson = numberLesson
	if strings.Contains(subject, "\n") { // если в строчке с предметом более 1 строки
		lessons = LessonToLessons(subject, typeOfLesson, teacherName, cabinet)
		SubgroupLessonsSort(&lessons)
		for _, lesson := range lessons {
			//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			//fmt.Println("Предметы:")
			//fmt.Println(lesson.Subject)
			//fmt.Println("Тип занятий:")
			//fmt.Println(lesson.TypeOfLesson)
			//fmt.Println("ФИО:")
			//fmt.Println(lesson.TeacherName)
			//fmt.Println("Кабинет:")
			//fmt.Println(lesson.Cabinet)
			//SubgroupLessonParse(&[]Lesson{lesson})
			if lesson.Exists {

			}
		}
	} else { // в строке нет энтеров
		lesson := NewLesson()
		lesson.Subject = subject
		lesson.TypeOfLesson = typeOfLesson
		lesson.TeacherName = teacherName
		lesson.Cabinet = cabinet
		SubgroupLessonsSort(&([]Lesson{lesson}))
		//SubgroupLessonParse(&lesson)
	}
	for i, _ := range lessons {
		SubgroupLessonParse(&lessons[i])
		lessons[i].NumberLesson, _ = strconv.Atoi(numberLesson)
		lessons[i].DayOfWeek = dayOfWeek
		//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		//fmt.Println("Предметы:")
		//fmt.Println((lessons)[i].Subject)
		//fmt.Println("Тип занятий")
		//fmt.Println((lessons)[i].TypeOfLesson)
		//fmt.Println("ФИО")
		//fmt.Println((lessons)[i].TeacherName)
		//fmt.Println("Кабинет")
		//fmt.Println((lessons)[i].Cabinet)
		fmt.Println(lessons[i])

	}
	return lessons
}

func SubgroupLessonParse(lesson *Lesson) {
	if SubgroupNumber == lesson.SubGroup || lesson.SubGroup == 0 {
		lesson.Exists = true
	} else {
		lesson.Exists = false
	}
}

func SubgroupLessonsSort(lessons *[]Lesson) {
	for getIdLesson(lessons) != -1 {
		i2 := getIdLesson(lessons)
		lesson1, lesson2 := Fix((*lessons)[i2])
		*lessons = append(append((*lessons)[:i2], (*lessons)[i2+1:]...), lesson1, lesson2)
	}
	for i, lesson := range *lessons {
		if !SubgroupRegexp.MatchString(lesson.Subject) {
			(*lessons)[i].SubGroup = 0
			(*lessons)[i].FillInWeeks(GlobalWeek)
			// lesson надо отправить в человеческий парс
		} else if CrutchRegexp3.MatchString(lesson.Subject) || CrutchRegexp3Lite.MatchString(lesson.Subject) { //парс с подгруппами
			temp := " " + strings.ReplaceAll(strings.ReplaceAll(CrutchRegexp3Mini.FindString(lesson.Subject), "-", ""), " ", "")
			(*lessons)[i].Subject = strings.ReplaceAll((*lessons)[i].Subject, CrutchRegexp3Mini.FindString(lesson.Subject), "") + temp
			digit, _ := strconv.Atoi(Digit.FindString(CrutchRegexp3Mini.FindString(lesson.Subject)))
			(*lessons)[i] = DefaultParse((*lessons)[i].Subject, (*lessons)[i].TypeOfLesson, (*lessons)[i].TeacherName, (*lessons)[i].Cabinet, GlobalDayOfWeek, GlobalNumberLesson, GlobalWeek)[0]
			(*lessons)[i].SubGroup = digit
		} else if SubgroupRegexp1.MatchString(lesson.Subject) && !strings.Contains(lesson.Subject, ")/И") {
			// ~15000 строчек
			temp := SubgroupRegexp1Subgroup.FindString(lesson.Subject)
			(*lessons)[i].Subject = strings.ReplaceAll(lesson.Subject, temp, "") // Строка для норм парса
			digit, _ := strconv.Atoi(Digit.FindString(temp))                     // номер подгруппы
			(*lessons)[i] = DefaultParse((*lessons)[i].Subject, (*lessons)[i].TypeOfLesson, (*lessons)[i].TeacherName, (*lessons)[i].Cabinet, GlobalDayOfWeek, GlobalNumberLesson, GlobalWeek)[0]
			(*lessons)[i].SubGroup = digit
			//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			//fmt.Println("Предметы:")
			//fmt.Println((*lessons)[i].Subject)
			//fmt.Println("Тип занятий")
			//fmt.Println((*lessons)[i].TypeOfLesson)
			//fmt.Println("ФИО")
			//fmt.Println((*lessons)[i].TeacherName)
			//fmt.Println("Кабинет")
			//fmt.Println((*lessons)[i].Cabinet)
		} else if SubgroupRegexp2.MatchString(lesson.Subject) {
			// ~120 строчек
			temp := SubgroupRegexp2Subgroup.FindString(lesson.Subject)
			(*lessons)[i].Subject = strings.ReplaceAll(lesson.Subject, temp, "") // Строка для норм парса
			digit, _ := strconv.Atoi(Digit.FindString(temp))                     // номер подгруппы
			(*lessons)[i] = DefaultParse((*lessons)[i].Subject, (*lessons)[i].TypeOfLesson, (*lessons)[i].TeacherName, (*lessons)[i].Cabinet, GlobalDayOfWeek, GlobalNumberLesson, GlobalWeek)[0]
			(*lessons)[i].SubGroup = digit
			//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			//fmt.Println("Предметы:")
			//fmt.Println((*lessons)[i].Subject)
			//fmt.Println("Тип занятий")
			//fmt.Println((*lessons)[i].TypeOfLesson)
			//fmt.Println("ФИО")
			//fmt.Println((*lessons)[i].TeacherName)
			//fmt.Println("Кабинет")
			//fmt.Println((*lessons)[i].Cabinet)
		} else if SubgroupRegexp3.MatchString(lesson.Subject) {
			// ~76 строчек
			temp := SubgroupRegexp3Subgroup.FindString(lesson.Subject)
			(*lessons)[i].Subject = strings.ReplaceAll(lesson.Subject, temp, "") // Строка для норм парса
			digit, _ := strconv.Atoi(Digit.FindString(temp))                     // номер подгруппы
			(*lessons)[i] = DefaultParse((*lessons)[i].Subject, (*lessons)[i].TypeOfLesson, (*lessons)[i].TeacherName, (*lessons)[i].Cabinet, GlobalDayOfWeek, GlobalNumberLesson, GlobalWeek)[0]
			(*lessons)[i].SubGroup = digit
			//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			//fmt.Println("Предметы:")
			//fmt.Println((*lessons)[i].Subject)
			//fmt.Println("Тип занятий")
			//fmt.Println((*lessons)[i].TypeOfLesson)
			//fmt.Println("ФИО")
			//fmt.Println((*lessons)[i].TeacherName)
			//fmt.Println("Кабинет")
			//fmt.Println((*lessons)[i].Cabinet)
		} else {
			(*lessons)[i].SubGroup = 0
			(*lessons)[i].FillInWeeks(GlobalWeek)
			//как есть так и есть нормально не запарсить и возможно и не стоит парсить
			//fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			//fmt.Println("Предметы:")
			//fmt.Println(lesson.Subject)
			//fmt.Println("Тип занятий")
			//fmt.Println(lesson.TypeOfLesson)
			//fmt.Println("ФИО")
			//fmt.Println(lesson.TeacherName)
			//fmt.Println("Кабинет")
			//fmt.Println(lesson.Cabinet)
		}
	}
}

//номер костыльного случая в массиве
func getIdLesson(lessons *[]Lesson) (array int) {
	for i, lesson := range *lessons {
		if strings.Contains(lesson.Subject, "1+2 гр") {
			(*lessons)[i].Subject = strings.ReplaceAll((*lessons)[i].Subject, "1+2 гр", "1 гр/2 гр")
		}
	}
	for i, lesson := range *lessons {
		if CrutchRegexp1.MatchString(lesson.Subject) || CrutchRegexp2.MatchString(lesson.Subject) || CrutchRegexp4.MatchString(lesson.Subject) ||
			CrutchRegexp5.MatchString(lesson.Subject) || CrutchRegexp6.MatchString(lesson.Subject) || CrutchRegexp7.MatchString(lesson.Subject) { //
			return i
		}
	}
	return -1
}

//такой убогий метод для слэшей
func Fix(lesson Lesson) (lesson1, lesson2 Lesson) {
	lesson1 = NewLesson()
	lesson2 = NewLesson()
	if CrutchRegexp1.MatchString(lesson.Subject) {
		subgroupStr := CrutchRegexp1.FindString(lesson.Subject)
		str := strings.ReplaceAll(lesson.Subject, subgroupStr, "")
		array := strings.Split(subgroupStr, "/")
		lesson1.Subject = str + array[0]
		lesson2.Subject = str + array[1]
		arrayTypes := strings.Split(lesson.TypeOfLesson, "/")
		arrayTeachers := strings.Split(lesson.TeacherName, "/")
		arrayCabinets := strings.Split(lesson.Cabinet, "/")
		if len(arrayTypes) != 1 {
			lesson1.TypeOfLesson = arrayTypes[0]
			lesson2.TypeOfLesson = arrayTypes[1]
		} else {
			lesson1.TypeOfLesson = arrayTypes[0]
			lesson2.TypeOfLesson = arrayTypes[0]
		}
		if len(arrayTeachers) != 1 {
			lesson1.TeacherName = arrayTeachers[0]
			lesson2.TeacherName = arrayTeachers[1]
		} else {
			lesson1.TeacherName = arrayTeachers[0]
			lesson2.TeacherName = arrayTeachers[0]
		}
		if len(arrayCabinets) != 1 {
			lesson1.Cabinet = arrayCabinets[0]
			lesson2.Cabinet = arrayCabinets[1]
		} else {
			lesson1.Cabinet = arrayCabinets[0]
			lesson2.Cabinet = arrayCabinets[0]
		}
	} else if CrutchRegexp2.MatchString(lesson.Subject) {
		arraySubject := strings.Split(lesson.Subject, "/")
		lesson1.Subject = arraySubject[0]
		lesson2.Subject = arraySubject[1]
		lesson1.Cabinet = lesson.Cabinet
		lesson2.Cabinet = lesson.Cabinet
		lesson1.TypeOfLesson = lesson.TypeOfLesson
		lesson2.TypeOfLesson = lesson.TypeOfLesson
		lesson1.TeacherName = lesson.TeacherName
		lesson2.TeacherName = lesson.TeacherName
	} else if CrutchRegexp4.MatchString(lesson.Subject) {
		lesson2.Subject = CrutchRegexp4Normal.FindString(lesson.Subject)
		lesson1.Subject = CrutchRegexp4Subgroup.FindString(lesson.Subject)
		lesson1.TypeOfLesson = lesson.TypeOfLesson
		lesson2.TypeOfLesson = lesson.TypeOfLesson
		str := strings.Join(Initials.FindAllString(lesson.TeacherName, -1), " ? ")
		lesson1.TeacherName = str
		lesson2.TeacherName = str
		str = strings.ReplaceAll(lesson.Cabinet, Space.FindString(lesson.Cabinet), " ? ")
		lesson1.Cabinet = str
		lesson2.Cabinet = str
	} else if CrutchRegexp5.MatchString(lesson.Subject) {
		arraySubject := SubgroupRegexp4.FindAllString(lesson.Subject, -1)
		lesson1.Subject = arraySubject[0]
		lesson2.Subject = arraySubject[1]
		lesson1.TypeOfLesson = lesson.TypeOfLesson
		lesson2.TypeOfLesson = lesson.TypeOfLesson
		arrayTeachers := Initials.FindAllString(lesson.TeacherName, -1)
		lesson1.TeacherName = arrayTeachers[0]
		lesson2.TeacherName = arrayTeachers[1]
		arrayCabinets := strings.Split(lesson.Cabinet, ", ")
		lesson1.Cabinet = arrayCabinets[0]
		lesson2.Cabinet = arrayCabinets[1]
	} else if CrutchRegexp6.MatchString(lesson.Subject) {
		subgroups := CrutchRegexp6Subgroup.FindAllString(lesson.Subject, -1)
		subject := lesson.Subject
		for _, subgroup := range subgroups {
			subject = strings.ReplaceAll(subject, subgroup, "")
		}
		subgroupsWithoutWeeks := CrutchRegexp6Mini.FindAllString(lesson.Subject, -1)
		lesson1.Subject = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(subgroups[0], subgroupsWithoutWeeks[0], ""), ";", "")+
			subject+" "+strings.ReplaceAll(strings.ReplaceAll(GroupRegexp.FindString(subgroupsWithoutWeeks[0]), "гр", " подгр"), ".", ""), "=", "")
		lesson2.Subject = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(subgroups[1], subgroupsWithoutWeeks[1], ""), ";", "")+
			subject+" "+strings.ReplaceAll(strings.ReplaceAll(GroupRegexp.FindString(subgroupsWithoutWeeks[1]), "гр", " подгр"), ".", ""), "=", "")
		lesson1.Cabinet = lesson.Cabinet
		lesson2.Cabinet = lesson.Cabinet
		lesson1.TypeOfLesson = lesson.TypeOfLesson
		lesson2.TypeOfLesson = lesson.TypeOfLesson
		lesson1.TeacherName = lesson.TeacherName
		lesson2.TeacherName = lesson.TeacherName
	} else if CrutchRegexp7.MatchString(lesson.Subject) {
		subgroups := CrutchRegexp7Subgroup.FindAllString(lesson.Subject, -1)
		subject := lesson.Subject
		for _, subgroup := range subgroups {
			subject = strings.ReplaceAll(subject, subgroup, "")
		}
		subgroupsWithoutWeeks := CrutchRegexp7Mini.FindAllString(lesson.Subject, -1)
		lesson1.Subject = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(subgroups[0], subgroupsWithoutWeeks[0], ""), ";", "")+
			subject+" "+strings.ReplaceAll(strings.ReplaceAll(GroupRegexp.FindString(subgroupsWithoutWeeks[0]), "гр", " подгр"), ".", ""), "-", ""), "  ", " ")
		lesson2.Subject = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(subgroups[1], subgroupsWithoutWeeks[1], ""), ";", "")+
			subject+" "+strings.ReplaceAll(strings.ReplaceAll(GroupRegexp.FindString(subgroupsWithoutWeeks[1]), "гр", " подгр"), ".", ""), "-", ""), "  ", " ")
		lesson1.Cabinet = lesson.Cabinet
		lesson2.Cabinet = lesson.Cabinet
		lesson1.TypeOfLesson = lesson.TypeOfLesson
		lesson2.TypeOfLesson = lesson.TypeOfLesson
		lesson1.TeacherName = lesson.TeacherName
		lesson2.TeacherName = lesson.TeacherName
	}
	return lesson1, lesson2
}

func LessonToLessons(subject, typeOfLesson, teacherName, cabinet string) []Lesson {
	var lessons []Lesson
	if strings.Contains(cabinet, "В-78*\n") || strings.Contains(cabinet, "В-86*\n") || strings.Contains(cabinet, "МП-1*\n") {
		strings.ReplaceAll(cabinet, "В-78*\n", "В-78* ")
		strings.ReplaceAll(cabinet, "В-86*\n", "В-86* ")
		strings.ReplaceAll(cabinet, "МП-1*\n", "МП-1* ")
	}
	subjects := strings.Split(subject, "\n")
	typesOfLessons := strings.Split(typeOfLesson, "\n")
	teachersNames := strings.Split(teacherName, "\n")
	cabinets := strings.Split(cabinet, "\n")
	RemoveLastEmptyElement(&typesOfLessons)
	RemoveLastEmptyElement(&teachersNames)
	RemoveLastEmptyElement(&cabinets)
	if Contains(subjects, "…………………") { // =)
		for i, s := range subjects {
			if s == "…………………" {
				RemoveElement(&subjects, i)
			}
		}
	}
	if Contains(subjects, "") {
		for i, s := range subjects {
			if s == "" {
				RemoveElement(&subjects, i)
			}
		}

		if Contains(typesOfLessons, "") {
			for i, s := range typesOfLessons {
				if s == "" {
					RemoveElement(&typesOfLessons, i)
				}
			}
		}
		if Contains(teachersNames, "") {
			for i, s := range teachersNames {
				if s == "" {
					RemoveElement(&teachersNames, i)
				}
			}
		}
		if Contains(cabinets, "") {
			for i, s := range cabinets {
				if s == "" {
					RemoveElement(&cabinets, i)
				}
			}
		}
	}
	FixSameSubjectParameters(&subjects, &typesOfLessons, &teachersNames, &cabinets)

	parameterConversion(&subjects, &typesOfLessons)
	parameterConversion(&subjects, &teachersNames)
	parameterConversion(&subjects, &cabinets)

	collection := [][]string{
		subjects, typesOfLessons, teachersNames, cabinets,
	}

	length := len(collection[0])

	for i := 0; i < length; i++ {
		if collection[0][i] == "" && collection[1][i] == "" && collection[2][i] == "" && collection[3][i] == "" {
			RemoveElement(&collection[0], i)
			RemoveElement(&collection[1], i)
			RemoveElement(&collection[2], i)
			RemoveElement(&collection[3], i)
			length--
			continue
		}
		someLesson := NewLesson()
		someLesson.Subject = collection[0][i]
		someLesson.TypeOfLesson = collection[1][i]
		someLesson.TeacherName = collection[2][i]
		someLesson.Cabinet = collection[3][i]
		lessons = append(lessons, someLesson) // массив с уроками "предмет с/без п/г" "тип" "фио" "кабинет"
	}
	return lessons
}

//Функция для случаев где предметов меньше чем других параметров и дублирует все лишние параметры в одной строке разделяя их вопросом
func parameterConversion(subjects, array *[]string) {
	if len(*subjects) < len(*array) {
		sum := strings.Join(*array, " ? ")
		*array = make([]string, len(*subjects))
		for i := range *array {
			(*array)[i] = sum
		}
	}
}

func Max(number ...int) int {
	max := 0
	for _, num := range number {
		if max < num {
			max = num
		}
	}
	return max
}

//Удаляет последний пустой элемент
func RemoveLastEmptyElement(array *[]string) {
	if (*array)[len(*array)-1] == "" {
		RemoveElement(array, len(*array)-1)
	}
}

//Удаляет элемент из среза строк по индексу
func RemoveElement(a *[]string, i int) {
	*a = append((*a)[:i], (*a)[i+1:]...)
}

var regexpForFixSameSubjectParametersFunc = regexp.MustCompile("[^\\d()][А-я -]+")

//Если в строках содержатся одинаковые предметы, то подтянуть из нужной строки вид занятия или кабинет или препода или задублировать последним найденным
func FixSameSubjectParameters(subjects, typeOfLessons, teachersNames, cabinets *[]string) {
	for len(*typeOfLessons) < len(*subjects) {
		*typeOfLessons = append(*typeOfLessons, "")
	}

	for len(*teachersNames) < len(*subjects) {
		*teachersNames = append(*teachersNames, "")
	}

	for len(*cabinets) < len(*subjects) {
		*cabinets = append(*cabinets, "")
	}
	for i, subject := range *subjects {
		for i2, s := range (*subjects)[i+1:] { // i2 относителен поэтому i + i2 + 1

			index := i2 + i + 1
			str1 := strings.ReplaceAll(strings.ToLower(strings.ReplaceAll(LongestString(regexpForFixSameSubjectParametersFunc.FindAllString(subject, -1)), " ", "")), "н", "")
			str2 := strings.ReplaceAll(strings.ToLower(strings.ReplaceAll(LongestString(regexpForFixSameSubjectParametersFunc.FindAllString(s, -1)), " ", "")), "н", "")
			if str1 == str2 { // если 2 предмета одинаковые
				if (*typeOfLessons)[i] == "" {
					(*typeOfLessons)[i] = (*typeOfLessons)[index]
				}
				if (*typeOfLessons)[index] == "" {
					(*typeOfLessons)[index] = (*typeOfLessons)[i]
				}

				if (*teachersNames)[i] == "" {
					(*teachersNames)[i] = (*teachersNames)[index]
				}
				if (*teachersNames)[index] == "" {
					(*teachersNames)[index] = (*teachersNames)[i]
				}

				if (*cabinets)[i] == "" {
					(*cabinets)[i] = (*cabinets)[index]
				}
				if (*cabinets)[index] == "" {
					(*cabinets)[index] = (*cabinets)[i]
				}
			}
		}
	}
	RepeatFunc(typeOfLessons)
	RepeatFunc(teachersNames)
	RepeatFunc(cabinets)
}

//Заполнение пустых элементов (дублирование)
func RepeatFunc(array *[]string) {
	flag := false
	for _, s := range *array {
		if s != "" {
			flag = true
		}
	}
	if flag {
		for i := 1; i < len(*array); i++ {
			if (*array)[i] == "" {
				(*array)[i] = (*array)[i-1]
			}
		}

		for i := len(*array) - 1; i > 0; i-- {
			if (*array)[i] == "" {
				(*array)[i] = (*array)[i+1]
			}
		}
	}
}

//Возвращает самую длинную строку
func LongestString(s []string) string {
	if len(s) == 0 {
		return ""
	}
	max := 0
	result := 0
	for i, s2 := range s {
		if len(s2) > max {
			max = len(s2)
			result = i
		}
	}
	return s[result]
}
