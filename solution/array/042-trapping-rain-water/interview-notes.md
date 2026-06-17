# Interview Notes — Trapping Rain Water

## What Interviewer Is Testing

### 1. Problem Decomposition

The interviewer wants to see whether you can transform a visual problem into a mathematical model.

Many candidates focus on the picture of bars and water. Strong candidates quickly identify:

```text
water[i] =
min(leftMax, rightMax)
- height[i]
```

The challenge is not calculating water itself.

The challenge is determining the required boundaries efficiently.

---

### 2. Optimization Skills

A common interview flow:

```text
Brute Force
    ↓
Prefix/Suffix Arrays
    ↓
Two Pointers
```

Interviewers want to observe:

- Can you identify redundant work?
- Can you reduce repeated scans?
- Can you improve space complexity?
- Can you justify the optimization?

---

### 3. Two-Pointer Reasoning

This problem is famous because many candidates can memorize the solution but cannot explain it.

Interviewers often ask:

> Why is it safe to move the left pointer when leftMax < rightMax?

A strong candidate can explain:

```text
The smaller boundary determines
the water level.
```

Therefore:

```text
min(leftMax, rightMax)
=
leftMax
```

and the current left position can be finalized immediately.

---

### 4. Complexity Analysis

Expected discussion:

| Approach | Time | Space |
|-----------|--------|--------|
| Brute Force | O(n²) | O(1) |
| DP Arrays | O(n) | O(n) |
| Stack | O(n) | O(n) |
| Two Pointers | O(n) | O(1) |

Interviewers expect candidates to compare tradeoffs.

---

### 5. Communication Ability

Strong candidates narrate their thought process:

```text
First I will solve it correctly.

Then I will optimize.

Then I will reduce memory usage.
```

This demonstrates engineering maturity.

---

# Typical Follow-up Questions

## Follow-up 1

### Can you solve it in O(n)?

Expected answer:

```text
Yes.
```

Use:

```text
leftMax[]
rightMax[]
```

arrays.

---

## Follow-up 2

### Can you reduce O(n) space?

Expected answer:

```text
Use Two Pointers.
```

Space becomes:

```text
O(1)
```

---

## Follow-up 3

### Can you solve it using a Stack?

Expected answer:

```text
Yes.
```

Use a monotonic decreasing stack.

Each valley can be evaluated when a higher bar appears.

---

## Follow-up 4

### Why does the smaller boundary matter?

Expected answer:

```text
Water overflows through
the shorter wall.
```

Therefore:

```text
waterLevel =
min(leftBoundary, rightBoundary)
```

---

## Follow-up 5

### What if bar widths vary?

Now the formula becomes:

```text
water =
heightDifference × width
```

The implementation must account for varying widths.

---

## Follow-up 6

### What if the elevation map is 2D?

This becomes:

```text
LeetCode 407
Trapping Rain Water II
```

Typical solution:

```text
Priority Queue + BFS
```

---

# Optimization Journey

## Stage 1 — Brute Force

For each position:

```text
Find left maximum.
Find right maximum.
```

Complexity:

```text
O(n²)
```

Reason:

Each index scans the array twice.

---

## Stage 2 — Dynamic Programming

Precompute:

```text
leftMax[]
rightMax[]
```

Example:

```text
Height:

[4,2,0,3,2,5]

Left:

[4,4,4,4,4,5]

Right:

[5,5,5,5,5,5]
```

Complexity:

```text
Time  O(n)
Space O(n)
```

---

## Stage 3 — Two Pointers

Observation:

```text
We do not need all
prefix/suffix values.
```

We only need:

```text
current leftMax
current rightMax
```

Complexity:

```text
Time  O(n)
Space O(1)
```

This is the optimal solution.

---

# Whiteboard Strategy

## Step 1

Draw an elevation map.

Example:

```text
4 2 0 3 2 5
```

---

## Step 2

Ask:

```text
What determines water
at one position?
```

Answer:

```text
Left boundary
Right boundary
```

---

