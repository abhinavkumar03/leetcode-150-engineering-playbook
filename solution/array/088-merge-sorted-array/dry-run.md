# Merge Sorted Array — Dry Run

## Goal

Merge:

```text
nums2
```

into:

```text
nums1
```

while maintaining sorted order and using **O(1)** extra space.

---

# Example Input

```text
nums1 = [1,2,3,0,0,0]
m = 3

nums2 = [2,5,6]
n = 3
```

---

# Initial Setup

Valid portions:

```text
nums1 = [1,2,3]
nums2 = [2,5,6]
```

Available empty slots:

```text
[0,0,0]
```

---

# Pointer Initialization

```text
i = m - 1 = 2
j = n - 1 = 2
k = m + n - 1 = 5
```

### Meaning

| Pointer | Points To                     |
| ------- | ----------------------------- |
| i       | Last valid element in nums1   |
| j       | Last element in nums2         |
| k       | Last position in merged array |

---

# Visual Representation

```text
nums1 = [1,2,3,0,0,0]
             ^
             i

nums2 = [2,5,6]
             ^
             j

Merged Position
               ^
               k
```

---

# Iteration 1

## Compare

```text
nums1[i] = 3
nums2[j] = 6
```

Since:

```text
6 > 3
```

Place 6 at index k.

---

## Action

```text
nums1[5] = 6
```

Array becomes:

```text
[1,2,3,0,0,6]
```

Move:

```text
j--
k--
```

---

## State After Iteration

| Variable | Value |
| -------- | ----- |
| i        | 2     |
| j        | 1     |
| k        | 4     |

```text
nums1 = [1,2,3,0,0,6]
nums2 = [2,5,6]
```

---

# Iteration 2

## Compare

```text
nums1[i] = 3
nums2[j] = 5
```

Since:

```text
5 > 3
```

Place 5.

---

## Action

```text
nums1[4] = 5
```

Array:

```text
[1,2,3,0,5,6]
```

Move:

```text
j--
k--
```

---

## State

| Variable | Value |
| -------- | ----- |
| i        | 2     |
| j        | 0     |
| k        | 3     |

---

# Iteration 3

## Compare

```text
nums1[i] = 3
nums2[j] = 2
```

Since:

```text
3 > 2
```

Place 3.

---

## Action

```text
nums1[3] = 3
```

Array:

```text
[1,2,3,3,5,6]
```

Move:

```text
i--
k--
```

---

## State

| Variable | Value |
| -------- | ----- |
| i        | 1     |
| j        | 0     |
| k        | 2     |

---

# Iteration 4

## Compare

```text
nums1[i] = 2
nums2[j] = 2
```

Equal values.

Choose nums2 value.

---

## Action

```text
nums1[2] = 2
```

Array:

```text
[1,2,2,3,5,6]
```

Move:

```text
j--
k--
```

---

## State

| Variable | Value |
| -------- | ----- |
| i        | 1     |
| j        | -1    |
| k        | 1     |

---

# Main Loop Ends

Condition:

```text
j >= 0
```

fails.

```text
j = -1
```

Therefore all elements from nums2 have been merged.

---

# Why Remaining nums1 Elements Are Ignored

Current state:

```text
nums1 = [1,2,2,3,5,6]
```

Pointers:

```text
i = 1
j = -1
```

Remaining values:

```text
nums1[0...i]
```

are already:

* sorted
* correctly positioned

No additional work required.

---

# State Transition Table

| Step    | i | j | k | nums1[i] | nums2[j] | Chosen Value | Array State   |
| ------- | - | - | - | -------- | -------- | ------------ | ------------- |
| Initial | 2 | 2 | 5 | 3        | 6        | -            | [1,2,3,0,0,0] |
| 1       | 2 | 2 | 5 | 3        | 6        | 6            | [1,2,3,0,0,6] |
| 2       | 2 | 1 | 4 | 3        | 5        | 5            | [1,2,3,0,5,6] |
| 3       | 2 | 0 | 3 | 3        | 2        | 3            | [1,2,3,3,5,6] |
| 4       | 1 | 0 | 2 | 2        | 2        | 2            | [1,2,2,3,5,6] |

---

# Complete Pointer Movement Diagram

```text
Initial

nums1 = [1,2,3,0,0,0]
             i     k

nums2 = [2,5,6]
             j


After placing 6

nums1 = [1,2,3,0,0,6]
             i   k

nums2 = [2,5,6]
           j


After placing 5

nums1 = [1,2,3,0,5,6]
             i k

nums2 = [2,5,6]
         j


After placing 3

nums1 = [1,2,3,3,5,6]
           i k

nums2 = [2,5,6]
         j


After placing 2

nums1 = [1,2,2,3,5,6]
         i
```

---

# Dry Run Summary

### Key Insight

Always place the **largest remaining element** into the **last available position**.

```text
Compare:
nums1[i]
nums2[j]

Place larger value at:
nums1[k]
```

### Invariant

At every step:

```text
nums1[k+1 ... end]
```

already contains the correct sorted values.

This guarantees:

* No overwriting
* O(m+n) runtime
* O(1) extra space

---

# Final Output

```text
[1,2,2,3,5,6]
```
