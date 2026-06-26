# Dry Run — 209. Minimum Size Subarray Sum

This document provides a detailed walkthrough of the Sliding Window algorithm used to solve **LeetCode 209 - Minimum Size Subarray Sum**.

---

# Example

## Input

```text
target = 7
nums = [2, 3, 1, 2, 4, 3]
```

## Expected Output

```text
2
```

### Explanation

The smallest contiguous subarray whose sum is at least `7` is:

```text
[4, 3]
```

Length:

```text
2
```

---

# Initial State

Before processing any elements:

| Variable | Value |
|----------|------:|
| left | 0 |
| right | - |
| windowSum | 0 |
| minLength | ∞ |

Window:

```text
[]
```

---

# Sliding Window Visualization

Legend:

- **L** → Left Pointer
- **R** → Right Pointer

Initially:

```text
Index : 0  1  2  3  4  5
Value : 2  3  1  2  4  3

L,R
 ↓
[2]
```

---

# Step-by-Step Execution

## Step 1 — Expand Window

Move `right` to index **0**.

Window:

```text
[2]
```

Calculation:

```text
windowSum = 2
```

Since:

```text
2 < 7
```

Expand further.

| Left | Right | Window | Sum | Action |
|------|-------|--------|----:|--------|
| 0 | 0 | [2] | 2 | Expand |

---

## Step 2 — Expand Window

Move `right` to index **1**.

Window:

```text
[2,3]
```

```text
windowSum = 5
```

Still less than target.

| Left | Right | Window | Sum | Action |
|------|-------|--------|----:|--------|
| 0 | 1 | [2,3] | 5 | Expand |

---

## Step 3 — Expand Window

Move `right` to index **2**.

Window:

```text
[2,3,1]
```

```text
windowSum = 6
```

Still not enough.

| Left | Right | Window | Sum | Action |
|------|-------|--------|----:|--------|
| 0 | 2 | [2,3,1] | 6 | Expand |

---

## Step 4 — Expand Window

Move `right` to index **3**.

Window:

```text
[2,3,1,2]
```

```text
windowSum = 8
```

Now:

```text
8 ≥ 7
```

A valid window is found.

Current length:

```text
4
```

Update:

```text
minLength = 4
```

Now attempt to shrink.

| Left | Right | Window | Sum | Action |
|------|-------|--------|----:|--------|
| 0 | 3 | [2,3,1,2] | 8 | Shrink |

---

## Step 5 — Shrink Window

Remove:

```text
2
```

Window becomes:

```text
[3,1,2]
```

```text
windowSum = 6
```

Since:

```text
6 < 7
```

Stop shrinking.

| Left | Right | Window | Sum | Minimum |
|------|-------|--------|----:|---------:|
| 1 | 3 | [3,1,2] | 6 | 4 |

---

## Step 6 — Expand Window

Move `right` to index **4**.

Window:

```text
[3,1,2,4]
```

```text
windowSum = 10
```

Valid window.

Current length:

```text
4
```

Minimum remains:

```text
4
```

Shrink again.

| Left | Right | Window | Sum | Action |
|------|-------|--------|----:|--------|
| 1 | 4 | [3,1,2,4] | 10 | Shrink |

---

## Step 7 — Shrink Again

Remove:

```text
3
```

Window:

```text
[1,2,4]
```

```text
windowSum = 7
```

Still valid.

Length:

```text
3
```

Update:

```text
minLength = 3
```

Continue shrinking.

| Left | Right | Window | Sum | Minimum |
|------|-------|--------|----:|---------:|
| 2 | 4 | [1,2,4] | 7 | 3 |

---

## Step 8 — Continue Shrinking

Remove:

```text
1
```

Window:

```text
[2,4]
```

```text
windowSum = 6
```

Now invalid.

Stop shrinking.

| Left | Right | Window | Sum | Minimum |
|------|-------|--------|----:|---------:|
| 3 | 4 | [2,4] | 6 | 3 |

---

## Step 9 — Expand Window

Move `right` to index **5**.

Window:

```text
[2,4,3]
```

```text
windowSum = 9
```

Valid.

Shrink.

| Left | Right | Window | Sum | Action |
|------|-------|--------|----:|--------|
| 3 | 5 | [2,4,3] | 9 | Shrink |

---

## Step 10 — Shrink Window

Remove:

```text
2
```

Window:

```text
[4,3]
```

```text
windowSum = 7
```

Still valid.

Length:

```text
2
```

Update:

```text
minLength = 2
```

| Left | Right | Window | Sum | Minimum |
|------|-------|--------|----:|---------:|
| 4 | 5 | [4,3] | 7 | **2** |

---

## Step 11 — Final Shrink

Remove:

```text
4
```

Window:

```text
[3]
```

```text
windowSum = 3
```

Window is no longer valid.

Traversal ends.

---

# Complete State Transition Table

| Step | Left | Right | Added | Removed | Window | Window Sum | Min Length |
|-----:|-----:|------:|------:|--------:|--------|-----------:|-----------:|
| Start | 0 | - | - | - | [] | 0 | ∞ |
| 1 | 0 | 0 | 2 | - | [2] | 2 | ∞ |
| 2 | 0 | 1 | 3 | - | [2,3] | 5 | ∞ |
| 3 | 0 | 2 | 1 | - | [2,3,1] | 6 | ∞ |
| 4 | 0 | 3 | 2 | - | [2,3,1,2] | 8 | 4 |
| 5 | 1 | 3 | - | 2 | [3,1,2] | 6 | 4 |
| 6 | 1 | 4 | 4 | - | [3,1,2,4] | 10 | 4 |
| 7 | 2 | 4 | - | 3 | [1,2,4] | 7 | 3 |
| 8 | 3 | 4 | - | 1 | [2,4] | 6 | 3 |
| 9 | 3 | 5 | 3 | - | [2,4,3] | 9 | 3 |
| 10 | 4 | 5 | - | 2 | [4,3] | 7 | **2** |
| 11 | 5 | 5 | - | 4 | [3] | 3 | **2** |

---

# Pointer Movement

```text
Right Pointer
0 → 1 → 2 → 3 → 4 → 5

Left Pointer
0 → 1 → 2 → 3 → 4 → 5
```

Each pointer moves only forward and never revisits an index.

---

# Why the Algorithm is O(n)

Although there is a nested `while` loop, the total work remains linear because:

- Each element is added to the window exactly once.
- Each element is removed from the window exactly once.
- Neither pointer moves backward.

Therefore:

```text
Right Pointer Movements = n
Left Pointer Movements  = n

Total Operations ≈ 2n

Time Complexity = O(n)
```

---

# Final Result

```text
Minimum Length = 2
```

Smallest valid subarray:

```text
[4,3]
```

---

# Key Takeaways

- Maintain a running sum instead of recalculating subarray sums.
- Expand the window until it satisfies the target.
- Shrink the window immediately to minimize its length.
- The positivity of array elements guarantees that pointer movements are monotonic, enabling an **O(n)** solution.