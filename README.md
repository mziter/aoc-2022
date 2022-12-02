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

## Day Status Execution Times
No execution time means day or part is not complete. These times capture the execution times
for each part and may not capture some of the initial file i/o. These is because we are mostly
focused on the execution time of the core logic and algorithms used.

| Day # |  Part One  |   Part Two  | 
| ----: |  -------:  |  -------:   | 
| 1     |  71.65 µs  |  71.22 µs   | 
| 2     | 376.41 µs  | 363.78 µs   |