# Dry Run — 274. H-Index

## Goal

Find the largest value `h` such that:

- At least `h` papers have received **at least `h` citations**.
- The remaining papers have received **no more than `h` citations**.

---

# Example Input

```text
citations = [3,0,6,1,5]
```

Expected Output:

```text
3
```

---

# Step 1 — Sort Citations

Sort in descending order.

Before:

```text
[3,0,6,1,5]
```

After:

```text
[6,5,3,1,0]
```

---

# Why Sorting Helps

After sorting:

```text
Index      0  1  2  3  4
Citation   6  5  3  1  0
```

At position `i`:

```text
i + 1
```

represents how many papers are currently being considered.

Example:

```text
Index 0 → 1 paper
Index 1 → 2 papers
Index 2 → 3 papers
Index 3 → 4 papers
Index 4 → 5 papers
```

We check:

```text
citations[i] >= i + 1
```

If true:

```text
At least (i + 1) papers
have at least (i + 1) citations.
```

---

# Visual Walkthrough

Sorted Array:

```text
[6,5,3,1,0]
```

```text
Position:   0   1   2   3   4
Citation:   6   5   3   1   0
Paper #:    1   2   3   4   5
```

---

# Iteration-by-Iteration Trace

## Initial State

| Variable | Value |
|-----------|--------|
| h | 0 |

---

## Iteration 1

### Current Position

```text
i = 0
```

### Citation Count

```text
6
```

### Papers Considered

```text
i + 1 = 1
```

### Check

```text
6 >= 1
```

✅ True

Update:

```text
h = 1
```

### State

| Variable | Value |
|-----------|--------|
| h | 1 |

---

## Iteration 2

### Current Position

```text
i = 1
```

### Citation Count

```text
5
```

### Papers Considered

```text
2
```

### Check

```text
5 >= 2
```

✅ True

Update:

```text
h = 2
```

### State

| Variable | Value |
|-----------|--------|
| h | 2 |

---

## Iteration 3

### Current Position

```text
i = 2
```

### Citation Count

```text
3
```

### Papers Considered

```text
3
```

### Check

```text
3 >= 3
```

✅ True

Update:

```text
h = 3
```

### State

| Variable | Value |
|-----------|--------|
| h | 3 |

---

## Iteration 4

### Current Position

```text
i = 3
```

### Citation Count

```text
1
```

### Papers Considered

```text
4
```

### Check

```text
1 >= 4
```

❌ False

Cannot increase H-index.

Since array is sorted descending:

```text
All remaining values
will be <= 1
```

Therefore:

```text
No future position can satisfy the condition.
```

Stop processing.

---

# Complete State Transition Table

| Step | Index | Citation | Papers Considered (i+1) | Condition | H |
|--------|--------|-----------|-------------------------|-----------|---|
| Start | - | - | - | - | 0 |
| 1 | 0 | 6 | 1 | 6 ≥ 1 ✅ | 1 |
| 2 | 1 | 5 | 2 | 5 ≥ 2 ✅ | 2 |
| 3 | 2 | 3 | 3 | 3 ≥ 3 ✅ | 3 |
| 4 | 3 | 1 | 4 | 1 ≥ 4 ❌ | 3 |
| End | - | - | - | Stop | 3 |

---

# Final Result

```text
H-Index = 3
```

---

# Visual Interpretation

Sorted citations:

```text
[6,5,3,1,0]
```

Check candidate H values:

---

## H = 1

Papers with at least 1 citation:

```text
4 papers
```

```text
4 >= 1
```

✅ Valid

---

## H = 2

Papers with at least 2 citations:

```text
3 papers
```

```text
3 >= 2
```

✅ Valid

---

## H = 3

Papers with at least 3 citations:

```text
3 papers
```

```text
3 >= 3
```

✅ Valid

---

## H = 4

Papers with at least 4 citations:

```text
2 papers
```

```text
2 >= 4
```

❌ Invalid

---

Largest valid value:

```text
3
```

---

# Edge Case Walkthrough

## Case 1

Input:

```text
[0]
```

Sorted:

```text
[0]
```

Check:

```text
0 >= 1
```

❌ False

Result:

```text
0
```

---

## Case 2

Input:

```text
[100]
```

Sorted:

```text
[100]
```

Check:

```text
100 >= 1
```

✅ True

Result:

```text
1
```

---

## Case 3

Input

```text
[4,4,4,4]
```

Sorted:

```text
[4,4,4,4]
```

Checks:

```text
4 >= 1 ✅
4 >= 2 ✅
4 >= 3 ✅
4 >= 4 ✅
```

Result:

```text
4
```

---

## Case 4

Input

```text
[0,0,0,0]
```

Checks:

```text
0 >= 1 ❌
```

Result:

```text
0
```

---

# Key Insight

The critical observation is:

```text
citations[i] >= i + 1
```

means:

"There are at least (i + 1) papers
with at least (i + 1) citations."
```

The largest index satisfying this condition determines the H-index.

---

# Complexity Recap

| Metric | Value |
|----------|--------|
| Time Complexity | O(n log n) |
| Space Complexity | O(1)* |

\* Ignoring sorting implementation overhead.

---

# Interview Summary

When explaining this problem:

1. Define H-index carefully.
2. Sort citations in descending order.
3. Interpret each position as the number of papers considered.
4. Check whether:

```text
citations[i] >= i + 1
```

5. Track the largest valid value.
6. Mention bucket sort optimization for O(n) time.

This progression demonstrates both problem-solving and optimization skills.