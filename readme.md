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

### Notes
Was kind of a rough start, need to pay more attention to the specifics moving forward / start thinking of edge cases sooner. Just gotta get back in the groove. Overall it took an hour to come to a solution while implementing in a new language. 