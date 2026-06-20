# Longest Common Prefix — Cheat Sheet

## Pattern Summary

### Pattern Name

**Horizontal Scanning**

### Category

**String Processing**

### Difficulty

**Easy**

### Core Idea

Use the first string as a candidate prefix and continuously shrink it until every string starts with that prefix.

```text
Initial Prefix
      ↓
Compare
      ↓
Mismatch
      ↓
Shrink Prefix
      ↓
Match
      ↓
Continue
```

---

# Recognition Signals

Look for these clues in a problem statement:

✅ Multiple strings

✅ Common beginning characters

✅ Prefix matching

✅ Shared string segment from the start

✅ Need to find the longest matching start sequence

Common keywords:

```text
common prefix
shared beginning
starts with
matching prefix
autocomplete
search suggestion
```

If you see these phrases, think:

```text
String Processing
Prefix Matching
Horizontal Scanning
Trie (advanced alternative)
```

---

# Key Formula

### Horizontal Scanning Template

```text
prefix = first string

for each remaining string:
    while string does not start with prefix:
        remove last character from prefix

return prefix
```

---

### Prefix Shrinking Rule

```text
prefix = prefix[:-1]
```

Keep shrinking until:

```text
currentString.startsWith(prefix)
```

becomes true.

---

# Complexity Cheatsheet

| Approach            | Time           | Space          | Interview Rating |
| ------------------- | -------------- | -------------- | ---------------- |
| Horizontal Scanning | O(N × M)       | O(1)           | ⭐⭐⭐⭐⭐            |
| Vertical Scanning   | O(N × M)       | O(1)           | ⭐⭐⭐⭐             |
| Divide & Conquer    | O(N × M)       | O(log N)       | ⭐⭐⭐              |
| Binary Search       | O(N × M log M) | O(1)           | ⭐⭐⭐              |
| Trie                | O(total chars) | O(total chars) | ⭐⭐               |

Where:

```text
N = number of strings
M = length of common prefix candidate
```

---

# Similar Problems

## Easy

### 14. Longest Common Prefix

Pattern:

```text
Prefix Matching
```

---

## Medium

### 208. Implement Trie (Prefix Tree)

Pattern:

```text
Trie
```

---

### 211. Design Add and Search Words Data Structure

Pattern:

```text
Trie + DFS
```

---

### 1268. Search Suggestions System

Pattern:

```text
Sorting + Trie
```

---

## Hard

### 212. Word Search II

Pattern:

```text
Trie + Backtracking
```

---

# Quick Revision Notes

### Observation

The answer can only:

```text
Stay the same
or
Become shorter
```

It can never become longer.

---

### Invariant

After processing each string:

```text
prefix =
longest common prefix among
all processed strings
```

---

### Early Exit

If:

```text
prefix == ""
```

Return immediately.

No common prefix exists.

---

### Common Mistakes

❌ Forgetting empty array handling

```text
[]
```

---

❌ Shrinking from the wrong side

```text
Wrong:
prefix = prefix[1:]

Correct:
prefix = prefix[:-1]
```

---

❌ Ignoring single-element input

```text
["leetcode"]
```

Answer:

```text
leetcode
```

---

❌ Continuing after prefix becomes empty

```text
No need for further comparisons
```

---

# Interview Answer Framework

### Step 1

State the observation:

> The longest common prefix can only shrink as more strings are processed.

---

### Step 2

Choose the first string as a candidate prefix.

```text
prefix = strs[0]
```

---

### Step 3

Compare against every remaining string.

```text
while current string does not start with prefix:
    shrink prefix
```

---

### Step 4

Return the remaining prefix.

---

# One-Minute Revision

```text
Pattern:
    Horizontal Scanning

Idea:
    Start with first string as prefix

Process:
    Compare with each string
    Shrink on mismatch

Invariant:
    Prefix is valid for all processed strings

Time:
    O(N × M)

Space:
    O(1)

Early Exit:
    Return "" if prefix becomes empty

Alternative Solutions:
    Vertical Scanning
    Divide & Conquer
    Binary Search
    Trie
```

---

# Recruiter / Portfolio Talking Point

This problem demonstrates:

* String manipulation fundamentals
* Pattern recognition
* Edge case handling
* Complexity analysis
* Optimization reasoning
* Ability to compare multiple solution strategies

Although simple, it showcases disciplined problem-solving and clear communication—qualities interviewers consistently value.
