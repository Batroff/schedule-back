package parse

import (
	"fmt"
	. "schedule/structure"
	"strings"
)

//случаи
//#1
//1,5,9,13 н. Цифровая обработка оптических сигналов (1 п/г)
//3,7,11,15 н. Цифровая обработка оптических сигналов (2 п/г)
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//лаб
//лаб
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//Маняк А.П.
//Маняк А.П.
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//В-401-1
//В-401-1
//#2
//2,6 н. Системы автоматизированного проектирования и виртуальные приборы в оптотехнике (1 п/г)
//4,8 н. Системы автоматизированного проектирования и виртуальные приборы в оптотехнике (2 п/г)
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//лаб
//лаб
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//Кретушев А.В.
//Кретушев А.В.
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//В-401-1
//#3
//3,7,11,15н Метрология, стандартизация и сертификация
//9,13 н ТКМ, 1 гр / 2 гр
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//лр
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//Чернова А.В.
//Зуев В.В./Баранова Н.С.
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//В-78* В-419
//254/244
//#4
//2-8 н Теория соединения материалов
//10,14н-1гр 12,16н-2 гр Тепл. проц. в ТС
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//пр
//лр
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//Кудрявцев И.В.
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//242
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//КМБО-05-20
//3,7,11,15 н. Правоведение
//5 н. Методы и стандарты программирования (2 п/г)
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//пр
//пр
//
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//Леонова С.Л.
//Смирнов А.В.
//
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//А-203
//Б-209
//
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//ЧЕТВЕРГ
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//4
//Ин.яз(1 подгр.)/Ин.яз. 2 подгр
//Ин. язык (нем.)
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//пр
//пр
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//Курсевич Д.В./ Каппушева И.Ш.
//Гриценко С.А.
//~~~~~~~~~~~~~~~~~~~~~~~~~~
//И-303
//И-311
//Б-402
//надо удалить энтеры после этого
//В-78*
//В-419

func SubGroupParse(subject, typeOfLesson, teacherName, cabinet, dayOfWeek, numberLesson, week string) []Lesson {
	//несколько уроков в 1 дне надо раскидать по строкам и если одинаковые предметы почему они раскинуты(тип работы/преподы)
	var lessons []Lesson
	if strings.Contains(subject, "\n") { // если в строчке с предметом более 1 строки
		subjects := strings.Split(subject, "\n")
		typesOfLessons := strings.Split(typeOfLesson, "\n")
		teachersNames := strings.Split(teacherName, "\n")
		cabinets := strings.Split(cabinet, "\n")
		max := Max(len(subjects), len(typesOfLessons), len(teachersNames), len(cabinets)) // максимальное число строк в ячейке
		collection := [][]string{
			subjects, typesOfLessons, teachersNames, cabinets,
		}
		if Contains(subjects, "…………………") {
			for i, s := range subjects {
				if s == "…………………" {
					RemoveElement(subjects, i)
				}
			}
		}
		flag := false // количество строк в каждой ячейке неодинаковое
		for _, strings1 := range collection {
			if len(strings1) != max {
				flag = true // в какой-то ячейке кол-во строк неодинаковое
			}
		}
		if flag {
			fmt.Println(subject)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(typeOfLesson)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(teacherName)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(cabinet)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(dayOfWeek)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
			fmt.Println(numberLesson)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~")
		} else { //раскидать по урокам дублировав необходимые значения и передать в дальнейший парс
			for i := 0; i < max; i++ {
				if collection[0][i] == "" && collection[1][i] == "" && collection[2][i] == "" && collection[3][i] == "" {
					collection[0] = RemoveElement(collection[0], i)
					collection[1] = RemoveElement(collection[1], i)
					collection[2] = RemoveElement(collection[2], i)
					collection[3] = RemoveElement(collection[3], i)
					max--
					continue
				}
				someLesson := NewLesson()
				someLesson.Subject = collection[0][i]
				someLesson.TypeOfLesson = collection[1][i]
				someLesson.TeacherName = collection[2][i]
				someLesson.Cabinet = collection[3][i]
				lessons = append(lessons, someLesson) // массив с уроками "предмет с/без п/г" "тип" "фио" "кабинет"
			}

			//for i, lesson := range lessons {
			//	fmt.Println(lesson.Subject)
			//	fmt.Println(lesson.TypeOfLesson)
			//	fmt.Println(lesson.TeacherName)
			//	fmt.Println(lesson.Cabinet)
			//	fmt.Println(i)
			//	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
			//}
		}
	}
	return []Lesson{NewLesson(), NewLesson()}
}

func SortDataBySubgroup(subjects, typesOfLesson, teachersName, cabinets []string) {

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

//func CheckForEmptyElements(array [][]string) [][]string { // очистка массива от пустых элементов
//	flag := false
//	for i, s := range array {
//		for i2, s2 := range s {
//			if s2 == "" {
//				flag = true
//			} else if flag{
//				return array
//			}
//		}
//	}
//	for i, i2 := range array {
//
//	}
//	return array
//}

func RemoveElement(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}
