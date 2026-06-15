# Dry Run - 134. Gas Station

## Purpose

This document provides a detailed walkthrough of the Greedy solution used for the Gas Station problem.

The goal is to understand:

- Why the algorithm works
- How candidate starting stations are eliminated
- How the running fuel balance changes
- Why only one pass is required

---

# Example 1

## Input

```text
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
```

## Expected Output

```text
3
```

---

# Step 1: Calculate Fuel Difference

For each station:

```text
difference = gas[i] - cost[i]
```

| Station | Gas | Cost | Difference |
|----------|-----|------|------------|
| 0 | 1 | 3 | -2 |
| 1 | 2 | 4 | -2 |
| 2 | 3 | 5 | -2 |
| 3 | 4 | 1 | +3 |
| 4 | 5 | 2 | +3 |

---

# Initial State

```text
startStation = 0
currentTank  = 0
totalBalance = 0
```

---

# Iteration Walkthrough

## Station 0

### Fuel Difference

```text
1 - 3 = -2
```

### Update Balances

```text
currentTank  = -2
totalBalance = -2
```

### Check Tank

```text
currentTank < 0
```

Cannot reach next station.

Therefore:

```text
startStation = 1
currentTank  = 0
```

### Why?

If station 0 cannot even reach station 1,
then station 0 cannot be the answer.

---

### State After Station 0

| Variable | Value |
|-----------|---------|
| startStation | 1 |
| currentTank | 0 |
| totalBalance | -2 |

---

## Station 1

### Fuel Difference

```text
2 - 4 = -2
```

### Update

```text
currentTank  = -2
totalBalance = -4
```

Again:

```text
currentTank < 0
```

Reset:

```text
startStation = 2
currentTank = 0
```

---

### State After Station 1

| Variable | Value |
|-----------|---------|
| startStation | 2 |
| currentTank | 0 |
| totalBalance | -4 |

---

## Station 2

### Fuel Difference

```text
3 - 5 = -2
```

### Update

```text
currentTank  = -2
totalBalance = -6
```

Failure again.

Reset:

```text
startStation = 3
currentTank = 0
```

---

### State After Station 2

| Variable | Value |
|-----------|---------|
| startStation | 3 |
| currentTank | 0 |
| totalBalance | -6 |

---

## Station 3

### Fuel Difference

```text
4 - 1 = +3
```

### Update

```text
currentTank  = 3
totalBalance = -3
```

Tank remains positive.

Continue.

---

### State After Station 3

| Variable | Value |
|-----------|---------|
| startStation | 3 |
| currentTank | 3 |
| totalBalance | -3 |

---

## Station 4

### Fuel Difference

```text
5 - 2 = +3
```

### Update

```text
currentTank  = 6
totalBalance = 0
```

Still positive.

Continue.

---

### State After Station 4

| Variable | Value |
|-----------|---------|
| startStation | 3 |
| currentTank | 6 |
| totalBalance | 0 |

---

# Final Check

```text
totalBalance = 0
```

Since:

```text
totalBalance >= 0
```

A valid route exists.

Return:

```text
startStation = 3
```

---

# Visual Route Simulation

Starting at station 3:

```text
Start at 3
Tank = 0
```

---

## Visit Station 3

```text
Tank += 4
Tank -= 1

Tank = 3
```

Move to station 4.

---

## Visit Station 4

```text
Tank += 5
Tank -= 2

Tank = 6
```

Move to station 0.

---

## Visit Station 0

```text
Tank += 1
Tank -= 3

Tank = 4
```

Move to station 1.

---

## Visit Station 1

```text
Tank += 2
Tank -= 4

Tank = 2
```

Move to station 2.

---

## Visit Station 2

```text
Tank += 3
Tank -= 5

Tank = 0
```

Returned successfully.

---

# Circular Journey Diagram

```text
                +--------+
                |   0    |
                +--------+
                  gas=1
                 cost=3
                     ↑
                     |
                     |
+--------+       +--------+
|   4    | ----> |   1    |
+--------+       +--------+
 gas=5            gas=2
 cost=2           cost=4
     ↑                |
     |                |
     |                ↓
+--------+ <---- +--------+
|   3    |       |   2    |
+--------+       +--------+
 gas=4            gas=3
 cost=1           cost=5

Valid Start = 3
```

---

# Candidate Elimination Visualization

Initially:

```text
Possible Starts:
[0,1,2,3,4]
```

---

After failing at station 0:

```text
Remove:
[0]
```

Remaining:

```text
[1,2,3,4]
```

---

After failing at station 1:

```text
Remove:
[1]
```

Remaining:

```text
[2,3,4]
```

---

After failing at station 2:

```text
Remove:
[2]
```

Remaining:

```text
[3,4]
```

---

Greedy proof tells us:

```text
If start fails at i,
every station between start and i
also fails.
```

Thus we safely skip entire ranges.

---

# Example 2

## Input

```text
gas  = [2,3,4]
cost = [3,4,3]
```

---

## Difference Array

```text
[-1,-1,+1]
```

---

## Total Balance

```text
-1 + -1 + 1
=
-1
```

Since:

```text
totalBalance < 0
```

Answer:

```text
-1
```

No route can complete the circuit.

---

# Key Learning

The critical insight is:

```text
Local Failure
        ↓
Eliminate Entire Range
        ↓
Single Pass Solution
        ↓
O(n) Time
```

This is a classic Greedy optimization where failed states allow us to discard many future candidates without explicitly testing them.