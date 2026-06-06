# Majority Element — Dry Run

## Purpose

This document provides a detailed walkthrough of the Boyer-Moore Voting Algorithm used to solve LeetCode 169: Majority Element.

The goal is to understand:

- How the candidate is selected
- How pair cancellation works
- Why the majority element survives
- How the algorithm achieves O(n) time and O(1) space

---

# Algorithm Recap

We maintain two variables:

```text
candidate
count
```

Rules:

1. If count becomes 0:
   - Choose current number as candidate

2. If current number equals candidate:
   - Increment count

3. Otherwise:
   - Decrement count

---

# Example 1

## Input

```text
nums = [3,2,3]
```

## Initial State

| Candidate | Count |
|------------|--------|
| None | 0 |

---

## Iteration Walkthrough

### Step 1

Current Number:

```text
3
```

Count is 0.

Select new candidate:

```text
candidate = 3
count = 1
```

| Index | Value | Candidate | Count |
|---------|---------|---------|---------|
| 0 | 3 | 3 | 1 |

---

### Step 2

Current Number:

```text
2
```

Different from candidate.

```text
count--
```

| Index | Value | Candidate | Count |
|---------|---------|---------|---------|
| 1 | 2 | 3 | 0 |

Candidate survives but count resets.

---

### Step 3

Current Number:

```text
3
```

Count is 0.

Choose new candidate:

```text
candidate = 3
count = 1
```

| Index | Value | Candidate | Count |
|---------|---------|---------|---------|
| 2 | 3 | 3 | 1 |

---

## Final Result

```text
Majority Element = 3
```

---

# Example 2

## Input

```text
nums = [2,2,1,1,1,2,2]
```

Frequency:

```text
2 → 4 times
1 → 3 times
```

Majority:

```text
2
```

---

# Detailed Iteration Table

| Step | Number | Candidate Before | Count Before | Action | Candidate After | Count After |
|--------|--------|--------|--------|--------|--------|--------|
| 1 | 2 | None | 0 | New Candidate | 2 | 1 |
| 2 | 2 | 2 | 1 | Match | 2 | 2 |
| 3 | 1 | 2 | 2 | Cancel | 2 | 1 |
| 4 | 1 | 2 | 1 | Cancel | 2 | 0 |
| 5 | 1 | 2 | 0 | New Candidate | 1 | 1 |
| 6 | 2 | 1 | 1 | Cancel | 1 | 0 |
| 7 | 2 | 1 | 0 | New Candidate | 2 | 1 |

---

## Final Candidate

```text
2
```

Return:

```text
2
```

---

# Visual Pair Cancellation

Original Array

```text
[2,2,1,1,1,2,2]
```

Pair removals:

```text
2 cancels 1
2 cancels 1
2 cancels 1
```

Remaining:

```text
[2]
```

The majority element survives.

---

# Why Cancellation Works

Majority element occurs more than half the time.

Let:

```text
Majority Count = M
Other Elements = O
```

Condition:

```text
M > O
```

Each cancellation removes:

```text
1 majority element
1 non-majority element
```

Even after maximum cancellations:

```text
M - O > 0
```

At least one majority element remains.

Therefore:

```text
Final Candidate = Majority Element
```

---

# State Transition Visualization

## Input

```text
[2,2,1,1,1,2,2]
```

### Start

```text
candidate = -
count = 0
```

---

### Read 2

```text
candidate = 2
count = 1
```

---

### Read 2

```text
candidate = 2
count = 2
```

---

### Read 1

```text
candidate = 2
count = 1
```

---

### Read 1

```text
candidate = 2
count = 0
```

---

### Read 1

```text
candidate = 1
count = 1
```

---

### Read 2

```text
candidate = 1
count = 0
```

---

### Read 2

```text
candidate = 2
count = 1
```

---

### End

```text
return 2
```

---

# Edge Case Walkthrough

## Single Element

Input:

```text
[7]
```

### Step 1

```text
candidate = 7
count = 1
```

Return:

```text
7
```

---

## All Same Elements

Input:

```text
[5,5,5,5]
```

| Number | Candidate | Count |
|----------|----------|----------|
| 5 | 5 | 1 |
| 5 | 5 | 2 |
| 5 | 5 | 3 |
| 5 | 5 | 4 |

Return:

```text
5
```

---

## Negative Numbers

Input:

```text
[-1,-1,-1,2,3]
```

| Number | Candidate | Count |
|----------|----------|----------|
| -1 | -1 | 1 |
| -1 | -1 | 2 |
| -1 | -1 | 3 |
| 2 | -1 | 2 |
| 3 | -1 | 1 |

Return:

```text
-1
```

---

# Common Interview Explanation

A concise explanation:

> "I maintain a candidate and a vote count. Matching values increase the vote, while different values decrease it. Whenever the vote count reaches zero, I choose a new candidate. Because the majority element appears more than all other elements combined, it cannot be completely cancelled out and must remain as the final candidate."

---

# Complexity Verification

## Time Complexity

Single traversal:

```text
O(n)
```

---

## Space Complexity

Only two variables:

```text
candidate
count
```

Therefore:

```text
O(1)
```

---

# Final Takeaway

The Boyer-Moore Voting Algorithm works because every non-majority element can cancel out at most one occurrence of the majority element.

Since the majority element appears more than all remaining elements combined, it is guaranteed to survive the cancellation process and become the final candidate returned by the algorithm.