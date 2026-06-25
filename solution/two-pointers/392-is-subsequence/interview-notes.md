# Interview Notes — 392. Is Subsequence

This document focuses on the interview perspective of solving **LeetCode 392 - Is Subsequence**. It covers what interviewers evaluate, how to communicate your solution effectively, common follow-up discussions, and senior-level insights.

---

# Problem Summary

Given two strings:

- `s` (candidate subsequence)
- `t` (target string)

Determine whether every character in `s` appears in `t` **in the same relative order**.

Characters do **not** need to be contiguous.

---

# What Interviewer Is Testing

Although this is an Easy problem, interviewers use it to evaluate several fundamental skills.

## 1. Pattern Recognition

Can you immediately recognize this as a **Two Pointer** problem?

Candidates who understand common algorithmic patterns usually reach the optimal solution quickly.

---

## 2. Greedy Thinking

Can you justify why matching the **earliest valid occurrence** of each character is always correct?

Expected explanation:

> Once a character in `t` is skipped, revisiting it cannot help satisfy the remaining characters of `s`. Therefore, greedily matching the first valid occurrence never reduces future matching possibilities.

---

## 3. Complexity Awareness

Interviewers expect you to identify that:

- Every character is visited at most once.
- No nested loops are required.
- Constant extra memory is sufficient.

Expected answer:

| Metric | Complexity |
|---------|------------|
| Time | **O(|s| + |t|)** |
| Space | **O(1)** |

---

## 4. Edge Case Handling

Strong candidates proactively discuss edge cases before being asked.

Examples include:

- Empty `s`
- Empty `t`
- Both strings empty
- Single-character strings
- Duplicate characters
- Order violations
- `s` longer than `t`

---

## 5. Code Readability

Interviewers value code that is easy to understand.

Prefer:

```text
sIndex
tIndex
```

instead of:

```text
i
j
```

Descriptive variable names improve maintainability and communication.

---

# Whiteboard Strategy

A structured explanation demonstrates clear thinking.

## Step 1 — Clarify the Problem

Restate the problem in your own words:

> We need to determine whether every character of `s` appears in `t` while preserving order, but not necessarily contiguity.

---

## Step 2 — Identify the Pattern

Explain why Two Pointers fit naturally:

- Both strings are traversed from left to right.
- Characters are never revisited.
- Relative ordering is preserved.

---

## Step 3 — Explain Pointer Movement

Define the pointers:

- `sIndex` tracks the current character in `s`.
- `tIndex` scans through `t`.

Rules:

- If characters match, advance both pointers.
- Otherwise, advance only `tIndex`.

---

## Step 4 — Dry Run

Use a small example:

```text
s = "abc"
t = "ahbgdc"
```

Draw pointer movement step by step.

Interviewers appreciate visual reasoning.

---

## Step 5 — Analyze Complexity

Explain:

- Each pointer moves only forward.
- No character is processed more than once.

Result:

```text
Time  : O(|s| + |t|)
Space : O(1)
```

---

# Typical Follow-up Questions

## Q1. Why is this a greedy algorithm?

Because we always match the earliest possible occurrence of each character.

This choice cannot negatively affect future matches.

---

## Q2. Why not use nested loops?

Nested loops perform unnecessary comparisons.

The Two Pointer approach already guarantees a linear scan.

---

## Q3. Could recursion solve this problem?

Yes.

However:

- Higher memory usage
- Function call overhead
- Exponential exploration in naive implementations

The iterative solution is simpler and more efficient.

---

## Q4. What if `t` is reused for millions of queries?

This is the most common follow-up.

Example:

```text
t = "ahbgdc"
```

Queries:

```text
abc
ac
adc
abdc
...
```

Scanning `t` for every query becomes inefficient.

---

### Better Solution

Preprocess `t`:

```text
a → [0]
b → [2]
c → [5]
d → [4]
g → [3]
h → [1]
```

Then:

- Use Binary Search to locate the next valid occurrence of each character.

Complexity:

| Operation | Complexity |
|-----------|------------|
| Preprocessing | O(|t|) |
| Each Query | O(|s| log |t|) |
| Space | O(|t|) |

