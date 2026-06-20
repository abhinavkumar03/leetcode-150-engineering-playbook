# Horizontal Scanning Pattern

## Pattern Definition

Horizontal Scanning is a technique where a candidate answer is progressively compared against multiple elements in a collection and adjusted until it satisfies all constraints.

The pattern is commonly used when:

* Comparing multiple strings
* Finding common prefixes
* Maintaining a valid candidate solution
* Reducing a search space incrementally

Unlike Vertical Scanning, which processes data column-by-column, Horizontal Scanning processes one element at a time while continuously updating the candidate result.

---

# Recognition Signals

Look for these clues in problem statements:

### String-Based Signals

```text id="7n7bmi"
common prefix
shared beginning
matching start
prefix search
starts with
```

### Collection-Based Signals

```text id="xtdfu5"
compare all elements
find common property
reduce candidate answer
maintain valid result
```

### Interview Clues

When the interviewer asks:

> "Can you start with one answer and keep refining it?"

or

> "Can you progressively reduce the search space?"

Horizontal Scanning is often a good fit.

---

# Core Idea

Start with an initial candidate.

```text id="wn7ulh"
candidate = first element
```

Compare against each remaining element.

If the candidate becomes invalid:

```text id="m09mr0"
shrink
adjust
refine
```

Continue until the candidate works for all elements.

---

# Generic Template

```pseudo
candidate = first_element

for each item in collection:
    while candidate is invalid:
        adjust candidate

return candidate
```

---

# Longest Common Prefix Template

```pseudo
prefix = strs[0]

for each string in strs:
    while string does not start with prefix:
        remove last character from prefix

return prefix
```

---

# Step-by-Step Workflow

```text id="dzho5e"
Initialize Candidate
          │
          ▼
Compare With Next Element
          │
          ▼
Candidate Valid?
     ┌────┴────┐
     │         │
    Yes        No
     │         │
     ▼         ▼
 Continue   Adjust Candidate
     │         │
     └────┬────┘
          ▼
     Process Next
```

---

# Complexity

## Time Complexity

Typical:

```text id="vzt3ut"
O(N × M)
```

Where:

* N = number of elements
* M = size of candidate solution

---

## Space Complexity

Typical:

```text id="5x8w5g"
O(1)
```

Only a few variables are maintained.

---

# Advantages

### Easy to Understand

Very intuitive for interviews.

---

### Easy to Implement

Requires minimal code.

---

### Supports Early Termination

Can stop immediately when:

```text id="sawh88"
candidate becomes empty
candidate becomes impossible
```

---

### Excellent for String Problems

Particularly effective for:

* Prefix matching
* Shared substrings
* Incremental validation

---

# Common Pitfalls

## Pitfall 1

### Forgetting Early Exit

Bad:

```pseudo
continue processing
even when answer is impossible
```

Good:

```pseudo
if candidate is empty:
    return immediately
```

---

## Pitfall 2

### Incorrect Candidate Updates

Example:

```text id="5f3wkg"
Removing characters from wrong side
```

For Longest Common Prefix:

Correct:

```text id="znxhjl"
prefix = prefix[:-1]
```

Incorrect:

```text id="znhq1r"
prefix = prefix[1:]
```

---

## Pitfall 3

### Recomputing From Scratch

Avoid restarting comparisons repeatedly.

Maintain and refine the current candidate.

---

## Pitfall 4

### Ignoring Edge Cases

Examples:

```text id="eh2pje"
[]
["abc"]
["abc","xyz"]
```

Always validate input assumptions.

---

# When NOT to Use Horizontal Scanning

Avoid when:

### Dataset Is Extremely Large

More advanced structures may be better:

* Trie
* Hashing
* Suffix structures

---

### Multiple Queries Are Required

If thousands of prefix lookups are needed:

```text id="s7dtcg"
Trie
Radix Tree
Prefix Index
```

are usually better choices.

---

# Related Problems

## Easy

### 14. Longest Common Prefix

Pattern Usage:

```text id="i7t8rm"
Candidate prefix is continuously reduced
until it matches every string.
```

Complexity:

```text id="y9gzkz"
O(N × M)
```

---

## Medium

### 1268. Search Suggestions System

Uses:

```text id="sl9nqz"
Prefix matching
String processing
```

Often combined with Trie structures.

---

## Medium

### 208. Implement Trie (Prefix Tree)

Natural next step after learning prefix problems.

---

## Hard

### 212. Word Search II

Extends prefix matching concepts into:

```text id="oj02qn"
Trie + Backtracking
```

---

# Comparison With Other Patterns

| Pattern             | Idea                       | Typical Use Case      |
| ------------------- | -------------------------- | --------------------- |
| Horizontal Scanning | Compare row-by-row         | Common prefix         |
| Vertical Scanning   | Compare column-by-column   | Character alignment   |
| Sliding Window      | Expand/Shrink range        | Subarrays, substrings |
| Two Pointers        | Move indices strategically | Sorted arrays         |
| Trie                | Prefix tree structure      | Fast prefix lookup    |

---

# Interview Strategy

When discussing this pattern:

### Mention the Invariant

For Longest Common Prefix:

```text id="b2mx5e"
prefix =
longest common prefix
among processed strings
```

---

### Explain Why It Works

The candidate answer:

```text id="0j9nnk"
only shrinks
never grows
```

Therefore every adjustment moves closer to the final solution.

---

### Highlight Early Exit

```text id="3m4fsh"
if prefix == ""
return ""
```

This demonstrates practical optimization awareness.

---

# Pattern Checklist

Before using Horizontal Scanning, verify:

✅ Candidate solution exists

✅ Candidate can be refined incrementally

✅ Each comparison reduces uncertainty

✅ Early termination is possible

✅ Constant extra space is preferred

---

# Problems Using This Pattern

| LeetCode # | Problem               | Difficulty |
| ---------- | --------------------- | ---------- |
| 14         | Longest Common Prefix | Easy       |

---

# Latest Addition

### LeetCode 14 — Longest Common Prefix

**Folder**

```text id="g5d8jl"
solutions/
└── strings/
    └── 014-longest-common-prefix/
```

**Key Insight**

```text id="u9dr0v"
The common prefix can only shrink,
never expand.
```

**Pattern**

```text id="7m4ov0"
Horizontal Scanning
```

**Complexity**

```text id="q4lylb"
Time: O(N × M)
Space: O(1)
```
