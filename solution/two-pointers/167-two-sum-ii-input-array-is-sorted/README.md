# Two Sum II - Input Array Is Sorted

## Problem Statement

Given a **1-indexed** array of integers `numbers` that is already sorted in **non-decreasing order**, find two numbers such that they add up to a given `target`.

Return the indices of the two numbers (1-indexed) as an array of size two.

Constraints guarantee that:

* Exactly one valid solution exists.
* The same element cannot be used twice.
* The solution must use only constant extra space.

---

## Difficulty

**Easy**

---

## Tags

* Array
* Two Pointers
* Binary Search (Alternative)
* Sorted Array

---

## Pattern

**Primary Pattern:** Two Pointers

**Secondary Pattern:** Exploiting Sorted Arrays

---

# Intuition

A brute-force approach would compare every pair of numbers until the target sum is found.

However, since the array is already sorted, we can do much better.

Imagine placing one pointer at the beginning and another at the end.

* The left pointer starts with the smallest value.
* The right pointer starts with the largest value.

At each step:

* If the current sum is too small, move the left pointer right to increase the sum.
* If the current sum is too large, move the right pointer left to decrease the sum.
* If the sum equals the target, we have found the answer.

Because the array is sorted, every pointer movement is guaranteed to move us closer to the solution without missing any valid pair.

---

# Key Observation

The sorted property provides direction.

Suppose:

```
numbers[left] + numbers[right] < target
```

Increasing the left pointer increases the sum because all elements to the right are greater than or equal to the current one.

Similarly,

```
numbers[left] + numbers[right] > target
```

Decreasing the right pointer decreases the sum because all elements to the left are smaller than or equal to the current one.

Without a sorted array, these decisions would not be valid.

---

# Brute Force Approach

Compare every pair of numbers until a pair equals the target.

### Algorithm

1. Iterate through every element.
2. For each element, iterate through all remaining elements.
3. Compute the pair sum.
4. If the sum equals the target, return the two indices.

### Complexity

| Metric | Complexity |
| ------ | ---------- |
| Time   | O(n²)      |
| Space  | O(1)       |

### Limitations

* Performs many unnecessary comparisons.
* Does not leverage the sorted input.
* Too slow for large arrays.
* Poor scalability.

---

# Optimized Approach

Use two pointers.

* `left` starts at index `0`.
* `right` starts at index `n - 1`.

At each iteration:

* Compute the current sum.
* If the sum equals the target, return the answer.
* If the sum is smaller than the target, increment `left`.
* If the sum is greater than the target, decrement `right`.

Continue until the pair is found.

### Algorithm

```
left = 0
right = n - 1

while left < right

    sum = numbers[left] + numbers[right]

    if sum == target
        return [left + 1, right + 1]

    if sum < target
        left++

    else
        right--
```

### Why It Works

Since the array is sorted:

* Moving the left pointer always increases (or maintains) the left value.
* Moving the right pointer always decreases (or maintains) the right value.

Each move eliminates impossible pairs without revisiting previous combinations.

Every element is processed at most once by each pointer.

This yields a linear-time solution.

### Complexity

| Metric | Complexity |
| ------ | ---------- |
| Time   | **O(n)**   |
| Space  | **O(1)**   |

---

# Edge Cases

| Case                    | Example              | Expected Behavior                                                       |
| ----------------------- | -------------------- | ----------------------------------------------------------------------- |
| Empty input             | `[]`                 | Invalid per constraints, but defensive code may return an empty result. |
| Single element          | `[5]`                | No valid pair exists.                                                   |
| Two elements            | `[2, 7]`, target `9` | Return `[1, 2]`.                                                        |
| Duplicate values        | `[3, 3]`, target `6` | Return both indices.                                                    |
| Negative values         | `[-5, -2, 7, 10]`    | Pointer logic still works.                                              |
| Mixed positive/negative | `[-4, -1, 2, 6]`     | No change in algorithm.                                                 |
| Large inputs            | `100000` elements    | Still runs in linear time.                                              |

---

# Dry Run

### Example

```
numbers = [2, 7, 11, 15]
target = 9
```

| Step | Left | Right | Values | Sum | Action                     |
| ---- | ---- | ----- | ------ | --- | -------------------------- |
| 1    | 0    | 3     | 2 + 15 | 17  | Sum too large → Move right |
| 2    | 0    | 2     | 2 + 11 | 13  | Sum too large → Move right |
| 3    | 0    | 1     | 2 + 7  | 9   | Target found               |

Returned indices:

```
[1, 2]
```

---

### Another Example

```
numbers = [1, 2, 3, 4, 4, 9]
target = 8
```

| Step | Left | Right | Sum | Action     |
| ---- | ---- | ----- | --- | ---------- |
| 1    | 0    | 5     | 10  | Move right |
| 2    | 0    | 4     | 5   | Move left  |
| 3    | 1    | 4     | 6   | Move left  |
| 4    | 2    | 4     | 7   | Move left  |
| 5    | 3    | 4     | 8   | Found      |

---

# Common Mistakes

### Forgetting that the answer is **1-indexed**

Incorrect:

```text
return [left, right]
```

Correct:

```text
return [left + 1, right + 1]
```

---

### Moving both pointers together

Only one pointer should move based on the comparison with the target.

---

### Using a HashMap

While a hash map solves the original Two Sum problem, this problem explicitly benefits from the sorted input and expects a constant-space solution.

---

### Ignoring the sorted property

Using nested loops misses the main optimization and results in O(n²) time.

---

### Incorrect pointer update

```text
sum < target
```

Always move the **left** pointer.

```text
sum > target
```

Always move the **right** pointer.

---

# Interview Discussion

A strong interview answer should explain:

* Why sorting changes the strategy.
* Why each pointer movement is safe.
* Why no valid pair is skipped.
* Why each pointer moves at most `n` times.
* Why the algorithm is O(n) with O(1) extra space.

A common follow-up is asking for the proof that moving one pointer cannot discard the correct solution. Explaining the monotonic nature of the sorted array demonstrates solid reasoning.

---

# Follow-up Questions

1. How would you solve this if the array were **not sorted**?
2. Can you solve it using binary search?
3. What changes if multiple valid pairs exist?
4. How would you return all valid pairs?
5. Can this be generalized to Three Sum or Four Sum?
6. What if the array is sorted in descending order?
7. How would you solve this on a linked list?

---

# Real World Applications

* Financial systems matching transactions that sum to a target amount.
* Inventory systems combining item costs to meet a budget.
* Resource allocation where two capacities must meet a required limit.
* Analytics pipelines searching sorted datasets efficiently.
* Search engines using pointer techniques on ordered indexes.
* Scheduling systems pairing time slots or resources.

---

# Related Problems

| LeetCode | Problem                   | Relation                                      |
| -------- | ------------------------- | --------------------------------------------- |
| 1        | Two Sum                   | Hash Map version for unsorted arrays          |
| 15       | Three Sum                 | Extension of the two-pointer technique        |
| 16       | 3Sum Closest              | Two pointers with optimization                |
| 18       | Four Sum                  | Generalization using sorting and two pointers |
| 1679     | Max Number of K-Sum Pairs | Pair formation using similar reasoning        |
| 633      | Sum of Square Numbers     | Two pointers on a sorted search space         |
| 11       | Container With Most Water | Another classic two-pointer problem           |
