# Design Notes
This folder will contain some notes on design choices made during each day. Not every day will
have detailed choices depending on how froggy I feel. Usually I will try to eek out performance on days
with easier tasks and discuss more logic on days with more difficult tasks.

The general theme here for at least the first few days is to use the most basic data structures required when possible and iterating over existing resources to minimize time and space wasted by making copies. This goes A LONG WAY and brought some days
benchmarks from 500µs to < 100µs. Some examples are used fixed size vectors instead of slices, reusing vectors with counts 
instead of allocating maps (see Day 3, speed vs size tradeoff), avoid escaping values, etc...