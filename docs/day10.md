# Day 10
Fairly straightforward as far as modeling the problem. Some key points that I used when considering performance:
* scan input as bytes to avoid allocating strings
* parse integers using the newer fast functions in pkg/strutil
* parse instructions on the fly to avoid allocating temporary variables
* store pending executions as single variable instead of recording history in array, slice, ringbuffer, etc...
* store CRT pixel output in string builder until it is ready to be printed