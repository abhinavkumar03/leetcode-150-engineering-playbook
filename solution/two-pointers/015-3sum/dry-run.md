# Dry Run — 15. 3Sum

This document demonstrates how the optimized **Sorting + Two Pointers** solution works internally.

---

# Example

## Input

```text
nums = [-1,0,1,2,-1,-4]
```

---

## Step 1 — Sort the Array

Before applying the Two Pointer technique, sort the array.

```text
Original

[-1,0,1,2,-1,-4]
```

↓

```text
Sorted

[-4,-1,-1,0,1,2]
```

Sorted order allows us to:

- Move **left** pointer to increase the sum.
- Move **right** pointer to decrease the sum.

---

# Initial State

```text
Index

0   1   2   3   4   5
--------------------------------
-4 -1 -1  0  1  2
```

We iterate through every element and treat it as the first number of the triplet.

---

# Iteration 1

## Fixed Element

```text
i = 0

Value = -4
```

Pointers

```text
L = 1
R = 5
```

Visualization

```text
        L           R
        ↓           ↓
[-4, -1, -1, 0, 1, 2]
 ↑
 i
```

---

## Step-by-Step

| Step | i | Left | Right | Values | Sum | Decision |
|-----:|--:|-----:|------:|--------|----:|----------|
| 1 | 0 | 1 | 5 | (-4,-1,2) | -3 | Sum < 0 → Left++ |
| 2 | 0 | 2 | 5 | (-4,-1,2) | -3 | Sum < 0 → Left++ |
| 3 | 0 | 3 | 5 | (-4,0,2) | -2 | Sum < 0 → Left++ |
| 4 | 0 | 4 | 5 | (-4,1,2) | -1 | Sum < 0 → Left++ |

Pointers meet.

No valid triplet.

---

# Iteration 2

## Fixed Element

```text
i = 1

Value = -1
```

Pointers

```text
L = 2
R = 5
```

Visualization

```text
            L           R
            ↓           ↓
[-4, -1, -1, 0, 1, 2]
     ↑
     i
```

---

## Step 1

Triplet

```text
(-1,-1,2)
```

Sum

```text
0
```

Triplet found.

Store

```text
[-1,-1,2]
```

Move both pointers.

```text
L++
R--
```

---

### State After Movement

```text
L = 3
R = 4
```

Visualization

```text
               L    R
               ↓    ↓
[-4,-1,-1,0,1,2]
    ↑
    i
```

---

## Step 2

Triplet

```text
(-1,0,1)
```

Sum

```text
0
```

Triplet found.

Store

```text
[-1,0,1]
```

Move pointers again.

Pointers cross.

Iteration complete.

---

# Iteration 3

```text
i = 2
```

Current value

```text
-1
```

Previous value

```text
-1
```

Duplicate fixed element.

Skip it.

```text
continue
```

Reason:

Using the same fixed value again would generate duplicate triplets.

---

# Iteration 4

```text
i = 3

Value = 0
```

Pointers

```text
L = 4
R = 5
```

Visualization

```text
                  L  R
                  ↓  ↓
[-4,-1,-1,0,1,2]
          ↑
          i
```

Compute

```text
0 + 1 + 2 = 3
```

Since

```text
3 > 0
```

Move

```text
R--
```

Pointers meet.

Stop.

---

# Final Output

```text
[
  [-1,-1,2],
  [-1,0,1]
]
```

---

# Complete State Transition Table

| Iteration | Fixed Value | Left Value | Right Value | Sum | Action | Result |
|-----------|------------:|-----------:|------------:|----:|--------|--------|
| 1 | -4 | -1 | 2 | -3 | Left++ | — |
| 1 | -4 | -1 | 2 | -3 | Left++ | — |
| 1 | -4 | 0 | 2 | -2 | Left++ | — |
| 1 | -4 | 1 | 2 | -1 | Left++ | — |
| 2 | -1 | -1 | 2 | 0 | Store | [-1,-1,2] |
| 2 | -1 | 0 | 1 | 0 | Store | [-1,0,1] |
| 3 | -1 | — | — | — | Skip Duplicate | — |
| 4 | 0 | 1 | 2 | 3 | Right-- | — |

---

# Pointer Movement Rules

## Case 1 — Sum Too Small

```text
sum < 0
```

Move

```text
Left++
```

Reason:

The array is sorted.

Moving left increases the total sum.

---

## Case 2 — Sum Too Large

```text
sum > 0
```

Move

```text
Right--
```

Reason:

Moving right decreases the total sum.

---

## Case 3 — Perfect Match

```text
sum == 0
```

1. Save the triplet.
2. Move both pointers.
3. Skip duplicate values on both sides.

---

# Duplicate Handling Walkthrough

Consider

```text
[-2,0,0,0,2,2]
```

Sorted

```text
[-2,0,0,0,2,2]
```

Without duplicate handling, the algorithm would repeatedly produce:

```text
[-2,0,2]
```

To avoid this:

After finding a valid triplet:

```go
left++

while left < right && nums[left] == nums[left-1]
    left++
```

Similarly,

```go
right--

while left < right && nums[right] == nums[right+1]
    right--
```

This guarantees each unique triplet appears exactly once.

---

# Visual Timeline

```text
Input
 │
 ▼
Sort Array
 │
 ▼
Fix First Element
 │
 ▼
Initialize Left & Right
 │
 ▼
Compute Current Sum
 │
 ├───────────────┐
 │               │
 ▼               ▼
Sum < 0       Sum > 0
 │               │
Left++        Right--
 │               │
 └──────┬────────┘
        │
        ▼
   Sum == 0
        │
        ▼
 Store Triplet
        │
        ▼
Move Both Pointers
        │
        ▼
Skip Duplicates
        │
        ▼
Continue Until Left >= Right
        │
        ▼
Repeat for Next Fixed Element
```

---

# Key Takeaways

- Sorting enables efficient pointer movement.
- Each element is fixed exactly once.
- The remaining search is linear using two pointers.
- Duplicate skipping is required both:
  - for the fixed element (`i`)
  - after storing a valid triplet (`left` and `right`)
- The algorithm reduces the brute-force **O(n³)** solution to **O(n²)** while guaranteeing unique triplets.