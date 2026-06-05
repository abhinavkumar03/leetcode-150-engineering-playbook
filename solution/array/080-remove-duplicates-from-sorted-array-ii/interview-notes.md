# Interview Notes — Remove Duplicates from Sorted Array II

## What Interviewer Is Testing

This problem looks simple on the surface, but it evaluates several important engineering and interview skills.

### 1. Understanding of Sorted Arrays

The interviewer wants to see whether you recognize that:

```text
The array is already sorted.
```

This property is the foundation of the optimal solution.

Because duplicates are adjacent:

```text
1,1,1,2,2,3
```

we can process the array in a single pass without additional data structures.

---

### 2. Two-Pointer Technique

The key pattern is:

```text
Read Pointer
+
Write Pointer
```

Interviewers expect candidates to:

- Separate reading from writing.
- Maintain a valid output region.
- Modify the array in-place.

---

### 3. Space Optimization

A common beginner solution uses:

```text
HashMap
ArrayList
Extra Array
```

Interviewers intentionally include the constraint:

```text
O(1) extra space
```

to force candidates toward an in-place solution.

---

### 4. Invariant-Based Thinking

Strong candidates identify and maintain an invariant.

For this problem:

```text
nums[0...write-1]
```

always contains the valid result where every value appears at most twice.

This invariant remains true throughout execution.

---

### 5. Generalization Ability

Senior interviewers often ask:

> What if we allow K duplicates instead of 2?

Candidates who recognize the reusable pattern demonstrate stronger algorithmic thinking.

---

# Typical Follow-up Questions

## Follow-up 1

### Allow At Most K Duplicates

Instead of:

```text
nums[write - 2]
```

use:

```text
nums[write - K]
```

Template:

```text
if nums[read] != nums[write - K]
```

Complexity remains:

```text
Time:  O(n)
Space: O(1)
```

---

## Follow-up 2

### What If Array Is Not Sorted?

The current solution depends on duplicates being adjacent.

Without sorting:

```text
[1,2,1,3,1]
```

the comparison trick fails.

Possible solution:

```text
HashMap<Integer, Integer>
```

Complexity:

```text
Time:  O(n)
Space: O(n)
```

---

## Follow-up 3

### Remove All Duplicates

Requirement:

```text
Every element appears exactly once.
```

This becomes:

### LeetCode 26

Remove Duplicates from Sorted Array.

Allowed occurrences:

```text
1
```

instead of:

```text
2
```

---

## Follow-up 4

### Keep At Most Three Copies

Only one change:

```text
write = 3
```

and

```text
nums[write - 3]
```

Everything else remains identical.

---

## Follow-up 5

### Can We Do Better Than O(n)?

No.

Reason:

Every element must be examined at least once.

Lower bound:

```text
Ω(n)
```

Therefore:

```text
O(n)
```

is optimal.

---

# Optimization Journey

Interviewers love seeing the evolution of thought.

---

## Stage 1 — Brute Force

### Idea

Build a new array.

Track frequency.

Keep at most two copies.

Example:

```text
result = [1,1,2,2,3]
```

Complexity:

```text
Time:  O(n)
Space: O(n)
```

Problem:

```text
Extra memory.
```

---

## Stage 2 — Better Observation

Array is sorted.

Duplicates appear together.

Example:

```text
1,1,1,1,2,2,3
```

This means:

```text
No HashMap required.
```

---

## Stage 3 — In-Place Construction

Maintain:

```text
write pointer
```

that tracks the valid portion.

Overwrite unwanted values.

Complexity:

```text
Time:  O(n)
Space: O(1)
```

---

## Stage 4 — Elegant Constraint Check

Observation:

```text
A value may appear at most twice.
```

Therefore:

```text
Compare against
write - 2
```

This eliminates frequency counting completely.

Final solution:

```text
nums[read] != nums[write - 2]
```

---

# Whiteboard Strategy

When solving on a whiteboard:

### Step 1

Write the requirement.

```text
Keep at most two copies.
```

---

### Step 2

Draw a sample.

```text
[1,1,1,2,2,3]
```

---

### Step 3

Create pointers.

```text
read
write
```

---

### Step 4

Define invariant.

```text
nums[0...write-1]
```

contains valid output.

---

### Step 5

Explain comparison logic.

```text
nums[read] != nums[write - 2]
```

---

### Step 6

Walk through one example.

Interviewers often judge communication more than syntax.

---

# Communication Tips

## Good Explanation

> Since the array is sorted, duplicates are adjacent. I can use a read pointer to scan the array and a write pointer to build the valid result in-place. By comparing the current value with the element two positions before the write pointer, I can determine whether two copies already exist.

---

## Avoid Saying

```text
I memorized this pattern.
```

Instead explain:

```text
Why it works.
```

---

## Mention the Invariant

A strong signal for interviewers:

> At every step, the region before the write pointer already satisfies the problem constraints.

This demonstrates mature problem-solving skills.

---

# Senior-Level Discussion Points

## 1. Why Sorted Input Matters

The algorithm relies on:

```text
Duplicates being contiguous.
```

Without this property:

```text
O(1) space solution breaks.
```

---

## 2. Reusable Template

This is actually a generalized frequency-limited compression algorithm.

Template:

```text
Allow at most K occurrences.
```

Implementation:

```text
if nums[read] != nums[write - K]
```

---

## 3. Streaming Analogy

Think of:

```text
write pointer
```

as a stream output buffer.

The algorithm continuously decides:

```text
Accept
or
Reject
```

each incoming value.

---

## 4. Cache Efficiency

The solution:

- Traverses sequentially.
- Uses contiguous memory.
- Avoids allocations.

Benefits:

```text
Good cache locality.
```

---

## 5. Production Perspective

This pattern appears in:

- Data deduplication
- Event stream filtering
- Log compression
- Analytics preprocessing

---

# FAANG-Level Variations

## Variation 1

Allow At Most K Duplicates

Most common follow-up.

Expected answer:

```text
nums[write - K]
```

---

## Variation 2

Return Compressed Array

Instead of length only:

```text
Return resulting array.
```

Trivial modification.

---

## Variation 3

Keep Only Unique Values

Example:

```text
[1,1,2,2,3]
```

becomes:

```text
[3]
```

Requires counting frequencies.

---

## Variation 4

Unsorted Input

Need:

```text
HashMap
```

or

```text
Sorting + Two Pointers
```

Trade-off discussion expected.

---

## Variation 5

Streaming Data Version

Input arrives continuously.

Need:

```text
Online processing
```

Keep only K recent occurrences.

Tests system-design thinking.

---

# Red Flags Interviewers Notice

### Red Flag 1

Using extra arrays despite O(1) requirement.

---

### Red Flag 2

Not leveraging sorted input.

---

### Red Flag 3

Unable to explain:

```text
write - 2
```

---

### Red Flag 4

Memorizing solution without understanding invariant.

---

### Red Flag 5

Incorrect edge-case handling.

Examples:

```text
[]
[1]
[1,1]
```

---

# One-Minute Interview Summary

### Problem

Remove duplicates from a sorted array while allowing each value to appear at most twice.

### Pattern

Two Pointers

### Key Insight

Because the array is sorted, duplicates are adjacent.

Keep a write pointer representing the valid output.

Accept an element only when:

```text
nums[read] != nums[write - 2]
```

### Complexity

```text
Time:  O(n)
Space: O(1)
```

### Generalization

Replace:

```text
2
```

with:

```text
K
```

to allow at most K duplicates.

This turns the solution into a reusable frequency-limited array compression template.