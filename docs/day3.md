# Day 3
Chose to use fixed sized vectors instead of maps to lower CPU time with maybe an expense of memory. Since
the domain of values is fixed to pretty small it seems like it is worth the cost here. Since we have to calculate the
priority anyway, it is faster to store/lookup in a vector based on index than to compute a hash and lookup in a map. We
also save time by not having to recreate the map each time we need a fresh one as opposed to just zero'ing our array.

Choose to iterate over the lines in part one to get lengths before iterating again so that we know the halfway point
ahead of time. This seemed okay since it was linear runtime and also could save time by knowing what compartment we
are in as well as short circuiting some of the logic by knowing when we found an item early.

## Possible Improvements
Instead of scanning entire contents and allocating an array to line sizes I could iterate ahead to find newline and then just starting back at beginning of line using index values of byte array.