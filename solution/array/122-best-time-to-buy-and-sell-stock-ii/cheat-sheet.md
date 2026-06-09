# Best Time to Buy and Sell Stock II — Quick Revision Sheet

## Pattern Summary

### Primary Pattern

**Greedy**

### Secondary Patterns

* Array Traversal
* Profit Accumulation
* Peak & Valley
* Stock Trading

### Difficulty

**Medium**

### LeetCode

```text
122. Best Time to Buy and Sell Stock II
```

---

# Recognition Signals

Look for clues such as:

✅ Unlimited transactions allowed

✅ Buy and sell multiple times

✅ Only one stock can be held at a time

✅ Need maximum profit

✅ Array of prices over time

✅ Profit from repeated opportunities

Typical wording:

```text
You may complete as many transactions as you like.
```

This is the strongest indicator of the Greedy solution.

---

# Core Insight

Instead of finding:

```text
Best buy day
Best sell day
```

Think:

```text
Capture every profitable increase
```

Whenever:

```text
prices[i] > prices[i-1]
```

take the profit.

---

# Key Formula

Whenever today's price is greater than yesterday's:

Profit += \max(0,\ prices[i]-prices[i-1])

Equivalent implementation:

```text
if prices[i] > prices[i-1]:
    profit += prices[i] - prices[i-1]
```

---

# Why Greedy Works

Consider:

```text
1 → 3 → 5 → 8
```

Single transaction:

```text
Buy at 1
Sell at 8

Profit = 7
```

Greedy accumulation:

```text
(3-1)
+
(5-3)
+
(8-5)

=
2 + 2 + 3

=
7
```

Therefore:

```text
Entire upward trend
=
Sum of daily gains
```

This proves that taking every positive difference is optimal.

---

# Mental Model

Think of the stock chart as:

```text
Price Goes Up  → Collect Profit
Price Goes Down → Ignore
```

Visual:

```text
7 → 1 ↓
1 → 5 ↑ +4
5 → 3 ↓
3 → 6 ↑ +3
6 → 4 ↓

Total = 7
```

---

# Complexity Cheatsheet

| Approach            | Time  | Space |
| ------------------- | ----- | ----- |
| Brute Force         | O(2ⁿ) | O(n)  |
| Dynamic Programming | O(n)  | O(n)  |
| Greedy              | O(n)  | O(1)  |

Recommended:

```text
Greedy
```

---

# Implementation Template

### Generic Pseudocode

```text
profit = 0

for i from 1 to n-1:

    if prices[i] > prices[i-1]:
        profit += prices[i] - prices[i-1]

return profit
```

---

# Common Mistakes

## Mistake 1

Treating it like Stock I.

```text
Wrong:
Only one transaction
```

Problem 122 allows:

```text
Multiple transactions
```

---

## Mistake 2

Trying Sliding Window.

```text
No fixed-size window exists.
```

This is not a Sliding Window problem.

---

## Mistake 3

Overusing Dynamic Programming.

DP works but is unnecessary.

Greedy is simpler and optimal.

---

## Mistake 4

Searching for Peaks and Valleys Explicitly.

Example:

```text
Find local minimum
Find local maximum
```

Works but adds complexity.

Simply summing positive differences is cleaner.

---

# Edge Cases Checklist

Before submitting:

### Empty Array

```text
[]
```

Result:

```text
0
```

---

### One Element

```text
[5]
```

Result:

```text
0
```

---

### Increasing Prices

```text
[1,2,3,4,5]
```

Result:

```text
4
```

---

### Decreasing Prices

```text
[7,6,4,3,1]
```

Result:

```text
0
```

---

### Duplicate Prices

```text
[3,3,3]
```

Result:

```text
0
```

---

# Interview Sound Bites

### 15-Second Explanation

> Since unlimited transactions are allowed, every positive day-to-day increase contributes directly to the final profit. Therefore, summing all positive differences produces the maximum achievable profit.

---

### 30-Second Explanation

> Any increasing sequence can be split into multiple smaller profitable transactions without changing total profit. Because of this property, we do not need to find exact buy and sell points. We simply accumulate every positive price increase while traversing the array once.

---

# Similar Problems

## Stock Trading Series

| Problem                                            | Pattern         |
| -------------------------------------------------- | --------------- |
| 121. Best Time to Buy and Sell Stock               | Running Minimum |
| 122. Best Time to Buy and Sell Stock II            | Greedy          |
| 123. Best Time to Buy and Sell Stock III           | DP              |
| 188. Best Time to Buy and Sell Stock IV            | DP              |
| 309. Best Time to Buy and Sell Stock with Cooldown | DP              |
| 714. Best Time to Buy and Sell Stock with Fee      | DP              |

---

# Pattern Relatives

Problems that use similar Greedy thinking:

* 55. Jump Game
* 45. Jump Game II
* 134. Gas Station
* 376. Wiggle Subsequence
* 435. Non-overlapping Intervals

---

# One-Minute Revision

```text
Pattern:
Greedy

Observation:
Take every positive difference.

Formula:
profit += max(0, prices[i] - prices[i-1])

Time:
O(n)

Space:
O(1)

Why It Works:
Sum of all daily gains equals total gain of an upward trend.

Interview Keyword:
Local optimum forms global optimum.
```

---

# Memorization Trigger

If you see:

```text
Unlimited transactions allowed
```

Immediately think:

```text
Greedy

Sum all positive price differences.
```