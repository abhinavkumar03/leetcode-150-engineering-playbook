# Remove Element — Detailed Dry Run

## Goal

Remove all occurrences of `val` from the array **in-place** and return the count of remaining elements.

---

# Example Input

```text
nums = [0,1,2,2,3,0,4,2]
val = 2
```

Expected Output:

```text
k = 5

nums = [0,1,3,0,4,_,_,_]
```

Only the first `k` elements matter.

---

# Visualization

## Initial State

```text
Index: 0 1 2 3 4 5 6 7

Nums:  0 1 2 2 3 0 4 2

Write Pointer
↓
0
```

---

# Pointer Meaning

| Pointer | Purpose                   |
| ------- | ------------------------- |
| read    | Scans every element       |
| write   | Stores next valid element |

---

# Iteration 1

### Read Index = 0

```text
nums[0] = 0
```

Since:

```text
0 != 2
```

Keep the value.

Write it at position `write`.

```text
nums[0] = 0
```

Increment write.

### State

```text
Array : [0,1,2,2,3,0,4,2]

Read  : 0
Write : 1
```

---

# Iteration 2

### Read Index = 1

```text
nums[1] = 1
```

Since:

```text
1 != 2
```

Keep it.

```text
nums[1] = 1
```

Increment write.

### State

```text
Array : [0,1,2,2,3,0,4,2]

Read  : 1
Write : 2
```

---

# Iteration 3

### Read Index = 2

```text
nums[2] = 2
```

Target value found.

Skip it.

### State

```text
Array : [0,1,2,2,3,0,4,2]

Read  : 2
Write : 2
```

---

# Iteration 4

### Read Index = 3

```text
nums[3] = 2
```

Target value.

Skip.

### State

```text
Array : [0,1,2,2,3,0,4,2]

Read  : 3
Write : 2
```

---

# Iteration 5

### Read Index = 4

```text
nums[4] = 3
```

Valid value.

Move it to write position.

```text
nums[2] = 3
```

### State

```text
Array : [0,1,3,2,3,0,4,2]

Read  : 4
Write : 3
```

---

# Iteration 6

### Read Index = 5

```text
nums[5] = 0
```

Keep it.

```text
nums[3] = 0
```

### State

```text
Array : [0,1,3,0,3,0,4,2]

Read  : 5
Write : 4
```

---

# Iteration 7

### Read Index = 6

```text
nums[6] = 4
```

Keep it.

```text
nums[4] = 4
```

### State

```text
Array : [0,1,3,0,4,0,4,2]

Read  : 6
Write : 5
```

---

# Iteration 8

### Read Index = 7

```text
nums[7] = 2
```

Target value.

Skip.

### State

```text
Array : [0,1,3,0,4,0,4,2]

Read  : 7
Write : 5
```

---

# Complete Iteration Table

| Step | Read | Value | Action | Write After | Array Snapshot    |
| ---- | ---- | ----- | ------ | ----------- | ----------------- |
| 1    | 0    | 0     | Keep   | 1           | [0,1,2,2,3,0,4,2] |
| 2    | 1    | 1     | Keep   | 2           | [0,1,2,2,3,0,4,2] |
| 3    | 2    | 2     | Skip   | 2           | [0,1,2,2,3,0,4,2] |
| 4    | 3    | 2     | Skip   | 2           | [0,1,2,2,3,0,4,2] |
| 5    | 4    | 3     | Move   | 3           | [0,1,3,2,3,0,4,2] |
| 6    | 5    | 0     | Move   | 4           | [0,1,3,0,3,0,4,2] |
| 7    | 6    | 4     | Move   | 5           | [0,1,3,0,4,0,4,2] |
| 8    | 7    | 2     | Skip   | 5           | [0,1,3,0,4,0,4,2] |

---

# Final State

```text
Array:

[0,1,3,0,4,0,4,2]
```

Only first `k` positions matter.

```text
k = 5
```

Valid portion:

```text
[0,1,3,0,4]
```

Ignored portion:

```text
[0,4,2]
```

---

# Pointer Movement Diagram

```text
Initial

R
↓
[0,1,2,2,3,0,4,2]
↑
W


After Processing

                R
                ↓
[0,1,3,0,4,0,4,2]
          ↑
          W

k = 5
```

---

# Key Insight

The write pointer always represents:

> "The next position where a valid element should be placed."

Every non-`val` element is copied forward exactly once.

This makes the solution:

* O(n) Time
* O(1) Space
* Stable (preserves relative order)

---

# Reusable Pattern

```text
write = 0

for read in range(n):
    if nums[read] should be kept:
        nums[write] = nums[read]
        write++

return write
```

This same pattern is used in:

* LeetCode 26 — Remove Duplicates from Sorted Array
* LeetCode 283 — Move Zeroes
* LeetCode 905 — Sort Array By Parity
* Many stream filtering problems
