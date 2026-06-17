# Roman to Integer

## Problem Statement

Roman numerals are represented by seven different symbols:

| Symbol | Value |
|----------|----------|
| I | 1 |
| V | 5 |
| X | 10 |
| L | 50 |
| C | 100 |
| D | 500 |
| M | 1000 |

Given a Roman numeral string `s`, convert it to an integer.

Roman numerals are usually written largest to smallest from left to right. However, there are six special subtraction cases:

- I before V or X
- X before L or C
- C before D or M

Examples:

```text
III      -> 3
LVIII    -> 58
MCMXCIV  -> 1994
```

Return the integer value represented by the Roman numeral.

---

## Difficulty

**Easy**

---

## Tags

- String
- Hash Map
- Character Mapping
- Simulation
- Greedy

---

## Pattern

**Primary Pattern:** Hash Map

**Secondary Pattern:** String Traversal

---

## Intuition

Each Roman symbol corresponds to a fixed numeric value.

Most symbols contribute positively to the final result.

However, when a smaller value appears before a larger value, it indicates subtraction rather than addition.

Examples:

```text
IV = 5 - 1 = 4
IX = 10 - 1 = 9
XL = 50 - 10 = 40
CM = 1000 - 100 = 900
```

While traversing the string, we can compare the current symbol with the next symbol.

- If current < next → subtract current value.
- Otherwise → add current value.

This allows us to solve the problem in a single pass.

---

## Key Observation

A Roman numeral can be evaluated locally.

For every character:

```text
current >= next  → add current
current < next   → subtract current
```

Therefore, we only need to inspect the next character while iterating through the string.

This eliminates the need for complex parsing.

---

## Brute Force Approach

### Idea

Handle every subtraction combination separately.

Examples:

```text
IV
IX
XL
XC
CD
CM
```

Check each pair while traversing the string and manually process them.

### Algorithm

1. Create mappings for all Roman symbols.
2. Traverse the string.
3. Check whether current and next character form a subtraction pair.
4. If yes:
   - Add the pair value.
   - Skip the next character.
5. Otherwise:
   - Add current value.
6. Return result.

### Complexity

#### Time Complexity

```text
O(n)
```

#### Space Complexity

```text
O(1)
```

### Limitations

- Requires handling multiple special cases explicitly.
- More conditional branches.
- Less scalable and harder to maintain.
- Easy to introduce bugs.

---

## Optimized Approach

### Algorithm

1. Create a hash map:

```text
I → 1
V → 5
X → 10
L → 50
C → 100
D → 500
M → 1000
```

2. Traverse the string from left to right.

3. For each character:

   - If current value is smaller than next value:
     - Subtract current value.
   - Otherwise:
     - Add current value.

4. Return the accumulated result.

---

### Why It Works

Roman numeral subtraction is always represented by:

```text
smaller value before larger value
```

Examples:

```text
IV
IX
XL
XC
CD
CM
```

Whenever such a situation occurs, subtracting the current value automatically accounts for the subtraction rule.

All other symbols contribute positively.

Thus every symbol is processed exactly once.

---

### Complexity

#### Time Complexity

```text
O(n)
```

Where:

```text
n = length of Roman numeral string
```

#### Space Complexity

```text
O(1)
```

The hash map contains only seven fixed entries.

---

## Edge Cases

### Empty Input

```text
Input: ""
Output: 0
```

(Not required by LeetCode constraints, but useful for robustness.)

---

### Single Character

```text
Input: "V"
Output: 5
```

---

### Consecutive Same Symbols

```text
Input: "III"
Output: 3
```

---

### Subtraction Cases

```text
Input: "IV"
Output: 4
```

```text
Input: "CM"
Output: 900
```

---

### Large Roman Numerals

```text
Input: "MMMCMXCIX"
Output: 3999
```

Maximum valid Roman numeral under standard constraints.

---

### Mixed Additions and Subtractions

```text
Input: "MCMXCIV"
Output: 1994
```

Contains multiple subtraction pairs.

---

## Dry Run

### Example

```text
Input: s = "MCMXCIV"
```

### Mapping

| Symbol | Value |
|----------|----------|
| M | 1000 |
| C | 100 |
| M | 1000 |
| X | 10 |
| C | 100 |
| I | 1 |
| V | 5 |

### Iteration Table

| Index | Current | Next | Action | Result |
|---------|---------|---------|---------|---------|
| 0 | M(1000) | C(100) | Add | 1000 |
| 1 | C(100) | M(1000) | Subtract | 900 |
| 2 | M(1000) | X(10) | Add | 1900 |
| 3 | X(10) | C(100) | Subtract | 1890 |
| 4 | C(100) | I(1) | Add | 1990 |
| 5 | I(1) | V(5) | Subtract | 1989 |
| 6 | V(5) | End | Add | 1994 |

Final Answer:

```text
1994
```

---

## Common Mistakes

### Mistake 1: Only Adding Values

```text
IV

Wrong:
1 + 5 = 6

Correct:
5 - 1 = 4
```

---

### Mistake 2: Forgetting Look-Ahead

Failing to compare with the next symbol causes incorrect handling of subtraction pairs.

---

### Mistake 3: Index Out of Bounds

Accessing:

```text
s[i + 1]
```

without checking whether it exists.

---

### Mistake 4: Hardcoding Too Many Cases

Manually handling every subtraction pair increases complexity and reduces readability.

---

## Interview Discussion

This is a classic introductory string-processing problem.

Interviewers typically want to evaluate:

- Hash map usage
- Traversal logic
- Edge-case awareness
- Code clarity

A strong candidate should quickly identify:

```text
current < next → subtract
otherwise → add
```

and implement it in one pass.

---

## Follow-up Questions

### 1. Can you solve it without a hash map?

Using switch statements or arrays.

---

### 2. Can you convert Integer → Roman Numeral?

Reverse version of the problem.

Related LeetCode problem:

```text
12. Integer to Roman
```

---

### 3. Can invalid Roman numerals be detected?

Examples:

```text
IIII
VV
IC
```

Requires validation logic.

---

### 4. Can this be generalized for other numeral systems?

Useful for parser and compiler discussions.

---

## Real World Applications

### Symbol Parsing

Converting symbolic representations into numeric values.

---

### Compiler Design

Token interpretation and translation.

---

### Configuration Parsing

Mapping textual formats into machine-readable values.

---

### Data Transformation Pipelines

Transforming encoded representations into numerical values.

---

## Related Problems

### Easy

- 12. Integer to Roman
- 58. Length of Last Word
- 383. Ransom Note

### Medium

- 394. Decode String
- 227. Basic Calculator II

### Pattern-Based

- Character Mapping
- String Parsing
- Simulation
- Hash Map Lookup