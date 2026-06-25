# Valid Palindrome — Dry Run

## Goal

Determine whether a string is a palindrome after:

1. Removing non-alphanumeric characters.
2. Converting letters to lowercase.

We use the **Two Pointers** approach:

- Left pointer starts from the beginning.
- Right pointer starts from the end.
- Skip invalid characters.
- Compare valid characters.
- Move inward until pointers cross.

---

# Example 1

## Input

```text
"A man, a plan, a canal: Panama"
```

---

## Expected Output

```text
true
```

---

## Visual Representation

```text
A man, a plan, a canal: Panama
^                            ^
L                            R
```

After normalization:

```text
amanaplanacanalpanama
```

This string reads the same forward and backward.

---

# State Transition Overview

```text
Initialize Pointers
        ↓
Skip Invalid Characters
        ↓
Compare Characters
        ↓
Match?
   /        \
 Yes         No
  ↓           ↓
Move       Return
Pointers    False
  ↓
Continue
  ↓
Pointers Cross
  ↓
Return True
```

---

# Iteration-by-Iteration Walkthrough

## Iteration 1

### Current Pointers

```text
A man, a plan, a canal: Panama
^                            ^
L                            R
```

### Characters

| Pointer | Character |
|----------|----------|
| Left | A |
| Right | a |

### Normalize

```text
a == a
```

### Result

Match

### Move Pointers

```text
Left++
Right--
```

---

## Iteration 2

### Current Positions

```text
A man, a plan, a canal: Panama
  ^                        ^
```

Left points to space.

Right points to m.

### Skip Invalid Left Character

Space is ignored.

```text
Left moves to 'm'
```

### Compare

| Pointer | Character |
|----------|----------|
| Left | m |
| Right | m |

Match

---

## Iteration 3

### Compare

| Left | Right |
|--------|--------|
| a | a |

Match

Move inward.

---

## Iteration 4

### Compare

| Left | Right |
|--------|--------|
| n | n |

Match

Move inward.

---

## Iteration 5

### Left Encounters Comma

```text
A man, a plan, a canal: Panama
     ^
```

Comma is not alphanumeric.

Skip it.

```text
Left++
```

Continue until valid character found.

---

## Iteration 6

### Right Encounters Colon

```text
A man, a plan, a canal: Panama
                     ^
```

Colon is ignored.

Skip it.

---

## Remaining Comparisons

| Left Character | Right Character | Result |
|----------------|----------------|----------|
| a | a | Match |
| p | p | Match |
| l | l | Match |
| a | a | Match |
| n | n | Match |
| a | a | Match |
| c | c | Match |
| a | a | Match |
| n | n | Match |
| a | a | Match |

All comparisons succeed.

---

# Final State

Pointers cross:

```text
amanaplanacanalpanama
          ^
```

No mismatches found.

---

## Output

```text
true
```

---

# Example 2

## Input

```text
"race a car"
```

---

## Expected Output

```text
false
```

---

## Pointer Walkthrough

### Step 1

```text
r == r
```

Match

---

### Step 2

```text
a == a
```

Match

---

### Step 3

```text
c == c
```

Match

---

### Step 4

Skip space.

Compare:

```text
e != a
```

Mismatch detected.

---

## Immediate Return

```text
false
```

No further processing required.

---

# Example 3

## Input

```text
" "
```

---

## Expected Output

```text
true
```

---

## Walkthrough

String contains only spaces.

### Left Pointer

```text
Skip space
```

### Right Pointer

```text
Skip space
```

No valid characters remain.

Pointers cross immediately.

---

## Output

```text
true
```

---

# Detailed Pointer Movement Table

## Input

```text
"A man, a plan, a canal: Panama"
```

| Step | Left Index | Right Index | Left Char | Right Char | Action |
|--------|--------|--------|--------|--------|--------|
| 1 | 0 | 29 | A | a | Match |
| 2 | 2 | 28 | m | m | Match |
| 3 | 3 | 27 | a | a | Match |
| 4 | 4 | 26 | n | n | Match |
| 5 | Skip | Skip | Space/Comma | Space/Colon | Ignore |
| 6 | Continue | Continue | a | a | Match |
| 7 | Continue | Continue | p | p | Match |
| 8 | Continue | Continue | l | l | Match |
| 9 | Continue | Continue | a | a | Match |
| 10 | Continue | Continue | n | n | Match |
| End | Crossed | Crossed | — | — | Return True |

---

# Why the Algorithm Works

A palindrome requires:

```text
Character(i) == Character(n - 1 - i)
```

for every valid character.

The algorithm:

1. Removes irrelevant characters logically.
2. Normalizes case.
3. Compares mirrored positions.
4. Stops immediately on mismatch.

Because every character is processed at most once:

```text
Time Complexity: O(n)
Space Complexity: O(1)
```

---

# Key Takeaway

Whenever a problem asks you to:

- Compare from both ends
- Check symmetry
- Validate palindrome conditions

Think:

```text
Two Pointers
Left → Beginning
Right → End
Move Inward
```

This is the core pattern behind many palindrome-based interview questions.