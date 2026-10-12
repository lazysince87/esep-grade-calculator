package esepunittests

import (
	"strconv"
	"testing"
)

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestLetterGradeBoundaries(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{100, "A"}, {90, "A"}, {89, "B"}, {80, "B"}, {79, "C"},
		{70, "C"}, {69, "D"}, {60, "D"}, {59, "F"}, {0, "F"},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.score), func(t *testing.T) {
			gc := NewGradeCalculator()
			gc.AddGrade("assignment", tt.score, Assignment)
			gc.AddGrade("exam", tt.score, Exam)
			gc.AddGrade("essay", tt.score, Essay)
			if got := gc.GetFinalGrade(); got != tt.want {
				t.Errorf("score %d: want %s, got %s", tt.score, tt.want, got)
			}
		})
	}
}

func TestCategoryAveragesAndWeights(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment 1", 60, Assignment)
	gc.AddGrade("assignment 2", 80, Assignment)
	gc.AddGrade("exam 1", 100, Exam)
	gc.AddGrade("essay 1", 0, Essay)
	// 70*0.50 + 100*0.35 + 0*0.15 = 70 (C).
	// Reusing the exam for essays would produce 85 (B).
	if got := gc.GetFinalGrade(); got != "C" {
		t.Errorf("want C from weighted category averages, got %s", got)
	}
}

func TestFractionalCategoryAverage(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("assignment 1", 89, Assignment)
	gc.AddGrade("assignment 2", 90, Assignment)
	gc.AddGrade("exam", 91, Exam)
	gc.AddGrade("essay", 89, Essay)
	// 89.5*0.50 + 91*0.35 + 89*0.15 = 89.95, below A.
	if got := gc.GetFinalGrade(); got != "B" {
		t.Errorf("want B below the 90 cutoff, got %s", got)
	}

	gc = NewGradeCalculator()
	gc.AddGrade("assignment 1", 89, Assignment)
	gc.AddGrade("assignment 2", 90, Assignment)
	gc.AddGrade("exam", 90, Exam)
	gc.AddGrade("essay", 92, Essay)
	// 89.5*0.50 + 90*0.35 + 92*0.15 = 90.05 (A).
	// Truncating the category average first would incorrectly give B.
	if got := gc.GetFinalGrade(); got != "A" {
		t.Errorf("want A with fractional category averages preserved, got %s", got)
	}
}

func TestEmptyAndMissingCategories(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("empty calculator: want F, got %s", got)
	}
	gc.AddGrade("assignment", 100, Assignment)
	// An absent category contributes zero; weights are not redistributed.
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("assignments only: want F for weighted score 50, got %s", got)
	}
	gc.AddGrade("unknown type", 100, GradeType(99))
	if got := gc.GetFinalGrade(); got != "F" {
		t.Errorf("unknown grade type should be ignored, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	tests := []struct {
		kind GradeType
		want string
	}{
		{Assignment, "assignment"}, {Exam, "exam"}, {Essay, "essay"}, {GradeType(99), ""},
	}
	for _, tt := range tests {
		if got := tt.kind.String(); got != tt.want {
			t.Errorf("type %d: want %q, got %q", tt.kind, tt.want, got)
		}
	}
}
