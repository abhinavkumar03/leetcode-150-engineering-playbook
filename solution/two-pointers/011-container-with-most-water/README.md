# Container With Most Water

## Problem Statement

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `iᵗʰ` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container capable of holding the maximum amount of water.

Return the maximum amount of water that can be stored.

**Constraints**

* `n >= 2`
* `0 <= height[i] <= 10^4`

---

## Difficulty

**Medium**

---

## Tags

* Array
* Two Pointers
* Greedy
* Optimization

---

## Pattern

**Primary Pattern:** Two Pointers

**Secondary Pattern:** Greedy Optimization

---

# Intuition

A container is formed by choosing two vertical lines.

The amount of water stored depends on:

* The shorter of the two heights.
* The distance between the two lines.

The area is calculated as:

```text
Area = min(leftHeight, rightHeight) × width
```

A brute-force solution checks every pair of lines, but this requires evaluating every possible combination.

Instead, we can observe that the width continuously decreases as pointers move inward. Therefore, whenever we reduce the width, the only hope of increasing the area is to find a taller limiting height.

This insight naturally leads to the Two Pointer approach.

---

# Key Observation

The shorter line always limits the amount of water.

Suppose:

```text
Left Height  = 3
Right Height = 10
```

The maximum possible area is limited by height **3**.

Moving the taller line inward only decreases the width while keeping the limiting height at most 3.

Therefore, moving the taller pointer can never produce a better answer.

Only moving the shorter pointer gives us a chance to discover a taller line and increase the limiting height.

This greedy decision reduces the search space from **O(n²)** to **O(n)**.

---

# Brute Force Approach

Check every possible pair of lines.

## Algorithm

1. Select every index as the first line.
2. Pair it with every line to its right.
3. Compute:

   * Width
   * Minimum height
   * Area
4. Track the maximum area.
5. Return the maximum.

---

## Complexity

| Metric |     Value |
| ------ | --------: |
| Time   | **O(n²)** |
| Space  |  **O(1)** |

---

## Limitations

* Evaluates every pair.
* Performs many unnecessary calculations.
* Does not leverage the relationship between width and height.
* Fails to scale efficiently for large inputs.

---

# Optimized Approach

Use two pointers:

* One at the beginning.
* One at the end.

At every step:

1. Calculate the current container area.
2. Update the maximum area.
3. Move the pointer pointing to the shorter line.
4. Continue until both pointers meet.

---

## Algorithm

```text
Initialize:

left = 0
right = n - 1
maxArea = 0

While left < right

    width = right - left

    currentArea =
        min(height[left], height[right]) × width

    Update maxArea

    If height[left] < height[right]
        left++
    Else
        right--

Return maxArea
```

---

## Why It Works

At every iteration:

* Width decreases by exactly one.
* To compensate for the smaller width, the limiting height must increase.
* The shorter line is the bottleneck.
* Keeping the shorter line while shrinking the width cannot improve the answer.
* Therefore, discarding the shorter line is always the optimal greedy choice.

Every index is processed at most once.

This results in linear time complexity.

---

## Complexity

| Metric |    Value |
| ------ | -------: |
| Time   | **O(n)** |
| Space  | **O(1)** |

---

# Edge Cases

| Case                    | Expected Behavior                                              |
| ----------------------- | -------------------------------------------------------------- |
| Empty input             | Not applicable (problem guarantees at least two elements).     |
| Two elements            | Area is calculated directly using those two lines.             |
| Single element          | Not allowed by constraints.                                    |
| Duplicate heights       | Algorithm works correctly since comparisons remain valid.      |
| All heights equal       | Maximum area is determined by the widest pair.                 |
| Increasing heights      | Left pointer advances until a better limiting height is found. |
| Decreasing heights      | Right pointer retreats for the same reason.                    |
| Heights containing zero | Containers involving a zero-height line may produce zero area. |
| Very large input        | Linear solution remains efficient.                             |

---

# Dry Run

### Example

```text
Input

height = [1,8,6,2,5,4,8,3,7]
```

| Step | Left | Right | Heights | Width | Area | Max Area | Move    |
| ---: | ---: | ----: | ------- | ----: | ---: | -------: | ------- |
|    1 |    0 |     8 | 1, 7    |     8 |    8 |        8 | Left++  |
|    2 |    1 |     8 | 8, 7    |     7 |   49 |       49 | Right-- |
|    3 |    1 |     7 | 8, 3    |     6 |   18 |       49 | Right-- |
|    4 |    1 |     6 | 8, 8    |     5 |   40 |       49 | Right-- |
|    5 |    1 |     5 | 8, 4    |     4 |   16 |       49 | Right-- |
|    6 |    1 |     4 | 8, 5    |     3 |   15 |       49 | Right-- |
|    7 |    1 |     3 | 8, 2    |     2 |    4 |       49 | Right-- |
|    8 |    1 |     2 | 8, 6    |     1 |    6 |       49 | Right-- |

```text
Final Answer = 49
```

---

# Common Mistakes

### 1. Moving the taller pointer

This may discard a potentially optimal width while keeping the limiting height unchanged.

---

### 2. Forgetting to update the answer before moving pointers

Always compute the current area first.

---

### 3. Using the larger height

The area depends on the shorter line:

```text
Incorrect:
max(height[left], height[right])

Correct:
min(height[left], height[right])
```

---

### 4. Incorrect width calculation

The width is:

```text
right - left
```

not

```text
right - left + 1
```

---

### 5. Using nested loops after learning Two Pointers

This unnecessarily increases time complexity from **O(n)** to **O(n²)**.

---

# Interview Discussion

Interviewers typically expect candidates to explain:

* Why brute force is inefficient.
* How the area formula is derived.
* Why only the shorter pointer moves.
* The greedy proof behind the optimization.
* Time and space complexity.
* Why each index is visited only once.

A strong candidate explains the reasoning rather than simply memorizing the algorithm.

---

# Follow-up Questions

1. Can you prove why moving the taller pointer is never beneficial?
2. Can this problem be solved using Dynamic Programming? Why or why not?
3. What changes if the width between lines is not uniform?
4. How would you return the indices instead of only the maximum area?
5. Can this technique be generalized to similar optimization problems?

---

# Real World Applications

The underlying optimization strategy appears in many engineering scenarios:

* Maximizing storage capacity under physical constraints.
* Network bandwidth optimization.
* Resource allocation problems.
* Computational geometry.
* Image processing and boundary detection.
* Performance optimization using bidirectional scanning.

---

# Related Problems

| LeetCode | Problem                            | Pattern      |
| -------- | ---------------------------------- | ------------ |
| 125      | Valid Palindrome                   | Two Pointers |
| 167      | Two Sum II – Input Array Is Sorted | Two Pointers |
| 15       | 3Sum                               | Two Pointers |
| 42       | Trapping Rain Water                | Two Pointers |
| 392      | Is Subsequence                     | Two Pointers |
| 844      | Backspace String Compare           | Two Pointers |
| 977      | Squares of a Sorted Array          | Two Pointers |

---
