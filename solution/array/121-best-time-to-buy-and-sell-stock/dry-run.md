# Dry Run — Best Time to Buy and Sell Stock

## Goal

Find the maximum profit that can be achieved from exactly one buy transaction followed by one sell transaction.

Rules:

- Buy before sell.
- Only one transaction allowed.
- If no profit is possible, return 0.

---

# Example Input

```text
prices = [7,1,5,3,6,4]
```

Expected Output:

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

# Core Idea

While traversing the array:

1. Track the lowest stock price seen so far.
2. Treat it as the best buying opportunity.
3. Calculate profit if selling today.
4. Keep updating the maximum profit.

Maintain:

```text
minPrice
maxProfit
```

---

# Initialization

Before processing:

| Variable | Value |
|-----------|---------|
| minPrice | 7 |
| maxProfit | 0 |

Visual State:

```text
Prices

[7, 1, 5, 3, 6, 4]
 ↑

minPrice = 7
maxProfit = 0
```

---

# Iteration Walkthrough

## Day 0

Price:

```text
7
```

Current State:

```text
minPrice = 7
maxProfit = 0
```

Profit if sold today:

```text
7 - 7 = 0
```

Updated State:

| Variable | Value |
|-----------|---------|
| minPrice | 7 |
| maxProfit | 0 |

---

## Day 1

Price:

```text
1
```

Compare with current minimum:

```text
1 < 7
```

New minimum found.

Update:

```text
minPrice = 1
```

Current State:

| Variable | Value |
|-----------|---------|
| minPrice | 1 |
| maxProfit | 0 |

Visual:

```text
[7, 1, 5, 3, 6, 4]
    ↑

Best Buy So Far = 1
```

---

## Day 2

Price:

```text
5
```

Potential Profit:

```text
5 - 1 = 4
```

Compare:

```text
max(0,4) = 4
```

Update:

```text
maxProfit = 4
```

State:

| Variable | Value |
|-----------|---------|
| minPrice | 1 |
| maxProfit | 4 |

Visual:

```text
Buy  = 1
Sell = 5

Profit = 4
```

---

## Day 3

Price:

```text
3
```

Potential Profit:

```text
3 - 1 = 2
```

Compare:

```text
max(4,2) = 4
```

No update.

State:

| Variable | Value |
|-----------|---------|
| minPrice | 1 |
| maxProfit | 4 |

Visual:

```text
Buy  = 1
Sell = 3

Profit = 2

Current Best Profit = 4
```

---

## Day 4

Price:

```text
6
```

Potential Profit:

```text
6 - 1 = 5
```

Compare:

```text
max(4,5) = 5
```

Update:

```text
maxProfit = 5
```

State:

| Variable | Value |
|-----------|---------|
| minPrice | 1 |
| maxProfit | 5 |

Visual:

```text
Buy  = 1
Sell = 6

Profit = 5
```

---

## Day 5

Price:

```text
4
```

Potential Profit:

```text
4 - 1 = 3
```

Compare:

```text
max(5,3) = 5
```

No update.

State:

| Variable | Value |
|-----------|---------|
| minPrice | 1 |
| maxProfit | 5 |

---

# Complete State Transition Table

| Day | Price | Min Price Before | Action | Current Profit | Max Profit |
|------|--------|------------------|----------|----------------|------------|
| 0 | 7 | 7 | Initialize | 0 | 0 |
| 1 | 1 | 7 | Update Min | 0 | 0 |
| 2 | 5 | 1 | Calculate Profit | 4 | 4 |
| 3 | 3 | 1 | Calculate Profit | 2 | 4 |
| 4 | 6 | 1 | Calculate Profit | 5 | 5 |
| 5 | 4 | 1 | Calculate Profit | 3 | 5 |

---

# Visual Timeline

```text
Day:    0  1  2  3  4  5
Price: [7, 1, 5, 3, 6, 4]

Lowest Price Seen
        ↓
       [1]

Best Sell Opportunity
                ↓
               [6]

Maximum Profit

6 - 1 = 5
```

---

# Why We Never Need Nested Loops

Brute Force:

```text
For every buy day
    Check every sell day
```

Complexity:

```text
O(n²)
```

Example:

```text
(7,1)
(7,5)
(7,3)
(7,6)
(7,4)

(1,5)
(1,3)
(1,6)
(1,4)

...
```

Many unnecessary comparisons.

---

# Optimized Thinking

Instead of remembering every previous price:

Store only:

```text
minimum price seen so far
```

Then:

```text
profit = currentPrice - minimumPrice
```

This immediately gives the best profit if we sell today.

Result:

```text
Single Pass
O(n) Time
O(1) Space
```

---

# Edge Case Walkthrough

## Case 1

```text
[7,6,4,3,1]
```

### Traversal

| Price | Min Price | Profit | Max Profit |
|---------|------------|----------|------------|
| 7 | 7 | 0 | 0 |
| 6 | 6 | 0 | 0 |
| 4 | 4 | 0 | 0 |
| 3 | 3 | 0 | 0 |
| 1 | 1 | 0 | 0 |

Answer:

```text
0
```

No profitable transaction exists.

---

## Case 2

```text
[5]
```

Only one day.

Cannot sell after buying.

Answer:

```text
0
```

---

## Case 3

```text
[2,2,2,2]
```

Profit:

```text
2 - 2 = 0
```

Answer:

```text
0
```

---

# Final Result

Input:

```text
[7,1,5,3,6,4]
```

Best Transaction:

```text
Buy  at 1
Sell at 6
```

Maximum Profit:

```text
5
```

---

# Key Takeaway

This problem demonstrates a powerful interview pattern:

```text
Track the best state so far.
```

Here the state is:

```text
Minimum Price Seen So Far
```

Using that state allows us to convert a brute-force O(n²) solution into an optimal:

```text
Time  = O(n)
Space = O(1)
```