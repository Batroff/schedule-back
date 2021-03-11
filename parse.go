package main

import (
	TTS "Schedule/Structure"
	"fmt"
	"github.com/plandem/xlsx"
	"io"
	"net/http"
	"os"
	str "strings"
)

//type lesson struct {
//	subject      string //название предмета
//	typeOfLesson string //тип занятия
//	teacherName  string //фио преподавателя
//	cabinet      string //кабинет
//	numberLesson int    //номер пары
//	//occurrenceLesson []int//номера недель в которых присутствует эта пара
//	occurrenceLesson []bool //номера недель в которых присутствует эта пара
//	exists           bool   //для пустых пар??
//}
//
//type day struct {
//	lessons []lesson
//}
//
//type week struct {
//	days []day
//}
//
//type group struct {
//	weeks []week
//}

//func newGroup() group {
//	var g group
//	g.weeks = make([]week, 17)
//	for i := range g.weeks {
//		g.weeks[i] = newWeek()
//	}
//	return g
//}
//
//func newWeek() week {
//	var w week
//	w.days = make([]day, 6)
//	for i := range w.days {
//		w.days[i] = newDay()
//	}
//	return w
//}
//
//func newDay() day {
//	var d day
//	d.lessons = make([]lesson, 8)
//	for i := range d.lessons {
//		d.lessons[i] = newLesson()
//	}
//	return d
//}
//
//func newLesson() lesson {
//	var l lesson
//	l.occurrenceLesson = make([]bool, 17)
//	return l
//}
//
//func (g group) AddLesson(lessons []lesson) {
//	//
//}

//

func Contains(strings []string, s string) bool {
	for _, s2 := range strings {
		if s2 == s {
			return true
		}
	}
	return false
}

//func Index() {
//	url := "https://webservices.mirea.ru/upload/iblock/fac/%D0%A4%D0%A2%D0%98_1%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx"
//
//	resp, err := http.Get(url)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(len(body))
//	//download the file in browser
//
//}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	err := DownloadFile(`C:/Users/Kolya/test.xlsx`, `https://webservices.mirea.ru/upload/iblock/fac/%D0%A4%D0%A2%D0%98_1%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx`)
	if err != nil {
		panic(err)
	}

	//xl, err := xlsx.Open(`C:/Users/Kolya/Downloads/КБиСП 2 курс 2 сем-Д.xlsx`)
	group := TTS.NewGroup()
	table := GetTable(`C:/Users/Kolya/Downloads/ФТИ_1к_20-21_весна.xlsx`)

	fmt.Println(len(table))
	fmt.Println(len(table[0]))
	groupNumber := "ЭЛБО-01-20"
	x, y := GetCoords(table, groupNumber)
	fmt.Println(x, y)
	fmt.Println(table[1][5])
	fmt.Println(GetRaws(table))
	xInfo, yInfo := GetCoords(table, "день недели") //коориданаты панели с днём недели №пары и т.д.
	xInfo += 2
	yInfo = yInfo + yInfo - yInfo
	parse := GetParseFunc(groupNumber)
	//result := newGroup()
	for i := xInfo; i < GetRaws(table); i++ {
		//table[i][y] предмет
		//table[i][y+1] вид занятия
		//table[i][y+2] ФИО преподавателя
		//table[i][y+3] № аудитории
		//надо из этих 4 данных получать несколько уроков.
		lessons := parse(table[i][y], table[i][y+1], table[i][y+2], table[i][y+3])
		group.AddLesson(lessons)
	}

}

func GetParseFunc(group string) func(subject, typoOfLesson, teacherName, cabinet string) []TTS.Lesson {
	//проверка на институт
}

//func ParseFTI(subject, typeOfLesson, teacherName, cabinet string) []TTS.Lesson {
//	return []TTS.Lesson{TTS.NewLesson(), TTS.NewLesson()}
//}
//
//func ParseKIB(subject, typeOfLesson, teacherName, cabinet string) []TTS.Lesson {
//	return []TTS.Lesson{TTS.NewLesson(), TTS.NewLesson()}
//}

func GetRaws(table [][]string) int { //количество строк в таблице
	days := [6]string{"ПОНЕДЕЛЬНИК", "ВТОРНИК", "СРЕДА", "ЧЕТВЕРГ", "ПЯТНИЦА", "СУББОТА"}
	x, y := GetCoords(table, "день недели") //если они где-то непоставили пробел или написали по-другому....
	x += 2
	for Contains(days[0:], table[x][y]) {
		x++
	}
	return x + 1
}

func GetCoords(table [][]string, group string) (a, b int) { //координаты по строке
	for i, strings := range table {
		for j, s := range strings {
			if str.Contains(str.ToLower(s), str.ToLower(group)) {
				return i, j
			}
		}
	}
	return -1, -1
}

func GetTable(path string) [][]string {
	xl, err := xlsx.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	sheet := xl.Sheet(0)
	fmt.Println(sheet.Cell(5, 1))
	c, r := sheet.Dimension()
	var table = make([][]string, r)
	for i := range table {
		table[i] = make([]string, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			table[i][j] = sheet.Cell(j, i).String()
		}
	}
	return table
}
