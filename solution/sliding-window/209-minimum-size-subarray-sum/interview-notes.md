# Interview Notes — 209. Minimum Size Subarray Sum

---

# Problem Summary

Given an array of **positive integers** and a target value, find the **minimum length** of a contiguous subarray whose sum is greater than or equal to the target.

If no such subarray exists, return `0`.

---

# What Interviewer Is Testing

This problem evaluates much more than your ability to implement a Sliding Window.

## 1. Pattern Recognition ⭐⭐⭐⭐⭐

The interviewer wants to see whether you can immediately recognize:

- Variable Size Sliding Window
- Two Pointer Optimization
- Running Sum Maintenance

Recognizing the correct pattern is often more valuable than writing code quickly.

---

## 2. Optimization Skills ⭐⭐⭐⭐⭐

Most candidates naturally begin with:

```text
Brute Force
```

The interviewer expects you to optimize it by asking:

> "Can we avoid recalculating every subarray?"

Strong candidates identify that:

- Every element is positive.
- Window expansion only increases the sum.
- Window contraction only decreases the sum.

This observation enables an **O(n)** solution.

---

## 3. Constraint Analysis ⭐⭐⭐⭐☆

A common interview question is:

> Why does Sliding Window work here?

The correct explanation:

Because every number is positive.

If negative numbers were allowed:

- Expanding the window could decrease the sum.
- Shrinking the window could increase the sum.

The monotonic property would no longer hold, making the standard Sliding Window invalid.

Understanding **why** a pattern works is more important than memorizing it.

---

## 4. Complexity Analysis ⭐⭐⭐⭐⭐

The interviewer expects you to explain why the algorithm is linear despite the nested `while` loop.

A strong explanation:

- The `right` pointer visits each index exactly once.
- The `left` pointer also visits each index at most once.
- Therefore, the total pointer movements are bounded by `2n`.

```text
Time Complexity = O(n)
Space Complexity = O(1)
```

---

## 5. Code Quality ⭐⭐⭐⭐☆

Interviewers also assess whether your implementation is:

- Readable
- Correct
- Easy to explain
- Free of off-by-one errors
- Consistent in naming and formatting

---

# Typical Follow-up Questions

## Follow-up 1

**Can you solve it using Binary Search?**

Expected discussion:

- Build a Prefix Sum array.
- Binary search for the smallest valid ending index.
- Complexity:

```text
O(n log n)
```

Useful when introducing alternative approaches.

---

## Follow-up 2

**What if negative numbers are allowed?**

Sliding Window no longer works.

Possible alternatives include:

- Prefix Sum + Monotonic Queue (e.g., LeetCode 862)
- Advanced data structures depending on the problem constraints

---

## Follow-up 3

**Can you return the actual subarray instead of just its length?**

Maintain:

- Best starting index
- Best ending index

Whenever a smaller valid window is found.

---

## Follow-up 4

**How many valid subarrays exist?**

This becomes a counting problem.

The Sliding Window logic changes because every valid ending position may contribute multiple valid subarrays.

---

## Follow-up 5

**What if the array is extremely large?**

The algorithm already uses:

- Constant extra memory
- Single pass

It scales well and is suitable for large datasets.

---

# Optimization Journey

A common interview progression is:

## Stage 1 — Brute Force

### Idea

Check every possible subarray.

```text
for every start
    for every end
```

### Complexity

```text
Time: O(n²)
Space: O(1)
```

### Drawbacks

- Repeated computations
- Poor scalability

---

## Stage 2 — Key Observation

Notice that:

```text
All numbers are positive.
```

Therefore:

- Expanding the window always increases the sum.
- Shrinking the window always decreases the sum.

This is the turning point.

---

## Stage 3 — Sliding Window

Maintain:

```text
left
right
windowSum
minimumLength
```

Algorithm:

```text
Expand until sum >= target.

Shrink while still valid.

Repeat.
```

---

## Stage 4 — Optimal Solution

```text
Time: O(n)

Space: O(1)
```

Interviewers appreciate candidates who clearly articulate this optimization path.

---

# Whiteboard Strategy

When solving on a whiteboard:

## Step 1

Draw the array.

```text
2 3 1 2 4 3
```

---

## Step 2

Place the pointers.

```text
L
↓

2 3 1 2 4 3

↑
R
```

---

## Step 3

Track the running sum.

```text
Window Sum = 2
```

---

## Step 4

Expand the window.

```text
2
2+3
2+3+1
2+3+1+2
```

---

## Step 5

Shrink once the target is reached.

Show:

```text
Remove nums[left]
left++
```

Repeat until the window becomes invalid.

---

## Step 6

Continue expanding.

Repeat the process until the end of the array.

---

# Communication Tips

During the interview, explain your reasoning before writing code.

A structured explanation might be:

> "Since every element is positive, increasing the window size can only increase the sum. Once the sum reaches the target, I try to shrink the window to find the smallest valid subarray. Each pointer only moves forward, so the overall complexity is linear."

Avoid narrating every line of code. Focus on the algorithm's invariants and why it is correct.

---

# Senior-Level Discussion Points

## Why positivity matters

The correctness of the algorithm relies on monotonicity:

```text
Expand → Sum never decreases

Shrink → Sum never increases
```

This property guarantees that moving pointers only forward is sufficient.

---

## Loop Invariant

At the start of each outer-loop iteration:

- `windowSum` equals the sum of the current window `[left, right]`.
- Any window already discarded cannot lead to a shorter valid solution.
- `minLength` stores the smallest valid window seen so far.

Clearly stating loop invariants demonstrates strong algorithmic reasoning.

---

## Why `while` instead of `if`

Using `while` ensures the algorithm shrinks the window as much as possible before expanding again.

Replacing it with `if` may leave a larger-than-necessary valid window and miss the optimal answer.

---

## Space Optimization

The algorithm already uses the minimum additional memory required.

No auxiliary arrays, prefix sums, or hash maps are needed.

---

## Scalability

For an array with millions of elements:

- One sequential scan
- Constant memory
- Cache-friendly access pattern

This makes the approach practical in production systems handling large streams of data.

---

# FAANG-Level Variations

Interviewers at large tech companies may extend the problem with variations such as:

### 1. Allow Negative Numbers

Requires a different strategy, such as Prefix Sum with a Monotonic Queue.

---

### 2. Return the Subarray

Output the start and end indices (or the elements) of the smallest valid window.

---

### 3. Count Valid Subarrays

Instead of the minimum length, determine how many contiguous subarrays satisfy the target condition.

---

### 4. Circular Array

Adapt the solution when the array wraps around from the end to the beginning.

---

### 5. Streaming Input

Process numbers as they arrive without storing the entire array, maintaining only the current window.

---

### 6. Multiple Target Queries

If many target values are queried against the same array, discuss preprocessing techniques and trade-offs.

---

# Common Interview Mistakes

- Using `if` instead of `while` when shrinking the window.
- Forgetting to update the minimum length before removing the left element.
- Miscalculating the window length (`right - left + 1`).
- Returning the sentinel value instead of `0` when no valid subarray exists.
- Applying the Sliding Window technique to arrays containing negative numbers without verifying the constraints.

---

# Key Takeaways

- Recognize the **Variable Size Sliding Window** pattern quickly.
- Always justify **why** the pattern applies.
- Explain the optimization journey from **O(n²)** to **O(n)**.
- Emphasize that each pointer moves only forward.
- Clearly state loop invariants and edge cases.
- Connect the solution to real-world streaming and monitoring scenarios to demonstrate engineering maturity.