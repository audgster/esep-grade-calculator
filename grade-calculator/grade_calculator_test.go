package esepunittests

import (
	"fmt"
	"testing"
)

// Ensure the wrapper is executed (lines in GetGrade)
func TestGetGradeWrapper(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 90, Assignment)
	gc.AddGrade("e1", 90, Exam)
	gc.AddGrade("es1", 90, Essay)
	// Call the wrapper specifically
	got := gc.GetGrade()
	if got != "A" {
		t.Fatalf("GetGrade wrapper: want A, got %s", got)
	}
}

// Cover GradeType.String() for all enum values
func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String = %q", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String = %q", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String = %q", Essay.String())
	}
}

// Drive computeAverage len==0 for *other* categories too
func TestEmptyCategoriesComputeAverage(t *testing.T) {
	// No assignments (avg=0), exams & essays present
	// final = 0.50*0 + 0.35*90 + 0.15*90 = 0 + 31.5 + 13.5 = 45 -> F
	{
		gc := NewGradeCalculator()
		gc.AddGrade("e1", 90, Exam)
		gc.AddGrade("es1", 90, Essay)
		if got := gc.GetFinalGrade(); got != "F" {
			t.Fatalf("No assignments: want F, got %s", got)
		}
	}

	// No exams (avg=0), assignments & essays present
	// final = 0.50*90 + 0.35*0 + 0.15*90 = 45 + 0 + 13.5 = 58.5 -> F
	{
		gc := NewGradeCalculator()
		gc.AddGrade("a1", 90, Assignment)
		gc.AddGrade("es1", 90, Essay)
		if got := gc.GetFinalGrade(); got != "F" {
			t.Fatalf("No exams: want F, got %s", got)
		}
	}
}

// (Optional) Explicitly hit each GetFinalGrade branch again (A,B,C,D,F)
// You already cover many via other tests, but this guarantees 100%.
func TestAllLetterBranches(t *testing.T) {
	type ints = []int
	cases := []struct {
		name   string
		as, ex, es ints
		want   string
	}{
		{"A_>=90", ints{90}, ints{90}, ints{90}, "A"},
		{"B_>=80", ints{80}, ints{80}, ints{80}, "B"},
		{"C_>=70", ints{70}, ints{70}, ints{70}, "C"},
		{"D_>=60", ints{60}, ints{60}, ints{60}, "D"},
		{"F_<60",  ints{59}, ints{59}, ints{59}, "F"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gc := NewGradeCalculator()
			for i, v := range tc.as { gc.AddGrade(fmt.Sprintf("A%d", i), v, Assignment) }
			for i, v := range tc.ex { gc.AddGrade(fmt.Sprintf("E%d", i), v, Exam) }
			for i, v := range tc.es { gc.AddGrade(fmt.Sprintf("Es%d", i), v, Essay) }
			if got := gc.GetFinalGrade(); got != tc.want {
				t.Fatalf("%s: want %s, got %s", tc.name, tc.want, got)
			}
		})
	}
}
