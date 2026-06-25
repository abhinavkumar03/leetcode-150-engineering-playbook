# Cheat Sheet — 392. Is Subsequence

A one-page revision guide for **LeetCode 392 – Is Subsequence**.

---

# Pattern Summary

| Category | Value |
|----------|-------|
| Pattern | Two Pointers |
| Difficulty | Easy |
| Time Complexity | **O(|s| + |t|)** |
| Space Complexity | **O(1)** |
| Technique | Greedy Sequential Matching |

### Core Idea

Traverse both strings from left to right.

- Compare the current characters.
- If they match, advance both pointers.
- Otherwise, advance only the pointer in `t`.
- If all characters in `s` are matched, `s` is a subsequence of `t`.

---

# Recognition Signals

Consider the **Two Pointer** pattern when you see:

- ✔ Two sequences processed from left to right.
- ✔ Order matters.
- ✔ Characters/elements do **not** need to be adjacent.
- ✔ No revisiting of previous elements.
- ✔ Linear traversal is sufficient.
- ✔ Looking for an ordered match rather than a contiguous match.

### Keywords

- Subsequence
- Relative order
- Sequential matching
- String traversal
- Greedy
- Compare two strings

---

# Algorithm Template

```text
Initialize:

pointerS = 0
pointerT = 0

while pointerS < len(s) AND pointerT < len(t)

    if s[pointerS] == t[pointerT]
        pointerS++

    pointerT++

return pointerS == len(s)
```

---

# Go Template

```go
func isSubsequence(s string, t string) bool {
    sIndex := 0
    tIndex := 0

    for sIndex < len(s) && tIndex < len(t) {
        if s[sIndex] == t[tIndex] {
            sIndex++
        }
        tIndex++
    }

    return sIndex == len(s)
}
```

---

# Pointer Movement

## Match

```text
s : a b c
    ↑

t : a h b g d c
    ↑

Match

↓

Move both pointers
```

---

## No Match

```text
s : a b c
      ↑

t : a h b g d c
      ↑

No Match

↓

Move only pointer in t
```

---

# Greedy Invariant

Always match the **earliest valid occurrence** of each character.

Why?

Skipping an earlier valid match cannot improve the possibility of matching the remaining characters later.

This guarantees correctness while avoiding backtracking.

---

# Complexity Cheatsheet

| Approach | Time | Space | Interview Rating |
|----------|------|-------|------------------|
| Generate all subsequences | O(2ⁿ) | O(n) | ❌ Poor |
| Recursive matching | Exponential (naive) | O(n) | ⚠️ Acceptable only for discussion |
| Two Pointers | **O(|s| + |t|)** | **O(1)** | ✅ Optimal |
| Preprocessing + Binary Search (many queries) | O(|t|) preprocess + O(|s| log |t|) per query | O(|t|) | ✅ Best for repeated queries |

---

# Edge Cases Checklist

| Case | Expected Result |
|------|-----------------|
| `s = ""` | `true` |
| `t = ""`, `s != ""` | `false` |
| Both empty | `true` |
| `s` longer than `t` | `false` |
| Duplicate characters | Preserve order |
| Characters out of order | `false` |
| Single-character strings | Compare directly |

---

# Common Mistakes

### ❌ Treating subsequence as substring

Wrong assumption:

```text
abc
```

must appear contiguously.

Correct:

Characters only need to preserve relative order.

---

### ❌ Advancing both pointers on mismatch

Incorrect:

```text
sIndex++
tIndex++
```

Correct:

```text
tIndex++
```

Advance `sIndex` **only** when characters match.

---

### ❌ Using nested loops

This increases complexity unnecessarily.

A single pass with two pointers is sufficient.

---

### ❌ Forgetting Empty String Cases

Remember:

```text
"" is a subsequence of every string.
```

---

# Follow-up Optimization

If the same string `t` is queried many times:

Instead of scanning `t` repeatedly:

```text
a → [0]
b → [2]
c → [5]
d → [4]
g → [3]
h → [1]
```

Use binary search to find the next valid index for each character in `s`.

### Complexity

| Operation | Complexity |
|----------|------------|
| Preprocess | O(|t|) |
| Each Query | O(|s| log |t|) |
| Space | O(|t|) |

---

# Similar Problems

| Problem | Pattern |
|----------|---------|
| 125. Valid Palindrome | Two Pointers |
| 344. Reverse String | Two Pointers |
| 167. Two Sum II – Input Array Is Sorted | Two Pointers |
| 11. Container With Most Water | Two Pointers |
| 524. Longest Word in Dictionary through Deleting | Two Pointers + Greedy |
| 844. Backspace String Compare | Two Pointers |
| 925. Long Pressed Name | Two Pointers |
| 1768. Merge Strings Alternately | Two Pointers |

---

# Interview Sound Bites

> "A subsequence preserves order but does not require contiguous characters."

> "The two-pointer technique works because each character in both strings is processed at most once."

> "The greedy strategy is safe since matching the earliest valid occurrence never reduces future matching opportunities."

> "For a single query, O(|s| + |t|) with O(1) extra space is optimal."

> "If the target string is reused across many queries, preprocess character indices and use binary search."

---

# Quick Revision Notes

- ✅ Pattern: **Two Pointers**
- ✅ Traverse both strings once.
- ✅ Match → advance both pointers.
- ✅ Mismatch → advance only the pointer in `t`.
- ✅ Stop when either string is exhausted.
- ✅ Return `true` if all characters in `s` have been matched.
- ✅ Time: **O(|s| + |t|)**
- ✅ Space: **O(1)**
- ✅ Follow-up: **Preprocessing + Binary Search** for multiple queries.