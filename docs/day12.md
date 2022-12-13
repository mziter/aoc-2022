# Day 12
Today we accomplished the task by doing a BFS over the grid to find the shortest path from one point to another. The second part was simply to do the same thing but from all of the 'a' points on the grid.

After first implementing the solution in the most straightforward way, using a map to track seen points and breaking up
the code how I thought made since, the benchmark clocked in at 1848341 ns/op, or ~1.85ms. 

After running the profiling tools, I saw that a lot of time was spent allocating on the map and also allocating a slice of neighbors every time we wanted to get valid neighbors on a point. At this point I also realized that a flat slice could be used to track whether or not a point was seen just by doing simple math on the x and y coordinates. This would allow for all allocation at once and to avoid hashing time.

After the making the above changes the benchmark was now 246169 ns or 246.17µs, roughly 7x faster.