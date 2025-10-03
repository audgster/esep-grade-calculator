# ESEP Grade Calculator

A small Go library that computes a student’s **final letter grade** from three categories:
**Assignments**, **Exams**, and **Essays**.

> **Note:** In this repo the code lives in `grade-calculator/` (with a hyphen).

---

## Requirements

- **Go** 1.25+ (`go version` should print 1.25.x or newer)
- Go modules enabled (default in modern Go)

---

## Quick Start

```bash
git clone https://github.com/<your-username>/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
go test ./...

# overall coverage in stdout
go test -cover ./...

# detailed coverage file + per-function breakdown
go test -coverprofile=coverage.out ./...
go tool cover -func coverage.out

# optional HTML report
go tool cover -html=coverage.out

How It Works

Compute the average for each category:

assignments average

exams average

essays average

Compute the weighted numeric grade:

final = 0.50*avg(assignments) + 0.35*avg(exams) + 0.15*avg(essays)


Map the numeric to a letter grade:

A ≥ 90

B ≥ 80

C ≥ 70 (passing threshold)

D ≥ 60

F < 60

Averages and weighting use floating-point math (no integer truncation).

Example (pseudocode)
gc := NewGradeCalculator()
gc.AddGrade("assignment 1", 95, Assignment)
gc.AddGrade("exam 1", 88,       Exam)
gc.AddGrade("essay 1", 90,      Essay)

letter := gc.GetFinalGrade() // "A" | "B" | "C" | "D" | "F"

Project Structure
.
├── README.md                # this file
├── go.mod                   # module metadata
└── grade-calculator/
    ├── grade_calculator.go  # library implementation
    └── grade_calculator_test.go  # unit tests

Development Notes

Tests are table-driven and cover:

exact boundaries (90/80/70/60) and just-below cases

multiple items per category (averaging)

empty category behavior (avg = 0.0)

wrapper call (GetGrade) and enum string mapping

Goal coverage: 100% (use commands above to verify)