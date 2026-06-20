# Longest Common Prefix — Dry Run

## Visual Walkthrough

We will use the Horizontal Scanning approach.

### Input

```text
strs = ["flower", "flow", "flight"]
```

### Output

```text
"fl"
```

---

# Initial State

Use the first string as the candidate prefix.

| Variable | Value  |
| -------- | ------ |
| prefix   | flower |

---

# Iteration 1

Current string:

```text
flow
```

Current prefix:

```text
flower
```

Check:

```text
Does "flow" start with "flower"?
```

Result:

```text
No
```

Remove last character from prefix.

| Step    | Prefix |
| ------- | ------ |
| Initial | flower |
| Trim 1  | flowe  |

Check again:

```text
Does "flow" start with "flowe"?
```

Result:

```text
No
```

Trim again.

| Step   | Prefix |
| ------ | ------ |
| Trim 2 | flow   |

Check:

```text
Does "flow" start with "flow"?
```

Result:

```text
Yes
```

State after processing:

| Variable | Value |
| -------- | ----- |
| prefix   | flow  |

---

# Iteration 2

Current string:

```text
flight
```

Current prefix:

```text
flow
```

Check:

```text
Does "flight" start with "flow"?
```

Result:

```text
No
```

Trim prefix.

| Step    | Prefix |
| ------- | ------ |
| Initial | flow   |
| Trim 1  | flo    |

Check again:

```text
Does "flight" start with "flo"?
```

Result:

```text
No
```

Trim again.

| Step   | Prefix |
| ------ | ------ |
| Trim 2 | fl     |

Check:

```text
Does "flight" start with "fl"?
```

Result:

```text
Yes
```

State after processing:

| Variable | Value |
| -------- | ----- |
| prefix   | fl    |

---

# Final Result

All strings processed.

| String | Starts With "fl" |
| ------ | ---------------- |
| flower | Yes              |
| flow   | Yes              |
| flight | Yes              |

Return:

```text
fl
```

---

# Complete State Transition Table

| Iteration | Current String | Prefix Before | Action     | Prefix After |
| --------- | -------------- | ------------- | ---------- | ------------ |
| Start     | —              | flower        | Initialize | flower       |
| 1         | flow           | flower        | Trim       | flowe        |
| 1         | flow           | flowe         | Trim       | flow         |
| 1         | flow           | flow          | Match      | flow         |
| 2         | flight         | flow          | Trim       | flo          |
| 2         | flight         | flo           | Trim       | fl           |
| 2         | flight         | fl            | Match      | fl           |

---

# Example 2

## Input

```text
strs = ["dog", "racecar", "car"]
```

---

## Initial State

| Variable | Value |
| -------- | ----- |
| prefix   | dog   |

---

## Compare With "racecar"

| Step | Prefix   |
| ---- | -------- |
| dog  | mismatch |
| do   | mismatch |
| d    | mismatch |
| ""   | empty    |

Prefix becomes empty.

Return immediately:

```text
""
```

---

# Example 3

## Input

```text
strs = ["interview"]
```

Only one string exists.

| Variable | Value     |
| -------- | --------- |
| prefix   | interview |

Return:

```text
interview
```

---

# Edge Case Walkthroughs

## Empty Array

Input:

```text
[]
```

Result:

```text
""
```

Reason:

No strings exist to compare.

---

## Identical Strings

Input:

```text
["test", "test", "test"]
```

Processing:

```text
prefix = test
```

Every comparison succeeds.

Result:

```text
test
```

---

## No Common Prefix

Input:

```text
["abc", "xyz", "pqr"]
```

Processing:

```text
abc
ab
a
""
```

Result:

```text
""
```

---

# Key Learning

The prefix can only:

1. Stay the same
2. Become shorter

It can never become longer.

This property makes the Horizontal Scanning solution both efficient and easy to reason about during interviews.
