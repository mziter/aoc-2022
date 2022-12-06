# Day 6

## Part I
For first part I figured it would be nice to have a ringbuffer since we knew we wanted to look at the last n values and it would automatically handle old values falling off. I thought I would need a frequency table so I added the ability to see what value was replaced so that it could be decremented.

After thinking it over however, I thought it would be faster to just check every combination of the existing data in the ring buffer (worst case 4*3*2=24 loops) vs incrementing, decrementing values in another datastructure and looping through (26 loops). In practice this may not be true,  but I didn't test both.

## Part II
Here I introduced the frequency slice to track the number of occurrences of each character. To keep a zero based index we subtract 97 (ascii value of 'a') from each byte value for its index. Now that we had to check through 14 of them, it made more sense to loop through a slice and check the counts for each. We can increment the count for each new byte read and decrement the one that was replaced. 