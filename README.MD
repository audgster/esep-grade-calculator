# ESEP Grade Calculator

A tiny Go library that computes final course grades from assignments, exams, and essays.

## What this does (quick summary)

- Stores individual grades (assignment/exam/essay)
- Computes average per category
- Applies weights to produce a final numeric score
- Maps the numeric score to a letter grade (A/B/C/D/F)
- (Bonus in my fork) Optional Pass/Fail mode

---

## Prerequisites

- **Go 1.24+** (tested with Go 1.25)
- macOS/Linux/Windows shell

Verify your Go install:

```bash

go version
Running Tests

This project uses Go’s built-in test framework.

Run all tests:

go test ./...


Run tests with verbose output:

go test -v ./...


Run tests with coverage:

go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt


This shows what percentage of the code is executed by the test suite.

Project Structure

go.mod – Go module definition (name and Go version).

grade_calculator.go – Core implementation of the grade calculator:

GradeCalculator struct with methods for adding grades and computing results.

Category weighting (Assignments 50%, Exams 35%, Essays 15%).

Conversion from numeric to letter grades.

grade_calculator_test.go – Unit tests that validate grade conversion and calculator behavior.

grade_calculator_additional_test.go – Extra tests for boundary cutoffs, averages, and helper coverage.

Example Usage
gc := NewGradeCalculator()
gc.AddGrade("hw1", 95, Assignment)
gc.AddGrade("midterm", 85, Exam)
gc.AddGrade("essay1", 75, Essay)

final := gc.GetFinalGrade()
fmt.Println("Final grade:", final) // Output: B

Development Notes

Tests cover both correct behavior and boundary conditions.

Coverage reports ensure we are testing edge cases like empty grade slices.

Commits are separated by logical units of work (fixes, test updates, refactors, docs).

License

This project is for educational purposes only.