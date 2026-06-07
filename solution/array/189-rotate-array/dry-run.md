# Rotate Array — Dry Run

## Goal

Rotate the array to the right by `k` positions.

### Input

```text
nums = [1,2,3,4,5,6,7]
k = 3
```

### Expected Output

```text
[5,6,7,1,2,3,4]
```

---

# High-Level Idea

Instead of moving elements one by one:

1. Reverse the entire array.
2. Reverse the first k elements.
3. Reverse the remaining elements.

This achieves rotation using only constant extra space.

---

# Initial State

| Variable | Value |
|-----------|---------|
| nums | [1,2,3,4,5,6,7] |
| n | 7 |
| k | 3 |

---

# Step 0 — Normalize k

Formula:

```text
k = k % n
```

Calculation:

```text
3 % 7 = 3
```

No change.

| Variable | Value |
|-----------|---------|
| k | 3 |

---

# Step 1 — Reverse Entire Array

Operation:

```text
reverse(0, 6)
```

---

## Iteration 1

### Before Swap

```text
[1,2,3,4,5,6,7]
 ↑           ↑
left=0   right=6
```

Swap:

```text
1 ↔ 7
```

### After Swap

```text
[7,2,3,4,5,6,1]
```

| Left | Right | Array |
|--------|---------|---------|
| 1 | 5 | [7,2,3,4,5,6,1] |

---

## Iteration 2

### Before Swap

```text
[7,2,3,4,5,6,1]
   ↑       ↑
 left=1 right=5
```

Swap:

```text
2 ↔ 6
```

### After Swap

```text
[7,6,3,4,5,2,1]
```

| Left | Right | Array |
|--------|---------|---------|
| 2 | 4 | [7,6,3,4,5,2,1] |

---

## Iteration 3

### Before Swap

```text
[7,6,3,4,5,2,1]
     ↑   ↑
 left=2 right=4
```

Swap:

```text
3 ↔ 5
```

### After Swap

```text
[7,6,5,4,3,2,1]
```

| Left | Right | Array |
|--------|---------|---------|
| 3 | 3 | [7,6,5,4,3,2,1] |

Loop stops.

---

## State After Step 1

```text
[7,6,5,4,3,2,1]
```

---

# Step 2 — Reverse First k Elements

Operation:

```text
reverse(0, 2)
```

Current Array:

```text
[7,6,5,4,3,2,1]
```

---

## Iteration 1

### Before Swap

```text
[7,6,5,4,3,2,1]
 ↑   ↑
0   2
```

Swap:

```text
7 ↔ 5
```

### After Swap

```text
[5,6,7,4,3,2,1]
```

| Left | Right | Array |
|--------|---------|---------|
| 1 | 1 | [5,6,7,4,3,2,1] |

Loop stops.

---

## State After Step 2

```text
[5,6,7,4,3,2,1]
```

Notice:

```text
[5,6,7]
```

is now correctly positioned at the front.

---

# Step 3 — Reverse Remaining Elements

Operation:

```text
reverse(3, 6)
```

Current Array:

```text
[5,6,7,4,3,2,1]
```

---

## Iteration 1

### Before Swap

```text
[5,6,7,4,3,2,1]
       ↑     ↑
       3     6
```

Swap:

```text
4 ↔ 1
```

### After Swap

```text
[5,6,7,1,3,2,4]
```

| Left | Right | Array |
|--------|---------|---------|
| 4 | 5 | [5,6,7,1,3,2,4] |

---

## Iteration 2

### Before Swap

```text
[5,6,7,1,3,2,4]
         ↑ ↑
         4 5
```

Swap:

```text
3 ↔ 2
```

### After Swap

```text
[5,6,7,1,2,3,4]
```

| Left | Right | Array |
|--------|---------|---------|
| 5 | 4 | [5,6,7,1,2,3,4] |

Loop stops.

---

# Final State

```text
[5,6,7,1,2,3,4]
```

---

# Complete Transformation

```text
Original
[1,2,3,4,5,6,7]

Reverse Entire Array
[7,6,5,4,3,2,1]

Reverse First k Elements
[5,6,7,4,3,2,1]

Reverse Remaining Elements
[5,6,7,1,2,3,4]
```

---

# Visual Rotation Explanation

Desired rotation:

```text
[1,2,3,4 | 5,6,7]
             ↑
          move front
```

Target:

```text
[5,6,7 | 1,2,3,4]
```

Reversal technique achieves exactly this without creating another array.

---

# Edge Case Walkthrough

## Case 1

```text
nums = [1]
k = 10
```

### Normalize

```text
10 % 1 = 0
```

Result:

```text
[1]
```

---

## Case 2

```text
nums = [1,2,3]
k = 0
```

Result:

```text
[1,2,3]
```

No rotation required.

---

## Case 3

```text
nums = [1,2,3]
k = 5
```

Normalize:

```text
5 % 3 = 2
```

Rotate by 2:

```text
[2,3,1]
```

---

# Complexity Verification

## Time Complexity

Three reversals:

```text
O(n)
```

---

## Space Complexity

Only pointers used:

```text
O(1)
```

---

# Key Takeaway

The most important insight is:

```text
Rotate Right
=
Reverse Whole Array
+
Reverse First k Elements
+
Reverse Remaining Elements
```

This transforms the array in-place and produces the optimal solution expected in interviews.