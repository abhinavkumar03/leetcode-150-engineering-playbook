# Interview Notes

## Problem

**30. Substring with Concatenation of All Words**

**Difficulty:** Hard

**Primary Pattern:** Sliding Window

---

# What Interviewer Is Testing

This problem evaluates whether you can recognize a **fixed-size sliding window** hidden inside a string processing problem. Rather than checking every possible substring independently, the interviewer wants to see if you can exploit the fact that **all words have the same length** to process the string efficiently.

---

## 1. Pattern Recognition

The interviewer expects you to notice:

- Every word has identical length.
- The window size is fixed:
  ```
  totalWindowLength = wordLength × numberOfWords
  ```
- Words should be processed as **chunks**, not as individual characters.
- A frequency map is required to handle duplicate words correctly.

A candidate who identifies these observations early is more likely to arrive at the optimal solution.

---

## 2. Sliding Window Expertise

The interviewer is looking for your ability to:

- Expand the window one word at a time.
- Track word frequencies dynamically.
- Shrink the window when a word appears too many times.
- Maintain a valid window without rebuilding frequency maps.

This demonstrates mastery of advanced sliding window techniques.

---

## 3. Hash Map Usage

The solution relies on two hash maps:

### Target Frequency Map

Stores the expected frequency of each word.

Example:

```text
words = ["foo", "bar", "foo"]

Target Map

foo → 2
bar → 1
```

### Current Window Map

Tracks the frequencies of words currently inside the sliding window.

The interviewer wants to ensure you can compare and update these maps efficiently.

---

## 4. Window Alignment

A common mistake is sliding character by character.

The interviewer expects you to recognize that valid words can only begin at specific offsets.

For a word length of `3`, the valid alignments are:

```text
Offset 0
Offset 1
Offset 2
```

Each alignment must be processed independently.

This is one of the key insights that separates the optimal solution from less efficient approaches.

---

## 5. Complexity Analysis

You should be able to explain why the algorithm runs in linear time.

Each word-sized chunk:

- Enters the window at most once.
- Leaves the window at most once.

Therefore:

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(m)`, where `m` is the number of unique words.

Avoid stating `O(n × wordLength)` unless you clarify that `wordLength` is treated as a constant.

---

# Typical Follow-up Questions

### Q1. Why do we need multiple offsets?

**Answer:**

Because every word has the same length, valid word boundaries only occur at positions with the same remainder modulo `wordLength`. Processing each offset ensures that every possible concatenation is examined exactly once.

---

### Q2. Why can't we move one character at a time?

**Answer:**

Moving one character at a time would split words incorrectly, forcing unnecessary substring reconstruction and additional comparisons. Advancing by one full word preserves alignment and keeps the algorithm efficient.

---

### Q3. How are duplicate words handled?

**Answer:**

A frequency map stores the required count of each word. If the current window contains a word more times than allowed, the window is shrunk from the left until the frequency becomes valid again.

---

### Q4. What happens when an invalid word is encountered?

**Answer:**

The current window cannot be part of a valid answer. Reset the window:

- Clear the current frequency map.
- Reset the word count.
- Move the left pointer to the position immediately after the invalid word.

---

### Q5. Can this solution work if words have different lengths?

**Answer:**

No.

The algorithm depends on equal-length words to define fixed-size chunks and predictable window movement. Variable-length words would require a different strategy, such as backtracking, tries, or more advanced string matching techniques.

---

# Optimization Journey

Interviewers often expect candidates to improve the solution step by step rather than jumping directly to the optimal approach.

## Stage 1 — Brute Force

### Idea

For every starting index:

1. Extract a substring with the required total length.
2. Split it into equal-sized words.
3. Count the frequencies.
4. Compare with the target map.

### Complexity

- **Time:** `O(n × m)`
- **Space:** `O(m)`

### Drawback

The same words are counted repeatedly in overlapping windows.

---

## Stage 2 — Frequency Map Optimization

Reuse the target frequency map while still rebuilding the current map for every candidate window.

This reduces some repeated work but still performs unnecessary processing.

---

## Stage 3 — Sliding Window

Instead of rebuilding the current map:

- Expand the window one word at a time.
- Update frequencies incrementally.
- Shrink only when constraints are violated.

This eliminates redundant work.

---

## Stage 4 — Multi-Offset Sliding Window

Run the sliding window separately for each possible offset:

```text
0
1
2
...
wordLength - 1
```

This ensures complete coverage of all valid word alignments while maintaining linear time complexity.

---

# Whiteboard Strategy

A structured explanation helps interviewers follow your reasoning.

## Step 1

Write the given information:

```text
wordLength = len(words[0])
wordCount = len(words)

