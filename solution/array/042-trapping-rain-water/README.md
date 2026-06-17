# Trapping Rain Water

## Problem Statement

Given `n` non-negative integers representing an elevation map where the width of each bar is 1, compute how much water can be trapped after raining.

### Example 1

Input:

```text
height = [0,1,0,2,1,0,1,3,2,1,2,1]
```

Output:

```text
6
```

### Example 2

Input:

```text
height = [4,2,0,3,2,5]
```

Output:

```text
9
```

---

## Difficulty

**Hard**

---

## Tags

- Array
- Two Pointers
- Dynamic Programming
- Monotonic Stack
- Prefix Maximum
- Space Optimization

---

## Pattern

**Primary Pattern:** Two Pointers

**Secondary Patterns:**

- Prefix/Suffix Maximum Arrays
- Dynamic Programming
- Monotonic Stack

---

## Intuition

Water can only be trapped if there are taller bars on both sides.

For any position:

```text
Water = min(leftMax, rightMax) - currentHeight
```

Where:

- leftMax = tallest bar to the left
- rightMax = tallest bar to the right

The amount of water above a bar is determined by the shorter boundary because water would overflow from that side first.

---

## Key Observation

For every index:

```text
trappedWater[i] =
min(maxHeightOnLeft, maxHeightOnRight)
- height[i]
```

Example:

```text
Height:     [4,2,0,3,2,5]
Left Max:   [4,4,4,4,4,5]
Right Max:  [5,5,5,5,5,5]

Index 2:

min(4,5) - 0 = 4
```

Instead of recomputing left and right maximum values repeatedly, we can store or dynamically maintain them.

---

## Brute Force Approach

For every index:

1. Scan left side to find maximum height.
2. Scan right side to find maximum height.
3. Compute trapped water.
4. Add to total answer.

### Algorithm

```text
for each index i:

    leftMax = max(height[0...i])

    rightMax = max(height[i...n-1])

    water += min(leftMax, rightMax) - height[i]
```

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n²) |
| Space | O(1) |

### Limitations

- Repeatedly scans the same elements.
- Inefficient for large inputs.
- Fails scalability requirements in interviews.

---

## Optimized Approach

### Two Pointer Technique

Maintain:

```text
left
right

leftMax
rightMax
```

Move pointers inward.

At any moment:

```text
If leftMax <= rightMax

Then trapped water at left
depends only on leftMax.
```

Similarly:

```text
If rightMax < leftMax

Then trapped water at right
depends only on rightMax.
```

This eliminates the need for extra arrays.

---

### Algorithm

```text
Initialize:

left = 0
right = n - 1

leftMax = 0
rightMax = 0

while left < right

    update leftMax
    update rightMax

    if leftMax < rightMax

        water += leftMax - height[left]
        left++

    else

        water += rightMax - height[right]
        right--

return water
```

---

### Why It Works

When:

```text
leftMax < rightMax
```

the left side is the limiting boundary.

Even if a taller bar exists farther right, the amount of water above the current left position is already determined.

Therefore:

```text
waterAtLeft = leftMax - height[left]
```

can be calculated immediately.

This guarantees correctness while processing each index exactly once.

---

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n) |
| Space | O(1) |

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

No bars exist.

---

### Single Element

```text
[5]
```

Output:

```text
0
```

Water requires two boundaries.

---

### Two Elements

```text
[5,2]
```

Output:

```text
0
```

Still insufficient boundaries.

---

### Duplicates

```text
[3,3,3]
```

Output:

```text
0
```

No valleys are formed.

---

### Negative Values

Not applicable.

LeetCode guarantees:

```text
height[i] >= 0
```

---

### Large Inputs

```text
100000+ elements
```

The O(n) two-pointer solution remains efficient.

---

## Dry Run

Input:

```text
height = [4,2,0,3,2,5]
```

| Step | Left | Right | LeftMax | RightMax | Water Added | Total |
|--------|--------|--------|--------|--------|--------|--------|
| 1 | 0 | 5 | 4 | 5 | 0 | 0 |
| 2 | 1 | 5 | 4 | 5 | 2 | 2 |
| 3 | 2 | 5 | 4 | 5 | 4 | 6 |
| 4 | 3 | 5 | 4 | 5 | 1 | 7 |
| 5 | 4 | 5 | 4 | 5 | 2 | 9 |

Result:

```text
9
```

---

## Common Mistakes

### Forgetting the Minimum Boundary

Incorrect:

```text
max(leftMax, rightMax)
```

Correct:

```text
min(leftMax, rightMax)
```

---

### Adding Negative Water

Incorrect:

```text
water += boundary - height[i]
```

without proper boundary handling.

Always ensure boundary values are updated first.

---

### Using O(n²) During Interviews

Many candidates stop at brute force.

Interviewers usually expect:

```text
O(n)
```

solution after discussion.

---

### Off-by-One Errors

Incorrect pointer movement order can skip bars or double count them.

---

## Interview Discussion

### Expected Progression

Most interviewers expect:

```text
Brute Force
     ↓
Prefix/Suffix Arrays
     ↓
Two Pointers
```

Candidates should demonstrate the optimization journey.

---

### Key Insight To Communicate

Say:

> Water level at any position is determined by the shorter of the tallest bars on both sides.

This statement often leads directly to the optimal solution.

---

## Follow-up Questions

### Can you solve it using Dynamic Programming?

Yes.

Store:

```text
leftMax[]
rightMax[]
```

Time:

```text
O(n)
```

Space:

```text
O(n)
```

---

### Can you solve it using a Stack?

Yes.

Use a monotonic decreasing stack to identify valleys.

Time:

```text
O(n)
```

Space:

```text
O(n)
```

---

### Which approach is best?

| Approach | Time | Space |
|-----------|--------|--------|
| Brute Force | O(n²) | O(1) |
| DP Arrays | O(n) | O(n) |
| Monotonic Stack | O(n) | O(n) |
| Two Pointers | O(n) | O(1) |

The two-pointer solution is generally preferred.

---

## Real World Applications

### Terrain Simulation

Estimating water accumulation between hills and valleys.

---

### Flood Prediction Systems

Analyzing low-lying areas where water may collect.

---

### 3D Graphics

Calculating volume containment and surface effects.

---

### Urban Drainage Planning

Modeling rainwater retention between structures.

---

### Resource Capacity Analysis

Computing capacity between bounded structures.

---

## Related Problems

### Easy

- 11. Container With Most Water
- 26. Remove Duplicates from Sorted Array

### Medium

- 739. Daily Temperatures
- 84. Largest Rectangle in Histogram

### Hard

- 407. Trapping Rain Water II
- 85. Maximal Rectangle

---

## Summary

The core formula is:

```text
waterAtIndex =
min(leftMax, rightMax)
- height[i]
```

While multiple approaches exist, the optimal solution uses:

```text
Two Pointers
Time  : O(n)
Space : O(1)
```

making it the preferred interview solution.