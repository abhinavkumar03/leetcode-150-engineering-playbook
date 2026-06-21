# Reverse Words in a String — Interview Notes

## What Interviewer Is Testing

This problem appears simple but evaluates several fundamental engineering skills.

### 1. String Manipulation Fundamentals

The interviewer wants to verify that you can:

- Parse strings correctly
- Handle delimiters
- Extract meaningful tokens
- Reconstruct output accurately

Many candidates make mistakes when handling whitespace.

---

### 2. Edge Case Awareness

Common edge cases include:

```text
""
"     "
"hello"
"  hello world  "
"a   good   example"
```

Interviewers want to see whether you naturally think about:

- Leading spaces
- Trailing spaces
- Multiple consecutive spaces
- Empty input
- Single-word input

---

### 3. Problem Decomposition

Strong candidates quickly realize:

```text
Reverse characters ❌
Reverse words      ✅
```

The interviewer is looking for the ability to simplify a problem before coding.

---

### 4. Complexity Analysis

Candidates should identify:

```text
Time Complexity  : O(n)
Space Complexity : O(n)
```

and justify why.

---

### 5. Communication Skills

This problem is frequently used to assess how clearly you explain your thought process.

A good explanation often matters more than the code itself.

---

# Typical Follow-up Questions

## Follow-up 1

### Can you solve it without using split()?

Expected discussion:

```text
Manual parsing
Two pointers
Word extraction
```

The interviewer wants to ensure you understand the underlying mechanics.

---

## Follow-up 2

### Can you do it in-place?

This leads to the classic approach:

```text
Reverse entire string
Reverse each word
Clean extra spaces
```

Example:

```text
the sky is blue

↓ Reverse Entire String

eulb si yks eht

↓ Reverse Individual Words

blue is sky the
```

Often seen in senior-level interviews.

---

## Follow-up 3

### What if the input is a mutable character array?

Expected answer:

```text
O(1) extra space solution
```

using in-place reversals.

---

## Follow-up 4

### What if multiple delimiters exist?

Example:

```text
hello,world;foo|bar
```

Discussion points:

- Custom tokenization
- Parsing rules
- Delimiter abstraction

---

## Follow-up 5

### How would you handle Unicode?

Discussion:

- UTF-8 characters
- Rune handling in Go
- Character boundaries
- Internationalization

---

# Optimization Journey

Interviewers often want to see the progression of thought.

---

## Stage 1 — Naive Thinking

Treat string as characters.

```text
Reverse entire string
```

Problem:

```text
the sky

becomes

yks eht
```

Incorrect.

---

## Stage 2 — Extract Words

Recognize that words are the true units.

```text
["the", "sky", "is", "blue"]
```

Reverse:

```text
["blue", "is", "sky", "the"]
```

Join:

```text
"blue is sky the"
```

Correct.

---

## Stage 3 — Space Optimization

Discuss:

```text
In-place reversal
```

Useful when memory constraints matter.

---

## Interview Progression Summary

```text
Brute Force
      ↓
Word Extraction
      ↓
Reverse Array
      ↓
In-Place Optimization
```

---

# Whiteboard Strategy

## Step 1 — Clarify Requirements

Ask:

```text
Can there be leading spaces?
Can there be trailing spaces?
Can there be multiple spaces?
```

This demonstrates attention to detail.

---

## Step 2 — Walk Through Example

Use:

```text
"the sky is blue"
```

Show:

```text
["the","sky","is","blue"]
```

then:

```text
["blue","is","sky","the"]
```

---

## Step 3 — Discuss Complexity

Explain:

```text
Traversal → O(n)
Reversal  → O(k)
Join      → O(n)
```

Overall:

```text
O(n)
```

---

## Step 4 — Code Incrementally

Recommended order:

1. Extract words
2. Reverse words
3. Join words
4. Test edge cases

---

## Step 5 — Validate

Run through:

```text
""
" "
"hello"
"a good   example"
```

before finishing.

---

# Communication Tips

## Strong Candidate Response

A structured explanation:

> The key insight is that we need to reverse words rather than characters. I'll first extract valid words while ignoring extra spaces, then reverse the word sequence, and finally join them using a single space.

This immediately signals understanding.

---

## Avoid Saying

```text
I'll just reverse the string.
```

This often leads to the wrong approach.

---

## Explain Tradeoffs

Mention:

```text
Split + Reverse + Join
```

Advantages:

- Readable
- Maintainable
- Easy to verify

Disadvantages:

- O(n) extra memory

---

## Think Out Loud

Interviewers value reasoning.

Explain:

- Why you're choosing the approach
- Why edge cases are handled
- Why complexity is optimal

---

# Senior-Level Discussion Points

Senior engineers should move beyond simply solving the problem.

---

## API Design Considerations

Questions to discuss:

```text
Should delimiters be configurable?
Should whitespace normalization be optional?
```

---

## Memory Efficiency

Discuss:

```text
In-place transformations
Streaming approaches
Large-scale text processing
```

---

## Production Considerations

Potential concerns:

- Unicode handling
- Localization
- Performance on large text inputs
- Input validation

---

## Readability vs Optimization

Senior engineers should recognize:

```text
Most production systems prefer readability
over micro-optimizations.
```

Therefore:

```text
Split → Reverse → Join
```

is often preferred.

---

# FAANG-Level Variations

## Variation 1

### Reverse Words in a String II

LeetCode 186

Input is a character array.

Requirements:

```text
O(1) extra space
```

---

## Variation 2

### Reverse Each Word

Example:

```text
hello world

↓

olleh dlrow
```

Related:

```text
LeetCode 557
```

---

## Variation 3

### Reverse Sentences in a Stream

Input size may exceed memory.

Discussion:

```text
Chunk processing
Streaming algorithms
External storage
```

---

## Variation 4

### Preserve Delimiters

Example:

```text
hello,world;foo
```

Output requirements may differ.

Requires custom parsing logic.

---

## Variation 5

### Unicode-Aware Reversal

Example:

```text
English
Hindi
Chinese
Japanese
Emoji
```

Requires character-safe operations.

---

# Red Flags Interviewers Notice

### Red Flag #1

Ignoring multiple spaces.

---

### Red Flag #2

Reversing characters instead of words.

---

### Red Flag #3

No complexity discussion.

---

### Red Flag #4

No edge-case testing.

---

### Red Flag #5

Jumping into code before explaining approach.

---

# Hiring Manager Evaluation

### Junior Candidate

Expected:

- Correct solution
- Basic complexity analysis

---

### Mid-Level Candidate

Expected:

- Clean implementation
- Edge case awareness
- Good communication

---

### Senior Candidate

Expected:

- Multiple approaches
- Tradeoff discussion
- In-place optimization discussion
- Production considerations
- Strong reasoning and clarity

---

# Quick Interview Summary

### Recognition Signal

```text
Reverse order of words
Ignore extra spaces
```

### Recommended Approach

```text
Extract Words
      ↓
Reverse Array
      ↓
Join Words
```

### Complexity

```text
Time  : O(n)
Space : O(n)
```

### Senior-Level Follow-Up

```text
Can we do it in O(1) extra space?
```

That is the most common interviewer extension for this problem.