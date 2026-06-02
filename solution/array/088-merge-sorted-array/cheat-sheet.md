## Pattern Summary

### Primary Pattern

```text
Two Pointers
```

### Secondary Pattern

```text
Reverse Traversal
In-Place Merge
```

### Difficulty

```text
Easy
```

---

# Recognition Signals

Look for this pattern when:

✅ Two arrays are already sorted

✅ Arrays need to be merged

✅ Space optimization is required

✅ Interview mentions:

```text
"in-place"

"without extra memory"

"O(1) space"
```

✅ Need to preserve sorted order

---

# Core Insight

Instead of merging from the beginning:

```text
❌ Left → Right
```

Merge from the end:

```text
✅ Right → Left
```

Reason:

```text
nums1 already contains free space at the end.
```

---

# Template

```go
i := m - 1
j := n - 1
k := m + n - 1

for i >= 0 && j >= 0 {
    if nums1[i] > nums2[j] {
        nums1[k] = nums1[i]
        i--
    } else {
        nums1[k] = nums2[j]
        j--
    }
    k--
}

for j >= 0 {
    nums1[k] = nums2[j]
    j--
    k--
}
```

---

# Pointer Meaning

| Pointer | Purpose                                 |
| ------- | --------------------------------------- |
| i       | Last valid element in nums1             |
| j       | Last element in nums2                   |
| k       | Last available position in merged array |

---

# Visual Memory Trick

```text
nums1 = [1,2,3,0,0,0]
             ↑
             i

nums2 = [2,5,6]
             ↑
             j

Merged Position
               ↑
               k
```

Always place:

```text
max(nums1[i], nums2[j])
```

at:

```text
nums1[k]
```

---

# Key Formula

```text
i = m - 1
j = n - 1
k = m + n - 1
```

Comparison:

```text
if nums1[i] > nums2[j]
```

Placement:

```text
nums1[k] = larger value
```

---

# Complexity Cheatsheet

| Approach             | Time              | Space  |
| -------------------- | ----------------- | ------ |
| Merge + Sort         | O((m+n) log(m+n)) | O(m+n) |
| Temp Array Merge     | O(m+n)            | O(m+n) |
| Reverse Two Pointers | O(m+n)            | O(1)   |

### Optimal

```text
Time  : O(m+n)
Space : O(1)
```

---

# Common Pitfalls

## Mistake 1

Merging from the front.

```text
❌ Overwrites valid values.
```

---

## Mistake 2

Forgetting remaining nums2 elements.

Need:

```go
for j >= 0 {
    nums1[k] = nums2[j]
    j--
    k--
}
```

---

## Mistake 3

Copying remaining nums1 elements.

```text
❌ Unnecessary
```

They are already correctly positioned.

---

## Mistake 4

Wrong pointer initialization.

Correct:

```text
i = m - 1
j = n - 1
k = m + n - 1
```

---

# Edge Cases Checklist

### nums2 Empty

```text
nums1 = [1]
nums2 = []
```

---

### nums1 Empty

```text
nums1 = [0]
nums2 = [1]
```

---

### Single Element Arrays

```text
nums1 = [2,0]
nums2 = [1]
```

---

### Duplicate Values

```text
nums1 = [1,2,2,0,0]
nums2 = [2,2]
```

---

### Negative Values

```text
nums1 = [-5,-2,0,0]
nums2 = [-4,-1]
```

---

# Interview One-Liner

> Since both arrays are sorted and nums1 already has free space at the end, we can merge from right to left using three pointers, achieving O(m+n) time and O(1) extra space.

---

# Similar Problems

### Same Pattern

* LeetCode 21: Merge Two Sorted Lists
* LeetCode 977: Squares of a Sorted Array
* LeetCode 167: Two Sum II
* LeetCode 283: Move Zeroes

---

### Merge-Based Problems

* LeetCode 56: Merge Intervals
* LeetCode 23: Merge k Sorted Lists
* LeetCode 148: Sort List

---

# Quick Revision (30 Seconds Before Interview)

```text
Pattern:
Two Pointers

Observation:
Both arrays are sorted.

Problem:
Forward merge overwrites data.

Solution:
Merge from end.

Pointers:
i = m - 1
j = n - 1
k = m + n - 1

Action:
Place larger value at nums1[k]

Complexity:
Time  = O(m+n)
Space = O(1)
```

---

# Recruiter / Portfolio Takeaway

This problem demonstrates:

* Array manipulation skills
* Space optimization thinking
* Two-pointer mastery
* In-place algorithm design
* Ability to move from brute force to optimal solutions

These are foundational skills expected from backend, systems, and product engineers.