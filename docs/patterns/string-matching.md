# String Matching Pattern

## Pattern Definition

String Matching is a technique used to determine whether a pattern string exists within a larger text string.

The goal may be to:

- Find the first occurrence
- Find all occurrences
- Count occurrences
- Verify existence
- Match patterns with constraints

This pattern is one of the most important foundations for:

- Search engines
- Text editors
- Compilers
- DNA sequence analysis
- Log processing systems

---

# Recognition Signals

Look for phrases such as:

```text
find substring
first occurrence
search pattern
contains
match string
locate text
find index
pattern exists
```

Common input structure:

```text
text
pattern
```

or

```text
haystack
needle
```

Questions typically ask:

```text
Return index
Return true/false
Return occurrences
Return count
```

---

# When To Use This Pattern

Use String Matching when:

- One string must be searched inside another
- Characters are compared sequentially
- Pattern existence is important
- Substring positions must be identified

Examples:

```text
Find "abc" inside "xyzabcdef"
```

```text
Search a keyword inside a document
```

```text
Locate a DNA sequence
```

---

# Core Idea

Given:

```text
Text    = haystack
Pattern = needle
```

Try every valid starting position.

At each position:

```text
Compare characters one by one.
```

If all characters match:

```text
Pattern found.
```

Otherwise:

```text
Move to next position.
```

---

# Generic Template

## Brute Force Template

```text
for each valid start position:

    compare pattern characters

    if all characters match:
        return start

return not found
```

---

## Generic Pseudocode

```text
n = length(text)
m = length(pattern)

for start = 0 to n - m:

    matched = true

    for offset = 0 to m - 1:

        if text[start + offset] != pattern[offset]:
            matched = false
            break

    if matched:
        return start

return -1
```

---

# Complexity

## Brute Force

### Time

```text
O((n - m + 1) × m)
```

Often simplified as:

```text
O(n × m)
```

---

### Space

```text
O(1)
```

---

# Optimization Levels

## Level 1

### Brute Force

Characteristics:

- Simple
- Easy to implement
- Interview friendly

Complexity:

```text
O(n × m)
```

---

## Level 2

### KMP (Knuth-Morris-Pratt)

Observation:

```text
Reuse information from previous matches.
```

Complexity:

```text
O(n + m)
```

Additional Space:

```text
O(m)
```

Uses:

```text
LPS Array
(Longest Prefix Suffix)
```

---

## Level 3

### Rabin-Karp

Observation:

```text
Compare hashes instead of characters.
```

Average Complexity:

```text
O(n + m)
```

Used when:

```text
Multiple pattern searches
Rolling hash applications
```

---

## Level 4

### Boyer-Moore

Observation:

```text
Skip large sections after mismatch.
```

Often faster than KMP in practice.

Used in:

```text
Search tools
Text editors
Production search systems
```

---

# Recognition Checklist

Before choosing this pattern ask:

✅ Am I searching one string inside another?

✅ Do I need an index of occurrence?

✅ Is substring comparison required?

✅ Is sequential character matching sufficient?

If yes:

```text
String Matching
```

is likely the correct pattern.

---

# Common Pitfalls

## Pitfall 1

Incorrect boundary.

Wrong:

```text
for start < n
```

Correct:

```text
for start <= n - m
```

---

## Pitfall 2

Not returning first occurrence.

Many problems explicitly require:

```text
first occurrence
```

---

## Pitfall 3

Continuing comparisons after mismatch.

Always:

```text
break immediately
```

---

## Pitfall 4

Ignoring repeated comparisons.

Example:

```text
aaaaaaaaab
aaaaab
```

Brute force performs redundant work.

This motivates:

```text
KMP
```

---

# Interview Strategy

## Junior Level

Expected:

```text
Correct brute-force solution
```

Focus on:

- Correctness
- Boundaries
- Complexity

---

## Mid Level

Expected:

```text
Brute force + optimization discussion
```

Mention:

```text
KMP
```

---

## Senior Level

Expected:

Discussion around:

- KMP
- Rabin-Karp
- Boyer-Moore
- Search systems
- Text indexing

---

# Related Problems

## Easy

### 28. Find the Index of the First Occurrence in a String

Pattern:

```text
Basic substring search
```

Concepts:

- Sequential comparison
- First occurrence detection
- Search boundaries

Complexity:

```text
Time  : O(n × m)
Space : O(1)
```

---

# Pattern Problem Index

| Problem # | Problem Name | Difficulty | Concepts |
|------------|------------|------------|------------|
| 28 | Find the Index of the First Occurrence in a String | Easy | Substring Search, Pattern Matching |

---

# Advanced Problems

### 459. Repeated Substring Pattern

Introduces:

```text
Prefix/Suffix relationships
```

---

### 686. Repeated String Match

Combines:

```text
String construction
Pattern matching
```

---

### 1392. Longest Happy Prefix

Strong KMP preparation.

---

### 214. Shortest Palindrome

Advanced prefix-function usage.

---

### 28 → KMP Learning Path

Recommended progression:

```text
28. Find the Index of the First Occurrence in a String
        ↓
459. Repeated Substring Pattern
        ↓
1392. Longest Happy Prefix
        ↓
KMP Mastery
```

---

# Quick Revision

Remember:

```text
Search pattern inside text.
```

```text
Check all valid start positions.
```

```text
Break on mismatch.
```

```text
Return first complete match.
```

```text
Brute Force → O(n × m)
```

```text
KMP → O(n + m)
```

This pattern forms the foundation of advanced text-search algorithms and frequently appears in technical interviews.