package Parse

import (
	. "Schedule/Structure"
	"fmt"
	"regexp"
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
	// flag на кр regexp на кр
	if regexp.MustCompile("((^)|( ))кр((\\.)|(  ??))").MatchString(subject) {
		//flag := true
	}
	if regexp.MustCompile("^((\\d{2}|\\d)( *)|( *?, *?)|( *?- *?)){1,17}").MatchString(subject) {
		loc := regexp.MustCompile("^((\\d{2}|\\d)( *)|( *?, *?)|( *?- *?)){1,17}").FindStringIndex(subject)
		string := subject[loc[0]:loc[1]]
		fmt.Println(string)
		//fmt.Println(subject)
	}

	//fmt.Println(subject)
	return []Lesson{NewLesson(), NewLesson()}
}
func NewLineSeparator(line string) []string {
	line = strings.ReplaceAll(line, "/", "\n")
	return strings.Split(line, "\n")
}
