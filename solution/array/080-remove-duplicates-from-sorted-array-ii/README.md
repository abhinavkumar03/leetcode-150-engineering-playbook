# Remove Duplicates from Sorted Array II

## Problem Statement

Given an integer array `nums` sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice.

The relative order of the elements should remain the same.

Since it is impossible to change the length of the array in some languages, you must instead place the result in the first part of the array and return the new length `k`.

After removing duplicates, the first `k` elements of `nums` should contain the final result.

### Example 1

Input:

nums = [1,1,1,2,2,3]

Output:

k = 5

nums = [1,1,2,2,3,_]

### Example 2

Input:

nums = [0,0,1,1,1,1,2,3,3]

Output:

k = 7

nums = [0,0,1,1,2,3,3,_,_]

---

## Difficulty

Medium

---

## Tags

- Array
- Two Pointers

---

## Pattern

### Primary Pattern

Two Pointers

### Secondary Pattern

In-Place Array Modification

---

## Intuition

The array is already sorted.

This means:

- Duplicate values always appear next to each other.
- We can process elements from left to right.
- We only need to determine whether the current value should be kept.

The problem allows:

- First occurrence → keep
- Second occurrence → keep
- Third occurrence and beyond → discard

Instead of creating a new array, we overwrite unwanted elements while maintaining a write position.

---

## Key Observation

For every element after the first two positions:

If the current value is different from the element located two positions before the write pointer, then it can be safely included.

Why?

Because:

- If nums[i] equals nums[write - 2]
- Then we already have two identical values in the valid portion.
- Adding another would violate the constraint.

This allows duplicate control using a single comparison.

Example:

```text
Valid Portion:

[1,1]

Current = 1

Compare:
current == nums[write - 2]

1 == 1

Already have two copies.

Skip.
```

---

## Brute Force Approach

### Idea

Create a separate result array.

Count occurrences of each value.

Insert a value only if its frequency does not exceed 2.

### Algorithm

1. Create an empty result array.
2. Traverse nums.
3. Count occurrences.
4. Insert at most two copies.
5. Copy result back if required.

### Complexity

Time Complexity:

O(n)

Space Complexity:

O(n)

### Limitations

- Uses additional memory.
- Violates the constant-space requirement.
- Less attractive in interviews.

---

## Optimized Approach

### Idea

Use two pointers:

- Read Pointer → scans the array
- Write Pointer → marks where the next valid value should be written

Keep the first two elements automatically.

Starting from index 2:

- Compare current value with nums[write - 2]
- If different → write it
- If same → skip it

### Algorithm

1. If array length <= 2, return length.
2. Initialize write = 2.
3. Iterate read from index 2 to n - 1.
4. Compare nums[read] with nums[write - 2].
5. If different:
   - nums[write] = nums[read]
   - write++
6. Return write.

### Why It Works

The first two occurrences of a number are always accepted.

When examining a new element:

```text
nums[read] != nums[write - 2]
```

means:

The value does not already appear twice in the valid portion.

Therefore it is safe to include.

This guarantees:

- At most two occurrences
- Correct ordering
- Constant extra space

---

## Complexity

### Time Complexity

```text
O(n)
```

Each element is visited once.

### Space Complexity

```text
O(1)
```

No extra array is used.

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

---

### Two Elements

```text
[5,5]
```

Output:

```text
2
```

Both are allowed.

---

### Multiple Duplicates

```text
[1,1,1,1]
```

Output:

```text
[1,1]
```

Only two copies remain.

---

### Negative Values

```text
[-2,-2,-2,-1,-1]
```

Output:

```text
[-2,-2,-1,-1]
```

Algorithm works identically.

---

### Large Inputs

```text
100,000+ elements
```

Still:

```text
O(n)
```

with constant memory.

---

## Dry Run

Input:

```text
[1,1,1,2,2,3]
```

Initial:

```text
write = 2
```

| Read | Current | Compare With nums[write-2] | Action | Array State |
|--------|----------|----------|----------|----------|
| 2 | 1 | 1 | Skip | [1,1,1,2,2,3] |
| 3 | 2 | 1 | Write | [1,1,2,2,2,3] |
| 4 | 2 | 1 | Write | [1,1,2,2,2,3] |
| 5 | 3 | 2 | Write | [1,1,2,2,3,3] |

Final:

```text
Length = 5

[1,1,2,2,3]
```

---

## Common Mistakes

### Mistake 1

Comparing with previous element only.

```text
nums[i] != nums[i - 1]
```

This solves Problem 26, not Problem 80.

---

### Mistake 2

Using extra arrays.

```text
result = []
```

Violates O(1) space requirement.

---

### Mistake 3

Starting write pointer at wrong position.

Correct:

```text
write = 2
```

because two duplicates are allowed.

---

### Mistake 4

Forgetting small-array handling.

```text
if len(nums) <= 2
```

must be handled immediately.

---

## Interview Discussion

### What Makes This Problem Interesting?

The challenge is not removing duplicates.

The challenge is enforcing a frequency constraint while:

- Preserving order
- Using constant memory
- Modifying the array in-place

---

### Generalization

A common follow-up:

```text
Allow at most K duplicates.
```

Replace:

```text
nums[write - 2]
```

with:

```text
nums[write - K]
```

This creates a reusable template.

---

### Why Two Pointers?

Because:

- Read pointer explores input.
- Write pointer builds output.
- No shifting required.
- O(n) runtime.

---

## Follow-up Questions

### 1. Allow At Most K Duplicates

How would you modify the solution?

Answer:

```text
Compare against nums[write - K]
```

---

### 2. What If Array Is Not Sorted?

Need:

- HashMap frequency tracking

Complexities:

```text
Time: O(n)
Space: O(n)
```

---

### 3. Can This Be Done With One Pointer?

Not cleanly.

Two pointers provide clearer separation of concerns.

---

### 4. Why Is Sorting Important?

Sorting groups duplicates together.

Without sorting, the simple comparison trick fails.

---

## Real World Applications

### Data Compression

Removing excessive duplicate records.

---

### Log Processing

Keeping only a limited number of repeated events.

---

### Analytics Pipelines

Reducing noise from repeated values.

---

### Streaming Systems

Maintaining bounded occurrences of repeated data.

---

## Related Problems

### Easy

- 26. Remove Duplicates from Sorted Array
- 283. Move Zeroes

### Medium

- 27. Remove Element
- 75. Sort Colors
- 189. Rotate Array

### Advanced Two-Pointer Problems

- 15. 3Sum
- 167. Two Sum II
- 977. Squares of a Sorted Array

---

## Key Takeaway

This is a classic in-place array compression problem.

The crucial insight is:

```text
Keep an element only when:

nums[read] != nums[write - 2]
```

This single comparison ensures:

- Maximum two occurrences
- O(n) runtime
- O(1) space
- Stable ordering