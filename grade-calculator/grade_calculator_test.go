package esepunittests

import "testing"

// basic sanity check: all 100s should be an A
func TestGetGradeA_AllHundreds(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 100, Assignment)
	gc.AddGrade("e1", 100, Exam)
	gc.AddGrade("s1", 100, Essay)

	got := gc.GetFinalGrade()
	if got != "A" {
		t.Errorf("expected A, got %s", got)
	}
}

// numbers around low 80s should average to a B with the given weights
func TestGetGradeB_MixedLow80s(t *testing.T) {
	// 0.5*80 + 0.35*81 + 0.15*85 = 81.1 → B
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 80, Assignment)
	gc.AddGrade("e1", 81, Exam)
	gc.AddGrade("s1", 85, Essay)

	got := gc.GetFinalGrade()
	if got != "B" {
		t.Errorf("expected B, got %s", got)
	}
}

// this used to be tested as F, but the math is actually ~96.9 which is an A
func TestGetGrade_ShouldBeA_NotF(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 100, Assignment)
	gc.AddGrade("e1", 95, Exam)
	gc.AddGrade("s1", 91, Essay)

	got := gc.GetFinalGrade()
	if got != "A" {
		t.Errorf("expected A, got %s", got)
	}
}

// quick boundary checks for the >= rules
func TestBoundaries(t *testing.T) {
	t.Run("exactly 90 is A", func(t *testing.T) {
		gc := NewGradeCalculator()
		gc.AddGrade("a1", 90, Assignment)
		gc.AddGrade("e1", 90, Exam)
		gc.AddGrade("s1", 90, Essay)
		if g := gc.GetFinalGrade(); g != "A" {
			t.Errorf("expected A at 90, got %s", g)
		}
	})

	t.Run("exactly 80 is B", func(t *testing.T) {
		gc := NewGradeCalculator()
		gc.AddGrade("a1", 80, Assignment)
		gc.AddGrade("e1", 80, Exam)
		gc.AddGrade("s1", 80, Essay)
		if g := gc.GetFinalGrade(); g != "B" {
			t.Errorf("expected B at 80, got %s", g)
		}
	})
}

// no essays entered → essay avg should be 0, not a crash
func TestEmptyEssays_CountsAsZero(t *testing.T) {
	// 0.5*100 + 0.35*100 + 0.15*0 = 85 → B
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 100, Assignment)
	gc.AddGrade("e1", 100, Exam)

	if g := gc.GetFinalGrade(); g != "B" {
		t.Errorf("expected B with no essays, got %s", g)
	}
}

// just making sure the String() names are wired up
func TestGradeTypeString_NotEmpty(t *testing.T) {
	if Assignment.String() == "" || Exam.String() == "" || Essay.String() == "" {
		t.Fatalf("expected non-empty names for grade types")
	}
}

func TestGetGradeC_AllSeventies(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 70, Assignment)
	gradeCalculator.AddGrade("e1", 70, Exam)
	gradeCalculator.AddGrade("s1", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD_AllSixties(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 60, Assignment)
	gradeCalculator.AddGrade("e1", 60, Exam)
	gradeCalculator.AddGrade("s1", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF_AllZeros(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.AddGrade("a1", 0, Assignment)
	gradeCalculator.AddGrade("e1", 0, Exam)
	gradeCalculator.AddGrade("s1", 0, Essay)

	actual_value := gradeCalculator.GetFinalGrade()
	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

