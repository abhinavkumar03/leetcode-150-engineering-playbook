# Dry Run — Integer to Roman

## Goal

Convert an integer into its Roman numeral representation using a greedy approach.

---

# Algorithm Recap

We maintain an ordered list of Roman numeral values:

| Value | Symbol |
|---------|---------|
| 1000 | M |
| 900 | CM |
| 500 | D |
| 400 | CD |
| 100 | C |
| 90 | XC |
| 50 | L |
| 40 | XL |
| 10 | X |
| 9 | IX |
| 5 | V |
| 4 | IV |
| 1 | I |

For each value:

1. Check if current number is greater than or equal to the value.
2. If yes:
   - Append the symbol.
   - Subtract the value.
3. Repeat until the number becomes smaller.
4. Move to the next value.

---

# Example 1

## Input

```text
num = 1994
```

## Expected Output

```text
MCMXCIV
```

---

# Initial State

| Variable | Value |
|-----------|--------|
| num | 1994 |
| result | "" |

---

# Iteration Walkthrough

## Step 1

Current number:

```text
1994
```

Largest value ≤ 1994:

```text
1000 -> M
```

Append:

```text
M
```

Subtract:

```text
1994 - 1000 = 994
```

State:

| num | result |
|------|---------|
| 994 | M |

---

## Step 2

Current number:

```text
994
```

Largest value ≤ 994:

```text
900 -> CM
```

Append:

```text
CM
```

Result becomes:

```text
MCM
```

Subtract:

```text
994 - 900 = 94
```

State:

| num | result |
|------|---------|
| 94 | MCM |

---

## Step 3

Current number:

```text
94
```

Largest value ≤ 94:

```text
90 -> XC
```

Append:

```text
XC
```

Result becomes:

```text
MCMXC
```

Subtract:

```text
94 - 90 = 4
```

State:

| num | result |
|------|---------|
| 4 | MCMXC |

---

## Step 4

Current number:

```text
4
```

Largest value ≤ 4:

```text
4 -> IV
```

Append:

```text
IV
```

Result becomes:

```text
MCMXCIV
```

Subtract:

```text
4 - 4 = 0
```

State:

| num | result |
|------|---------|
| 0 | MCMXCIV |

---

## Step 5

Number becomes:

```text
0
```

Algorithm stops.

---

# Final Output

```text
MCMXCIV
```

---

# Complete State Transition Table

| Step | Current Number | Chosen Value | Roman Symbol | Remaining Number | Result |
|--------|--------|--------|--------|--------|--------|
| Start | 1994 | - | - | 1994 | "" |
| 1 | 1994 | 1000 | M | 994 | M |
| 2 | 994 | 900 | CM | 94 | MCM |
| 3 | 94 | 90 | XC | 4 | MCMXC |
| 4 | 4 | 4 | IV | 0 | MCMXCIV |
| End | 0 | - | - | 0 | MCMXCIV |

---

# Visual Representation

```text
1994
│
├── 1000 → M
│      Remaining = 994
│
├── 900 → CM
│      Remaining = 94
│
├── 90 → XC
│      Remaining = 4
│
├── 4 → IV
│      Remaining = 0
│
└── Result = MCMXCIV
```

---

# Example 2

## Input

```text
num = 58
```

---

## Execution

### Pick 50

```text
58 >= 50
```

Append:

```text
L
```

Remaining:

```text
8
```

---

### Pick 5

```text
8 >= 5
```

Append:

```text
V
```

Remaining:

```text
3
```

---

### Pick 1

```text
3 >= 1
```

Append:

```text
I
```

Remaining:

```text
2
```

Append again:

```text
I
```

Remaining:

```text
1
```

Append again:

```text
I
```

Remaining:

```text
0
```

---

## Final Result

```text
LVIII
```

---

# Example 3

## Input

```text
num = 3749
```

---

## Detailed Execution

| Current Number | Symbol Added | Remaining |
|----------------|-------------|------------|
| 3749 | M | 2749 |
| 2749 | M | 1749 |
| 1749 | M | 749 |
| 749 | D | 249 |
| 249 | C | 149 |
| 149 | C | 49 |
| 49 | XL | 9 |
| 9 | IX | 0 |

Result:

```text
MMMDCCXLIX
```

---

# Why Greedy Works

At every step:

```text
Choose the largest Roman value possible.
```

This ensures:

- Maximum reduction of remaining value.
- Valid Roman numeral ordering.
- Minimal number of operations.
- Correct handling of subtractive notation.

Because Roman numeral rules are fixed and complete, the greedy choice is always optimal.

---

# Dry Run Summary

## Input

```text
1994
```

## Output

```text
MCMXCIV
```

## Key Operations

```text
1994
→ M
→ CM
→ XC
→ IV
```

## Final Result

```text
MCMXCIV
```

## Pattern Used

```text
Greedy
```

## Core Insight

```text
Always use the largest Roman value
that does not exceed the current number.
```