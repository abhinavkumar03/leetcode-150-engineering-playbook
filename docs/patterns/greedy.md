# Greedy Pattern

## Pattern Definition

The **Greedy Pattern** is an algorithmic strategy where we make the best possible local decision at each step with the expectation that these local optimal choices will lead to a globally optimal solution.

Unlike Dynamic Programming, Greedy algorithms do not revisit previous decisions or store subproblem results.

The key requirement for a greedy solution is:

```text
Greedy Choice Property
```

Meaning:

```text
A locally optimal choice
→
Leads to a globally optimal solution
```

---

## When To Use Greedy

Consider a Greedy approach when:

* The problem asks for maximum or minimum optimization.
* Decisions are made sequentially.
* Earlier choices do not need to be reconsidered.
* Local optimal decisions contribute directly to the final answer.
* The solution can be built incrementally.

Common keywords:

```text
Maximum
Minimum
Fewest
Largest
Smallest
Optimal
Profit
Cost
Intervals
Scheduling
Reachability
```

---

## Recognition Signals

### Signal 1: Local Decisions Determine the Answer

Example:

```text
Take the best option available now.
```

---

### Signal 2: No Need To Backtrack

If choosing an option now never hurts future decisions, Greedy is often possible.

---

### Signal 3: Continuous Optimization

Problems involving:

* Profit maximization
* Cost minimization
* Resource allocation
* Scheduling

frequently use Greedy.

---

### Signal 4: Single Pass Solutions

If an optimal solution emerges while scanning the data once, Greedy is a strong candidate.

---

## Generic Greedy Template

### Selection-Based Greedy

```text
answer = initialValue

for each item:

    if currentChoice improves answer:
        take currentChoice

return answer
```

---

### Profit Accumulation Greedy

```text
profit = 0

for i from 1 to n-1:

    if gain exists:
        profit += gain

return profit
```

---

### Interval Greedy

```text
sort intervals

choose best interval

continue selecting valid intervals
```

---

## Complexity Characteristics

| Operation        | Typical Complexity |
| ---------------- | ------------------ |
| Traversal        | O(n)               |
| Sorting + Greedy | O(n log n)         |
| Extra Space      | O(1) to O(n)       |

---

## Correctness Requirements

A Greedy solution should satisfy at least one of the following:

### Greedy Choice Property

```text
Local optimum
=
Part of global optimum
```

---

### Optimal Substructure

```text
Optimal solution
=
Optimal choice
+
Optimal remaining solution
```

---

## Common Pitfalls

### Pitfall 1

Assuming Greedy Always Works

Not every optimization problem can be solved greedily.

Example:

```text
0/1 Knapsack
```

requires Dynamic Programming.

---

### Pitfall 2

Skipping Proof of Correctness

Interviewers often ask:

> Why does this greedy choice produce the optimal answer?

Always justify the decision.

---

### Pitfall 3

Confusing Greedy With Dynamic Programming

Greedy:

```text
Make choice once
Never revisit
```

DP:

```text
Explore states
Store results
```

---

### Pitfall 4

Choosing Based on Intuition Alone

A solution that "feels correct" is not enough.

Provide reasoning or proof.

---

# Problem Added

## 122. Best Time to Buy and Sell Stock II

### Difficulty

Medium

### Pattern Usage

This problem is a classic example of:

```text
Profit Accumulation Greedy
```

Observation:

```text
Every positive day-to-day increase
contributes directly to total profit.
```

Instead of finding explicit buy and sell points:

```text
Buy Low
Sell High
```

we accumulate every profitable increase.

---

### Greedy Rule

Whenever:

prices[i] > prices[i-1]

Add:

prices[i]-prices[i-1]

to the answer.

---

### Implementation Template

```text
profit = 0

for i from 1 to n-1:

    if prices[i] > prices[i-1]:
        profit += prices[i] - prices[i-1]

return profit
```

---

### Complexity

| Metric | Value |
| ------ | ----- |
| Time   | O(n)  |
| Space  | O(1)  |

---

### Key Interview Insight

For an increasing sequence:

```text
1 → 3 → 5 → 8
```

Single transaction:

```text
8 - 1 = 7
```

Greedy accumulation:

```text
(3-1)
+
(5-3)
+
(8-5)
=
7
```

Thus:

```text
Entire trend profit
=
Sum of daily gains
```

which proves the greedy strategy is optimal.

---

# Related Greedy Problems

## Easy

| Problem                              | Pattern            |
| ------------------------------------ | ------------------ |
| 121. Best Time to Buy and Sell Stock | Running Minimum    |
| 860. Lemonade Change                 | Transaction Greedy |

---

## Medium

| Problem                                         | Pattern                    |
| ----------------------------------------------- | -------------------------- |
| 45. Jump Game II                                | Reachability Greedy        |
| 55. Jump Game                                   | Reachability Greedy        |
| 122. Best Time to Buy and Sell Stock II         | Profit Accumulation Greedy |
| 134. Gas Station                                | Circular Greedy            |
| 435. Non-overlapping Intervals                  | Interval Scheduling Greedy |
| 452. Minimum Number of Arrows to Burst Balloons | Interval Greedy            |
| 763. Partition Labels                           | Boundary Expansion Greedy  |

---

## Hard

| Problem                                | Pattern              |
| -------------------------------------- | -------------------- |
| 135. Candy                             | Bidirectional Greedy |
| 321. Create Maximum Number             | Monotonic Greedy     |
| 871. Minimum Number of Refueling Stops | Heap + Greedy        |

---

# Quick Recognition Checklist

Before choosing Greedy, ask:

```text
□ Can I make the best local choice?
□ Will that choice never need to be changed?
□ Does a local optimum contribute to the global optimum?
□ Can I solve the problem in one pass?
□ Can I prove why the choice is safe?
```

If most answers are **Yes**, Greedy is likely the correct pattern.

---

# Pattern Summary

```text
Pattern:
Greedy

Recognition:
Local optimum contributes to global optimum

Template:
Take best choice now

Time:
Usually O(n) or O(n log n)

Space:
Usually O(1)

Representative Problem:
122. Best Time to Buy and Sell Stock II

Interview Focus:
Justifying why the greedy choice is correct
```