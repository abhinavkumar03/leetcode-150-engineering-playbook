# Jump Game II — Dry Run

## Goal

Find the minimum number of jumps needed to reach the last index.

---

# Example 1

## Input

```text
nums = [2,3,1,1,4]
```

## Expected Output

```text
2
```

---

# Greedy Visualization

Each value represents the maximum jump length.

```text
Index:  0  1  2  3  4
Value: [2, 3, 1, 1, 4]
```

From index 0:

```text
Can reach:
1
2
```

From index 1:

```text
Can reach:
2
3
4
```

Optimal path:

```text
0 → 1 → 4
```

Total jumps:

```text
2
```

---

# Variable Definitions

| Variable | Meaning |
|-----------|-----------|
| jumps | Total jumps used |
| currentEnd | End of current jump range |
| farthestReach | Farthest position reachable from current range |

---

# Initialization

```text
jumps = 0
currentEnd = 0
farthestReach = 0
```

---

# Initial State

```text
Index:  0  1  2  3  4
Value: [2, 3, 1, 1, 4]

^
i
```

---

# Iteration 1

## Current Index

```text
i = 0
nums[i] = 2
```

## Calculate Reach

```text
i + nums[i]
= 0 + 2
= 2
```

Update:

```text
farthestReach = max(0, 2)
               = 2
```

Current state:

| Variable | Value |
|-----------|---------|
| jumps | 0 |
| currentEnd | 0 |
| farthestReach | 2 |

---

## Boundary Check

```text
i == currentEnd

0 == 0
```

TRUE

Current range has been fully explored.

Take a jump.

```text
jumps++
```

```text
jumps = 1
```

Move boundary:

```text
currentEnd = farthestReach
           = 2
```

State becomes:

| Variable | Value |
|-----------|---------|
| jumps | 1 |
| currentEnd | 2 |
| farthestReach | 2 |

---

# Visual Range Expansion

Current jump reaches:

```text
[0] ---> [1,2]
```

```text
Index:  0  1  2  3  4
         └─────┘
      current range
```

---

# Iteration 2

## Current Index

```text
i = 1
nums[i] = 3
```

Reach:

```text
1 + 3 = 4
```

Update:

```text
farthestReach = max(2,4)
               = 4
```

State:

| Variable | Value |
|-----------|---------|
| jumps | 1 |
| currentEnd | 2 |
| farthestReach | 4 |

---

# Visual

```text
Index:  0  1  2  3  4
               ↑
     From index 1
     can reach 4
```

Destination becomes reachable.

---

# Iteration 3

## Current Index

```text
i = 2
nums[i] = 1
```

Reach:

```text
2 + 1 = 3
```

Update:

```text
farthestReach = max(4,3)
               = 4
```

No change.

State:

| Variable | Value |
|-----------|---------|
| jumps | 1 |
| currentEnd | 2 |
| farthestReach | 4 |

---

# Boundary Check

```text
i == currentEnd

2 == 2
```

TRUE

Current range finished.

Take another jump.

```text
jumps++
```

```text
jumps = 2
```

Move boundary:

```text
currentEnd = farthestReach
           = 4
```

State:

| Variable | Value |
|-----------|---------|
| jumps | 2 |
| currentEnd | 4 |
| farthestReach | 4 |

---

# Visual Range Expansion

Previous range:

```text
[1,2]
```

Generated next range:

```text
[3,4]
```

Visualization:

```text
Index:  0  1  2  3  4

Level 0:
[0]

Level 1:
[1,2]

Level 2:
[3,4]
```

Destination is now inside current range.

---

# Final Answer

```text
jumps = 2
```

Return:

```text
2
```

---

# Complete State Transition Table

| i | nums[i] | Reach (i + nums[i]) | farthestReach | currentEnd | jumps |
|---|---|---|---|---|---|
| Start | - | - | 0 | 0 | 0 |
| 0 | 2 | 2 | 2 | 0 | 0 |
| Jump | - | - | 2 | 2 | 1 |
| 1 | 3 | 4 | 4 | 2 | 1 |
| 2 | 1 | 3 | 4 | 2 | 1 |
| Jump | - | - | 4 | 4 | 2 |

---

# BFS Interpretation

The greedy solution behaves exactly like BFS.

## Level 0

```text
Index 0
```

Reachable:

```text
[0]
```

---

## Level 1

Using one jump:

```text
[1,2]
```

---

## Level 2

Using two jumps:

```text
[3,4]
```

Destination found.

Answer:

```text
2 jumps
```

---

# Example 2

## Input

```text
nums = [2,3,0,1,4]
```

---

## Iteration Summary

| i | farthestReach | currentEnd | jumps |
|---|---|---|---|
| 0 | 2 | 0 | 0 |
| Jump | - | 2 | 1 |
| 1 | 4 | 2 | 1 |
| 2 | 4 | 2 | 1 |
| Jump | - | 4 | 2 |

Return:

```text
2
```

---

# Key Insight

The algorithm never decides:

```text
"Which exact index should I jump to?"
```

Instead it decides:

```text
"What is the farthest position reachable from the current jump range?"
```

That single observation reduces:

```text
Exponential Search
        ↓
Linear Time Greedy Solution
```

---

# Dry Run Takeaway

Think of the array as BFS levels:

```text
Current Range
      ↓
Explore Everything
      ↓
Find Farthest Reach
      ↓
Commit One Jump
      ↓
Expand Next Range
```

This guarantees the minimum number of jumps while maintaining:

```text
Time  : O(n)
Space : O(1)
```