---

## Q5. What if characters are streamed?

Maintain the current pointer in `s`.

As each new character arrives:

- Compare it with the current required character.
- Advance the pointer if it matches.

Useful in streaming systems where the full input is unavailable upfront.

---

# Optimization Journey

Interviewers often expect candidates to explain how they refined their solution.

## Stage 1 — Brute Force

Idea:

Generate all possible subsequences of `t` and compare them with `s`.

Complexity:

```text
Time: O(2^|t|)
```

Clearly impractical.

---

## Stage 2 — Recursive Matching

Recursively choose whether to match or skip characters.

Better conceptually but still inefficient due to repeated subproblems.

---

## Stage 3 — Two Pointers (Optimal)

Observation:

- Order matters.
- Characters are never revisited.

Solution:

Traverse both strings once.

Complexity:

```text
Time : O(|s| + |t|)
Space: O(1)
```

---

## Stage 4 — Multiple Queries Optimization

If `t` is fixed:

- Preprocess character positions.
- Use Binary Search.

Ideal for large-scale query workloads.

---

# Communication Tips

During interviews:

✅ Clearly define what each pointer represents.

✅ Explain why only one pointer moves on mismatch.

✅ State the greedy invariant:

> Once a character is skipped in `t`, revisiting it is unnecessary.

✅ Mention complexity before writing code.

✅ Walk through one example while coding.

Avoid jumping directly into implementation.

---

# Common Mistakes

## Moving both pointers on mismatch

Incorrect:

```text
i++
j++
```

Correct:

```text
Only j++
```

---

## Treating subsequences like substrings

A subsequence does **not** require consecutive characters.

---

## Forgetting Empty String Cases

Remember:

```text
"" is always a subsequence.
```

---

## Overcomplicating the Solution

No stacks, queues, hash maps, or dynamic programming are needed for the basic problem.

---

# Senior-Level Discussion Points

Experienced engineers often extend the conversation beyond the basic algorithm.

## Production Considerations

### Single Query

Use the Two Pointer solution.

Simple and efficient.

---

### Batch Processing

If thousands of queries use the same target string:

- Preprocess once.
- Reuse the index map.
- Reduce overall runtime significantly.

---

### Memory vs Performance

Two Pointer:

- Minimal memory.
- Repeated scans.

Preprocessing:

- Higher memory usage.
- Faster repeated lookups.

Choosing between them depends on workload characteristics.

---

### Streaming Systems

The algorithm adapts naturally to streaming input because it processes characters incrementally without backtracking.

---

# FAANG-Level Variations

Interviewers may extend the problem in several directions.

## Variation 1 — Multiple Queries

LeetCode follow-up.

Preprocess `t` and answer many subsequence queries efficiently.

---

## Variation 2 — Unicode Support

Replace fixed-size character arrays with hash maps keyed by Unicode code points.

---

## Variation 3 — Extremely Large Input

When `t` cannot fit into memory:

- Process it as a stream.
- Maintain only the current pointer into `s`.

---

## Variation 4 — Return Matching Indices

Instead of returning `true` or `false`, return the indices in `t` where each character of `s` was matched.

Example:

```text
s = "abc"
t = "ahbgdc"

Output:

[0, 2, 5]
```

---

## Variation 5 — Longest Matching Prefix

If a complete subsequence does not exist, return how many initial characters of `s` were matched before the scan ended.

Useful in search engines and autocomplete systems.

---

# Interview Checklist

Before submitting your solution, verify that you can confidently answer:

- [ ] Why is Two Pointers the correct pattern?
- [ ] Why does the greedy strategy work?
- [ ] Why is the algorithm O(|s| + |t|)?
- [ ] Why is the space complexity O(1)?
- [ ] What happens when `s` is empty?
- [ ] How would you optimize for millions of queries?
- [ ] Can you explain the solution without looking at the code?
- [ ] Can you dry-run the algorithm on paper?

If you can answer all of these comfortably, you're well-prepared for interview discussions around this problem.