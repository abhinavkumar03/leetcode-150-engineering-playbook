# Cheat Sheet — Find the Index of the First Occurrence in a String

## Pattern Summary

### Primary Pattern
**String Matching**

### Secondary Pattern
**Two Pointers**

### Difficulty
Easy

### LeetCode
28. Find the Index of the First Occurrence in a String

---

# Recognition Signals

Use the String Matching pattern when the problem asks:

✅ Find a substring inside a string

✅ Search for a pattern in text

✅ Return first occurrence

✅ Return all occurrences

✅ Check if one string exists inside another

✅ Compare characters sequentially

Typical keywords:

```text
find substring
contains
first occurrence
pattern
match
search
index
```

---

# Core Idea

For every valid starting position:

```text
0 → n - m
```

Try matching every character of the pattern.

If all characters match:

```text
return start index
```

Otherwise:

```text
move to next position
```

---

# Visual Template

```text
haystack = sadbutsad
needle   = sad

Position 0

s a d b u t s a d
| | |
s a d

Match Found
Return 0
```

---

# Generic Algorithm

```text
For every start position:

    Compare all characters
    of needle

    If mismatch:
        Stop checking

    If complete match:
        Return start

Return -1
```

---

# Recognition Formula

Given:

```text
n = len(haystack)
m = len(needle)
```

Last valid starting index:

```text
n - m
```

Loop:

```text
for start = 0 to n - m
```

---

# Key Formula

### Character Match Condition

```text
haystack[start + offset]
==
needle[offset]
```

for all:

```text
0 <= offset < m
```

---

# Complexity Cheatsheet

## Brute Force

| Metric | Complexity |
|----------|----------|
| Time | O((n - m + 1) × m) |
| Space | O(1) |

---

## KMP

| Metric | Complexity |
|----------|----------|
| Time | O(n + m) |
| Space | O(m) |

---

## Rabin-Karp

| Metric | Complexity |
|----------|----------|
| Average Time | O(n + m) |
| Worst Time | O(n × m) |
| Space | O(1) / O(m) |

---

## Boyer-Moore

| Metric | Complexity |
|----------|----------|
| Average Time | Sublinear |
| Space | O(m) |

---

# Interview Template

### Step 1

State the idea:

```text
I'll try every valid starting
position and compare characters
one by one.
```

---

### Step 2

Boundary observation:

```text
Last valid start = n - m
```

---

### Step 3

Early exit:

```text
Stop comparison immediately
after mismatch.
```

---

### Step 4

Return:

```text
First complete match index
```

---

### Step 5

Complexity:

```text
Time  : O((n - m + 1) × m)
Space : O(1)
```

---

# Common Pitfalls

## Pitfall 1

Incorrect loop boundary.

Wrong:

```text
start < n
```

Correct:

```text
start <= n - m
```

---

## Pitfall 2

Returning last occurrence.

Requirement:

```text
Return first occurrence.
```

---

## Pitfall 3

Continuing after mismatch.

Bad:

```text
compare remaining characters
```

Good:

```text
break immediately
```

---

## Pitfall 4

Index out of range.

Always ensure:

```text
start + offset < n
```

The loop boundary guarantees this.

---

# Edge Cases Checklist

### Empty Needle

```text
needle = ""
```

Return:

```text
0
```

---

### Exact Match

```text
abc
abc
```

Return:

```text
0
```

---

### No Match

```text
abc
xyz
```

Return:

```text
-1
```

---

### Repeated Characters

```text
aaaaa
aaa
```

Return:

```text
0
```

---

### Needle Longer Than Haystack

```text
abc
abcdef
```

Return:

```text
-1
```

---

# Optimization Journey

```text
Brute Force
      ↓
Early Exit
      ↓
KMP
      ↓
Boyer-Moore
      ↓
Search Engine Scale Algorithms
```

---

# Similar Problems

## Easy

### 14. Longest Common Prefix

Common theme:

```text
character-by-character comparison
```

---

### 58. Length of Last Word

Common theme:

```text
string traversal
```

---

## Medium

### 3. Longest Substring Without Repeating Characters

Pattern:

```text
string processing
sliding window
```

---

### 438. Find All Anagrams in a String

Pattern:

```text
window-based matching
```

---

### 567. Permutation in String

Pattern:

```text
substring matching
frequency tracking
```

---

## Advanced String Matching

### 459. Repeated Substring Pattern

Introduces:

```text
KMP concepts
```

---

### 686. Repeated String Match

Pattern:

```text
substring search
```

---

### 1392. Longest Happy Prefix

Strong KMP practice.

---

# Quick Revision Notes

### Remember

```text
Search pattern inside text.
```

```text
Check every valid start position.
```

```text
Stop comparing on mismatch.
```

```text
Return first complete match.
```

```text
Last valid start = n - m.
```

---

### Complexity

```text
Time  : O((n - m + 1) × m)

Space : O(1)
```

---

### Senior-Level Insight

This problem is the gateway to:

```text
KMP
Rabin-Karp
Boyer-Moore
Aho-Corasick
```

Understanding this brute-force solution thoroughly makes advanced pattern matching algorithms much easier to learn.