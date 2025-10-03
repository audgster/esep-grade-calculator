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

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 20, Assignment)
	gradeCalculator.AddGrade("exam 1", 15, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 31, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeType(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected Assignment.String() to return '%s'; got '%s' instead", "assignment", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected Exam.String() to return '%s'; got '%s' instead", "exam", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected Essay.String() to return '%s'; got '%s' instead", "essay", Essay.String())
	}
}

func TestAddGradeStores(t *testing.T) {
	gc := NewGradeCalculator()

	gc.AddGrade("assignment1", 100, Assignment)
	gc.AddGrade("exam1", 90, Exam)
	gc.AddGrade("essay1", 80, Essay)

	if len(gc.assignments) != 1 || gc.assignments[0].Type != Assignment {
		t.Errorf("Expected one assignment with type Assignment; got %+v", gc.assignments)
	}
	if len(gc.exams) != 1 || gc.exams[0].Type != Exam {
		t.Errorf("Expected one exam with type Exam; got %+v", gc.exams)
	}
	if len(gc.essays) != 1 || gc.essays[0].Type != Essay {
		t.Errorf("Expected one essay with type Essay; got %+v", gc.essays)
	}
}

