package main

import (
	"Schedule/Structure"
	"fmt"
	"github.com/plandem/xlsx"
	"io"
	"net/http"
	"os"
	"regexp"
	//"strconv"
	str "strings"
)

func Contains(strings []string, s string) bool {
	for _, s2 := range strings {
		if s2 == s {
			return true
		}
	}
	return false
}

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
	//err := DownloadFile(`C:/Users/Kolya/test.xlsx`, `https://webservices.mirea.ru/upload/iblock/fac/%D0%A4%D0%A2%D0%98_1%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx`)
	//if err != nil {
	//	panic(err)
	//}
	Parse()
	//xl, err := xlsx.Open(`C:/Users/Kolya/Downloads/КБиСП 2 курс 2 сем-Д.xlsx`)
	//group := TTS.NewGroup()
	//table := GetTable(`C:/Users/Kolya/Downloads/ФТИ_1к_20-21_весна.xlsx`)
	//
	//fmt.Println(len(table))
	//fmt.Println(len(table[0]))
	//groupNumber := "ЭЛБО-01-20"
	//x, y := GetCoords(table, groupNumber)
	//fmt.Println(x, y)
	//fmt.Println(table[1][5])
	//fmt.Println(GetRows(table))
	//xInfo, yInfo := GetCoords(table, "день недели") //коориданаты панели с днём недели №пары и т.д.
	//xInfo += 2
	//yInfo = yInfo + yInfo - yInfo
	////parse := GetParseFunc(groupNumber)
	////result := newGroup()
	//for i := xInfo; i < GetRows(table); i++ {
	//	//table[i][y] предмет
	//	//table[i][y+1] вид занятия
	//	//table[i][y+2] ФИО преподавателя
	//	//table[i][y+3] № аудитории
	//	//table[i][yInfo] день недели
	//	//table[i][yInfo+1] №пары
	//	//table[i][yInfo+4] Неделя
	//	//надо из этих 4 данных получать несколько уроков.
	//	//	lessons := parse(table[i][y], table[i][y+1], table[i][y+2], table[i][y+3], table[i][yInfo], table[i][yInfo+1], table[i][yInfo+4])
	//	//	group.AddLesson(lessons)
	//}

}

func Parse() {
	//var links []string = make([]string, 2)//массив с ссылками на excel файлы
	//for i, link := range links {
	//path := "./Excel/" + strconv.Itoa(i) + ".xlsx"
	//err := DownloadFile(path, link)
	//if err != nil {
	//	panic(err)
	//}
	//xl, err := xlsx.Open(path)
	xl, err := xlsx.Open(`C:/Users/Kolya/Downloads/КБиСП 2 курс 2 сем-Д.xlsx`)
	if err != nil {
		panic(err)
	}
	test := xl.Sheets()
	for test.HasNext() { //следующий лист
		_, sheet := test.Next()
		table := GetTable(sheet)
		rowsTable := GetRows(table)
		rowInfo, colInfo := GetCoords(table, "день недели") //коориданаты панели с днём недели №пары и т.д.
		rowInfo += 2
		stringsUnique := make([]string, len(table))
		for rowGroup, strings := range table {
			for colGroup, s := range strings {
				regexpGroupNumber := regexp.MustCompile(`[А-Я]{4}[-]\d{2}[-]\d{2}`)
				if regexpGroupNumber.MatchString(str.ToTitle(s)) && !Contains(stringsUnique, regexpGroupNumber.FindString(s)) {
					stringsUnique = append(stringsUnique, regexpGroupNumber.FindString(s))
					GetGroup(table, rowGroup, colGroup, colInfo, rowInfo, rowsTable)
				}
			}
		}
	}
	//}
}

func GetGroup(table [][]string, rowGroup int, colGroup int, colInfo int, rowInfo int, rows int) Structure.Group {
	group := Structure.NewGroup()
	for i := rowInfo; i < rows; i++ {

		fmt.Println(table[i][colGroup])   //предмет
		fmt.Println(table[i][colGroup+1]) //вид занятия
		fmt.Println(table[i][colGroup+2]) //ФИО преподавателя
		fmt.Println(table[i][colGroup+3]) //№ аудитории
		fmt.Println(table[i][colInfo])    //день недели
		fmt.Println(table[i][colInfo+1])  //№пары
		fmt.Println(table[i][colInfo+4])  //Неделя
		//надо из этих 4 данных получать несколько уроков.
		//lessons := parse(table[i][y], table[i][y+1], table[i][y+2], table[i][y+3], table[i][yInfo], table[i][yInfo+1], table[i][yInfo+4])
		//group.AddLesson(lessons)
	}
	return group
}

func GetRows(table [][]string) int { //количество строк в таблице
	days := [6]string{"ПОНЕДЕЛЬНИК", "ВТОРНИК", "СРЕДА", "ЧЕТВЕРГ", "ПЯТНИЦА", "СУББОТА"}
	x, y := GetCoords(table, "день недели") //если они где-то непоставили пробел или написали по-другому....
	x += 2
	for Contains(days[0:], table[x][y]) {
		x++
		if x >= len(table) {
			return x
		}
	}
	return x
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

func GetTable(sheet xlsx.Sheet) [][]string {
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
