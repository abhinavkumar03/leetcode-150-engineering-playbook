# 209. Minimum Size Subarray Sum

## Problem Statement

Given an array of **positive integers** `nums` and a positive integer `target`, return the **minimal length** of a contiguous subarray whose sum is greater than or equal to `target`.

If there is no such subarray, return `0`.

### Example 1

**Input**

```text
target = 7
nums = [2,3,1,2,4,3]
```

**Output**

```text
2
```

**Explanation**

The subarray `[4,3]` has a sum of `7`, so the minimum length is `2`.

---

### Example 2

**Input**

```text
target = 4
nums = [1,4,4]
```

**Output**

```text
1
```

---

### Example 3

**Input**

```text
target = 11
nums = [1,1,1,1,1,1,1,1]
```

**Output**

```text
0
```

---

## Difficulty

**Medium**

---

## Tags

- Array
- Sliding Window
- Two Pointers
- Prefix Sum (Related Concept)
- Binary Search (Alternative Solution)

---

## Pattern

**Primary Pattern**

- Variable Size Sliding Window

**Secondary Pattern**

- Two Pointers

---

# Intuition

The brute force approach checks every possible subarray until it finds one whose sum reaches the target. Since there are **O(n²)** possible subarrays, this quickly becomes inefficient for large inputs.

A key observation is that **every number in the array is positive**.

This means:

- Expanding the window always increases the sum.
- Shrinking the window always decreases the sum.

Because of this monotonic behavior, we never need to move the left pointer backward. Each element enters and leaves the window at most once, allowing a linear-time solution.

---

# Key Observation

Since all numbers are positive:

- If the current window sum is less than the target, the only way to reach the target is to expand the window.
- Once the window sum reaches or exceeds the target, we should immediately try shrinking it to obtain the smallest valid window.

This "expand then shrink" strategy naturally leads to a Sliding Window solution.

---

# Brute Force Approach

Check every possible starting position and keep extending the subarray until the target is reached.

## Algorithm

1. Initialize answer as infinity.
2. For every starting index:
   - Initialize current sum to zero.
   - Extend the ending index one element at a time.
   - Add the current element to the sum.
   - If the sum becomes at least the target:
     - Update the minimum length.
     - Stop checking further endings for this start.
3. Return the minimum length if found; otherwise return `0`.

---

## Complexity

| Metric | Value |
|---------|-------|
| Time | **O(n²)** |
| Space | **O(1)** |

---

## Limitations

- Recomputes many overlapping subarrays.
- Inefficient for large arrays.
- Cannot satisfy optimal interview expectations.
- Does not leverage the positivity constraint of the input.

---

# Optimized Approach

Use a **variable-size Sliding Window**.

Maintain:

- `left` pointer
- `right` pointer
- Current window sum

Expand the window by moving the right pointer.

Whenever the window sum becomes at least the target:

- Record the current window length.
- Shrink the window from the left while the condition remains satisfied.

Continue until the entire array has been processed.

---

## Algorithm

1. Initialize:
   - `left = 0`
   - `windowSum = 0`
   - `minimumLength = ∞`
2. Iterate `right` from left to right.
3. Add `nums[right]` to the window sum.
4. While the window sum is greater than or equal to the target:
   - Update the minimum answer.
   - Remove `nums[left]`.
   - Move `left` forward.
5. If no valid window exists, return `0`; otherwise return the minimum length.

---

## Why It Works

Because all values are positive:

- Expanding never decreases the sum.
- Shrinking never increases the sum.

This guarantees that every possible minimal valid window is considered exactly once.

Each element:

- Enters the window once.
- Leaves the window once.

Therefore the total work is linear.

---

## Complexity

| Metric | Value |
|---------|-------|
| Time | **O(n)** |
| Space | **O(1)** |

---

# Edge Cases

