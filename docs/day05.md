# Day 5
Today's challenge was a bit messier given the odd form of the input data and trying to stay as cpu efficient and
non-pessimistic as possible. The logic of the puzzle was fairly straightforward though. After the initial pile layout
was known, the items were moved into stack data structures where they could be popped and pushed according to the
input instructions.

I took this opportunity with this day to give generics a try and created two generic stack implementations, a node
based one, and an slice based one. I would like to benchmark both once I refactor.

Eventually I would like to refactor the day in general, it is a bit messy with giant loops and if statements, but I
didn't want to refactor prematurely. I would also like to pull out some of the IO looping logic into some
common components and maybe do some more comparison vs bufio.