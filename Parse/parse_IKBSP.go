package Parse

import (
	. "../Structure"
	"fmt"
	"strings"
)

func ParseIKBSP(subject, typeOfLesson, teacherName, cabinet, dayOfWeek, numberLesson, week string) []Lesson {

	//someLesson := Lesson{}
	//matched, _ := regexp.Match("((\\d{2}|\\d)+((, *)|( *)|(\\s+,\\s+))){1,17}н\\s+", []byte(subject))
	//if matched {
	//	re := regexp.MustCompile("((\\d{2}|\\d)+((, *)|( *)|(\\s+,\\s+))){1,17}н\\s+")
	//	includedString := re.FindString(subject)
	//	fmt.Println(subject)
	//	fmt.Println(includedString)
	//	for re.MatchString(subject){
	//
	//	}
	//	for i, _ := range someLesson.OccurrenceLesson{
	//		someLesson.OccurrenceLesson[i] = true
	//	}
	//}
	//for _,v := range NewLineSeparator(subject){
	//	fmt.Println(v)
	//	fmt.Println("----------")
	//}
	fmt.Println(subject)
	return []Lesson{NewLesson(), NewLesson()}
}
func NewLineSeparator(line string) []string {
	line = strings.ReplaceAll(line, "/", "\n")
	return strings.Split(line, "\n")
}
