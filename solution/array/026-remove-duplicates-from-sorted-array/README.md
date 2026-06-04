# 26. Remove Duplicates from Sorted Array

---

## Problem Statement

Given an integer array `nums` sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.

The relative order of the elements should remain the same.

Since it is not possible to change the length of the array in some languages, you must place the unique elements in the first part of the array and return the number of unique elements `k`.

After removing duplicates:

* The first `k` elements should contain the unique values.
* The remaining elements beyond `k` do not matter.

### Example 1

```text
Input:
nums = [1,1,2]

Output:
2

Modified Array:
[1,2,_]
```

### Example 2

```text
Input:
nums = [0,0,1,1,1,2,2,3,3,4]

Output:
5

Modified Array:
[0,1,2,3,4,_,_,_,_,_]
```

---

## Difficulty

**Easy**

---

## Tags

* Array
* Two Pointers
* In-Place Modification
* Fast & Slow Pointer

---

## Pattern

### Primary Pattern

Two Pointers

### Secondary Pattern

Fast Pointer + Slow Pointer

---

## Intuition

The array is already sorted.

This means:

```text
Duplicate values always appear next to each other.
```

Example:

```text
[1,1,2,2,2,3,4,4]
```

We do not need a hash set or extra storage.

Instead:

* One pointer scans the array.
* Another pointer tracks where the next unique value should be placed.

Whenever we find a new unique value, we write it into the next available position.

---

## Key Observation

Because the array is sorted:

```text
nums[i] != nums[i - 1]
```

indicates the start of a new unique element.

Therefore:

* Keep the first occurrence.
* Skip all consecutive duplicates.
* Compact unique values toward the front.

This allows:

```text
O(n) Time
O(1) Extra Space
```

which is optimal.

---

## Brute Force Approach

### Idea

Use an additional data structure such as a set to store unique values.

### Algorithm

1. Create an empty set.
2. Traverse the array.
3. Insert unseen values into the set.
4. Copy unique values back into the array.
5. Return the count of unique elements.

### Complexity

| Metric | Value |
| ------ | ----- |
| Time   | O(n)  |
| Space  | O(n)  |

### Limitations

* Uses extra memory.
* Violates the problem requirement of in-place modification.
* Not the optimal interview solution.

---

## Optimized Approach

### Algorithm

1. Handle edge case for empty array.
2. Initialize:

```text
slow = 0
```

3. Traverse array using:

```text
fast = 1 → n-1
```

4. Whenever:

```text
nums[fast] != nums[slow]
```

a new unique value is found.

5. Move `slow` forward and place the unique value:

```text
slow++
nums[slow] = nums[fast]
```

6. Return:

```text
slow + 1
```

as the count of unique elements.

---

### Why It Works

The region:

```text
nums[0...slow]
```

always contains unique values.

The fast pointer explores the remaining elements.

Whenever a new unique value is found:

* Extend the unique region.
* Copy the value into the next position.

By the end:

```text
nums[0...slow]
```

contains exactly one occurrence of each distinct number.

---

### Complexity

| Metric | Value |
| ------ | ----- |
| Time   | O(n)  |
| Space  | O(1)  |

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

Array remains unchanged.

---

### All Duplicates

```text
[2,2,2,2]
```

Output:

```text
1
```

---

### No Duplicates

```text
[1,2,3,4]
```

Output:

```text
4
```

---

### Negative Values

```text
[-3,-3,-2,-1,-1]
```

Output:

```text
3
```

Modified Array:

```text
[-3,-2,-1]
```

---

### Large Inputs

```text
100,000+ elements
```

Still performs efficiently because:

```text
O(n) time
O(1) space
```

---

## Dry Run

### Example

```text
nums = [0,0,1,1,1,2,2,3,3,4]
```

Initial:

```text
slow = 0
```

| fast | nums[fast] | nums[slow] | Action       | Array State           |
| ---- | ---------- | ---------- | ------------ | --------------------- |
| 1    | 0          | 0          | Duplicate    | [0,0,1,1,1,2,2,3,3,4] |
| 2    | 1          | 0          | Unique found | [0,1,1,1,1,2,2,3,3,4] |
| 3    | 1          | 1          | Duplicate    | unchanged             |
| 4    | 1          | 1          | Duplicate    | unchanged             |
| 5    | 2          | 1          | Unique found | [0,1,2,1,1,2,2,3,3,4] |
| 6    | 2          | 2          | Duplicate    | unchanged             |
| 7    | 3          | 2          | Unique found | [0,1,2,3,1,2,2,3,3,4] |
| 8    | 3          | 3          | Duplicate    | unchanged             |
| 9    | 4          | 3          | Unique found | [0,1,2,3,4,2,2,3,3,4] |

Final:

```text
k = 5

Valid Portion:
[0,1,2,3,4]
```

---

## Common Mistakes

### 1. Using Extra Space

```text
HashSet
ArrayList
Map
```

Works but violates in-place requirement.

---

### 2. Returning Slow Instead of Slow + 1

Incorrect:

```java
return slow;
```

Correct:

```java
return slow + 1;
```

because indices start from 0.

---

### 3. Starting Fast Pointer at 0

Incorrect:

```java
fast = 0
```

The first element is already unique.

Start from:

```java
fast = 1
```

---

### 4. Forgetting Empty Array Check

Without it:

```java
nums[0]
```

causes an index error for empty input.

---

### 5. Modifying Relative Order

Interviewers expect:

```text
Original order preserved.
```

Do not sort again.

---

## Interview Discussion

This problem is often used as an introduction to:

* Two Pointer Pattern
* In-Place Array Modification
* Array Compression
* Fast/Slow Pointer Technique

A strong candidate should recognize:

```text
Sorted Array
+
Need O(1) Space
=
Two Pointers
```

within the first few minutes.

---

## Follow-up Questions

### 1. What if the array is not sorted?

Possible answer:

Use a hash set.

Complexity:

```text
Time: O(n)
Space: O(n)
```

---

### 2. What if each element can appear at most twice?

Leads to:

Remove Duplicates from Sorted Array II

---

### 3. Can we solve it recursively?

Yes, but iterative is simpler and avoids stack overhead.

---

### 4. Why is O(1) space important?

Large datasets may not fit into additional memory structures.

---

### 5. What happens if the array is extremely large?

The algorithm remains efficient because:

```text
Single pass
Constant extra space
```

---

## Real World Applications

### Data Cleaning

Removing duplicate records from sorted datasets.

---

### Log Processing

Compressing repeated log entries.

---

### Database Systems

Creating unique ordered indexes.

---

### Analytics Pipelines

Deduplicating sorted event streams.

---

### Search Engines

Removing repeated sorted document identifiers.

---

## Related Problems

| Problem                                | Pattern             |
| -------------------------------------- | ------------------- |
| Remove Duplicates from Sorted Array II | Two Pointers        |
| Remove Element                         | Fast & Slow Pointer |
| Move Zeroes                            | Array Compaction    |
| Merge Sorted Array                     | Sorted Arrays       |
| Squares of a Sorted Array              | Two Pointers        |
