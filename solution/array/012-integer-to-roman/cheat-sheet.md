# Integer to Roman — Cheat Sheet

## Pattern Summary

### Primary Pattern

```text
Greedy
```

### Secondary Pattern

```text
String Construction
```

### Difficulty

```text
Medium
```

### Category

```text
Math
String
```

---

# Recognition Signals

You should immediately think about this pattern when:

### Signal 1

Input must be converted from one representation to another.

Examples:

```text
Integer → Roman
Roman → Integer
Number → Words
Excel Column Conversion
```

---

### Signal 2

A fixed set of symbols and values exists.

Example:

```text
1000 → M
900  → CM
500  → D
...
```

---

### Signal 3

Choosing the largest valid option repeatedly appears natural.

Example:

```text
1994

Choose:
1000
900
90
4
```

---

### Signal 4

Problem contains special predefined combinations.

Example:

```text
IV
IX
XL
XC
CD
CM
```

---

### Signal 5

Rules are deterministic.

Meaning:

```text
Same input
Always same output
```

No backtracking required.

No dynamic programming required.

---

# Key Formula

## Roman Mapping

```text
1000 → M
900  → CM
500  → D
400  → CD
100  → C
90   → XC
50   → L
40   → XL
10   → X
9    → IX
5    → V
4    → IV
1    → I
```

---

## Greedy Rule

```text
While num >= value

    Append symbol

    num -= value
```

---

## Generic Template

```text
result = ""

for every value in descending order

    while number >= value

        append symbol

        subtract value

return result
```

---

# Complexity Cheatsheet

| Metric | Complexity |
|----------|----------|
| Time | O(1) |
| Space | O(1) |

---

## Why O(1)?

Constraint:

```text
1 <= num <= 3999
```

Maximum Roman numeral length is bounded.

The mapping size is fixed:

```text
13 entries
```

Therefore runtime does not grow with input size.

---

# Similar Problems

## Directly Related

### LeetCode 13

```text
Roman to Integer
```

Pattern:

```text
Hash Map
Traversal
```

---

### LeetCode 171

```text
Excel Sheet Column Number
```

Pattern:

```text
Base Conversion
```

---

### LeetCode 168

```text
Excel Sheet Column Title
```

Pattern:

```text
Symbol Conversion
```

---

### LeetCode 273

```text
Integer to English Words
```

Pattern:

```text
Number Representation
```

---

## Greedy Practice Problems

### LeetCode 55

```text
Jump Game
```

---

### LeetCode 45

```text
Jump Game II
```

---

### LeetCode 134

```text
Gas Station
```

---

### LeetCode 135

```text
Candy
```

---

### LeetCode 763

```text
Partition Labels
```

---

# Quick Revision Notes

## Core Idea

```text
Store Roman values and symbols
in descending order.
```

---

## Correct Mapping

```text
1000 M
900  CM
500  D
400  CD
100  C
90   XC
50   L
40   XL
10   X
9    IX
5    V
4    IV
1    I
```

---

## Important Observation

Roman numerals are not purely additive.

Special cases exist:

```text
4  → IV
9  → IX
40 → XL
90 → XC
400 → CD
900 → CM
```

Always include them in the mapping.

---

## Interview Explanation (30 Seconds)

```text
Roman numerals follow a fixed ordered
mapping of values and symbols.

I store these mappings in descending
order and repeatedly choose the
largest value that does not exceed
the remaining number.

After appending the corresponding
symbol, I subtract the value and
continue until the number becomes zero.

Because the Roman numeral system
already contains subtractive forms,
the greedy strategy always works.
```

---

## Common Mistakes

### Forgetting Subtractive Values

Wrong:

```text
1000
500
100
50
10
5
1
```

Correct:

```text
1000
900
500
400
100
90
50
40
10
9
5
4
1
```

---

### Processing in Random Order

Must process:

```text
Largest → Smallest
```

---

### Massive If/Else Chains

Avoid:

```text
if num >= 1000
if num >= 900
if num >= 500
...
```

Prefer:

```text
Mapping Arrays
```

---

### Ignoring String Builder

For production-quality code:

- Go → strings.Builder
- Java → StringBuilder
- JavaScript → Array + join()

---

# Whiteboard Template

```text
values =
[
1000,900,500,400,
100,90,50,40,
10,9,5,4,1
]

symbols =
[
"M","CM","D","CD",
"C","XC","L","XL",
"X","IX","V","IV","I"
]

result = ""

for each value

    while num >= value

        result += symbol
        num -= value

return result
```

---

# Edge Cases Checklist

### Minimum

```text
1 → I
```

---

### Subtractive

```text
4 → IV
9 → IX
```

---

### Tens

```text
40 → XL
90 → XC
```

---

### Hundreds

```text
400 → CD
900 → CM
```

---

### Maximum

```text
3999 → MMMCMXCIX
```

---

# Pattern Memory Hook

Think:

```text
Currency Notes Problem
```

If you need:

```text
1994
```

Use the biggest denomination first:

```text
1000
900
90
4
```

Exactly how Roman numeral conversion works.

---

# One-Line Revision

```text
Greedy + Descending Roman Mapping:
repeatedly take the largest Roman value
that fits into the remaining number.
```