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

func Parse() ([]structure.Group, []structure.GroupMini) {
	var count int
	var groups []structure.Group
	var groupsMini []structure.GroupMini
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
							group1, groupMini1 := GetGroup(&table, rowGroup, colGroup, colInfo, rowInfo, rowsTable)
							group1.Name = regexpGroupNumber.FindString(str.ToTitle(s)) + "-1"
							groupMini1.Name = regexpGroupNumber.FindString(str.ToTitle(s)) + "-1"
							group1.SubGroup = 1
							groupMini1.SubGroup = 1
							groups = append(groups, group1)
							groupsMini = append(groupsMini, groupMini1)

							SubgroupNumber = 2
							group2, groupMini2 := GetGroup(&table, rowGroup, colGroup, colInfo, rowInfo, rowsTable)
							group2.Name = regexpGroupNumber.FindString(str.ToTitle(s)) + "-2"
							groupMini2.Name = regexpGroupNumber.FindString(str.ToTitle(s)) + "-2"
							group2.SubGroup = 2
							groupMini2.SubGroup = 2
							groups = append(groups, group2)
							groupsMini = append(groupsMini, groupMini2)
						} else {
							group, groupMini := GetGroup(&table, rowGroup, colGroup, colInfo, rowInfo, rowsTable)
							group.Name = regexpGroupNumber.FindString(str.ToTitle(s))
							groupMini.Name = regexpGroupNumber.FindString(str.ToTitle(s))
							group.SubGroup = 0
							groupMini.SubGroup = 0
							groups = append(groups, group)
							groupsMini = append(groupsMini, groupMini)
						}
						SubgroupNumber = 0
					}
				}
			}
		}
	}
	fmt.Println("Кол-во групп")
	fmt.Println(count)
	return groups, groupsMini
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

func GetGroup(table *[][]string, rowGroup int, colGroup int, colInfo int, rowInfo int, rows int) (structure.Group, structure.GroupMini) {
	result := structure.NewGroup()
	resultTest := structure.NewGroupMini()
	var lessons []structure.Lesson
	for i := rowInfo; i < rows; i++ {
		if SubgroupRegexp.MatchString((*table)[i][colGroup]) { //проверка на подгруппы
			//if str.Contains(table[i][colGroup], "гр") || str.Contains(table[i][colGroup], "п/г") {
			//	fmt.Println(table[i][colGroup])   //предмет
			//	fmt.Println(table[i][colGroup+1]) //вид занятия
			//	fmt.Println(table[i][colGroup+2]) //ФИО преподавателя
			//	fmt.Println(table[i][colGroup+3]) //№ аудитории
			//  fmt.Println((*table)[i][colInfo])    //день недели
			//	fmt.Println(table[i][colInfo+1])  //№пары
			//	fmt.Println(table[i][colInfo+4])  //Неделя
			//	fmt.Println("-------------------------------------------------")
			lessons = SubGroupParse((*table)[i][colGroup], (*table)[i][colGroup+1], (*table)[i][colGroup+2], (*table)[i][colGroup+3], (*table)[i][colInfo], (*table)[i][colInfo+1], (*table)[i][colInfo+4])
			//for _, lesson := range lessons {
			//	if lesson.Exists {
			//		fmt.Println(lesson)
			//	}
			//}
			//resultTest.Days[(*table)[i][colInfo]] = lessons
			result.AddLesson(lessons)
		} else {

			lessons = DefaultParse((*table)[i][colGroup], (*table)[i][colGroup+1], (*table)[i][colGroup+2], (*table)[i][colGroup+3], (*table)[i][colInfo], (*table)[i][colInfo+1], (*table)[i][colInfo+4])
			//for _, lesson := range lessons {
			//	if lesson.Exists {
			//		fmt.Println(lesson)
			//	}
			//}
			//fmt.Println((*table)[i][colInfo+1])
			//resultTest.Days[(*table)[i][colInfo]] = lessons
			result.AddLesson(lessons)
		}

		for Exist(&lessons) != -1 {
			RemoveElementLesson(&lessons, Exist(&lessons))
		}
		test := structure.Day{Lessons: lessons}
		resultTest.Days[(*table)[i][colInfo]] = test
		//fmt.Println((*table)[i][colGroup])   //предмет
		//fmt.Println((*table)[i][colGroup+1]) //вид занятия
		//fmt.Println()table[i][colGroup+2])//ФИО преподавателя
		//table[i][colGroup+3]) //№ аудитории
		//table[i][colInfo])    //день недели
		//table[i][colInfo+1])  //№пары
		//table[i][colInfo+4])  //Неделя
		////надо из этих 4 данных получать несколько уроков.
		//fmt.Println((*table)[i][colInfo+1])

	}
	return result, resultTest
}

func Exist(lessons *[]structure.Lesson) int {
	for i, lesson := range *lessons {
		if lesson.Exists == false {
			return i
		}
	}
	return -1
}

func RemoveElementLesson(a *[]structure.Lesson, i int) {
	//*a = append((*a)[:i], (*a)[i+1:]...)
	(*a)[i] = (*a)[len(*a)-1]            // Copy last element to index i.
	(*a)[len(*a)-1] = structure.Lesson{} // Erase last element (write zero value).
	*a = (*a)[:len(*a)-1]
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
