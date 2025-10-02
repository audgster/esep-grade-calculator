# ESEP Grade Calculator

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

Dev Use Notes:

# Grade Calculator (Go)

This package provides a simple way to calculate final letter grades from **assignments**, **exams**, and **essays**. It applies weighted averages to compute a numerical grade and then maps it to a standard letter grade (A–F).

---

## Usage

### Create a Grade Calculator

gc := esepunittests.NewGradeCalculator()

### Add Grades

gc.AddGrade("Assignment 1", 95, esepunittests.Assignment)
gc.AddGrade("Midterm Exam", 88, esepunittests.Exam)
gc.AddGrade("Essay on AI", 92, esepunittests.Essay)

### Get Final Grade

final := gc.GetFinalGrade()
fmt.Println("Final Grade:", final) // "A"

---

## API Reference

### Types

- GradeCalculator  
  Stores all grades and computes results.

- GradeType  
  Enum for grade categories:

  - Assignment
  - Exam
  - Essay

- Grade  
  Represents an individual grade (Name, Grade, Type).

---

## Testing

Run tests with:  
go test ./...
