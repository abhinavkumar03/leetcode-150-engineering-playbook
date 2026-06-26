# 30. Substring with Concatenation of All Words

## Problem Statement

You are given a string `s` and an array of strings `words`, where every word has the same length.

A valid substring is formed by concatenating every word in `words` exactly once, in any order, without any extra characters.

Return all starting indices of such substrings in `s`.

---

## Difficulty

**Hard**

---

## Tags

- Sliding Window
- Hash Map
- String
- Two Pointers
- Frequency Counting

---

## Pattern

**Primary Pattern**

- Sliding Window

**Secondary Patterns**

- Fixed-Length Window
- Hash Map
- String Processing

---

# Intuition

The brute-force idea is straightforward:

For every possible starting index:

1. Extract a substring whose length equals the total length of all words.
2. Split it into equal-sized pieces.
3. Count the frequency of every piece.
4. Compare it with the frequency of the given words.

Although simple, this repeats a tremendous amount of work because neighboring windows overlap significantly.

A better observation is that every word has the **same length**.

Instead of moving one character at a time and rebuilding everything, we can:

- move by one complete word,
- keep a running frequency map,
- expand and shrink the window as needed.

This converts repeated computation into incremental updates.

---

# Key Observation

Suppose:

```
wordLength = 3
```

Every valid word boundary must begin at one of only

```
0
1
2
```

offsets.

Example:

```
barfoothefoobarman

Offset 0

bar foo the foo bar man

Offset 1

arf oot hef oob arm an

Offset 2

rfo oth efo oba rma n
```

Only Offset **0** can ever produce valid words.

Therefore we independently process every possible alignment.

For each alignment:

- move right pointer one word at a time
- maintain current frequencies
- shrink window if a word appears too many times
- whenever window size equals total number of words, record the answer

---

# Brute Force Approach

For every index:

- Take substring of length

```
numberOfWords × wordLength
```

Split into chunks.

Example:

```
barfoo

↓

bar
foo
```

Count frequencies.

Compare with target frequency map.

If equal:

Add index to answer.

### Algorithm

```
For every starting index:

    Extract candidate substring

    Split into equal words

    Build frequency map

    Compare with target map

    If equal
        Save index
```

### Complexity

| Operation | Complexity |
|------------|------------|
| Candidate windows | O(n) |
| Checking each window | O(totalWords) |
| Overall | O(n × totalWords) |

---

### Limitations

- Rebuilds frequency map repeatedly.
- Re-processes overlapping substrings.
- Performs unnecessary comparisons.
- Poor scalability for larger inputs.

---

# Optimized Approach

Instead of rebuilding the entire frequency map every time:

Maintain a sliding window over complete words.

For each possible alignment:

```
0

1

2

...

wordLength-1
```

Maintain:

- left pointer
- right pointer
- current frequency map
- number of matched words

Whenever a frequency exceeds the allowed count:

Move the left pointer forward until the window becomes valid again.

Whenever window size equals total words:

Store the starting index.

Continue sliding.

---

## Algorithm

```
Create target frequency map

For every possible offset

    left = offset
    right = offset

    currentMap = {}

    While right is valid

        Read next word

        If word not present

            Reset window

        Else

            Add word

            While frequency exceeds limit

                Remove left word
                Move left

            If window size equals required size

                Save answer
```

---

## Why It Works

The algorithm guarantees:

- every word enters the window once
- every word leaves the window once

Instead of rebuilding counts repeatedly, frequencies are updated incrementally.

Each alignment is processed independently, ensuring all possible valid concatenations are examined exactly once.

This reduces redundant computation dramatically.

---

## Complexity

Let

```
N = length of string

M = number of words

L = word length
```

### Time Complexity

```
O(N)
```

Every word-sized chunk enters and leaves the sliding window at most once.

---

### Space Complexity

```
O(M)
```

For the frequency hash maps.

---

# Edge Cases

## 1. Empty String

```
s = ""

Output:

[]
```

No substring exists.

---

## 2. Empty Words Array

```
words = []

Output:

[]
```

