# Text Justification

## Problem Statement

Given an array of words and a maximum line width `maxWidth`, format the text such that each line has exactly `maxWidth` characters and is fully justified.

Rules:

- Pack words greedily into each line.
- Distribute spaces as evenly as possible between words.
- If spaces cannot be evenly distributed, assign extra spaces to the left gaps first.
- The last line should be left-justified.
- A line containing only one word should be left-justified.

Return the fully justified text as an array of strings.

---

## Difficulty

Hard

---

## Tags

- String
- Greedy
- Simulation
- Array

---

## Pattern

**Primary Pattern:** Simulation

**Secondary Pattern:** Greedy

---

## Intuition

The challenge is not finding an optimal arrangement but accurately following formatting rules.

The key observation is:

1. Build one line at a time.
2. Greedily add words until the next word exceeds `maxWidth`.
3. Once a line is determined:
   - Calculate total spaces needed.
   - Distribute spaces evenly across gaps.
   - Assign extra spaces from left to right.
4. Handle special cases:
   - Single-word lines.
   - Last line.

This mirrors how a text editor formats paragraphs.

---

## Key Observation

For a completed line:

```text
Total Spaces Needed
=
maxWidth - totalCharactersOfWords
```

If there are:

```text
gaps = numberOfWords - 1
```

Then:

```text
baseSpaces = totalSpaces / gaps
extraSpaces = totalSpaces % gaps
```

Every gap gets:

```text
baseSpaces
```

The first:

```text
extraSpaces
```

gaps receive one additional space.

---

## Brute Force Approach

### Idea

Generate every possible spacing configuration and test whether it satisfies justification requirements.

### Algorithm

1. Generate all possible space distributions.
2. Validate line width.
3. Check formatting rules.
4. Choose valid configuration.

### Complexity

#### Time Complexity

```text
Exponential
```

#### Space Complexity

```text
Exponential
```

### Limitations

- Completely impractical.
- Large number of possible distributions.
- Fails for interview-scale inputs.

---

## Optimized Approach

### Algorithm

#### Step 1

Initialize:

```text
currentLineWords
currentLineLength
```

#### Step 2

Greedily add words until:

```text
currentLineLength
+ nextWordLength
+ numberOfRequiredSpaces
>
maxWidth
```

#### Step 3

Process the current line.

##### Case A: Last Line

Join words using a single space.

Pad remaining spaces at the end.

Example:

```text
hello world
```

becomes

```text
"hello world    "
```

##### Case B: Single Word Line

Append trailing spaces.

Example:

```text
"interview      "
```

##### Case C: Normal Justified Line

Compute:

```text
totalSpaces
baseSpaces
extraSpaces
```

Distribute spaces:

```text
Left gaps receive extras first
```

#### Step 4

Repeat until all words are processed.

#### Step 5

Return all formatted lines.

---

### Why It Works

Greedy packing guarantees:

- Maximum utilization of each line.
- Compliance with problem constraints.
- Correct distribution of spaces.

Each word is processed exactly once, making the algorithm efficient.

The formatting logic guarantees:

- Exact width.
- Even spacing.
- Left-biased remainder distribution.
- Proper last-line formatting.

---

### Complexity

#### Time Complexity

```text
O(N)
```

Where:

- N = total characters across all words

Each character is processed a constant number of times.

#### Space Complexity

```text
O(maxWidth)
```

excluding output storage.

Including output:

```text
O(total output size)
```

---

## Edge Cases

### Empty Input

```text
words = []
```

Expected:

```text
[]
```

---

### Single Element

```text
["hello"]
```

Output:

```text
"hello     "
```

---

### Single Word Per Line

```text
["extraordinary"]
```

Must be left justified.

---

### Duplicates

```text
["a","a","a","a"]
```

Spacing logic remains unchanged.

---

### Negative Values

Not applicable.

Words contain only valid characters.

---

### Large Inputs

```text
1000+ words
```

Solution remains linear.

---

### Uneven Space Distribution

Example:

```text
Words Length = 10
maxWidth = 16
```

Spaces Needed:

```text
6
```

Gaps:

```text
4
```

Result:

```text
2 2 1 1
```

Extra spaces go left first.

---

## Dry Run

### Input

```text
words =
["This","is","an","example","of","text","justification."]

maxWidth = 16
```

---

### Line 1 Construction

| Step | Words | Length |
|--------|--------|--------|
| Add This | [This] | 4 |
| Add is | [This,is] | 6 |
| Add an | [This,is,an] | 8 |
| Cannot add example | Stop | |

Words:

```text
This is an
```

Characters:

```text
4 + 2 + 2 = 8
```

Spaces Needed:

```text
16 - 8 = 8
```

Gaps:

```text
2
```

Distribution:

```text
4 spaces each
```

Output:

```text
"This    is    an"
```

---

### Line 2 Construction

| Step | Words | Length |
|--------|--------|--------|
| example | 7 |
| of | 9 |
| text | 13 |
| stop | |

Characters:

```text
13
```

Spaces Needed:

```text
3
```

Gaps:

```text
2
```

Distribution:

```text
2 + 1
```

Output:

```text
"example  of text"
```

---

### Final Line

```text
justification.
```

Left justified:

```text
"justification.  "
```

---

### Final Output

```text
[
 "This    is    an",
 "example  of text",
 "justification.  "
]
```

---

## Common Mistakes

### Forgetting Last Line Rule

Incorrect:

```text
Fully justify last line
```

Correct:

```text
Left justify last line
```

---

### Forgetting Single Word Case

Division by zero occurs when:

```text
gaps = 0
```

Must handle separately.

---

### Wrong Extra Space Distribution

Incorrect:

```text
Right side gets extras
```

Correct:

```text
Left side gets extras
```

---

### Miscounting Spaces

Common mistake:

```text
totalLength includes spaces
```

Keep:

```text
wordLength
```

and

```text
spaceLength
```

separate.

---

## Interview Discussion

Key topics interviewers often explore:

1. Why greedy packing works.
2. How space distribution is computed.
3. Handling uneven spacing.
4. Supporting streaming text input.
5. Formatting huge files efficiently.
6. Internationalization considerations.
7. Word wrapping in document editors.

---

## Follow-up Questions

### Follow-up 1

How would you center-align text?

### Follow-up 2

How would you right-align text?

### Follow-up 3

How would you justify text containing Unicode characters?

### Follow-up 4

How would you process gigabyte-scale files?

### Follow-up 5

How would you support custom spacing rules?

---

## Real World Applications

### Text Editors

- Microsoft Word
- Google Docs

### Publishing Systems

- Newspapers
- Magazines
- Books

### PDF Generation

- Invoice formatting
- Report generation

### Terminal Rendering

- CLI output formatting
- Log presentation

### E-book Readers

- Kindle-style rendering

---

## Related Problems

### Easy

- 58. Length of Last Word

### Medium

- 151. Reverse Words in a String
- 71. Simplify Path
- 6. Zigzag Conversion

### Hard

- 68. Text Justification
- 44. Wildcard Matching

---

## Key Takeaway

This problem is an excellent example of a real-world engineering simulation task.

Success depends on:

- Careful requirement analysis
- Precise edge-case handling
- Correct space distribution
- Clean implementation

While the algorithm is straightforward, implementation accuracy is what makes this a Hard problem.