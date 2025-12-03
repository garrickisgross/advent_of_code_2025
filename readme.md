# Advent of Code 2025

## Day 1 - Secret Entrance

### Overview

You arrive at the secret entrance to the North Pole base ready to start decorating. Unfortunately, the password seems to have been changed, so you can't get in. A document taped to the wall helpfully explains:

The safe has a dial with only an arrow on it; around the dial are the numbers 0 through 99 in order. As you turn the dial, it makes a small click noise as it reaches each number.

The attached document (your puzzle input) contains a sequence of rotations, one per line, which tell you how to open the safe. A rotation starts with an L or R which indicates whether the rotation should be to the left (toward lower numbers) or to the right (toward higher numbers). Then, the rotation has a distance value which indicates how many clicks the dial should be rotated in that direction.

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