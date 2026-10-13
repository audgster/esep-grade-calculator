# ESEP Grade Calculator

This library is a simple utility to calculate your grades. It stores a student's assignment, exam, and essay grades, weighs them and returns a final letter grade.

## Files

- go.mod: Defines the Go module and sets the Go version to 1.24.4.
- grade_calculator.go: Defines GradeCalculator, which stores a student's grades. AddGrade adds a grade to the matching list. GetFinalGrade weighs them by 50/35/15 and returns the final grade.
- grade_calculator_test.go: Has three unit tests. Each one builds a calculator, adds sample grades, and checks the letter returned.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```