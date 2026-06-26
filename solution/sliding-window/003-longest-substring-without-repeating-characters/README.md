# 3. Longest Substring Without Repeating Characters

## Problem Statement

Given a string `s`, find the length of the **longest substring** without repeating characters.

A **substring** is a contiguous sequence of characters within a string.

Return the maximum possible length.

---

## Difficulty

**Medium**

---

## Tags

- String
- Hash Map
- Sliding Window
- Two Pointers

---

## Pattern

**Primary Pattern**

- Sliding Window

**Secondary Pattern**

- Hash Map
- Two Pointers

---

# Intuition

The brute-force approach considers every possible substring and checks whether all characters are unique. This leads to excessive repeated work because many substrings overlap.

Instead, notice that:

- Characters are processed sequentially.
- Once a duplicate appears, only part of the current substring becomes invalid.
- We do **not** need to restart from scratch.

The Sliding Window technique allows us to:

- Expand the window by moving the right pointer.
- Shrink the window only when duplicates appear.
- Maintain a valid substring throughout the traversal.

Since each character enters and leaves the window at most once, the algorithm runs in linear time.

---

# Key Observation

The current window always satisfies:

> **Every character inside the window is unique.**

Whenever adding a new character violates this rule:

- Move the left pointer forward.
- Remove characters until the duplicate disappears.
- Continue expanding.

This guarantees the window remains valid after every iteration.

---

# Brute Force Approach

Generate every possible substring and verify whether it contains duplicate characters.

### Algorithm

1. Start from every index.
2. Extend the substring one character at a time.
3. Use a HashSet to check uniqueness.
4. Stop extending once a duplicate is found.
5. Update the maximum length.
6. Repeat for every starting position.

### Complexity

| Operation | Complexity |
|-----------|------------|
| Time | **O(n²)** (or O(n³) if uniqueness is checked repeatedly) |
| Space | **O(n)** |

### Limitations

- Repeatedly checks overlapping substrings.
- Performs unnecessary duplicate validations.
- Not scalable for large input sizes.
- Inefficient during coding interviews.

---

# Optimized Approach

Maintain a Sliding Window representing the current substring without duplicate characters.

A HashSet stores all characters currently inside the window.

Whenever a duplicate is encountered:

- Remove characters from the left.
- Continue until the duplicate is removed.
- Insert the current character.
- Update the maximum window size.

Each character is added once and removed once.

---

## Algorithm

1. Initialize:
   - left = 0
   - right = 0
   - HashSet
   - answer = 0

2. Iterate through the string using the right pointer.

3. While the current character already exists inside the HashSet:

   - Remove `s[left]`
   - Increment left

4. Insert the current character.

5. Update:

```
answer = max(answer, right - left + 1)
```

6. Continue until the end of the string.

7. Return the maximum length.

---

## Why It Works

The window always satisfies the invariant:

> No character appears more than once.

Whenever the invariant is violated:

- The left pointer moves just enough to restore validity.
- No unnecessary characters are revisited.

Each pointer moves only forward.

Therefore:

- Left pointer moves at most **n** times.
- Right pointer moves exactly **n** times.

Overall complexity becomes linear.

---

## Complexity

| Metric | Complexity |
|---------|------------|
| Time | **O(n)** |
| Space | **O(min(n, charset))** |

For ASCII strings:

- Space is effectively constant.

For Unicode strings:

- Space grows with the number of distinct characters.

---

# Edge Cases

| Case | Example | Expected |
|------|---------|----------|
| Empty string | `""` | `0` |
| Single character | `"a"` | `1` |
| All unique | `"abcdef"` | `6` |
| All duplicates | `"aaaaaa"` | `1` |
| Duplicate in middle | `"abcabcbb"` | `3` |
| Duplicate at beginning | `"aabcdef"` | `6` |
| Duplicate at end | `"abcdefa"` | `6` |
| Alternating duplicates | `"abababab"` | `2` |
| Numbers | `"112233"` | `2` |
| Symbols | `"!@#$%^&"` | `7` |
| Mixed characters | `"aA1!bB"` | `6` |
| Large input | Length = 50,000 | Still O(n) |

