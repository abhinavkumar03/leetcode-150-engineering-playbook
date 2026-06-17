# Roman to Integer — Cheat Sheet

## Pattern Summary

**Primary Pattern:** Hash Map + String Traversal

**Core Idea:**

Roman numerals are generally additive.

```text
VI = 5 + 1 = 6
XV = 10 + 5 = 15
```

But when a smaller value appears before a larger value, it becomes subtraction.

```text
IV = 5 - 1 = 4
IX = 10 - 1 = 9
XL = 50 - 10 = 40
CM = 1000 - 100 = 900
```

The entire solution is based on detecting this pattern during traversal.

---

# Recognition Signals

Look for this pattern when:

✅ Characters represent fixed numeric values

✅ A lookup table naturally maps symbols → values

✅ Current element depends on the next element

✅ Input is processed sequentially

✅ Local comparisons determine behavior

Typical interview phrases:

```text
Convert symbolic representation into a number

Character mapping

Custom numeral system

Special subtraction rules

Evaluate encoded string
```

---

# Key Formula

## Roman Rule

```text
current < next
    => subtract current

otherwise
    => add current
```

---

## Pseudocode

```text
result = 0

for each character i

    current = value[s[i]]

    if current < value[s[i + 1]]
        result -= current
    else
        result += current

return result
```

---

# Roman Mapping Table

| Symbol | Value |
|----------|----------|
| I | 1 |
| V | 5 |
| X | 10 |
| L | 50 |
| C | 100 |
| D | 500 |
| M | 1000 |

---

# Valid Subtraction Pairs

| Pair | Value |
|--------|--------|
| IV | 4 |
| IX | 9 |
| XL | 40 |
| XC | 90 |
| CD | 400 |
| CM | 900 |

Observation:

```text
Every valid subtraction pair satisfies:

smaller value < larger value
```

---

# Complexity Cheatsheet

| Operation | Complexity |
|------------|------------|
| Hash Map Lookup | O(1) |
| String Traversal | O(n) |
| Total Time | O(n) |
| Extra Space | O(1) |

Where:

```text
n = length of Roman numeral string
```

---

# Quick Dry Run

## Input

```text
MCMXCIV
```

### Processing

| Character | Action | Result |
|------------|------------|------------|
| M | +1000 | 1000 |
| C | -100 | 900 |
| M | +1000 | 1900 |
| X | -10 | 1890 |
| C | +100 | 1990 |
| I | -1 | 1989 |
| V | +5 | 1994 |

Output:

```text
1994
```

---

# Common Mistakes

## Mistake 1

Adding every value.

```text
IV

1 + 5 = 6 ❌

4 ✅
```

---

## Mistake 2

Forgetting to check the next character.

```text
current < next
```

is the key comparison.

---

## Mistake 3

Accessing out of bounds.

Wrong:

```text
s[i + 1]
```

when

```text
i == n - 1
```

Always verify:

```text
i < n - 1
```

---

## Mistake 4

Hardcoding subtraction pairs.

Avoid:

```text
if pair == "IV"
if pair == "IX"
...
```

Prefer:

```text
current < next
```

---

# Alternative Solution

## Right-to-Left Traversal

Traverse from the end.

Rule:

```text
current < previous
    => subtract

otherwise
    => add
```

Example:

```text
IV

Start from V = 5

I < V

Subtract 1

Answer = 4
```

Complexity remains:

```text
Time  : O(n)
Space : O(1)
```

---

# Similar Problems

## Directly Related

- 12. Integer to Roman
- 58. Length of Last Word
- 171. Excel Sheet Column Number

---

## Hash Map + Parsing

- 205. Isomorphic Strings
- 290. Word Pattern
- 383. Ransom Note

---

## String Processing

- 394. Decode String
- 227. Basic Calculator II
- 224. Basic Calculator

---

# Interview One-Liner

> Roman numerals are additive except when a smaller symbol appears before a larger symbol. During a single left-to-right traversal, subtract when current < next; otherwise add.

---

# 30-Second Revision

### Mapping

```text
I=1
V=5
X=10
L=50
C=100
D=500
M=1000
```

### Core Rule

```text
current < next
    => subtract

otherwise
    => add
```

### Complexity

```text
Time  : O(n)

Space : O(1)
```

### Key Insight

```text
No need to memorize:

IV
IX
XL
XC
CD
CM

Just detect:

current < next
```

### Interview Goal

Demonstrate:

- Hash map usage
- String traversal
- Pattern recognition
- Edge-case handling
- Clean implementation