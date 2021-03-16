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

func SubGroupParse(subject, typeOfLesson, teacherName, cabinet, dayOfWeek, numberLesson, week string) []Lesson {
	//несколько уроков в 1 дне надо раскидать по строкам и если одинаковые предметы почему они раскинуты(тип работы/преподы)
	if strings.Contains(subject, "\n") {
		subjects := strings.Split(subject, "\n")
		typesOfLessons := strings.Split(typeOfLesson, "\n")
		teachersNames := strings.Split(teacherName, "\n")
		cabinets := strings.Split(cabinet, "\n")
		length := len(subjects) + len(typesOfLessons) + len(teachersNames) + len(cabinets)
		collection := [][]string{
			subjects, typesOfLessons, teachersNames, cabinets,
		}
		flag := false
		for _, strings1 := range collection {
			if len(strings1) != length/4 {
				flag = true
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
		} //раскидать по урокам дублировав необходимые значения и передать в дальнейший парс
	}
	return []Lesson{NewLesson(), NewLesson()}
}
