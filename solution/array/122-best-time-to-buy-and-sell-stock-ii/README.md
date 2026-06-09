# Best Time to Buy and Sell Stock II

## Problem Statement

You are given an integer array `prices` where `prices[i]` is the price of a given stock on the `iᵗʰ` day.

On each day, you may decide to:

* Buy one stock
* Sell the stock you currently hold

You can hold at most one stock at a time.

However, you may buy and sell the stock multiple times.

Return the maximum profit you can achieve.

---

## Difficulty

**Medium**

---

## Tags

* Array
* Greedy
* Dynamic Programming
* Stock Trading

---

## Pattern

### Primary Pattern

**Greedy**

### Secondary Pattern

* Peak & Valley
* Array Traversal
* Profit Accumulation

---

## Intuition

At first glance, the problem appears to require finding the best buy and sell pairs.

Many candidates try to:

* Find all local minima and maxima
* Use dynamic programming
* Simulate every transaction

However, there is a simpler observation.

Whenever the price increases from one day to the next, that increase can be safely added to the profit.

Instead of finding one large transaction:

```text
Buy at 1
Sell at 5
Profit = 4
```

We can split it into:

```text
Buy at 1 Sell at 2 => +1
Buy at 2 Sell at 3 => +1
Buy at 3 Sell at 5 => +2

Total Profit = 4
```

Both approaches yield the same profit.

Therefore, every positive difference contributes to the optimal answer.

---

## Key Observation

For every consecutive pair:

```text
prices[i] > prices[i - 1]
```

the profit gained is:

```text
prices[i] - prices[i - 1]
```

Adding all positive gains produces the maximum possible profit.

Example:

```text
Prices:
[1, 5, 3, 6, 8]

Profit:
(5 - 1) + (6 - 3) + (8 - 6)

= 4 + 3 + 2

= 9
```

This is equivalent to:

```text
Buy at 1 Sell at 5
Buy at 3 Sell at 8

Profit = 4 + 5 = 9
```

---

## Brute Force Approach

### Idea

Try every possible buy and sell combination recursively.

For each day:

* Buy
* Sell
* Skip

Explore all possibilities and return maximum profit.

### Algorithm

1. Start from day 0.
2. Decide:

   * Buy
   * Sell
   * Skip
3. Recurse for remaining days.
4. Return maximum achievable profit.

### Complexity

| Metric | Value |
| ------ | ----- |
| Time   | O(2ⁿ) |
| Space  | O(n)  |

### Limitations

* Exponential growth.
* Not practical for interview constraints.
* Generates many repeated states.
* Easily leads to TLE.

---

## Optimized Approach

### Idea

Add every positive price increase.

If today's price is higher than yesterday's price:

```text
profit += prices[i] - prices[i - 1]
```

This captures every profitable opportunity.

### Algorithm

1. Initialize `profit = 0`.
2. Traverse from day 1 to n-1.
3. If current price is greater than previous price:

   * Add the difference to profit.
4. Return total profit.

### Why It Works

Consider an increasing sequence:

```text
1 → 2 → 4 → 7
```

Profit from one transaction:

```text
7 - 1 = 6
```

Profit from accumulating gains:

```text
(2-1) + (4-2) + (7-4)

= 1 + 2 + 3

= 6
```

Both produce the same result.

Therefore, capturing every upward movement guarantees the maximum achievable profit.

This is a classic Greedy Optimization.

### Complexity

| Metric | Value |
| ------ | ----- |
| Time   | O(n)  |
| Space  | O(1)  |

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

No transactions possible.

---

### Single Element

```text
[5]
```

Output:

```text
0
```

Cannot sell after buying.

---

### Strictly Increasing Prices

```text
[1,2,3,4,5]
```

Output:

```text
4
```

Profit:

```text
(2-1)+(3-2)+(4-3)+(5-4)
```

---

### Strictly Decreasing Prices

```text
[7,6,4,3,1]
```

Output:

```text
0
```

No profitable transaction.

---

### Duplicate Prices

```text
[3,3,3,3]
```

Output:

```text
0
```

No gain exists.

---

### Negative Values (Hypothetical)

LeetCode constraints use positive stock prices, but if negatives existed:

```text
[-2,-1,4]
```

The same greedy rule still works.

---

### Large Inputs

```text
100,000+ prices
```

The O(n) solution scales efficiently.

---

## Dry Run

### Example

```text
prices = [7,1,5,3,6,4]
```

### Execution Table

| Day | Price | Previous Price | Profit Added | Total Profit |
| --- | ----- | -------------- | ------------ | ------------ |
| 0   | 7     | -              | -            | 0            |
| 1   | 1     | 7              | 0            | 0            |
| 2   | 5     | 1              | 4            | 4            |
| 3   | 3     | 5              | 0            | 4            |
| 4   | 6     | 3              | 3            | 7            |
| 5   | 4     | 6              | 0            | 7            |

Final Answer:

```text
7
```

---

### Visual Representation

```text
7 → 1 ↓

1 → 5 ↑ (+4)

5 → 3 ↓

3 → 6 ↑ (+3)

6 → 4 ↓

Total Profit = 7
```

---

## Common Mistakes

### Mistake 1: Searching for One Best Transaction

This is LeetCode 121 behavior.

Problem 122 allows:

```text
Multiple transactions
```

---

### Mistake 2: Buying Multiple Stocks

The problem allows:

```text
At most one stock at a time
```

---

### Mistake 3: Overengineering with DP

While DP works, Greedy is simpler and optimal.

---

### Mistake 4: Missing Consecutive Gains

Incorrect:

```text
Buy 1
Sell 5
Ignore intermediate gains
```

Correct:

```text
Accumulate all positive differences
```

---

## Interview Discussion

### Expected Interview Path

Interviewer typically expects:

1. Understand transaction rules.
2. Consider brute force.
3. Observe increasing sequences.
4. Derive greedy insight.
5. Implement O(n) solution.

---

### Discussion Topics

* Why Greedy works.
* Local vs global optimum.
* Comparison with Stock I.
* DP alternative formulation.
* Transaction constraints.

---

### Alternative DP View

State:

```text
hold
notHold
```

Transitions:

```text
hold = max(hold, notHold - price)

notHold = max(notHold, hold + price)
```

Produces the same answer but with additional complexity.

---

## Follow-up Questions

### 1. What changes if only one transaction is allowed?

Answer:

LeetCode 121.

Track minimum price and maximum profit.

---

### 2. What if there is a cooldown day?

Answer:

LeetCode 309.

Requires Dynamic Programming.

---

### 3. What if each transaction has a fee?

Answer:

LeetCode 714.

Modify profit transitions.

---

### 4. What if only k transactions are allowed?

Answer:

LeetCode 188.

DP with transaction count.

---

### 5. Can the solution be done in one pass?

Yes.

The Greedy solution is already one-pass.

---

## Real World Applications

### Financial Trading Systems

Identify profitable price movements.

---

### Inventory Arbitrage

Buy low and sell high repeatedly.

---

### Energy Market Optimization

Profit from fluctuating energy prices.

---

### Commodity Trading

Capture gains from multiple market swings.

---

### Resource Allocation Systems

Maximize value through repeated exchanges.

---

## Related Problems

### Easy

* 121. Best Time to Buy and Sell Stock

### Medium

* 714. Best Time to Buy and Sell Stock with Transaction Fee
* 309. Best Time to Buy and Sell Stock with Cooldown
* 188. Best Time to Buy and Sell Stock IV

### Pattern Relatives

* 55. Jump Game
* 45. Jump Game II
* 134. Gas Station
* 376. Wiggle Subsequence
