package esepunittests

import "fmt"

type GradeCalculator struct {
	assignments []Grade
	exams       []Grade
	essays      []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		assignments: make([]Grade, 0),
		exams:       make([]Grade, 0),
		essays:      make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numerical := gc.calculateNumericalGrade()

	switch {
	case numerical >= 90:
		return "A"
	case numerical >= 80:
		return "B"
	case numerical >= 70:
		return "C"
	case numerical >= 60:
		return "D"
	default:
		return "F"
	}
}

func GetGrade(assignments, exams, essays []int) string {
	gc := NewGradeCalculator()
	for i, v := range assignments {
		gc.AddGrade(fmt.Sprintf("assignment-%d", i+1), v, Assignment)
	}
	for i, v := range exams {
		gc.AddGrade(fmt.Sprintf("exam-%d", i+1), v, Exam)
	}
	for i, v := range essays {
		gc.AddGrade(fmt.Sprintf("essay-%d", i+1), v, Essay)
	}
	return gc.GetFinalGrade()
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{Name: name, Grade: grade, Type: Assignment})
	case Exam:
		gc.exams = append(gc.exams, Grade{Name: name, Grade: grade, Type: Exam})
	case Essay:
		gc.essays = append(gc.essays, Grade{Name: name, Grade: grade, Type: Essay})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	assignmentAvg := computeAverage(gc.assignments)
	examAvg := computeAverage(gc.exams)
	essayAvg := computeAverage(gc.essays)

	return 0.50*assignmentAvg + 0.35*examAvg + 0.15*essayAvg
}

func computeAverage(grades []Grade) float64 {
	if len(grades) == 0 {
		return 0
	}
	sum := 0
	for _, g := range grades {
		sum += g.Grade
	}
	return float64(sum) / float64(len(grades))
}
