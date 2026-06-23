# Dry Run — Find the Index of the First Occurrence in a String

## Goal

Find the first index where `needle` appears inside `haystack`.

If the substring does not exist, return `-1`.

---

# Example 1

## Input

```text
haystack = "sadbutsad"
needle   = "sad"
```

## Expected Output

```text
0
```

---

## Initial Setup

### Variables

```text
n = len(haystack) = 9
m = len(needle)   = 3
```

Valid starting positions:

```text
0 → (n - m)
0 → 6
```

We only need to check positions:

```text
0, 1, 2, 3, 4, 5, 6
```

---

# Iteration 1

## Start Position = 0

Visual Alignment

```text
haystack : sadbutsad
needle   : sad
           ^^^
```

Character Comparisons

| Offset | haystack[start + offset] | needle[offset] | Match |
|----------|----------|----------|----------|
| 0 | s | s | ✅ |
| 1 | a | a | ✅ |
| 2 | d | d | ✅ |

---

### Result

All characters matched.

```text
"sad" == "sad"
```

Return:

```text
0
```

Algorithm Stops.

---

# Execution Summary

| Start Index | Comparison Result |
|------------|-------------------|
| 0 | Full Match Found |

Returned Value:

```text
0
```

---

# Example 2

## Input

```text
haystack = "leetcode"
needle   = "leeto"
```

## Expected Output

```text
-1
```

---

## Initial Setup

```text
n = 8
m = 5
```

Valid start positions:

```text
0 → 3
```

---

# Iteration 1

## Start Position = 0

Visual Alignment

```text
haystack : leetcode
needle   : leeto
           ^^^^^
```

---

### Character Comparisons

| Offset | Haystack | Needle | Match |
|----------|----------|----------|----------|
| 0 | l | l | ✅ |
| 1 | e | e | ✅ |
| 2 | e | e | ✅ |
| 3 | t | t | ✅ |
| 4 | c | o | ❌ |

Mismatch occurs.

Stop checking this position.

---

# Iteration 2

## Start Position = 1

Visual Alignment

```text
haystack : leetcode
needle   :  leeto
            ^^^^^
```

---

### Character Comparisons

| Offset | Haystack | Needle | Match |
|----------|----------|----------|----------|
| 0 | e | l | ❌ |

Immediate mismatch.

Move to next position.

---

# Iteration 3

## Start Position = 2

Visual Alignment

```text
haystack : leetcode
needle   :   leeto
             ^^^^^
```

---

### Character Comparisons

| Offset | Haystack | Needle | Match |
|----------|----------|----------|----------|
| 0 | e | l | ❌ |

Immediate mismatch.

Move forward.

---

# Iteration 4

## Start Position = 3

Visual Alignment

```text
haystack : leetcode
needle   :    leeto
              ^^^^^
```

---

### Character Comparisons

| Offset | Haystack | Needle | Match |
|----------|----------|----------|----------|
| 0 | t | l | ❌ |

Immediate mismatch.

---

# Final Result

No valid position produced a full match.

Return:

```text
-1
```

---

# State Transition Walkthrough

## Example 1

```text
Start = 0
│
├─ Compare s == s ✅
├─ Compare a == a ✅
├─ Compare d == d ✅
│
└─ Full Match Found
      ↓
   Return 0
```

---

## Example 2

```text
Start = 0
│
├─ l == l ✅
├─ e == e ✅
├─ e == e ✅
├─ t == t ✅
└─ c != o ❌

Start = 1
└─ e != l ❌

Start = 2
└─ e != l ❌

Start = 3
└─ t != l ❌

No Match Found
      ↓
   Return -1
```

---

# Worst-Case Dry Run

## Input

```text
haystack = "aaaaaaaaaa"
needle   = "aaaab"
```

---

### Start = 0

```text
a=a ✅
a=a ✅
a=a ✅
a=a ✅
a!=b ❌
```

---

### Start = 1

```text
a=a ✅
a=a ✅
a=a ✅
a=a ✅
a!=b ❌
```

---

### Start = 2

```text
a=a ✅
a=a ✅
a=a ✅
a=a ✅
a!=b ❌
```

Continues similarly.

---

## Why This Is Worst Case

Almost the entire pattern matches before failing.

This causes repeated work.

Complexity becomes:

```text
O((n - m + 1) × m)
```

which is the worst-case runtime of the brute-force approach.

---

# Key Learning

The algorithm works because:

1. Every valid starting position is checked.
2. Comparison stops immediately after mismatch.
3. The first complete match is returned.
4. If no complete match exists, `-1` is returned.

This guarantees correctness while keeping the implementation simple and interview-friendly.