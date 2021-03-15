package main

import (
	. "./Parse"
	"./Structure"
	"fmt"
	"github.com/plandem/xlsx"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
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
	links := []string{
		"https://webservices.mirea.ru/upload/iblock/2c3/%D0%A4%D0%A2%D0%98_1%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/f03/%D0%A4%D0%A2%D0%98_2%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/a4e/%D0%A4%D0%A2%D0%98_3%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/142/%D0%A4%D0%A2%D0%98_4%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/b00/%D0%A4%D0%A2%D0%98_%D0%A1%D1%82%D1%80%D0%BE%D0%BC%D1%8B%D0%BD%D0%BA%D0%B0%201%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/d00/%D0%A4%D0%A2%D0%98_%D0%A1%D1%82%D1%80%D0%BE%D0%BC%D1%8B%D0%BD%D0%BA%D0%B0%202%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/9f2/%D0%A4%D0%A2%D0%98_%D0%A1%D1%82%D1%80%D0%BE%D0%BC%D1%8B%D0%BD%D0%BA%D0%B0%203%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/e34/%D0%A4%D0%A2%D0%98_%D0%A1%D1%82%D1%80%D0%BE%D0%BC%D1%8B%D0%BD%D0%BA%D0%B0%204%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/9ef/%D0%98%D0%98%D0%9D%D0%A2%D0%95%D0%93%D0%A3_1%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/974/%D0%98%D0%98%D0%9D%D0%A2%D0%95%D0%93%D0%A3_2%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/134/%D0%98%D0%98%D0%9D%D0%A2%D0%95%D0%93%D0%A3_3%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/b47/%D0%98%D0%98%D0%9D%D0%A2%D0%95%D0%93%D0%A3_4%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/a52/%D0%98%D0%98%D0%9D%D0%A2%D0%95%D0%93%D0%A3_%D0%B7%D0%B0%D0%BE%D1%87_4%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/c20/%D0%98%D0%98%D0%A2_1%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/4f2/%D0%98%D0%98%D0%A2_2%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/e08/%D0%98%D0%98%D0%A2_3%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/0f0/%D0%98%D0%98%D0%A2_4%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/8b9/%D0%98%D0%9A_1%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/245/%D0%98%D0%9A_2%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/b66/%D0%98%D0%9A_3%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/0b4/%D0%98%D0%9A_4%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/962/%D0%98%D0%9A_5%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/59e/%D0%9A%D0%91%D0%B8%D0%A1%D0%9F%201%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/7e3/%D0%9A%D0%91%D0%B8%D0%A1%D0%9F%202%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/f5f/%D0%9A%D0%91%D0%B8%D0%A1%D0%9F%203%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/9ab/%D0%9A%D0%91%D0%B8%D0%A1%D0%9F%204%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/641/%D0%9A%D0%91%D0%B8%D0%A1%D0%9F%205%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/8e7/%D0%98%D0%A0%D0%A2%D0%A1_1%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/0c5/%D0%98%D0%A0%D0%A2%D0%A1_2%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/9ec/%D0%98%D0%A0%D0%A2%D0%A1_3%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/234/%D0%98%D0%A0%D0%A2%D0%A1_4%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/452/%D0%98%D0%A0%D0%A2%D0%A1_5%D0%BA_20-21_%D0%B2%D0%B5%D1%81%D0%BD%D0%B0.xlsx",
		"https://webservices.mirea.ru/upload/iblock/6e7/%D0%98%D0%A2%D0%A5%D0%A2_%D0%B1%D0%B0%D0%BA_1%D0%BA_20-21_%D0%BB%D0%B5%D1%82%D0%BE.xlsx",
		"https://webservices.mirea.ru/upload/iblock/60c/%D0%98%D0%A2%D0%A5%D0%A2_%D0%B1%D0%B0%D0%BA_2%D0%BA_20-21_%D0%BB%D0%B5%D1%82%D0%BE.xlsx",
		"https://webservices.mirea.ru/upload/iblock/955/%D0%98%D0%A2%D0%A5%D0%A2_%D0%B1%D0%B0%D0%BA_3%D0%BA_20-21_%D0%BB%D0%B5%D1%82%D0%BE.xlsx",
		"https://webservices.mirea.ru/upload/iblock/c16/%D0%98%D0%A2%D0%A5%D0%A2_%D0%B1%D0%B0%D0%BA_4%D0%BA_20-21_%D0%BB%D0%B5%D1%82%D0%BE.xlsx",
		"https://webservices.mirea.ru/upload/iblock/420/%D0%98%D0%AD%D0%9F%201%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/c01/%D0%98%D0%AD%D0%9F%202%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/6f6/%D0%98%D0%AD%D0%9F%203%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/667/%D0%98%D0%AD%D0%9F%204%20%D0%BA%D1%83%D1%80%D1%81%202%20%D1%81%D0%B5%D0%BC-%D0%94.xlsx",
		"https://webservices.mirea.ru/upload/iblock/11e/%D0%98%D0%AD%D0%9F_4%20%D0%BA%D1%83%D1%80%D1%81_%D0%B7%D0%B0%D0%BE%D1%87%D0%BD%D0%BE%D0%B5.xlsx",
		"https://webservices.mirea.ru/upload/iblock/9d0/%D0%9E-%D0%97%205%20%D0%BA%D1%83%D1%80%D1%81%20%D0%A0%D0%A2%D0%A3%20%D0%9C%D0%98%D0%A0%D0%AD%D0%90%20(%D0%B2%D0%B5%D1%81%D0%BD%D0%B0)..xlsx",
		"https://webservices.mirea.ru/upload/iblock/db2/%D0%97_4_%D0%BA%D1%83%D1%80%D1%81_%D0%A0%D0%A2%D0%A3_%D0%9C%D0%98%D0%A0%D0%AD%D0%90_(%D0%B2%D0%B5%D1%81%D0%BD%D0%B0).xlsx",
		"https://webservices.mirea.ru/upload/iblock/6b4/%D0%97%205%20%D0%BA%D1%83%D1%80%D1%81%20%D0%A0%D0%A2%D0%A3%20%D0%9C%D0%98%D0%A0%D0%AD%D0%90%20(%D0%B2%D0%B5%D1%81%D0%BD%D0%B0)..xlsx",
	} //массив с ссылками на excel файлы
	for i, _ := range links {
		path := "C:/Excel/" + strconv.Itoa(i) + ".xlsx"
		//err := DownloadFile(path, link)
		//if err != nil {
		//	panic(err)
		//}
		xl, err := xlsx.Open(path)
		//xl, err := xlsx.Open(`C:/Excel/1.xlsx`)
		if err != nil {
			panic(err)
		}
		test := xl.Sheets()
		for test.HasNext() { //следующий лист
			_, sheet := test.Next()
			table := GetTable(sheet)
			rowInfo, colInfo := GetCoords(table, "день недели") //коориданаты панели с днём недели №пары и т.д.
			if rowInfo == -1 {                                  // Чек на пустую страницу excel
				continue
			}
			rowInfo += 2
			stringsUnique := make([]string, len(table))
			rowsTable := GetRows(table)
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
	}
}

func GetGroup(table [][]string, rowGroup int, colGroup int, colInfo int, rowInfo int, rows int) Structure.Group {
	group := Structure.NewGroup()
	for i := rowInfo; i < rows; i++ {
		if regexp.MustCompile(" *?кр((\\.)|(  ??))").MatchString(table[i][colGroup]) {
			fmt.Println(table[i][colGroup])   //предмет
			fmt.Println(table[i][colGroup+1]) //вид занятия
			fmt.Println(table[i][colGroup+2]) //ФИО преподавателя
			fmt.Println(table[i][colGroup+3]) //№ аудитории
			fmt.Println(table[i][colInfo])    //день недели
			fmt.Println(table[i][colInfo+1])  //№пары
			fmt.Println(table[i][colInfo+4])  //Неделя
			fmt.Println("-------------------------------------------------")
		}
		//fmt.Println(table[i][colGroup])   //предмет
		//fmt.Println(table[i][colGroup+1]) //вид занятия
		//fmt.Println(table[i][colGroup+2]) //ФИО преподавателя
		//fmt.Println(table[i][colGroup+3]) //№ аудитории
		//fmt.Println(table[i][colInfo])    //день недели
		//fmt.Println(table[i][colInfo+1])  //№пары
		//fmt.Println(table[i][colInfo+4])  //Неделя
		//надо из этих 4 данных получать несколько уроков.
		lessons := ParseIKBSP(table[i][colGroup], table[i][colGroup+1], table[i][colGroup+2], table[i][colGroup+3], table[i][colInfo], table[i][colInfo+1], table[i][colInfo+4])
		group.AddLesson(lessons)
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
