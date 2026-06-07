# Rotate Array

## Problem Statement

Given an integer array `nums`, rotate the array to the right by `k` steps, where `k` is non-negative.

### Example 1

Input:

nums = [1,2,3,4,5,6,7], k = 3

Output:

[5,6,7,1,2,3,4]

### Example 2

Input:

nums = [-1,-100,3,99], k = 2

Output:

[3,99,-1,-100]

### Constraints

- 1 <= nums.length <= 10⁵
- -2³¹ <= nums[i] <= 2³¹ - 1
- 0 <= k <= 10⁵

---

## Difficulty

**Medium**

---

## Tags

- Array
- Two Pointers
- In-Place Algorithm
- Reversal Technique
- Simulation

---

## Pattern

### Primary Pattern

Array Manipulation

### Secondary Pattern

Reversal Technique

---

## Intuition

Rotating an array means moving elements from one position to another while preserving their relative order.

For example:

```text
[1,2,3,4,5,6,7]
```

Rotate by 3:

```text
[5,6,7,1,2,3,4]
```

A straightforward approach is to create a new array and place each element into its rotated position.

However, interviewers often expect an **O(1) extra space** solution.

The key observation is that rotation can be achieved through a sequence of reversals.

---

## Key Observation

Consider:

```text
[1,2,3,4,5,6,7]
```

Rotate by:

```text
k = 3
```

### Step 1: Reverse Entire Array

```text
[7,6,5,4,3,2,1]
```

### Step 2: Reverse First k Elements

```text
[5,6,7,4,3,2,1]
```

### Step 3: Reverse Remaining Elements

```text
[5,6,7,1,2,3,4]
```

Result achieved without extra memory.

This transformation forms the foundation of the optimized solution.

---

## Brute Force Approach

### Idea

Create a new array and place each element into its final rotated position.

Formula:

```text
newIndex = (currentIndex + k) % n
```

### Algorithm

1. Create an auxiliary array of size n.
2. Iterate through original array.
3. Compute destination index.
4. Store element in destination position.
5. Copy auxiliary array back to original array.

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n) |
| Space | O(n) |

### Limitations

- Requires additional memory.
- Not considered optimal for interview settings.
- Less efficient for very large arrays.

---

## Optimized Approach

### Idea

Use the Reversal Technique.

Instead of moving every element individually:

1. Reverse entire array.
2. Reverse first k elements.
3. Reverse remaining elements.

### Algorithm

Let:

```text
n = nums.length
k = k % n
```

#### Step 1

Reverse:

```text
0 → n-1
```

#### Step 2

Reverse:

```text
0 → k-1
```

#### Step 3

Reverse:

```text
k → n-1
```

Return modified array.

---

### Why It Works

Suppose:

```text
[1,2,3,4,5,6,7]
```

Desired split:

```text
[1,2,3,4] [5,6,7]
```

After reversing entire array:

```text
[7,6,5] [4,3,2,1]
```

Each segment appears reversed.

Reversing each segment again restores internal order while keeping the rotated positions.

Final result:

```text
[5,6,7] [1,2,3,4]
```

This achieves rotation using only swaps.

---

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n) |
| Space | O(1) |

Because each element participates in a constant number of swaps.

---

## Edge Cases

### Empty Input

Although constraints guarantee at least one element, production systems should handle:

```text
[]
```

Safely.

---

### Single Element

```text
[10]
k = 100
```

Output:

```text
[10]
```

No change.

---

### k Greater Than Array Length

```text
nums = [1,2,3]
k = 8
```

Effective rotation:

```text
8 % 3 = 2
```

Output:

```text
[2,3,1]
```

---

### Duplicates

```text
[1,1,1,2,2]
```

Rotation remains valid because position changes are independent of value uniqueness.

---

### Negative Values

```text
[-1,-100,3,99]
```

Rotation logic remains unchanged because values are never compared.

---

### Large Inputs

```text
n = 100000
```

Reversal solution remains efficient:

```text
O(n) time
O(1) space
```

---

## Dry Run

### Input

```text
nums = [1,2,3,4,5,6,7]
k = 3
```

### Effective Rotation

```text
k = 3 % 7 = 3
```

### Step 1: Reverse Entire Array

| Before | After |
|----------|----------|
| [1,2,3,4,5,6,7] | [7,6,5,4,3,2,1] |

---

### Step 2: Reverse First k Elements

Reverse:

```text
[7,6,5]
```

Result:

| Before | After |
|----------|----------|
| [7,6,5,4,3,2,1] | [5,6,7,4,3,2,1] |

---

### Step 3: Reverse Remaining Elements

Reverse:

```text
[4,3,2,1]
```

Result:

| Before | After |
|----------|----------|
| [5,6,7,4,3,2,1] | [5,6,7,1,2,3,4] |

---

### Final Output

```text
[5,6,7,1,2,3,4]
```

---

## Common Mistakes

### Forgetting k % n

Incorrect:

```text
k = 15
n = 7
```

Without modulo handling, index calculations become inefficient or incorrect.

---

### Reversing Wrong Ranges

Correct:

```text
reverse(0, n-1)
reverse(0, k-1)
reverse(k, n-1)
```

A small boundary mistake causes incorrect output.

---

### Using Extra Space Unnecessarily

Many candidates immediately choose:

```text
extra array
```

While valid, it misses the intended optimization.

---

### Not Handling k = 0

```text
k = 0
```

Array should remain unchanged.

---

## Interview Discussion

Interviewers often evaluate whether a candidate can progress through:

### Level 1

Brute force shifting.

### Level 2

Auxiliary array solution.

### Level 3

Cyclic replacement approach.

### Level 4

Reversal technique.

The reversal solution is generally considered the cleanest and most interview-friendly answer.

Topics frequently discussed:

- Why reversal works
- Space optimization
- In-place transformations
- Modulo arithmetic
- Tradeoffs between readability and performance

---

## Follow-up Questions

### Can you solve it with O(1) extra space?

Expected answer:

Reversal Technique.

---

### Can you solve it using cyclic replacements?

Expected discussion:

Track cycles formed by index mapping.

---

### What if rotation is to the left?

Adjust split point accordingly.

---

### What if the array is immutable?

Use an auxiliary array.

---

### Can rotation be done in a stream?

Requires different data structures such as circular buffers.

---

## Real World Applications

### Circular Buffers

Used in:

- Logging systems
- Audio processing
- Network packet queues

---

### Scheduling Systems

Rotating task assignments among workers.

---

### Distributed Systems

Leader rotation and token passing.

---

### Game Development

Rotating player turns in multiplayer games.

---

### Data Processing Pipelines

Shifting processing windows over datasets.

---

## Related Problems

### Easy

- 26. Remove Duplicates from Sorted Array
- 27. Remove Element
- 88. Merge Sorted Array

### Medium

- 48. Rotate Image
- 238. Product of Array Except Self
- 280. Wiggle Sort

### Advanced Rotation Concepts

- Circular Array Loop
- Rotate Image
- Shift 2D Grid
- Cyclic Replacement Problems