---

# Dry Run

Example:

```
Input:

s = "abcabcbb"
```

| Step | Right | Character | Left | Window | Max Length |
|------|-------|-----------|------|--------|------------|
| 1 | 0 | a | 0 | a | 1 |
| 2 | 1 | b | 0 | ab | 2 |
| 3 | 2 | c | 0 | abc | 3 |
| 4 | 3 | a | 1 | bca | 3 |
| 5 | 4 | b | 2 | cab | 3 |
| 6 | 5 | c | 3 | abc | 3 |
| 7 | 6 | b | 5 | cb | 3 |
| 8 | 7 | b | 7 | b | 3 |

Final Answer:

```
3
```

---

# Common Mistakes

### 1. Forgetting to remove characters

Many implementations only move the left pointer but forget to remove the character from the HashSet.

Result:

- Incorrect duplicate detection.
- Window becomes inconsistent.

---

### 2. Updating answer too early

Wrong:

```
Update answer
Then remove duplicates
```

Correct order:

```
Remove duplicates
Insert character
Update answer
```

---

### 3. Resetting the window completely

Some beginners restart from the duplicate position.

This loses useful information and increases complexity.

Sliding Window only removes the minimum required characters.

---

### 4. Off-by-one error

Correct window size:

```
right - left + 1
```

Missing the `+1` is a common mistake.

---

### 5. Using nested loops unnecessarily

The optimized solution requires only one traversal with two pointers.

Nested iterations indicate the Sliding Window optimization has not been fully applied.

---

# Interview Discussion

Interviewers often expect candidates to progress through these stages:

### Stage 1

Describe the brute-force solution.

### Stage 2

Analyze its time complexity.

### Stage 3

Identify overlapping work.

### Stage 4

Recognize Sliding Window as the appropriate optimization.

### Stage 5

Choose a suitable data structure:

- HashSet
- HashMap (last seen index)

### Stage 6

Explain why each pointer moves only forward.

### Stage 7

Derive the O(n) complexity.

Strong candidates clearly explain the invariant maintained by the window and justify the correctness of pointer movements.

---

# Follow-up Questions

1. Can you solve it using a HashMap instead of a HashSet?

2. What changes if the string contains Unicode characters?

3. Can you return the actual longest substring instead of just its length?

4. Can you count all substrings without repeating characters?

5. How would you solve this problem in a streaming scenario where characters arrive one at a time?

6. What if only lowercase English letters are allowed?

7. Can the auxiliary space be reduced further?

---

# Real World Applications

This Sliding Window pattern appears in many practical systems:

### Text Editors

Finding the longest sequence of unique characters for formatting or validation.

---

### Network Packet Analysis

Detecting the longest sequence of non-repeating packet identifiers.

---

### Log Processing

Finding the longest interval without duplicate events or request IDs.

---

### Data Compression

Identifying unique segments before encoding.

---

### Security Systems

Detecting repeated authentication tokens or session identifiers.

---

### Streaming Analytics

Maintaining unique-event windows in real-time data pipelines.

---

### Compiler Design

Scanning token streams while ensuring uniqueness constraints.

---

# Related Problems

| Problem | Pattern | Difficulty |
|----------|----------|------------|
| Two Sum | Hash Map | Easy |
| Minimum Size Subarray Sum | Sliding Window | Medium |
| Minimum Window Substring | Sliding Window | Hard |
| Permutation in String | Sliding Window | Medium |
| Find All Anagrams in a String | Sliding Window | Medium |
| Longest Repeating Character Replacement | Sliding Window | Medium |
| Fruit Into Baskets | Sliding Window | Medium |
| Maximum Erasure Value | Sliding Window | Medium |
| Longest Repeating Substring After Replacement | Sliding Window | Medium |

---

## Key Takeaways

- Sliding Window avoids recomputing overlapping substrings.
- The window always maintains a valid state.
- Hash-based data structures enable constant-time duplicate checks.
- Each character is processed at most twice (once entering and once leaving the window).
- Time complexity improves from quadratic to linear, making the solution suitable for large inputs and technical interviews.