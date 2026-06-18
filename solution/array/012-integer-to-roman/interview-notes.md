# Interview Notes — Integer to Roman

# What Interviewer Is Testing

Although this problem appears simple, it evaluates several important engineering skills.

---

## 1. Pattern Recognition

The interviewer wants to see whether you can identify that:

```text
Roman numerals follow a fixed ordered system.
```

Instead of building complicated logic, strong candidates recognize that a predefined mapping is sufficient.

Expected observation:

```text
This is not a mathematical problem.

This is a conversion problem using
a fixed set of rules.
```

---

## 2. Greedy Thinking

A major objective of this problem is validating greedy reasoning.

Interviewer expectation:

```text
At every step,
take the largest Roman value possible.
```

Questions they may ask:

- Why does greedy work?
- Can greedy ever produce an invalid answer?
- What guarantees correctness?

Strong answer:

```text
Roman numeral rules are complete and
contain all subtractive combinations.

Therefore choosing the largest valid value
always leads to the correct representation.
```

---

## 3. Data Representation Skills

Software engineers frequently convert data between formats.

Examples:

- Serialization
- Encoding systems
- Protocol translation
- Currency formatting
- Legacy system integrations

This problem measures comfort with:

```text
Input Format
      ↓
Transformation Logic
      ↓
Output Format
```

---

## 4. Clean Coding Ability

This problem can easily become messy.

Weak candidates often write:

```text
if num >= 1000
if num >= 900
if num >= 500
...
```

Large chains of conditionals indicate poor abstraction.

Strong candidates build:

```text
Value Array
+
Symbol Array
```

and iterate through them.

---

## 5. Edge Case Awareness

Interviewers want confidence that you understand:

- Small values
- Large values
- Subtractive notation
- Repeated symbols

Common checks:

```text
1
4
9
40
90
400
900
3999
```

---

# Typical Follow-up Questions

## Follow-up 1

### Why does greedy work?

Expected answer:

```text
Roman numerals have a complete ordered
symbol system.

Choosing the largest valid symbol
always leaves a smaller valid subproblem.
```

---

## Follow-up 2

### Can you solve it without loops?

Possible solution:

Use lookup tables:

```text
Thousands
Hundreds
Tens
Ones
```

Example:

```text
1994

M
CM
XC
IV
```

Concatenate results.

---

## Follow-up 3

### Can you convert Roman → Integer?

Related problem:

```text
LeetCode 13
Roman to Integer
```

Expected discussion:

- Left-to-right scanning
- Subtractive notation
- Mapping table

---

## Follow-up 4

### How would you validate a Roman numeral?

Potential rules:

```text
IIII  ❌
VV    ❌
IC    ❌
XM    ❌
```

Discussion points:

- Symbol repetition limits
- Legal subtraction pairs
- Ordering constraints

---

## Follow-up 5

### Support numbers greater than 3999

Interview discussion:

Traditional Roman numerals stop around:

```text
3999
```

Possible extension methods:

- Overline notation
- Custom symbols
- Alternative numeral systems

---

# Optimization Journey

Interviewers often expect candidates to explain how they arrived at the final solution.

---

## Stage 1 — Naive Conditional Approach

Example:

```text
if num >= 1000
if num >= 900
if num >= 500
...
```

Pros:

- Easy to write

Cons:

- Repetitive
- Hard to maintain
- Error-prone

---

## Stage 2 — Mapping-Based Solution

Store:

```text
Values
Symbols
```

Example:

```text
1000 -> M
900  -> CM
500  -> D
...
```

Pros:

- Cleaner
- Easier to maintain
- Interview friendly

---

## Stage 3 — Lookup Table Optimization

Use:

```text
Thousands[]
Hundreds[]
Tens[]
Ones[]
```

Pros:

```text
Pure O(1)
```

Cons:

```text
Less intuitive
Harder to explain
```

Most interviewers prefer the mapping-based greedy solution.

