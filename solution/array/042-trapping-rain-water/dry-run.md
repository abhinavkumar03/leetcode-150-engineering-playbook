# Dry Run — Trapping Rain Water

## Objective

Understand exactly how the optimal Two Pointer solution calculates trapped rain water.

We will trace every iteration and observe:

- Pointer movement
- Maximum boundaries
- Water trapped at each step
- Running total

---

# Example Input

```text
height = [4,2,0,3,2,5]
```

Expected Output:

```text
9
```

---

# Visual Representation

```text
Index:    0 1 2 3 4 5

Height:   4 2 0 3 2 5


            █
            █
█           █
█     █     █
█ █   █ █   █
█ █ █ █ █ █ █
----------------
0 1 2 3 4 5
```

Water trapped:

```text
            █
            █
█ ~ ~ ~ ~   █
█ ~ ~ █ ~   █
█ █ ~ █ █   █
█ █ █ █ █ █ █
----------------
```

Total trapped water:

```text
2 + 4 + 1 + 2 = 9
```

---

# Initial State

```text
left      = 0
right     = 5

leftMax   = 0
rightMax  = 0

water     = 0
```

---

# Core Rule

If:

```text
leftMax < rightMax
```

Process left side.

Otherwise:

```text
Process right side.
```

Reason:

The smaller boundary determines the water level.

---

# Iteration-by-Iteration Walkthrough

---

## Iteration 1

Pointers:

```text
L                           R
↓                           ↓

[4,2,0,3,2,5]
```

Update maximums:

```text
leftMax  = max(0,4) = 4
rightMax = max(0,5) = 5
```

Comparison:

```text
4 < 5
```

Process left.

Water added:

```text
leftMax - height[left]

4 - 4 = 0
```

Running total:

```text
0
```

Move:

```text
left++
```

New state:

```text
left = 1
right = 5
```

---

## Iteration 2

Pointers:

```text
    L                       R
    ↓                       ↓

[4,2,0,3,2,5]
```

Update maximums:

```text
leftMax  = 4
rightMax = 5
```

Comparison:

```text
4 < 5
```

Process left.

Water added:

```text
4 - 2 = 2
```

Running total:

```text
0 + 2 = 2
```

Move:

```text
left++
```

New state:

```text
left = 2
```

---

## Iteration 3

Pointers:

```text
      L                     R
      ↓                     ↓

[4,2,0,3,2,5]
```

Maximums:

```text
leftMax  = 4
rightMax = 5
```

Process left.

Water added:

```text
4 - 0 = 4
```

Running total:

```text
2 + 4 = 6
```

Move:

```text
left++
```

New state:

```text
left = 3
```

---

## Iteration 4

Pointers:

```text
        L                   R
        ↓                   ↓

[4,2,0,3,2,5]
```

Maximums:

```text
leftMax = 4
rightMax = 5
```

Process left.

Water added:

```text
4 - 3 = 1
```

Running total:

```text
6 + 1 = 7
```

Move:

```text
left++
```

New state:

```text
left = 4
```

---

## Iteration 5

Pointers:

```text
          L                 R
          ↓                 ↓

[4,2,0,3,2,5]
```

Maximums:

```text
leftMax = 4
rightMax = 5
```

Process left.

Water added:

```text
4 - 2 = 2
```

Running total:

```text
7 + 2 = 9
```

Move:

```text
left++
```

New state:

```text
left = 5
```

---

## Loop Ends

Condition:

```text
left < right
```

becomes:

```text
5 < 5
```

False.

Stop.

---

# State Transition Table

| Iteration | Left | Right | LeftMax | RightMax | Water Added | Total Water |
|------------|--------|--------|----------|-----------|-------------|-------------|
| Start | 0 | 5 | 0 | 0 | 0 | 0 |
| 1 | 0 | 5 | 4 | 5 | 0 | 0 |
| 2 | 1 | 5 | 4 | 5 | 2 | 2 |
| 3 | 2 | 5 | 4 | 5 | 4 | 6 |
| 4 | 3 | 5 | 4 | 5 | 1 | 7 |
| 5 | 4 | 5 | 4 | 5 | 2 | 9 |

Final Answer:

```text
9
```

---

# Why Water Can Be Calculated Immediately

Consider:

```text
leftMax = 4
rightMax = 5
```

Current position:

```text
height = 2
```

Water level:

```text
min(4,5) = 4
```

Therefore:

```text
4 - 2 = 2
```

Even if there is a taller wall farther right:

```text
10
20
100
```

it doesn't matter.

The limiting wall is already:

```text
leftMax = 4
```

Thus the trapped water is known immediately.

---

# Alternative Verification Using Formula

For every index:

```text
water[i] =
min(leftMax[i], rightMax[i])
- height[i]
```

| Index | Height | Left Max | Right Max | Water |
|---------|---------|----------|-----------|--------|
| 0 | 4 | 4 | 5 | 0 |
| 1 | 2 | 4 | 5 | 2 |
| 2 | 0 | 4 | 5 | 4 |
| 3 | 3 | 4 | 5 | 1 |
| 4 | 2 | 4 | 5 | 2 |
| 5 | 5 | 5 | 5 | 0 |

Total:

```text
0 + 2 + 4 + 1 + 2 + 0
=
9
```

---

# Key Takeaways

### Observation 1

Water depends on:

```text
min(leftMax, rightMax)
```

not the larger boundary.

---

### Observation 2

The smaller boundary completely determines the water level.

---

### Observation 3

Two pointers allow us to process each index once.

---

### Observation 4

The optimal solution achieves:

```text
Time  : O(n)
Space : O(1)
```

which is the expected interview solution.

---

# Final Result

Input:

```text
[4,2,0,3,2,5]
```

Output:

```text
9
```

Total trapped rain water:

```text
9 units
```