# 134. Gas Station

## Problem Statement

There are `n` gas stations along a circular route.

You are given two integer arrays:

- `gas[i]` = amount of gas available at station `i`
- `cost[i]` = gas required to travel from station `i` to station `i + 1`

You begin with an empty tank at one of the stations.

Return the starting gas station's index if you can travel around the circuit exactly once in the clockwise direction.

If it is impossible, return `-1`.

It is guaranteed that if a solution exists, it is unique.

---

## Difficulty

**Medium**

---

## Tags

- Array
- Greedy

---

## Pattern

### Primary Pattern
Greedy

### Secondary Pattern
Circular Array Traversal

---

## Intuition

A brute-force approach would try every station as a starting point and simulate the entire journey.

However, this leads to repeated work.

The key insight is:

If we start from station `start` and fail at station `i`, then none of the stations between `start` and `i` can be valid starting stations.

Why?

Because those stations would begin with even less fuel than the current failed attempt.

This allows us to skip many unnecessary checks and build a linear-time greedy solution.

---

## Key Observation

Let:

```text
balance = gas[i] - cost[i]
```

Two important facts:

### Observation 1

If:

```text
sum(gas) < sum(cost)
```

then completing the circuit is impossible.

Return:

```text
-1
```

immediately.

---

### Observation 2

If the current fuel becomes negative:

```text
tank < 0
```

at station `i`, then:

```text
start = i + 1
```

because every station before `i + 1` has already been proven invalid.

---

## Brute Force Approach

Try every station as a starting point.

For each station:

1. Start with empty tank
2. Simulate full circular traversal
3. Check whether tank ever becomes negative
4. Return first valid station

### Algorithm

For each station:

```text
for start in [0..n-1]:
    tank = 0

    traverse n stations:

        tank += gas[current]
        tank -= cost[current]

        if tank < 0:
            fail
```

Return valid start if found.

Otherwise:

```text
-1
```

---

### Complexity

#### Time Complexity

```text
O(n²)
```

For every station we may traverse nearly the entire circuit.

#### Space Complexity

```text
O(1)
```

---

### Limitations

- Too slow for large inputs
- Repeats many failed traversals
- Does not leverage greedy observations

---

## Optimized Approach

Use a greedy scan while maintaining:

- totalBalance
- currentTank
- candidateStart

### Algorithm

Initialize:

```text
totalBalance = 0
currentTank = 0
start = 0
```

For every station:

```text
diff = gas[i] - cost[i]

totalBalance += diff
currentTank += diff
```

If:

```text
currentTank < 0
```

then:

```text
start = i + 1
currentTank = 0
```

After processing all stations:

If:

```text
totalBalance < 0
```

return:

```text
-1
```

Otherwise return:

```text
start
```

---

### Why It Works

Suppose we started from station `start`.

If we fail at station `i`:

```text
currentTank < 0
```

Then every station between:

```text
[start, i]
```

must also fail.

Reason:

Any intermediate station would start with less accumulated fuel.

Therefore we can safely skip all of them and move directly to:

```text
i + 1
```

This greedy elimination guarantees a single linear scan.

---

### Complexity

#### Time Complexity

```text
O(n)
```

Single traversal of the array.

#### Space Complexity

```text
O(1)
```

Only a few variables are maintained.

---

## Edge Cases

### Empty Input

```text
gas = []
cost = []
```

Return:

```text
-1
```

or handle according to platform constraints.

---

### Single Station (Valid)

```text
gas  = [5]
cost = [4]
```

Answer:

```text
0
```

---

### Single Station (Invalid)

```text
gas  = [3]
cost = [5]
```

Answer:

```text
-1
```

---

### Exact Fuel Match

```text
gas  = [2,2,2]
cost = [2,2,2]
```

Answer:

```text
0
```

---

### Large Deficit Early

```text
gas  = [1,10,1]
cost = [5,1,1]
```

Need to restart after failure.

---

### Duplicates

```text
gas  = [4,4,4]
cost = [3,3,3]
```

Should still work correctly.

---

### Large Inputs

```text
n = 100000
```

Greedy solution remains efficient.

---

## Dry Run

### Example

```text
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
```

### Initial State

```text
start = 0
tank = 0
total = 0
```

| i | gas | cost | diff | tank | total | start |
|---|------|------|------|------|------|------|
| 0 | 1 | 3 | -2 | -2 | -2 | 0 |
|   |   |   |   | reset |   | 1 |
| 1 | 2 | 4 | -2 | -2 | -4 | 1 |
|   |   |   |   | reset |   | 2 |
| 2 | 3 | 5 | -2 | -2 | -6 | 2 |
|   |   |   |   | reset |   | 3 |
| 3 | 4 | 1 | +3 | 3 | -3 | 3 |
| 4 | 5 | 2 | +3 | 6 | 0 | 3 |

Final:

```text
total = 0
```

Since:

```text
total >= 0
```

Answer:

```text
3
```

---

## Common Mistakes

### Forgetting Global Feasibility Check

Wrong:

```text
return start
```

without checking:

```text
totalBalance
```

This may return an invalid answer.

---

### Resetting Start Incorrectly

Wrong:

```text
start = i
```

Correct:

```text
start = i + 1
```

---

### Not Resetting Current Tank

After failure:

```text
currentTank = 0
```

must be performed.

---

### Using Nested Simulation

Many candidates stop at:

```text
O(n²)
```

and miss the greedy optimization.

---

## Interview Discussion

### Expected Progression

1. Explain brute force
2. Analyze repeated work
3. Identify failure elimination property
4. Derive greedy solution
5. Prove correctness
6. Discuss complexity

---

### What Interviewers Like To Hear

- "If station A fails before reaching B, every station between them also fails."
- "We can eliminate an entire range of candidates."
- "Total gas must be at least total cost."

These statements demonstrate true understanding.

---

## Follow-up Questions

### 1. What if multiple answers were allowed?

The algorithm still finds one valid answer.

---

### 2. What if the route was not circular?

Problem becomes significantly different.

---

### 3. Can we solve it using prefix sums?

Yes, but greedy is simpler and optimal.

---

### 4. What if gas values change dynamically?

A different data structure would be required.

---

## Real World Applications

### Logistics Route Planning

Determining whether a vehicle can complete a delivery route.

---

### Electric Vehicle Charging

Finding a charging station from which an EV can complete a circuit.

---

### Resource Allocation Systems

Checking whether cumulative resources can sustain cyclic workloads.

---

### Network Token Circulation

Verifying sufficient resources to complete a ring traversal.

---

## Related Problems

### Easy

- 121. Best Time to Buy and Sell Stock
- 53. Maximum Subarray

### Medium

- 55. Jump Game
- 45. Jump Game II
- 452. Minimum Number of Arrows to Burst Balloons

### Advanced Greedy

- 135. Candy
- 406. Queue Reconstruction by Height
- 763. Partition Labels

---

## Key Takeaway

The core insight is not the fuel calculation.

The real lesson is the greedy elimination rule:

> If a starting station fails, every station before the failure point can be discarded.

Recognizing and proving this property transforms an O(n²) simulation into an O(n) optimal solution.