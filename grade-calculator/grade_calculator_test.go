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
    expected_value := "A" 

    gradeCalculator := NewGradeCalculator()
    gradeCalculator.AddGrade("open source assignment", 100, Assignment)
    gradeCalculator.AddGrade("exam 1", 95, Exam)
    gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

    actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
        t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    }
}

func TestGetGradeC(t *testing.T) {
	expected := "C"
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 75, Assignment)
	gc.AddGrade("e1", 75, Exam)
	gc.AddGrade("es1", 75, Essay)
	if got := gc.GetFinalGrade(); got != expected {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

func TestGetGradeD(t *testing.T) {
	expected := "D"
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 65, Assignment)
	gc.AddGrade("e1", 65, Exam)
	gc.AddGrade("es1", 65, Essay)
	if got := gc.GetFinalGrade(); got != expected {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

func TestGetGradeFProper(t *testing.T) {
	expected := "F"
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 55, Assignment)
	gc.AddGrade("e1", 55, Exam)
	gc.AddGrade("es1", 55, Essay)
	if got := gc.GetFinalGrade(); got != expected {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

func TestComputeAverage_EmptyInputsYieldF(t *testing.T) {
	gc := NewGradeCalculator()
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("empty inputs should yield F, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String() = %q", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String() = %q", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String() = %q", Essay.String())
	}
}

func TestGetGrade_FreeFunctionAdapter(t *testing.T) {
	assignments := []int{100}
	exams := []int{80}
	essays := []int{70}
	if got := GetGrade(assignments, exams, essays); got != "B" {
		t.Fatalf("expected B from free function, got %s", got)
	}
}
