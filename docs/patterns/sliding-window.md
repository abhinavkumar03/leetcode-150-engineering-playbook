# Sliding Window Pattern

> The Sliding Window pattern is an optimization technique for solving problems involving contiguous sequences (arrays or strings) by maintaining a "window" of elements instead of repeatedly recomputing overlapping ranges.

---

# Pattern Definition

The Sliding Window pattern maintains a contiguous range of elements while moving through a collection. Instead of evaluating every possible subarray or substring independently, the algorithm incrementally updates the current window by adding new elements and removing old ones.

This reduces many brute-force solutions from **O(n²)** to **O(n)**.

There are two major variants:

1. **Fixed Size Sliding Window**
2. **Variable Size Sliding Window**

---

# When to Use Sliding Window

Use this pattern when the problem contains one or more of the following signals:

- Contiguous subarray
- Contiguous substring
- Consecutive elements
- Window
- Range
- Segment
- Running sum
- Running frequency
- Minimum window
- Maximum window
- Longest valid sequence
- Shortest valid sequence

---

# Recognition Signals

## Fixed Size Window

Common keywords:

- Size = K
- Exactly K elements
- Fixed-length substring
- Average of K elements
- Maximum sum of K elements

Typical examples:

- Maximum Average Subarray I
- Maximum Sum Subarray of Size K

---

## Variable Size Window

Common keywords:

- At least
- At most
- Minimum length
- Maximum length
- Smallest window
- Longest window
- Valid window
- Constraint satisfied

Typical examples:

- Minimum Size Subarray Sum
- Minimum Window Substring
- Fruit Into Baskets
- Longest Repeating Character Replacement

---

# Fixed vs Variable Window

| Feature | Fixed Size | Variable Size |
|----------|------------|---------------|
| Window Length | Constant | Dynamic |
| Expand | Always | Until constraint is met |
| Shrink | Rarely | Frequently |
| Typical Goal | Sum/Average | Minimum/Maximum valid window |
| Time Complexity | O(n) | O(n) |

---

# Variable Sliding Window Template

```text
left = 0
windowState = initial value

for right in range(0, n):

    Expand window
    Update state

    while window satisfies condition:

        Update answer

        Shrink window

        left++
```

---

# Core Invariant

At every iteration:

```text
windowState
=
State of elements
between left and right (inclusive)
```

The window always represents the current contiguous segment under consideration.

---

# Complexity

| Operation | Complexity |
|-----------|------------|
| Time | O(n) |
| Space | O(1) to O(k), depending on maintained state |

Why is it linear?

- Right pointer moves from left to right once.
- Left pointer also moves from left to right once.
- No pointer ever moves backward.

Total pointer movements:

```text
≤ 2n
```

Therefore:

```text
O(n)
```

---

# Common Pitfalls

## 1. Choosing the Wrong Pattern

Sliding Window works only when the problem involves **contiguous** elements.

---

## 2. Using Variable Window with Negative Numbers

Many variable-size window problems (including LC 209) rely on all numbers being positive.

If negative values exist, expanding the window may decrease the sum, breaking the monotonic property.

---

## 3. Using `if` Instead of `while`

When shrinking is required, use:

```text
while condition is true
```

instead of:

```text
if condition is true
```

to fully optimize the current window.

---

## 4. Incorrect Window Size

Always calculate:

```text
right - left + 1
```

---

## 5. Forgetting to Update the State

Whenever a pointer moves:

- Update the running sum.
- Update the frequency map.
- Update any maintained statistics.

The maintained state must always reflect the current window.

---

# Optimization Journey

### Step 1

Brute Force

```text
O(n²)
```

---

### Step 2

Observe overlapping work.

---

### Step 3

Maintain the current window instead of recomputing it.

---

### Step 4

Update the window incrementally.

---

### Step 5

Achieve linear time.

```text
O(n)
```

---

# Related Problems

## Fixed Size Sliding Window

| Problem | Focus |
|----------|-------|
| 643. Maximum Average Subarray I | Fixed window sum |
| 1456. Maximum Number of Vowels in a Substring of Given Length | Fixed window frequency |
| 1343. Number of Sub-arrays of Size K and Average ≥ Threshold | Running average |

---

## Variable Size Sliding Window

| Problem | Focus |
|----------|-------|
| **209. Minimum Size Subarray Sum** | Minimum valid window (positive integers) |
| 3. Longest Substring Without Repeating Characters | Unique characters |
| 76. Minimum Window Substring | Minimum covering window |
| 424. Longest Repeating Character Replacement | Character replacement budget |
| 904. Fruit Into Baskets | At most two distinct values |
| 1004. Max Consecutive Ones III | Flip budget |
| 713. Subarray Product Less Than K | Product constraint |
| 1493. Longest Subarray of 1's After Deleting One Element | One deletion allowed |
| 1658. Minimum Operations to Reduce X to Zero | Prefix/Suffix transformation |

---

# Current Problem Addition

## 209. Minimum Size Subarray Sum

### Why It Belongs Here

This problem is a textbook example of the **Variable Size Sliding Window** pattern.

Key characteristics:

- Contiguous subarray
- Positive integers
- Running sum
- Expand until the target is reached
- Shrink to minimize the window
- Linear-time optimization

### Pattern-Specific Insight

The algorithm depends on the monotonic behavior of the running sum:

- Expanding the window can only increase the sum.
- Shrinking the window can only decrease the sum.

This property allows both pointers to move only forward, guaranteeing an **O(n)** solution.

### Complexity

| Metric | Value |
|--------|-------|
| Time | **O(n)** |
| Space | **O(1)** |

### Interview Takeaway

If you see:

- A contiguous subarray,
- Positive integers,
- A minimum or maximum valid window,

then **Variable Size Sliding Window** should be one of the first patterns you consider.

---

# Quick Reference

| Question | Answer |
|----------|--------|
| Contiguous data? | ✅ Required |
| Positive values? | Often required for variable windows |
| Two pointers? | Yes |
| Running state? | Yes (sum, count, frequency, etc.) |
| Window expands? | Yes |
| Window shrinks? | Yes (variable window) |
| Typical Time Complexity | **O(n)** |
| Typical Space Complexity | **O(1)** to **O(k)** |

---

# Pattern Checklist

Before choosing Sliding Window, verify:

- ✅ The problem involves contiguous elements.
- ✅ The maintained window state can be updated incrementally.
- ✅ Recomputing overlapping ranges would be inefficient.
- ✅ The window boundaries move only forward.
- ✅ The problem asks for a minimum, maximum, longest, shortest, or count over a contiguous range.

If all of these conditions are satisfied, Sliding Window is likely the appropriate optimization.