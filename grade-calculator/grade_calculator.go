package esepunittests

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
	numericalGrade := gc.calculateNumericalGrade()

	// A >=90, B >=80, C >=70, D >=60, F <60
	switch {
	case numericalGrade >= 90:
		return "A"
	case numericalGrade >= 80:
		return "B"
	case numericalGrade >= 70:
		return "C"
	case numericalGrade >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.assignments = append(gc.assignments, Grade{
			Name:  name,
			Grade: grade,
			Type:  Assignment,
		})
	case Exam:
		gc.exams = append(gc.exams, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})
	case Essay:
		gc.essays = append(gc.essays, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	assignmentAvg := computeAverage(gc.assignments)
	examAvg := computeAverage(gc.exams)
	essayAvg := computeAverage(gc.essays)

	return assignmentAvg*0.50 + examAvg*0.35 + essayAvg*0.15
}

func computeAverage(grades []Grade) float64 {
	if len(grades) == 0 {
		return 0.0 
	}
	sum := 0
	for _, g := range grades { // sum  grade value
		sum += g.Grade
	}
	return float64(sum) / float64(len(grades))
}