No concatenation can be formed.

---

## 3. Single Word

```
s = "leetcode"

words = ["leet"]

Output

[0]
```

The algorithm still works.

---

## 4. Duplicate Words

```
words =

["foo","foo","bar"]
```

Frequency counting becomes essential.

Without frequencies, duplicate handling is incorrect.

---

## 5. Word Not Present

```
s = "abcdef"

words =

["xyz"]
```

Entire window resets immediately.

---

## 6. Large Input

Thousands of characters with many repeated words.

Sliding Window avoids rebuilding maps repeatedly and remains efficient.

---

# Dry Run

Example

```
s = "barfoothefoobarman"

words =

["foo","bar"]
```

Target map

| Word | Count |
|------|-------|
| foo | 1 |
| bar | 1 |

Window length

```
2 × 3 = 6
```

### Offset = 0

| Step | Window | Current Map | Action |
|------|---------|------------|--------|
|1|bar|bar=1|Expand|
|2|bar foo|bar=1 foo=1|Valid → Answer = 0|
|3|foo the|Reset|Invalid word|
|4|foo|foo=1|Expand|
|5|foo bar|foo=1 bar=1|Valid → Answer = 9|

Output

```
[0,9]
```

---

# Common Mistakes

### Forgetting Different Offsets

Many solutions only start from index `0`.

Correct solution starts from

```
0
1
2
...
wordLength-1
```

---

### Sliding by One Character

Move the window by

```
wordLength
```

not by one character.

---

### Ignoring Duplicate Words

Using a Set instead of a Frequency Map fails for:

```
["foo","foo","bar"]
```

---

### Not Shrinking the Window

If frequency exceeds the required count:

```
while current[word] > target[word]
```

must shrink from the left.

---

### Incorrect Window Size

The window is valid only when

```
matchedWords == totalWords
```

---

# Interview Discussion

### Expected Progression

Most interviewers expect:

1. Brute Force
2. Frequency Map Optimization
3. Sliding Window
4. Multi-offset Sliding Window

Skipping directly to the optimal solution without explanation often misses an opportunity to demonstrate engineering thinking.

---

### Key Insights to Communicate

- All words have identical length.
- Process string in fixed-size chunks.
- Sliding window prevents repeated counting.
- Frequency maps handle duplicates.
- Multiple offsets guarantee complete coverage.

---

# Follow-up Questions

### What if words have different lengths?

The fixed-size sliding window no longer works.

A more sophisticated string matching approach would be required.

---

### Can this be solved without Hash Maps?

Practically no.

Frequency comparison is fundamental to the problem.

---

### Can the algorithm be parallelized?

Yes.

Each offset is independent and can be processed concurrently.

---

### How would you optimize for extremely large inputs?

- Parallelize offsets.
- Reduce substring allocations.
- Compare using string slices where supported.
- Reuse frequency maps.

---

# Real World Applications

This pattern appears in:

- Log sequence detection
- Event stream matching
- Token sequence analysis
- Malware signature scanning
- DNA sequence matching
- Network packet inspection
- Search engine token processing
- Natural Language Processing (NLP)

---

# Related Problems

| Problem | Pattern |
|----------|----------|
| 3. Longest Substring Without Repeating Characters | Sliding Window |
| 76. Minimum Window Substring | Sliding Window |
| 438. Find All Anagrams in a String | Sliding Window + Hash Map |
| 567. Permutation in String | Sliding Window |
| 209. Minimum Size Subarray Sum | Sliding Window |
| 424. Longest Repeating Character Replacement | Sliding Window |
| 904. Fruit Into Baskets | Sliding Window |
| 1004. Max Consecutive Ones III | Sliding Window |

---

# Key Takeaways

- Equal word lengths enable fixed-size chunk processing.
- Sliding Window eliminates repeated frequency reconstruction.
- Frequency Hash Maps correctly manage duplicate words.
- Processing each possible alignment independently guarantees correctness.
- The optimal solution achieves **O(N)** time with **O(M)** auxiliary space, making it suitable for large inputs.