windowLength = wordLength × wordCount
```

---

## Step 2

Draw the target frequency map.

Example:

```text
foo → 2
bar → 1
```

---

## Step 3

Illustrate the possible offsets:

```text
Offset 0

bar | foo | the | foo | bar

Offset 1

arf | oot | hef | oob | arm

Offset 2

rfo | oth | efo | oba | rma
```

Explain why only aligned chunks can form valid words.

---

## Step 4

Draw the sliding window:

```text
Left                Right

bar | foo | the | foo | bar
^^^^^^
```

Show how the window expands and shrinks while maintaining valid frequencies.

---

## Step 5

Explain the shrinking condition:

```text
while current[word] > target[word]
```

Remove words from the left until the window becomes valid again.

---

## Step 6

State the completion condition:

```text
wordsInWindow == totalWords
```

Record the current left index as a valid answer.

---

# Communication Tips

During the interview:

- Start with the brute-force solution before introducing optimizations.
- Emphasize that all words have equal length.
- Explain why processing by fixed-size chunks is more efficient than character-by-character scanning.
- Clearly describe the roles of the target and current frequency maps.
- Mention that each alignment is processed independently.
- Conclude with a concise complexity analysis.

A clear progression demonstrates strong problem-solving skills and engineering thought process.

---

# Senior-Level Discussion Points

Experienced engineers often discuss considerations beyond algorithm correctness.

## 1. Avoiding Substring Allocations

Repeated substring creation can increase memory usage.

Possible optimizations include:

- Using string slices where supported.
- Reusing buffers.
- Comparing indices directly when feasible.

---

## 2. Parallel Processing

Each offset is independent.

For very large inputs, different offsets can be processed concurrently to utilize multiple CPU cores.

---

## 3. Streaming Data

If the input arrives as a stream:

- Maintain a rolling window of word-sized chunks.
- Update frequency maps incrementally.
- Emit valid indices as they are discovered.

This adapts the solution to real-time text analysis.

---

## 4. Memory Considerations

Reuse data structures between offsets instead of allocating new maps repeatedly when performance is critical.

---

## 5. Production Applications

The same approach appears in:

- Log sequence detection.
- Tokenized search engines.
- Event stream processing.
- Malware signature matching.
- DNA sequence analysis.

Understanding these applications demonstrates practical engineering awareness.

---

# FAANG-Level Variations

Interviewers may extend the problem with additional constraints.

### Variation 1 — Variable-Length Words

Words no longer have equal lengths.

**Challenge:** Fixed-size chunking is impossible.

**Possible approaches:**

- Trie-based matching.
- Dynamic programming.
- Backtracking with memoization.

---

### Variation 2 — Return the Matching Substrings

Instead of indices, return the actual concatenated substrings.

The sliding window remains unchanged; only the output format differs.

---

### Variation 3 — Allow Extra Characters Between Words

Words may appear in order but are separated by arbitrary characters.

This becomes a more general pattern-matching problem and cannot be solved using a fixed-size sliding window alone.

---

### Variation 4 — Case-Insensitive Matching

Normalize all words and input characters (e.g., convert to lowercase) before processing.

---

### Variation 5 — Massive Dictionary

If `words` contains millions of entries:

- Compress word representations.
- Use integer IDs instead of strings.
- Optimize hash lookups and memory layout.

---

# Common Pitfalls

- Sliding the window by one character instead of one word.
- Forgetting to process every possible starting offset.
- Using a set instead of a frequency map, which fails for duplicate words.
- Not shrinking the window when a word exceeds its allowed count.
- Failing to remove the leftmost word after recording a valid answer, causing overlapping matches to be missed.
- Miscalculating the total window length.

---

# Key Takeaways

- Equal-length words enable fixed-size chunk processing.
- Multi-offset sliding windows guarantee complete coverage.
- Frequency maps are essential for duplicate handling.
- Window expansion and contraction keep the algorithm linear.
- Explaining the optimization journey is just as important as presenting the final solution.