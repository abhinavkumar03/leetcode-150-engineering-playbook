# 274. H-Index

## Problem Statement

Given an array `citations` where `citations[i]` is the number of citations a researcher received for their ith paper, return the researcher's H-index.

According to the definition of H-index:

A researcher has an index `h` if:

- At least `h` papers have received at least `h` citations each.
- The remaining papers have no more than `h` citations each.

Return the maximum value of `h`.

### Example 1

Input:

```text
citations = [3,0,6,1,5]
```

Output:

```text
3
```

Explanation:

- The researcher has 5 papers.
- 3 papers have at least 3 citations.
- No larger value satisfies the definition.

### Example 2

Input:

```text
citations = [1,3,1]
```

Output:

```text
1
```

---

## Difficulty

**Medium**

---

## Tags

- Array
- Sorting
- Counting Sort
- Bucket Sort
- Greedy

---

## Pattern

### Primary Pattern

**Counting / Bucket Sort**

### Secondary Pattern

**Sorting + Threshold Evaluation**

---

## Intuition

The H-index is essentially asking:

> What is the largest number `h` such that at least `h` papers have `h` or more citations?

Instead of focusing on individual citation counts, we focus on how many papers satisfy a citation threshold.

A useful observation:

If papers are sorted in descending order, the position of a paper tells us how many papers have citations greater than or equal to it.

This transforms the problem into finding the largest valid threshold.

---

## Key Observation

After sorting citations in descending order:

```text
[6,5,3,1,0]
```

Index:

```text
0 1 2 3 4
```

Paper count considered:

```text
1 2 3 4 5
```

At each position:

```text
citations[i] >= i + 1
```

means there are at least `(i + 1)` papers with at least `(i + 1)` citations.

The largest valid value becomes the answer.

---

# Brute Force Approach

For every possible H-index from `n` down to `0`:

1. Count papers having citations ≥ h.
2. Verify whether count ≥ h.
3. Return the largest valid h.

---

## Algorithm

1. Let n be number of papers.
2. For h from n down to 0:
   - Count papers with citations ≥ h.
   - If count ≥ h:
     - Return h.
3. Return 0.

---

## Complexity

### Time Complexity

```text
O(n²)
```

### Space Complexity

```text
O(1)
```

---

## Limitations

- Repeated scanning of the array.
- Inefficient for large inputs.
- Does not utilize ordering information.

---

# Optimized Approach

## Sorting-Based Solution

Sort citations in descending order.

Traverse the sorted array and determine the largest position where:

```text
citations[i] >= i + 1
```

The answer becomes the maximum valid paper count.

---

### Algorithm

1. Sort citations descending.
2. Initialize h = 0.
3. Traverse array:
   - If citations[i] ≥ i + 1:
     - h = i + 1
4. Return h.

---

### Why It Works

After sorting:

- First paper has highest citations.
- First two papers have citations at least equal to second paper.
- First three papers have citations at least equal to third paper.

When:

```text
citations[i] >= i + 1
```

there are at least `(i + 1)` papers with `(i + 1)` citations.

The largest such value is the H-index.

---

### Complexity

#### Time Complexity

```text
O(n log n)
```

Sorting dominates.

#### Space Complexity

```text
O(1)
```

or

```text
O(log n)
```

depending on sorting implementation.

---

## Bucket Sort Optimization

Since H-index cannot exceed number of papers `n`:

- Any citation count larger than `n` can be grouped into bucket `n`.
- Count frequency of citation values.
- Traverse buckets from high to low.
- Maintain cumulative paper count.
- First position where cumulative count ≥ current index is the answer.

---

### Complexity

#### Time Complexity

```text
O(n)
```

#### Space Complexity

```text
O(n)
```

---

## Edge Cases

### Empty Input

```text
[]
```

Output:

```text
0
```

No papers exist.

---

### Single Element

```text
[5]
```

Output:

```text
1
```

One paper has at least one citation.

---

### All Zero Citations

```text
[0,0,0]
```

Output:

```text
0
```

No paper satisfies H-index > 0.

---

### Duplicate Values

```text
[4,4,4,4]
```

Output:

```text
4
```

All four papers have at least four citations.

---

### Negative Values

LeetCode constraints do not allow negative citations.

If encountered in real systems:

- Validate input.
- Reject invalid data.

---

### Large Inputs

```text
100000 papers
```

Sorting approach remains efficient:

```text
O(n log n)
```

Bucket solution achieves:

```text
O(n)
```

---

## Dry Run

Input:

```text
[3,0,6,1,5]
```

Sorted:

```text
[6,5,3,1,0]
```

| Index | Citations | Papers Considered | Condition | H |
|---------|---------|---------|---------|---------|
| 0 | 6 | 1 | 6 ≥ 1 | 1 |
| 1 | 5 | 2 | 5 ≥ 2 | 2 |
| 2 | 3 | 3 | 3 ≥ 3 | 3 |
| 3 | 1 | 4 | 1 ≥ 4 ❌ | 3 |
| 4 | 0 | 5 | 0 ≥ 5 ❌ | 3 |

Final Answer:

```text
3
```

---

## Common Mistakes

### Mistake 1

Using:

```text
citations[i] > i + 1
```

instead of:

```text
citations[i] >= i + 1
```

Equality is valid.

---

### Mistake 2

Sorting ascending and applying the same logic.

The condition changes when order changes.

---

### Mistake 3

Confusing citation count with paper count.

H-index measures:

```text
number of papers
```

not

```text
citation value itself
```

---

### Mistake 4

Returning first valid H-index.

We need the **maximum** valid H-index.

---

## Interview Discussion

Topics commonly discussed:

- Why sorting works.
- Alternative bucket-sort solution.
- Why H-index never exceeds number of papers.
- Complexity comparison between approaches.
- Handling extremely large citation counts.
- Tradeoffs between memory and runtime.

---

## Follow-up Questions

### Follow-up 1

Can you solve this in O(n)?

Expected answer:

- Bucket counting approach.

---

### Follow-up 2

Can you solve it without modifying input?

Expected answer:

- Copy array before sorting.

---

### Follow-up 3

What if citation values are extremely large?

Expected answer:

- Cap values at `n` in bucket solution.

---

### Follow-up 4

Can H-index exceed number of papers?

No.

Maximum H-index is exactly `n`.

---

## Real World Applications

### Academic Research Metrics

- Research impact measurement.
- Faculty evaluation.
- Publication ranking.

### Search Ranking Systems

Threshold-based scoring techniques.

### Reputation Systems

Determining influence based on activity counts.

### Analytics Platforms

Computing performance metrics from large datasets.

### Recommendation Engines

Ranking content creators using engagement thresholds.

---

## Related Problems

### Easy

- 169. Majority Element
- 217. Contains Duplicate

### Medium

- 274. H-Index
- 275. H-Index II
- 347. Top K Frequent Elements
- 451. Sort Characters By Frequency

### Hard

- 295. Find Median from Data Stream

---

## Key Takeaway

The most important insight is:

> H-index is a threshold problem.

Sort the papers (or count frequencies), then find the largest value where at least `h` papers have at least `h` citations.

This transforms a seemingly academic metric into a classic array and counting problem frequently used in technical interviews.