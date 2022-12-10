# Day 9
I solved the first part pretty quickly and then moved on to tuning performance. I took the day to do some low level optimizations that I haven't done before and to learn to use more of the pprof tool.

In the first iteration for Part I, I used the scanner to scan lines as strings, used `string.Split()` to split the instructions into direction and amount, used strconv to convert string numbers to ints, and used a map to track seen locations as Point structs.

After generating a profile of the benchmark of Part I using pprof, I realized there were quite a few problems locations:

* Using `Scan.Text()` and `String.Split()` are both pretty slow because they need to allocate new strings to return to us, even though we don't need to keep them for any reason. Also, `strconv.Atoi()` required us to allocate a string from our `[]byte` so
that we can convert it to an int.

After replacing the scan and split functions with their byte alternatives and copying a []byte to int conversion function the code was already 22% faster.

* Next up that we could see were growth functions in our map were taking a lot of the cpu time.

I saw that I didn't set an initial capacity, so I set a larger capacity. Now we are 42% faster than when we started.

* Lastly, I saw that our map of type `map[Point]bool` for tracking visited spots, where `Point { x: int, y : int }` contained the point, was taking a lot time allocating space. I thought it would be a good learning experience to build my own hash table. I could use a rather simple hashing function and gear everything towards my use case. This ended up working okay, but the hashing function was very bad for the number domain that we were using and didn't give that great of spread across the buckets. But the biggest issue I learned was that the built-in map functionality has a lot of access to private functions in the core language for getting quicker allocations. So without wanting to dig in too deeply, I gave up on my plans to build one myself (for now!). In my search, I did stumble upon some descriptions of how the built-in map works, and how it optimized for different keys. So I did pull out an idea from my hash table, which was to store x and y as a single number, where x is bit shifted left so that they could be stored as a single unique value. When this was used in the map instead of our custom struct the performance improvement was noticeable. Our final benchmark with all of the above optimizations was 545.21 µs, 60% faster