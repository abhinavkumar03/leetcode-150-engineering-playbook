# Dry Run — Jump Game

## Goal

Determine whether it is possible to reach the last index of the array.

We maintain:

```text
farthestReach
```

which represents:

> The farthest index reachable so far.

---

# Example 1

## Input

```text
nums = [2,3,1,1,4]
```

## Expected Output

```text
true
```

---

## Initial State

| Variable | Value |
|-----------|-----------|
| farthestReach | 0 |
| lastIndex | 4 |

---

## Visual Walkthrough

```text
Index:  0  1  2  3  4
Value: [2, 3, 1, 1, 4]

Start at index 0

Reachable Range:
[0]
```

---

## Iteration 1

### Current Position

```text
index = 0
nums[0] = 2
```

### Reachability Check

```text
0 <= farthestReach (0)
```

Reachable.

### Update Reach

```text
farthestReach =
max(0, 0 + 2)

= 2
```

### State

| Index | Value |
|---------|---------|
| Current Index | 0 |
| Jump Length | 2 |
| Farthest Reach Before | 0 |
| Farthest Reach After | 2 |

### Reachable Region

```text
[0 ----- 2]

Index: 0 1 2 3 4
       ^^^^^
```

---

## Iteration 2

### Current Position

```text
index = 1
nums[1] = 3
```

### Reachability Check

```text
1 <= 2
```

Reachable.

### Update Reach

```text
farthestReach =
max(2, 1 + 3)

= 4
```

### State

| Index | Value |
|---------|---------|
| Current Index | 1 |
| Jump Length | 3 |
| Farthest Reach Before | 2 |
| Farthest Reach After | 4 |

### Reachable Region

```text
[0 ----------- 4]

Index: 0 1 2 3 4
       ^^^^^^^^^
```

Now:

```text
farthestReach = 4
```

and

```text
lastIndex = 4
```

Therefore:

```text
destination reached
```

Return:

```text
true
```

---

# Complete Iteration Table

| Index | nums[i] | Reachable? | Farthest Reach Before | Farthest Reach After |
|---------|---------|---------|---------|---------|
| 0 | 2 | Yes | 0 | 2 |
| 1 | 3 | Yes | 2 | 4 |

Algorithm exits early.

---

# Example 2

## Input

```text
nums = [3,2,1,0,4]
```

## Expected Output

```text
false
```

---

## Initial State

| Variable | Value |
|-----------|-----------|
| farthestReach | 0 |
| lastIndex | 4 |

---

## Visual Walkthrough

```text
Index:  0  1  2  3  4
Value: [3, 2, 1, 0, 4]
```

---

## Iteration 1

### Current Position

```text
index = 0
nums[0] = 3
```

### Update

```text
farthestReach =
max(0, 0 + 3)

= 3
```

Reachable region:

```text
0 → 3
```

---

## Iteration 2

### Current Position

```text
index = 1
nums[1] = 2
```

### Update

```text
farthestReach =
max(3, 1 + 2)

= 3
```

No improvement.

---

## Iteration 3

### Current Position

```text
index = 2
nums[2] = 1
```

### Update

```text
farthestReach =
max(3, 2 + 1)

= 3
```

No improvement.

---

## Iteration 4

### Current Position

```text
index = 3
nums[3] = 0
```

### Update

```text
farthestReach =
max(3, 3 + 0)

= 3
```

Still cannot move further.

---

## State Before Index 4

Current reach:

```text
farthestReach = 3
```

Next index:

```text
index = 4
```

Check:

```text
4 > 3
```

This means:

```text
index 4 is unreachable
```

Return:

```text
false
```

---

# Failure Visualization

```text
Index:  0  1  2  3  4
Value: [3, 2, 1, 0, 4]

Reachable:
^^^^

Blocked Here:
            X
```

More accurately:

```text
Index:  0  1  2  3  4
Value: [3, 2, 1, 0, 4]

Reach Limit = 3

0 -----> 3

Index 4 cannot be reached.
```

---

# State Transition Summary

## Success Case

```text
Start
  |
Reach = 0
  |
Process Index 0
  |
Reach = 2
  |
Process Index 1
  |
Reach = 4
  |
Reach >= Last Index
  |
SUCCESS
```

---

## Failure Case

```text
Start
  |
Reach = 0
  |
Reach = 3
  |
Process Index 1
  |
Reach = 3
  |
Process Index 2
  |
Reach = 3
  |
Process Index 3
  |
Reach = 3
  |
Need Index 4
  |
4 > 3
  |
FAILURE
```

---

# Key Insight

The algorithm never decides:

```text
Which jump should I take?
```

Instead, it continuously asks:

```text
What is the farthest index
that can currently be reached?
```

This single observation reduces the problem from exploring many possible paths to maintaining one value:

```text
farthestReach
```

which enables:

```text
Time Complexity  : O(n)
Space Complexity : O(1)
```