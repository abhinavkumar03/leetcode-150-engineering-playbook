# `interview-notes.md`

# Product of Array Except Self — Interview Notes

## What Interviewer Is Testing

This problem appears simple on the surface, but it evaluates several important engineering and algorithmic skills.

### 1. Pattern Recognition

The interviewer wants to see whether you recognize:

```text
Product Except Self
=
Left Contribution
×
Right Contribution
```

This is a classic Prefix/Suffix pattern.

Candidates who immediately identify this relationship usually reach the optimal solution faster.

---

### 2. Optimization Thinking

Most interviews are less interested in the final answer and more interested in how you improve it.

Expected progression:

```text
Brute Force
↓
Division Solution
↓
Prefix + Suffix Arrays
↓
Space Optimized Prefix/Suffix
```

The interviewer wants to evaluate:

* Can you identify inefficiencies?
* Can you reduce time complexity?
* Can you reduce space complexity?
* Can you justify the tradeoffs?

---

### 3. Constraint Awareness

A common interview trap:

```text
"Why not calculate the total product
and divide by nums[i]?"
```

Strong candidates immediately mention:

* Division is explicitly forbidden.
* Division fails with zeros.
* Precision issues can occur in other variants.

---

### 4. Edge Case Handling

Interviewers often ask:

```text
What happens if the array contains:
- One zero?
- Multiple zeros?
- Negative numbers?
```

The optimal prefix/suffix approach handles all of these naturally.

No special-case logic is required.

---

### 5. Space Complexity Knowledge

Many candidates claim:

```text
Space = O(n)
```

But the interviewer may ask:

```text
Does the output array count?
```

Correct answer:

```text
No.
LeetCode excludes the output array.
```

Therefore:

```text
Extra Space = O(1)
```

---

# Typical Follow-up Questions

## Follow-up 1

### Can you solve it using division?

Yes.

Example:

```text
totalProduct = product of all elements

answer[i] =
totalProduct / nums[i]
```

But:

* Violates constraints
* Fails with zeros
* Usually rejected

---

## Follow-up 2

### Can you reduce extra space?

Initial solution:

```text
left[]
right[]
answer[]
```

Space:

```text
O(n)
```

Optimized:

```text
answer[]
+
suffix variable
```

Space:

```text
O(1)
```

---

## Follow-up 3

### How are zeros handled?

Case 1:

```text
[1,2,0,4]
```

Output:

```text
[0,0,8,0]
```

Case 2:

```text
[1,0,2,0]
```

Output:

```text
[0,0,0,0]
```

Prefix/suffix multiplication naturally produces correct results.

---

## Follow-up 4

### What if numbers are extremely large?

Discussion points:

* Integer overflow
* Using long/long long
* BigInteger for arbitrary precision
* Language-specific limits

LeetCode guarantees valid ranges.

---

## Follow-up 5

### Can you solve it in one traversal?

Not practically.

Reason:

```text
Need information from both sides
of every index.
```

At least two directional accumulations are required.

---

# Optimization Journey

## Stage 1 — Brute Force

### Idea

For each index:

```text
Multiply all remaining elements
```

Example:

```text
for each i:
    for each j:
        if i != j:
            multiply
```

### Complexity

```text
Time:  O(n²)
Space: O(1)
```

### Problem

Repeated work.

---

## Stage 2 — Division

### Idea

```text
answer[i]
=
totalProduct / nums[i]
```

### Complexity

```text
Time:  O(n)
Space: O(1)
```

### Problem

* Division forbidden
* Zero handling breaks

---

## Stage 3 — Prefix + Suffix Arrays

Store:

```text
left[i]
right[i]
```

Then:

```text
answer[i]
=
left[i] × right[i]
```

### Complexity

```text
Time:  O(n)
Space: O(n)
```

---

## Stage 4 — Space Optimized Solution

Store prefix values directly inside:

```text
answer[]
```

Maintain:

```text
suffixProduct
```

