# Jump Game

## Problem Statement

You are given an integer array `nums`.

You are initially positioned at the first index, and each element in the array represents your maximum jump length at that position.

Determine whether you can reach the last index.

### Example 1

Input:

```text
nums = [2,3,1,1,4]
```

Output:

```text
true
```

Explanation:

```text
Jump 1 step from index 0 to 1.
Then jump 3 steps to the last index.
```

### Example 2

Input:

```text
nums = [3,2,1,0,4]
```

Output:

```text
false
```

Explanation:

```text
You will always arrive at index 3.
Since nums[3] = 0, you cannot move further.
```

---

## Difficulty

Medium

---

## Tags

- Array
- Greedy
- Dynamic Programming

---

## Pattern

### Primary Pattern

Greedy

### Secondary Pattern

Reachability Tracking

---

## Intuition

At first glance, the problem looks like a path-finding or Dynamic Programming problem because there are many possible jumps from each position.

A common mistake is attempting to explore every jump possibility.

Instead, observe that:

> We do not care how we reach an index.
> We only care whether that index is reachable.

If we continuously track the farthest position reachable so far, we can determine whether the end is reachable without exploring all paths.

---

## Key Observation

For every index:

```text
farthestReach = max(
    farthestReach,
    currentIndex + nums[currentIndex]
)
```

If we ever encounter:

```text
currentIndex > farthestReach
```

then that index is unreachable.

Therefore, reaching the end becomes impossible.

---

## Brute Force Approach

### Idea

Try every possible jump recursively.

For each position:

- Jump 1 step
- Jump 2 steps
- ...
- Jump nums[i] steps

Continue until:

- Last index is reached
- No moves remain

### Algorithm

1. Start from index 0.
2. Recursively try all valid jumps.
3. Return true if any path reaches the end.
4. Return false otherwise.

### Complexity

#### Time Complexity

```text
O(2^n)
```

Worst-case exponential exploration.

#### Space Complexity

```text
O(n)
```

Recursion stack.

### Limitations

- Massive repeated calculations
- Exponential runtime
- Not suitable for interview constraints

---

## Optimized Approach

### Idea

Track the farthest index that can currently be reached.

As we move through the array:

- If current index is reachable
- Extend the reachable range

Eventually:

- Reach the end → return true
- Encounter unreachable index → return false

### Algorithm

1. Initialize:

```text
farthestReach = 0
```

2. Traverse array from left to right.

3. If:

```text
i > farthestReach
```

return false.

4. Update:

```text
farthestReach = max(
    farthestReach,
    i + nums[i]
)
```

5. If farthestReach reaches or exceeds last index:

```text
return true
```

6. Finish traversal.

### Why It Works

Suppose:

```text
farthestReach = 7
```

Then every index from:

```text
0 → 7
```

is reachable.

When processing any reachable index, we may extend the reachable region.

Therefore, maintaining only the farthest reachable position is sufficient.

No additional state is required.

### Complexity

#### Time Complexity

```text
O(n)
```

Single traversal.

#### Space Complexity

```text
O(1)
```

Constant extra memory.

---

## Edge Cases

### Empty Input

```text
[]
```

Typically excluded by constraints.

---

### Single Element

```text
[0]
```

Already at destination.

Output:

```text
true
```

---

### Contains Zeroes But Reachable

```text
[2,0,1]
```

Output:

```text
true
```

---

### Blocking Zero

```text
[3,2,1,0,4]
```

Output:

```text
false
```

---

### Large Jump at Start

```text
[10,0,0,0]
```

Output:

```text
true
```

---

### Large Input

```text
10000+ elements
```

Greedy solution remains efficient.

---

## Dry Run

Input:

```text
nums = [2,3,1,1,4]
```

Target Index:

```text
4
```

| Index | nums[i] | Farthest Reach Before | Farthest Reach After |
|---------|---------|---------|---------|
| 0 | 2 | 0 | 2 |
| 1 | 3 | 2 | 4 |
| 2 | 1 | 4 | 4 |
| 3 | 1 | 4 | 4 |
| 4 | 4 | 4 | 8 |

Result:

```text
true
```

Since:

```text
farthestReach >= lastIndex
```

---

### Failure Example

Input:

```text
nums = [3,2,1,0,4]
```

| Index | nums[i] | Farthest Reach |
|---------|---------|---------|
| 0 | 3 | 3 |
| 1 | 2 | 3 |
| 2 | 1 | 3 |
| 3 | 0 | 3 |
| 4 | 4 | unreachable |

At:

```text
i = 4
```

we have:

```text
4 > farthestReach
```

Therefore:

```text
false
```

---

## Common Mistakes

### Mistake 1

Trying every jump recursively.

Leads to:

```text
O(2^n)
```

complexity.

---

### Mistake 2

Using Dynamic Programming unnecessarily.

DP works but consumes extra memory.

Greedy is simpler and optimal.

---

### Mistake 3

Not checking reachability before processing index.

Incorrect:

```text
farthestReach =
max(farthestReach, i + nums[i])
```

without verifying:

```text
i <= farthestReach
```

---

### Mistake 4

Confusing minimum jumps with reachability.

This problem only asks:

```text
Can we reach?
```

Not:

```text
How many jumps?
```

That is a different problem:

```text
Jump Game II
```

---

## Interview Discussion

### Expected Progression

Interviewer often expects:

1. Recursive solution
2. Memoized solution
3. Dynamic Programming
4. Greedy optimization

---

### Key Insight To Communicate

State clearly:

> I don't need to know the exact path.
> I only need to know the farthest reachable position.

This demonstrates strong optimization thinking.

---

### What Interviewers Evaluate

- Pattern recognition
- Greedy reasoning
- Complexity reduction
- Correctness proof
- Communication skills

---

## Follow-up Questions

### Follow-up 1

What if we need the minimum number of jumps?

Answer:

```text
Jump Game II
```

---

### Follow-up 2

Can you return the actual path?

Requires storing predecessor information.

---

### Follow-up 3

Can this be solved from right to left?

Yes.

A backward greedy solution also exists.

---

### Follow-up 4

What if negative jumps are allowed?

Problem transforms into a graph reachability problem.

---

## Real World Applications

### Network Reachability

Determining whether a destination node can be reached through available connections.

---

### Workflow Execution

Checking whether a process can reach completion based on allowed transitions.

---

### Game Development

Character movement systems with variable jump ranges.

---

### Resource Planning

Determining whether future milestones remain attainable given current capacity.

---

## Related Problems

### Easy

- Two Sum
- Best Time to Buy and Sell Stock

### Medium

- Jump Game II
- Gas Station
- Partition Labels
- Maximum Subarray

### Hard

- Minimum Number of Refueling Stops
- Candy
- Frog Jump

---