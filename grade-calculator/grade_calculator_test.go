package esepunittests

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetGradeA(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("open source assignment", 100, Assignment)
	g.AddGrade("exam 1", 100, Exam)
	g.AddGrade("essay on ai ethics", 100, Essay)
	if got := g.GetFinalGrade(); got != "A" {
		t.Fatalf("GetFinalGrade() = %q; want %q", got, "A")
	}
}

func TestGetGradeB(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("open source assignment", 80, Assignment)
	g.AddGrade("exam 1", 81, Exam)
	g.AddGrade("essay on ai ethics", 85, Essay)
	if got := g.GetFinalGrade(); got != "B" {
		t.Fatalf("GetFinalGrade() = %q; want %q", got, "B")
	}
}

func TestGetGradeF(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("open source assignment", 100, Assignment) // 50%
	g.AddGrade("exam 1", 0, Exam)                         // 35%
	g.AddGrade("essay on ai ethics", 0, Essay)            // 15%
	if got := g.GetFinalGrade(); got != "F" {
		t.Fatalf("GetFinalGrade() = %q; want %q", got, "F")
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() == "" || Exam.String() == "" || Essay.String() == "" {
		t.Fatalf("GradeType.String() returned empty for a valid type")
	}
}

func TestGradeTypeString_viaFmt(t *testing.T) {
	s := fmt.Sprintf("%s-%s-%s", Assignment, Exam, Essay)
	if !strings.Contains(s, "assignment") || !strings.Contains(s, "exam") || !strings.Contains(s, "essay") {
		t.Fatalf("unexpected String() formatting: %q", s)
	}
}

func TestGetFinalGrade_Boundaries(t *testing.T) {
	type tc struct {
		score int
		want  string
	}
	cases := []tc{
		{score: 90, want: "A"},
		{score: 80, want: "B"},
		{score: 70, want: "C"},
		{score: 60, want: "D"},
		{score: 50, want: "F"},
	}
	for _, c := range cases {
		g := NewGradeCalculator()
		g.AddGrade("a1", c.score, Assignment)
		g.AddGrade("e1", c.score, Exam)
		g.AddGrade("s1", c.score, Essay)
		if got := g.GetFinalGrade(); got != c.want {
			t.Fatalf("for score %d got %q; want %q", c.score, got, c.want)
		}
	}
}

func TestComputeAverage_EmptyAndNonEmpty(t *testing.T) {
	if got := computeAverage([]Grade{}); got != 0.0 {
		t.Fatalf("computeAverage(empty) = %v, want 0", got)
	}
	got := computeAverage([]Grade{
		{Name: "x", Grade: 100, Type: Assignment},
		{Name: "y", Grade: 80, Type: Assignment},
	})
	if got != 90.0 {
		t.Fatalf("computeAverage(non-empty) = %v, want %v", got, 90.0)
	}
}
