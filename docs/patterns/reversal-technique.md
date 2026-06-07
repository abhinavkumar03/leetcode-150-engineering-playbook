# Reversal Technique Pattern

## Pattern Definition

The Reversal Technique is an in-place array/string transformation pattern where one or more segments are reversed to achieve a desired rearrangement without using additional memory.

Instead of physically moving elements one by one, we:

1. Reverse a larger structure.
2. Reverse one or more substructures.
3. Restore internal ordering while preserving new positions.

This pattern frequently appears in:

- Array rotation
- String rotation
- Matrix transformations
- In-place rearrangement problems

---

## When To Use This Pattern

Consider the Reversal Technique when:

### Signal 1 — Rotation Problems

Problem statements contain:

```text
Rotate
Shift
Circular movement
Move elements to front/back
```

Examples:

- Rotate Array
- Rotate String
- Shift Grid

---

### Signal 2 — O(1) Space Requirement

Problem asks:

```text
Can you solve it in-place?
```

or

```text
Use constant extra space.
```

Reversal often replaces the need for an auxiliary array.

---

### Signal 3 — Group Rearrangement

The desired output can be viewed as:

```text
Part A + Part B
```

becoming:

```text
Part B + Part A
```

Reversal is especially effective for these transformations.

---

## Generic Template

### Rotate Right by k

```text
reverse(whole array)

reverse(first k elements)

reverse(remaining elements)
```

---

### Rotate Left by k

```text
reverse(first k elements)

reverse(remaining elements)

reverse(whole array)
```

---

## Visual Explanation

### Original

```text
[A A A A | B B B]
```

Goal:

```text
[B B B | A A A A]
```

---

### Reverse Entire Array

```text
[B B B | A A A A]
↓
[B B B reversed | A A A A reversed]
```

Example:

```text
[7,6,5 | 4,3,2,1]
```

---

### Reverse First Segment

```text
[5,6,7 | 4,3,2,1]
```

---

### Reverse Remaining Segment

```text
[5,6,7 | 1,2,3,4]
```

Desired arrangement achieved.

---

## Complexity

| Operation | Complexity |
|------------|------------|
| Time | O(n) |
| Space | O(1) |

Because each element is swapped a constant number of times.

---

## Advantages

### Memory Efficient

Uses:

```text
O(1)
```

extra space.

---

### Interview Friendly

Easy to explain and implement.

---

### Production Ready

Simple logic.

Low bug risk.

Good maintainability.

---

## Common Pitfalls

### Forgetting Normalization

Always reduce:

```text
k %= n
```

before processing.

---

### Incorrect Reverse Order

The order of reversals matters.

Wrong sequence produces incorrect output.

---

### Boundary Errors

Be careful with:

```text
reverse(0, k - 1)

reverse(k, n - 1)
```

Off-by-one mistakes are common.

---

### Empty Arrays

Production code should safely handle:

```text
[]
```

even when constraints exclude them.

---

## Comparison With Other Approaches

| Approach | Time | Space |
|-----------|--------|--------|
| Repeated Shift | O(n × k) | O(1) |
| Extra Array | O(n) | O(n) |
| Cyclic Replacement | O(n) | O(1) |
| Reversal Technique | O(n) | O(1) |

---

## Recognition Checklist

Ask yourself:

### 1

Do elements wrap around?

✅ Yes

---

### 2

Is the operation a rotation?

✅ Yes

---

### 3

Is O(1) extra space desired?

✅ Yes

---

### 4

Can the array be viewed as:

```text
A + B
```

becoming:

```text
B + A
```

✅ Yes

---

If most answers are yes:

```text
Think Reversal Technique
```

---

# Pattern Examples

## Easy

### Reverse String

Basic reversal operation.

---

## Medium

### 189. Rotate Array

Pattern Use:

```text
Reverse All
Reverse First k
Reverse Remaining
```

Complexity:

```text
O(n)
O(1)
```

---

### 48. Rotate Image

Uses:

- Reverse rows
- Matrix transpose

to achieve rotation.

---

## Hard

### Advanced String Rotation Problems

Use reversal-based transformations combined with indexing logic.

---

# Related Patterns

### Array Manipulation

Frequently paired with reversal.

---

### Two Pointers

Used to perform swaps efficiently.

---

### Cyclic Replacement

Alternative solution for rotation problems.

---

### In-Place Algorithms

Reversal is a common in-place transformation strategy.

---

# Problem Registry

| LeetCode # | Problem | Difficulty |
|------------|---------|------------|
| 189 | Rotate Array | Medium |

---

# Key Takeaway

The Reversal Technique is one of the most valuable in-place transformation patterns.

Memorize:

```text
Rotate Right

↓

Reverse Entire Array

↓

Reverse First k

↓

Reverse Remaining Elements
```

This single pattern solves many array and string rotation problems with:

```text
O(n) Time
O(1) Space
```