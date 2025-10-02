# ESEP Grade Calculator

A tiny Go library that computes a student's final **letter grade** from three categories:

- **Assignments** — 50%
- **Exams** — 35%
- **Essays** — 15%

Each category is averaged, weights are applied, and the weighted score is mapped to letter grades:
**A ≥ 90, B ≥ 80, C ≥ 70, D ≥ 60, F < 60.**

> **Assumptions for this assignment**
> - Weights are fixed as above.
> - If a category has no grades, it contributes **0** (no re-weighting).
> - The weighted score is **floored to an integer** before letter mapping.

---

## Requirements

- Go **1.24.4+**

## Setup

```bash
# From repo root, navigate to the grade-calculator directory
cd grade-calculator
go mod tidy
```

## Run Tests

```bash
# From grade-calculator directory
go test
```

## Coverage

```bash
# From grade-calculator directory
go test -coverprofile=coverage.txt
go tool cover -func=coverage.txt

# Optional HTML report:
go tool cover -html=coverage.txt
```

## Usage Example

```go
package main

import (
    "fmt"
    esep "esep/grade-calculator"
)

func main() {
    gc := esep.NewGradeCalculator()
    gc.AddGrade("hw1", 92, esep.Assignment)
    gc.AddGrade("midterm", 88, esep.Exam)
    gc.AddGrade("essay1", 94, esep.Essay)
    fmt.Println(gc.GetFinalGrade()) // e.g., "A"
}
```

## Project Layout

- `grade-calculator/grade_calculator.go` — Library implementation
- `grade-calculator/grade_calculator_test.go` — Unit tests
- `grade-calculator/go.mod` — Module metadata