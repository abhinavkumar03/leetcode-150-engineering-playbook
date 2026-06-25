# Dry Run — 392. Is Subsequence

This document provides a detailed walkthrough of the Two Pointer solution for **LeetCode 392 - Is Subsequence**.

---

# Algorithm Overview

We maintain two pointers:

- `sIndex` → Points to the current character in string `s`.
- `tIndex` → Points to the current character in string `t`.

Rules:

- If characters match, move **both pointers**.
- Otherwise, move **only `tIndex`**.
- If `sIndex` reaches the end of `s`, then every character in `s` has been matched in order.

---

# Example 1

## Input

```text
s = "abc"
t = "ahbgdc"
```

## Expected Output

```text
true
```

---

# Initial State

| Variable | Value |
|----------|-------|
| sIndex | 0 |
| tIndex | 0 |
| s | "abc" |
| t | "ahbgdc" |

---

# Visual Walkthrough

```text
s : a  b  c
    ↑

t : a  h  b  g  d  c
    ↑
```

Both pointers start at the first character.

---

# Iteration-by-Iteration State Transition

| Iteration | sIndex | tIndex | s[sIndex] | t[tIndex] | Match | Action | Next State |
|-----------|--------|--------|-----------|-----------|-------|--------|------------|
| 1 | 0 | 0 | a | a | ✅ Yes | Move both pointers | sIndex=1, tIndex=1 |
| 2 | 1 | 1 | b | h | ❌ No | Move `tIndex` | sIndex=1, tIndex=2 |
| 3 | 1 | 2 | b | b | ✅ Yes | Move both pointers | sIndex=2, tIndex=3 |
| 4 | 2 | 3 | c | g | ❌ No | Move `tIndex` | sIndex=2, tIndex=4 |
| 5 | 2 | 4 | c | d | ❌ No | Move `tIndex` | sIndex=2, tIndex=5 |
| 6 | 2 | 5 | c | c | ✅ Yes | Move both pointers | sIndex=3, tIndex=6 |

---

# Pointer Movement Diagram

## Step 1

```text
s : a  b  c
    ↑

t : a  h  b  g  d  c
    ↑

Match ✓
```

Move both.

---

## Step 2

```text
s : a  b  c
       ↑

t : a  h  b  g  d  c
       ↑

No Match ✗
```

Move only `tIndex`.

---

## Step 3

```text
s : a  b  c
       ↑

t : a  h  b  g  d  c
          ↑

Match ✓
```

Move both.

---

## Step 4

```text
s : a  b  c
          ↑

t : a  h  b  g  d  c
             ↑

No Match ✗
```

Move only `tIndex`.

---

## Step 5

```text
s : a  b  c
          ↑

t : a  h  b  g  d  c
                ↑

No Match ✗
```

Move only `tIndex`.

---

## Step 6

```text
s : a  b  c
          ↑

t : a  h  b  g  d  c
                   ↑

Match ✓
```

Move both.

---

# Final State

| Variable | Value |
|----------|-------|
| sIndex | 3 |
| tIndex | 6 |

Since:

```text
sIndex == s.length()
```

Return:

```text
true
```

---

# Example 2

## Input

```text
s = "axc"
t = "ahbgdc"
```

---

# Initial State

```text
s : a  x  c
    ↑

t : a  h  b  g  d  c
    ↑
```

---

# State Transition Table

| Iteration | sIndex | tIndex | s[sIndex] | t[tIndex] | Match | Action |
|-----------|--------|--------|-----------|-----------|--------|--------|
| 1 | 0 | 0 | a | a | ✅ | Move both |
| 2 | 1 | 1 | x | h | ❌ | Move `tIndex` |
| 3 | 1 | 2 | x | b | ❌ | Move `tIndex` |
| 4 | 1 | 3 | x | g | ❌ | Move `tIndex` |
| 5 | 1 | 4 | x | d | ❌ | Move `tIndex` |
| 6 | 1 | 5 | x | c | ❌ | Move `tIndex` |

---

# Final State

| Variable | Value |
|----------|-------|
| sIndex | 1 |
| tIndex | 6 |

Since:

```text
sIndex != s.length()
```

Return:

```text
false
```

---

# Edge Case Walkthroughs

## Case 1 — Empty `s`

### Input

```text
s = ""
t = "abc"
```

Since there are no characters to match:

```text
Return true
```

---

## Case 2 — Empty `t`

### Input

```text
s = "abc"
t = ""
```

No characters are available for comparison:

```text
Return false
```

---

## Case 3 — Both Strings Empty

```text
s = ""
t = ""
```

An empty string is always a subsequence.

```text
Return true
```

---

## Case 4 — Order Violation

### Input

```text
s = "acb"
t = "abc"
```

Traversal:

```text
Match a
↓

Match c

↓

Need b

↓

End of string
```

Return:

```text
false
```

---

# Key Observations

- `sIndex` advances **only when a matching character is found**.
- `tIndex` advances **during every iteration**.
- Each character of `t` is processed at most once.
- No backtracking is required.
- The algorithm greedily matches the earliest valid occurrence of each character.

---

# Complexity Summary

| Metric | Complexity |
|---------|------------|
| Time | **O(|s| + |t|)** |
| Space | **O(1)** |

---

# Takeaways

- This is a classic **Two Pointer** problem.
- Matching must preserve **relative order**, not adjacency.
- A greedy left-to-right scan is sufficient.
- The algorithm is optimal for a single subsequence check.
- For repeated queries against the same `t`, preprocess character indices and use binary search for faster lookups.