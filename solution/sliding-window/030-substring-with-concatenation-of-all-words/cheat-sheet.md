# Sliding Window Cheat Sheet

## Problem

**30. Substring with Concatenation of All Words**

**Difficulty:** Hard

**Pattern:** Sliding Window + Hash Map + Fixed-Length String Processing

---

# Pattern Summary

This problem is an advanced variation of the Sliding Window pattern.

Unlike traditional sliding window problems that process one character at a time, this problem processes the string in **fixed-length word chunks**.

The key idea is to maintain a window containing complete words while dynamically tracking their frequencies using a hash map.

---

# Recognition Signals

If a problem contains most of the following clues, consider this pattern.

✅ All strings have the same length

✅ Find every valid starting position

✅ Words can appear in any order

✅ Every word must appear exactly once

✅ Duplicate words may exist

✅ Frequency counting is required

✅ Consecutive substring matching

---

# Core Insight

Instead of checking every substring independently,

process the string using **word-sized jumps**.

For every possible alignment:

```
Offset 0

word word word ...

Offset 1

word word word ...

Offset 2

word word word ...
```

Maintain

- left pointer
- right pointer
- current frequency map

Expand when valid.

Shrink whenever a frequency exceeds the allowed count.

---

# Recognition Flow

```
Equal Length Words?

        │
        ▼
Yes

        │
        ▼
Window Size Fixed?

        │
        ▼
Yes

        │
        ▼
Need Frequencies?

        │
        ▼
Yes

        │
        ▼
Sliding Window + Hash Map
```

---

# Algorithm Template

```text
Build target frequency map

for each offset

    initialize

    left = offset

    current map = {}

    wordsInWindow = 0

    for right += wordLength

        currentWord = next chunk

        if invalid

            reset window

        else

            add word

            while frequency exceeded

                remove left word

            if window complete

                save answer

                remove left word
```

---

# Visual Window

```
Input

bar | foo | the | foo | bar | man

Window

^^^^^^
bar foo

↓

      ^^^^^^
      foo the

↓

                ^^^^^^
                foo bar
```

---

# State Machine

```
Read Word

      │
      ▼

Is Word Valid?

      │

 ┌────┴─────┐

No          Yes

│            │

Reset     Add Word

             │

             ▼

Frequency OK?

      │

 ┌────┴─────┐

No          Yes

│            │

Shrink     Continue

             │

             ▼

Window Full?

      │

 ┌────┴─────┐

No          Yes

│            │

Continue   Save Answer
```

---

# Recognition Formula

```
Window Length

=

Word Length

×

Number Of Words
```

---

# Complexity Cheatsheet

| Approach | Time | Space |
|----------|------|-------|
| Brute Force | O(n × m) | O(m) |
| Sliding Window | O(n) | O(m) |

Where:

- **n** = length of input string
- **m** = number of words

---

# Window Invariants

Always maintain:

```
currentFrequency[word]

≤

targetFrequency[word]
```

Window size never exceeds

```
wordCount
```

Window always starts on a valid word boundary.

---

# Reset Condition

```
Current Word

NOT IN

Target Map
```

Action

```
Clear Window

Move Left

Continue
```

---

# Shrink Condition

```
current[word]

>

target[word]
```

Action

```
Remove Left Word

Move Left
```

Repeat until valid.

---

# Success Condition

```
wordsInWindow

==

wordCount
```

Record

```
left
```

as the answer.

---

# Common Mistakes

### ❌ Sliding Character by Character

Move by

```
wordLength
```

instead.

---

### ❌ Ignoring Offsets

Always process

```
0

1

...

wordLength - 1
```

---

### ❌ Using a Set

Duplicate words require

```
Hash Map
```

not a Set.

---

### ❌ Forgetting to Shrink

When frequency exceeds limit:

```
while current[word] > target[word]
```

---

### ❌ Incorrect Window Length

Correct formula:

```
wordLength

×

wordCount
```

---

# Edge Cases Checklist

✔ Empty string

✔ Empty words array

✔ String shorter than required window

✔ Duplicate words

✔ Invalid words

✔ Repeated words

✔ Overlapping answers

✔ Large inputs

---

# Similar Problems

| LeetCode | Problem | Pattern |
|-----------|---------|---------|
| 3 | Longest Substring Without Repeating Characters | Sliding Window |
| 76 | Minimum Window Substring | Sliding Window |
| 209 | Minimum Size Subarray Sum | Sliding Window |
| 424 | Longest Repeating Character Replacement | Sliding Window |
| 438 | Find All Anagrams in a String | Sliding Window + Hash Map |
| 567 | Permutation in String | Sliding Window |
| 904 | Fruit Into Baskets | Sliding Window |
| 1004 | Max Consecutive Ones III | Sliding Window |

---

# Interview Memory Trick

Remember the sequence:

```
Equal Word Length

↓

Multiple Offsets

↓

Sliding Window

↓

Hash Map

↓

Shrink if Needed

↓

Window Full?

↓

Record Answer
```

---

# One-Minute Revision

- Equal-length words allow **fixed-size chunk processing**.
- Process **every alignment** (`0` to `wordLength - 1`).
- Maintain **target** and **current** frequency maps.
- Expand by one word at a time.
- Shrink when a frequency exceeds its allowed count.
- Record the left index when the window contains exactly all required words.
- Each chunk enters and leaves the window at most once.
- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(m)`

---

# Quick Interview Pitch

> "Since every word has the same length, I process the string in word-sized chunks instead of characters. I run a sliding window for each possible alignment, maintain a frequency map of words in the current window, shrink the window whenever a word appears too many times, and record the starting index whenever the window contains exactly all required words. This avoids rebuilding frequency maps and achieves linear time complexity."