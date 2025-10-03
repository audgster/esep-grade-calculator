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

// Wrapper to match tests that call GetGrade()
func (gc *GradeCalculator) GetGrade() string {
	return gc.GetFinalGrade()
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	switch {
	case numericalGrade >= 90.0:
		return "A"
	case numericalGrade >= 80.0:
		return "B"
	case numericalGrade >= 70.0:
		return "C"
	case numericalGrade >= 60.0:
		return "D"
	default:
		return "F"
	}
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
	assignmentAverage := computeAverage(gc.assignments)
	examAverage := computeAverage(gc.exams)
	essayAverage := computeAverage(gc.essays) // <-- fixed to essays

	// Required weights: Assignments 50%, Exams 35%, Essays 15%
	return 0.50*assignmentAverage + 0.35*examAverage + 0.15*essayAverage
}

func computeAverage(grades []Grade) float64 {
	if len(grades) == 0 {
		return 0.0
	}
	sum := 0
	for _, g := range grades { // use the value, not the index
		sum += g.Grade
	}
	return float64(sum) / float64(len(grades)) // float division (no truncation)
}
