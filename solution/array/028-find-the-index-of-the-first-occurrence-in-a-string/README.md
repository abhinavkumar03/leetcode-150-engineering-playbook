# Find the Index of the First Occurrence in a String

## Problem Statement

Given two strings:

- `haystack`
- `needle`

Return the index of the first occurrence of `needle` in `haystack`.

If `needle` is not part of `haystack`, return `-1`.

### Example 1

Input:

```text
haystack = "sadbutsad"
needle = "sad"
```

Output:

```text
0
```

Explanation:

The first occurrence of `"sad"` starts at index `0`.

---

### Example 2

Input:

```text
haystack = "leetcode"
needle = "leeto"
```

Output:

```text
-1
```

Explanation:

The substring does not exist in the given string.

---

## Difficulty

Easy

---

## Tags

- String
- Two Pointers
- String Matching
- Pattern Matching

---

## Pattern

### Primary Pattern

String Matching

### Secondary Pattern

Two Pointers

---

## Intuition

To find a pattern inside a larger string, we can try every possible starting position in the main string.

At each position:

1. Compare characters one by one.
2. Continue while characters match.
3. If all characters of `needle` match, return that starting index.
4. Otherwise move to the next position.

This mimics how a human searches for a word inside a sentence.

---

## Key Observation

If:

```text
haystack length = n
needle length = m
```

Then the last valid starting position is:

```text
n - m
```

Any position beyond that cannot contain the entire pattern.

Therefore we only need to check:

```text
0 → n - m
```

starting indices.

---

# Brute Force Approach

For every possible starting position:

- Compare characters of `needle`
- Stop comparison immediately when mismatch occurs
- If all characters match, return current index

---

## Algorithm

1. Let:

   ```text
   n = len(haystack)
   m = len(needle)
   ```

2. Iterate through every valid starting index.

3. For each index:

   - Compare all characters of `needle`
   - If mismatch occurs, stop checking current position

4. If all characters match:
   - Return current index

5. If loop finishes:
   - Return `-1`

---

## Complexity

### Time Complexity

```text
O((n - m + 1) × m)
```

Worst case:

```text
haystack = "aaaaaaaaaa"
needle   = "aaaab"
```

Many repeated comparisons occur.

---

### Space Complexity

```text
O(1)
```

No extra storage is used.

---

## Limitations

The brute-force approach repeatedly compares the same characters.

Example:

```text
aaaaaaaaab
aaaaab
```

Many comparisons are duplicated.

For very large strings, advanced algorithms like KMP can reduce complexity.

---

# Optimized Approach

For this problem, the accepted optimal interview solution is still the direct string matching approach.

More advanced solutions include:

- KMP Algorithm
- Rabin-Karp
- Boyer-Moore

KMP reduces time complexity to:

```text
O(n + m)
```

but introduces additional preprocessing and complexity.

For LeetCode 28, the straightforward matching solution is preferred unless explicitly asked for optimization.

---

## Algorithm

For each valid starting position:

1. Attempt to match all characters.
2. Exit early on mismatch.
3. Return index immediately on full match.

---

## Why It Works

A substring match exists if and only if:

```text
haystack[i + j] == needle[j]
```

for every:

```text
0 <= j < needle length
```

Checking all valid starting positions guarantees we find the first occurrence.

---

## Complexity

### Time Complexity

```text
O((n - m + 1) × m)
```

### Space Complexity

```text
O(1)
```

---

# Edge Cases

## Empty Input

```text
haystack = ""
needle = ""
```

Depending on language specifications, an empty pattern is typically found at index `0`.

LeetCode constraints guarantee valid inputs.

---

## Single Character Match

```text
haystack = "a"
needle = "a"
```

Output:

```text
0
```

---

## Single Character Mismatch

```text
haystack = "a"
needle = "b"
```

Output:

```text
-1
```

---

## Duplicates

```text
haystack = "aaaaa"
needle = "aaa"
```

Output:

```text
0
```

Must return first occurrence.

---

## Negative Values

Not applicable because input contains strings rather than numeric values.

---

## Large Inputs

```text
haystack length = 10,000+
needle length = 1,000+
```

Brute force remains acceptable but performs repeated comparisons.

KMP becomes more attractive at larger scales.

---

# Dry Run

## Example

```text
haystack = "sadbutsad"
needle = "sad"
```

### Lengths

```text
n = 9
m = 3
```

Valid starting positions:

```text
0 → 6
```

---

### Iteration Table

| Start Index | Compared Characters | Match |
|------------|---------------------|--------|
| 0 | s=s, a=a, d=d | Yes |

All characters matched.

Return:

```text
0
```

---

## Example 2

```text
haystack = "leetcode"
needle = "leeto"
```

### Iteration Table

| Start Index | Result |
|------------|---------|
| 0 | mismatch at character 5 |
| 1 | mismatch |
| 2 | mismatch |
| 3 | mismatch |

No match found.

Return:

```text
-1
```

---

# Common Mistakes

### 1. Iterating Too Far

Wrong:

```text
for i := 0; i < n; i++
```

Correct:

```text
for i := 0; i <= n-m; i++
```

---

### 2. Forgetting Early Exit

Continue comparing after mismatch.

This wastes time.

Stop immediately when mismatch occurs.

---

### 3. Returning Last Match

Problem requires:

```text
first occurrence
```

Return immediately once a complete match is found.

---

### 4. Incorrect Boundary Checks

Accessing:

```text
haystack[i+j]
```

without ensuring indices remain valid.

---

# Interview Discussion

### Expected Solution

Brute-force string matching.

---

### Follow-Up Discussion

Interviewer may ask:

> Can you improve the worst-case time complexity?

Possible answers:

- KMP
- Rabin-Karp
- Boyer-Moore

---

### What Interviewers Evaluate

- String handling
- Loop boundaries
- Complexity reasoning
- Ability to discuss optimizations

---

# Follow-up Questions

### 1. Can you solve it using KMP?

Expected complexity:

```text
O(n + m)
```

---

### 2. What if many searches are performed on the same text?

Consider:

- Suffix Arrays
- Trie
- Search Indexes

---

### 3. What if pattern matching allows wildcards?

Examples:

```text
?
*
```

Leads to regex-style matching problems.

---

### 4. How would you search billions of characters?

Discussion topics:

- Distributed indexing
- Search engines
- Streaming pattern matching

---

# Real World Applications

### Search Feature

Searching text within documents.

---

### IDE Find Function

```text
Ctrl + F
```

String search functionality.

---

### Log Analysis

Finding specific events inside large log files.

---

### DNA Sequence Matching

Searching genetic patterns.

---

### Search Engines

Token and phrase matching.

---

### Intrusion Detection Systems

Pattern matching within network traffic.

---

# Related Problems

### Easy

- 14. Longest Common Prefix
- 58. Length of Last Word

### Medium

- 3. Longest Substring Without Repeating Characters
- 438. Find All Anagrams in a String
- 567. Permutation in String

### Advanced Pattern Matching

- 459. Repeated Substring Pattern
- 686. Repeated String Match
- 214. Shortest Palindrome

### KMP Related

- 459. Repeated Substring Pattern
- 1392. Longest Happy Prefix
