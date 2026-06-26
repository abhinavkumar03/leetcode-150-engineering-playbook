# Cheat Sheet — 209. Minimum Size Subarray Sum

> **Goal:** One-page revision guide for quickly recalling the pattern, algorithm, complexity, and interview insights.

---

# Pattern Summary

| Property | Value |
|----------|-------|
| Pattern | Variable Size Sliding Window |
| Category | Array |
| Difficulty | Medium |
| Time Complexity | **O(n)** |
| Space Complexity | **O(1)** |
| Interview Frequency | ⭐⭐⭐⭐⭐ |
| Key Requirement | **All numbers must be positive** |

---

# Recognition Signals

Consider a **Variable Size Sliding Window** when you see:

- ✅ Contiguous subarray or substring
- ✅ Need to find the **minimum** or **maximum** window
- ✅ Array elements are **all positive**
- ✅ Window expands until a condition is met
- ✅ Window shrinks to optimize the answer
- ✅ Single-pass optimization is expected

### Typical Keywords

- Minimum length
- Maximum length
- Contiguous
- At least
- Sum ≥ Target
- Smallest window
- Expand and shrink

---

# Algorithm Template

```text
Initialize:

left = 0
windowSum = 0
answer = Infinity

For right from 0 to n-1:

    Add nums[right] to windowSum

    While windowSum >= target:

        Update answer

        Remove nums[left]

        left++

Return answer if found, otherwise 0
```

---

# Visual Workflow

```text
             Expand Window
                  │
                  ▼
      +----------------------+
      | Add nums[right]      |
      +----------------------+
                  │
                  ▼
     Is windowSum >= target?
          /              \
        No                Yes
        │                  │
        ▼                  ▼
 Move right          Update answer
 pointer                │
                        ▼
               Remove nums[left]
                        │
                        ▼
                  Move left
                    pointer
                        │
                        └────── Repeat while valid
```

---

# Key Formula

### Window Sum

```text
windowSum += nums[right]
```

---

### Shrink Window

```text
windowSum -= nums[left]
left++
```

---

### Window Length

```text
right - left + 1
```

---

### Update Answer

```text
answer = min(answer, right - left + 1)
```

---

# Complexity Cheatsheet

| Approach | Time | Space |
|----------|------|-------|
| Brute Force | **O(n²)** | **O(1)** |
| Sliding Window | **O(n)** | **O(1)** |
| Prefix Sum + Binary Search | **O(n log n)** | **O(n)** |

---

# Sliding Window Invariant

At every iteration:

```text
windowSum
=
Sum of elements
between left and right (inclusive)
```

Once:

```text
windowSum >= target
```

The current window is valid, and we immediately try to shrink it.

---

# Why It Works

Because every element is **positive**:

```text
Expand Window
      ↓
Sum Increases
```

```text
Shrink Window
      ↓
Sum Decreases
```

This monotonic behavior ensures that each pointer only moves forward.

---

# Common Pitfalls

### ❌ Using `if` Instead of `while`

```java
if (windowSum >= target) { ... }
```

This only shrinks once and may miss a smaller valid window.

✅ Correct:

```java
while (windowSum >= target) { ... }
```

---

### ❌ Wrong Window Length

Incorrect:

```text
right - left
```

Correct:

```text
right - left + 1
```

---

### ❌ Updating the Answer Too Late

Always record the current window length **before** removing `nums[left]`.

---

### ❌ Ignoring the "Positive Numbers" Constraint

This approach **does not work** if negative values are allowed because the window sum is no longer monotonic.

---

# Edge Cases Checklist

| Scenario | Expected Output |
|----------|-----------------|
| Empty array | `0` |
| No valid subarray | `0` |
| Single element equals target | `1` |
| Single element less than target | `0` |
| Entire array needed | Array length |
| Multiple valid windows | Smallest length |
| Duplicates | Supported |
| Large arrays | Efficient (`O(n)`) |
| Negative values | ❌ Requires a different approach |

---

# Interview Checklist

Before coding, confirm:

- ✅ Input contains **positive integers**
- ✅ Need a **contiguous** subarray
- ✅ Looking for the **minimum** valid window
- ✅ Sliding Window is applicable

While coding:

- ✅ Expand using `right`
- ✅ Maintain `windowSum`
- ✅ Shrink using `left`
- ✅ Update answer before shrinking
- ✅ Return `0` if no solution exists

After coding:

- ✅ Explain why each pointer moves at most `n` times
- ✅ State `O(n)` time and `O(1)` space
- ✅ Mention why the algorithm fails with negative numbers

---

# Similar Problems

## Easy

- **643. Maximum Average Subarray I**
- **485. Max Consecutive Ones**

---

## Medium

- **3. Longest Substring Without Repeating Characters**
- **1004. Max Consecutive Ones III**
- **424. Longest Repeating Character Replacement**
- **904. Fruit Into Baskets**
- **713. Subarray Product Less Than K**
- **1493. Longest Subarray of 1's After Deleting One Element**
- **1658. Minimum Operations to Reduce X to Zero**

---

## Hard

- **76. Minimum Window Substring**
- **862. Shortest Subarray with Sum at Least K** *(supports negative numbers using Prefix Sum + Monotonic Queue)*

---

# Pattern Comparison

| Problem Type | Fixed Window | Variable Window |
|--------------|--------------|-----------------|
| Window Size | Constant | Changes dynamically |
| Expansion | Fixed | Until condition is met |
| Shrinking | Not required | Required |
| Typical Goal | Maximum/Average/Sum | Minimum/Maximum valid window |
| Example | LC 643 | LC 209 |

---

# Memory Trick

```text
Need a contiguous subarray?

            │
            ▼

Are all numbers positive?

      Yes            No
       │              │
       ▼              ▼
Sliding Window   Consider Prefix Sum,
                 Monotonic Queue,
                 or other techniques
```

---

# 30-Second Revision

- Pattern → **Variable Size Sliding Window**
- Keep two pointers: `left`, `right`
- Maintain a running `windowSum`
- Expand until `windowSum >= target`
- Shrink while the condition remains true
- Update the minimum length before shrinking
- Each element enters and exits the window once
- **Time:** `O(n)`
- **Space:** `O(1)`
- **Works only because all numbers are positive**