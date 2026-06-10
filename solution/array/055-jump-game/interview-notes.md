# Interview Notes — Jump Game

## What Interviewer Is Testing

This problem looks simple but evaluates several important engineering and interview skills.

### 1. Pattern Recognition

Can you identify that this is a **Greedy** problem rather than:

- Backtracking
- BFS
- DFS
- Dynamic Programming

Strong candidates quickly recognize that:

> The exact path is irrelevant.
> Only the farthest reachable position matters.

---

### 2. Optimization Ability

Interviewers often expect candidates to progress through:

```text
Brute Force
    ↓
Memoization
    ↓
Dynamic Programming
    ↓
Greedy
```

The key evaluation is:

> Can you reduce unnecessary state and derive a linear solution?

---

### 3. Complexity Analysis

Candidates should be able to explain:

| Approach | Time | Space |
|-----------|-----------|-----------|
| Recursive | O(2ⁿ) | O(n) |
| Memoized | O(n²) | O(n) |
| DP | O(n²) | O(n) |
| Greedy | O(n) | O(1) |

---

### 4. Correctness Reasoning

Interviewers want proof, not just code.

You should explain:

```text
farthestReach
```

always represents:

```text
The maximum reachable index discovered so far.
```

If:

```text
currentIndex > farthestReach
```

then:

```text
currentIndex is unreachable
```

and reaching the destination becomes impossible.

---

## Typical Follow-up Questions

### Follow-up 1

#### Why does the greedy solution work?

Expected Answer:

```text
Every reachable position contributes to
extending the reachable range.

We only need to know the farthest
reachable index, not the exact path.
```

---

### Follow-up 2

#### Can this be solved using Dynamic Programming?

Yes.

Define:

```text
dp[i]
```

as:

```text
Can index i reach the end?
```

However:

```text
Time: O(n²)
Space: O(n)
```

which is inferior to Greedy.

---

### Follow-up 3

#### Can you solve it from right to left?

Yes.

Use a backward greedy strategy.

Initialize:

```text
goal = lastIndex
```

Traverse backwards:

```text
if i + nums[i] >= goal
    goal = i
```

Final condition:

```text
goal == 0
```

---

### Follow-up 4

#### What changes if we need minimum jumps?

This becomes:

```text
Jump Game II
```

Different objective:

```text
Reachability
    vs
Minimum Moves
```

---

### Follow-up 5

#### Can you return the actual jump path?

Yes.

Store:

```text
parent[index]
```

while exploring reachable positions.

This increases memory usage.

---

### Follow-up 6

#### What if negative jumps are allowed?

The problem changes significantly.

Now:

```text
Array → Graph
```

Possible techniques:

- BFS
- DFS
- Graph Reachability
- Shortest Path Variants

---

## Optimization Journey

### Stage 1 — Brute Force

Try every possible jump.

Example:

```text
Index 0
 ├── Jump 1
 ├── Jump 2
 └── Jump 3
```

Repeated subproblems appear quickly.

Complexity:

```text
O(2ⁿ)
```

---

### Stage 2 — Memoization

Cache results:

```text
canReach(index)
```

Avoid recomputation.

Complexity improves but remains expensive.

```text
O(n²)
```

---

### Stage 3 — Dynamic Programming

Build reachability information.

Example:

```text
dp[i] = reachable
```

Still requires checking many future states.

```text
O(n²)
```

---

### Stage 4 — Greedy

Observation:

```text
We never need every reachable position.

We only need the farthest reachable one.
```

Track:

```text
farthestReach
```

Complexity:

```text
O(n)
```

---

## Whiteboard Strategy

### Step 1

Restate the problem.

Example:

> We need to determine whether the last index can be reached from index 0.

---

### Step 2

Discuss brute force.

Show interviewer that you understand all possibilities.

---

### Step 3

Identify redundancy.

Explain:

```text
Many paths lead to the same index.
```

---

### Step 4

Introduce reachability.

Say:

> Instead of tracking paths, I will track the farthest reachable position.

This is the turning point.

---

### Step 5

Draw Reach Expansion

Example:

```text
[2,3,1,1,4]

Index 0 → Reach 2

Reachable:
0 1 2

Index 1 extends reach to 4

Reachable:
0 1 2 3 4
```

Visual explanations score highly in interviews.

---

### Step 6

Provide Complexity

```text
Time: O(n)
Space: O(1)
```

Always justify both.

---

## Communication Tips

### Good Explanation

> At each position, I update the farthest index that can be reached. If I ever encounter an index beyond that boundary, I know it is impossible to proceed.

---

### Stronger Explanation

> The problem is fundamentally a reachability problem. The exact sequence of jumps is irrelevant. The only information needed is the maximum reachable boundary.

---

### Avoid Saying

```text
It just works.
```

or

```text
This is the standard solution.
```

Interviewers expect reasoning.

---

### Mention Invariants

A powerful statement:

> Before processing index i, every index up to farthestReach is guaranteed reachable.

This demonstrates algorithmic maturity.

---

## Senior-Level Discussion Points

### Invariant-Based Reasoning

Invariant:

```text
All indices ≤ farthestReach
are reachable.
```

This remains true throughout execution.

---

### Early Exit Optimization

Once:

```text
farthestReach >= lastIndex
```

we can immediately return:

```text
true
```

No need to continue scanning.

---

### Failure Detection

Failure is discovered immediately when:

```text
currentIndex > farthestReach
```

This avoids unnecessary work.

---

### Production Considerations

For very large arrays:

- Linear runtime scales well
- Constant memory usage
- Cache-friendly traversal
- No recursion risk

---

### Why This Is a Greedy Problem

Greedy choice:

```text
Maintain the largest reachable boundary.
```

Local decision:

```text
Expand reach whenever possible.
```

Global result:

```text
Determine reachability of destination.
```

---

## FAANG-Level Variations

### Variation 1

Return the actual path.

Example:

```text
0 → 1 → 4
```

Requires predecessor tracking.

---

### Variation 2

Minimum number of jumps.

Leads to:

```text
Jump Game II
```

---

### Variation 3

Count all valid ways to reach the end.

Potential solutions:

- Dynamic Programming
- Recursion + Memoization

---

### Variation 4

Weighted jumps.

Now each jump has a cost.

Potential techniques:

- Dijkstra
- Dynamic Programming

---

### Variation 5

Bidirectional jumps

Example:

```text
+2
-1
+3
```

Now the problem becomes graph traversal.

Techniques:

- BFS
- DFS
- Cycle Detection

---

### Variation 6

Streaming Input

Array arrives incrementally.

Discussion topics:

- Online algorithms
- Reachability maintenance
- State persistence

---

## Interview Sound-Bite

A concise answer suitable for interviews:

> Jump Game is a Greedy reachability problem. Instead of exploring every possible jump, we maintain the farthest reachable index. If we ever reach an index beyond that boundary, the destination is unreachable. Otherwise, continuously expanding the boundary guarantees an O(n) time and O(1) space solution.
