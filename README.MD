# ESEP Grade Calculator

A Go library that calculates a final course grade from assignments (50%),
exams (35%), and essays (15%). Each category contributes its average score
multiplied by its weight. Letter cutoffs are A: 90+, B: 80+, C: 70+,
D: 60+, and F: below 60. A C or better is passing.

## Requirements and setup

Install Go 1.24.4 or newer. The module uses only the Go standard library;
there are no external packages, environment variables, databases, or services.
From the repository root, enter the actual module directory (with a hyphen):

```sh
cd grade-calculator
go version
go mod download
go build ./...
```

The module path is `esep/grade-calculator` and the package name is
`esepunittests`. This is a library, so `go run .` is not an application entry point.

## Run tests and measure coverage

```sh
go test -v -count=1 ./...
go test -count=1 -coverprofile=/tmp/esep-coverage.txt ./...
go tool cover -func=/tmp/esep-coverage.txt
go tool cover -html=/tmp/esep-coverage.txt -o /tmp/esep-coverage.html
go vet ./...
```

Keep generated reports outside the checkout. A nonzero test exit status is a
failure even if a coverage report was generated. Coverage measures executed
statements, not whether assertions check the right behavior.

## Use the API

Construct a calculator with `NewGradeCalculator()`, call
`AddGrade(name, score, gradeType)` using `Assignment`, `Exam`, or `Essay`,
then read `GetFinalGrade()`. Use scores from 0 through 100. The API currently
does not validate scores or return input errors.

## Starting repository status (before assignment fixes)

The three provided tests execute, but `TestGetGradeA` and `TestGetGradeB`
fail. The implementation sums slice indexes instead of score values and
uses exams for the essay average. `TestGetGradeF` passes with all-high scores
because the first bug masks its incorrect assertion. Empty categories also
divide by zero. These are source/test defects, not missing dependencies.

## Project files and references

- `README.MD`: setup, commands, API, and known starting behavior.
- `grade-calculator/go.mod`: module identity and minimum Go version.
- `grade-calculator/grade_calculator.go`: grade types and calculator logic.
- `grade-calculator/grade_calculator_test.go`: Go `testing` unit tests.
- Go testing tutorial: https://go.dev/doc/tutorial/add-a-test
- Go coverage explanation: https://go.dev/blog/cover
