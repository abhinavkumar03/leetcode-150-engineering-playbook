# Valid Palindrome (LeetCode 125) — Cheat Sheet

## Pattern Summary

### Primary Pattern

```text
Two Pointers
```

### Secondary Pattern

```text
String Processing
Character Validation
```

### Core Idea

A palindrome is symmetric.

Instead of:

```text
Clean String
Reverse String
Compare
```

Use:

```text
Left Pointer  → Start
Right Pointer → End
```

Compare valid characters while moving inward.

---

# Recognition Signals

When you see requirements like:

### Signal 1

```text
Check if string is palindrome
```

Think:

```text
Two Pointers
```

---

### Signal 2

```text
Compare from both ends
```

Think:

```text
Symmetry
```

---

### Signal 3

```text
Ignore spaces
Ignore punctuation
Ignore symbols
```

Think:

```text
Skip invalid characters
```

---

### Signal 4

```text
Case insensitive comparison
```

Think:

```text
Convert to lowercase
```

---

### Signal 5

```text
Optimize memory usage
```

Think:

```text
Do not create extra strings
```

Use O(1) space solution.

---

# Key Formula

## Palindrome Property

```text
s[i] == s[n - 1 - i]
```

for all valid positions.

---

## Two Pointer Framework

```text
left = 0
right = n - 1

while left < right:
    skip invalid left
    skip invalid right

    compare characters

    left++
    right--
```

---

# Complexity Cheatsheet

| Approach | Time | Space |
|-----------|--------|--------|
| Filter + Reverse | O(n) | O(n) |
| Filter + Compare | O(n) | O(n) |
| Two Pointers | O(n) | O(1) |

---

## Optimal Solution

```text
Time  = O(n)
Space = O(1)
```

---

# Interview Template

## Generic Palindrome Template

```text
Initialize left pointer.

Initialize right pointer.

While left < right:

    Skip invalid left characters.

    Skip invalid right characters.

    Normalize case.

    Compare.

    If mismatch:
        return false

    Move inward.

Return true.
```

---

# Common Pitfalls

## Pitfall 1

Forgetting case conversion.

Wrong:

```text
A != a
```

Correct:

```text
toLower(A) == toLower(a)
```

---

## Pitfall 2

Not skipping punctuation.

Example:

```text
"A,"
```

Comma must be ignored.

---

## Pitfall 3

Ignoring digits.

Example:

```text
"1a1"
```

Digits are valid alphanumeric characters.

---

## Pitfall 4

Using unnecessary extra memory.

Avoid:

```text
Build cleaned string
Build reversed string
```

when O(1) space is possible.

---

## Pitfall 5

Pointer movement mistakes.

Always:

```text
Skip invalid characters first
Compare second
Move pointers third
```

---

# Edge Cases Checklist

Before submitting:

### Empty String

```text
""
```

Expected:

```text
true
```

---

### Only Spaces

```text
" "
```

Expected:

```text
true
```

---

### Only Symbols

```text
"!!!"
```

Expected:

```text
true
```

---

### Single Character

```text
"a"
```

Expected:

```text
true
```

---

### Mixed Case

```text
"Aa"
```

Expected:

```text
true
```

---

### Digits Included

```text
"1a1"
```

Expected:

```text
true
```

---

### Non-Palindrome

```text
"race a car"
```

Expected:

```text
false
```

---

# Similar Problems

## Easy

### LeetCode 344

```text
Reverse String
```

Pattern:

```text
Two Pointers
```

---

### LeetCode 9

```text
Palindrome Number
```

Pattern:

```text
Two Pointers / Math
```

---

### LeetCode 680

```text
Valid Palindrome II
```

Pattern:

```text
Two Pointers + Greedy
```

---

## Medium

### LeetCode 151

```text
Reverse Words in a String
```

Pattern:

```text
Two Pointers
```

---

### LeetCode 5

```text
Longest Palindromic Substring
```

Pattern:

```text
Expand Around Center
```

---

### LeetCode 647

```text
Palindromic Substrings
```

Pattern:

```text
Expand Around Center
```

---

# Pattern Expansion Roadmap

```text
Valid Palindrome
        ↓
Valid Palindrome II
        ↓
Palindrome Number
        ↓
Palindromic Substrings
        ↓
Longest Palindromic Substring
```

Mastering this problem builds intuition for nearly all palindrome interview questions.

---

# One-Minute Revision

### Pattern

```text
Two Pointers
```

---

### Observation

```text
Palindrome = Symmetry
```

---

### Strategy

```text
Start from both ends.

Skip invalid characters.

Convert to lowercase.

Compare.

Move inward.
```

---

### Return False When

```text
Mismatch found.
```

---

### Return True When

```text
Pointers cross successfully.
```

---

### Complexity

```text
Time  = O(n)

Space = O(1)
```

---

### Interview Keyword

```text
Compare mirrored characters
while skipping non-alphanumeric
characters in a single traversal.
```

---

# Memory Hook

Think:

```text
Mirror Test
```

A palindrome behaves like a mirror.

The left side must match the right side after normalization.

```text
Left ←→ Right
```

If every mirrored pair matches:

```text
Valid Palindrome
```