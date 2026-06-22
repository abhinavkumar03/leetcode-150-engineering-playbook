# Simulation Pattern

## Pattern Definition

Simulation problems require reproducing a process, movement, system behavior, or state transition exactly as described.

Instead of searching for a mathematical shortcut immediately, the goal is to model the system's rules and execute them correctly.

Typical simulations involve:

- Movement
- Direction changes
- State transitions
- Traversals
- Object interactions
- Process execution

The challenge is usually correctness rather than advanced algorithmic complexity.

---

## Core Principle

> Model the state, not the visualization.

Many simulation problems provide a visual representation that can distract from the actual solution.

Strong engineers identify:

- What state must be stored
- What state changes occur
- What events trigger transitions

and ignore unnecessary visual details.

---

# Recognition Signals

You are likely dealing with a Simulation problem if the statement includes:

### Movement

Examples:

```text
Move up
Move down
Move left
Move right
Traverse
Walk
Travel
Navigate
```

---

### Direction Changes

Examples:

```text
Reverse direction
Bounce
Turn
Rotate
Alternate
```

---

### Process Execution

Examples:

```text
Perform steps
Execute instructions
Follow rules
Process events
```

---

### Visual Layout

Examples:

```text
Arrange
Convert
Draw
Transform
Display
Render
```

Often the visual layout is only a description of the process.

---

# Generic Simulation Template

```pseudo
initialize state

for each input item

    process current state

    if transition condition occurs
        update state

    move to next state

return result
```

---

# State Machine View

Many simulation problems can be represented as a finite state machine.

Example:

```text
State A
   ↓
State B
   ↓
State C
```

Transitions occur when predefined conditions are met.

Thinking in terms of states often simplifies implementation.

---

# Complexity Characteristics

Simulation problems commonly have:

| Metric | Typical Value |
|----------|----------|
| Time | O(n) |
| Space | O(1) to O(n) |
| Optimization Focus | State Reduction |

Most optimizations come from:

- Storing less state
- Avoiding unnecessary structures
- Recognizing repeated patterns

---

# Common Pitfalls

## Building Unnecessary Structures

Many candidates build:

```text
Full matrix
Full graph
Full visualization
```

when only partial state is needed.

---

## Incorrect State Updates

Updating state in the wrong order often causes:

```text
Off-by-one errors
Boundary bugs
Incorrect traversal
```

---

## Missing Edge Cases

Simulation problems frequently break at:

```text
Empty input
Single element
Boundary positions
```

Always test these first.

---

## Ignoring Transition Conditions

Direction changes and boundary events are usually the most important part of the problem.

---

# Pattern Checklist

Before coding ask:

- What state do I need?
- What events change state?
- Can I avoid storing the full structure?
- What are the transition rules?
- Are there boundaries that trigger behavior changes?

---

# Problem Template

## State

```text
currentPosition
direction
mode
counter
```

---

## Transition Conditions

```text
Top reached
Bottom reached
Boundary crossed
Instruction executed
```

---

## Output Construction

```text
Collect results
Build answer
Return output
```

---

# Related Problems

## Easy

- N/A

---

## Medium

- Zigzag Conversion (#6)
- Spiral Matrix (#54)
- Spiral Matrix II (#59)
- Diagonal Traverse (#498)
- Walking Robot Simulation (#874)

---

## Hard

- Simulation-heavy grid traversal problems
- Multi-state robot movement problems
- Advanced matrix navigation problems

---

# Problem Entry

## LeetCode #6 — Zigzag Conversion

### Difficulty

Medium

### Pattern

Simulation

### Secondary Pattern

String Manipulation

### Key Insight

The zigzag matrix never needs to be constructed.

Track only:

```text
currentRow
direction
rows[]
```

Characters are appended to row buffers while moving:

```text
DOWN
↓
↓
BOTTOM
↑
↑
UP
```

When traversal completes:

```text
Concatenate all rows
```

---

### Complexity

| Metric | Value |
|----------|--------|
| Time | O(n) |
| Space | O(n) |

---

### Interview Takeaway

Zigzag Conversion demonstrates a classic simulation principle:

> Store the state required to reproduce behavior, not the visual representation itself.

This mindset appears frequently in real-world systems such as:

- Workflow engines
- Event processors
- Game loops
- State machines
- Stream processors