| Case | Expected Result |
|------|-----------------|
| Empty array | `0` |
| Single element equals target | `1` |
| Single element less than target | `0` |
| Entire array required | Return array length |
| No valid subarray | `0` |
| Target already satisfied by first element | `1` |
| Multiple minimum windows | Return smallest length only |
| Very large input | Linear-time solution performs efficiently |
| Duplicate values | Works correctly |
| Negative values | **Sliding Window approach does NOT work** because this algorithm depends on all numbers being positive. A different approach is required when negative values are allowed. |

---

# Dry Run

### Input

```text
target = 7
nums = [2,3,1,2,4,3]
```

| Step | Right | Value Added | Window | Sum | Action | Minimum |
|------|------:|------------:|--------|----:|--------|---------|
| 1 | 0 | 2 | [2] | 2 | Expand | ∞ |
| 2 | 1 | 3 | [2,3] | 5 | Expand | ∞ |
| 3 | 2 | 1 | [2,3,1] | 6 | Expand | ∞ |
| 4 | 3 | 2 | [2,3,1,2] | 8 | Shrink | 4 |
| 5 | Left++ | -2 | [3,1,2] | 6 | Stop shrinking | 4 |
| 6 | 4 | 4 | [3,1,2,4] | 10 | Shrink | 4 |
| 7 | Left++ | -3 | [1,2,4] | 7 | Shrink | 3 |
| 8 | Left++ | -1 | [2,4] | 6 | Stop shrinking | 3 |
| 9 | 5 | 3 | [2,4,3] | 9 | Shrink | 3 |
| 10 | Left++ | -2 | [4,3] | 7 | Shrink | **2** |
| 11 | Left++ | -4 | [3] | 3 | Finish | **2** |

Final Answer:

```text
2
```

---

# Common Mistakes

### Forgetting to shrink repeatedly

Using `if` instead of `while` leaves the window larger than necessary.

---

### Updating the answer too late

Always record the current window length **before** removing the left element.

---

### Incorrect window length

The window size is:

```text
right - left + 1
```

---

### Returning infinity

If no valid window exists, return `0` instead of the initialized maximum value.

---

### Applying Sliding Window with negative numbers

This algorithm only works because every element is positive.

---

# Interview Discussion

Interviewers often expect candidates to progress through the following optimization journey:

1. Brute Force (`O(n²)`)
2. Observe positivity constraint.
3. Identify Sliding Window.
4. Implement variable-size window.
5. Explain why each pointer moves only forward.
6. Derive linear complexity.

Be prepared to justify **why the algorithm is correct**, not just how it works.

---

# Follow-up Questions

1. Can you solve this using Binary Search?
2. What changes if negative numbers are allowed?
3. Can this be generalized to finding the longest valid window?
4. How would you count all valid subarrays instead of finding the shortest?
5. What if the array were streamed instead of stored entirely in memory?
6. Can you solve the problem recursively? What are the trade-offs?

---

# Real World Applications

This Sliding Window pattern appears in many production systems:

- Network traffic monitoring over rolling intervals.
- Log analysis for detecting bursts of events.
- Financial analytics on moving transaction windows.
- CPU and memory utilization monitoring.
- Streaming analytics pipelines.
- Sensor data processing.
- Event aggregation systems.
- Time-series anomaly detection.

Understanding this pattern is valuable beyond coding interviews because it enables efficient processing of continuous data streams.

---

# Related Problems

| Problem | Pattern |
|---------|---------|
| 3. Longest Substring Without Repeating Characters | Sliding Window |
| 76. Minimum Window Substring | Sliding Window |
| 1004. Max Consecutive Ones III | Sliding Window |
| 424. Longest Repeating Character Replacement | Sliding Window |
| 904. Fruit Into Baskets | Variable Sliding Window |
| 1493. Longest Subarray of 1's After Deleting One Element | Sliding Window |
| 713. Subarray Product Less Than K | Sliding Window |
| 2090. K Radius Subarray Averages | Fixed Window |
| 643. Maximum Average Subarray I | Fixed Window |
| 1658. Minimum Operations to Reduce X to Zero | Sliding Window + Prefix Sum |