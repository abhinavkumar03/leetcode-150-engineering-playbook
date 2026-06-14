# Product of Array Except Self — Dry Run

## Goal

For every index `i`, compute:

```text
answer[i] =
(product of all elements before i)
×
(product of all elements after i)
```

without using division.

---

# Example Input

```text
nums = [1, 2, 3, 4]
```

Expected Output:

```text
[24, 12, 8, 6]
```

---

# High-Level Idea

The optimized solution works in two phases:

### Phase 1: Prefix Products

Store the product of all elements to the left of each index.

### Phase 2: Suffix Products

Traverse from right to left and multiply the suffix product into the result.

---

# Initial State

```text
nums   = [1, 2, 3, 4]

result = [0, 0, 0, 0]
```

Initialize:

```text
result[0] = 1
```

Current state:

| Index  | 0 | 1 | 2 | 3 |
| ------ | - | - | - | - |
| result | 1 | 0 | 0 | 0 |

---

# Phase 1 — Build Prefix Products

## Iteration 1

### i = 1

Formula:

```text
result[i] = result[i-1] × nums[i-1]
```

Calculation:

```text
result[1]
=
result[0] × nums[0]

=
1 × 1

=
1
```

State:

| Index  | 0 | 1 | 2 | 3 |
| ------ | - | - | - | - |
| result | 1 | 1 | 0 | 0 |

---

## Iteration 2

### i = 2

Calculation:

```text
result[2]
=
result[1] × nums[1]

=
1 × 2

=
2
```

State:

| Index  | 0 | 1 | 2 | 3 |
| ------ | - | - | - | - |
| result | 1 | 1 | 2 | 0 |

---

## Iteration 3

### i = 3

Calculation:

```text
result[3]
=
result[2] × nums[2]

=
2 × 3

=
6
```

State:

| Index  | 0 | 1 | 2 | 3 |
| ------ | - | - | - | - |
| result | 1 | 1 | 2 | 6 |

---

# Prefix Phase Complete

Result array now stores:

```text
Product of all elements to the LEFT
```

| Index | Value Stored |
| ----- | ------------ |
| 0     | 1            |
| 1     | 1            |
| 2     | 2            |
| 3     | 6            |

Current:

```text
result = [1, 1, 2, 6]
```

---

# Phase 2 — Inject Suffix Products

Initialize:

```text
suffixProduct = 1
```

Current state:

```text
result = [1, 1, 2, 6]
suffix = 1
```

---

## Iteration 1

### i = 3

Current:

```text
result[3] = 6
suffix = 1
```

Update:

```text
result[3]
=
6 × 1

=
6
```

Update suffix:

```text
suffix
=
1 × nums[3]

=
1 × 4

=
4
```

State:

```text
result = [1, 1, 2, 6]
suffix = 4
```

---

## Iteration 2

### i = 2

Current:

```text
result[2] = 2
suffix = 4
```

Update result:

```text
result[2]
=
2 × 4

=
8
```

Update suffix:

```text
suffix
=
4 × nums[2]

=
4 × 3

=
12
```

State:

```text
result = [1, 1, 8, 6]
suffix = 12
```

---

## Iteration 3

### i = 1

Current:

```text
result[1] = 1
suffix = 12
```

Update result:

```text
result[1]
=
1 × 12

=
12
```

Update suffix:

```text
suffix
=
12 × nums[1]

=
12 × 2

=
24
```

State:

```text
result = [1, 12, 8, 6]
suffix = 24
```

---

## Iteration 4

### i = 0

Current:

```text
result[0] = 1
suffix = 24
```

Update result:

```text
result[0]
=
1 × 24

=
24
```

Update suffix:

```text
suffix
=
24 × nums[0]

=
24 × 1

=
24
```

Final State:

```text
result = [24, 12, 8, 6]
```

---

# Complete Execution Table

## Prefix Pass

| i | nums[i] | result after update |
| - | ------- | ------------------- |
| 0 | 1       | [1,0,0,0]           |
| 1 | 2       | [1,1,0,0]           |
| 2 | 3       | [1,1,2,0]           |
| 3 | 4       | [1,1,2,6]           |

---

## Suffix Pass

| i | suffix before | result[i] before | result[i] after | suffix after |
| - | ------------- | ---------------- | --------------- | ------------ |
| 3 | 1             | 6                | 6               | 4            |
| 2 | 4             | 2                | 8               | 12           |
| 1 | 12            | 1                | 12              | 24           |
| 0 | 24            | 1                | 24              | 24           |

---

# Visual Representation

```text
nums:

[1, 2, 3, 4]

---------------------------------

Prefix Products:

Index 0 -> 1
Index 1 -> 1
Index 2 -> 1×2 = 2
Index 3 -> 1×2×3 = 6

result:

[1, 1, 2, 6]

---------------------------------

Right-to-Left Suffix Multiplication:

suffix=1

Index 3:
6 × 1 = 6

suffix=4

Index 2:
2 × 4 = 8

suffix=12

Index 1:
1 × 12 = 12

suffix=24

Index 0:
1 × 24 = 24

---------------------------------

Final:

[24, 12, 8, 6]
```

---

# Why This Works

For every position:

```text
answer[i]
=
(prefix product before i)
×
(suffix product after i)
```

The first traversal stores:

```text
left product
```

The second traversal contributes:

```text
right product
```

Combining both gives the required answer without division.

---

# Complexity Summary

| Metric       | Complexity       |
| ------------ | ---------------- |
| Time         | O(n)             |
| Space        | O(1) Extra Space |
| Output Array | O(n)             |

---

# Key Interview Takeaway

The core insight is:

```text
Product Except Self
=
Left Product
×
Right Product
```

Use:

* One forward pass for prefix products
* One backward pass for suffix products
* Reuse the output array

This achieves the optimal:

```text
Time  : O(n)
Space : O(1)
```

**PHASE 4 Complete.**

Reply with **Next** to continue to **PHASE 5 — interview-notes.md**.
