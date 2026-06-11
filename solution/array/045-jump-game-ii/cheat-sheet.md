# Jump Game II — Cheat Sheet

## Pattern Summary

### Primary Pattern

```text
Greedy
```

### Secondary Pattern

```text
Range Expansion
Implicit BFS
Array Traversal
```

### Difficulty

```text
Medium
```

### LeetCode

```text
45. Jump Game II
```

---

# Recognition Signals

Use this pattern when you see:

### Signal 1

You need:

```text
Minimum number of jumps
Minimum moves
Minimum steps
```

---

### Signal 2

Each position provides:

```text
Reachability
Range
Movement options
```

---

### Signal 3

The problem asks:

```text
Minimum transitions to destination
```

instead of:

```text
Can we reach destination?
```

---

### Signal 4

You can process reachable positions in layers.

Think:

```text
BFS Levels
```

without building an actual graph.

---

### Signal 5

A decision made now affects how far you can go later.

Think:

```text
Greedy Future Reach
```

---

# Core Insight

Do NOT decide:

```text
Which index should I jump to?
```

Instead decide:

```text
How far can the next jump reach?
```

This changes the problem from:

```text
Path Search
```

to:

```text
Range Expansion
```

---

# Mental Model

Treat the array like BFS levels.

Example:

```text
[2,3,1,1,4]
```

Level 0:

```text
[0]
```

Level 1:

```text
[1,2]
```

Level 2:

```text
[3,4]
```

Answer:

```text
2 jumps
```

---

# Key Variables

## jumps

Tracks:

```text
How many jumps have been used
```

---

## currentEnd

Tracks:

```text
End of current BFS level
```

---

## farthestReach

Tracks:

```text
Farthest position reachable
from the current level
```

---

# Visual Formula

```text
Current Range
      ↓
Explore All Indices
      ↓
Find Farthest Reach
      ↓
Range Ends
      ↓
Take Jump
      ↓
Expand Next Range
```

---

# Algorithm Template

```go
jumps := 0
currentEnd := 0
farthestReach := 0

for i := 0; i < n-1; i++ {

    farthestReach = max(
        farthestReach,
        i + nums[i]
    )

    if i == currentEnd {
        jumps++
        currentEnd = farthestReach
    }
}

return jumps
```

---

# Why It Works

At every boundary:

```text
All positions reachable
using current jumps
have already been explored.
```

Therefore:

```text
Choosing the maximum future reach
cannot miss a shorter path.
```

Equivalent to:

```text
BFS Level Expansion
```

---

# Complexity Cheatsheet

| Approach | Time | Space |
|-----------|--------|--------|
| Recursion | Exponential | O(n) |
| Memoization | O(n²) | O(n) |
| Dynamic Programming | O(n²) | O(n) |
| Greedy | O(n) | O(1) |

---

# Complexity To Memorize

```text
Time  = O(n)

Space = O(1)
```

---

# Quick Dry Run

Input:

```text
[2,3,1,1,4]
```

---

### Start

```text
jumps = 0
currentEnd = 0
farthestReach = 0
```

---

### i = 0

```text
reach = 2

farthestReach = 2
```

Boundary reached.

```text
jumps = 1
currentEnd = 2
```

---

### i = 1

```text
reach = 4

farthestReach = 4
```

---

### i = 2

Boundary reached.

```text
jumps = 2
currentEnd = 4
```

Destination reachable.

Answer:

```text
2
```

---

# Common Mistakes

## Mistake 1

Looping until:

```text
i < n
```

Correct:

```text
i < n - 1
```

Last index does not need expansion.

---

## Mistake 2

Incrementing jumps at every index.

Wrong.

Increment only when:

```text
i == currentEnd
```

---

## Mistake 3

Tracking the exact jump path.

Unnecessary.

Only the farthest reachable position matters.

---

## Mistake 4

Confusing with Jump Game I.

Jump Game I:

```text
Reachability
```

Jump Game II:

```text
Minimum Jumps
```

---

# Interview Sound Bites

### 15-Second Explanation

> I treat each reachable range as a BFS level. While scanning the current level, I compute the farthest position reachable in the next level. When the current level ends, I increment the jump count and move to the next level.

---

### 30-Second Explanation

> The brute-force solution explores every jump path, which becomes exponential. Instead, I use a greedy strategy that tracks the current reachable range and the farthest reachable position. Whenever the current range ends, I take a jump and extend the range. This simulates BFS levels and guarantees the minimum number of jumps in O(n) time and O(1) space.

---

# Similar Problems

## Directly Related

```text
55. Jump Game
```

Pattern:

```text
Greedy Reachability
```

---

## Same Family

```text
1306. Jump Game III
1340. Jump Game V
1345. Jump Game IV
1696. Jump Game VI
```

---

## Similar Greedy Problems

```text
134. Gas Station
435. Non-overlapping Intervals
452. Minimum Arrows to Burst Balloons
763. Partition Labels
```

---

# Pattern Connections

| Problem | Pattern |
|-----------|-----------|
| Jump Game | Greedy Reachability |
| Jump Game II | Greedy + BFS Levels |
| Rotting Oranges | Explicit BFS |
| Word Ladder | BFS |
| Minimum Genetic Mutation | BFS |

---

# One-Line Memory Trick

```text
Jump boundary reached?
Take a jump.

While scanning boundary?
Keep extending farthest reach.
```

---

# 10-Second Revision

```text
Goal:
Minimum jumps

Pattern:
Greedy + BFS Levels

Track:
currentEnd
farthestReach
jumps

When:
i == currentEnd

Then:
jumps++
currentEnd = farthestReach

Complexity:
O(n) time
O(1) space
```