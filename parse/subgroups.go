package parse

import (
	"fmt"
	"regexp"
	. "schedule/structure"
	"strings"
)

func SubGroupParse(subject, typeOfLesson, teacherName, cabinet, dayOfWeek, numberLesson, week string) []Lesson {
	//несколько уроков в 1 дне надо раскидать по строкам и если одинаковые предметы почему они раскинуты(тип работы/преподы)
	var lessons []Lesson
	if strings.Contains(subject, "\n") { // если в строчке с предметом более 1 строки
		//fmt.Println("ПРЕДМЕТ:")
		//fmt.Println(subject)
		//fmt.Println("ТИП:")
		//fmt.Println(typeOfLesson)
		//fmt.Println("ФИО:")
		//fmt.Println(teacherName)
		//fmt.Println("КАБИНЕТ:")
		//fmt.Println(cabinet)
		if strings.Contains(cabinet, "В-78*\n") || strings.Contains(cabinet, "В-86*\n") || strings.Contains(cabinet, "МП-1*\n") {
			strings.ReplaceAll(cabinet, "В-78*\n", "В-78* ")
			strings.ReplaceAll(cabinet, "В-86*\n", "В-86* ")
			strings.ReplaceAll(cabinet, "МП-1*\n", "МП-1* ")
		}
		subjects := strings.Split(subject, "\n")
		typesOfLessons := strings.Split(typeOfLesson, "\n")
		teachersNames := strings.Split(teacherName, "\n")
		cabinets := strings.Split(cabinet, "\n")
		//max := Max(len(subjects), len(typesOfLessons), len(teachersNames), len(cabinets)) // максимальное число строк в ячейке
		typesOfLessons = RemoveLastEmptyElement(typesOfLessons)
		teachersNames = RemoveLastEmptyElement(teachersNames)
		cabinets = RemoveLastEmptyElement(cabinets)
		if Contains(subjects, "…………………") { // =)
			for i, s := range subjects {
				if s == "…………………" {
					subjects = RemoveElement(subjects, i)
				}
			}
		}
		if Contains(subjects, "") {
			for i, s := range subjects {
				if s == "" {
					subjects = RemoveElement(subjects, i)
				}
			}

			if Contains(typesOfLessons, "") {
				for i, s := range typesOfLessons {
					if s == "" {
						subjects = RemoveElement(typesOfLessons, i)
					}
				}
			}
			if Contains(teachersNames, "") {
				for i, s := range teachersNames {
					if s == "" {
						subjects = RemoveElement(teachersNames, i)
					}
				}
			}
			if Contains(cabinets, "") {
				for i, s := range cabinets {
					if s == "" {
						subjects = RemoveElement(cabinets, i)
					}
				}
			}
		}
		subjects, typesOfLessons, teachersNames, cabinets = FixSameSubjectParameters(subjects, typesOfLessons, teachersNames, cabinets)

		collection := [][]string{
			subjects, typesOfLessons, teachersNames, cabinets,
		}

		length := len(collection[0])

		for i := 0; i < length; i++ {
			if collection[0][i] == "" && collection[1][i] == "" && collection[2][i] == "" && collection[3][i] == "" {
				collection[0] = RemoveElement(collection[0], i)
				collection[1] = RemoveElement(collection[1], i)
				collection[2] = RemoveElement(collection[2], i)
				collection[3] = RemoveElement(collection[3], i)
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
		//b := len(subjects)
		//if (b != len(typesOfLessons)) || (b!= len(teachersNames)) || (b!= len(cabinets)) || (len(typesOfLessons) != len(teachersNames)) || (len(typesOfLessons) != len(cabinets)) || (len(teachersNames) != len(cabinets)) {
		//	fmt.Println("Предметы:")
		//	for _, s := range subjects {
		//		fmt.Println(s)
		//	}
		//	fmt.Println("Тип занятий")
		//	for _, lesson := range typesOfLessons {
		//		fmt.Println(lesson)
		//	}
		//	fmt.Println("ФИО")
		//	for _, teachersName := range teachersNames {
		//		fmt.Println(teachersName)
		//	}
		//	fmt.Println("Кабинет")
		//	for _, s := range cabinets {
		//		fmt.Println(s)
		//
		//	}
		//}

	} else { // в строке нет энтеров
		if !strings.Contains(subject, "яз") { // надо где языки то что возможно запарсить а то что нет так кинуть в урок
			fmt.Println("Предметы:")
			fmt.Println(subject)
			fmt.Println("Тип занятий")
			fmt.Println(typeOfLesson)
			fmt.Println("ФИО")
			fmt.Println(teacherName)
			fmt.Println("Кабинет")
			fmt.Println(cabinet)
		} else { // надо парсить одиночные строки желательно бы начать с тех строк где несколько предметов и разелить их как-то

		}
	}
	return lessons
}

func SortDataBySubgroup(subjects, typesOfLesson, teachersName, cabinets []string) {

}

//Удаляет последний пустой элемент
func RemoveLastEmptyElement(array []string) []string {
	if array[len(array)-1] == "" {
		array = RemoveElement(array, len(array)-1)
	}
	return array
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

//Удаляет элемент из среза строк по индексу
func RemoveElement(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

//Если в строках содержатся одинаковые предметы, то подтянуть из нужной строки вид занятия или кабинет или препода или задублировать последним найденным
func FixSameSubjectParameters(subjects, typeOfLessons, teachersNames, cabinets []string) (rsubjects []string, rtypeOfLessons []string, rteachersNames []string, rcabinets []string) {
	for len(typeOfLessons) < len(subjects) {
		typeOfLessons = append(typeOfLessons, "")
	}

	for len(teachersNames) < len(subjects) {
		teachersNames = append(teachersNames, "")
	}

	for len(cabinets) < len(subjects) {
		cabinets = append(cabinets, "")
	}
	for i, subject := range subjects {
		for i2, s := range subjects[i+1:] { // i2 относителен поэтому i + i2 + 1
			reg := regexp.MustCompile("[^\\d()][А-я -]+")
			index := i2 + i + 1
			str1 := strings.ReplaceAll(strings.ToLower(strings.ReplaceAll(LongestString(reg.FindAllString(subject, -1)), " ", "")), "н", "")
			str2 := strings.ReplaceAll(strings.ToLower(strings.ReplaceAll(LongestString(reg.FindAllString(s, -1)), " ", "")), "н", "")
			if str1 == str2 { // если 2 предмета одинаковые
				if typeOfLessons[i] == "" {
					typeOfLessons[i] = typeOfLessons[index]
				}
				if typeOfLessons[index] == "" {
					typeOfLessons[index] = typeOfLessons[i]
				}

				if teachersNames[i] == "" {
					teachersNames[i] = teachersNames[index]
				}
				if teachersNames[index] == "" {
					teachersNames[index] = teachersNames[i]
				}

				if cabinets[i] == "" {
					cabinets[i] = cabinets[index]
				}
				if cabinets[index] == "" {
					cabinets[index] = cabinets[i]
				}
			}
		}
	}
	typeOfLessons = RepeatFunc(typeOfLessons)
	teachersNames = RepeatFunc(teachersNames)
	cabinets = RepeatFunc(cabinets)
	return subjects, typeOfLessons, teachersNames, cabinets
}

//Заполнение пустых элементов (дублирование)
func RepeatFunc(array []string) []string {
	flag := false
	for _, s := range array {
		if s != "" {
			flag = true
		}
	}
	if flag {
		for i := 1; i < len(array); i++ {
			if array[i] == "" {
				array[i] = array[i-1]
			}
		}

		for i := len(array) - 1; i > 0; i-- {
			if array[i] == "" {
				array[i] = array[i+1]
			}
		}
	}
	return array
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
