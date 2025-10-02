package esepunittests

import "testing"

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

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 55, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestNoGradesIsF(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("expected F with no grades; got %s", got)
	}
}

func TestBoundaries_C_and_D(t *testing.T) {
	// Exactly 70 -> C
	gc := NewGradeCalculator()
	gc.AddGrade("a", 70, Assignment)
	gc.AddGrade("e", 70, Exam)
	gc.AddGrade("s", 70, Essay)
	if got := gc.GetFinalGrade(); got != "C" {
		t.Fatalf("want C, got %s", got)
	}

	// Exactly 60 -> D
	gc = NewGradeCalculator()
	gc.AddGrade("a", 60, Assignment)
	gc.AddGrade("e", 60, Exam)
	gc.AddGrade("s", 60, Essay)
	if got := gc.GetFinalGrade(); got != "D" {
		t.Fatalf("want D, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String wrong: %q", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String wrong: %q", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String wrong: %q", Essay.String())
	}
}
