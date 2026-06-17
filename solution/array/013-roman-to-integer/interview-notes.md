# Interview Notes — Roman to Integer

---

# What Interviewer Is Testing

Although classified as an Easy problem, Roman to Integer is frequently used as a screening question because it evaluates several fundamental engineering skills.

---

## 1. Problem Translation Skills

The interviewer wants to see whether you can translate a real-world rule set into code.

Roman numerals contain:

- Symbol mapping
- Addition rules
- Subtraction rules

Candidates must convert these requirements into a deterministic algorithm.

---

## 2. String Traversal

The problem checks whether you can:

- Iterate through a string correctly
- Access neighboring characters safely
- Avoid out-of-bounds errors

Common mistakes often occur when examining the next character.

---

## 3. Hash Map Usage

Interviewers expect candidates to recognize that Roman symbols naturally map to numeric values.

Example:

```text
I → 1
V → 5
X → 10
L → 50
C → 100
D → 500
M → 1000
```

A hash map provides:

```text
Lookup = O(1)
```

---

## 4. Pattern Recognition

The critical observation is:

```text
Smaller value before larger value
⇒ subtraction
```

Recognizing this pattern eliminates the need for multiple special-case conditions.

---

## 5. Edge Case Awareness

Interviewers often watch whether candidates consider:

- Single-character inputs
- Consecutive identical symbols
- Multiple subtraction pairs
- Last-character handling

---

# Typical Follow-up Questions

---

## Follow-up 1

### Can you solve this without a hash map?

Possible approaches:

- Switch statement
- Array lookup
- Character indexing

Example:

```java
switch(c) {
    case 'I': return 1;
    case 'V': return 5;
}
```

Discussion focuses on readability versus flexibility.

---

## Follow-up 2

### Can you traverse from right to left?

Yes.

Alternative logic:

```text
If current < previous
    subtract
Else
    add
```

Example:

```text
IV

Start at V = 5

I < V
Subtract 1

Answer = 4
```

---

## Follow-up 3

### How would you validate Roman numerals?

Examples of invalid numerals:

```text
IIII
VV
IC
XM
```

Discussion may include:

- Grammar rules
- Input validation
- Parser construction

---

## Follow-up 4

### Convert Integer to Roman

Reverse problem:

```text
12. Integer to Roman
```

Candidates may be asked to design both directions.

---

## Follow-up 5

### What if the numeral system changes?

The interviewer may generalize the problem:

```text
Custom symbols
Different subtraction rules
Different bases
```

This tests abstraction and extensibility.

---

# Optimization Journey

---

## Stage 1 — Manual Pair Handling

Many candidates start by explicitly checking:

```text
IV
IX
XL
XC
CD
CM
```

Pseudo-code:

```text
if pair == "IV"
    add 4
else if pair == "IX"
    add 9
...
```

### Problems

- Many conditions
- Hard to maintain
- Easy to miss cases

---

## Stage 2 — Pattern Recognition

Observe:

```text
IV -> 1 before 5
IX -> 1 before 10
XL -> 10 before 50
```

Common property:

```text
current < next
```

---

## Stage 3 — Single Pass Solution

Replace special cases with:

```text
if current < next
    subtract
else
    add
```

Benefits:

- Cleaner code
- Easier explanation
- Fewer bugs
- Better scalability

---

# Whiteboard Strategy

When solving on a whiteboard:

---

## Step 1

Write the Roman mapping.

```text
I → 1
V → 5
X → 10
L → 50
C → 100
D → 500
M → 1000
```

---

## Step 2

Ask:

```text
What makes IV different from VI?
```

Explanation:

```text
IV
1 before 5

VI
5 before 1
```

This leads naturally to the key observation.

---

## Step 3

State the rule.

```text
current < next
⇒ subtract

otherwise
⇒ add
```

---

## Step 4

Walk through an example.

```text
MCMXCIV
```

Show:

```text
+1000
-100
+1000
-10
+100
-1
+5
```

Final:

```text
1994
```

---

## Step 5

Analyze complexity.

```text
Time: O(n)

Space: O(1)
```

---

# Communication Tips

---

## Good Interview Explanation

A strong explanation sounds like:

> Roman numerals are mostly additive, except when a smaller symbol appears before a larger symbol. I can detect those cases by comparing the current symbol with the next symbol. If the current value is smaller, I subtract it; otherwise, I add it. This lets me process the string in a single pass.

---

## Avoid This Explanation

```text
I memorized the six special cases.
```

Why?

Because it demonstrates memorization rather than pattern recognition.

Interviewers prefer:

```text
general rule
```

over

```text
hardcoded cases
```

---

## Narrate Your Thinking

While coding:

```text
I'm checking the next character.

If the current value is smaller,
it belongs to a subtraction pair.

Otherwise I add it normally.
```

This keeps interviewers aligned with your reasoning.

---

# Senior-Level Discussion Points

---

## Maintainability

A senior engineer should discuss:

```text
How easy is the solution to modify?
```

The pattern-based solution is preferable because:

- Fewer conditions
- Lower bug risk
- Easier testing

---

## Extensibility

Suppose the business introduces:

```text
New symbols
Different subtraction rules
```

A map-driven approach allows configuration changes without modifying core logic.

---

## Input Validation

Production systems rarely assume valid input.

Additional validation may include:

```text
Maximum repetitions
Illegal symbol ordering
Invalid subtraction pairs
```

---

## Testing Strategy

Test categories:

### Basic

```text
III
VIII
```

### Subtraction

```text
IV
IX
XL
CM
```

### Mixed

```text
MCMXCIV
```

### Boundary

```text
I
MMMCMXCIX
```

---

# FAANG-Level Variations

---

## Variation 1

Validate whether a Roman numeral is legal.

Example:

```text
Input:
IC

Output:
Invalid
```

---

## Variation 2

Roman Numeral Calculator

Example:

```text
X + IV
```

Output:

```text
XIV
```

Requires:

- Parsing
- Arithmetic
- Re-encoding

---

## Variation 3

Bidirectional Converter

Implement:

```text
romanToInt()
intToRoman()
```

and guarantee:

```text
romanToInt(intToRoman(x)) == x
```

---

## Variation 4

Streaming Roman Numeral Parser

Characters arrive one at a time:

```text
M
C
M
X
C
I
V
```

Process without storing the full string.

---

## Variation 5

Generic Symbol System

Create a framework that supports:

```text
Roman numerals
Custom numeral systems
Configurable subtraction rules
```

This shifts the discussion from algorithm implementation to software design.

---

# Interview Cheat Answer

If asked:

> What's the key insight?

Answer:

```text
Roman numerals are additive except when a smaller value
appears before a larger value. By comparing each symbol
with the next symbol, we can determine whether to add or
subtract and solve the problem in a single pass.
```

---

# Hiring Manager Perspective

Strong Candidate:

- Identifies the subtraction pattern quickly
- Uses a clean hash map solution
- Explains reasoning clearly
- Handles edge cases
- Produces bug-free code

Average Candidate:

- Hardcodes all six subtraction cases
- Requires several hints
- Misses boundary checks

Weak Candidate:

- Cannot identify subtraction logic
- Produces incorrect results for IV, IX, CM, etc.
- Struggles to explain the approach

For an Easy-level interview question, clarity and correctness matter more than cleverness.