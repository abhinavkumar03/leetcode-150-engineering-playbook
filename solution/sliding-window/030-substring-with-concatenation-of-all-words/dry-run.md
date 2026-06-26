# Dry Run

## Problem

**30. Substring with Concatenation of All Words**

---

# Goal

Given:

```text
s = "barfoothefoobarman"

words = ["foo", "bar"]
```

Find every starting index where all words appear exactly once, consecutively, and in any order.

Expected Output

```text
[0, 9]
```

---

# Step 1 — Initial Setup

## Calculate Window Information

| Property | Value |
|----------|-------|
| Word Length | 3 |
| Number of Words | 2 |
| Total Window Length | 6 |

Target Frequency Map

| Word | Count |
|------|------:|
| foo | 1 |
| bar | 1 |

We process the string using **multiple offsets** because every word has the same length.

Possible offsets:

```text
Offset 0
Offset 1
Offset 2
```

---

# Offset 0 Walkthrough

String divided into word-sized chunks:

```text
bar | foo | the | foo | bar | man
```

Initialize

```text
left = 0
right = 0

currentMap = {}

wordsInWindow = 0
```

---

## Iteration 1

### Right Pointer

```text
right = 0

Word = "bar"
```

Current Map

```text
{
    bar : 1
}
```

Window

```text
bar
```

Everything is valid.

Move right.

---

## State

| Variable | Value |
|----------|-------|
| left | 0 |
| right | 3 |
| wordsInWindow | 1 |

---

## Iteration 2

Read next word

```text
foo
```

Current Map

```text
{
    bar : 1,
    foo : 1
}
```

Window

```text
bar foo
```

Window size

```text
2 words
```

Matches required number.

Record answer.

```text
Answer = [0]
```

Now slide the window by removing the leftmost word.

Remove

```text
bar
```

Window becomes

```text
foo
```

Updated State

| Variable | Value |
|----------|-------|
| left | 3 |
| right | 6 |
| wordsInWindow | 1 |

---

## Iteration 3

Read

```text
the
```

Not present in target map.

Reset entire window.

Current Map

```text
{}
```

Pointers

```text
left = 9
right = 9
```

---

## State After Reset

| Variable | Value |
|----------|-------|
| left | 9 |
| right | 9 |
| wordsInWindow | 0 |

---

## Iteration 4

Read

```text
foo
```

Current Map

```text
{
    foo : 1
}
```

Window

```text
foo
```

Move right.

---

## Iteration 5

Read

```text
bar
```

Current Map

```text
{
    foo : 1,
    bar : 1
}
```

Window

```text
foo bar
```

Window size equals required number.

Record answer.

```text
Answer = [0, 9]
```

Remove leftmost word.

Remove

```text
foo
```

Window becomes

```text
bar
```

---

## Iteration 6

Read

```text
man
```

Not present.

Reset window.

End of Offset 0.

---

# Offset 1 Walkthrough

Split into chunks from index 1.

```text
arf | oot | hef | oob | arm
```

Every chunk is invalid.

Each invalid word immediately resets the window.

No answers found.

---

# Offset 2 Walkthrough

Chunks

```text
rfo | oth | efo | oba | rma
```

Again, none exist in the target map.

No answers.

---

# Visual Sliding Window

```
Initial

bar | foo | the | foo | bar | man
^^^^^^
Answer → 0

Slide

bar | foo | the | foo | bar | man
      ^^^^^^
Reset

bar | foo | the | foo | bar | man
                ^^^^^^
Answer → 9
```

---

# Complete State Transition Table

| Step | Offset | Left | Right | Current Word | Current Map | Action | Result |
|-----:|-------:|----:|------:|--------------|-------------|--------|--------|
| 1 | 0 | 0 | 0 | bar | {bar:1} | Expand | |
| 2 | 0 | 0 | 3 | foo | {bar:1, foo:1} | Valid Window | Add 0 |
| 3 | 0 | 3 | 6 | the | {} | Reset | |
| 4 | 0 | 9 | 9 | foo | {foo:1} | Expand | |
| 5 | 0 | 9 | 12 | bar | {foo:1, bar:1} | Valid Window | Add 9 |
| 6 | 0 | 12 | 15 | man | {} | Reset | |
| 7 | 1 | - | - | arf | {} | Reset | |
| 8 | 1 | - | - | oot | {} | Reset | |
| 9 | 2 | - | - | rfo | {} | Reset | |
| 10 | 2 | - | - | oth | {} | Reset | |

---

# Example with Duplicate Words

Input

```text
s = "barfoofoo"

words = ["foo", "foo"]
```

Target Map

```text
foo : 2
```

Processing

```
bar → invalid

Reset

foo → count = 1

foo → count = 2

Window valid
```

Answer

```text
[3]
```

This demonstrates why a **frequency map** is required instead of a simple set.

---

# Window Shrinking Example

Suppose

```text
words =

["foo", "bar"]
```

Current Window

```text
foo
bar
foo
```

Current Frequency

```text
foo = 2
bar = 1
```

Target Frequency

```text
foo = 1
bar = 1
```

The window is invalid because `"foo"` appears too many times.

Shrink from the left:

Remove first `"foo"`.

Updated Window

```text
bar
foo
```

Updated Frequency

```text
foo = 1
bar = 1
```

The window becomes valid again.

---

# Why the Algorithm is O(n)

Although the window sometimes shrinks, each word-sized chunk:

- enters the window **once**
- leaves the window **once**

No chunk is processed more than twice.

Therefore,

```text
Time Complexity = O(n)
```

---

# Key Takeaways

- Treat the string as fixed-size word chunks rather than individual characters.
- Process every possible alignment (`0` to `wordLength - 1`).
- Maintain a sliding window with a frequency map.
- Shrink the window only when a word exceeds its allowed frequency.
- Record the starting index whenever the window contains exactly the required number of words.
- This incremental approach avoids rebuilding maps and achieves linear-time performance.