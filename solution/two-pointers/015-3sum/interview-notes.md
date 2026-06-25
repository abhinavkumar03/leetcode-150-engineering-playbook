# Interview Notes — 15. 3Sum

This document focuses on the interview perspective of solving the **3Sum** problem. It covers what interviewers evaluate, how to communicate your solution effectively, common follow-up discussions, and advanced variations.

---

# What Interviewer Is Testing

This problem is less about implementing a known algorithm and more about demonstrating structured problem-solving and optimization skills.

## 1. Brute Force Recognition

The interviewer expects you to first identify the straightforward approach.

```text
Choose every combination of three numbers.

for i
    for j
        for k
```

### Expected Complexity

| Time | Space |
|------|-------|
| O(n³) | O(1) |

A good candidate quickly recognizes that this approach is not scalable.

---

## 2. Optimization Ability

The key transition is:

```text
3Sum

↓

Fix one element

↓

Remaining problem becomes

2Sum
```

This observation demonstrates algorithmic thinking rather than memorization.

Interviewers value this reasoning more than simply stating the optimal solution.

---

## 3. Understanding Why Sorting Helps

Sorting is not performed because the problem asks for ordered output.

Sorting is used because it enables deterministic pointer movement.

Without sorting:

- No predictable direction exists.
- Pointer movement becomes meaningless.

After sorting:

```text
Small ------------------> Large
```

Now:

```
Sum too small

↓

Move Left
```

```
Sum too large

↓

Move Right
```

---

## 4. Duplicate Handling

One of the most common interview pitfalls.

Interviewers usually ask:

> "How do you ensure the output contains only unique triplets?"

There are three duplicate scenarios.

### Fixed Element

```
-1 -1 ...
```

Skip duplicate fixed values.

```java
if (i > 0 && nums[i] == nums[i - 1])
    continue;
```

---

### Left Pointer

After finding a valid triplet:

```
0 0 0
```

Move until a different value is found.

---

### Right Pointer

Similarly,

```
2 2
```

Skip duplicate right values.

---

## 5. Complexity Analysis

Candidates should justify complexity rather than memorize it.

Sorting

```
O(n log n)
```

Outer loop

```
O(n)
```

Each pointer scan

```
O(n)
```

Overall

```
O(n²)
```

Since

```
O(n²) > O(n log n)
```

The dominant complexity is

```
O(n²)
```

---

# Typical Follow-up Questions

## Q1. Why not use a HashSet?

Possible answer:

A hash-based solution can solve the inner Two Sum in linear time, but:

- Uses extra memory.
- Duplicate handling becomes more complex.
- Sorting + Two Pointers is cleaner and uses constant auxiliary space.

---

## Q2. Can this be solved without sorting?

Yes.

One approach is:

- Fix one element.
- Use a HashSet for the remaining values.

Trade-offs:

Pros

- No sorting.

Cons

- Extra memory.
- More complicated duplicate removal.
- Less elegant.

---

## Q3. Why move the left pointer when the sum is too small?

Because the array is sorted.

Increasing the left pointer selects a larger value.

This increases the total sum.

---

## Q4. Why move the right pointer when the sum is too large?

The right pointer currently references the largest remaining value.

Moving it left decreases the sum.

---

## Q5. Why do we move both pointers after finding a triplet?

Because:

```
Current triplet already recorded.
```

Keeping either pointer unchanged would revisit the same combination and may generate duplicates.

---

## Q6. What if the target is not zero?

Instead of checking:

```
sum == 0
```

Check

```
sum == target
```

Or equivalently,

```
nums[left] + nums[right] == target - nums[i]
```

The algorithm remains identical.

---

## Q7. Can this idea solve 4Sum?

Yes.

General strategy:

```
Fix first element

↓

Fix second element

↓

Run Two Pointers
```

Complexity

```
O(n³)
```

---

# Optimization Journey

Interviewers like hearing the evolution of your thought process.

## Stage 1

Brute Force

```
Choose every triplet.
```

Complexity

```
O(n³)
```

---

## Stage 2

Observation

```
One value is fixed.

↓

Need Two Sum.
```

---

## Stage 3

Sort Array

Sorting enables ordered traversal.

---

## Stage 4

Apply Two Pointers

```
Left

↓

Increase sum
```

```
Right

↓

Decrease sum
```

---

## Stage 5

Remove Duplicates

Skip duplicates for:

- Fixed element
- Left pointer
- Right pointer

Final Complexity

```
O(n²)
```

---

# Whiteboard Strategy

