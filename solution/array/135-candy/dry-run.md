# Candy — Dry Run

## Goal

Distribute the minimum number of candies such that:

1. Every child receives at least one candy.
2. A child with a higher rating than an adjacent child receives more candies.

---

# Example Input

```text
ratings = [1, 0, 2]
```

Expected Output:

```text
5
```

---

# High-Level Strategy

We solve the problem in three phases:

1. Initialize every child with one candy.
2. Traverse Left → Right to satisfy increasing rating relationships.
3. Traverse Right → Left to satisfy decreasing rating relationships.
4. Sum all candies.

---

# Initial State

Each child receives one candy.

| Index | Rating | Candies |
|---------|---------|---------|
| 0 | 1 | 1 |
| 1 | 0 | 1 |
| 2 | 2 | 1 |

Current candy array:

```text
[1, 1, 1]
```

---

# Pass 1: Left → Right

Rule:

```text
If ratings[i] > ratings[i - 1]
then candies[i] = candies[i - 1] + 1
```

---

## Iteration 1

Index:

```text
i = 1
```

Compare:

```text
ratings[1] = 0
ratings[0] = 1
```

Condition:

```text
0 > 1 ?
```

Result:

```text
False
```

No update required.

Candy array:

```text
[1, 1, 1]
```

---

## Iteration 2

Index:

```text
i = 2
```

Compare:

```text
ratings[2] = 2
ratings[1] = 0
```

Condition:

```text
2 > 0 ?
```

Result:

```text
True
```

Update:

```text
candies[2] = candies[1] + 1
           = 1 + 1
           = 2
```

Candy array:

```text
[1, 1, 2]
```

---

# State After Left Pass

| Index | Rating | Candies |
|---------|---------|---------|
| 0 | 1 | 1 |
| 1 | 0 | 1 |
| 2 | 2 | 2 |

Current candies:

```text
[1, 1, 2]
```

The increasing relationship has been satisfied.

However, we have not yet checked decreasing relationships.

---

# Pass 2: Right → Left

Rule:

```text
If ratings[i] > ratings[i + 1]
then candies[i] =
max(
    candies[i],
    candies[i + 1] + 1
)
```

---

## Iteration 1

Index:

```text
i = 1
```

Compare:

```text
ratings[1] = 0
ratings[2] = 2
```

Condition:

```text
0 > 2 ?
```

Result:

```text
False
```

No update required.

Candy array:

```text
[1, 1, 2]
```

---

## Iteration 2

Index:

```text
i = 0
```

Compare:

```text
ratings[0] = 1
ratings[1] = 0
```

Condition:

```text
1 > 0 ?
```

Result:

```text
True
```

Required candies:

```text
candies[1] + 1
= 1 + 1
= 2
```

Update:

```text
candies[0] =
max(1, 2)
= 2
```

Candy array:

```text
[2, 1, 2]
```

---

# Final State

| Index | Rating | Candies |
|---------|---------|---------|
| 0 | 1 | 2 |
| 1 | 0 | 1 |
| 2 | 2 | 2 |

Final candy distribution:

```text
[2, 1, 2]
```

---

# Total Candies

```text
2 + 1 + 2 = 5
```

Answer:

```text
5
```

---

# Visual Walkthrough

Ratings:

```text
Index:    0   1   2
Rating:   1   0   2
```

Initial candies:

```text
1   1   1
```

After Left → Right:

```text
1   1   2
```

After Right → Left:

```text
2   1   2
```

Final answer:

```text
5
```

---

# State Transition Summary

| Phase | Candy State |
|---------|---------|
| Initial | [1,1,1] |
| After Left Pass | [1,1,2] |
| After Right Pass | [2,1,2] |
| Final | [2,1,2] |

---

# Larger Example

Input:

```text
ratings = [1, 2, 2]
```

---

## Initialization

```text
[1,1,1]
```

---

## Left Pass

### i = 1

```text
2 > 1
```

Update:

```text
[1,2,1]
```

### i = 2

```text
2 > 2
```

False

Remain:

```text
[1,2,1]
```

---

## Right Pass

### i = 1

```text
2 > 2
```

False

### i = 0

```text
1 > 2
```

False

Remain:

```text
[1,2,1]
```

---

## Result

```text
1 + 2 + 1 = 4
```

Answer:

```text
4
```

---

# Why max() Is Necessary

Consider:

```text
ratings = [1, 3, 4, 5, 2]
```

After Left Pass:

```text
[1,2,3,4,1]
```

When processing right-to-left:

```text
5 > 2
```

We need:

```text
candies[3] >= 2
```

But candies[3] is already:

```text
4
```

Using:

```text
max(existing, required)
```

prevents overwriting a valid larger value.

Without `max()`, earlier constraints could be broken.

---

# Key Learning

This problem demonstrates a classic greedy pattern:

1. Solve constraints from one direction.
2. Solve constraints from the opposite direction.
3. Merge both requirements.
4. Use the minimum valid assignment.

This bidirectional greedy strategy appears in many advanced interview problems involving neighboring dependencies and local ordering constraints.