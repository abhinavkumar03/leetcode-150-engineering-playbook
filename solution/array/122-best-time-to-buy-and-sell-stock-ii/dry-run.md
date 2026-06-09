# Dry Run — Best Time to Buy and Sell Stock II

## Goal

Given an array of stock prices, determine the maximum profit that can be achieved by making as many transactions as desired.

Rules:

* You may buy and sell multiple times.
* You can only hold one stock at a time.
* You must sell before buying again.

---

# Core Greedy Idea

Whenever:

```text
prices[i] > prices[i - 1]
```

add the profit:

```text
prices[i] - prices[i - 1]
```

to the answer.

This captures every profitable upward movement.

---

# Example 1

## Input

```text
prices = [7,1,5,3,6,4]
```

## Expected Output

```text
7
```

---

# Visual Price Movement

```text
Day:    0   1   2   3   4   5
Price:  7 → 1 → 5 → 3 → 6 → 4
            ↑     ↑
           +4    +3
```

---

# Initialization

| Variable | Value |
| -------- | ----- |
| profit   | 0     |

---

# Iteration Walkthrough

## Day 1

### Compare

```text
Current Price = 1
Previous Price = 7
```

### Difference

```text
1 - 7 = -6
```

Not profitable.

### Action

```text
Skip
```

### State

| Profit |
| ------ |
| 0      |

---

## Day 2

### Compare

```text
Current Price = 5
Previous Price = 1
```

### Difference

```text
5 - 1 = 4
```

Profitable.

### Action

```text
profit += 4
```

### State

| Profit |
| ------ |
| 4      |

---

## Day 3

### Compare

```text
Current Price = 3
Previous Price = 5
```

### Difference

```text
3 - 5 = -2
```

Not profitable.

### Action

```text
Skip
```

### State

| Profit |
| ------ |
| 4      |

---

## Day 4

### Compare

```text
Current Price = 6
Previous Price = 3
```

### Difference

```text
6 - 3 = 3
```

Profitable.

### Action

```text
profit += 3
```

### State

| Profit |
| ------ |
| 7      |

---

## Day 5

### Compare

```text
Current Price = 4
Previous Price = 6
```

### Difference

```text
4 - 6 = -2
```

Not profitable.

### Action

```text
Skip
```

### State

| Profit |
| ------ |
| 7      |

---

# Final Result

```text
Maximum Profit = 7
```

---

# Transaction Interpretation

The greedy solution implicitly performs:

```text
Buy  at 1
Sell at 5
Profit = 4

Buy  at 3
Sell at 6
Profit = 3
```

Total:

```text
4 + 3 = 7
```

---

# Example 2

## Input

```text
prices = [1,2,3,4,5]
```

## Expected Output

```text
4
```

---

# Step-by-Step Table

| Day | Previous | Current | Gain | Total Profit |
| --- | -------- | ------- | ---- | ------------ |
| 1   | 1        | 2       | 1    | 1            |
| 2   | 2        | 3       | 1    | 2            |
| 3   | 3        | 4       | 1    | 3            |
| 4   | 4        | 5       | 1    | 4            |

---

# Visual Explanation

```text
1 → 2 → 3 → 4 → 5
  +1  +1  +1  +1

Total = 4
```

Equivalent to:

```text
Buy at 1
Sell at 5

Profit = 4
```

The greedy approach captures the same profit.

---

# Example 3

## Input

```text
prices = [7,6,4,3,1]
```

## Expected Output

```text
0
```

---

# Step-by-Step Table

| Day | Previous | Current | Gain Added |
| --- | -------- | ------- | ---------- |
| 1   | 7        | 6       | 0          |
| 2   | 6        | 4       | 0          |
| 3   | 4        | 3       | 0          |
| 4   | 3        | 1       | 0          |

---

# Visual Explanation

```text
7 → 6 → 4 → 3 → 1
↓   ↓   ↓   ↓

No profitable increase
```

Result:

```text
Profit = 0
```

---

# State Transition View

## Decision Rule

```text
               prices[i] > prices[i-1] ?
                      |
            +---------+---------+
            |                   |
          Yes                  No
            |                   |
profit += difference       Do Nothing
            |                   |
            +---------+---------+
                      |
                Next Day
```

---

# Why Greedy Works

Consider:

```text
1 → 3 → 5 → 8
```

Single transaction:

```text
Buy 1
Sell 8

Profit = 7
```

Greedy accumulation:

```text
(3 - 1)
+
(5 - 3)
+
(8 - 5)

=
2 + 2 + 3

=
7
```

Both produce the same profit.

Therefore:

```text
Total upward movement
=
Maximum achievable profit
```

---

# Dry Run Summary

| Property             | Value                        |
| -------------------- | ---------------------------- |
| Pattern              | Greedy                       |
| Traversals           | 1                            |
| Extra Space          | O(1)                         |
| Time Complexity      | O(n)                         |
| Key Insight          | Sum all positive differences |
| Interview Difficulty | Medium                       |
| Optimization Level   | Optimal                      |
