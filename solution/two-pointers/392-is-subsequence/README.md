# 392. Is Subsequence

## Problem Statement

Given two strings `s` and `t`, return `true` if `s` is a subsequence of `t`; otherwise, return `false`.

A **subsequence** of a string is formed by deleting zero or more characters without changing the relative order of the remaining characters.

### Examples

#### Example 1

Input

```text
s = "abc"
t = "ahbgdc"
```

Output

```text
true
```

Explanation

```
a h b g d c
↑   ↑     ↑
a   b     c
```

---

#### Example 2

Input

```text
s = "axc"
t = "ahbgdc"
```

Output

```text
false
```

---

## Difficulty

**Easy**

---

## Tags

- String
- Two Pointers
- Greedy

---

## Pattern

**Primary Pattern**

- Two Pointers

**Secondary Pattern**

- Sequential Matching
- Greedy Traversal

---

# Intuition

A subsequence only requires that characters appear **in the same order**.

They do **not** need to be adjacent.

Instead of checking every possible combination, we simply walk through both strings simultaneously.

- One pointer scans `s`.
- Another pointer scans `t`.

Whenever characters match, we move both pointers.

Otherwise, we continue scanning only `t`.

If every character of `s` is matched before `t` ends, then `s` is a subsequence.

---

# Key Observation

The order of characters is the only requirement.

Since every character in `t` is examined at most once, we never need to revisit previously scanned characters.

This naturally leads to a greedy linear solution.

---

# Brute Force Approach

One possible approach is to recursively decide whether to:

- match the current character
- skip the current character

This explores many possibilities and becomes inefficient.

### Algorithm

1. Start from the beginning of both strings.
2. At each position:
   - Match characters if possible.
   - Otherwise skip characters in `t`.
3. Continue until all possibilities are explored.

### Complexity

- Time Complexity: **O(2ⁿ)** (recursive exploration)
- Space Complexity: **O(n)** (recursion stack)

### Limitations

- Exponential growth.
- Unnecessary repeated work.
- Unsuitable for interview settings.
- Poor scalability for long strings.

---

# Optimized Approach

Use the **Two Pointer** technique.

Maintain:

- `i` → current character in `s`
- `j` → current character in `t`

Process `t` only once.

Whenever

```
s[i] == t[j]
```

advance both pointers.

Otherwise,

advance only `j`.

If `i` reaches the end of `s`, every character has been matched.

---

## Algorithm

1. Initialize two pointers.

```
i = 0
j = 0
```

2. While both pointers are within bounds:

- Compare `s[i]` and `t[j]`.

3. If equal:

```
i++
j++
```

4. Otherwise:

```
j++
```

5. Return

```
i == len(s)
```

---

## Why It Works

The algorithm always matches the earliest possible occurrence of each character.

Choosing an earlier occurrence can never reduce future matching opportunities.

This greedy choice guarantees correctness.

Each character of `t` is visited once, producing an optimal linear-time solution.

---

## Complexity

| Metric | Complexity |
|---------|------------|
| Time | **O(n)** |
| Space | **O(1)** |

where

- `n = length of t`

More precisely:

```
O(|s| + |t|)
```

since each pointer moves only forward.

---

# Edge Cases

## Empty `s`

```
s = ""
t = "abc"

Answer = true
```

An empty string is a subsequence of every string.

---

## Empty `t`

```
s = "abc"
t = ""

Answer = false
```

No characters are available for matching.

---

## Both Empty

```
s = ""
t = ""

Answer = true
```

---

## Single Character Match

```
s = "a"
t = "a"

Answer = true
```

---

## Single Character Mismatch

```
s = "a"
t = "b"

Answer = false
```

---

## Duplicate Characters

```
s = "aaa"
t = "aaab"

Answer = true
```

Pointers correctly preserve order.

---

## Order Violated

```
s = "acb"
t = "abc"

Answer = false
```

Characters exist but not in the required order.

---

## Negative Values

Not applicable.

The problem operates exclusively on strings.

---

## Large Inputs

The linear algorithm scales efficiently even when `t` contains hundreds of thousands of characters.

---

# Dry Run

Input

```text
s = "abc"
t = "ahbgdc"
```

| Step | i | j | s[i] | t[j] | Match | Action |
|------|---|---|------|------|-------|--------|
| 1 | 0 | 0 | a | a | Yes | i++, j++ |
| 2 | 1 | 1 | b | h | No | j++ |
| 3 | 1 | 2 | b | b | Yes | i++, j++ |
| 4 | 2 | 3 | c | g | No | j++ |
| 5 | 2 | 4 | c | d | No | j++ |
| 6 | 2 | 5 | c | c | Yes | i++, j++ |

Final

```
i == len(s)

Return true
```

---

# Common Mistakes

### Forgetting that subsequence does not require contiguous characters

Incorrect assumption:

```
abc
```

must appear consecutively.

---

### Moving both pointers on mismatch

Incorrect

```
i++
j++
```

Only the pointer for `t` should move when characters differ.

---

### Ignoring empty string cases

Always remember:

```
"" is a subsequence of every string.
```

---

### Using nested loops

This increases complexity unnecessarily.

The Two Pointer solution already achieves optimal performance.

---

# Interview Discussion

### Why is this a greedy algorithm?

Because we always match the earliest valid occurrence of each character.

---

### Why are Two Pointers sufficient?

The relative order is fixed.

Once a character has been passed in `t`, revisiting it provides no benefit.

---

### Can this be solved recursively?

Yes.

However, recursion explores unnecessary states and is significantly slower than the linear solution.

---

### Can it be optimized further?

The asymptotic complexity cannot be improved beyond linear because every relevant character may need to be examined.

---

### Follow-up

Suppose there are **billions of queries** asking whether different strings `s` are subsequences of the **same** string `t`.

A better approach is:

- preprocess `t`
- store indices of every character
- use Binary Search for each character of `s`

Complexities:

Preprocessing:

```
O(|t|)
```

Each query:

```
O(|s| log |t|)
```

This is the intended follow-up for LeetCode 392.

---

# Follow-up Questions

1. How would you solve this if there were millions of queries?
2. Why is Binary Search useful after preprocessing?
3. Can HashMap + Binary Search improve performance?
4. What changes if characters are Unicode?
5. How would you process a streaming input?

---

# Real World Applications

This pattern appears in many software engineering problems.

### Search Engines

Matching abbreviated search queries.

---

### Version Control Systems

Comparing ordered edit sequences.

---

### Text Editors

Incremental search and fuzzy matching.

---

### DNA Sequence Analysis

Checking whether one gene sequence appears within another while preserving order.

---

### Event Stream Processing

Verifying that expected events occurred in chronological order.

---

# Related Problems

| Problem | Pattern |
|----------|---------|
| 125. Valid Palindrome | Two Pointers |
| 344. Reverse String | Two Pointers |
| 167. Two Sum II | Two Pointers |
| 11. Container With Most Water | Two Pointers |
| 392. Is Subsequence | Two Pointers |
| 524. Longest Word in Dictionary through Deleting | Two Pointers + Greedy |
| 844. Backspace String Compare | Two Pointers |
| 925. Long Pressed Name | Two Pointers |
| 1768. Merge Strings Alternately | Two Pointers |