## Step 3

Write the core formula:

```text
water[i] =
min(leftMax, rightMax)
- height[i]
```

---

## Step 4

Present the brute-force solution.

Interviewers usually prefer seeing a correct solution first.

---

## Step 5

Optimize using precomputed arrays.

---

## Step 6

Eliminate arrays using two pointers.

---

## Step 7

Explain pointer movement carefully.

Most interview mistakes occur here.

---

# Communication Tips

## Good Explanation

Say:

> Water is constrained by the shorter boundary. If the left boundary is shorter, I can determine the water level at the left pointer immediately.

This demonstrates understanding.

---

## Avoid Saying

```text
I memorized this solution.
```

or

```text
I know the trick.
```

Interviewers care about reasoning.

---

## Narrate Decisions

Example:

> I first solved it with O(n²). Then I noticed repeated scans. I stored left and right maximum values. Finally I realized only the current maxima are needed.

This creates a strong optimization narrative.

---

# Senior-Level Discussion Points

A senior engineer should go beyond coding.

---

## Tradeoff Analysis

Compare:

### DP Arrays

Pros:

```text
Easy to understand
Easy to debug
```

Cons:

```text
Extra memory
```

---

### Two Pointers

Pros:

```text
Optimal memory
Optimal runtime
```

Cons:

```text
More difficult reasoning
```

---

### Stack Solution

Pros:

```text
Useful for histogram problems
```

Cons:

```text
Harder implementation
```

---

## Production Considerations

Questions a senior engineer may discuss:

### Input Validation

```text
null input
small arrays
```

---

### Memory Constraints

For huge datasets:

```text
O(1) memory
```

becomes valuable.

---

### Maintainability

Sometimes:

```text
O(n) space
```

may be preferred if readability is more important than memory savings.

---

## Related Engineering Patterns

This problem appears in:

- Resource allocation
- Capacity planning
- Flood simulation
- Terrain analysis
- Histogram processing
- Computational geometry

---

# FAANG-Level Variations

## Variation 1

### Trapping Rain Water II

Problem:

2D grid.

Approach:

```text
Min Heap
+
BFS
```

Difficulty:

```text
Hard
```

---

## Variation 2

### Container With Most Water

Key Difference:

```text
Maximum area
```

instead of:

```text
Total trapped volume
```

Pattern:

```text
Two Pointers
```

---

## Variation 3

### Largest Rectangle in Histogram

Pattern:

```text
Monotonic Stack
```

Frequently discussed alongside this problem.

---

## Variation 4

### Daily Temperatures

Pattern:

```text
Monotonic Stack
```

Tests similar boundary reasoning.

---

## Variation 5

### Streaming Elevation Data

Suppose bars arrive continuously:

```text
1,4,2,5,3...
```

Questions:

- Can trapped water be updated incrementally?
- Can the algorithm work online?

This shifts the discussion toward system design and data streaming.

---

# Red Flags During Interviews

Interviewers often notice:

### Red Flag 1

Jumping directly to code.

---

### Red Flag 2

Unable to explain why the smaller boundary matters.

---

### Red Flag 3

Memorized solution without deriving it.

---

### Red Flag 4

Incorrect complexity analysis.

---

### Red Flag 5

Cannot compare alternative solutions.

---

# Interview Success Checklist

Before finishing, ensure you can explain:

- Why water depends on both sides.
- Why the minimum boundary determines water level.
- Why brute force is O(n²).
- How DP reduces time complexity.
- How two pointers reduce space complexity.
- Why the optimal solution is O(n) time and O(1) space.
- Edge cases.
- Tradeoffs between all approaches.

---

# Final Interview Summary

Core Formula:

```text
water[i] =
min(leftMax, rightMax)
- height[i]
```

Optimization Path:

```text
Brute Force
    ↓
DP Arrays
    ↓
Two Pointers
```

Expected Interview Solution:

```text
Time  : O(n)
Space : O(1)
```

Most Important Insight:

```text
The smaller boundary determines
the water level.
```