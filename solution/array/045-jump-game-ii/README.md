# Jump Game II

## Problem Statement

You are given a 0-indexed array `nums` of length `n`.

Each element `nums[i]` represents the maximum jump length from index `i`.

Your goal is to reach the last index using the minimum number of jumps.

You can assume that it is always possible to reach the last index.

Return the minimum number of jumps required.

### Example 1

Input:

nums = [2,3,1,1,4]

Output:

2

Explanation:

Jump from index 0 → index 1 → index 4

### Example 2

Input:

nums = [2,3,0,1,4]

Output:

2

---

## Difficulty

Medium

---

## Tags

- Array
- Greedy
- Dynamic Programming
- BFS Concept
- Optimization

---

## Pattern

### Primary Pattern

Greedy

### Secondary Pattern

Range Expansion / Implicit BFS

---

## Intuition

A straightforward approach would be to explore every possible jump and determine the minimum path to the end.

However, that creates many overlapping decisions.

Instead, observe that:

- Every jump creates a range of reachable positions.
- While scanning the current reachable range, we can determine the farthest position reachable by the next jump.
- Once we finish scanning the current range, we commit to a jump and move to the next range.

This is similar to BFS where:

- One jump = one BFS level.
- Reachable indices form a level.
- Expanding the farthest reachable index builds the next level.

---

## Key Observation

At any point:

- `currentEnd` represents the boundary of the current jump.
- `farthestReach` represents the farthest position we can reach while exploring the current range.

When the current index reaches `currentEnd`:

- We have finished exploring all positions available in the current jump.
- A new jump becomes necessary.
- The next boundary becomes `farthestReach`.

This guarantees the minimum number of jumps because we always maximize future reach before committing to the next jump.

---

## Brute Force Approach

### Idea

Try every possible jump from every index and recursively compute the minimum jumps required to reach the end.

### Algorithm

1. Start from index 0.
2. Explore all reachable indices.
3. Recursively calculate minimum jumps from each reachable index.
4. Return the minimum among all choices.

### Complexity

| Metric | Value |
|----------|----------|
| Time | Exponential |
| Space | O(n) recursion stack |

### Limitations

- Explores many duplicate states.
- Extremely slow for large inputs.
- Not suitable for interview constraints.

---

## Optimized Approach

### Idea

Use a greedy strategy to track:

- Current jump range
- Farthest reachable index

Whenever the current range ends, perform a jump and extend the range.

### Algorithm

1. Initialize:
   - `jumps = 0`
   - `currentEnd = 0`
   - `farthestReach = 0`

2. Iterate from index `0` to `n - 2`.

3. Update:

   ```text
   farthestReach = max(farthestReach, i + nums[i])
   ```

4. If `i == currentEnd`:
   - Increment jump count.
   - Move boundary to `farthestReach`.

5. Continue until reaching the end.

6. Return total jumps.

---

### Why It Works

For every jump, we evaluate all positions reachable using the current number of jumps.

Among them, we select the position that provides the maximum future reach.

Since every possible position in the current layer is considered before increasing the jump count, the solution behaves exactly like BFS and guarantees the minimum number of jumps.

---

### Complexity

| Metric | Value |
|----------|----------|
| Time | O(n) |
| Space | O(1) |

---

## Edge Cases

### Empty Input

```text
[]
```

Not valid according to problem constraints.

---

### Single Element

```text
[0]
```

Already at destination.

Answer:

```text
0
```

---

### All Ones

```text
[1,1,1,1,1]
```

Must move one step each time.

Answer:

```text
4
```

---

### Large Jump At Start

```text
[10,1,1,1,1]
```

Can directly reach the end.

Answer:

```text
1
```

---

### Multiple Optimal Paths

```text
[2,3,1,1,4]
```

Several paths exist.

Greedy still finds the minimum jump count.

---

### Large Inputs

```text
Length = 10^4+
```

Greedy remains efficient due to linear traversal.

---

## Dry Run

### Input

```text
nums = [2,3,1,1,4]
```

### Initialization

| Variable | Value |
|-----------|---------|
| jumps | 0 |
| currentEnd | 0 |
| farthestReach | 0 |

---

### Iteration Table

| i | nums[i] | farthestReach | currentEnd | jumps |
|---|---|---|---|---|
| 0 | 2 | 2 | 0 | 0 |
| Jump Trigger | - | 2 | 2 | 1 |
| 1 | 3 | 4 | 2 | 1 |
| 2 | 1 | 4 | 2 | 1 |
| Jump Trigger | - | 4 | 4 | 2 |

Destination reached.

Answer:

```text
2
```

---

## Common Mistakes

### Mistake 1

Incrementing jumps at every index.

Incorrect because a jump should only occur when the current reachable range is exhausted.

---

### Mistake 2

Iterating until `n - 1`.

The last index does not need expansion.

Correct loop:

```text
for i := 0; i < n-1; i++
```

---

### Mistake 3

Confusing Jump Game I and Jump Game II.

Jump Game I:

```text
Can we reach the end?
```

Jump Game II:

```text
Minimum jumps needed?
```

---

### Mistake 4

Using Dynamic Programming unnecessarily.

DP works but is slower than the greedy solution.

---

## Interview Discussion

### Expected Progression

A strong candidate should discuss:

1. Recursive exploration
2. DP optimization
3. Greedy observation
4. BFS-layer interpretation
5. O(n) solution

---

### Key Insight To Communicate

Instead of choosing where to jump immediately:

- Explore the entire current reachable range.
- Determine the best future reach.
- Commit to a jump only when necessary.

This produces the minimum jump count.

---

## Follow-up Questions

### Follow-up 1

What if reaching the end is not guaranteed?

Answer:

Need additional checks to detect unreachable states.

---

### Follow-up 2

Can this be solved using Dynamic Programming?

Answer:

Yes.

Time complexity is typically O(n²).

---

### Follow-up 3

Why is the greedy solution optimal?

Answer:

Because each jump corresponds to a BFS level, and BFS guarantees the shortest path in an unweighted graph.

---

### Follow-up 4

Can we reconstruct the actual path?

Answer:

Yes.

Store parent indices while tracking transitions.

---

## Real World Applications

### Network Routing

Choosing the farthest beneficial hop while minimizing transitions.

---

### Logistics Optimization

Minimizing transfers between transportation hubs.

---

### Wireless Communication

Reducing relay hops between nodes.

---

### Game Development

Calculating minimum moves to reach a destination.

---

### Distributed Systems

Minimizing intermediary nodes in message propagation.

---

## Related Problems

### Easy

- Jump Game (55)

### Medium

- Minimum Number of Refueling Stops (871)
- Gas Station (134)
- Partition Labels (763)

### Hard

- Frog Jump (403)
- Jump Game IV (1345)
- Jump Game V (1340)
- Jump Game VI (1696)

---

## Key Takeaway

This problem is one of the most important greedy interview questions.

The core idea is:

> Treat reachable indices as BFS layers and use greedy range expansion to maximize future reach.

This transforms an exponential search problem into a linear-time optimal solution.