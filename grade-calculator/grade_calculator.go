package esepunittests

type GradeCalculator struct {
	grades []Grade
	mode   EvaluationMode
}

// EvaluationMode selects the output of GetFinalGrade.
type EvaluationMode int

const (
	Letter EvaluationMode = iota
	PassFail
)

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

// NewGradeCalculator defaults to letter grades. Supply PassFail for pass/fail
// output. Existing calls without arguments remain valid; supply at most one mode.
func NewGradeCalculator(mode ...EvaluationMode) *GradeCalculator {
	evaluationMode := Letter
	if len(mode) > 0 {
		evaluationMode = mode[0]
	}
	return &GradeCalculator{
		grades: make([]Grade, 0),
		mode:   evaluationMode,
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if gc.mode == PassFail {
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
	switch gradeType {
	case Assignment, Exam, Essay:
		gc.grades = append(gc.grades, Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average := computeAverage(gc.grades, Assignment)
	exam_average := computeAverage(gc.grades, Exam)
	essay_average := computeAverage(gc.grades, Essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade, gradeType GradeType) float64 {
	sum := 0
	count := 0

	for _, grade := range grades {
		if grade.Type == gradeType {
			sum += grade.Grade
			count++
		}
	}

	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}
