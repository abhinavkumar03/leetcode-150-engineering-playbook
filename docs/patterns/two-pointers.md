# Pattern Definition

## Two Pointers

The Two Pointers pattern uses two indices that traverse a data structure (typically an array, string, or linked list) to efficiently solve problems involving searching, comparison, partitioning, merging, or window processing.

Instead of using nested loops, two pointers often reduce time complexity from **O(n²)** to **O(n)**.

---

## When To Use

Use Two Pointers when:

✅ Input is sorted

✅ Need pair/triplet relationships

✅ Need merging of sorted data

✅ Need partitioning

✅ Need in-place modifications

✅ Need left/right traversal

✅ Need sliding window optimization

---

## Recognition Signals

### Strong Indicators

```text
Input array is sorted

Find pair/triplet

Merge arrays

Remove duplicates

Move elements

In-place transformation

Compare elements from both ends
```

---

### Common Keywords

```text
sorted array

merge

pair sum

target

remove duplicates

move zeroes

in-place

constant space
```

---

# Generic Template

## Opposite-End Pointers

```go
left := 0
right := len(nums) - 1

for left < right {

    if condition {
        left++
    } else {
        right--
    }
}
```

---

## Same-Direction Pointers

```go
slow := 0

for fast := 0; fast < len(nums); fast++ {

    if valid(nums[fast]) {
        nums[slow] = nums[fast]
        slow++
    }
}
```

---

## Merge Two Sorted Arrays

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

# Complexity

| Operation          | Complexity         |
| ------------------ | ------------------ |
| Traversal          | O(n)               |
| Space              | O(1)               |
| Sorting Dependency | Sometimes Required |

---

# Common Pitfalls

## 1. Missing Boundary Conditions

```go
for left <= right
```

vs

```go
for left < right
```

Always verify requirements.

---

## 2. Pointer Update Bugs

Incorrect pointer movement often causes:

* Infinite loops
* Skipped elements
* Wrong answers

---

## 3. Ignoring Sorted Property

Many Two Pointer problems become easy once the sorted property is recognized.

---

## 4. Overwriting Data

Common in:

* Merge problems
* Array compression problems

Always verify whether traversal should be:

```text
Left → Right
```

or

```text
Right → Left
```

---

# Related Problems

## Beginner

| #   | Problem                             | Difficulty |
| --- | ----------------------------------- | ---------- |
| 88  | Merge Sorted Array                  | Easy       |
| 27  | Remove Element                      | Easy       |
| 26  | Remove Duplicates from Sorted Array | Easy       |
| 283 | Move Zeroes                         | Easy       |

---

## Intermediate

| #   | Problem                   | Difficulty |
| --- | ------------------------- | ---------- |
| 167 | Two Sum II                | Medium     |
| 11  | Container With Most Water | Medium     |
| 15  | 3Sum                      | Medium     |

---

## Advanced

| #    | Problem                                             | Difficulty |
| ---- | --------------------------------------------------- | ---------- |
| 42   | Trapping Rain Water                                 | Hard       |
| 76   | Minimum Window Substring                            | Hard       |
| 2958 | Length of Longest Subarray With At Most K Frequency | Hard       |

---

# Current Problem Addition

## LeetCode 88 — Merge Sorted Array

### Pattern Usage

**Variant:** Reverse Two Pointers

Unlike the classic left/right pointer approach, this problem uses:

```text
i -> end of valid nums1
j -> end of nums2
k -> end of merge destination
```

### Key Insight

```text
Merge from the end to avoid overwriting values.
```

### Why It Belongs In Two Pointers

Because:

* Multiple pointers track positions simultaneously.
* Each pointer moves independently.
* The solution avoids nested loops.
* The merge completes in linear time.

### Complexity

```text
Time  : O(m+n)
Space : O(1)
```

### Interview Importance

This is one of the foundational Two Pointer problems and is frequently used as an introduction to:

* In-place array manipulation
* Merge Sort concepts
* Sorted data processing
* Space optimization techniques

---

# Pattern Statistics Update

```diff
Two Pointers Pattern

Before:
Problems Covered: T

After:
Problems Covered: T + 1

Added:
+ LeetCode 88 - Merge Sorted Array
```

---

# Pattern Learning Roadmap

```text
Easy
├── 26 Remove Duplicates from Sorted Array
├── 27 Remove Element
├── 88 Merge Sorted Array
├── 283 Move Zeroes

Medium
├── 167 Two Sum II
├── 11 Container With Most Water
├── 15 3Sum

Hard
├── 42 Trapping Rain Water
├── 76 Minimum Window Substring
```
