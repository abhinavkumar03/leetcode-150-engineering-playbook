# interview-notes.md

# Interview Notes

This document focuses on the interview perspective of **LeetCode 167 — Two Sum II: Input Array Is Sorted**. Beyond implementing the algorithm, it explains what interviewers are evaluating, how to communicate your reasoning, and how to handle common follow-up questions.

---

# What Interviewer Is Testing

## 1. Pattern Recognition

The interviewer wants to see whether you recognize that:

* The input array is **already sorted**.
* Sorting is not just additional information—it fundamentally changes the optimal approach.
* A sorted array often suggests:

  * Two Pointers
  * Binary Search
  * Sliding Window (in some problems)

A strong candidate immediately identifies the **Two Pointers** pattern instead of defaulting to a Hash Map.

---

## 2. Optimization Skills

Most candidates can write a brute-force solution:

```text
for every i
    for every j
        if numbers[i] + numbers[j] == target
```

The interviewer is looking for the ability to ask:

> "Can I use the sorted property to avoid unnecessary comparisons?"

Recognizing this optimization is often more valuable than writing the final code.

---

## 3. Algorithmic Reasoning

Writing the code is only part of the evaluation.

The interviewer expects you to justify why pointer movements are correct.

For example:

* Why do we move the left pointer when the sum is too small?
* Why is moving the right pointer the only correct choice when the sum is too large?
* Why are no valid pairs skipped?

A strong explanation demonstrates understanding rather than memorization.

---

## 4. Complexity Analysis

You should confidently explain:

| Approach     | Time     | Space    |
| ------------ | -------- | -------- |
| Brute Force  | O(n²)    | O(1)     |
| Hash Map     | O(n)     | O(n)     |
| Two Pointers | **O(n)** | **O(1)** |

Interviewers appreciate candidates who compare alternatives and justify the chosen solution.

---

## 5. Communication

Good candidates narrate their thought process:

1. Understand the constraints.
2. Consider a simple solution.
3. Identify inefficiencies.
4. Leverage the sorted property.
5. Arrive at the optimal solution.

This structured communication reflects strong engineering thinking.

---

# Typical Follow-up Questions

## Follow-up 1

**Why does the Two Pointer approach work?**

Expected answer:

Because the array is sorted, moving:

* the left pointer increases the sum,
* the right pointer decreases the sum.

This monotonic behavior guarantees that we never discard the correct solution.

---

## Follow-up 2

**What if the array is not sorted?**

Use a Hash Map.

| Method   | Time | Space |
| -------- | ---- | ----- |
| Hash Map | O(n) | O(n)  |

---

## Follow-up 3

**Can you solve it using Binary Search?**

Yes.

For each element:

* Compute the complement.
* Perform binary search on the remaining array.

Complexity:

* Time: O(n log n)
* Space: O(1)

This is correct but less efficient than Two Pointers.

---

## Follow-up 4

**What if multiple answers exist?**

Possible variations include:

* Return the first pair.
* Return all pairs.
* Count all valid pairs.
* Return unique pairs only.

Each variation may require changes to pointer movement or duplicate handling.

---

## Follow-up 5

**What if duplicate values are present?**

The current solution still works because the problem guarantees exactly one solution.

If multiple solutions are allowed, additional logic may be needed to avoid duplicate pairs.

---

## Follow-up 6

**Can this technique be extended?**

Yes.

The Two Pointer strategy is foundational for:

* Three Sum
* Four Sum
* Closest Sum problems
* Pair counting problems
* Sorted array optimizations

---

# Optimization Journey

## Stage 1 — Brute Force

### Idea

Compare every possible pair.

```text
for every pair
    check sum
```

Time Complexity:

```text
O(n²)
```

Advantages:

* Simple
* Easy to implement

Disadvantages:

* Repeats unnecessary comparisons.
* Does not use the sorted property.

---

## Stage 2 — Hash Map

Idea:

Store previously seen values and look for the complement.

Complexity:

```text
Time: O(n)
Space: O(n)
```

Good for unsorted arrays but violates the constant-space requirement.

