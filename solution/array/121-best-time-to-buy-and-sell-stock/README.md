# Best Time to Buy and Sell Stock

## Problem Statement

You are given an array `prices` where `prices[i]` represents the price of a given stock on the `iᵗʰ` day.

You want to maximize your profit by choosing:

- One day to buy a stock
- A different future day to sell that stock

Return the maximum profit you can achieve.

If no profit is possible, return `0`.

### Example 1

Input:

```text
prices = [7,1,5,3,6,4]
```

Output:

```text
5
```

Explanation:

```text
Buy at 1
Sell at 6
Profit = 5
```

---

### Example 2

Input:

```text
prices = [7,6,4,3,1]
```

Output:

```text
0
```

Explanation:

```text
No profitable transaction exists.
```

---

## Difficulty

Easy

---

## Tags

- Array
- Greedy
- Dynamic Programming
- Sliding State
- Optimization

---

## Pattern

### Primary Pattern

Running Minimum

### Secondary Pattern

Greedy Optimization

---

## Intuition

A naive approach would compare every possible buy day with every possible sell day.

For every day:

- Buy on day i
- Sell on day j (j > i)

Calculate profit and keep the maximum.

However, this requires checking all pairs and results in O(n²) time.

Instead, while scanning the array from left to right:

- Keep track of the lowest price seen so far.
- Treat that price as the best buying opportunity.
- Calculate profit if selling today.
- Update maximum profit when beneficial.

This allows us to solve the problem in one pass.

---

## Key Observation

For each day:

```text
Profit = Current Price - Lowest Price Seen So Far
```

The problem becomes:

```text
Find the largest difference where
buy occurs before sell.
```

Instead of checking every pair, continuously maintain:

```text
minimumPrice
maximumProfit
```

---

## Brute Force Approach

### Algorithm

1. Iterate through every day as a potential buy day.
2. Iterate through all future days as sell days.
3. Compute profit.
4. Track maximum profit.
5. Return the result.

### Pseudocode

```text
maxProfit = 0

for i in prices:
    for j after i:
        profit = prices[j] - prices[i]
        maxProfit = max(maxProfit, profit)

return maxProfit
```

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n²) |
| Space | O(1) |

### Limitations

- Too slow for large inputs.
- Repeatedly recalculates unnecessary comparisons.
- Does not scale efficiently.

---

## Optimized Approach

### Algorithm

1. Initialize:

```text
minimumPrice = prices[0]
maximumProfit = 0
```

2. Traverse the array once.
3. Update minimum price when a lower value is found.
4. Calculate current profit:

```text
currentProfit = currentPrice - minimumPrice
```

5. Update maximum profit if larger.
6. Return maximum profit.

---

### Why It Works

At every index:

```text
minimumPrice
```

represents the best buying opportunity encountered so far.

Therefore:

```text
currentPrice - minimumPrice
```

is the best possible profit if selling today.

Checking this for every day guarantees that we discover the maximum achievable profit.

Since each element is processed once:

```text
Time = O(n)
Space = O(1)
```

---

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n) |
| Space | O(1) |

---

## Edge Cases

### Empty Input

```text
[]
```

Expected:

```text
0
```

---

### Single Element

```text
[5]
```

Expected:

```text
0
```

Cannot sell after buying.

---

### Strictly Decreasing Prices

```text
[7,6,5,4,3]
```

Expected:

```text
0
```

No profitable transaction exists.

---

### Duplicate Prices

```text
[3,3,3,3]
```

Expected:

```text
0
```

No gain can be achieved.

---

### Negative Values

Although stock prices are typically positive, if negatives are allowed:

```text
[-5,-3,-1]
```

Profit:

```text
4
```

The algorithm still works correctly.

---

### Large Inputs

```text
100,000+ elements
```

The optimized solution remains efficient due to linear complexity.

---

## Dry Run

Input:

```text
[7,1,5,3,6,4]
```

| Day | Price | Min Price So Far | Current Profit | Max Profit |
|------|--------|------------------|----------------|------------|
| 0 | 7 | 7 | 0 | 0 |
| 1 | 1 | 1 | 0 | 0 |
| 2 | 5 | 1 | 4 | 4 |
| 3 | 3 | 1 | 2 | 4 |
| 4 | 6 | 1 | 5 | 5 |
| 5 | 4 | 1 | 3 | 5 |

Final Answer:

```text
5
```

---

## Common Mistakes

### Mistake 1

Selling before buying.

Incorrect:

```text
Use global max and min.
```

This ignores ordering constraints.

---

### Mistake 2

Updating profit before updating minimum price.

Always ensure:

```text
minimumPrice
```

represents the smallest value seen before the current day.

---

### Mistake 3

Returning negative profit.

Example:

```text
[7,6,5]
```

Expected:

```text
0
```

Not:

```text
-1
```

---

### Mistake 4

Using nested loops unnecessarily.

This increases complexity from:

```text
O(n)
```

to

```text
O(n²)
```

---

## Interview Discussion

This question is commonly used to test:

- Greedy thinking
- Optimization skills
- State tracking
- Complexity reduction
- Ability to derive a one-pass solution

A strong candidate should:

1. Present brute force.
2. Analyze complexity.
3. Identify redundant work.
4. Derive the running minimum approach.
5. Explain correctness clearly.

---

## Follow-up Questions

### Follow-up 1

What if multiple transactions are allowed?

Related Problem:

```text
LeetCode 122
Best Time to Buy and Sell Stock II
```

---

### Follow-up 2

What if there is a cooldown period?

Related Problem:

```text
LeetCode 309
```

---

### Follow-up 3

What if only two transactions are allowed?

Related Problem:

```text
LeetCode 123
```

---

### Follow-up 4

What if transaction fees exist?

Related Problem:

```text
LeetCode 714
```

---

## Real World Applications

### Financial Trading Systems

Determine optimal buy and sell points.

---

### Profit Maximization Analytics

Identify best entry and exit opportunities.

---

### Time-Series Optimization

Track minimum state and maximum gain over time.

---

### Resource Cost Monitoring

Find best purchase timing based on historical pricing.

---

### Business Intelligence

Detect maximum growth opportunity within a sequence of values.

---

## Related Problems

### Easier Foundations

- Two Sum
- Maximum Subarray
- Running Sum of 1D Array

### Similar Stock Problems

- Best Time to Buy and Sell Stock II
- Best Time to Buy and Sell Stock III
- Best Time to Buy and Sell Stock IV
- Best Time to Buy and Sell Stock with Cooldown
- Best Time to Buy and Sell Stock with Transaction Fee

### Similar Pattern Problems

- Maximum Subarray
- Maximum Difference Between Increasing Elements
- Longest Continuous Increasing Subsequence

---

## Summary

The key insight is to maintain the lowest stock price seen so far while scanning the array.

For every day:

```text
profit = currentPrice - minimumPrice
```

By updating the running minimum and maximum profit in a single traversal, we achieve:

```text
Time Complexity: O(n)
Space Complexity: O(1)
```

This is the optimal solution and a foundational interview pattern for state-tracking and greedy optimization problems.