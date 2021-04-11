package parse

import (
	"fmt"
	"github.com/plandem/xlsx"
	"log"
	"regexp"
	"schedule/download"
	"schedule/html"
	"schedule/structure"
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

func Parse() []structure.Group {
	var count int
	var groups []structure.Group
	links, err := html.Parse()
	if err != nil {
		log.Panicf("Error occured while html parsing. %v", err)
	}

	regexpGroupNumber := regexp.MustCompile(`[А-Я]{4}[-]\d{2}[-]\d{2}`)
	for i, link := range links[0] {
		path := "C:/Excel/" + strconv.Itoa(i) + ".xlsx"
		defer fmt.Println(path)
		err := download.GetFile(path, link)
		if err != nil {
			panic(err)
		}
		xl, err := xlsx.Open(path)
		if err != nil {
			fmt.Println(link)
			fmt.Println(path)
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
			for table[rowInfo][colInfo] == table[rowInfo][colInfo+1] { //фикс скрытого A столбика в ТХТ
				colInfo++
			}
			rowInfo += 2
			stringsUnique := make([]string, len(table))
			rowsTable := GetRows(table) // количество строк
			for rowGroup, strings := range table {
				for colGroup, s := range strings {

					if regexpGroupNumber.MatchString(str.ToTitle(s)) && !Contains(stringsUnique, regexpGroupNumber.FindString(s)) {
						count++
						stringsUnique = append(stringsUnique, regexpGroupNumber.FindString(s))
						if CheckSubgroups(&table, colGroup, rowInfo, rowsTable) {
							SubgroupNumber = 1
							group1 := GetGroup(&table, rowGroup, colGroup, colInfo, rowInfo, rowsTable)
							group1.Name = regexpGroupNumber.FindString(str.ToTitle(s)) + "-1"
							group1.SubGroup = 1
							group1.Clear()
							groups = append(groups, group1)

							SubgroupNumber = 2
							group2 := GetGroup(&table, rowGroup, colGroup, colInfo, rowInfo, rowsTable)
							group2.Name = regexpGroupNumber.FindString(str.ToTitle(s)) + "-2"
							group2.SubGroup = 2
							group2.Clear()
							groups = append(groups, group2)
						} else {
							group := GetGroup(&table, rowGroup, colGroup, colInfo, rowInfo, rowsTable)
							group.Name = regexpGroupNumber.FindString(str.ToTitle(s))
							group.SubGroup = 0
							group.Clear()
							groups = append(groups, group)
						}
						SubgroupNumber = 0
					}
				}
			}
		}
	}
	log.Println("Кол-во групп")
	log.Println(count)
	return groups
}

var SubgroupRegexp = regexp.MustCompile("[^А-Яа-я](п/гр|гр|подгр|подгруп|п/г|подгруппа)([^А-Яа-я]|$)")

func CheckSubgroups(table *[][]string, colGroup int, rowInfo int, rows int) bool {
	for i := rowInfo; i < rows; i++ {
		if SubgroupRegexp.MatchString((*table)[i][colGroup]) {
			return true
		}
	}
	return false
}

func GetGroup(table *[][]string, rowGroup int, colGroup int, colInfo int, rowInfo int, rows int) structure.Group {
	result := structure.NewGroup()
	var lessons []structure.Lesson
	for i := rowInfo; i < rows; i++ {
		/*
			table[i][colGroup]		предмет
			table[i][colGroup+1]	вид занятия
			table[i][colGroup+2]	ФИО преподавателя
			table[i][colGroup+3]	№ аудитории
			table[i][colInfo]		день недели
			table[i][colInfo+1]  	№пары
			table[i][colInfo+4]  	Неделя
		*/
		if SubgroupRegexp.MatchString((*table)[i][colGroup]) { //проверка на подгруппы
			lessons = SubGroupParse((*table)[i][colGroup], (*table)[i][colGroup+1], (*table)[i][colGroup+2], (*table)[i][colGroup+3], (*table)[i][colInfo], (*table)[i][colInfo+1], (*table)[i][colInfo+4])
		} else {
			lessons = DefaultParse((*table)[i][colGroup], (*table)[i][colGroup+1], (*table)[i][colGroup+2], (*table)[i][colGroup+3], (*table)[i][colInfo], (*table)[i][colInfo+1], (*table)[i][colInfo+4])
		}
		for j := 0; j < len(lessons); j++ {
			if !lessons[j].Exists {
				structure.RemoveElementLesson(&lessons, j)
				j--
			}
		}
		if !(Exist(&lessons) != -1 && len(lessons) == 1) {
			result.Days[(*table)[i][colInfo]] = append(result.Days[(*table)[i][colInfo]], lessons...)
		}
	}
	return result
}

func Exist(lessons *[]structure.Lesson) int {
	for i, lesson := range *lessons {
		if lesson.Exists == false {
			return i
		}
	}
	return -1
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
