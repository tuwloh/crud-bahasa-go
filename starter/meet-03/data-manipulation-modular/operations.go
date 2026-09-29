package main

import "strings"

func updateScore(students []Student, id int, score int) bool {
	for i := range students {
		if students[i].ID == id {
			students[i].Score = score
			return true
		}
	}
	return false
}

func removeInactive(students []Student) []Student {
	activeStudents := students[:0]
	for _, student := range students {
		if student.Active {
			activeStudents = append(activeStudents, student)
		}
	}
	return activeStudents
}

func filterByMajor(students []Student, major string) []Student {
	var result []Student
	for _, student := range students {
		if strings.EqualFold(student.Major, major) {
			result = append(result, student)
		}
	}
	return result
}

func indexByID(students []Student) map[int]Student {
	studentByID := make(map[int]Student, len(students))
	for _, student := range students {
		studentByID[student.ID] = student
	}
	return studentByID
}

func averageScore(students []Student) float64 {
	if len(students) == 0 {
		return 0
	}

	total := 0
	for _, student := range students {
		total += student.Score
	}
	return float64(total) / float64(len(students))
}

func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "D"
	}
}
