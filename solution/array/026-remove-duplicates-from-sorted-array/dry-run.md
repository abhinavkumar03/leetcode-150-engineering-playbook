# Dry Run — Remove Duplicates from Sorted Array

## Goal

Given a sorted array, remove duplicates in-place and return the number of unique elements.

The first `k` positions of the array should contain all unique values.

---

## Example Input

```text
nums = [0,0,1,1,1,2,2,3,3,4]
```

Expected Output:

```text
k = 5

nums = [0,1,2,3,4,_,_,_,_,_]
```

---

## Visual Idea

We maintain two pointers:

```text
slow → Last unique element position
fast → Current scanning position
```

Initial State:

```text
Index:  0 1 2 3 4 5 6 7 8 9
Value: [0,0,1,1,1,2,2,3,3,4]

slow
 ↓

fast
   ↓
```

---

## Invariant

At any moment:

```text
nums[0...slow]
```

contains all unique elements discovered so far.

Everything after `slow` is still being processed.

---

# Iteration Walkthrough

## Initial State

| Variable | Value |
|-----------|---------|
| slow | 0 |
| fast | 1 |

Array:

```text
[0,0,1,1,1,2,2,3,3,4]
```

Unique Region:

```text
[0]
```

---

## Iteration 1

### Compare

```text
nums[fast] = 0
nums[slow] = 0
```

Duplicate found.

### Action

```text
Do nothing
```

Move:

```text
fast++
```

| slow | fast |
|--------|--------|
| 0 | 2 |

Array:

```text
[0,0,1,1,1,2,2,3,3,4]
```

---

## Iteration 2

### Compare

```text
nums[fast] = 1
nums[slow] = 0
```

New unique value found.

### Action

```text
slow++
nums[slow] = nums[fast]
```

After update:

```text
slow = 1
nums[1] = 1
```

Array:

```text
[0,1,1,1,1,2,2,3,3,4]
```

Unique Region:

```text
[0,1]
```

---

## Iteration 3

### Compare

```text
nums[fast] = 1
nums[slow] = 1
```

Duplicate.

Move forward.

| slow | fast |
|--------|--------|
| 1 | 4 |

Array unchanged.

---

## Iteration 4

### Compare

```text
nums[fast] = 1
nums[slow] = 1
```

Duplicate.

Move forward.

| slow | fast |
|--------|--------|
| 1 | 5 |

Array unchanged.

---

## Iteration 5

### Compare

```text
nums[fast] = 2
nums[slow] = 1
```

Unique value found.

### Action

```text
slow++
nums[slow] = 2
```

Array:

```text
[0,1,2,1,1,2,2,3,3,4]
```

Unique Region:

```text
[0,1,2]
```

---

## Iteration 6

### Compare

```text
nums[fast] = 2
nums[slow] = 2
```

Duplicate.

Move forward.

---

## Iteration 7

### Compare

```text
nums[fast] = 3
nums[slow] = 2
```

Unique value found.

### Action

```text
slow++
nums[slow] = 3
```

Array:

```text
[0,1,2,3,1,2,2,3,3,4]
```

Unique Region:

```text
[0,1,2,3]
```

---

## Iteration 8

### Compare

```text
nums[fast] = 3
nums[slow] = 3
```

Duplicate.

Move forward.

---

## Iteration 9

### Compare

```text
nums[fast] = 4
nums[slow] = 3
```

Unique value found.

### Action

```text
slow++
nums[slow] = 4
```

Array:

```text
[0,1,2,3,4,2,2,3,3,4]
```

Unique Region:

```text
[0,1,2,3,4]
```

---

# State Transition Table

| Step | fast | nums[fast] | slow Before | Action | slow After | Valid Region |
|--------|--------|--------|--------|--------|--------|--------|
| Start | - | - | 0 | Initialize | 0 | [0] |
| 1 | 1 | 0 | 0 | Duplicate | 0 | [0] |
| 2 | 2 | 1 | 0 | Write unique | 1 | [0,1] |
| 3 | 3 | 1 | 1 | Duplicate | 1 | [0,1] |
| 4 | 4 | 1 | 1 | Duplicate | 1 | [0,1] |
| 5 | 5 | 2 | 1 | Write unique | 2 | [0,1,2] |
| 6 | 6 | 2 | 2 | Duplicate | 2 | [0,1,2] |
| 7 | 7 | 3 | 2 | Write unique | 3 | [0,1,2,3] |
| 8 | 8 | 3 | 3 | Duplicate | 3 | [0,1,2,3] |
| 9 | 9 | 4 | 3 | Write unique | 4 | [0,1,2,3,4] |

---

# Final Result

Final Array:

```text
[0,1,2,3,4,2,2,3,3,4]
```

Only the first `k` elements matter:

```text
[0,1,2,3,4]
```

Returned Value:

```text
k = slow + 1
k = 5
```

---

# Pointer Movement Diagram

```text
Initial

[0,0,1,1,1,2,2,3,3,4]
 S F

After finding 1

[0,1,1,1,1,2,2,3,3,4]
   S   F

After finding 2

[0,1,2,1,1,2,2,3,3,4]
     S       F

After finding 3

[0,1,2,3,1,2,2,3,3,4]
       S         F

After finding 4

[0,1,2,3,4,2,2,3,3,4]
         S           F
```

---

# Key Takeaway

The algorithm succeeds because:

```text
Sorted Array
+
Two Pointers
+
In-Place Writes
=
O(n) Time
O(1) Space
```

The `slow` pointer always marks the end of the unique region, while the `fast` pointer explores the remaining elements looking for the next unique value.