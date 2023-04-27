package userservice

import (
	"testing"

	"bookingrooms/internal/models"
)

func Test_GetSumGrade_A_A_A_Equal_A(t *testing.T) {
	expected := "A"
	grade := models.UserGrade{AGrade: "A", BGrade: "A", CGrade: "A"}
	s := service{}
	actual := s.GetSumGrade(grade)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_GetSumGrade_C_C_D_Equal_C(t *testing.T) {
	expected := "C"
	grade := models.UserGrade{AGrade: "C", BGrade: "C", CGrade: "D"}
	s := service{}
	actual := s.GetSumGrade(grade)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_GetSumGrade_C_D_D_Equal_D(t *testing.T) {
	expected := "D"
	grade := models.UserGrade{AGrade: "C", BGrade: "D", CGrade: "D"}
	s := service{}
	actual := s.GetSumGrade(grade)

	if expected != actual {
		t.Errorf("Expect %s but got %s", expected, actual)
	}
}

func Test_CalculateSumGrade_A_Equal_4(t *testing.T) {
	expected := 4
	grade := "A"
	s := service{}
	actual := s.calculateSumGrade(grade)

	if expected != actual {
		t.Errorf("Expect %d but got %d", expected, actual)
	}
}

func Test_CalculateSumGrade_B_Equal_2(t *testing.T) {
	expected := 2
	grade := "B"
	s := service{}
	actual := s.calculateSumGrade(grade)

	if expected != actual {
		t.Errorf("Expect %d but got %d", expected, actual)
	}
}

func Test_CalculateSumGrade_C_Equal_1(t *testing.T) {
	expected := 1
	grade := "C"
	s := service{}
	actual := s.calculateSumGrade(grade)

	if expected != actual {
		t.Errorf("Expect %d but got %d", expected, actual)
	}
}

func Test_CalculateSumGrade_D_Equal_0(t *testing.T) {
	expected := 0
	grade := "D"
	s := service{}
	actual := s.calculateSumGrade(grade)

	if expected != actual {
		t.Errorf("Expect %d but got %d", expected, actual)
	}
}
