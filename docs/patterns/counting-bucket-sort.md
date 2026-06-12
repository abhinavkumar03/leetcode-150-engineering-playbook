# Counting / Bucket Sort Pattern

---

# Pattern Definition

The Counting / Bucket Sort pattern is used when:

- Values fall within a predictable range.
- We care about frequencies rather than exact ordering.
- Counting occurrences is more efficient than comparison-based sorting.
- The problem can be transformed into cumulative frequency analysis.

Instead of sorting elements directly, we count how many times each value appears and use those counts to derive the answer.

---

# When To Use This Pattern

Use Counting / Bucket Sort when:

✅ Values belong to a bounded range

✅ Maximum value can be capped

✅ Frequency matters more than ordering

✅ O(n log n) sorting can be replaced by O(n)

✅ Problem involves:

- Rankings
- Thresholds
- Frequency counting
- Distribution analysis

---

# Recognition Signals

Common interview clues:

### Signal 1

```text
Value range is limited
```

Example:

```text
0 <= value <= n
```

---

### Signal 2

```text
Need frequency counts
```

Examples:

- Top K elements
- Majority element
- H-index

---

### Signal 3

```text
Sorting feels unnecessary
```

If only counts matter:

```text
Count first
Sort later (or avoid sorting entirely)
```

---

### Signal 4

```text
Need cumulative frequencies
```

Examples:

```text
How many values are >= X?
How many values are <= X?
```

---

### Signal 5

```text
Answer depends on thresholds
```

Threshold problems frequently benefit from bucket counting.

---

# Generic Template

## Step 1

Create buckets.

```text
bucket[maxValue + 1]
```

---

## Step 2

Count frequencies.

```pseudo
for value in nums:
    bucket[value]++
```

---

## Step 3

Process buckets.

Options:

- Left to right
- Right to left
- Prefix sum
- Suffix sum

depending on problem requirements.

---

## Generic Pseudocode

```pseudo
bucket = new int[maxValue + 1]

for value in values:
    bucket[value]++

answer = ...

for i in range(...):
    process(bucket[i])

return answer
```

---

# Complexity

## Time Complexity

```text
O(n + k)
```

Where:

```text
n = number of elements
k = bucket range
```

---

## Space Complexity

```text
O(k)
```

---

# Advantages

### Fast

Often improves:

```text
O(n log n)
```

to

```text
O(n)
```

---

### Simple Frequency Tracking

Directly represents distribution.

---

### Great For Threshold Problems

Useful when answer depends on:

```text
count of elements
```

instead of:

```text
element ordering
```

---

# Common Pitfalls

---

## Pitfall 1

Huge Bucket Range

Bad:

```text
values up to 10^9
```

Creating buckets becomes impossible.

Use:

- HashMap
- TreeMap
- Heap

instead.

---

## Pitfall 2

Off-by-One Errors

Typical mistake:

```text
bucket[n]
```

vs

```text
bucket[n + 1]
```

Always verify bounds.

---

## Pitfall 3

Forgetting Value Compression

Sometimes:

```text
values > n
```

can be grouped.

Example:

H-index.

---

## Pitfall 4

Using Buckets When Sorting Is Easier

For small inputs:

```text
O(n log n)
```

may be preferable due to readability.

---

# Problem-Specific Example

## LeetCode 274 — H-Index

### Observation

For:

```text
n papers
```

maximum H-index is:

```text
n
```

Therefore:

```text
citations > n
```

can be grouped together.

---

### Bucket Construction

Example:

```text
citations = [3,0,6,1,5]
n = 5
```

Buckets:

| Citation | Count |
|-----------|---------|
| 0 | 1 |
| 1 | 1 |
| 2 | 0 |
| 3 | 1 |
| 4 | 0 |
| 5+ | 2 |

---

### Right-to-Left Accumulation

| h | Papers ≥ h |
|----|------------|
| 5 | 2 |
| 4 | 2 |
| 3 | 3 |

Condition:

```text
papers >= h
```

becomes true at:

```text
h = 3
```

Answer:

```text
3
```

---

# Related Problems

## Easy

### 169. Majority Element

Count occurrences.

---

### 242. Valid Anagram

Frequency comparison.

---

## Medium

### 274. H-Index

Threshold evaluation using buckets.

---

### 347. Top K Frequent Elements

Frequency buckets.

---

### 451. Sort Characters By Frequency

Frequency-based ordering.

---

### 75. Sort Colors

Counting sort variant.

---

## Hard

### 295. Find Median from Data Stream

Frequency concepts appear in alternative solutions.

---

# Comparison With Other Patterns

| Pattern | Use Case |
|----------|----------|
| Sorting | Need full ordering |
| Heap | Need top-k values |
| HashMap | Need frequency counts |
| Bucket Sort | Need bounded-frequency analysis |
| Binary Search | Need monotonic decision space |

---

# Interview Strategy

When you recognize:

```text
bounded values
+
frequency counting
+
threshold evaluation
```

consider:

```text
Bucket Sort
```

before reaching for comparison sorting.

A strong interview answer often follows:

```text
Brute Force
    ↓
Sorting
    ↓
Bucket Optimization
```

This demonstrates both algorithmic understanding and optimization skills.