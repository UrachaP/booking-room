package userservice

import (
	"bookingrooms/internal/models"
)

func (s service) GetSumGrade(g models.UserGrade) string {
	var sum int
	sum += s.calculateSumGrade(g.AGrade)
	sum += s.calculateSumGrade(g.BGrade)
	sum += s.calculateSumGrade(g.CGrade)

	var grade string
	if sum >= 9 {
		grade = "A"
	} else if sum >= 5 {
		grade = "B"
	} else if sum >= 2 {
		grade = "C"
	} else {
		grade = "D"
	}
	return grade
}

func (s service) calculateSumGrade(grade string) int {
	var sum int
	if grade == "A" {
		sum = 4
	} else if grade == "B" {
		sum = 2
	} else if grade == "C" {
		sum = 1
	} else if grade == "D" {
		sum = 0
	}
	return sum
}
