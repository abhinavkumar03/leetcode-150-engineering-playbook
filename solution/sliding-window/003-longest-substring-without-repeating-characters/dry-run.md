# Dry Run

## Problem

**LeetCode 3 — Longest Substring Without Repeating Characters**

---

# Goal

Find the length of the **longest substring** that contains **no repeated characters**.

We maintain a **Sliding Window** using two pointers:

- **Left Pointer (`left`)** → Start of the current window.
- **Right Pointer (`right`)** → End of the current window.

A **HashSet** stores the characters currently inside the window.

Whenever a duplicate character is encountered:

- Shrink the window from the left.
- Remove characters until the duplicate no longer exists.
- Expand the window again.

---

# Example

```
Input

s = "abcabcbb"
```

Expected Output

```
3
```

The longest substrings are:

```
"abc"
"bca"
"cab"
```

Each has length **3**.

---

# Initial State

| Variable | Value |
|----------|-------|
| left | 0 |
| right | 0 |
| window | {} |
| maxLength | 0 |

---

# Visual Walkthrough

```
String

Index : 0 1 2 3 4 5 6 7
Char  : a b c a b c b b
```

---

# Iteration-by-Iteration Dry Run

| Step | Right | Character | Action | Left | Window | Current Length | Max Length |
|------|------:|-----------|--------|-----:|--------|---------------:|-----------:|
| 1 | 0 | a | Add | 0 | {a} | 1 | 1 |
| 2 | 1 | b | Add | 0 | {a,b} | 2 | 2 |
| 3 | 2 | c | Add | 0 | {a,b,c} | 3 | 3 |
| 4 | 3 | a | Duplicate → Remove 'a', move left | 1 | {b,c,a} | 3 | 3 |
| 5 | 4 | b | Duplicate → Remove 'b', move left | 2 | {c,a,b} | 3 | 3 |
| 6 | 5 | c | Duplicate → Remove 'c', move left | 3 | {a,b,c} | 3 | 3 |
| 7 | 6 | b | Duplicate → Remove 'a', 'b', move left twice | 5 | {c,b} | 2 | 3 |
| 8 | 7 | b | Duplicate → Remove 'c', 'b', move left twice | 7 | {b} | 1 | 3 |

---

# Detailed State Transitions

---

## Step 1

Current character:

```
a
```

Window before:

```
{}
```

Add `a`.

Window after:

```
[a]
```

```
left = 0
right = 0
length = 1
max = 1
```

---

## Step 2

Current character:

```
b
```

No duplicate.

Window becomes:

```
[a b]
```

```
length = 2
max = 2
```

---

## Step 3

Current character:

```
c
```

Window:

```
[a b c]
```

```
length = 3
max = 3
```

---

## Step 4

Current character:

```
a
```

Duplicate found.

Window before:

```
[a b c]
```

Remove from the left:

```
remove a
```

Window:

```
[b c]
```

Move left:

```
left = 1
```

Now add current `a`.

Window:

```
[b c a]
```

```
length = 3
max = 3
```

---

## Step 5

Current character:

```
b
```

Duplicate detected.

Remove:

```
b
```

Window:

```
[c a]
```

Move left:

```
left = 2
```

Insert current `b`.

Window:

```
[c a b]
```

Maximum remains:

```
3
```

---

## Step 6

Current character:

```
c
```

Duplicate.

Remove:

```
c
```

Window:

```
[a b]
```

Move left.

Insert `c`.

Window:

```
[a b c]
```

Still:

```
max = 3
```

---

## Step 7

Current character:

```
b
```

Duplicate exists.

Current window:

```
[a b c]
```

Remove left:

```
remove a
```

Still duplicate.

Remove again:

```
remove b
```

Now window:

```
[c]
```

Insert current `b`.

Window:

```
[c b]
```

```
left = 5
length = 2
```

---

## Step 8

Current character:

```
b
```

Duplicate.

Current window:

```
[c b]
```

Remove:

```
c
```

Still duplicate.

Remove:

```
b
```

Window becomes:

```
{}
```

Insert current `b`.

Window:

```
[b]
```

Maximum remains:

```
3
```

---

# Pointer Movement Visualization

```
Step 1

L
R
a b c a b c b b

Window = a
```

```
Step 2

L
  R
a b c a b c b b

Window = ab
```

```
Step 3

L
    R
a b c a b c b b

Window = abc
```

```
Step 4

  L
      R
a b c a b c b b

Window = bca
```

```
Step 5

    L
        R
a b c a b c b b

Window = cab
```

```
Step 6

      L
          R
a b c a b c b b

Window = abc
```

```
Step 7

          L
            R
a b c a b c b b

Window = cb
```

```
Step 8

              L
              R
a b c a b c b b

Window = b
```

---

# Algorithm State Summary

| Iteration | Left | Right | Window | Max Length |
|-----------|-----:|------:|--------|-----------:|
| 1 | 0 | 0 | a | 1 |
| 2 | 0 | 1 | ab | 2 |
| 3 | 0 | 2 | abc | 3 |
| 4 | 1 | 3 | bca | 3 |
| 5 | 2 | 4 | cab | 3 |
| 6 | 3 | 5 | abc | 3 |
| 7 | 5 | 6 | cb | 3 |
| 8 | 7 | 7 | b | 3 |

---

# Final Output

```
Input

"abcabcbb"
```

```
Longest Substring

"abc"
```

```
Length

3
```

---

# Key Observations

- The window **always contains unique characters**.
- The `right` pointer expands the window.
- The `left` pointer contracts the window only when necessary.
- Each character enters and exits the window at most once.
- This guarantees **O(n)** time complexity.

---

# Dry Run Takeaways

- Sliding Window avoids recomputing overlapping substrings.
- The HashSet provides **O(1)** average-time duplicate checks.
- Maintaining the window invariant (all unique characters) is the core idea.
- The algorithm is efficient, readable, and a common interview favorite for mastering the Sliding Window pattern.