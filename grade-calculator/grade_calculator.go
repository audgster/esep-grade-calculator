package esepunittests

type GradeType int

type GradingMode int

const (
	ModeLetter GradingMode = iota
	ModePassFail
)

const (
	Assignment GradeType = iota
	Exam
	Essay
)

func (gt GradeType) String() string {
	switch gt {
	case Assignment:
		return "assignment"
	case Exam:
		return "exam"
	case Essay:
		return "essay"
	default:
		return "unknown"
	}
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

type GradeCalculator struct {
	grades []Grade
	mode   GradingMode
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
		mode:   ModeLetter,
	}
}

func NewGradeCalculatorPassFail() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
		mode:   ModePassFail,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.mode == ModePassFail {
		if numericalGrade >= 70 {
			return "Pass"
		}
		return "Fail"
	}

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverageByType(gc.grades, Assignment)
	exam_average := computeAverageByType(gc.grades, Exam)
	essay_average := computeAverageByType(gc.grades, Essay)

	weighted_grade := float64(assignment_average)*0.5 + float64(exam_average)*0.35 + float64(essay_average)*0.15

	return int(weighted_grade)
}

func computeAverageByType(all []Grade, t GradeType) int {
	filtered := make([]Grade, 0)
	for _, grade := range all {
		if grade.Type == t {
			filtered = append(filtered, grade)
		}
	}

	return computeAverage(filtered)
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0
	}

	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}
