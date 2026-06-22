# Zigzag Conversion

## Problem Statement

Given a string `s` and an integer `numRows`, arrange the characters of the string in a zigzag pattern across the specified number of rows.

After writing the characters in a zigzag fashion, read the rows one by one and concatenate them to produce the final string.

### Example

Input:

s = "PAYPALISHIRING"
numRows = 3

Zigzag Pattern:

P   A   H   N
A P L S I I G
Y   I   R

Output:

"PAHNAPLSIIGYIR"

---

## Difficulty

Medium

---

## Tags

- String
- Simulation
- Array
- State Management

---

## Pattern

### Primary Pattern
Simulation

### Secondary Pattern
String Manipulation

---

## Intuition

The challenge is not performing a mathematical transformation but accurately simulating how characters move through rows.

Imagine writing characters row by row:

1. Move downward until reaching the bottom row.
2. Reverse direction.
3. Move upward diagonally until reaching the top row.
4. Reverse direction again.
5. Continue until all characters are placed.

Instead of constructing an entire zigzag matrix, we only need to store the characters belonging to each row.

After processing the string, concatenate all rows.

---

## Key Observation

The actual visual zigzag structure is unnecessary.

We only care about:

- Current row
- Current direction
- Characters assigned to each row

By storing row strings independently, we can achieve an efficient solution.

Example:

Input:

s = "PAYPALISHIRING"
numRows = 3

Rows become:

Row 0: PAHN
Row 1: APLSIIG
Row 2: YIR

Final Answer:

PAHNAPLSIIGYIR

---

## Brute Force Approach

Construct the complete zigzag grid and place characters into matrix positions.

### Algorithm

1. Create a 2D matrix.
2. Simulate downward movement.
3. Simulate upward diagonal movement.
4. Place characters into matrix cells.
5. Traverse matrix row-by-row.
6. Build final string.

### Complexity

Time Complexity: O(n)

Space Complexity: O(n + matrix size)

### Limitations

- Unnecessary matrix construction.
- Extra memory consumption.
- More complicated implementation.
- Difficult to maintain during interviews.

---

## Optimized Approach

Use row buffers to simulate character placement.

### Algorithm

1. Handle edge case:
   - If `numRows == 1`, return original string.

2. Create an array of row builders.

3. Maintain:
   - currentRow
   - direction

4. For each character:
   - Append character to current row.
   - Reverse direction when reaching:
     - Top row
     - Bottom row
   - Move to next row.

5. Concatenate all rows.

### Why It Works

Every character belongs to exactly one row.

The zigzag pattern only affects the order in which rows receive characters.

Tracking row movement reproduces the pattern without storing the full grid.

The direction variable acts as a simple state machine.

States:

DOWN → DOWN → DOWN

then

UP → UP → UP

and repeat.

---

## Complexity

### Time Complexity

O(n)

Where:

- n = length of string

Each character is processed exactly once.

### Space Complexity

O(n)

We store each character once across row buffers.

---

## Edge Cases

### Empty Input

Input:

s = ""
numRows = 3

Output:

""

---

### Single Character

Input:

s = "A"
numRows = 4

Output:

"A"

---

### Single Row

Input:

s = "HELLO"
numRows = 1

Output:

"HELLO"

Special case because no zigzag exists.

---

### Number of Rows Greater Than String Length

Input:

s = "ABC"
numRows = 5

Output:

"ABC"

Each character occupies a separate row.

---

### Duplicate Characters

Input:

s = "AAAAAA"
numRows = 3

Output remains valid because placement depends on position rather than value.

---

### Negative Values

Not applicable.

The problem only deals with strings and positive row counts.

---

### Large Inputs

Input size can reach thousands of characters.

The O(n) solution scales efficiently because each character is visited once.

---

## Dry Run

Input:

s = "PAYPALISHIRING"
numRows = 3

| Character | Row | Direction |
|------------|------|------------|
| P | 0 | Down |
| A | 1 | Down |
| Y | 2 | Up |
| P | 1 | Up |
| A | 0 | Down |
| L | 1 | Down |
| I | 2 | Up |
| S | 1 | Up |
| H | 0 | Down |
| I | 1 | Down |
| R | 2 | Up |
| I | 1 | Up |
| N | 0 | Down |
| G | 1 | Down |

Rows:

| Row | Content |
|------|---------|
| 0 | PAHN |
| 1 | APLSIIG |
| 2 | YIR |

Final Output:

PAHNAPLSIIGYIR

---

## Common Mistakes

### Forgetting numRows == 1

This is the most common bug.

Without handling it separately, row movement logic fails.

---

### Incorrect Direction Switching

Wrong:

Switch after moving.

Correct:

Switch immediately when reaching top or bottom row.

---

### Using a Full Matrix

A matrix makes the solution larger and harder to debug.

Only row buffers are needed.

---

### Off-by-One Errors

Typical mistakes:

- Row index becomes -1
- Row index becomes numRows

Always reverse direction before moving beyond boundaries.

---

## Interview Discussion

### What Makes This Problem Interesting?

The challenge is not the algorithmic complexity.

The challenge is translating a visual pattern into code.

Interviewers want to see:

- Careful simulation
- Clean state management
- Strong edge-case handling

---

### Expected Solution Evolution

1. Understand zigzag drawing.
2. Simulate manually.
3. Observe row assignment behavior.
4. Store row strings.
5. Eliminate matrix construction.
6. Reach O(n) implementation.

---

### Key Interview Insight

When a problem describes movement, patterns, or visual behavior:

Ask:

"Can I model only the state transitions instead of storing the whole structure?"

That observation often leads to the optimal solution.

---

## Follow-up Questions

### Can this be solved using mathematical indexing?

Yes.

Characters follow repeating cycles:

cycleLength = 2 × numRows − 2

Rows can be computed mathematically.

---

### Can we reduce space complexity?

Not meaningfully.

The output itself requires O(n) space.

---

### What if the string were streamed?

Maintain row builders while processing incoming characters.

---

### Can this be generalized?

Yes.

The same simulation pattern appears in:

- Matrix traversal
- Robot movement
- State machines
- Path simulations

---

## Real World Applications

### Text Rendering Engines

Rendering text in special layouts.

---

### Event Stream Routing

Messages routed through multiple channels following directional rules.

---

### Animation Systems

Objects moving between states.

---

### Workflow Engines

State transition tracking.

---

### Simulation Software

Directional movement modeling.

---

## Related Problems

### Easy

- Two Sum (#1)
- Length of Last Word (#58)

### Medium

- Reverse Words in a String (#151)
- String to Integer (atoi) (#8)
- Longest Substring Without Repeating Characters (#3)

### Pattern-Based

- Spiral Matrix
- Diagonal Traverse
- Matrix Simulation Problems

---

## Takeaway

Zigzag Conversion is a classic simulation problem.

The optimal solution comes from recognizing that the zigzag grid itself is unnecessary.

Track:

- Current row
- Direction
- Row contents

Process each character once, concatenate rows, and achieve an elegant O(n) solution.