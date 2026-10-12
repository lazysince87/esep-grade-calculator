# Automated Unit Tests — Thuy Le

Fork: https://github.com/lazysince87/esep-grade-calculator

## Reproduce the final checks

Use Go 1.24.4 or newer from the repository's `grade-calculator` directory:

```sh
go build ./...
go vet ./...
go test -v -count=1 -coverprofile=/tmp/esep-coverage.txt ./...
go tool cover -func=/tmp/esep-coverage.txt
```

There are no third-party dependencies or required services.

## Measured results

| Stage | Result | Statement coverage |
| --- | --- | --- |
| Starting code | A and B fail; F passes despite an incorrect expected result | 80.0% |
| Part 3 | All 3 provided tests pass after source/data fixes | 85.2% |
| Part 4 | All 8 top-level tests and 10 table subtests pass | 100.0% |
| Part 5 | All `t.Errorf` calls commented out in a temporary module copy | 100.0% |
| Part 6 | Same test file, byte-for-byte unchanged; single-list implementation passes | 100.0% |
| Part 7.1 | Tests added first fail to compile because the mode API is not present | Not measured |
| Part 7.2 | Mode API compiles; new Pass/Fail assertions fail against letter output | Not measured |
| Part 7.3 | All 10 top-level tests and 16 table subtests pass | 100.0% |

The Part 5 copy was outside this checkout; assertions remain active here.
The Part 6 comparison used a SHA-256 hash of the test file before and after.
No test changes were needed for the storage refactor.

## Separate assignment commits

- Part 2, documentation only: https://github.com/lazysince87/esep-grade-calculator/commit/560c3c74b03df3f4172a7d6718bddeda1e4ccdb2
- Part 3, source fixes and existing test-data correction: https://github.com/lazysince87/esep-grade-calculator/commit/490aecc15b1b4d653ca273b3732a9314a3da320a
- Part 4, added coverage tests: https://github.com/lazysince87/esep-grade-calculator/commit/8a259518485f43503f8d65797ce2ff1609025e6f
- Part 6, storage refactor only: https://github.com/lazysince87/esep-grade-calculator/commit/49bdaf815344e9cead20d7449d794460caf4d128
- Part 7.1, tests first: https://github.com/lazysince87/esep-grade-calculator/commit/7aa1c2475cc6b0aa8a51f6b1be031feae3e01cc1
- Part 7.2, mode and constructor: https://github.com/lazysince87/esep-grade-calculator/commit/6e4d4c338a20599310024ceac1776cd7ea224261
- Part 7.3, behavior implementation: https://github.com/lazysince87/esep-grade-calculator/commit/327cb0597f978f754612dc0f4112f1a6f14e8dd3

## Requirements and assumptions

The confirmed scale is A ≥90, B ≥80, C ≥70, D ≥60, F <60; passing is C or
better. Weights are assignments 50%, exams 35%, essays 15%. Means retain
fractions; the final weighted value is truncated, matching the original
numerical-return contract. Empty categories contribute zero without
redistribution. Unknown types are ignored as before. Scores should be 0–100;
validation and configurable weights were not added to the assignment scope.

Pass/Fail is selected through the optional constructor mode; existing no-argument
construction remains compatible, and the public methods are unchanged.

Resources consulted:

- https://go.dev/doc/tutorial/add-a-test
- https://go.dev/blog/cover
- The repository README, module file, implementation, and test file.
