# Length of Last Word

## Problem Statement

Given a string `s` consisting of words and spaces, return the length of the last word in the string.

A word is defined as a maximal substring consisting of non-space characters only.

### Examples

#### Example 1

Input:
s = "Hello World"

Output:
5

Explanation:
The last word is "World", which has length 5.

---

#### Example 2

Input:
s = "   fly me   to   the moon  "

Output:
4

Explanation:
The last word is "moon", which has length 4.

---

#### Example 3

Input:
s = "luffy is still joyboy"

Output:
6

Explanation:
The last word is "joyboy", which has length 6.

---

## Difficulty

Easy

---

## Tags

- String
- Simulation
- Traversal

---

## Pattern

### Primary Pattern
String Traversal

### Secondary Pattern
Reverse Traversal

---

## Intuition

The most direct way to solve this problem is to find the last word and count its characters.

Since trailing spaces may exist, starting from the end of the string is the most efficient approach:

1. Skip all trailing spaces.
2. Count characters until a space is encountered.
3. Return the count.

This avoids splitting the string and uses constant extra memory.

---

## Key Observation

The last word is always located before:

- The end of the string, or
- A sequence of trailing spaces.

Therefore:

- Ignore trailing spaces first.
- Count consecutive non-space characters.

A single reverse traversal is sufficient.

---

## Brute Force Approach

### Idea

Split the string into words and return the length of the last valid word.

### Algorithm

1. Split string using spaces.
2. Remove empty entries.
3. Select last word.
4. Return its length.

### Complexity

Time Complexity: O(n)

Space Complexity: O(n)

### Limitations

- Requires additional memory.
- Creates unnecessary substrings.
- Less efficient for very large inputs.

---

## Optimized Approach

### Algorithm

1. Start from the last index.
2. Skip trailing spaces.
3. Count characters until another space appears.
4. Return the count.

### Why It Works

After skipping trailing spaces, the first sequence of non-space characters encountered from the end must be the last word.

Because every character is visited at most once, the solution is both simple and optimal.

### Complexity

| Metric | Complexity |
|----------|----------|
| Time | O(n) |
| Space | O(1) |

Where `n` is the length of the string.

---

## Edge Cases

### Empty Input

Input:

```text
""
```

Output:

```text
0
```

---

### Only Spaces

Input:

```text
"     "
```

Output:

```text
0
```

---

### Single Word

Input:

```text
"leetcode"
```

Output:

```text
8
```

---

### Multiple Spaces Between Words

Input:

```text
"a   b   c"
```

Output:

```text
1
```

---

### Trailing Spaces

Input:

```text
"Hello World    "
```

Output:

```text
5
```

---

### Duplicates

Input:

```text
"test test"
```

Output:

```text
4
```

Duplicate words do not affect the logic.

---

### Negative Values

Not applicable because the input contains characters and spaces only.

---

### Large Inputs

Input:

```text
Length = 10^4+
```

The reverse traversal solution remains efficient because it performs a single pass with constant extra memory.

---

## Dry Run

### Example

```text
s = "Hello World  "
```

### Step 1: Skip Trailing Spaces

| Index | Character | Action |
|---------|---------|---------|
| 12 | ' ' | Skip |
| 11 | ' ' | Skip |
| 10 | 'd' | Stop |

---

### Step 2: Count Last Word

| Index | Character | Count |
|---------|---------|---------|
| 10 | d | 1 |
| 9 | l | 2 |
| 8 | r | 3 |
| 7 | o | 4 |
| 6 | W | 5 |
| 5 | space | Stop |

Result:

```text
5
```

---

## Common Mistakes

### 1. Forgetting Trailing Spaces

Incorrect:

```text
"Hello World   "
```

Direct counting from end without skipping spaces produces incorrect results.

---

### 2. Using Split Without Filtering Empty Strings

Example:

```text
"Hello World  "
```

Splitting creates empty elements at the end.

---

### 3. Not Handling Single Word Inputs

Example:

```text
"leetcode"
```

The entire string is the last word.

---

### 4. Accessing Invalid Indexes

Carefully check bounds when moving backward through the string.

---

## Interview Discussion

### Expected Solution

Most interviewers expect:

- Reverse traversal
- O(n) time
- O(1) extra space

### What Makes It Good

- Minimal memory usage
- Clean implementation
- Handles all edge cases naturally

### Alternative Solutions

1. Split string
2. Trim and scan
3. Reverse traversal (preferred)

Discussing these alternatives demonstrates engineering maturity.

---

## Follow-up Questions

### 1. Can you solve it without splitting the string?

Yes. Reverse traversal solves it in O(1) space.

---

### 2. Can you process the string in one pass?

Yes. Track current word length while traversing.

---

### 3. What if words are separated by tabs and spaces?

Generalize the delimiter check using whitespace detection.

---

### 4. What if the input is streamed?

Maintain the latest word length while reading characters incrementally.

---

## Real World Applications

### Text Editors

Finding the last typed word.

### Search Engines

Processing query tokens.

### Command Line Parsers

Extracting final arguments.

### Log Processing

Parsing final fields in log entries.

### Natural Language Processing

Token boundary detection.

---

## Related Problems

### Easy

- 344. Reverse String
- 151. Reverse Words in a String
- 125. Valid Palindrome

### Medium

- 49. Group Anagrams
- 3. Longest Substring Without Repeating Characters

### Pattern Progression

```text
58  Length of Last Word
      ↓
151 Reverse Words in a String
      ↓
3   Longest Substring Without Repeating Characters
      ↓
76  Minimum Window Substring
```

These problems gradually increase string-processing complexity.