while traversing backwards.

### Complexity

```text
Time:  O(n)
Space: O(1)
```

Optimal solution.

---

# Whiteboard Strategy

When solving on a whiteboard:

### Step 1

Write a small example.

```text
nums = [1,2,3,4]
```

Ask:

```text
What should answer[2] be?
```

Answer:

```text
1 × 2 × 4 = 8
```

---

### Step 2

Observe:

```text
Left Side Product
×
Right Side Product
```

---

### Step 3

Draw prefix products.

```text
Index:  0 1 2 3

Nums:   1 2 3 4

Prefix: 1 1 2 6
```

---

### Step 4

Show suffix traversal.

```text
suffix = 1

Move right → left
```

---

### Step 5

Combine both.

Interviewers usually accept the solution immediately after this explanation.

---

# Communication Tips

## Strong Candidate Response

Instead of saying:

```text
I know this problem already.
```

Explain:

```text
For each position,
I need everything on the left
and everything on the right.

That suggests a prefix/suffix decomposition.
```

This demonstrates reasoning.

---

## Explain Space Optimization Clearly

Say:

```text
The output array already exists.

I can reuse it to store prefix products,
then inject suffix products later.
```

This shows engineering maturity.

---

## Mention Constraints Early

Say:

```text
Division is prohibited,
so I'll look for a prefix/suffix approach.
```

Interviewers appreciate constraint-driven thinking.

---

# Senior-Level Discussion Points

Senior engineers are expected to discuss more than just code.

---

## Memory Optimization

Compare:

### Version A

```text
left[]
right[]
answer[]
```

Memory:

```text
3N
```

### Version B

```text
answer[]
suffix
```

Memory:

```text
N + 1
```

This demonstrates practical optimization.

---

## Cache Efficiency

The optimal solution performs:

```text
Sequential array access
```

Benefits:

* Cache friendly
* Predictable memory access
* Better real-world performance

---

## Production Considerations

Questions a senior engineer may raise:

### Overflow

What happens when products exceed integer range?

### Streaming Data

Can this be computed if the entire array is not available?

### Distributed Computation

How would prefix/suffix products be computed across partitions?

These discussions often differentiate senior candidates.

---

# FAANG-Level Variations

## Variation 1

### Product Except Self With Division Allowed

Tests:

```text
Edge-case reasoning
```

instead of algorithmic optimization.

---

## Variation 2

### Product Except Self in a Matrix

For every cell:

```text
Product of all other cells
```

Introduces multidimensional preprocessing.

---

## Variation 3

### Maximum Product Subarray

Related concept:

* Running products
* Sign handling
* Prefix/suffix thinking

Related problem:

Maximum Product Subarray

---

## Variation 4

### Prefix/Suffix Query System

Support:

```text
Query(i)
=
Product Except Self
```

multiple times efficiently.

Leads toward:

* Segment Trees
* Fenwick Trees
* Range Product Queries

---

## Variation 5

### Dynamic Updates

Array values change over time.

Need:

```text
Update(index, value)
Query(index)
```

Discussion naturally evolves into:

* Segment Trees
* Product Trees
* Distributed Aggregation Systems

---

# Interview Summary

### Core Insight

```text
answer[i]
=
(left product)
×
(right product)
```

### Optimal Solution

```text
Forward pass  -> Prefix Products
Backward pass -> Suffix Products
```

### Final Complexity

```text
Time  : O(n)
Space : O(1)
```

### What Makes This Problem Important

This is one of the most common interview questions for evaluating:

* Prefix/Suffix pattern recognition
* Optimization ability
* Space reduction techniques
* Constraint-driven problem solving
* Communication and reasoning skills

A candidate who can derive the O(1) extra-space solution and explain *why it works* is demonstrating the level of thinking expected in strong product-company interviews.

---

**PHASE 5 Complete.**

Reply with **Next** to continue to **PHASE 6 — cheat-sheet.md**.
