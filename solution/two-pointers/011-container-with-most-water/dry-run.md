# dry-run.md

# Dry Run — Container With Most Water

This document provides a detailed walkthrough of the Two Pointer algorithm used to solve **LeetCode 11 — Container With Most Water**.

---

# Example

## Input

```text
height = [1,8,6,2,5,4,8,3,7]
```

## Expected Output

```text
49
```

---

# Visual Representation

```text
Index

 0 1 2 3 4 5 6 7 8
 | | | | | | | | |
 1 8 6 2 5 4 8 3 7

 L               R
```

Initial pointers:

```text
Left  = 0
Right = 8
```

---

# Area Formula

At every iteration:

```text
Area = min(height[left], height[right]) × (right - left)
```

---

# Step-by-Step State Transition

| Step | Left Index | Right Index | Left Height | Right Height | Width | Min Height | Current Area | Max Area | Pointer Move | Reason                                               |
| ---: | ---------: | ----------: | ----------: | -----------: | ----: | ---------: | -----------: | -------: | ------------ | ---------------------------------------------------- |
|    1 |          0 |           8 |           1 |            7 |     8 |          1 |            8 |        8 | Left++       | Left height is smaller.                              |
|    2 |          1 |           8 |           8 |            7 |     7 |          7 |           49 |       49 | Right--      | Right height is smaller.                             |
|    3 |          1 |           7 |           8 |            3 |     6 |          3 |           18 |       49 | Right--      | Right height is smaller.                             |
|    4 |          1 |           6 |           8 |            8 |     5 |          8 |           40 |       49 | Right--      | Heights are equal. Either pointer works; move right. |
|    5 |          1 |           5 |           8 |            4 |     4 |          4 |           16 |       49 | Right--      | Right height is smaller.                             |
|    6 |          1 |           4 |           8 |            5 |     3 |          5 |           15 |       49 | Right--      | Right height is smaller.                             |
|    7 |          1 |           3 |           8 |            2 |     2 |          2 |            4 |       49 | Right--      | Right height is smaller.                             |
|    8 |          1 |           2 |           8 |            6 |     1 |          6 |            6 |       49 | Right--      | Right height is smaller.                             |

Pointers now meet.

Algorithm terminates.

---

# Pointer Movement Visualization

## Iteration 1

```text
L               R

1 8 6 2 5 4 8 3 7

Area = 1 × 8 = 8

Move Left
```

---

## Iteration 2

```text
  L             R

1 8 6 2 5 4 8 3 7

Area = 7 × 7 = 49

Maximum Updated

Move Right
```

---

## Iteration 3

```text
  L           R

1 8 6 2 5 4 8 3 7

Area = 3 × 6 = 18

Move Right
```

---

## Iteration 4

```text
  L         R

1 8 6 2 5 4 8 3 7

Area = 8 × 5 = 40

Move Right
```

---

## Remaining Iterations

```text
  L       R

1 8 6 2 5 4 8 3 7

Area = 16

↓

Area = 15

↓

Area = 4

↓

Area = 6
```

No larger area is found.

---

# State Transition Summary

```text
Start

Left = 0
Right = 8

↓

Area = 8

↓

Move Left

↓

Area = 49

↓

Maximum Updated

↓

Move Right

↓

Area = 18

↓

Move Right

↓

Area = 40

↓

Move Right

↓

Area = 16

↓

Move Right

↓

Area = 15

↓

Move Right

↓

Area = 4

↓

Move Right

↓

Area = 6

↓

Pointers Meet

↓

Return 49
```

---

# Why Move the Shorter Pointer?

Current state:

```text
Height Left  = 3
Height Right = 8

Width = 10

Area = 3 × 10
```

The shorter line limits the container.

If we move the taller pointer:

```text
Width decreases

Height still ≤ 3

Area cannot increase.
```

If we move the shorter pointer:

```text
Width decreases

BUT

A taller line may be discovered.

Area may increase.
```

This greedy decision makes the algorithm optimal.

---

# Invariant Maintained

During every iteration:

* The current area is evaluated before moving pointers.
* The maximum area found so far is preserved.
* One pointer moves inward.
* The search space decreases by one.
* No potentially optimal solution is skipped because only the limiting (shorter) line is discarded.

---

# Complexity Walkthrough

For an array of length **n**:

* Left pointer moves at most **n** times.
* Right pointer moves at most **n** times.
* Each iteration moves exactly one pointer.
* Every element is processed at most once.

Therefore:

| Metric | Complexity |
| ------ | ---------: |
| Time   |   **O(n)** |
| Space  |   **O(1)** |

---

# Final Result

```text
Maximum Water = 49
```

The optimal container is formed by:

```text
Index 1 → Height = 8

Index 8 → Height = 7
```

```text
Width = 7

Minimum Height = 7

Area = 7 × 7 = 49
```
