# String Traversal Pattern

## Pattern Definition

String Traversal is the process of iterating through a string character by character to extract information, validate conditions, count occurrences, or transform data.

The traversal can occur:

- Left → Right (Forward Traversal)
- Right → Left (Reverse Traversal)
- Both Directions (Two Pointers)

This pattern is one of the most fundamental techniques in coding interviews and forms the basis for many advanced string algorithms.

---

## When To Use This Pattern

Use String Traversal when:

- Processing characters individually
- Counting occurrences
- Validating string properties
- Detecting boundaries between words
- Parsing text
- Finding prefixes or suffixes
- Handling whitespace or delimiters

---

## Recognition Signals

Look for phrases such as:

```text
word
character
string processing
whitespace
delimiter
last word
first word
parse
token
substring conditions
```

Common interview clues:

```text
Find the last word
Count characters
Ignore spaces
Process text efficiently
Avoid extra memory
```

---

## Core Template

### Forward Traversal

```text
for each character in string:
    process character
```

Pseudo Code:

```text
for i from 0 to n - 1:
    process(s[i])
```

---

### Reverse Traversal

Pseudo Code:

```text
i = n - 1

while i >= 0:
    process(s[i])
    i--
```

---

### Two Pointer Traversal

Pseudo Code:

```text
left = 0
right = n - 1

while left <= right:
    process(left, right)
    left++
    right--
```

---

## Complexity

| Operation | Complexity |
|------------|------------|
| Traversal | O(n) |
| Extra Space | O(1) |

Where:

```text
n = length of string
```

---

## Common Pitfalls

### 1. Ignoring Edge Cases

Examples:

```text
""
" "
"hello"
```

Always verify boundary conditions.

---

### 2. Index Out Of Bounds

Incorrect:

```text
s[i]
```

without checking:

```text
i >= 0
```

or

```text
i < n
```

---

### 3. Unnecessary String Operations

Avoid:

```text
split()
substring()
trim()
```

when a traversal solution already exists.

These operations often create additional memory allocations.

---

### 4. Mishandling Spaces

Common bug:

```text
"Hello World   "
```

Trailing spaces must often be handled explicitly.

---

## Optimization Guidelines

Before writing code ask:

### Do I need the whole string?

If not:

```text
Use traversal
```

instead of:

```text
split()
```

---

### Do I only need the end of the string?

If yes:

```text
Use reverse traversal
```

---

### Do I need matching from both ends?

If yes:

```text
Use two pointers
```

---

# Problem Added

## 58. Length of Last Word

### Difficulty

Easy

### Pattern Usage

Reverse Traversal

### Key Observation

The last word is the first continuous block of non-space characters encountered after removing trailing spaces.

### Optimal Solution

```text
1. Start from the end
2. Skip trailing spaces
3. Count characters
4. Stop at space
5. Return count
```

### Complexity

```text
Time  : O(n)
Space : O(1)
```

### Recognition Signal

```text
Need information about the last word only.
```

This strongly suggests reverse traversal.

---

## Related Problems

### Easy

- 58. Length of Last Word
- 344. Reverse String
- 125. Valid Palindrome

### Medium

- 151. Reverse Words in a String
- 680. Valid Palindrome II
- 3. Longest Substring Without Repeating Characters

### Hard

- 76. Minimum Window Substring

---

## Pattern Progression

```text
58  Length of Last Word
        ↓
151 Reverse Words in a String
        ↓
680 Valid Palindrome II
        ↓
3   Longest Substring Without Repeating Characters
        ↓
76  Minimum Window Substring
```

---

## Interview Notes

When the problem asks for:

```text
last item
last word
suffix
rightmost segment
```

consider:

```text
Reverse Traversal
```

before using extra data structures.

This often leads to:

```text
O(1) Space
```

solutions that interviewers prefer.