# Integer to Roman

## Problem Statement

Given an integer, convert it to a Roman numeral.

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

Special subtractive cases:

| Roman | Value |
|---------|---------|
| IV | 4 |
| IX | 9 |
| XL | 40 |
| XC | 90 |
| CD | 400 |
| CM | 900 |

Convert the given integer into its Roman numeral representation.

---

## Difficulty

**Medium**

---

## Tags

- Math
- String
- Greedy

---

## Pattern

**Primary Pattern:** Greedy

**Secondary Pattern:** String Construction

---

## Intuition

Roman numerals follow a fixed set of symbols and combinations.

To generate the Roman representation efficiently:

1. Always use the largest Roman value possible.
2. Append its corresponding symbol.
3. Subtract that value from the number.
4. Repeat until the number becomes zero.

This is a classic greedy strategy because choosing the largest valid symbol at every step always leads to the optimal Roman representation.

---

## Key Observation

Roman numerals can be represented using a predefined ordered mapping:

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

If we process these values from largest to smallest, we can construct the answer greedily.

---

## Brute Force Approach

### Algorithm

1. Repeatedly determine which Roman symbol should be added.
2. Use conditional checks for every possible range.
3. Append symbols one by one.
4. Continue until the number becomes zero.

### Complexity

- Time Complexity: **O(N)**
- Space Complexity: **O(1)**

Where N is the input number.

### Limitations

- Too many conditional branches.
- Hard to maintain.
- Difficult to extend.
- Not elegant for interviews.

---

## Optimized Approach

Use a predefined ordered mapping of values and symbols.

### Algorithm

1. Create two arrays:
   - Roman values
   - Roman symbols
2. Iterate through the values array.
3. While current value ≤ num:
   - Append corresponding symbol.
   - Subtract value from num.
4. Continue until num becomes 0.
5. Return the generated string.

### Why It Works

At every step, choosing the largest valid Roman value leaves the smallest remaining number to process.

Roman numerals are designed so that this greedy choice always produces the correct representation.

Example:

Input = 1994

```
1994 -> M (1000)
994  -> CM (900)
94   -> XC (90)
4    -> IV (4)
```

Result:

```
MCMXCIV
```

Because the Roman system has predefined subtractive rules, the greedy strategy never creates an invalid representation.

### Complexity

| Metric | Complexity |
|----------|----------|
| Time | O(1) |
| Space | O(1) |

Reason:

- Roman numerals are limited to values 1–3999.
- Maximum number of iterations is bounded by a constant.

---

## Edge Cases

### Empty Input

Not applicable because constraints guarantee a valid integer.

---

### Single Element

Input:

```text
1
```

Output:

```text
I
```

---

### Subtractive Cases

Input:

```text
4
```

Output:

```text
IV
```

Input:

```text
9
```

Output:

```text
IX
```

---

### Large Values

Input:

```text
3999
```

Output:

```text
MMMCMXCIX
```

---

### Repeated Symbols

Input:

```text
30
```

Output:

```text
XXX
```

---

### Mixed Representation

Input:

```text
58
```

Output:

```text
LVIII
```

---

## Dry Run

### Example

Input:

```text
num = 1994
```

Mapping:

| Value | Symbol |
|---------|---------|
|1000|M|
|900|CM|
|500|D|
|400|CD|
|100|C|
|90|XC|
|50|L|
|40|XL|
|10|X|
|9|IX|
|5|V|
|4|IV|
|1|I|

---

### Step-by-Step Execution

| Step | Current Num | Selected Value | Symbol Added | Result |
|--------|--------|--------|--------|--------|
| 1 | 1994 | 1000 | M | M |
| 2 | 994 | 900 | CM | MCM |
| 3 | 94 | 90 | XC | MCMXC |
| 4 | 4 | 4 | IV | MCMXCIV |
| 5 | 0 | Done | - | MCMXCIV |

---

### Final Output

```text
MCMXCIV
```

---

## Common Mistakes

### Forgetting Subtractive Notation

Incorrect:

```text
4 -> IIII
```

Correct:

```text
4 -> IV
```

---

### Missing 900, 400, 90, 40, 9, 4

Without these values, the greedy algorithm produces invalid Roman numerals.

---

### Using Unordered Mapping

Roman values must be processed from largest to smallest.

Incorrect order can generate wrong results.

---

### Building Strings Inefficiently

Repeated string concatenation may be expensive in some languages.

Prefer:

- String Builder (Java)
- strings.Builder (Go)
- Array + join (JavaScript)

---

## Interview Discussion

### Expected Solution

Interviewers usually expect:

- Greedy solution
- Ordered mapping
- Clean implementation

---

### Key Talking Points

1. Why greedy works.
2. Why subtractive notation must be included.
3. Constant complexity due to bounded input range.
4. Tradeoffs between lookup tables and conditional logic.

---

### What Makes a Strong Answer

A strong candidate:

- Recognizes the greedy pattern quickly.
- Includes subtractive cases upfront.
- Produces concise code.
- Explains correctness clearly.

---

## Follow-up Questions

### 1. Can you convert Roman to Integer?

Related problem:

- Integer to Roman → Roman to Integer

---

### 2. Can you support numbers larger than 3999?

Would require extending Roman numeral conventions.

---

### 3. Can you implement both conversions?

Build a bidirectional converter:

- Integer → Roman
- Roman → Integer

---

### 4. Can you validate Roman numerals?

Check:

- Invalid symbol order
- Excessive repetitions
- Illegal subtractive combinations

---

## Real World Applications

### Data Representation Conversion

Converting between different symbolic systems.

---

### Formatting Systems

Used in:

- Book chapters
- Outlines
- Academic documents
- Legal references

---

### Rule-Based Encoding

Demonstrates how fixed symbolic rules can be translated into efficient algorithms.

---

### Legacy System Support

Some enterprise systems still use Roman numerals in reporting and formatting modules.

---

## Related Problems

### Easy

- Roman to Integer
- Excel Sheet Column Number

### Medium

- Integer to English Words
- Decode String

### Pattern Related

- Roman to Integer
- Remove K Digits
- Gas Station
- Partition Labels

### Greedy Practice

- Jump Game
- Jump Game II
- Candy
- Task Scheduler
