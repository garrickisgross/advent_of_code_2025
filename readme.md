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
This one **sucked**. At first I tried an implementation of the Choice function (https://en.wikipedia.org/wiki/Combination), but while it significantly sped up part 1, and solved the sample problem on part 2, it froze my computer upon trying to execute on part 2. Threw that out the window and then implemented a greedy sliding window algorithm that scans from left to right in a window that is `k - remaining` long, chooses the max, and then loops until all versions are compared. We could only implement this i believe because the **order is preserved** from left to right. so we can assume that the next max value is to the right of the 1st max value.  

## Day 4 - Printing Department

### Overview
We are given a grid and rolls of paper are identified by an "@". In Part 1, the goal is to count the number of rolls that can be removed, identified if they have fewer than 4 other rolls in its 8 adjacent spaces. Part 2, if we remove those rolls, how many can we remove in total if done until we can't anymore. 

#### Part 1
EZPZ. grab the input as a grid, loop through the 2d grid, define a canAccess function that looks at every adjacent square and adds rolls up, if less than four, increase the iterator. First try. 

#### Part 2
EZIERPZIER. take the first solution and when you can access one, remove it from the grid and replace with "x" (or "." doesn't matter). If you find one, set done to false, need to rescan the grid, loop until done. 

## Day 5 - Cafeteria

### Overview
We are given a list of ranges encoding a **min** and a **max** along with some numbers that may or may not fit into those ranges. Part 1, the goal is to count how many of the numbers fall in the ranges. Part two asks to identify how many numbers fit in the ranges overall. Ranges may overlap.

#### Part 1 
This one was fairly straightforward, create a 2d array of ranges with the min and max, for each number, check if it fits in any of the ranges, if yes, then increment result. 

#### Part 2
For me, the best way forward was to merge all the ranges. then for each range, take the max - min + 1, and increment result by that much. I struggled quite a bit with the merging, ultimately I was using a flag to show when I edited the result array and it was in the wrong spot, which led to a ton of debugging. everytime I updated the array in the result, I recursively called the merge function to correct any overlap in the new ranges. This worked fairly well, although I feel like there is a more optimal solution for both parts. 

## Day 6 - Trash Compactor

### Overview
We are given a list of numbers with the final row being operations. Part on asks us to use the operations in vertical order, part 2 in columns with the most significant digit starting at the top. 

#### Part 1
I really just loaded the numbers in, kept them in a 2d array, and then used the index of the operations. Happened fairly quickly. 

#### Part 2
The only difference between this one and part one was the way we prepared the initial 2d array, it took me a little to reason through the construction, but once we got there, the solution was fairly straightforward. 

## Day 7 - Laboritories

### Overview
We are given a grid of string characters, S denotes the starting pos of the tachyon beam. The beam moves downwards, and when it hits a splitter "^", it splits into two beams on the left and the right and ceases to move downwards. Part 1 asks us to count how many time the beam splits. Part 2, annoyingly, involves counting all of the paths the beam can take if it only chooses one each time it hits a splitter. 

#### Part 1
I did okay with this one. I legitimately simulated it in code to get the answer, was fairly straightforward.. Part 2 however...

#### Part 2
Had to lean on the LLM quite a bit here. But essentially it uses a row by row dynamic programming approach and uses a map to hold the number of active timelines at the current row. 

- [TODO] practice dynamic programming problems to get more comfortable with this type of thing. 

## Day 8 - Playground

### Overview

We are given a set of coordinates in 3d space. the task essentially is to find the pairs with the shortest distances and connect them together, and keep track of the circuits that form. Part 1 has to do with making a certain number of connections, and then measuring the 3 biggest circuits. Part 2 wants us to connect until there is just one connection left and then look at those two points. 

#### Part 1
Both solutions had similar solves. for the pairs, we loop through all possible pairs, calculate distance and store in a slice of pairs that contain the index of the boxes and then the distance between that pair, we then use a quicksort implementation to sort those pairs lowest -> highest distance. 

This is where I started struggling, I needed a way to keep track of circuits. I chose to do this through a find union pattern. We define a parent that represent the root to the circuit mapped to the same index of boxes. when we match one up, we update its index in the parent to the root index of the one that we merged it with and update its size. We create a find function and a union function, the find function will find the root. and the union function will merge the circuits if needed. We then go through all the roots where the index didn't change, and map those sizes to a slice. grab the biggest three and multiply together. 

#### Part 2
For part two, its pretty much the same solution as part one, except you want to keep pairing boxes until there is one left. our union function returns a bool that lets us know if it made a merge or not, so using that, its pretty straightforward and mostly copy/paste. 