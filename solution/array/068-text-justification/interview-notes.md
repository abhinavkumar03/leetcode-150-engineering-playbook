# Text Justification — Interview Notes

## What Interviewer Is Testing

### 1. Requirement Interpretation

This problem contains several formatting rules that must all be implemented correctly.

Interviewers want to see whether you can:

- Read requirements carefully.
- Translate specifications into code.
- Avoid missing edge cases.
- Build a correct solution from detailed constraints.

Unlike many algorithmic questions, success depends heavily on implementation accuracy.

---

### 2. Greedy Thinking

The problem statement explicitly says:

```text
Pack words greedily.
```

Interviewers expect candidates to recognize that:

- There is no search space.
- No dynamic programming is required.
- No backtracking is required.
- No optimization decision exists.

The next line is uniquely determined by the greedy packing rule.

---

### 3. Simulation Skills

This is primarily a simulation problem.

Candidates must correctly simulate:

- Line construction
- Space distribution
- Left-biased extra spaces
- Last-line formatting

Strong engineers are often expected to handle these real-world implementation problems efficiently.

---

### 4. Edge Case Awareness

Many candidates solve the main flow but fail on:

- Last line
- Single-word line
- Exact-width line
- Uneven space distribution

Interviewers intentionally test these cases.

---

### 5. Clean Code Organization

This problem becomes significantly easier when broken into helper functions.

Interviewers look for:

```text
Build Line
Distribute Spaces
Handle Last Line
```

instead of one giant function.

---

# Typical Follow-up Questions

## Follow-up 1

### Why is Greedy Correct?

Expected Answer:

```text
The problem itself requires greedily packing words.

Therefore, each line boundary is uniquely determined.
No alternative packing strategy is allowed.
```

---

## Follow-up 2

### What Happens When There Is Only One Word?

Expected Answer:

```text
No gaps exist.

Space distribution would cause division by zero.

Therefore the line is left-justified and padded on the right.
```

---

## Follow-up 3

### Why Do Extra Spaces Go Left?

Expected Answer:

```text
The specification explicitly requires
left-most gaps to receive additional spaces first.
```

---

## Follow-up 4

### Can This Be Done In One Pass?

Expected Answer:

```text
Yes.

We scan words once while forming lines.

Each line is immediately emitted after construction.
```

---

## Follow-up 5

### How Would You Handle Gigabytes Of Text?

Expected Answer:

```text
Stream words from disk.

Build one line at a time.

Write output immediately.

Avoid storing entire input in memory.
```

---

## Follow-up 6

### How Would Unicode Change The Solution?

Expected Answer:

```text
String length becomes more complicated.

Character count must be based on visual width
instead of byte count.
```

Examples:

```text
English characters
Chinese characters
Emoji
Combining characters
```

---

# Optimization Journey

## Stage 1 — Understand The Rules

Before coding:

```text
How are lines formed?
How are spaces distributed?
What is special about the last line?
```

Most mistakes occur here.

---

## Stage 2 — Greedy Line Selection

Collect as many words as possible:

```text
Current Length
+
Required Spaces
+
Next Word Length
<= maxWidth
```

Once the next word does not fit:

```text
Emit current line.
```

---

## Stage 3 — Calculate Space Distribution

Compute:

```text
totalSpaces
baseSpaces
extraSpaces
```

Formula:

```text
baseSpaces = totalSpaces / gaps
extraSpaces = totalSpaces % gaps
```

---

## Stage 4 — Handle Special Cases

### Case A

```text
Last Line
```

Use:

```text
Single spaces between words
Remaining spaces at end
```

---

### Case B

```text
Single Word
```

Use:

```text
Word + trailing spaces
```

---

## Final Complexity

### Time

```text
O(N)
```

Where:

```text
N = total characters
```

---

### Space

```text
O(maxWidth)
```

excluding output.

---

# Whiteboard Strategy

## Step 1

Explain line formation first.

Draw:

```text
This is an example

maxWidth = 16
```

Show where packing stops.

---

## Step 2

Show character counting.

Example:

```text
This
is
an
```

Characters:

```text
4 + 2 + 2 = 8
```

---

## Step 3

Show space computation.

```text
16 - 8 = 8 spaces
```

---

## Step 4

Show gap computation.

```text
3 words

2 gaps
```

---

## Step 5

Distribute spaces.

```text
8 / 2 = 4

4 spaces each
```

---

## Step 6

Handle last line separately.

Most interviewers expect this discussion.

---

# Communication Tips

## Good Explanation

```text
First I greedily determine which words belong
to the current line.

Then I calculate the remaining spaces.

For normal lines I distribute spaces evenly.

For the last line and single-word lines
I left justify.
```

This demonstrates a clear thought process.

---

## Avoid Saying

```text
I will just keep trying spaces until it works.
```

This sounds unstructured.

---

## Mention Explicitly

```text
Single word edge case

Last line edge case

Uneven space edge case
```

Interviewers appreciate proactive thinking.

---

# Senior-Level Discussion Points

## 1. Separation Of Concerns

A maintainable implementation typically uses:

```text
collectWords()
buildJustifiedLine()
buildLastLine()
```

This improves readability and testing.

---

## 2. Production Text Rendering

Real systems such as:

- Word processors
- PDF engines
- E-book readers

use more advanced layout algorithms.

Examples:

```text
Kerning
Hyphenation
Font metrics
Unicode rendering
```

---

## 3. Memory Efficiency

Large-scale systems should:

```text
Stream input
Emit output incrementally
Avoid storing all lines
```

---

## 4. Internationalization

Real-world text engines must support:

```text
RTL languages
Unicode
Variable-width characters
```

Examples:

- Arabic
- Hebrew
- Chinese
- Japanese

---

## 5. Testing Strategy

Senior engineers often discuss:

### Boundary Tests

```text
maxWidth = 1
```

### Single Word

```text
["hello"]
```

### Exact Fit

```text
Line width exactly matches maxWidth
```

### Large Dataset

```text
Thousands of words
```

---

# FAANG-Level Variations

## Variation 1 — Right Justification

Example:

```text
________hello world
```

Questions:

```text
How would you adapt your formatter?
```

---

## Variation 2 — Center Alignment

Example:

```text
____hello world____
```

Discuss:

```text
Left padding
Right padding
Odd spacing behavior
```

---

## Variation 3 — Hyphenation Support

Example:

```text
inter-
national
```

Challenges:

- Word splitting
- Formatting consistency

---

## Variation 4 — Full Text Layout Engine

Design:

```text
Input Stream
→ Line Builder
→ Formatter
→ Renderer
```

Potential discussion topics:

- Extensibility
- Performance
- Memory usage

---

## Variation 5 — Rich Text Support

Input:

```text
<b>Hello</b>
<i>World</i>
```

Questions:

```text
How do formatting tags affect width?
```

---

# Red Flags Interviewers Notice

### Red Flag 1

Ignoring last-line rules.

---

### Red Flag 2

Division-by-zero errors.

---

### Red Flag 3

Uneven space allocation.

---

### Red Flag 4

Nested logic that is difficult to read.

---

### Red Flag 5

Unable to explain why greedy works.

---

# Interview Summary

This problem is less about advanced algorithms and more about engineering discipline.

A strong solution demonstrates:

- Careful requirement analysis
- Clean greedy implementation
- Correct simulation
- Robust edge-case handling
- Readable code structure

Candidates who communicate the formatting logic clearly and handle special cases confidently typically perform very well on this question.