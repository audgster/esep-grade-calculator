# ESEP Grade Calculator

This library is a simple utility to calculate final grades for a class based on assignments, exams, and essays. The calculator will compute the final grade as a weighted average of the three categories, then convert that numerical grade to a letter grade.

## Repository Overview

There are four main files in this repository:

- grade_calculator.go
  Implements the grade calculator by defining structs and methods that allow defining grades with their respective names and categories (e.g., exam, essay, assignment) and calculating final grades.  

- grade_calculator_test.go
  Contains different unit tests for the `GetFinalGrade()` method.  

- README.md
  Explains what the project does, running requirements, and how to run the unit tests and coverage reports with Go.  

- go.mod
  Go module file that defines the module path and Go version used.  


## Tests

- Verify that the calculator outputs the correct letter grade for different grade inputs.  

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

## Tech Stack

- Go, an open-source programming language supported by Google  

Runtime requirements:
Golang 1.24+
