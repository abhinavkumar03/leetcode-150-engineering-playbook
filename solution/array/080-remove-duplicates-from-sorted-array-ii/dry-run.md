# Dry Run — Remove Duplicates from Sorted Array II

## Goal

Given a sorted array, remove duplicates in-place such that each unique value appears **at most twice**.

Return the new length of the modified array.

---

# Example 1

## Input

```text
nums = [1,1,1,2,2,3]
```

## Expected Output

```text
k = 5

nums = [1,1,2,2,3]
```

---

# Core Idea

We maintain two pointers:

```text
read  -> scans every element
write -> points to next valid insertion position
```

Rule:

```text
Keep nums[read]
only if

nums[read] != nums[write - 2]
```

Why?

Because:

```text
write - 2
```

represents the second-most recent valid element.

If both values are equal, we already have two copies.

Adding another would violate the constraint.

---

# Initial State

Array:

```text
[1,1,1,2,2,3]
```

Length:

```text
n = 6
```

Initialize:

```text
write = 2
```

The first two elements are always allowed.

Current valid portion:

```text
[1,1]
```

Visualization:

```text
Index:  0 1 2 3 4 5
Value: [1 1 1 2 2 3]
         ^
         ^
       valid
```

---

# Iteration 1

## Read Pointer

```text
read = 2
```

Current value:

```text
nums[2] = 1
```

Compare with:

```text
nums[write - 2]
=
nums[0]
=
1
```

Comparison:

```text
1 == 1
```

Decision:

```text
Skip
```

Reason:

Two copies already exist.

State:

```text
write = 2
```

Array remains:

```text
[1,1,1,2,2,3]
```

---

# Iteration 2

## Read Pointer

```text
read = 3
```

Current value:

```text
nums[3] = 2
```

Compare with:

```text
nums[0] = 1
```

Comparison:

```text
2 != 1
```

Decision:

```text
Keep
```

Write operation:

```text
nums[write] = nums[3]

nums[2] = 2
```

Array becomes:

```text
[1,1,2,2,2,3]
```

Increment:

```text
write = 3
```

Valid portion:

```text
[1,1,2]
```

---

# Iteration 3

## Read Pointer

```text
read = 4
```

Current value:

```text
nums[4] = 2
```

Compare with:

```text
nums[write - 2]

nums[1] = 1
```

Comparison:

```text
2 != 1
```

Decision:

```text
Keep
```

Write operation:

```text
nums[3] = 2
```

Array:

```text
[1,1,2,2,2,3]
```

Increment:

```text
write = 4
```

Valid portion:

```text
[1,1,2,2]
```

---

# Iteration 4

## Read Pointer

```text
read = 5
```

Current value:

```text
nums[5] = 3
```

Compare with:

```text
nums[write - 2]

nums[2] = 2
```

Comparison:

```text
3 != 2
```

Decision:

```text
Keep
```

Write operation:

```text
nums[4] = 3
```

Array becomes:

```text
[1,1,2,2,3,3]
```

Increment:

```text
write = 5
```

Valid portion:

```text
[1,1,2,2,3]
```

---

# Final Result

Return:

```text
write = 5
```

Valid array:

```text
[1,1,2,2,3]
```

---

# Step-by-Step Table

| Iteration | Read | Current | Compare With | Action | Write After | Valid Portion |
|------------|------|----------|--------------|---------|-------------|----------------|
| Start | - | - | - | Initialize | 2 | [1,1] |
| 1 | 2 | 1 | 1 | Skip | 2 | [1,1] |
| 2 | 3 | 2 | 1 | Keep | 3 | [1,1,2] |
| 3 | 4 | 2 | 1 | Keep | 4 | [1,1,2,2] |
| 4 | 5 | 3 | 2 | Keep | 5 | [1,1,2,2,3] |

---

# Example 2

## Input

```text
[0,0,1,1,1,1,2,3,3]
```

---

# Initial State

```text
write = 2
```

Valid portion:

```text
[0,0]
```

---

# Execution Table

| Read | Value | Compare With nums[write-2] | Action | Write |
|--------|--------|--------|--------|--------|
| 2 | 1 | 0 | Keep | 3 |
| 3 | 1 | 0 | Keep | 4 |
| 4 | 1 | 1 | Skip | 4 |
| 5 | 1 | 1 | Skip | 4 |
| 6 | 2 | 1 | Keep | 5 |
| 7 | 3 | 1 | Keep | 6 |
| 8 | 3 | 2 | Keep | 7 |

---

# Final State

Array:

```text
[0,0,1,1,2,3,3,3,3]
```

Valid portion:

```text
[0,0,1,1,2,3,3]
```

Length:

```text
7
```

---

# Pointer Movement Visualization

```text
Initial

[1,1,1,2,2,3]
     R
     W

-------------------

Skip duplicate

[1,1,1,2,2,3]
       R
     W

-------------------

Keep 2

[1,1,2,2,2,3]
         R
       W

-------------------

Keep 2

[1,1,2,2,2,3]
           R
         W

-------------------

Keep 3

[1,1,2,2,3,3]
             R
           W
```

Legend:

```text
R = Read Pointer
W = Write Pointer
```

---

# State Transition Summary

## Before Processing

```text
nums[0...write-1]
```

contains valid output.

---

## During Processing

For every element:

```text
if nums[read] != nums[write - 2]
```

then:

```text
Accept element
Move write forward
```

otherwise:

```text
Reject element
Continue scanning
```

---

## After Processing

Invariant remains true:

```text
nums[0...write-1]
```

contains every value at most twice.

---

# Key Insight

The most important observation is:

```text
nums[read] != nums[write - 2]
```

This single comparison guarantees:

- No value appears more than twice.
- Order remains unchanged.
- O(n) runtime.
- O(1) extra space.