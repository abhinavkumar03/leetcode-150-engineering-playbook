# Product of Array Except Self

## Problem Statement

Given an integer array `nums`, return an array `answer` such that:

`answer[i]` is equal to the product of all elements of `nums` except `nums[i]`.

Requirements:

* The solution must run in **O(n)** time.
* Division operation is **not allowed**.
* The product of any prefix or suffix of `nums` is guaranteed to fit within a 32-bit integer.

### Example 1

```text
Input:
nums = [1,2,3,4]

Output:
[24,12,8,6]
```

Explanation:

```text
answer[0] = 2 × 3 × 4 = 24
answer[1] = 1 × 3 × 4 = 12
answer[2] = 1 × 2 × 4 = 8
answer[3] = 1 × 2 × 3 = 6
```

### Example 2

```text
Input:
nums = [-1,1,0,-3,3]

Output:
[0,0,9,0,0]
```

---

## Difficulty

**Medium**

---

## Tags

* Array
* Prefix Product
* Suffix Product
* Space Optimization

---

## Pattern

### Primary Pattern

Prefix Product

### Secondary Pattern

Suffix Product

### Optimization Pattern

In-Place Prefix/Suffix Accumulation

---

## Intuition

For every position `i`, we need:

```text
Product of all elements left of i
×
Product of all elements right of i
```

Instead of recomputing these products repeatedly, we can precompute:

* Left products (prefix products)
* Right products (suffix products)

Then combine them.

The challenge is achieving this without using division and while maintaining O(1) extra space.

---

## Key Observation

For any index `i`:

```text
answer[i] =
(prefix product before i)
×
(suffix product after i)
```

Example:

```text
nums = [1,2,3,4]

Index 2 (value = 3)

Left Product:
1 × 2 = 2

Right Product:
4

Answer:
2 × 4 = 8
```

Instead of storing separate prefix and suffix arrays, we can:

1. Store prefix products directly inside the result array.
2. Traverse from right to left using a running suffix product.
3. Multiply suffix product into the result.

This achieves O(1) extra space.

---

## Brute Force Approach

### Algorithm

For every index:

1. Traverse the entire array.
2. Multiply all values except the current index.
3. Store the result.

Pseudo-code:

```text
for i in range(n):
    product = 1

    for j in range(n):
        if i != j:
            product *= nums[j]

    answer[i] = product
```

### Complexity

| Metric | Value |
| ------ | ----- |
| Time   | O(n²) |
| Space  | O(1)  |

### Limitations

* Too slow for large inputs.
* Repeatedly recomputes products.
* Does not meet interview expectations.

---

## Optimized Approach

### Algorithm

### Step 1 — Build Prefix Products

Store product of all elements to the left.

```text
nums = [1,2,3,4]

answer:

[1,1,2,6]
```

Explanation:

```text
answer[0] = 1
answer[1] = 1
answer[2] = 1×2
answer[3] = 1×2×3
```

---

### Step 2 — Traverse From Right

Maintain:

```text
suffixProduct = 1
```

Move from right to left:

```text
answer[i] *= suffixProduct
suffixProduct *= nums[i]
```

---

### Step 3 — Final Result

```text
answer = [24,12,8,6]
```

---

### Why It Works

At every index:

```text
answer[i]
=
(left product)
×
(right product)
```

The first pass stores:

```text
left product
```

The second pass injects:

```text
right product
```

Combining them produces the desired result.

Because each element is processed twice:

```text
Time = O(n)
```

And only one extra variable is used:

```text
Space = O(1)
```

(excluding output array)

---

### Complexity

| Metric | Value            |
| ------ | ---------------- |
| Time   | O(n)             |
| Space  | O(1) Extra Space |

---

## Edge Cases

### Empty Input

```text
[]
```

Output:

```text
[]
```

---

### Single Element

```text
[5]
```

Output:

```text
[1]
```

Reason:

No other elements exist.

---

### Duplicates

```text
[2,2,2]
```

Output:

```text
[4,4,4]
```

---

### Negative Values

```text
[-1,2,-3]
```

Output:

```text
[-6,3,-2]
```

---

### Contains One Zero

```text
[1,2,0,4]
```

Output:

```text
[0,0,8,0]
```

Only the zero position receives the product of non-zero elements.

---

### Contains Multiple Zeros

```text
[1,0,2,0]
```

Output:

```text
[0,0,0,0]
```

Every result becomes zero.

---

### Large Inputs

```text
100000 elements
```

Optimized solution remains:

```text
O(n)
```

and scales efficiently.

---

## Dry Run

### Input

```text
nums = [1,2,3,4]
```

### Prefix Pass

| Index | Value | Prefix Product Stored |
| ----- | ----- | --------------------- |
| 0     | 1     | 1                     |
| 1     | 2     | 1                     |
| 2     | 3     | 2                     |
| 3     | 4     | 6                     |

Result:

```text
answer = [1,1,2,6]
```

---

### Suffix Pass

Initial:

```text
suffix = 1
```

| Index | answer[i] Before | suffix | answer[i] After |
| ----- | ---------------- | ------ | --------------- |
| 3     | 6                | 1      | 6               |
| 2     | 2                | 4      | 8               |
| 1     | 1                | 12     | 12              |
| 0     | 1                | 24     | 24              |

Final:

```text
[24,12,8,6]
```

---

## Common Mistakes

### Using Division

```text
totalProduct / nums[i]
```

Rejected by problem constraints.

---

### Incorrect Prefix Initialization

Wrong:

```text
answer[0] = nums[0]
```

Correct:

```text
answer[0] = 1
```

---

### Forgetting to Update Suffix

Wrong order:

```text
suffix *= nums[i]
answer[i] *= suffix
```

Correct:

```text
answer[i] *= suffix
suffix *= nums[i]
```

---

### Mishandling Zeros

Division-based approaches usually fail when zeros exist.

Prefix/Suffix method naturally handles zeros.

---

## Interview Discussion

### Expected Progression

An interviewer typically expects:

```text
Brute Force
↓
Division Solution
↓
Prefix Array + Suffix Array
↓
O(1) Extra Space Optimization
```

Discuss:

* Why division is prohibited.
* Why prefix/suffix decomposition works.
* How space optimization is achieved.

---

## Follow-up Questions

### Follow-up 1

Can you solve it using division?

**Yes**, but it violates constraints.

---

### Follow-up 2

Can you solve it in O(1) extra space?

**Yes**, using the output array plus a suffix accumulator.

---

### Follow-up 3

How does the algorithm handle zeros?

The prefix/suffix multiplication naturally produces correct values without special logic.

---

### Follow-up 4

Can this be done in one pass?

Not cleanly, because both left and right products are required.

---

## Real World Applications

### Financial Analytics

Compute portfolio metrics excluding the current asset.

---

### Distributed Systems

Calculate aggregate system load excluding one server.

---

### Recommendation Engines

Compute influence scores excluding a specific user or item.

---

### Data Processing Pipelines

Generate aggregate statistics while excluding the current record.

---

## Related Problems

* 238. Product of Array Except Self
* 42. Trapping Rain Water
* 53. Maximum Subarray
* 152. Maximum Product Subarray
* 560. Subarray Sum Equals K
* 724. Find Pivot Index

---

**PHASE 2 Complete.**

Reply with **Next** to continue to **PHASE 3 — Code Implementations**.
