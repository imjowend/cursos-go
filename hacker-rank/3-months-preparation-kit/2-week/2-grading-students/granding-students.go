package main

import "fmt"

const (
	failingGrade = 38
)

func main() {
	var grades []int32

	grades = append(grades, 38, 84, 29, 57)
	fmt.Println(gradingStudents(grades))

}

func gradingStudents(grades []int32) []int32 {

	var newGradeNotes []int32

	for i := 0; i < len(grades); i++ {
		newGrade := grades[i]
		if grades[i] >= failingGrade {
			if grades[i]%5 >= 3 {
				newGrade = ((grades[i] / 5) + 1) * 5
			}
		}
		newGradeNotes = append(newGradeNotes, newGrade)
	}

	return newGradeNotes
}
