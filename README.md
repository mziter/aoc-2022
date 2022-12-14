# Advent of Code 2022 using Go

## Project Structure
As this project is just for fun, I am attempting to keep the simplest layout possible. Each day can be generated,
built, and run independently. 

```shell
go run ./cmd/day1
go build ./cmd/day5
```

## Goals 
My goal for this year is to focus on performance. I was recently inspired by
[Casey Muratori's Explanation of Non-Pessimization](https://www.youtube.com/watch?v=pgoetgxecw8) and wanted to really
focus on writing fast code by avoiding unnecessary work. 

You can find some design write-ups for each day in `/docs`.

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

| Day # |   Part One  |  Part Two   | 
| ----: |   -------:  |   -------:  | 
| 1     |   66.16 µs  |   66.61 µs  | 
| 2     |   51.33 µs  |   58.84 µs  |
| 3     |   51.24 µs  |   83.03 µs  |
| 4     |  116.55 µs  |  116.33 µs  |
| 5     |   81.96 µs  |   91.23 µs  |
| 6     |   18.37 µs  |   27.27 µs  |
| 7     |  182.39 µs  |  181.52 µs  |
| 8     |   89.55 µs  |  211.81 µs  |
| 9     |  545.21 µs  |    1.24 ms  |
| 10    |    7.67 µs  |   17.13 µs  |
| 11    |  111.09 µs  |   85.37 ms  |
| 12    |  246.17 µs  |  329.54 µs  |
| 13*   |    2.53 ms  |    6.89 ms  |
| 14    |  917.51 µs  |    5.54 ms  |

> Days marked with * were "heavily inspired" *cough cough* by someone else solution and are not entirely my own.