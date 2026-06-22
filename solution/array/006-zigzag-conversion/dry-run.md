# Zigzag Conversion — Dry Run

## Goal

Understand exactly how characters move through rows during zigzag traversal.

We simulate:

- Current row
- Direction of movement
- Row contents after each iteration

Instead of building the full zigzag matrix, we store characters directly inside row buffers.

---

# Example

Input:

```text
s = "PAYPALISHIRING"
numRows = 3
```

Expected Output:

```text
PAHNAPLSIIGYIR
```

---

# Visual Zigzag Layout

The actual zigzag looks like:

```text
Row 0: P   A   H   N
Row 1: A P L S I I G
Row 2: Y   I   R
```

Reading row-by-row:

```text
PAHN
APLSIIG
YIR
```

Result:

```text
PAHNAPLSIIGYIR
```

---

# Initial State

Create row buffers:

```text
Row 0 = ""
Row 1 = ""
Row 2 = ""
```

Variables:

```text
currentRow = 0
direction = DOWN (+1)
```

---

# Iteration Walkthrough

## Step 1

Character:

```text
P
```

Current Row:

```text
0
```

Add character:

```text
Row 0 = "P"
```

State:

| Variable | Value |
|-----------|--------|
| currentRow | 0 |
| direction | DOWN |

Move:

```text
currentRow = 1
```

---

## Step 2

Character:

```text
A
```

Add:

```text
Row 1 = "A"
```

State:

| Variable | Value |
|-----------|--------|
| currentRow | 1 |
| direction | DOWN |

Move:

```text
currentRow = 2
```

---

## Step 3

Character:

```text
Y
```

Add:

```text
Row 2 = "Y"
```

Bottom row reached.

Reverse direction.

```text
direction = UP
```

Move:

```text
currentRow = 1
```

Rows:

```text
Row 0 = P
Row 1 = A
Row 2 = Y
```

---

## Step 4

Character:

```text
P
```

Add:

```text
Row 1 = AP
```

Move upward:

```text
currentRow = 0
```

Rows:

```text
Row 0 = P
Row 1 = AP
Row 2 = Y
```

---

## Step 5

Character:

```text
A
```

Add:

```text
Row 0 = PA
```

Top row reached.

Reverse direction.

```text
direction = DOWN
```

Move:

```text
currentRow = 1
```

Rows:

```text
Row 0 = PA
Row 1 = AP
Row 2 = Y
```

---

## Step 6

Character:

```text
L
```

Add:

```text
Row 1 = APL
```

Move:

```text
currentRow = 2
```

Rows:

```text
Row 0 = PA
Row 1 = APL
Row 2 = Y
```

---

## Step 7

Character:

```text
I
```

Add:

```text
Row 2 = YI
```

Bottom row reached.

Reverse:

```text
direction = UP
```

Move:

```text
currentRow = 1
```

Rows:

```text
Row 0 = PA
Row 1 = APL
Row 2 = YI
```

---

## Step 8

Character:

```text
S
```

Add:

```text
Row 1 = APLS
```

Move:

```text
currentRow = 0
```

Rows:

```text
Row 0 = PA
Row 1 = APLS
Row 2 = YI
```

---

## Step 9

Character:

```text
H
```

Add:

```text
Row 0 = PAH
```

Top reached.

Reverse:

```text
direction = DOWN
```

Move:

```text
currentRow = 1
```

Rows:

```text
Row 0 = PAH
Row 1 = APLS
Row 2 = YI
```

---

## Step 10

Character:

```text
I
```

Add:

```text
Row 1 = APLSI
```

Move:

```text
currentRow = 2
```

Rows:

```text
Row 0 = PAH
Row 1 = APLSI
Row 2 = YI
```

---

## Step 11

Character:

```text
R
```

Add:

```text
Row 2 = YIR
```

Bottom reached.

Reverse:

```text
direction = UP
```

Move:

```text
currentRow = 1
```

Rows:

```text
Row 0 = PAH
Row 1 = APLSI
Row 2 = YIR
```

---

## Step 12

Character:

```text
I
```

Add:

```text
Row 1 = APLSII
```

Move:

```text
currentRow = 0
```

Rows:

```text
Row 0 = PAH
Row 1 = APLSII
Row 2 = YIR
```

---

## Step 13

Character:

```text
N
```

Add:

```text
Row 0 = PAHN
```

Top reached.

Reverse:

```text
direction = DOWN
```

Move:

```text
currentRow = 1
```

Rows:

```text
Row 0 = PAHN
Row 1 = APLSII
Row 2 = YIR
```

---

## Step 14

Character:

```text
G
```

Add:

```text
Row 1 = APLSIIG
```

Final rows:

```text
Row 0 = PAHN
Row 1 = APLSIIG
Row 2 = YIR
```

---

# Complete State Transition Table

| Step | Character | Row | Direction Before | Direction After |
|--------|-----------|-----|------------------|-----------------|
| 1 | P | 0 | Down | Down |
| 2 | A | 1 | Down | Down |
| 3 | Y | 2 | Down | Up |
| 4 | P | 1 | Up | Up |
| 5 | A | 0 | Up | Down |
| 6 | L | 1 | Down | Down |
| 7 | I | 2 | Down | Up |
| 8 | S | 1 | Up | Up |
| 9 | H | 0 | Up | Down |
|10 | I | 1 | Down | Down |
|11 | R | 2 | Down | Up |
|12 | I | 1 | Up | Up |
|13 | N | 0 | Up | Down |
|14 | G | 1 | Down | Down |

---

# Final Row Buffers

| Row | Content |
|------|---------|
| 0 | PAHN |
| 1 | APLSIIG |
| 2 | YIR |

---

# Building Final Answer

Concatenate rows from top to bottom.

```text
PAHN
+
APLSIIG
+
YIR
```

Result:

```text
PAHNAPLSIIGYIR
```

---

# Direction Flow Visualization

```text
Row 0  ●────●────●────●
         ↘  ↗  ↘  ↗

Row 1    ●──●──●──●──●──●──●
         ↗  ↘  ↗  ↘  ↗  ↘

Row 2      ●────●────●
```

Movement:

```text
DOWN
↓
↓
BOTTOM

UP
↑
↑
TOP

Repeat
```

---

# Key Insight

The zigzag matrix never needs to be built.

We only need three pieces of state:

1. Current row
2. Current direction
3. Characters stored per row

This transforms a visual drawing problem into a simple O(n) simulation.