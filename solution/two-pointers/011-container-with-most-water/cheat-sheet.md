# cheat-sheet.md

# Container With Most Water — Cheat Sheet

A one-page quick revision guide for **LeetCode 11 — Container With Most Water**.

---

# Pattern Summary

| Property              | Value              |
| --------------------- | ------------------ |
| **Problem Type**      | Array Optimization |
| **Primary Pattern**   | Two Pointers       |
| **Secondary Pattern** | Greedy             |
| **Difficulty**        | Medium             |
| **Time Complexity**   | **O(n)**           |
| **Space Complexity**  | **O(1)**           |

---

# Recognition Signals

Use the **Two Pointer** pattern when you notice:

* Two ends of an array are involved.
* The answer depends on a pair of elements.
* Brute force checks every pair (`O(n²)`).
* The search space can be reduced using a greedy observation.
* The input array does **not** need to be sorted.
* Moving one pointer can safely eliminate impossible candidates.

### Keywords

* Maximum area
* Pair of indices
* Optimize brute force
* Bidirectional traversal
* Linear scan
* Greedy decision

---

# Core Formula

```text
Area = min(height[left], height[right]) × (right - left)
```

Where:

* `min(height[left], height[right])` → Limiting height
* `(right - left)` → Width of the container

---

# Key Insight

> **The shorter line always limits the container's height.**

* Moving the taller pointer only decreases the width.
* The limiting height cannot increase by keeping the shorter line.
* Therefore, always move the pointer pointing to the **shorter** height.

---

# Algorithm Template

```text
left = 0
right = n - 1
maxArea = 0

while left < right

    width = right - left

    currentArea =
        min(height[left], height[right]) * width

    maxArea = max(maxArea, currentArea)

    if height[left] < height[right]
        left++
    else
        right--

return maxArea
```

---

# Pointer Movement Rules

| Condition                       | Action              | Reason                                                     |
| ------------------------------- | ------------------- | ---------------------------------------------------------- |
| `height[left] < height[right]`  | Move `left++`       | Try to find a taller limiting height.                      |
| `height[left] > height[right]`  | Move `right--`      | Right side is limiting the area.                           |
| `height[left] == height[right]` | Move either pointer | Both heights limit the current area; either move is valid. |

---

# Complexity Cheatsheet

| Approach     | Time       |      Space |
| ------------ | ---------- | ---------: |
| Brute Force  | `O(n²)`    |     `O(1)` |
| Two Pointers | **`O(n)`** | **`O(1)`** |

---

# Common Mistakes

❌ Using the taller height instead of the shorter height.

```text
Incorrect:
max(height[left], height[right])
```

```text
Correct:
min(height[left], height[right])
```

---

❌ Incorrect width calculation.

```text
Incorrect:
right - left + 1
```

```text
Correct:
right - left
```

---

❌ Moving the taller pointer.

This cannot increase the limiting height while always reducing the width.

---

❌ Updating pointers before calculating the area.

Always:

1. Compute the current area.
2. Update the maximum.
3. Move the appropriate pointer.

---

❌ Falling back to nested loops after identifying the Two Pointer pattern.

---

# Dry Run Snapshot

Example:

```text
height = [1,8,6,2,5,4,8,3,7]
```

| Left | Right | Area | Max | Move  |
| ---- | ----- | ---: | --: | ----- |
| 0    | 8     |    8 |   8 | Left  |
| 1    | 8     |   49 |  49 | Right |
| 1    | 7     |   18 |  49 | Right |
| 1    | 6     |   40 |  49 | Right |
| ...  | ...   |  ... |  49 | ...   |

**Final Answer:** `49`

---

# Edge Cases

* Minimum valid input (`n = 2`)
* All heights are equal
* Strictly increasing heights
* Strictly decreasing heights
* Heights containing zeros
* Very large arrays
* Duplicate heights

---

# Similar Problems

| Problem                                  | Pattern      |
| ---------------------------------------- | ------------ |
| LeetCode 125 — Valid Palindrome          | Two Pointers |
| LeetCode 167 — Two Sum II                | Two Pointers |
| LeetCode 15 — 3Sum                       | Two Pointers |
| LeetCode 42 — Trapping Rain Water        | Two Pointers |
| LeetCode 392 — Is Subsequence            | Two Pointers |
| LeetCode 977 — Squares of a Sorted Array | Two Pointers |

---

# Interview Sound Bites

Use these concise explanations during interviews:

### Explaining the Greedy Choice

> "The shorter line determines the current container height. Moving the taller line only reduces the width without increasing the limiting height, so it cannot produce a better result."

---

### Explaining Time Complexity

> "Each pointer moves inward at most once across the array, resulting in a single linear traversal."

---

### Explaining Space Complexity

> "The algorithm uses only a constant number of variables, so the extra space remains O(1)."

---

# Pattern Recognition Checklist

Before choosing the Two Pointer approach, ask:

* Is the answer based on **two positions** in an array?
* Can I start from opposite ends?
* Does moving one pointer eliminate impossible candidates?
* Is there a mathematical property that justifies the move?
* Can I improve a quadratic solution to linear time?

If most answers are **Yes**, the **Two Pointer** pattern is likely the correct choice.

---

# Quick Revision Notes

* Formula: `min(height[left], height[right]) × width`
* Width = `right - left`
* Shorter line limits the area.
* Always compute the area before moving pointers.
* Move the pointer at the shorter height.
* Maintain the maximum area seen so far.
* Every element is processed at most once.
* **Optimal Complexity:** `O(n)` time, `O(1)` space.

---

# Memory Trick

> **"Shorter wall? Move it."**

Why?

* Width will always shrink.
* The only hope of finding a larger area is discovering a taller limiting wall.
* Moving the taller wall cannot improve the current limitation.
