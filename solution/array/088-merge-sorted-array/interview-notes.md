# Merge Sorted Array — Interview Notes

## What Interviewer Is Testing

Although this is labeled as an Easy problem, interviewers are often evaluating much more than whether you can merge two arrays.

### Core Skills Being Evaluated

#### 1. Two Pointer Recognition

Can you immediately identify that two sorted arrays can be processed using pointers instead of nested loops?

Expected observation:

```text
Both arrays are sorted.
A merge process is possible.
```

---

#### 2. Space Optimization

Many candidates solve the problem using:

```text
Extra Array → Merge → Copy Back
```

Interviewers want to see whether you can improve:

```text
O(m+n) Space
```

to:

```text
O(1) Space
```

---

#### 3. In-Place Manipulation

Can you safely modify an existing data structure without losing information?

This is a common skill in:

* Arrays
* Strings
* Linked Lists
* Memory-sensitive systems

---

#### 4. Edge Case Thinking

Interviewers look for awareness of:

```text
nums2 empty
nums1 empty
duplicates
negative numbers
single element arrays
```

before coding.

---

## Typical Follow-up Questions

### Follow-up 1

Why can't we merge from the beginning?

Expected answer:

```text
Because nums1 contains valid values that would be overwritten
before they are processed.
```

Example:

```text
nums1 = [1,2,3,0,0,0]
```

Writing from index 0 destroys data.

---

### Follow-up 2

Why does merging from the end work?

Expected answer:

```text
The empty buffer exists at the end of nums1.

We always place the largest remaining element there.
```

No useful data is overwritten.

---

### Follow-up 3

Why don't we need to copy remaining nums1 elements?

Expected answer:

```text
If nums2 is exhausted,
the remaining nums1 values are already sorted
and already positioned correctly.
```

---

### Follow-up 4

What is the complexity?

Expected answer:

```text
Time  : O(m+n)
Space : O(1)
```

---

### Follow-up 5

Can this be generalized?

Expected answer:

The same merge idea appears in:

* Merge Sort
* Merging Sorted Lists
* External Sorting
* Stream Processing
* Database Query Engines

---

## Optimization Journey

A strong candidate should explain the evolution of solutions.

### Solution 1 — Merge Then Sort

```text
Copy all elements
Sort entire array
```

Complexity:

```text
Time  : O((m+n) log(m+n))
Space : O(m+n)
```

Problems:

* Wastes sorting work
* Ignores sorted input

---

### Solution 2 — Temporary Array Merge

```text
Classic merge process
Store result separately
Copy back
```

Complexity:

```text
Time  : O(m+n)
Space : O(m+n)
```

Better time.

Still not optimal space.

---

### Solution 3 — Reverse Two Pointers (Optimal)

Use:

```text
i = m - 1
j = n - 1
k = m + n - 1
```

Fill array from right to left.

Complexity:

```text
Time  : O(m+n)
Space : O(1)
```

Interview-preferred solution.

---

## Whiteboard Strategy

### Step 1

State observations.

```text
Both arrays are sorted.
nums1 has extra capacity.
```

---

### Step 2

Discuss brute force briefly.

```text
Merge into temp array.
```

Mention complexity.

---

### Step 3

Identify optimization opportunity.

```text
Can we use the empty space inside nums1?
```

---

### Step 4

Discover overwrite problem.

```text
Forward merge destroys data.
```

---

### Step 5

Propose reverse traversal.

```text
Take largest element first.
```

This is usually the breakthrough moment.

---

### Step 6

Write pointer logic.

```text
i = m - 1
j = n - 1
k = m + n - 1
```

---

### Step 7

Walk through an example.

Interviewers often care more about explanation than typing speed.

---

## Communication Tips

### Good Explanation

> Since both arrays are already sorted, the largest remaining value must be at the end of one of the arrays. I'll compare those values and place the larger one into the last available position of nums1.

---

### Better Explanation

> If I merge from the front, I risk overwriting values that I haven't processed yet. By filling from the end, I only write into unused space, allowing an O(1) space solution.

---

### What To Avoid

Avoid saying:

```text
I memorized this solution.
```

Instead explain:

```text
Why it works.
```

Interviewers care about reasoning.

---

## Senior-Level Discussion Points

### Memory Efficiency

This solution demonstrates:

```text
Data reuse
Memory optimization
In-place transformation
```

Common themes in large-scale systems.

---

### Streaming Systems

The same idea appears when merging:

* Sorted event streams
* Search indexes
* Log files

---

### Database Systems

Query engines frequently merge:

```text
Sorted result sets
```

using similar techniques.

---

### Distributed Systems

Workers often produce sorted outputs.

A coordinator merges those outputs efficiently using the same concept.

---

## FAANG-Level Variations

### Variation 1

Merge K Sorted Arrays

Common solution:

* Min Heap

Complexity:

```text
O(N log K)
```

---

### Variation 2

Merge Two Sorted Linked Lists

Related problem:

LeetCode 21: Merge Two Sorted Lists

Tests pointer manipulation.

---

### Variation 3

Merge Intervals

Related problem:

LeetCode 56: Merge Intervals

Uses sorting + merging logic.

---

### Variation 4

External Sorting

Used when data does not fit in memory.

Requires merging sorted chunks from disk.

---

### Variation 5

In-Place Merge Without Extra Buffer

Much harder variation.

Requires advanced partitioning techniques.

Typically appears in senior-level interviews.

---

# Quick Interview Recap

### Pattern

```text
Two Pointers
```

### Key Insight

```text
Merge from the end.
```

### Why?

```text
Avoid overwriting valid values.
```

### Complexity

```text
Time  : O(m+n)
Space : O(1)
```

### Interview Sound Bite

> The critical observation is that nums1 already provides unused space at the end. By processing both arrays from right to left, we can merge them in-place while preserving all unprocessed values.