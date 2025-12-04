# Advent of Code 2025 (https://adventofcode.com/2025)

## Day 1 - Secret Entrance

### Overview

We are given a list of rotations, and a dial numbered from 0 - 99 that starts at 50. Rotating left, decreases, Rotating right, increases. The solution involves counting the number of times 0 is seen. In part 1, we need to count how many times the dial lands on 0 **after** a rotation. In part 2, we need to count how many times the dial **sees** 0, whether during a rotation or ending on 0. 

#### Part 1 

I built a simple state machine to tackle this problem with. We build a struct that houses the current dial value, then after each rotation we see if 0 has been hit, if it has, then we increment the amount of times hit 0. 

 - Mistake: Failed to account for rotations > 100. A simple modulus fixed that for us fairly easily. 

#### Part 2

A modification of the first approach was needed, instead of checking the end, we had to increment if it hits 0 **at all**. Floor div solved the overall times it returned back to the same place, thus hitting 0 each time, then we loop through the remainder, increment 1 by 1. if it == 0 at any time throughout, then inc the counter. 

## Day 2 - Gift Shop

### Overview

We are given a range of product ID's that the elves got all mangled up. Each range gives it's **first id** and its **last id** seperated by a `-`. Invalid IDs are one where **any** sequence of digits are repeated twice [55, 6464, 123123, etc]. Find the invalid IDs for the given ranges and add them together. 

#### Part 1

Actually got this first try. I literally just took every int in the range, and converted it to a string, split the string in two halfs and compared them. I thought surely there would be an edge case. But it was the correct answer..

#### Part 2
For this solution, I ended up iterating through the string and appending the rune to the current sequence, if at any point the sequence repeated for the len(str)/len(seq) times **and** the string != to the sequence. Then we return true and add it to the list to add up. 


## Day 3 - Lobby

### Overview
We are given a list of integers that represent battery bank joltage ratings. In part 1 we can switch two on, and concatenate the values left to right 12345 we can make 23, 24, 25 etc.. Part ones task is to find the largest possible voltage each bank can produce and then sum them up for the answer. Part two is the same, but for 12 digits.. 



#### Part 1
I'm terrible at combinitorics, but for this part specifically, I can get through without.. What I ended up doing, is converting each line to an int array, found the first max for the 10's digit, then found the second max for the 1's digit, this guaruntees the correct solution.. unless you forget to omit the last battery in the first max calc.. which I definitely didn't do with test input. 

#### Part 2
This one **sucked**. I ended up realizing that my computer literally couldn't brute force it even if I tried. Even an implementation of the Choice function (https://en.wikipedia.org/wiki/Combination), but while it significantly sped up part 1, and solved the sample problem on part 2, it froze my computer upon trying to execute on part 2. Threw that out the window and then implemented a greedy sliding window algorithm that scans from left to right in a window that is `k - remaining` long, chooses the max, and then loops until all versions are compared. We could only implement this i believe because the **order is preserved** from left to right. so we can assume that the next max value is to the right of the 1st max value.  