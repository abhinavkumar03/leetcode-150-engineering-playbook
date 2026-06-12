# 274. H-Index — Cheat Sheet

---

# Pattern Summary

### Primary Pattern

**Counting / Bucket Sort**

### Secondary Pattern

**Sorting + Threshold Evaluation**

### Difficulty

Medium

---

# Recognition Signals

Look for these clues:

✅ Find the **largest valid value** satisfying a condition

✅ "At least X elements have value ≥ X"

✅ Threshold-based metrics

✅ Ranking or scoring systems

✅ Need to compare:

```text
count of elements
vs
element values
```

✅ Problem can be transformed by sorting

---

# Core Insight

H-index asks:

```text
Largest h such that
at least h papers
have at least h citations.
```

The answer is NOT:

```text
maximum citation count
```

The answer is a threshold.

---

# Sorting Approach Template

## Step 1

Sort descending.

```text
[6,5,3,1,0]
```

---

## Step 2

Traverse array.

For each position:

```text
papers = i + 1
```

---

## Step 3

Check:

```text
citations[i] >= papers
```

---

## Step 4

Update answer.

```text
h = papers
```

---

## Step 5

Largest valid h is the answer.

---

# Mental Model

After sorting:

```text
Index      0  1  2  3  4
Citation   6  5  3  1  0
```

Interpretation:

```text
1 paper has ≥ 6 citations
2 papers have ≥ 5 citations
3 papers have ≥ 3 citations
4 papers have ≥ 1 citation
5 papers have ≥ 0 citations
```

Check where:

```text
citation >= paper count
```

remains true.

---

# Key Formula

The critical condition:

:contentReference[oaicite:0]{index=0}

If true:

```text
There are at least (i + 1) papers
with at least (i + 1) citations.
```

---

# H-Index Constraints

For:

```text
n papers
```

Maximum possible H-index:

:contentReference[oaicite:1]{index=1}

Reason:

Need at least h papers to achieve H-index h.

---

# Bucket Sort Insight

Observation:

```text
Any citation value > n
does not matter individually.
```

Because:

```text
H-index cannot exceed n.
```

Store all values greater than n inside:

```text
bucket[n]
```

---

# Bucket Sort Template

```text
bucket = int[n + 1]

for citation:
    if citation >= n:
        bucket[n]++
    else:
        bucket[citation]++

papers = 0

for h from n down to 0:
    papers += bucket[h]

    if papers >= h:
        return h
```

---

# Complexity Cheatsheet

| Approach | Time | Space |
|-----------|--------|--------|
| Brute Force | O(n²) | O(1) |
| Sorting | O(n log n) | O(1)* |
| Bucket Sort | O(n) | O(n) |

\* Ignoring sorting implementation overhead.

---

# Optimization Journey

```text
Brute Force
    ↓
Sort Descending
    ↓
Threshold Check
    ↓
Bucket Counting
```

Interviewers often care more about this progression than the final code.

---

# Common Mistakes

### Mistake #1

Using:

```text
citation > i + 1
```

Instead of:

```text
citation >= i + 1
```

---

### Mistake #2

Returning citation value.

Wrong:

```text
return citations[i]
```

Correct:

```text
return i + 1
```

---

### Mistake #3

Confusing paper count and citation count.

Always remember:

```text
H-index measures papers.
```

---

### Mistake #4

Forgetting answer cannot exceed:

```text
n
```

---

### Mistake #5

Not handling:

```text
[0,0,0]
```

Answer:

```text
0
```

---

# Edge Cases

## Empty Array

```text
[]
```

Answer:

```text
0
```

---

## Single Paper

```text
[10]
```

Answer:

```text
1
```

---

## All Zeroes

```text
[0,0,0]
```

Answer:

```text
0
```

---

## Same Citations

```text
[4,4,4,4]
```

Answer:

```text
4
```

---

## Large Citation Values

```text
[1000,2000,3000]
```

Answer:

```text
3
```

---

# Interview Sound Bites

Useful statements during interviews:

### Sound Bite #1

> H-index is a threshold problem rather than a maximum-value problem.

---

### Sound Bite #2

> After sorting, position i represents i + 1 papers.

---

### Sound Bite #3

> The condition citations[i] >= i + 1 directly matches the H-index definition.

---

### Sound Bite #4

> Since H-index cannot exceed n, bucket counting provides an O(n) solution.

---

### Sound Bite #5

> I would implement the sorting solution first and mention bucket optimization afterward.

---

# Similar Problems

### Same Threshold Thinking

- 275. H-Index II
- 128. Longest Consecutive Sequence
- 347. Top K Frequent Elements
- 215. Kth Largest Element in an Array

---

### Same Counting Pattern

- 347. Top K Frequent Elements
- 451. Sort Characters By Frequency
- 75. Sort Colors
- 274. H-Index

---

### Same Sorting Pattern

- 56. Merge Intervals
- 179. Largest Number
- 435. Non-overlapping Intervals

---

# 30-Second Revision

1. Sort citations descending.

2. For each index:

```text
papers = i + 1
```

3. Check:

```text
citations[i] >= papers
```

4. Update:

```text
h = papers
```

5. Largest valid h is answer.

Complexities:

```text
Sorting:
O(n log n)

Bucket:
O(n)
```

Core idea:

```text
Count papers,
not citations.
```