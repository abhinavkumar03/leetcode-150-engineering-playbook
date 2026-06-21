# Reverse Words in a String — Dry Run

## Goal

Given a string containing words separated by spaces:

1. Remove extra spaces
2. Reverse the order of words
3. Return a clean string with a single space between words

---

## Example 1

### Input

```text
s = "the sky is blue"
```

### Expected Output

```text
"blue is sky the"
```

---

# High-Level Visualization

```text
Original String

the sky is blue

        │
        ▼

Extract Words

["the", "sky", "is", "blue"]

        │
        ▼

Reverse Order

["blue", "is", "sky", "the"]

        │
        ▼

Join With Single Space

"blue is sky the"
```

---

# Step 1 — Extract Words

Using whitespace parsing:

```text
Input:
"the sky is blue"
```

### Traversal

| Step | Current Word | Words Array |
|--------|--------|--------|
| 1 | the | ["the"] |
| 2 | sky | ["the", "sky"] |
| 3 | is | ["the", "sky", "is"] |
| 4 | blue | ["the", "sky", "is", "blue"] |

### Result

```text
["the", "sky", "is", "blue"]
```

---

# Step 2 — Reverse Word Order

### Initial State

```text
["the", "sky", "is", "blue"]
```

Pointers:

```text
L = 0
R = 3
```

---

## Iteration 1

### Before Swap

```text
["the", "sky", "is", "blue"]
  L                    R
```

Swap:

```text
the ↔ blue
```

### After Swap

```text
["blue", "sky", "is", "the"]
```

Move pointers:

```text
L = 1
R = 2
```

---

## Iteration 2

### Before Swap

```text
["blue", "sky", "is", "the"]
          L      R
```

Swap:

```text
sky ↔ is
```

### After Swap

```text
["blue", "is", "sky", "the"]
```

Move pointers:

```text
L = 2
R = 1
```

Loop stops because:

```text
L >= R
```

---

# Step 3 — Join Words

Array:

```text
["blue", "is", "sky", "the"]
```

Join using:

```text
" "
```

Result:

```text
"blue is sky the"
```

---

# Final Output

```text
"blue is sky the"
```

---

# Detailed State Transition Table

| Iteration | Left | Right | Array State |
|------------|------------|------------|------------|
| Start | 0 | 3 | ["the","sky","is","blue"] |
| 1 | 0 | 3 | ["blue","sky","is","the"] |
| 2 | 1 | 2 | ["blue","is","sky","the"] |
| End | 2 | 1 | Loop Terminates |

---

# Example 2 — Leading and Trailing Spaces

### Input

```text
s = "  hello world  "
```

---

## Word Extraction

Extra spaces are ignored.

```text
["hello", "world"]
```

---

## Reverse

```text
["world", "hello"]
```

---

## Join

```text
"world hello"
```

---

### Output

```text
"world hello"
```

---

# Example 3 — Multiple Spaces Between Words

### Input

```text
s = "a good   example"
```

---

## Extract

```text
["a", "good", "example"]
```

Notice:

```text
Multiple spaces are discarded.
```

---

## Reverse

```text
["example", "good", "a"]
```

---

## Join

```text
"example good a"
```

---

### Output

```text
"example good a"
```

---

# Visual Pointer Movement

Input:

```text
["the", "sky", "is", "blue"]
```

---

### Round 1

```text
L                          R

the      sky      is      blue
 │                         │
 └──────── swap ───────────┘
```

Result:

```text
blue     sky      is      the
```

---

### Round 2

```text
          L      R

blue     sky      is      the
           │       │
           └ swap ─┘
```

Result:

```text
blue      is      sky      the
```

---

# Edge Case Walkthroughs

## Case 1 — Single Word

Input:

```text
"hello"
```

Extract:

```text
["hello"]
```

Reverse:

```text
["hello"]
```

Output:

```text
"hello"
```

---

## Case 2 — Only Spaces

Input:

```text
"     "
```

Extract:

```text
[]
```

Reverse:

```text
[]
```

Output:

```text
""
```

---

## Case 3 — Empty String

Input:

```text
""
```

Extract:

```text
[]
```

Output:

```text
""
```

---

# Complexity Walkthrough

Assume:

```text
n = length of string
k = number of words
```

### Extraction

```text
O(n)
```

---

### Reversal

```text
O(k)
```

---

### Join

```text
O(n)
```

---

### Total

```text
Time  : O(n)
Space : O(n)
```

---

# Key Learning

The most important insight is:

```text
Do NOT reverse characters.
Reverse WORDS.
```

Transform the problem into:

Extract Words
      ↓
Reverse Array
      ↓
Join Words

This produces a clean and efficient O(n) solution.