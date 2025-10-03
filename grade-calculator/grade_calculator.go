package esepunittests

type GradeCalculator struct {
	grades   []Grade
	mapScale func(n float64) string
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
		grades:   make([]Grade, 0),
		mapScale: letterScale,
	}
}

func NewGradeCalculatorWithPassFail(passCutoff float64) *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
		mapScale: func(n float64) string {
			if n >= passCutoff {
				return "P"
			}
			return "F"
		},
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numerical := gc.calculateNumericalGrade()
	return gc.mapScale(numerical)
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.grades = append(gc.grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

func letterScale(n float64) string {
	switch {
	case n >= 90:
		return "A"
	case n >= 80:
		return "B"
	case n >= 70:
		return "C"
	case n >= 60:
		return "D"
	default:
		return "F"
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() float64 {
	assignmentAvg := gc.avgByType(Assignment)
	examAvg := gc.avgByType(Exam)
	essayAvg := gc.avgByType(Essay)
	return assignmentAvg*0.50 + examAvg*0.35 + essayAvg*0.15
}

func (gc *GradeCalculator) avgByType(t GradeType) float64 {
	sum, count := 0, 0
	for _, g := range gc.grades {
		if g.Type == t {
			sum += g.Grade
			count++
		}
	}
	if count == 0 {
		return 0.0
	}
	return float64(sum) / float64(count)
}

