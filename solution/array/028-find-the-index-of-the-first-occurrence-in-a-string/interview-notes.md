# Interview Notes — Find the Index of the First Occurrence in a String

# What Interviewer Is Testing

This problem appears simple, but interviewers use it to evaluate several fundamental skills simultaneously.

---

## 1. String Manipulation Fundamentals

The interviewer wants to verify that you can:

- Traverse strings correctly
- Access characters safely
- Handle indices without off-by-one errors
- Compare substrings efficiently

Common mistakes:

```text
Incorrect loop boundaries
Index out of range
Returning wrong occurrence
```

---

## 2. Brute Force Problem Solving

Many candidates immediately think about built-in functions.

Interviewers want to know:

> Can you implement the search yourself?

Expected thought process:

```text
Try every possible starting position
Compare characters one by one
Return when all characters match
```

---

## 3. Complexity Analysis

Candidates should be able to explain:

```text
n = length of haystack
m = length of needle
```

Worst-case:

```text
O((n − m + 1) × m)
```

Space:

```text
O(1)
```

Interviewers expect this explanation without hesitation.

---

## 4. Recognition of Advanced Algorithms

A strong candidate knows that this problem introduces:

- KMP
- Rabin-Karp
- Boyer-Moore

Even if not implemented.

Senior candidates should mention:

```text
Brute Force → O(n × m)

KMP → O(n + m)
```

---

# Typical Follow-up Questions

## Follow-up 1

### Can you do better than O(n × m)?

Expected discussion:

```text
KMP Algorithm
```

Topics:

- Prefix function
- LPS array
- Pattern preprocessing

Target complexity:

```text
O(n + m)
```

---

## Follow-up 2

### What if we search many patterns inside one text?

Possible discussion:

- Trie
- Aho-Corasick
- Search indexing

---

## Follow-up 3

### What if the pattern contains wildcards?

Example:

```text
ab?d
```

or

```text
ab*d
```

Discussion moves toward:

- Regex engines
- Dynamic Programming
- Backtracking

---

## Follow-up 4

### What if the text is extremely large?

Discussion areas:

- Streaming search
- Chunk processing
- Distributed search systems

---

## Follow-up 5

### How does your language's built-in search work?

Examples:

```text
indexOf()
contains()
strstr()
find()
```

A strong answer:

```text
Depends on implementation.

Many standard libraries use
optimized string matching algorithms
rather than simple brute force.
```

---

# Optimization Journey

Interviewers often want to hear the progression.

---

## Step 1 — Naive Search

Idea:

```text
Try every starting position
```

Complexity:

```text
O(n × m)
```

Pros:

- Easy to write
- Easy to verify

Cons:

- Repeated comparisons

---

## Step 2 — Early Termination

Improvement:

```text
Stop comparison immediately
when mismatch occurs
```

Benefit:

- Better average-case performance

Worst case remains:

```text
O(n × m)
```

---

## Step 3 — KMP

Observation:

```text
Repeated comparisons
contain reusable information
```

Solution:

```text
Store prefix/suffix relationships
```

Complexity:

```text
O(n + m)
```

---

## Step 4 — Specialized Search Algorithms

Examples:

- Boyer-Moore
- Rabin-Karp
- Aho-Corasick

Used in:

- Search engines
- IDEs
- Text processors

---

# Whiteboard Strategy

When solving during interviews:

---

## Step 1

Clarify inputs.

Ask:

```text
Can needle be empty?
```

```text
Can strings contain spaces?
```

```text
Are inputs ASCII or Unicode?
```

---

## Step 2

State brute-force idea first.

Example:

```text
I will check every possible
starting position and compare
characters one by one.
```

Interviewers like hearing a simple correct solution first.

---

## Step 3

Derive loop boundary.

Important observation:

```text
Last valid start = n - m
```

Mention it explicitly.

---

## Step 4

Walk through an example.

Example:

```text
sadbutsad
sad
```

Show:

```text
start = 0
match found
return 0
```

---

## Step 5

Analyze complexity.

Always finish with:

```text
Time: O((n − m + 1) × m)
Space: O(1)
```

---

## Step 6

Mention KMP

Only after presenting the working solution.

Example:

```text
If needed,
I can reduce worst-case complexity
to O(n + m) using KMP.
```

---

# Communication Tips

## Good Interview Communication

Say:

```text
I need to find the first occurrence.
```

```text
I only need to check until n - m.
```

```text
I can stop comparing as soon as a mismatch occurs.
```

```text
Once all characters match,
I immediately return the index.
```

---

## Avoid Saying

```text
I'll just use indexOf().
```

This bypasses the problem.

---

## Explain Your Variables

Instead of:

```text
i
j
```

Explain:

```text
start
offset
```

This improves readability and communication.

---

# Senior-Level Discussion Points

For Staff/Senior interviews, discussion often extends beyond coding.

---

## Text Search at Scale

Examples:

- Search engines
- Log analytics
- Monitoring platforms

Topics:

- Inverted indexes
- Distributed search
- Elasticsearch
- Query optimization

---

## Unicode Considerations

Real-world strings are not always ASCII.

Examples:

```text
English
Hindi
Japanese
Emoji
```

Questions:

```text
What is a character?
What is a byte?
```

Discussion:

- UTF-8
- UTF-16
- Rune handling in Go

---

## Memory vs Performance

Tradeoff:

```text
Brute Force:
O(1) space

KMP:
O(m) space
```

When memory is constrained, the simpler approach may still be preferred.

---

## Library Design

If implementing a production search API:

Consider:

- Case sensitivity
- Locale awareness
- Unicode normalization
- Streaming inputs

---

# FAANG-Level Variations

These are common extensions.

---

## Variation 1

### Implement KMP

Expected:

```text
Build LPS array
Perform efficient matching
```

Complexity:

```text
O(n + m)
```

---

## Variation 2

### Find All Occurrences

Example:

```text
haystack = "aaaaa"
needle = "aa"
```

Output:

```text
[0,1,2,3]
```

---

## Variation 3

### Count Occurrences

Return:

```text
Number of matches
```

instead of first index.

---

## Variation 4

### Multiple Pattern Search

Input:

```text
Patterns:
cat
dog
bird
```

Search all simultaneously.

Discussion:

```text
Aho-Corasick
```

---

## Variation 5

### Wildcard Matching

Examples:

```text
a?c
```

```text
a*c
```

Related problems:

- Regular Expression Matching
- Wildcard Matching

---

## Variation 6

### Search in a Stream

Text arrives continuously:

```text
a
ab
abc
abcd
...
```

Need online matching.

Topics:

- Rolling state
- Streaming KMP
- Finite-state machines

---

# Key Interview Takeaway

The primary goal is not implementing a sophisticated algorithm.

The interviewer wants to see:

1. Correct string traversal
2. Proper boundary handling
3. Clean implementation
4. Complexity awareness
5. Ability to discuss optimizations

A correct brute-force solution with clear reasoning is usually the expected answer.

Strong candidates additionally explain how KMP improves the worst-case runtime from:

```text
O(n × m)
```

to:

```text
O(n + m)
```

and when that optimization becomes valuable.