package esepunittests

import "testing"

func TestPassFail_DefaultCutoff70(t *testing.T) {
	gc := NewGradeCalculator()
	gc.passFailMode = true // flip internal switch
	gc.passCutoff = 70     // default per assignment

	addAll(gc, 72) // clearly passes cutoff
	if got := gc.GetFinalGrade(); got != "Pass" {
		t.Errorf("expected Pass, got %q", got)
	}

	gc2 := NewGradeCalculator()
	gc2.passFailMode = true
	gc2.passCutoff = 70

	addAll(gc2, 68) // clearly below cutoff
	if got := gc2.GetFinalGrade(); got != "Fail" {
		t.Errorf("expected Fail, got %q", got)
	}
}

func TestPassFail_CustomCutoff(t *testing.T) {
	gc := NewGradeCalculator()
	gc.passFailMode = true
	gc.passCutoff = 85 // stricter

	addAll(gc, 84)
	if got := gc.GetFinalGrade(); got != "Fail" {
		t.Errorf("with cutoff 85, expected Fail for 84, got %q", got)
	}

	addAll(gc, 86) // add more grades at/above cutoff — weighted avg still >= 85
	if got := gc.GetFinalGrade(); got != "Pass" {
		t.Errorf("with cutoff 85, expected Pass for 86, got %q", got)
	}
}

func TestLetterModeUnaffected(t *testing.T) {
	// sanity check: without passFailMode, we still get letter grades
	gc := NewGradeCalculator()
	addAll(gc, 90)
	if got := gc.GetFinalGrade(); got != "A" {
		t.Errorf("expected A in letter mode, got %q", got)
	}
}
