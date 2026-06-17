# Dry Run — Roman to Integer

## Objective

Understand exactly how the algorithm converts a Roman numeral into an integer using a single left-to-right traversal.

---

# Algorithm Recap

For each character:

```text
If current value < next value
    subtract current value
Else
    add current value
```

This automatically handles Roman numeral subtraction rules:

```text
IV = 4
IX = 9
XL = 40
XC = 90
CD = 400
CM = 900
```

---

# Roman Numeral Mapping

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

# Example 1

## Input

```text
s = "III"
```

## Expected Output

```text
3
```

---

## Step-by-Step Walkthrough

### Initial State

```text
result = 0
```

---

### Iteration Table

| Index | Current | Next | Comparison | Action | Result |
|---------|---------|---------|---------|---------|---------|
| 0 | I(1) | I(1) | 1 < 1 ❌ | Add 1 | 1 |
| 1 | I(1) | I(1) | 1 < 1 ❌ | Add 1 | 2 |
| 2 | I(1) | End | No Next | Add 1 | 3 |

---

## Final Answer

```text
3
```

---

# Example 2

## Input

```text
s = "LVIII"
```

## Expected Output

```text
58
```

---

## Visual Representation

```text
L   V   I   I   I
50  5   1   1   1
```

---

### Initial State

```text
result = 0
```

---

### Iteration Table

| Index | Current | Next | Action | Result |
|---------|---------|---------|---------|---------|
| 0 | L(50) | V(5) | Add 50 | 50 |
| 1 | V(5) | I(1) | Add 5 | 55 |
| 2 | I(1) | I(1) | Add 1 | 56 |
| 3 | I(1) | I(1) | Add 1 | 57 |
| 4 | I(1) | End | Add 1 | 58 |

---

## Final Answer

```text
58
```

---

# Example 3

## Input

```text
s = "IV"
```

## Expected Output

```text
4
```

---

## Visual Representation

```text
I   V
1   5
```

Since:

```text
1 < 5
```

we subtract the first value.

---

### Initial State

```text
result = 0
```

---

### Iteration Table

| Index | Current | Next | Action | Result |
|---------|---------|---------|---------|---------|
| 0 | I(1) | V(5) | Subtract 1 | -1 |
| 1 | V(5) | End | Add 5 | 4 |

---

## Final Answer

```text
4
```

---

# Example 4

## Input

```text
s = "MCMXCIV"
```

## Expected Output

```text
1994
```

---

# Visual Breakdown

```text
M    C    M    X    C    I    V
1000 100 1000 10  100   1    5
```

Subtraction pairs:

```text
CM = 900
XC = 90
IV = 4
```

---

## Initial State

```text
result = 0
```

---

## Detailed Iteration Table

| Step | Current | Value | Next | Next Value | Decision | Result |
|--------|--------|--------|--------|--------|--------|--------|
| 1 | M | 1000 | C | 100 | Add | 1000 |
| 2 | C | 100 | M | 1000 | Subtract | 900 |
| 3 | M | 1000 | X | 10 | Add | 1900 |
| 4 | X | 10 | C | 100 | Subtract | 1890 |
| 5 | C | 100 | I | 1 | Add | 1990 |
| 6 | I | 1 | V | 5 | Subtract | 1989 |
| 7 | V | 5 | End | - | Add | 1994 |

---

## State Transition Visualization

### Step 1

```text
result = 0 + 1000
result = 1000
```

---

### Step 2

```text
result = 1000 - 100
result = 900
```

---

### Step 3

```text
result = 900 + 1000
result = 1900
```

---

### Step 4

```text
result = 1900 - 10
result = 1890
```

---

### Step 5

```text
result = 1890 + 100
result = 1990
```

---

### Step 6

```text
result = 1990 - 1
result = 1989
```

---

### Step 7

```text
result = 1989 + 5
result = 1994
```

---

# Why Subtraction Works

Roman numerals use subtraction only when:

| Pair | Value |
|--------|--------|
| IV | 4 |
| IX | 9 |
| XL | 40 |
| XC | 90 |
| CD | 400 |
| CM | 900 |

Notice:

```text
smaller value before larger value
```

Therefore:

```text
current < next
```

is a complete signal that subtraction is required.

---

# Edge Case Walkthrough

## Single Character

### Input

```text
V
```

### Processing

```text
result = 0 + 5
```

### Output

```text
5
```

---

## Largest Standard Roman Numeral

### Input

```text
MMMCMXCIX
```

### Breakdown

```text
MMM     = 3000
CM      = 900
XC      = 90
IX      = 9
```

### Output

```text
3999
```

---

# Execution Summary

| Property | Value |
|------------|------------|
| Traversal Direction | Left → Right |
| Data Structure | Hash Map |
| Look-Ahead Required | Yes |
| Revisit Characters | No |
| Time Complexity | O(n) |
| Space Complexity | O(1) |

---

# Key Interview Takeaway

The entire solution depends on one observation:

```text
current < next
    => subtract

otherwise
    => add
```

Recognizing this rule transforms the problem from multiple special cases into a clean single-pass solution.