---

# Whiteboard Strategy

When solving on a whiteboard:

---

## Step 1

Write Roman numeral values.

```text
I   = 1
V   = 5
X   = 10
L   = 50
C   = 100
D   = 500
M   = 1000
```

---

## Step 2

Add subtractive pairs.

```text
IV
IX
XL
XC
CD
CM
```

This demonstrates understanding of Roman numeral rules.

---

## Step 3

Create ordered mapping.

```text
1000 -> M
900  -> CM
500  -> D
...
1    -> I
```

---

## Step 4

Explain greedy reasoning.

```text
Always choose the largest
value that fits.
```

---

## Step 5

Walk through example:

```text
1994
```

Demonstrate:

```text
1994 -> M
994  -> CM
94   -> XC
4    -> IV
```

Final:

```text
MCMXCIV
```

---

# Communication Tips

## Good Interview Explanation

```text
Roman numerals have a fixed ordered
set of values.

I can store these values and symbols
in descending order.

At each step I greedily choose the
largest value that does not exceed
the current number, append its symbol,
and subtract the value.

Because Roman numeral rules already
include subtractive cases, the greedy
choice always produces a valid answer.
```

---

## Avoid Saying

```text
I memorized this solution.
```

Instead explain:

```text
I derived it from the Roman numeral rules.
```

---

## Discuss Tradeoffs

Mention:

```text
Conditional chain
vs
Mapping array
vs
Lookup tables
```

This demonstrates engineering maturity.

---

# Senior-Level Discussion Points

A senior engineer should go beyond simply solving the problem.

---

## Maintainability

Prefer:

```text
Data-driven logic
```

over:

```text
Large if/else blocks
```

Benefits:

- Easier updates
- Lower bug risk
- Better readability

---

## Extensibility

Could support:

```text
Roman → Integer
Integer → Roman
Validation
Formatting APIs
```

using the same mapping structure.

---

## API Design Consideration

Potential interface:

```text
Convert(number)
ConvertRoman(string)
ValidateRoman(string)
```

Useful in:

- Libraries
- Formatting services
- Educational tools

---

## Testing Strategy

Important test cases:

```text
1
3
4
9
40
90
400
900
58
1994
3999
```

Cover:

- Standard symbols
- Subtractive notation
- Boundary values

---

# FAANG-Level Variations

Large companies sometimes extend the discussion.

---

## Variation 1

Roman to Integer

```text
MCMXCIV
```

Return:

```text
1994
```

Pattern:

```text
Hash Map + Traversal
```

---

## Variation 2

Bidirectional Converter

Implement:

```text
intToRoman()
romanToInt()
```

inside the same class.

---

## Variation 3

Roman Numeral Validator

Determine whether:

```text
MMXXIV
```

is valid.

Pattern:

```text
Parsing + Rules Engine
```

---

## Variation 4

Custom Numeral System

Given:

```text
A = 1
B = 5
C = 10
...
```

Build a generic converter.

Pattern:

```text
Greedy + Configuration
```

---

## Variation 5

Localization Engine

Convert numbers into:

- Roman numerals
- English words
- Scientific notation
- Regional numbering systems

Discussion:

```text
Strategy Pattern
Formatter Design
Conversion Engines
```

---

# Key Takeaways

## Pattern

```text
Greedy
```

---

## Core Insight

```text
Store Roman values in descending order.
Always choose the largest valid value.
```

---

## Interview Expectation

```text
Greedy + Ordered Mapping
```

---

## Complexity

```text
Time:  O(1)
Space: O(1)
```

---

## Related Problems

```text
13. Roman to Integer
273. Integer to English Words
171. Excel Sheet Column Number
168. Excel Sheet Column Title
```

---

## One-Sentence Summary

```text
Integer to Roman is a greedy conversion
problem where we repeatedly select the
largest Roman numeral value that fits
into the remaining number.
```