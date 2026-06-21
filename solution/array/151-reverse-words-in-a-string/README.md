# Reverse Words in a String

## Problem Statement

Given an input string `s`, reverse the order of the words.

A word is defined as a sequence of non-space characters.

The returned string should:

- Contain words in reverse order
- Have exactly one space between words
- Not contain leading or trailing spaces

### Examples

#### Example 1

Input:

```text
s = "the sky is blue"
```

Output:

```text
"blue is sky the"
```

---

#### Example 2

Input:

```text
s = "  hello world  "
```

Output:

```text
"world hello"
```

---

#### Example 3

Input:

```text
s = "a good   example"
```

Output:

```text
"example good a"
```

---

## Difficulty

**Medium**

---

## Tags

- String
- Two Pointers
- Parsing
- Simulation

---

## Pattern

**Primary Pattern:** Two Pointers

**Secondary Pattern:** String Manipulation

---

## Intuition

The problem is not simply reversing characters.

We need to:

1. Extract valid words
2. Ignore extra spaces
3. Reverse word order
4. Join words with a single space

The key observation is that the words themselves remain unchanged.

Only their positions change.

Example:

```text
Original:
the sky is blue

Words:
["the", "sky", "is", "blue"]

Reversed:
["blue", "is", "sky", "the"]

Result:
blue is sky the
```

---

## Key Observation

Extra spaces do not matter.

Instead of trying to manipulate spaces directly:

1. Extract words
2. Reverse the word sequence
3. Reconstruct the string

This converts a messy string problem into an array manipulation problem.

---

## Brute Force Approach

### Idea

Traverse the string character by character.

Build each word manually.

Store all words inside an array.

Reverse the array.

Join the words.

---

### Algorithm

1. Initialize an empty list of words.
2. Traverse the string.
3. Skip spaces.
4. Build a word until the next space.
5. Store the word.
6. Continue until end of string.
7. Reverse the words list.
8. Join using a single space.

---

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n) |
| Space | O(n) |

Where:

- n = length of string

---

### Limitations

Although linear, this approach requires:

- Additional word storage
- Additional output construction

It is not an in-place solution.

---

## Optimized Approach

### Idea

Use built-in string operations efficiently:

1. Split into words while ignoring extra spaces.
2. Reverse the words.
3. Join them with a single space.

Many languages internally optimize these operations.

---

### Algorithm

1. Trim unnecessary spaces.
2. Extract words.
3. Reverse word order.
4. Join with one space.

Pseudo Process:

```text
Input:
"  hello   world  "

Extract:
["hello", "world"]

Reverse:
["world", "hello"]

Join:
"world hello"
```

---

### Why It Works

The problem only requires reversing word order.

By treating words as independent units:

- Spaces become irrelevant
- Logic becomes simpler
- Edge cases become easier to handle

Every word is processed exactly once.

---

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n) |
| Space | O(n) |

Where:

- n = length of string

---

## Edge Cases

### Empty Input

```text
Input: ""
Output: ""
```

---

### Only Spaces

```text
Input: "     "
Output: ""
```

---

### Single Word

```text
Input: "hello"
Output: "hello"
```

---

### Multiple Spaces

```text
Input:
"a   good     example"

Output:
"example good a"
```

---

### Leading and Trailing Spaces

```text
Input:
"   hello world   "

Output:
"world hello"
```

---

### Duplicate Words

```text
Input:
"cat cat dog"

Output:
"dog cat cat"
```

---

### Negative Values

Not applicable because the input consists only of characters.

---

### Large Inputs

```text
Length = 10^4+
```

Solution remains linear and scalable.

---

## Dry Run

### Input

```text
s = "the sky is blue"
```

### Step 1 — Extract Words

| Index | Word Found | Words Array |
|---------|---------|---------|
| 0 | the | [the] |
| 1 | sky | [the, sky] |
| 2 | is | [the, sky, is] |
| 3 | blue | [the, sky, is, blue] |

---

### Step 2 — Reverse Array

| Before | After |
|----------|----------|
| [the, sky, is, blue] | [blue, is, sky, the] |

---

### Step 3 — Join

```text
blue is sky the
```

---

## Common Mistakes

### 1. Reversing Characters Instead of Words

Wrong:

```text
eulb si yks eht
```

Correct:

```text
blue is sky the
```

---

### 2. Keeping Extra Spaces

Wrong:

```text
"blue  is  sky  the"
```

Correct:

```text
"blue is sky the"
```

---

### 3. Forgetting Leading/Trailing Spaces

Input:

```text
"  hello world "
```

Output must be:

```text
"world hello"
```

---

### 4. Empty Token Handling

Splitting by a single space may produce:

```text
["", "", "hello", "", "world"]
```

These empty tokens must be ignored.

---

## Interview Discussion

### Expected Solution

Most interviewers expect:

- O(n) Time
- O(n) Space

using word extraction and reversal.

---

### Follow-Up Discussion

Can this be done:

- In-place?
- Without using split()?
- With O(1) extra space?

These questions often appear in senior-level interviews.

---

### What Interviewers Evaluate

- String parsing ability
- Edge case awareness
- Code cleanliness
- Complexity analysis
- Communication skills

---

## Follow-up Questions

### Follow-Up 1

Can you solve this without using built-in split()?

---

### Follow-Up 2

Can you perform the reversal in-place?

---

### Follow-Up 3

How would you handle Unicode strings?

---

### Follow-Up 4

What changes if punctuation must be preserved differently?

---

### Follow-Up 5

Can this be solved using a stack?

---

## Real World Applications

### Search Engines

Normalizing user-entered text.

---

### Text Editors

Sentence transformations and formatting.

---

### NLP Pipelines

Token extraction and sequence manipulation.

---

### Chat Applications

Message normalization and cleanup.

---

### Log Processing

Reorganizing textual data streams.

---

## Related Problems

| LeetCode | Problem |
|-----------|-----------|
| 58 | Length of Last Word |
| 344 | Reverse String |
| 541 | Reverse String II |
| 557 | Reverse Words in a String III |
| 125 | Valid Palindrome |
| 680 | Valid Palindrome II |
| 186 | Reverse Words in a String II |
| 151 | Reverse Words in a String |

---

### Key Takeaway

This problem demonstrates a classic interview pattern:

> Convert a difficult string-formatting problem into a simpler word-processing problem.

Extract words → Reverse order → Reconstruct output.

This leads to a clean and efficient O(n) solution.