# 15. 3Sum

## Problem Statement

Given an integer array `nums`, return all the unique triplets `[nums[i], nums[j], nums[k]]` such that:

- `i != j`
- `i != k`
- `j != k`
- `nums[i] + nums[j] + nums[k] == 0`

The solution set must not contain duplicate triplets.

---

## Difficulty

**Medium**

---

## Tags

- Array
- Sorting
- Two Pointers

---

## Pattern

**Primary Pattern:** Two Pointers

**Secondary Pattern:** Sorting

---

# Intuition

The brute force solution checks every possible combination of three numbers. While simple, it requires three nested loops, resulting in **O(n³)** time complexity.

A better observation is:

- If the array is sorted,
- We can fix one number,
- Then use two pointers to efficiently search for the remaining two numbers.

Sorting transforms an expensive search problem into a structured traversal problem.

Instead of searching blindly, the sorted order tells us exactly which pointer should move depending on whether the current sum is too small or too large.

---

# Key Observation

After sorting:

```
[-4, -1, -1, 0, 1, 2]
```

Fix one element:

```
-1
```

Now the problem becomes:

> Find two numbers whose sum equals **1**.

This is exactly a classic **Two Pointer** problem.

---

# Brute Force Approach

Check every possible triplet.

```
for i
    for j
        for k
            if sum == 0
```

### Algorithm

1. Iterate over every first element.
2. Iterate over every second element.
3. Iterate over every third element.
4. Check whether the sum equals zero.
5. Store unique triplets.

### Complexity

| Complexity | Value |
|------------|-------|
| Time | O(n³) |
| Space | O(1) (excluding output) |

### Limitations

- Extremely slow for large arrays.
- Generates many duplicate triplets.
- Requires additional duplicate checking.
- Poor scalability.

---

# Optimized Approach

## Idea

1. Sort the array.
2. Fix one element.
3. Use two pointers to find the remaining pair.
4. Skip duplicates while traversing.

---

### Algorithm

Sort the array.

For every index `i`:

- Skip duplicate values.
- Set

```
left = i + 1
right = n - 1
```

While `left < right`

Compute

```
sum = nums[i] + nums[left] + nums[right]
```

Three possibilities exist:

### Case 1

```
sum < 0
```

Increase the sum.

```
left++
```

---

### Case 2

```
sum > 0
```

Decrease the sum.

```
right--
```

---

### Case 3

```
sum == 0
```

Triplet found.

Store it.

Move both pointers.

Skip duplicate values on both sides.

Continue searching.

---

## Why It Works

Sorting guarantees:

- Numbers increase from left to right.
- Moving the left pointer always increases the sum.
- Moving the right pointer always decreases the sum.

Therefore, every pointer movement is meaningful.

Each pair is visited only once.

---

## Complexity

| Complexity | Value |
|------------|-------|
| Time | O(n²) |
| Space | O(1) (excluding output and sorting implementation) |

Sorting contributes:

```
O(n log n)
```

The two-pointer scan dominates:

```
O(n²)
```

Overall:

```
O(n²)
```

---

# Edge Cases

## Empty Input

```
[]
```

Output

```
[]
```

---

## Less Than Three Elements

```
[1]
```

```
[1,2]
```

Output

```
[]
```

---

## All Zeros

Input

```
[0,0,0]
```

Output

```
[[0,0,0]]
```

Only one unique triplet should be returned.

---

## Duplicate Values

Input

```
[-1,-1,-1,2,2]
```

Output

```
[[-1,-1,2]]
```

Duplicate triplets must be skipped.

---

## No Valid Triplets

Input

```
[5,7,9]
```

Output

```
[]
```

---

## Negative Values

Input

```
[-5,-2,-1,3,5]
```

Algorithm works naturally after sorting.

---

## Large Inputs

The optimized solution remains efficient because it avoids exploring every possible triplet.

---

# Dry Run

Input

```
[-1,0,1,2,-1,-4]
```

After sorting

```
[-4,-1,-1,0,1,2]
```

---

### Iteration 1

Fixed element

```
i = 0

-4
```

Target pair sum

```
4
```

