# Dry Run — Length of Last Word

## Goal

Given a string `s`, return the length of its last word.

A word consists only of non-space characters.

---

# Example 1

## Input

```text
s = "Hello World"
```

## Output

```text
5
```

## Visual Walkthrough

```text
String:

H e l l o _ W o r l d
0 1 2 3 4 5 6 7 8 9 10

Start from the end.
                ↓
H e l l o _ W o r l d
                10
```

The last word is:

```text
World
```

Length:

```text
5
```

---

## Step-by-Step State Transitions

### Phase 1: Skip Trailing Spaces

There are no trailing spaces.

| Step | Index | Character | Action |
|------|--------|-----------|---------|
| 1 | 10 | d | Stop skipping |

Current state:

```text
Index = 10
Length = 0
```

---

### Phase 2: Count Last Word

| Step | Index | Character | Length |
|------|--------|-----------|---------|
| 1 | 10 | d | 1 |
| 2 | 9 | l | 2 |
| 3 | 8 | r | 3 |
| 4 | 7 | o | 4 |
| 5 | 6 | W | 5 |
| 6 | 5 | space | Stop |

Final answer:

```text
5
```

---

# Example 2

## Input

```text
s = "   fly me   to   the moon  "
```

## Output

```text
4
```

---

## Visual Walkthrough

```text
String:

"   fly me   to   the moon  "

                           ↑
                      Start Here
```

Trailing spaces exist.

First we skip them.

```text
"   fly me   to   the moon"
                         ↑
```

Now count the final word.

```text
moon
```

Length:

```text
4
```

---

## Detailed Iterations

### Phase 1: Skip Trailing Spaces

| Step | Index | Character | Action |
|------|--------|-----------|---------|
| 1 | 27 | space | Skip |
| 2 | 26 | space | Skip |
| 3 | 25 | n | Stop |

Current state:

```text
Index = 25
Length = 0
```

---

### Phase 2: Count Last Word

| Step | Index | Character | Length |
|------|--------|-----------|---------|
| 1 | 25 | n | 1 |
| 2 | 24 | o | 2 |
| 3 | 23 | o | 3 |
| 4 | 22 | m | 4 |
| 5 | 21 | space | Stop |

Final answer:

```text
4
```

---

# Example 3

## Input

```text
s = "luffy is still joyboy"
```

## Output

```text
6
```

---

## Visual Walkthrough

```text
l u f f y _ i s _ s t i l l _ j o y b o y
                                    ↑
                               Start Here
```

Last word:

```text
joyboy
```

Length:

```text
6
```

---

## State Transition Table

| Step | Index | Character | Length |
|------|--------|-----------|---------|
| 1 | 20 | y | 1 |
| 2 | 19 | o | 2 |
| 3 | 18 | b | 3 |
| 4 | 17 | y | 4 |
| 5 | 16 | o | 5 |
| 6 | 15 | j | 6 |
| 7 | 14 | space | Stop |

Final answer:

```text
6
```

---

# Edge Case 1 — Single Word

## Input

```text
s = "leetcode"
```

### Traversal

| Index | Character | Length |
|---------|---------|---------|
| 7 | e | 1 |
| 6 | d | 2 |
| 5 | o | 3 |
| 4 | c | 4 |
| 3 | t | 5 |
| 2 | e | 6 |
| 1 | e | 7 |
| 0 | l | 8 |

Output:

```text
8
```

---

# Edge Case 2 — Only Spaces

## Input

```text
s = "     "
```

### Skip Trailing Spaces

| Step | Index | Character |
|------|--------|-----------|
| 1 | 4 | space |
| 2 | 3 | space |
| 3 | 2 | space |
| 4 | 1 | space |
| 5 | 0 | space |

Index becomes:

```text
-1
```

No word exists.

Output:

```text
0
```

---

# Edge Case 3 — Empty String

## Input

```text
s = ""
```

Length:

```text
0
```

No traversal occurs.

Output:

```text
0
```

---

# Execution Flow Summary

```text
Start
  │
  ▼
Move to Last Character
  │
  ▼
Skip Trailing Spaces
  │
  ▼
Count Consecutive Non-Spaces
  │
  ▼
Encounter Space or Start of String
  │
  ▼
Return Count
```

---

# Key Insight

The algorithm never needs to examine words that appear before the last word.

By traversing from the end:

1. Ignore irrelevant trailing spaces.
2. Find the last word immediately.
3. Count its characters.
4. Return the result.

This yields:

```text
Time  : O(n)
Space : O(1)
```

which is the optimal solution.