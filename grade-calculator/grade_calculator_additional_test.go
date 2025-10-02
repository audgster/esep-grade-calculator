package esepunittests

import "testing"

// helper to add the same score across all three categories
func addAll(gc *GradeCalculator, v int) {
	gc.AddGrade("a1", v, Assignment)
	gc.AddGrade("e1", v, Exam)
	gc.AddGrade("s1", v, Essay)
}

func TestGetGradeExactCutoffs(t *testing.T) {
	tests := []struct {
		name string
		v    int
		want string
	}{
		{"GradeA_Exact90", 90, "A"},
		{"GradeB_Exact89", 89, "B"},
		{"GradeB_Exact80", 80, "B"},
		{"GradeC_Exact70", 70, "C"},
		{"GradeD_Exact60", 60, "D"},
		{"GradeF_Exact59", 59, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gc := NewGradeCalculator()
			addAll(gc, tt.v) // ensure weighting doesn't distort the cutoff
			got := gc.GetFinalGrade()
			if got != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestComputeAverageEmptySlice(t *testing.T) {
	if got := computeAverage([]Grade{}); got != 0 {
		t.Errorf("expected 0 for empty slice, got %d", got)
	}
}

func TestComputeAverageSmallSlice(t *testing.T) {
	sample := []Grade{{Grade: 80}, {Grade: 90}, {Grade: 100}} // avg = 90
	if got := computeAverage(sample); got != 90 {
		t.Errorf("expected 90, got %d", got)
	}
}
