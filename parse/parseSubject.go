package parse

import (
	"fmt"
	"regexp"
	. "schedule/structure"
	"strconv"
	"strings"
)

func ParseIKBSP(subject, typeOfLesson, teacherName, cabinet, dayOfWeek, numberLesson, week string) []Lesson {
	//someLesson := NewLesson()
	//
	//if regexp.MustCompile("(((\\d{2}|\\d)( *[ \\-,] *)){1,17}(\\d{2}|\\d))|(\\d{2}|\\d)").MatchString(subject) {
	//flag := exceptFlag(subject)
	//	someLesson.FillInWeeks(flag, week)
	//loc := regexp.MustCompile("(((\\d{2}|\\d)( *[ \\-,] *)){1,17}(\\d{2}|\\d))|(\\d{2}|\\d)").FindStringIndex(subject)
	//	fmt.Println(subject)
	//	string := subject[loc[0]:loc[1]]
	//	string = removeSpaces(string)
	//	fmt.Println(string)
	//	fmt.Println(week)
	//FillingInOccurrenceLesson(flag,week,someLesson,string) // Заполнение массива недель на которых будет пара
	//	//for _, v := range numbersPresent(string) {
	//	//	if flag && strings.Contains(week, "II") && (v-1)%2 == 0 {
	//	//		someLesson.OccurrenceLesson[v-1] = false
	//	//	} else if flag && strings.Contains(week, "I") && (v-1)%2 == 0 {
	//	//		someLesson.OccurrenceLesson[v-1] = false
	//	//	} else if !flag && strings.Contains(week, "II") && (v-1)%2 == 0 {
	//	//		someLesson.OccurrenceLesson[v-1] = true
	//	//	} else if !flag && strings.Contains(week, "I") && (v-1)%2 == 0 {
	//	//		someLesson.OccurrenceLesson[v-1] = true
	//	//	}
	//	//}
	//	for _, v := range someLesson.OccurrenceLesson {
	//		fmt.Print(v)
	//		fmt.Print(" ")
	//	}
	//	fmt.Println()
	//}
	//	subject = RemoveJunk(subject)
	//	if strings.Contains(subject,"/"){
	//		fmt.Println("Предмет: ")
	//		fmt.Println(subject)
	//		fmt.Println("Препод: ")
	//		fmt.Println(teacherName)
	//		fmt.Println("Кабинет: ")
	//		fmt.Println(cabinet)
	//		fmt.Println("Разбиение на строки предмета: ")
	//		for _, v := range SeparateLessons(subject) {
	//			fmt.Println(v)
	//		}
	//		fmt.Println("Разбиение на строки преподов: ")
	//		for _, v := range SeparateTeachers(teacherName) {
	//			fmt.Println(v)
	//		}
	//		fmt.Println("Разбиение на строки кабинетов: ")
	//		for _, v := range SeparateTeachers(cabinet) {
	//			fmt.Println(v)
	//		}
	//		fmt.Println(typeOfLesson)
	//		fmt.Println("------------------------------")
	//	}
	subject = RemoveJunk(subject)
	if len(SeparateLessons(subject)) > 2 {
		fmt.Println("Предмет: ")
		fmt.Println(subject)
		fmt.Println("Препод: ")
		fmt.Println(teacherName)
		fmt.Println("Кабинет: ")
		fmt.Println(cabinet)
		fmt.Println("Тип: ")
		fmt.Println(typeOfLesson)
		a, b, c, d := countBalance(SlashManage(SeparateLessons(subject), SeparateTeachers(teacherName), SeparateCabinets(cabinet), SeparateCabinets(typeOfLesson)))
		fmt.Println("прдмт")
		for _, v := range a {
			fmt.Println(v)
		}
		fmt.Println("уч")
		for _, v := range b {
			fmt.Println(v)
		}
		fmt.Println("каб")
		for _, v := range c {
			fmt.Println(v)
		}
		fmt.Println("тип")
		for _, v := range d {
			fmt.Println(v)
		}
		fmt.Println("======================================")
	}
	return []Lesson{NewLesson(), NewLesson()}
}
func numbersPresent(subject string) []int { // Возвращает номера недель в предмете subject
	stringNumbers := orSplit(subject)
	intNumbers := []int{}
	for _, v := range stringNumbers {
		if strings.Contains(v, "-") {
			firstNum, err := strconv.Atoi(v[0:strings.Index(v, "-")])
			if err != nil {
				fmt.Println(err)
			}
			lastNum, err := strconv.Atoi(v[strings.Index(v, "-")+1 : len(v)])
			if err != nil {
				fmt.Println(err)
			}
			for i := firstNum; i <= lastNum; i++ {
				intNumbers = append(intNumbers, i)
			}
		} else {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			intNumbers = append(intNumbers, num)
		}
	}
	return intNumbers
}
func orSplit(subject string) []string { // делит номера недель
	if strings.Contains(subject, ",") {
		return strings.Split(subject, ",")
	} else {
		return strings.Split(subject, " ")
	}
}
func removeSpaces(subject string) string {
	subject = strings.ReplaceAll(subject, "  ", "")
	if subject[len(subject)-1:] == " " {
		subject = subject[:len(subject)-1]
	}
	return subject
}
func exceptFlag(subject string) bool {
	flag := false
	if regexp.MustCompile("((^)|( ))кр((\\.)|(  ??))").MatchString(subject) {
		flag = true
	}
	return flag
}
func NewLineSeparator(line string) []string {
	line = strings.ReplaceAll(line, "/", "\n")
	return strings.Split(line, "\n")
}
func RemoveJunk(line string) string {
	line = strings.ReplaceAll(line, ".", ",")
	line = strings.ReplaceAll(line, "\n", "")
	return line
}
func SeparateLessons(line string) []string {
	var lessons []string
	if regexp.MustCompile("\\((\\d{2}|\\d)-(\\d{2}|\\d) нед,\\/ (\\d{2}|\\d)-(\\d{2}|\\d) нед,\\)").MatchString(line) {
		lessons = append(lessons, line[0:strings.Index(line, "(")-1]+line[strings.Index(line, "(")+1:strings.Index(line, "/")-1])
		lessons = append(lessons, line[0:strings.Index(line, "(")-1]+line[strings.Index(line, "/")+1:strings.Index(line, ")")-1])
	} else if regexp.MustCompile("\\((\\d{2}|\\d)-(\\d{2}|\\d) нед \\/ (\\d{2}|\\d)-(\\d{2}|\\d) нед\\)").MatchString(line) {
		lessons = append(lessons, line[0:strings.Index(line, "(")-1]+line[strings.Index(line, "(")+1:strings.Index(line, "/")-1])
		lessons = append(lessons, line[0:strings.Index(line, "(")-1]+line[strings.Index(line, "/")+1:strings.Index(line, ")")])
	} else {
		for HasNextNumbers(line) > 0 {
			lessons = append(lessons, line[0:HasNextNumbers(line)])
			line = line[HasNextNumbers(line):]
		}
		if line != "" {
			lessons = append(lessons, line)
		}
	}
	return lessons
}
func NumbersCheck(line string) bool { //Проверяет есть ли в строке номера недель
	if regexp.MustCompile("(((\\d{2}|\\d)( *[ \\-,] *)){1,17}(\\d{2}|\\d))|(\\d{2}|\\d)").MatchString(line) {
		return true
	}
	return false
}
func NumbersIndex(line string) []int { //Возвращает начальный и конечный индексы вхождения номеров недель в строку
	if regexp.MustCompile("(((\\d{2}|\\d)( *[ \\-,] *)){1,17}(\\d{2}|\\d))|(\\d{2}|\\d)").FindStringIndex(line) != nil {
		return regexp.MustCompile("(((\\d{2}|\\d)( *[ \\-,] *)){1,17}(\\d{2}|\\d))|(\\d{2}|\\d)").FindStringIndex(line)
	} else {
		return []int{-1000, -1000}
	}
}
func HasNextNumbers(line string) int {
	count := len(line)
	if NumbersIndex(line)[0] != -1000 {
		line = line[NumbersIndex(line)[1]:len(line)]
		count -= len(line)
		return NumbersIndex(line)[0] + count
	} else {
		return -1
	}
}
func FillingInOccurrenceLesson(flag bool, week string, someLesson Lesson, line string) {
	for _, v := range numbersPresent(line) {
		if flag && strings.Contains(week, "II") && (v-1)%2 == 0 {
			someLesson.OccurrenceLesson[v-1] = false
		} else if flag && strings.Contains(week, "I") && (v-1)%2 == 0 {
			someLesson.OccurrenceLesson[v-1] = false
		} else if !flag && strings.Contains(week, "II") && (v-1)%2 == 0 {
			someLesson.OccurrenceLesson[v-1] = true
		} else if !flag && strings.Contains(week, "I") && (v-1)%2 == 0 {
			someLesson.OccurrenceLesson[v-1] = true
		}
	}
}
func SeparateTeachers(line string) []string {
	var teachers []string
	//for regexp.MustCompile("([А-Я]\\.){2}").MatchString(line){
	//	teachers = append(teachers, line[0:regexp.MustCompile("([А-Я]\\.){2}").FindStringIndex(line)[1]])
	//	line = line[regexp.MustCompile("([А-Я]\\.){2}").FindStringIndex(line)[1]:len(line)]
	//}
	for strings.Contains(line, "\n") {
		teachers = append(teachers, line[0:strings.Index(line, "\n")])
		line = line[strings.Index(line, "\n")+1 : len(line)]
	}
	teachers = append(teachers, line)
	return teachers
}
func SeparateCabinets(line string) []string {
	var cabinets []string
	//for regexp.MustCompile("([А-Я]\\.){2}").MatchString(line){
	//	teachers = append(teachers, line[0:regexp.MustCompile("([А-Я]\\.){2}").FindStringIndex(line)[1]])
	//	line = line[regexp.MustCompile("([А-Я]\\.){2}").FindStringIndex(line)[1]:len(line)]
	//}
	for strings.Contains(line, "\n") {
		cabinets = append(cabinets, line[0:strings.Index(line, "\n")])
		line = line[strings.Index(line, "\n")+1 : len(line)]
	}
	cabinets = append(cabinets, line)
	return cabinets
}
func SlashFix(lessons []string) []string {
	for i, v := range lessons {
		lessons[i] = removeSpaces(v)
	}
	for i := 0; i < len(lessons); i++ {
		if lessons[i][len(lessons[i])-1:] == "/" {
			lessons[i] = lessons[i] + lessons[i+1]
			copy(lessons[i+1:], lessons[i+2:])
			lessons = lessons[:len(lessons)-1]
		}
	}
	return lessons
}
func SlashManage(lessons, teachers, cabinets, types []string) ([]string, []string, []string, []string) {
	newLessons := lessons
	newTeachers := teachers
	newCabinets := cabinets
	newTypes := types
	for _, l := range lessons {
		if strings.Contains(l, "/") {
			for _, t := range teachers {
				if strings.Contains(t, "/") {
					newLessons = SliceSlashManage(lessons)
					newTeachers = SliceSlashManage(teachers)
				}
			}
			for _, c := range cabinets {
				if strings.Contains(c, "/") {
					newLessons = SliceSlashManage(lessons)
					newCabinets = SliceSlashManage(cabinets)
				}
			}
			for _, t := range types {
				if strings.Contains(t, "/") {
					newLessons = SliceSlashManage(lessons)
					newTypes = SliceSlashManage(types)
				}
			}
		}
	}
	return newLessons, newTeachers, newCabinets, newTypes
}
func SliceSlashManage(slice []string) []string { // Разбивает элементы слайса по слешам
	partBefore := ""
	for i := 0; i < len(slice); i++ {
		if HasNextNumbers(slice[i]) < 0 && NumbersIndex(slice[i])[0] > 0 {
			partBefore = slice[i][0 : NumbersIndex(slice[i])[1]+4] // часть с кроме или номерами недель
			slice[i] = slice[i][NumbersIndex(slice[i])[1]+4:]
		}
		if strings.Contains(slice[i], "/") {
			slice = append(slice, "")
			copy(slice[i+1:], slice[i:])
			slice[i] = partBefore + slice[i][0:strings.Index(slice[i], "/")]
			slice[i+1] = partBefore + slice[i+1][strings.Index(slice[i+1], "/")+1:len(slice[i+1])]
		}
	}
	return slice
}
func countBalance(lessons, teachers, cabinets, types []string) ([]string, []string, []string, []string) {
	if len(lessons) > len(teachers) {
		lessons, teachers = balanceSlices(lessons, teachers)
	}
	if len(lessons) > len(cabinets) {
		lessons, cabinets = balanceSlices(lessons, cabinets)
	}
	if len(lessons) > len(types) {
		lessons, types = balanceSlices(lessons, types)
	}
	return lessons, teachers, cabinets, types
}
func balanceSlices(lessons, teachers []string) ([]string, []string) {
	if len(lessons) == len(teachers)*2 {
		for len(lessons) != len(teachers) {
			teachers = append(teachers, "")
		}
		if len(lessons) == 4 {
			teachers[3] = teachers[1]
			teachers[2] = teachers[1]
			teachers[1] = teachers[0]
		} else if len(lessons) == 2 {
			teachers[1] = teachers[0]
		}
	}
	return lessons, teachers
}