---

## Stage 3 — Two Pointers (Optimal)

Exploit the sorted input:

* Small sum → move left.
* Large sum → move right.

Complexity:

```text
Time: O(n)
Space: O(1)
```

This satisfies all problem constraints and is the optimal solution.

---

# Whiteboard Strategy

When solving on a whiteboard:

## Step 1

Clarify the problem:

* Is the array sorted?
* Are indices 1-based or 0-based?
* Is exactly one solution guaranteed?
* Can I reuse an element?
* Is extra space allowed?

---

## Step 2

Discuss the brute-force approach briefly.

This shows you can establish a baseline before optimizing.

---

## Step 3

Recognize the key insight:

> "Since the array is sorted, I can use two pointers instead of checking every pair."

---

## Step 4

Draw a simple example:

```text
2   7   11   15
L            R
```

Walk through pointer movements while explaining each decision.

---

## Step 5

Write clean pseudocode before implementation.

This reduces coding mistakes and communicates your approach clearly.

---

## Step 6

Analyze complexity and explain why the algorithm is correct.

---

# Communication Tips

## Before Coding

State your plan:

> "I'll first consider a straightforward solution, then optimize it using the sorted property."

---

## While Coding

Explain pointer updates:

> "The current sum is too small, so I need a larger value. Since the array is sorted, moving the left pointer right is the only move that can increase the sum."

Avoid narrating syntax; focus on reasoning.

---

## After Coding

Verify:

* Example input
* Edge cases
* Complexity
* Constraint compliance

This demonstrates thoroughness and confidence.

---

# Senior-Level Discussion Points

A senior engineer should be able to discuss more than the implementation.

## Correctness Proof

Explain the invariant:

* At every iteration, all discarded pairs are guaranteed not to contain the solution.
* Pointer movements preserve the possibility of finding the valid pair.

---

## Why Sorting Matters

Without sorting, pointer movement has no predictable effect on the sum.

Sorting introduces monotonicity, enabling efficient elimination of impossible pairs.

---

## Space Efficiency

Compare:

| Solution     | Extra Space |
| ------------ | ----------- |
| Hash Map     | O(n)        |
| Two Pointers | O(1)        |

Choosing the latter aligns with the problem's constant-space constraint.

---

## Production Considerations

Discuss practical aspects:

* Deterministic runtime.
* Cache-friendly sequential memory access.
* Minimal memory allocation.
* Easy to test and maintain.

---

## Common Pitfalls

* Returning 0-based indices instead of 1-based.
* Moving the wrong pointer after comparing the sum.
* Forgetting the constant-space requirement.
* Using a Hash Map despite the sorted input.

---

# FAANG-Level Variations

Interviewers often extend this problem to evaluate adaptability.

## Variation 1 — Unsorted Array

Expected Pattern:

* Hash Map

---

## Variation 2 — Return All Valid Pairs

Requires:

* Duplicate handling
* Careful pointer movement
* Result collection

---

## Variation 3 — Three Sum

Approach:

* Sort the array.
* Fix one element.
* Use Two Pointers for the remaining range.

Complexity:

```text
O(n²)
```

---

## Variation 4 — Four Sum

Generalizes the same idea with two nested loops and a Two Pointer search.

---

## Variation 5 — Closest Sum

Instead of stopping at an exact match, track the closest sum encountered.

---

## Variation 6 — Pair Count

Return the total number of pairs whose sum equals the target instead of the indices.

Requires handling duplicates efficiently.

---

# Interview Checklist

Before submitting your solution, verify that you can confidently answer:

* ✅ Why is the Two Pointer technique applicable?
* ✅ Why is moving one pointer sufficient?
* ✅ Why are no valid pairs skipped?
* ✅ Why is the algorithm O(n)?
* ✅ Why is the extra space O(1)?
* ✅ Why does the solution return 1-based indices?
* ✅ How would you solve the unsorted version?
* ✅ How would you adapt the solution for multiple valid pairs?

If you can explain each of these points without referring to code, you have a strong interview-level understanding of this problem.
