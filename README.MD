# ESEP Grade Calculator

ESEP Grade Calculator is a small Go library that computes a student’s final letter grade from assignments (50%), exams (35%), and essays (15%) on a 0–100 scale, mapping to A ≥ 90, B ≥ 80, C ≥ 70, D ≥ 60, F < 60 (empty categories contribute 0). To run tests on Linux: cd grade-calculator && go test ./...; for coverage: go test -coverprofile=coverage.txt ./... && go tool cover -func=coverage.txt (optional HTML: go tool cover -html=coverage.txt -o coverage.html). Basic usage: create a calculator with NewGradeCalculator(), add scores with AddGrade(name, score, Assignment|Exam|Essay), then call GetFinalGrade() to get the final letter.

This library is a simple utility to calculate your grades.

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
