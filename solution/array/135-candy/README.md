# Candy

## Problem Statement

There are `n` children standing in a line. Each child is assigned a rating value.

You are giving candies to these children subject to the following requirements:

1. Each child must have at least one candy.
2. Children with a higher rating than their immediate neighbor must receive more candies than that neighbor.

Return the minimum number of candies needed to distribute.

---

## Difficulty

Hard

---

## Tags

- Array
- Greedy
- Two Pass Traversal

---

## Pattern

**Greedy + Bidirectional Constraint Satisfaction**

This problem involves satisfying local constraints between neighboring elements while minimizing the total resource allocation.

The key observation is that rating relationships must be respected from both directions:

- Left → Right
- Right → Left

A single traversal cannot fully satisfy both constraints.

---

## Intuition

Suppose every child initially receives one candy.

Whenever a child has a higher rating than the child on the left, they should receive more candies.

Likewise, whenever a child has a higher rating than the child on the right, they should also receive more candies.

The challenge is that satisfying one direction can violate requirements in the opposite direction.

Therefore, we must process both directions independently and combine the results.

---

## Key Observation

Two separate constraints exist:

### Left Constraint

If:

```text
ratings[i] > ratings[i - 1]
```

then:

```text
candies[i] > candies[i - 1]
```

### Right Constraint

If:

```text
ratings[i] > ratings[i + 1]
```

then:

```text
candies[i] > candies[i + 1]
```

A left-to-right pass handles the first rule.

A right-to-left pass handles the second rule.

The final candy count for each child must satisfy both rules simultaneously.

Therefore:

```text
candies[i] = max(
    leftRequirement,
    rightRequirement
)
```

---

## Brute Force Approach

Repeatedly adjust candy counts until all constraints are satisfied.

### Algorithm

1. Give every child one candy.
2. Traverse the array repeatedly.
3. If a constraint is violated:
   - Increase candy count.
4. Continue until no changes occur.
5. Sum all candies.

### Complexity

Time Complexity:

```text
O(n²)
```

Space Complexity:

```text
O(n)
```

### Limitations

- Multiple passes may be required.
- Inefficient for large inputs.
- Poor interview solution.
- Does not scale well.

---

## Optimized Approach

Use a greedy two-pass strategy.

### Algorithm

#### Step 1: Initialize

Give every child one candy.

```text
candies = [1,1,1,...]
```

---

#### Step 2: Left → Right Pass

For every child:

```text
if ratings[i] > ratings[i - 1]
    candies[i] = candies[i - 1] + 1
```

This satisfies all increasing slopes.

---

#### Step 3: Right → Left Pass

For every child:

```text
if ratings[i] > ratings[i + 1]
    candies[i] =
        max(candies[i],
            candies[i + 1] + 1)
```

This satisfies all decreasing slopes without breaking previous assignments.

---

#### Step 4: Sum Results

Add all candy counts.

Return total.

---

### Why It Works

The left pass guarantees:

```text
higher rating on left side
→ more candies
```

The right pass guarantees:

```text
higher rating on right side
→ more candies
```

Taking the maximum preserves both constraints simultaneously.

Since every adjustment is the minimum necessary increase, the resulting allocation is optimal.

---

### Complexity

Time Complexity:

```text
O(n)
```

Space Complexity:

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

---

### Single Element

```text
[5]
```

Output:

```text
1
```

Only one child.

---

### All Equal Ratings

```text
[3,3,3]
```

Output:

```text
3
```

Each child gets one candy.

---

### Strictly Increasing Ratings

```text
[1,2,3,4]
```

Candies:

```text
[1,2,3,4]
```

Output:

```text
10
```

---

### Strictly Decreasing Ratings

```text
[4,3,2,1]
```

Candies:

```text
[4,3,2,1]
```

Output:

```text
10
```

---

### Duplicate Ratings

```text
[1,2,2]
```

Candies:

```text
[1,2,1]
```

Output:

```text
4
```

---

### Negative Values

Ratings may be negative.

```text
[-3,-2,-1]
```

Only relative ordering matters.

Algorithm remains unchanged.

---

### Large Inputs

```text
n = 20,000
```

The O(n) solution scales efficiently.

---

## Dry Run

Input:

```text
ratings = [1,0,2]
```

### Initial State

| Index | Rating | Candies |
|---------|---------|---------|
| 0 | 1 | 1 |
| 1 | 0 | 1 |
| 2 | 2 | 1 |

---

### Left → Right Pass

| Index | Condition | Candies |
|---------|---------|---------|
| 1 | 0 > 1 ? No | [1,1,1] |
| 2 | 2 > 0 ? Yes | [1,1,2] |

---

### Right → Left Pass

| Index | Condition | Candies |
|---------|---------|---------|
| 1 | 0 > 2 ? No | [1,1,2] |
| 0 | 1 > 0 ? Yes | [2,1,2] |

---

### Final Sum

```text
2 + 1 + 2 = 5
```

Answer:

```text
5
```

---

## Common Mistakes

### 1. Using Only One Pass

Many candidates solve only:

```text
Left → Right
```

This fails for decreasing sequences.

Example:

```text
[3,2,1]
```

---

### 2. Forgetting max()

Incorrect:

```text
candies[i] = candies[i + 1] + 1
```

Correct:

```text
candies[i] =
max(candies[i],
    candies[i + 1] + 1)
```

Without `max()`, previous valid assignments may be overwritten.

---

### 3. Treating Equal Ratings Incorrectly

Wrong:

```text
>=
```

Correct:

```text
>
```

Equal ratings have no ordering requirement.

---

### 4. Overcomplicating With Heaps or DP

Greedy is sufficient.

Additional data structures add complexity without benefit.

---

## Interview Discussion

Expected progression:

### Level 1

Recognize neighbor constraints.

---

### Level 2

Propose brute-force simulation.

---

### Level 3

Notice constraints originate from both directions.

---

### Level 4

Develop two-pass greedy solution.

---

### Level 5

Explain correctness and minimality.

Interviewers often ask:

> Why does taking the maximum preserve optimality?

A strong answer discusses satisfying both directional constraints independently and combining them with the smallest valid value.

---

## Follow-up Questions

### Can you solve it using O(1) extra space?

Yes.

There is a slope-counting greedy approach using:

- Increasing runs
- Decreasing runs
- Peak adjustments

Time:

```text
O(n)
```

Space:

```text
O(1)
```

---

### What if candies had different costs?

The optimization strategy changes significantly and may require dynamic programming or graph-based modeling.

---

### What if constraints existed for both neighbors simultaneously?

The problem becomes a generalized resource-allocation constraint problem.

---

## Real World Applications

### Employee Compensation

Higher-performing employees may require higher compensation than nearby peers.

---

### Ranking Systems

Assigning rewards based on relative rankings.

---

### Resource Distribution

Allocating bonuses, credits, or incentives while preserving ordering constraints.

---

### Scheduling Priorities

Assigning priority levels based on comparative performance metrics.

---

## Related Problems

### Easier

- 55. Jump Game
- 121. Best Time to Buy and Sell Stock
- 605. Can Place Flowers

### Similar Greedy Problems

- 45. Jump Game II
- 134. Gas Station
- 406. Queue Reconstruction by Height

### Advanced Greedy Thinking

- 630. Course Schedule III
- 857. Minimum Cost to Hire K Workers
- 1354. Construct Target Array With Multiple Sums