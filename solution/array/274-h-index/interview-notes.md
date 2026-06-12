# Interview Notes — 274. H-Index

---

# What Interviewer Is Testing

This problem appears simple but evaluates several important interview skills.

## 1. Problem Comprehension

The biggest challenge is understanding the H-index definition correctly.

Interviewers want to see whether you can translate a business/domain metric into an algorithmic requirement.

Many candidates immediately start coding before fully understanding:

```text
At least h papers have at least h citations.
```

A strong candidate first verifies the definition with examples.

---

## 2. Pattern Recognition

Interviewers expect you to recognize that this is a:

- Threshold problem
- Sorting problem
- Counting problem

The key insight is identifying that the answer depends on the relationship between:

```text
number of papers
```

and

```text
citation count
```

rather than the exact citation values themselves.

---

## 3. Optimization Thinking

A good interview progression is:

```text
Brute Force
     ↓
Sorting
     ↓
Bucket Counting
```

Interviewers care more about the optimization journey than memorizing the final solution.

---

## 4. Complexity Analysis

You should clearly explain:

### Sorting Solution

```text
Time: O(n log n)
Space: O(1)
```

### Bucket Solution

```text
Time: O(n)
Space: O(n)
```

And discuss tradeoffs.

---

## 5. Communication Skills

This problem is frequently used because candidates must explain reasoning clearly.

Strong communication often matters more than writing code quickly.

---

# Typical Follow-up Questions

---

## Follow-up 1

### Can you solve it without sorting?

Expected Answer:

Yes.

Use bucket counting.

Since:

```text
H-index ≤ number of papers
```

Any citation count larger than n can be grouped into bucket n.

---

## Follow-up 2

### Why can H-index never exceed n?

Because:

```text
n = total number of papers
```

To have:

```text
H-index = h
```

you need at least:

```text
h papers
```

Therefore:

```text
h ≤ n
```

---

## Follow-up 3

### What if citations contain very large values?

Example:

```text
[1000000, 5000000, 9000000]
```

Bucket solution still works because:

```text
all values > n
```

can be stored in:

```text
bucket[n]
```

---

## Follow-up 4

### Can input be modified?

If modification is not allowed:

```go
sorted := append([]int(nil), citations...)
```

Sort the copy.

---

## Follow-up 5

### Which solution would you implement in production?

Typical answer:

```text
Sorting solution
```

because:

- Easier to maintain
- Easier to review
- Lower bug risk

Use bucket counting only if profiling shows sorting is a bottleneck.

---

# Optimization Journey

---

## Step 1 — Brute Force

Try every possible H-index.

For each candidate:

```text
Count papers with citations ≥ h
```

### Complexity

```text
Time: O(n²)
Space: O(1)
```

---

## Step 2 — Sorting

Sort descending.

Check:

```text
citations[i] >= i + 1
```

### Complexity

```text
Time: O(n log n)
Space: O(1)
```

This is the most common interview solution.

---

## Step 3 — Bucket Counting

Observation:

```text
Answer cannot exceed n
```

Use:

```text
bucket[n + 1]
```

Store citation frequencies.

Accumulate counts from right to left.

### Complexity

```text
Time: O(n)
Space: O(n)
```

Optimal solution.

---

# Whiteboard Strategy

When solving on a whiteboard:

---

## Step 1

Write the H-index definition.

```text
At least h papers
have at least h citations.
```

---

## Step 2

Take example:

```text
[3,0,6,1,5]
```

Sort:

```text
[6,5,3,1,0]
```

---

## Step 3

Create table:

| Index | Citation | Papers |
|---------|---------|---------|
| 0 | 6 | 1 |
| 1 | 5 | 2 |
| 2 | 3 | 3 |
| 3 | 1 | 4 |
| 4 | 0 | 5 |

---

## Step 4

Explain:

```text
citation >= papers
```

is the condition being tested.

---

## Step 5

Convert into code.

This progression demonstrates structured thinking.

---

# Communication Tips

---

## Good Explanation

> After sorting in descending order, position i represents i + 1 papers. If citations[i] is at least i + 1, then there are at least i + 1 papers with at least i + 1 citations. The largest such value becomes the H-index.

---

## Avoid Saying

> I remember this solution.

Interviewers prefer reasoning over memorization.

---

## Explain Why Sorting Helps

Do not simply say:

```text
Sort and iterate.
```

Explain:

```text
Sorting lets us compare
paper count against citation count.
```

---

## Mention Tradeoffs

A strong answer includes:

```text
Sorting:
O(n log n)

Bucket:
O(n)
```

and explains why one might be chosen.

---

# Senior-Level Discussion Points

---

## Production Considerations

### Input Mutation

Sorting modifies input.

Ask:

```text
Can I mutate the array?
```

If not:

```go
copyArray := append([]int(nil), citations...)
```

---

## Memory Tradeoffs

Bucket solution improves runtime:

```text
O(n log n) → O(n)
```

but increases memory usage:

```text
O(1) → O(n)
```

---

## Readability vs Performance

Many teams choose:

```text
Sorting solution
```

because:

- Easier maintenance
- Easier onboarding
- Easier debugging

even if asymptotically slower.

---

## Scalability Discussion

For millions of papers:

- Sorting cost becomes significant.
- Bucket counting becomes attractive.
- Performance testing should drive the decision.

---

# FAANG-Level Variations

---

## Variation 1 — H-Index II

Problem:

**LeetCode 275**

Input is already sorted.

Expected Solution:

Binary Search

Complexity:

```text
O(log n)
```

---

## Variation 2 — Dynamic H-Index

Papers are continuously added.

Need:

```text
Insert
Update
Query H-index
```

Possible Data Structures:

- Balanced BST
- Segment Tree
- Fenwick Tree

---

## Variation 3 — Streaming Citations

Citations arrive in real time.

Possible Solutions:

- Heap
- Frequency Map
- Order Statistics Tree

---

## Variation 4 — Top Researchers System

Design a service that:

- Stores papers
- Tracks citations
- Computes H-index

Discussion Topics:

- Database indexing
- Aggregation pipelines
- Caching
- Event-driven updates

---

## Variation 5 — Citation Analytics Platform

Design considerations:

- Batch processing
- Distributed aggregation
- Incremental updates
- Materialized views
- Cache invalidation

---

# Red Flags Interviewers Notice

❌ Misunderstanding H-index definition

❌ Using citation values directly as answer

❌ Forgetting answer cannot exceed n

❌ Incorrect condition

```text
citations[i] > i + 1
```

instead of:

```text
citations[i] >= i + 1
```

❌ Unable to explain why sorting works

❌ Giving complexity without justification

---

# Interview Takeaway

This problem is less about coding and more about reasoning.

The strongest interview performance usually follows:

```text
Understand Definition
        ↓
Brute Force
        ↓
Sorting Solution
        ↓
Bucket Optimization
        ↓
Tradeoff Discussion
```

Candidates who clearly communicate this progression typically perform much better than candidates who jump directly to implementation.