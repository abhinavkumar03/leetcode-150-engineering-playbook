# Text Justification — Dry Run

## Example Input

```text
words = [
  "This",
  "is",
  "an",
  "example",
  "of",
  "text",
  "justification."
]

maxWidth = 16
```

---

# High-Level Strategy

For each line:

1. Greedily pack as many words as possible.
2. Calculate remaining spaces.
3. Distribute spaces evenly.
4. Extra spaces go to left-most gaps.
5. Last line is left-justified.

---

# Iteration 1

## Current State

```text
Remaining Words:

[
  "This",
  "is",
  "an",
  "example",
  "of",
  "text",
  "justification."
]
```

---

## Building Line

### Add "This"

```text
Line Words = ["This"]

Length = 4
```

---

### Add "is"

Need:

```text
"This is"

4 + 1 + 2 = 7
```

Fits.

```text
Line Words = ["This", "is"]

Length = 7
```

---

### Add "an"

Need:

```text
"This is an"

4 + 1 + 2 + 1 + 2 = 10
```

Fits.

```text
Line Words = ["This", "is", "an"]

Length = 10
```

---

### Try Adding "example"

Need:

```text
"This is an example"

4 + 1 + 2 + 1 + 2 + 1 + 7
=
18
```

Does NOT fit.

Stop.

---

# Line 1 Finalization

## Selected Words

```text
["This", "is", "an"]
```

---

## Character Count

| Word | Length |
|--------|--------|
| This | 4 |
| is | 2 |
| an | 2 |

```text
Total Characters = 8
```

---

## Space Calculation

```text
maxWidth = 16

Spaces Needed
=
16 - 8
=
8
```

---

## Gap Calculation

```text
Number of Gaps
=
3 - 1
=
2
```

---

## Space Distribution

```text
baseSpaces
=
8 / 2
=
4

extraSpaces
=
8 % 2
=
0
```

---

## Visual Layout

```text
This____is____an
```

Where:

```text
_ = space
```

---

## Output Line

```text
"This    is    an"
```

Length:

```text
16
```

✓ Valid

---

# Iteration 2

## Remaining Words

```text
[
  "example",
  "of",
  "text",
  "justification."
]
```

---

## Building Line

### Add "example"

```text
Length = 7
```

---

### Add "of"

```text
7 + 1 + 2
=
10
```

Fits.

---

### Add "text"

```text
10 + 1 + 4
=
15
```

Fits.

---

### Try Adding "justification."

```text
15 + 1 + 14
=
30
```

Does NOT fit.

Stop.

---

# Line 2 Finalization

## Selected Words

```text
["example", "of", "text"]
```

---

## Character Count

| Word | Length |
|--------|--------|
| example | 7 |
| of | 2 |
| text | 4 |

```text
Total Characters = 13
```

---

## Space Calculation

```text
16 - 13
=
3
```

---

## Gap Calculation

```text
3 words

Gaps = 2
```

---

## Distribution

```text
baseSpaces
=
3 / 2
=
1

extraSpaces
=
3 % 2
=
1
```

---

## Gap Allocation

| Gap | Spaces |
|------|--------|
| example → of | 2 |
| of → text | 1 |

---

## Visual Layout

```text
example__of_text
```

---

## Output Line

```text
"example  of text"
```

Length:

```text
16
```

✓ Valid

---

# Iteration 3

## Remaining Words

```text
["justification."]
```

---

## Building Line

Only one word remains.

```text
Line Words = ["justification."]
```

---

# Last Line Rule

Last line must be:

```text
Left Justified
```

---

## Character Count

```text
justification.

Length = 14
```

---

## Padding Required

```text
16 - 14
=
2
```

---

## Output

```text
"justification.  "
```

Length:

```text
16
```

✓ Valid

---

# Final Output

```text
[
  "This    is    an",
  "example  of text",
  "justification.  "
]
```

---

# Complete Execution Timeline

| Iteration | Selected Words | Characters | Spaces Needed | Output |
|------------|----------------|-------------|----------------|---------|
| 1 | This, is, an | 8 | 8 | This____is____an |
| 2 | example, of, text | 13 | 3 | example__of_text |
| 3 | justification. | 14 | 2 trailing | justification.__ |

---

# Internal State Transition Table

| Step | Action | Current Words | Current Length |
|--------|---------|--------------|----------------|
| 1 | Add This | [This] | 4 |
| 2 | Add is | [This,is] | 7 |
| 3 | Add an | [This,is,an] | 10 |
| 4 | Reject example | [This,is,an] | 10 |
| 5 | Emit Line 1 | Reset | 0 |
| 6 | Add example | [example] | 7 |
| 7 | Add of | [example,of] | 10 |
| 8 | Add text | [example,of,text] | 15 |
| 9 | Reject justification. | [example,of,text] | 15 |
| 10 | Emit Line 2 | Reset | 0 |
| 11 | Add justification. | [justification.] | 14 |
| 12 | Emit Last Line | Done | 0 |

---

# Edge Case Dry Runs

## Case 1 — Single Word Line

### Input

```text
words = ["hello"]

maxWidth = 10
```

---

### Result

```text
"hello_____"
```

Length:

```text
10
```

---

## Case 2 — Uneven Space Distribution

### Input

```text
words = ["a", "b", "c"]

maxWidth = 8
```

---

### Characters

```text
1 + 1 + 1
=
3
```

---

### Spaces Needed

```text
8 - 3
=
5
```

---

### Gaps

```text
2
```

---

### Distribution

```text
baseSpaces = 2

extraSpaces = 1
```

---

### Allocation

```text
Gap1 = 3 spaces
Gap2 = 2 spaces
```

---

### Output

```text
"a___b__c"
```

---

# Key Dry Run Takeaways

1. Words are packed greedily.
2. Space calculation happens after line selection.
3. Extra spaces always go left first.
4. Single-word lines bypass gap calculations.
5. Last line follows different formatting rules.
6. Every output line must have exactly `maxWidth` characters.