| Left | Right | Sum | Action |
|------|-------|-----|--------|
| -1 | 2 | -3 | Move left |
| -1 | 2 | -3 | Move left |
| 0 | 2 | -2 | Move left |
| 1 | 2 | -1 | Move left |

No triplet.

---

### Iteration 2

Fixed element

```
i = 1

-1
```

| Left | Right | Sum | Action |
|------|-------|-----|--------|
| -1 | 2 | 0 | Store [-1,-1,2] |
| 0 | 1 | 0 | Store [-1,0,1] |

---

### Iteration 3

```
i = 2
```

Duplicate.

Skip.

---

### Iteration 4

```
i = 3
```

Pointers cross.

Stop.

---

Final Answer

```
[
[-1,-1,2],
[-1,0,1]
]
```

---

# Common Mistakes

### Forgetting to Sort

Two pointers only work on sorted arrays.

---

### Not Skipping Duplicate Fixed Elements

Incorrect

```
[-1,-1,-1,2]
```

Produces duplicate answers.

Always skip

```cpp
if (i > 0 && nums[i] == nums[i-1])
```

---

### Forgetting Duplicate Removal After Finding a Triplet

After storing a triplet:

```
left++
right--
```

Skip equal values on both pointers before continuing.

---

### Moving the Wrong Pointer

If

```
sum < 0
```

Increase the sum.

Move

```
left++
```

If

```
sum > 0
```

Decrease the sum.

Move

```
right--
```

---

### Pointer Out-of-Bounds

Always maintain

```
left < right
```

before accessing values.

---

# Interview Discussion

A strong interview explanation follows this progression:

1. Describe the brute force O(n³) approach.
2. Explain why it is inefficient.
3. Observe that sorting introduces order.
4. Show how fixing one element reduces the problem to Two Sum.
5. Apply the Two Pointer technique.
6. Explain duplicate handling carefully.
7. Analyze complexity and tradeoffs.

Interviewers value the reasoning process more than simply presenting the final code.

---

# Follow-up Questions

### Can you solve this without sorting?

Possible using hashing, but:

- Extra memory is required.
- Duplicate handling becomes more complicated.

---

### Why is sorting preferred?

Sorting enables:

- Constant extra space.
- Simpler logic.
- Efficient pointer movement.
- Easier duplicate elimination.

---

### What changes if the target is not zero?

Instead of searching for

```
0
```

search for

```
target - nums[i]
```

The overall approach remains the same.

---

### How would you solve 4Sum?

Fix two elements.

Apply Two Pointers to the remaining array.

Complexity:

```
O(n³)
```

---

### Can this idea generalize?

Yes.

This forms the foundation of the generic **k-Sum** algorithm using recursion and Two Pointers.

---

# Real World Applications

Although the exact "3Sum" problem is interview-oriented, the underlying techniques are widely used.

### Financial Analytics

Finding combinations of transactions that offset each other to achieve a target balance.

---

### Risk Management

Detecting groups of assets whose combined exposure equals a desired threshold.

---

### Recommendation Systems

Searching for combinations of items that satisfy pricing or compatibility constraints.

---

### Data Deduplication

Eliminating duplicate combinations while preserving unique results.

---

### Constraint Satisfaction

Many optimization and scheduling problems reduce to searching for combinations that meet a target value.

---

# Related Problems

| Problem | Difficulty | Pattern |
|----------|------------|---------|
| 1. Two Sum | Easy | Hash Map |
| 167. Two Sum II | Medium | Two Pointers |
| 11. Container With Most Water | Medium | Two Pointers |
| 16. 3Sum Closest | Medium | Two Pointers |
| 18. 4Sum | Medium | Two Pointers |
| 259. 3Sum Smaller | Medium | Two Pointers |
| 454. 4Sum II | Medium | Hash Map |
| k-Sum (Generalized) | Hard | Recursion + Two Pointers |

---

# Key Takeaways

- Sorting transforms an unstructured search into an ordered traversal.
- Fix one element and solve the remaining problem using Two Pointers.
- Duplicate handling is essential for correctness.
- The optimized solution improves from **O(n³)** to **O(n²)**.
- This problem is the foundation for many higher-order sum problems such as **3Sum Closest**, **4Sum**, and the generalized **k-Sum** family.