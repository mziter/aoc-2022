# Advent of Code 2022 using Go

## Project Structure
As this project is just for fun, I am attempting to keep the simplest layout possible. Each day can be generated,
built, and run independently. 

```shell
go run ./cmd/day1
go build ./cmd/day5
```

## Automation
You can automatically create a new day folder and download the input for that day using
the Makefile from the root of the project where ? is the day number.

```shell
make new year=2022 day=?
```

*This requires that you have a file `session_cookie` in the root folder that contains a valid session cookie
from your browser when visiting the adventofcode website*

## Day Status and Execution Times
No execution time means day or part is not complete. These times capture the execution times
for each part. These are meant to be rough estimate values to observe orders of magnitude and
are very much subject to wide degrees of changes depending on many external factors.

`cpu: Intel(R) Core(TM) i5-4258U CPU @ 2.40GHz`

| Day # |  Part One  |   Part Two  | 
| ----: |  -------:  |  -------:   | 
| 1     |  66.16 µs  |  66.61 µs   | 
| 2     |  51.33 µs  |  58.84 µs   |
| 3     |  51.24 µs  |  83.03 µs   |