# dry-run.md

# Dry Run

This document provides a step-by-step walkthrough of the optimal **Two Pointer** solution for **LeetCode 167 — Two Sum II: Input Array Is Sorted**.

---

# Goal

Given a sorted array and a target value, find the **1-indexed** positions of the two numbers whose sum equals the target.

---

# Algorithm Recap

```text
Initialize:

left = 0
right = n - 1

While left < right

    currentSum = numbers[left] + numbers[right]

    if currentSum == target
        return answer

    if currentSum < target
        left++

    else
        right--
```

---

# Why This Works

Since the array is sorted:

* Moving the **left pointer** increases the sum.
* Moving the **right pointer** decreases the sum.
* Every move removes impossible combinations from consideration.
* No pair is revisited.

This guarantees an **O(n)** solution.

---

# Example 1

## Input

```text
numbers = [2, 7, 11, 15]
target = 9
```

---

## Initial State

```text
                 L           R
                 ↓           ↓

numbers = [2, 7, 11, 15]

left = 0
right = 3
```

---

## Iteration 1

| Property    | Value |
| ----------- | ----- |
| Left Index  | 0     |
| Right Index | 3     |
| Left Value  | 2     |
| Right Value | 15    |
| Current Sum | 17    |

Comparison:

```text
17 > 9
```

Decision:

Move the **right pointer** left because the sum is too large.

---

## State After Iteration 1

```text
                 L       R
                 ↓       ↓

numbers = [2, 7, 11, 15]
```

---

## Iteration 2

| Property    | Value |
| ----------- | ----- |
| Left Index  | 0     |
| Right Index | 2     |
| Left Value  | 2     |
| Right Value | 11    |
| Current Sum | 13    |

Comparison:

```text
13 > 9
```

Decision:

Move the right pointer again.

---

## State After Iteration 2

```text
                 L   R
                 ↓   ↓

numbers = [2, 7, 11, 15]
```

---

## Iteration 3

| Property    | Value |
| ----------- | ----- |
| Left Index  | 0     |
| Right Index | 1     |
| Left Value  | 2     |
| Right Value | 7     |
| Current Sum | 9     |

Comparison:

```text
9 == 9
```

Target found.

Return:

```text
[1, 2]
```

---

# Iteration Summary

| Iteration | Left | Right | Values | Sum | Action        |
| --------- | ---- | ----- | ------ | --- | ------------- |
| 1         | 0    | 3     | 2 + 15 | 17  | Move Right    |
| 2         | 0    | 2     | 2 + 11 | 13  | Move Right    |
| 3         | 0    | 1     | 2 + 7  | 9   | Return Answer |

---

# Pointer Movement Visualization

```text
Step 1

2   7   11   15
L            R


↓

Sum Too Large

↓

Move Right



Step 2

2   7   11   15
L        R


↓

Sum Too Large

↓

Move Right



Step 3

2   7   11   15
L    R

↓

Target Found
```

---

# Example 2

## Input

```text
numbers = [1, 2, 3, 4, 4, 9]
target = 8
```

---

## Initial State

```text
1   2   3   4   4   9
L                   R
```

---

## Complete Execution

| Step | Left Value | Right Value | Sum | Decision      |
| ---- | ---------- | ----------- | --- | ------------- |
| 1    | 1          | 9           | 10  | Move Right    |
| 2    | 1          | 4           | 5   | Move Left     |
| 3    | 2          | 4           | 6   | Move Left     |
| 4    | 3          | 4           | 7   | Move Left     |
| 5    | 4          | 4           | 8   | Return Answer |

---

# Pointer Trace

```text
Initial

1 2 3 4 4 9
L         R


↓

Move Right


1 2 3 4 4 9
L       R


↓

Move Left


1 2 3 4 4 9
  L     R


↓

Move Left


1 2 3 4 4 9
    L   R


↓

Move Left


1 2 3 4 4 9
      L R

↓

Found
```

---

# State Transition Table

| Step   | Left Index | Right Index | Current Sum   | Comparison          | Action              |
| ------ | ---------- | ----------- | ------------- | ------------------- | ------------------- |
| Start  | 0          | n-1         | —             | —                   | Initialize pointers |
| Loop   | left       | right       | Compute sum   | Compare with target | Decide movement     |
| Case 1 | —          | —           | Sum < Target  | Increase sum        | `left++`            |
| Case 2 | —          | —           | Sum > Target  | Decrease sum        | `right--`           |
| Case 3 | —          | —           | Sum == Target | Solution found      | Return indices      |

---

# Why Pointer Movement Is Safe

### Case 1: Sum is Too Small

```text
numbers[left] + numbers[right] < target
```

Because the array is sorted:

* Every element to the left of `right` is **less than or equal** to `numbers[right]`.
* Keeping the same left value while moving the right pointer left can only keep the sum the same or make it smaller.

The only way to increase the sum is to move the **left pointer**.

---

### Case 2: Sum is Too Large

```text
numbers[left] + numbers[right] > target
```

Because the array is sorted:

* Every element to the right of `left` is **greater than or equal** to `numbers[left]`.
* Moving the left pointer right would only increase the sum further.

The only way to reduce the sum is to move the **right pointer**.

---

# Key Takeaways

* Always exploit the **sorted** property before considering extra data structures.
* The two-pointer technique eliminates impossible pairs in every iteration.
* Each pointer moves in only one direction, resulting in **O(n)** time complexity.
* The algorithm uses **O(1)** extra space.
* Understanding *why* each pointer moves is the key to explaining this solution confidently in interviews.
