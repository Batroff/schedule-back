package excel

import (
	"github.com/batroff/schedule-back/models"
	"github.com/pkg/errors"
	"github.com/plandem/xlsx"
	"regexp"
	str "strings"
)

var (
	regexpGroupNumber = regexp.MustCompile(`[А-Я]{4}[-]\d{2}[-]\d{2}`)
	SubgroupRegexp    = regexp.MustCompile("[^А-Яа-я](п/гр|гр|подгр|подгруп|п/г|подгруппа)([^А-Яа-я]|$)")
)

func ParseMultiple(excelPaths []string) (*[]models.Group, error) {
	var groups = new([]models.Group)

	for _, filepath := range excelPaths {
		group, parseErr := parse(filepath) // группа может быть с подгруппой поэтому массив
		if parseErr != nil {
			return nil, parseErr
		}

		for _, g := range *group {
			*groups = append(*groups, g)
		}
	}

	return groups, nil
}

func parse(excelPath string) (*[]models.Group, error) {
	var uniqueGroupsList []string
	var groups = new([]models.Group)

	xl, excelErr := xlsx.Open(excelPath)
	if excelErr != nil {
		return nil, errors.Wrapf(excelErr, "Excel open file %s error", excelPath)
	}
	sheetIterator := xl.Sheets()
	for sheetIterator.HasNext() { //следующий лист
		_, sheet := sheetIterator.Next()
		table := GetTable(sheet)
		rowInfo, colInfo := GetCoords(table, "день недели") //коориданаты панели с днём недели №пары и т.д.
		if rowInfo == -1 {                                  //чек на пустую страницу excel
			continue
		}
		for table[rowInfo][colInfo] == table[rowInfo][colInfo+1] { //фикс скрытого A столбика в ТХТ
			colInfo++
		}
		rowInfo += 2
		rowsTable := GetRows(table) // количество строк

		for _, strings := range table {
			for colGroup, s := range strings {
				if regexpGroupNumber.MatchString(str.ToTitle(s)) && !Contains(uniqueGroupsList, regexpGroupNumber.FindString(s)) {
					uniqueGroupsList = append(uniqueGroupsList, regexpGroupNumber.FindString(s))
					if CheckSubgroups(&table, colGroup, rowInfo, rowsTable) {
						GroupCreate(groups, &table, colGroup, colInfo, rowInfo, rowsTable, 1, s)
						GroupCreate(groups, &table, colGroup, colInfo, rowInfo, rowsTable, 2, s)
						models.GroupMap[regexpGroupNumber.FindString(s)] = true
					} else {
						GroupCreate(groups, &table, colGroup, colInfo, rowInfo, rowsTable, 2, s)
						models.GroupMap[regexpGroupNumber.FindString(s)] = false
					}
				}
			}
		}
	}

	return groups, nil
}

func GroupCreate(groups *[]models.Group, table *[][]string, colGroup int, colInfo int, rowInfo int, rows int, subgroup int, groupName string) {
	SubgroupNumber = subgroup
	group := GetGroup(table, colGroup, colInfo, rowInfo, rows)
	group.Name = regexpGroupNumber.FindString(str.ToTitle(groupName))
	group.SubGroup = subgroup
	group.Clear()
	*groups = append(*groups, group)
	SubgroupNumber = 0
}

func CheckSubgroups(table *[][]string, colGroup int, rowInfo int, rows int) bool {
	for i := rowInfo; i < rows; i++ {
		if SubgroupRegexp.MatchString((*table)[i][colGroup]) {
			return true
		}
	}
	return false
}

func GetGroup(table *[][]string, colGroup int, colInfo int, rowInfo int, rows int) models.Group {
	var result = models.NewGroup()
	var lessons []models.Lesson
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
			lessons = SubGroupParse((*table)[i][colGroup], (*table)[i][colGroup+1], (*table)[i][colGroup+2],
				(*table)[i][colGroup+3], (*table)[i][colInfo], (*table)[i][colInfo+1], (*table)[i][colInfo+4])
		} else {
			lessons = DefaultParse((*table)[i][colGroup], (*table)[i][colGroup+1], (*table)[i][colGroup+2],
				(*table)[i][colGroup+3], (*table)[i][colInfo], (*table)[i][colInfo+1], (*table)[i][colInfo+4])
		}
		for j := 0; j < len(lessons); j++ {
			if !lessons[j].Exists {
				models.RemoveElementLesson(&lessons, j)
				j--
			}
		}
		if !(Exist(&lessons) != -1 && len(lessons) == 1) {
			result.Days[(*table)[i][colInfo]] = append(result.Days[(*table)[i][colInfo]], lessons...)
		}
	}
	return result
}

func Exist(lessons *[]models.Lesson) int {
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

func Contains(strings []string, s string) bool {
	for _, s2 := range strings {
		if s2 == s {
			return true
		}
	}
	return false
}
