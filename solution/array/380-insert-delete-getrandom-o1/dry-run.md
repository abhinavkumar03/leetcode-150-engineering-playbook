# Dry Run — Insert Delete GetRandom O(1)

## Goal

Understand how the data structure maintains:

- O(1) Insert
- O(1) Remove
- O(1) GetRandom

using:

```text
Array/List
+
HashMap<Value, Index>
```

---

# Internal Data Structures

We maintain:

```text
values      -> stores elements
indexMap    -> stores element → index mapping
```

Example:

```text
values   = [10, 20, 30]

indexMap = {
    10 : 0,
    20 : 1,
    30 : 2
}
```

---

# Example Walkthrough

Operations:

```text
insert(10)
insert(20)
insert(30)
remove(20)
insert(40)
getRandom()
```

---

# Initial State

```text
values   = []

indexMap = {}
```

---

# Step 1 — insert(10)

### Check Existence

```text
10 not found
```

### Append to Array

```text
values = [10]
```

### Store Index

```text
indexMap = {
    10 : 0
}
```

### Result

```text
true
```

---

## State After Step 1

| Structure | Value |
|------------|--------|
| values | [10] |
| indexMap | {10:0} |

---

# Step 2 — insert(20)

### Check Existence

```text
20 not found
```

### Append

```text
values = [10, 20]
```

### Save Index

```text
indexMap = {
    10 : 0,
    20 : 1
}
```

### Result

```text
true
```

---

## State After Step 2

| Structure | Value |
|------------|--------|
| values | [10,20] |
| indexMap | {10:0,20:1} |

---

# Step 3 — insert(30)

### Append

```text
values = [10,20,30]
```

### Update Map

```text
indexMap = {
    10 : 0,
    20 : 1,
    30 : 2
}
```

### Result

```text
true
```

---

## State After Step 3

| Structure | Value |
|------------|--------|
| values | [10,20,30] |
| indexMap | {10:0,20:1,30:2} |

---

# Step 4 — remove(20)

Current state:

```text
values = [10,20,30]

indexMap = {
    10 : 0
    20 : 1
    30 : 2
}
```

---

## Locate Element

```text
index = 1
```

---

## Find Last Element

```text
lastValue = 30
lastIndex = 2
```

---

## Swap Target with Last Element

Move:

```text
30
```

into:

```text
index 1
```

Array becomes:

```text
[10,30,30]
```

Visualization:

```text
Before

Index:   0   1   2
Value:  10  20  30

After Swap

Index:   0   1   2
Value:  10  30  30
```

---

## Update HashMap

```text
30 : 1
```

Map becomes:

```text
{
    10 : 0,
    20 : 1,
    30 : 1
}
```

---

## Remove Last Element

```text
values = [10,30]
```

---

## Delete Removed Value

Delete:

```text
20
```

Final map:

```text
{
    10 : 0,
    30 : 1
}
```

---

## Result

```text
true
```

---

## State After Step 4

| Structure | Value |
|------------|--------|
| values | [10,30] |
| indexMap | {10:0,30:1} |

---

# Why Swap Works

Without swapping:

```text
[10,20,30]

remove(20)

[10,30]
```

The element:

```text
30
```

must shift left.

Cost:

```text
O(n)
```

---

With swapping:

```text
[10,20,30]

swap 20 ↔ 30

[10,30,20]

remove last

[10,30]
```

Cost:

```text
O(1)
```

---

# Step 5 — insert(40)

Current:

```text
values = [10,30]
```

Append:

```text
values = [10,30,40]
```

Update map:

```text
{
    10 : 0,
    30 : 1,
    40 : 2
}
```

---

## Result

```text
true
```

---

## State After Step 5

| Structure | Value |
|------------|--------|
| values | [10,30,40] |
| indexMap | {10:0,30:1,40:2} |

---

# Step 6 — getRandom()

Current array:

```text
[10,30,40]
```

Size:

```text
3
```

Generate:

```text
randomIndex = rand(0,2)
```

Possible outcomes:

| Random Index | Returned Value |
|-------------|----------------|
| 0 | 10 |
| 1 | 30 |
| 2 | 40 |

---

## Probability Distribution

Each element occupies exactly one slot.

```text
10 → 1/3
30 → 1/3
40 → 1/3
```

Uniform probability:

```text
✓ Correct
```

---

# Complete State Transition Table

| Step | Operation | values | indexMap |
|--------|----------|---------|----------|
| 0 | Start | [] | {} |
| 1 | insert(10) | [10] | {10:0} |
| 2 | insert(20) | [10,20] | {10:0,20:1} |
| 3 | insert(30) | [10,20,30] | {10:0,20:1,30:2} |
| 4 | remove(20) | [10,30] | {10:0,30:1} |
| 5 | insert(40) | [10,30,40] | {10:0,30:1,40:2} |
| 6 | getRandom() | [10,30,40] | {10:0,30:1,40:2} |

---

# Visual Summary

## Insert

```text
Array:
[10,20,30]

Map:
10 -> 0
20 -> 1
30 -> 2
```

---

## Remove

```text
Remove 20

Before:

[10,20,30]

Swap:

[10,30,20]

Pop Last:

[10,30]
```

---

## Get Random

```text
Array:

Index 0 -> 10
Index 1 -> 30
Index 2 -> 40

Pick random index.
Return value.
```

---

# Complexity Walkthrough

## Insert

```text
HashMap lookup      O(1)
Array append        O(1)

Total               O(1)
```

---

## Remove

```text
HashMap lookup      O(1)
Swap                O(1)
Pop last            O(1)
Map update          O(1)

Total               O(1)
```

---

## GetRandom

```text
Generate index      O(1)
Array access        O(1)

Total               O(1)
```

---

# Key Takeaway

The crucial insight is:

```text
HashMap
+
Array
+
Swap-and-Remove
=
All Operations O(1)
```

This is one of the most important data structure design techniques used in coding interviews and production systems.