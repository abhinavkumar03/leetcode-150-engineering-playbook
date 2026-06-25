# Interview Notes — Container With Most Water

This document focuses on the interview perspective of **LeetCode 11 — Container With Most Water**. It covers what interviewers evaluate, how to communicate your thought process, optimization strategies, and advanced discussion points expected in senior-level interviews.

---

# What Interviewer Is Testing

This problem is less about implementing two pointers and more about demonstrating algorithmic reasoning.

## 1. Pattern Recognition

The interviewer wants to see whether you can recognize that:

* The input is an array.
* The answer depends on choosing two indices.
* A brute-force solution exists but is inefficient.
* The problem can be optimized using the **Two Pointer** pattern.

---

## 2. Optimization Thinking

Most interviews expect the following progression:

```text
Brute Force
      ↓
Identify Redundant Work
      ↓
Observe Mathematical Property
      ↓
Apply Greedy Decision
      ↓
Two Pointer Optimization
```

Candidates who jump directly to the optimal solution without explaining *why* often receive lower communication scores.

---

## 3. Mathematical Reasoning

The key observation is:

```text
Area = min(leftHeight, rightHeight) × width
```

The interviewer wants to know whether you understand:

* Why the shorter height limits the container.
* Why decreasing the width is unavoidable.
* Why moving the taller pointer cannot improve the answer.

---

## 4. Greedy Proof

A common follow-up question is:

> **Why do we always move the shorter pointer?**

Expected explanation:

* The shorter line determines the current height.
* Moving the taller pointer only reduces the width while keeping the limiting height unchanged or lower.
* Therefore, it cannot produce a larger area.
* Moving the shorter pointer is the only action that can potentially increase the limiting height and compensate for the reduced width.

This is the mathematical justification behind the greedy strategy.

---

## 5. Complexity Analysis

Interviewers expect:

| Metric           | Expected Answer |
| ---------------- | --------------- |
| Time Complexity  | **O(n)**        |
| Space Complexity | **O(1)**        |

Be prepared to explain **why** each pointer moves at most once.

---

# Typical Follow-up Questions

### Q1. Why can't we move both pointers together?

**Answer**

Moving both pointers reduces the width twice as fast and may skip the optimal pair without evaluation.

---

### Q2. Why not move the taller pointer?

Because the shorter height already limits the area.

Reducing the width while keeping the limiting height unchanged cannot increase the area.

---

### Q3. Can Dynamic Programming solve this?

Not effectively.

The problem has:

* No overlapping subproblems.
* No reusable states.
* No optimal substructure suitable for DP.

Two Pointers is the optimal approach.

---

### Q4. Can we return the indices instead of the area?

Yes.

Store the pair `(left, right)` whenever a new maximum area is found.

Return both indices after the traversal.

---

### Q5. What happens if heights can be negative?

The original problem guarantees non-negative heights.

If negatives were allowed, the geometric interpretation would no longer be valid, and additional validation or problem redefinition would be required.

---

### Q6. What if the distance between adjacent lines is not uniform?

The width would no longer be:

```text
right - left
```

Instead, it would need to be computed using the actual x-coordinates of the selected lines.

The greedy reasoning remains similar, but the width calculation changes.

---

# Optimization Journey

A strong interview answer demonstrates how you evolved the solution.

## Step 1 — Brute Force

```text
Try every pair of lines.
```

Complexity:

```text
Time  : O(n²)
Space : O(1)
```

---

## Step 2 — Identify the Bottleneck

The bottleneck is checking every possible pair.

Most comparisons are unnecessary because many pairs cannot outperform the current best once the width shrinks without a taller limiting height.

---

## Step 3 — Key Insight

The shorter line limits the water.

Therefore:

* Keep searching for a taller limiting line.
* Discard the shorter one.
* Never move the taller line first.

---

## Step 4 — Final Solution

Use two pointers.

Each iteration:

1. Compute area.
2. Update answer.
3. Move the shorter pointer.

Linear traversal.

---

# Whiteboard Strategy

When solving on a whiteboard, communicate continuously.

## Step 1

Clarify the problem.

Repeat the area formula:

```text
Area = min(height[left], height[right]) × width
```

---

## Step 2

Discuss the brute-force approach before optimization.

Mention:

* Nested loops.
* O(n²) complexity.
* Inefficiency.

---

## Step 3

Introduce the key observation.

Explain:

> The shorter line limits the area.

---

## Step 4

Describe the greedy decision.

Explain why moving the shorter pointer is safe.

This is often the most important part of the interview.

---

## Step 5

Write the algorithm.

Keep it concise:

```text
while left < right

calculate area

update answer

move shorter pointer
```

---

## Step 6

Walk through an example.

Demonstrate pointer movement and explain each decision.

Interviewers value clear communication as much as correctness.

---

# Communication Tips

A polished explanation can distinguish you from other candidates.

## Start with the brute-force idea

This shows you understand the full search space before optimizing.

---

## Explain the bottleneck

Don't just say:

> "Use two pointers."

Instead, explain why the brute-force approach performs unnecessary work.

---

## Use mathematical language

Say:

* "The limiting factor is the shorter height."
* "The width decreases every iteration."
* "We need a taller limiting height to compensate."

This demonstrates deeper understanding.

---

## Avoid memorized statements

Instead of:

> "Move the smaller pointer."

Say:

> "The shorter line constrains the current area. Moving the taller line cannot increase the limiting height, so the only promising move is advancing the shorter pointer."

---

## Analyze complexity after coding

A complete answer should include:

* Time complexity.
* Space complexity.
* Justification.

---

# Senior-Level Discussion Points

Senior engineers are expected to discuss trade-offs and proofs, not just implementation.

Topics worth mentioning:

* Why the greedy choice is safe.
* Loop invariant maintained during traversal.
* Why every discarded state is provably non-optimal.
* Why the algorithm is optimal with respect to time complexity.
* Readability versus micro-optimizations.
* Extending the solution to return additional information (indices, coordinates, metadata).

---

# FAANG-Level Variations

Interviewers at larger companies may extend the problem with variations.

## Variation 1 — Return Indices

Instead of only the maximum area, return the pair of indices forming the optimal container.

---

## Variation 2 — Non-Uniform X Coordinates

Each vertical line has a custom x-coordinate.

Update the width calculation accordingly while preserving the greedy logic.

---

## Variation 3 — Streaming Heights

Heights arrive one at a time.

Discuss whether an exact online solution is possible and what trade-offs approximate solutions might require.

---

## Variation 4 — Top K Containers

Return the top **K** largest distinct container areas.

Requires additional data structures and changes to the search strategy.

---

## Variation 5 — Dynamic Updates

Heights can change after queries.

Discuss how segment trees, balanced trees, or other advanced structures might support efficient updates, noting that the original greedy algorithm no longer directly applies.

---

# Interview Checklist

Before finishing your interview solution, verify that you can confidently explain:

* ✅ Why brute force is O(n²).
* ✅ Why the shorter line limits the area.
* ✅ Why moving the taller pointer is not beneficial.
* ✅ Why moving the shorter pointer is the greedy choice.
* ✅ Why each pointer moves at most `n` times.
* ✅ Why the final algorithm runs in O(n) time and O(1) space.
* ✅ How to adapt the solution to return indices instead of only the maximum area.