A structured explanation often scores better than immediately writing code.

## Step 1

Write the brute-force idea.

```
for i

    for j

        for k
```

Discuss complexity.

---

## Step 2

Explain why it is inefficient.

```
Too many repeated searches.
```

---

## Step 3

Draw a sorted array.

```
-4 -1 -1 0 1 2
```

---

## Step 4

Fix one element.

```
-1
```

Remaining target

```
1
```

Now solve:

```
Two Sum
```

---

## Step 5

Explain pointer movement.

```
Too small

↓

Left++
```

```
Too large

↓

Right--
```

---

## Step 6

Discuss duplicate handling before coding.

Interviewers appreciate candidates who proactively address edge cases.

---

# Communication Tips

Strong candidates narrate their reasoning as they code.

### Good

> "I'll sort the array first because it gives me a monotonic ordering, allowing two pointers to move deterministically."

---

### Better

> "After fixing one element, the remaining problem becomes a Two Sum search. Since the array is sorted, I can move the left pointer to increase the sum and the right pointer to decrease it. This reduces the search from cubic to quadratic time."

---

### Best

> "I'll also skip duplicate fixed elements and duplicate pointer values after finding a valid triplet to ensure each combination is returned exactly once."

This demonstrates correctness, optimization, and attention to detail.

---

# Senior-Level Discussion Points

A senior engineer should be able to discuss broader engineering trade-offs beyond the algorithm itself.

## 1. Why Sorting is Acceptable

Sorting changes the input order.

Questions to consider:

- Is mutating the input acceptable?
- Should the function work on a copy instead?

Production code may require preserving the original input:

```go
copyNums := append([]int(nil), nums...)
sort.Ints(copyNums)
```

Trade-off:

- Preserves input.
- Requires O(n) additional space.

---

## 2. Memory vs Simplicity

Two common approaches:

| Approach | Time | Space | Notes |
|----------|------|-------|-------|
| Sorting + Two Pointers | O(n²) | O(1) | Simpler and memory efficient |
| Hash Set | O(n²) | O(n) | Easier inner search but more memory |

Choosing between them depends on system constraints.

---

## 3. Handling Very Large Inputs

For massive datasets:

- The output size may dominate runtime and memory.
- Returning all triplets may be impractical.
- Consider streaming results or limiting output if requirements allow.

---

## 4. Numeric Overflow

In languages with fixed-width integers, adding three large values may overflow.

Example in Java:

```java
long currentSum =
    (long) nums[i] +
    nums[left] +
    nums[right];
```

Using `long` improves robustness when constraints are large.

---

## 5. Immutability and API Design

If this function is part of a shared library:

- Document that it sorts the input in place.
- Or provide an immutable implementation that sorts a copy.

Clear API contracts prevent unexpected side effects.

---

# FAANG-Level Variations

Interviewers often extend the problem after the base solution.

---

## Variation 1 — 3Sum Closest

Return the triplet whose sum is closest to a target.

Key difference:

- Track the minimum absolute difference instead of exact matches.

Complexity:

```
O(n²)
```

---

## Variation 2 — 3Sum Smaller

Count the number of triplets with:

```
sum < target
```

Observation:

When a valid pair is found:

```
left ... right
```

All values between `left` and `right` also satisfy the condition.

Complexity:

```
O(n²)
```

---

## Variation 3 — 4Sum

Generalize by fixing two elements before applying Two Pointers.

Complexity:

```
O(n³)
```

---

## Variation 4 — k-Sum

Recursive approach:

```
k-Sum

↓

Fix one element

↓

(k-1)-Sum

↓

...

↓

2-Sum
```

This demonstrates recursive decomposition and pattern recognition.

---

## Variation 5 — Return Indices Instead of Values

Sorting changes indices.

Possible solutions:

- Store original indices before sorting.
- Use value-index pairs.
- Maintain a mapping from sorted values back to original positions.

This introduces additional implementation complexity.

---

# Interview Summary

## Key Signals of a Strong Candidate

- Starts with the brute-force solution before optimizing.
- Clearly explains why sorting enables the Two Pointer technique.
- Correctly handles duplicates for the fixed element and both pointers.
- Justifies the O(n²) complexity.
- Discusses trade-offs such as input mutation, memory usage, and overflow.
- Connects the problem to the broader k-Sum family.

Mastering **3Sum** demonstrates proficiency with one of the most important interview patterns: **Sorting + Two Pointers**, a technique that appears frequently across coding interviews and real-world optimization problems.