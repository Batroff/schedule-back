package Parse

import (
	. "Schedule/Structure"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ParseIKBSP(subject, typeOfLesson, teacherName, cabinet, dayOfWeek, numberLesson, week string) []Lesson {

	someLesson := NewLesson()
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

	if regexp.MustCompile("(((\\d{2}|\\d)( *[ \\-,] *)){1,17}(\\d{2}|\\d))|(\\d{2}|\\d)").MatchString(subject) {
		flag := exceptFlag(subject)
		someLesson.FillInWeeks(flag, week)
		loc := regexp.MustCompile("(((\\d{2}|\\d)( *[ \\-,] *)){1,17}(\\d{2}|\\d))|(\\d{2}|\\d)").FindStringIndex(subject)
		fmt.Println(subject)
		string := subject[loc[0]:loc[1]]
		string = removeSpaces(string)
		fmt.Println(string)
		fmt.Println(week)
		for _, v := range numbersPresent(string) {
			if flag && strings.Contains(week, "II") && (v-1)%2 == 0 {
				someLesson.OccurrenceLesson[v-1] = false
			} else if flag && strings.Contains(week, "I") && (v-1)%2 == 0 {
				someLesson.OccurrenceLesson[v-1] = false
			} else if !flag && strings.Contains(week, "II") && (v-1)%2 == 0 {
				someLesson.OccurrenceLesson[v-1] = true
			} else if !flag && strings.Contains(week, "I") && (v-1)%2 == 0 {
				someLesson.OccurrenceLesson[v-1] = true
			}
		}
		for _, v := range someLesson.OccurrenceLesson {
			fmt.Print(v)
			fmt.Print(" ")
		}
		fmt.Println()
	}

	//fmt.Println(subject)
	return []Lesson{NewLesson(), NewLesson()}
}
func numbersPresent(subject string) []int { // str - название предмета
	stringNumbers := orSplit(subject)
	intNumbers := []int{}
	for _, v := range stringNumbers {
		if strings.Contains(v, "-") {
			firstNum, err := strconv.Atoi(v[0:strings.Index(v, "-")])
			if err != nil {
				fmt.Println(err)
			}
			lastNum, err := strconv.Atoi(v[strings.Index(v, "-")+1 : len(v)])
			if err != nil {
				fmt.Println(err)
			}
			for i := firstNum; i <= lastNum; i++ {
				intNumbers = append(intNumbers, i)
			}
		} else {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
			}
			intNumbers = append(intNumbers, num)
		}
	}
	return intNumbers
}
func orSplit(subject string) []string {
	if strings.Contains(subject, ",") {
		return strings.Split(subject, ",")
	} else {
		return strings.Split(subject, " ")
	}
}
func removeSpaces(subject string) string {
	return strings.ReplaceAll(subject, " ", "")
}
func exceptFlag(subject string) bool {
	flag := false
	if regexp.MustCompile("((^)|( ))кр((\\.)|(  ??))").MatchString(subject) {
		flag = true
	}
	return flag
}
func NewLineSeparator(line string) []string {
	line = strings.ReplaceAll(line, "/", "\n")
	return strings.Split(line, "\n")
}
