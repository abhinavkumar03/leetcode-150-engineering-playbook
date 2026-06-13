# Interview Notes — 380. Insert Delete GetRandom O(1)

# What Interviewer Is Testing

This problem is not primarily about coding.

It is a **data structure design problem**.

Interviewers want to evaluate whether you can:

- Choose the correct data structure for a requirement.
- Combine multiple data structures effectively.
- Analyze tradeoffs.
- Optimize from a naive solution.
- Maintain complexity guarantees across operations.
- Explain design decisions clearly.

---

## Core Skills Being Evaluated

### 1. Data Structure Selection

Can you recognize that:

```text
HashMap
```

provides:

```text
O(1) lookup
O(1) insertion
O(1) deletion
```

but cannot provide:

```text
O(1) random selection
```

?

---

### 2. Understanding Array Strengths

Can you recognize that:

```text
Array/List
```

provides:

```text
O(1) append
O(1) random access
```

but deletion is:

```text
O(n)
```

?

---

### 3. Combining Structures

Strong candidates quickly realize:

```text
HashMap + Array
```

is required.

This is often the key interview breakthrough.

---

### 4. Complexity Awareness

Interviewers expect candidates to challenge their own solution:

```text
Can delete really be O(1)?
```

The ability to identify hidden complexity is important.

---

### 5. Engineering Thinking

This problem simulates real-world engineering.

Many production systems use:

```text
Storage Structure
+
Index Structure
```

to achieve performance goals.

Examples:

- Database indexes
- Search engines
- Distributed caches
- In-memory data stores

---

# Typical Follow-up Questions

## Follow-up #1

### Why Not Use Only a HashMap?

Expected Answer:

```text
HashMap does not support O(1) random access.
```

Selecting a random key requires:

```text
O(n)
```

iteration.

---

## Follow-up #2

### Why Not Use Only an Array?

Expected Answer:

```text
Deletion requires shifting elements.
```

Complexity becomes:

```text
O(n)
```

---

## Follow-up #3

### Why Does Swapping Work?

Expected Answer:

We don't care about element order.

Therefore:

```text
Swap target with last element
Remove last element
```

avoids shifting.

---

## Follow-up #4

### Does This Preserve Insertion Order?

Expected Answer:

No.

Example:

```text
[10,20,30]

remove(20)

[10,30]
```

The original order changes.

---

## Follow-up #5

### Can Duplicates Be Supported?

Expected Answer:

Yes.

Use:

```text
HashMap<Value, Set<Indices>>
```

instead of:

```text
HashMap<Value, Index>
```

Related problem:

```text
LeetCode 381
```

---

## Follow-up #6

### Is Random Selection Truly Uniform?

Expected Answer:

Yes.

Every element occupies exactly one position in the array.

Probability:

```text
1 / n
```

for every element.

---

## Follow-up #7

### What Happens at Scale?

Expected Answer:

Storage grows:

```text
O(n)
```

Operations remain:

```text
O(1)
```

average.

---

# Optimization Journey

Interviewers often want to hear the progression.

---

## Version 1 — Naive Array

Store everything in an array.

### Operations

```text
Insert     O(1)
Delete     O(n)
Random     O(1)
```

Problem:

```text
Delete too slow
```

---

## Version 2 — HashMap Only

Store everything in a HashMap.

### Operations

```text
Insert     O(1)
Delete     O(1)
Random     O(n)
```

Problem:

```text
Random too slow
```

---

## Version 3 — HashMap + Array

Store:

```text
Array
+
HashMap
```

### Operations

```text
Insert     O(1)
Delete     O(1)?
Random     O(1)
```

Question:

```text
How do we keep delete O(1)?
```

---

## Version 4 — Swap and Remove

Key breakthrough:

```text
Swap target with last element.
Pop last element.
```

Result:

```text
Insert     O(1)
Delete     O(1)
Random     O(1)
```

Final solution achieved.

---

# Whiteboard Strategy

If solving on a whiteboard:

---

## Step 1

State requirements.

```text
Need O(1) for:
Insert
Delete
GetRandom
```

---

## Step 2

Analyze common structures.

Create table:

| Structure | Insert | Delete | Random |
|------------|---------|---------|---------|
| Array | O(1) | O(n) | O(1) |
| HashMap | O(1) | O(1) | O(n) |

---

## Step 3

Explain why neither works alone.

---

## Step 4

Combine both structures.

Draw:

```text
Array

[10,20,30]

HashMap

10 -> 0
20 -> 1
30 -> 2
```

---

## Step 5

Explain deletion visually.

```text
Before

[10,20,30]

Swap

[10,30,20]

Pop

[10,30]
```

This visual explanation often impresses interviewers.

---

# Communication Tips

## Good Explanation

Say:

> "The challenge isn't insertion or random access. The real challenge is maintaining O(1) deletion."

This shows problem understanding.

---

## Explain Tradeoffs

Mention:

```text
Extra memory
for faster operations.
```

Interviewers appreciate explicit tradeoff discussion.

---

## Explain Why Order Doesn't Matter

A common interview question:

```text
Why are we allowed to swap?
```

Answer:

```text
The problem never requires preserving order.
```

---

## Discuss Average Complexity

HashMap operations are:

```text
Average O(1)
```

not strict worst-case O(1).

Showing this nuance demonstrates maturity.

---

# Senior-Level Discussion Points

A senior engineer should be able to discuss beyond the coding solution.

---

## Memory Tradeoffs

We store:

```text
Array
+
HashMap
```

which duplicates information.

Cost:

```text
O(n)
```

additional memory.

Benefit:

```text
O(1)
```

operations.

---

## Thread Safety

Current solution is not thread-safe.

Potential improvements:

```text
Mutex
RWLock
Concurrent Collections
```

depending on language.

---

## Production Considerations

Potential concerns:

- Memory fragmentation
- Resize costs
- Hash collisions
- Random generator quality

---

## API Design

Could expose:

```text
Contains()
Size()
Clear()
```

in production implementations.

---

## Testing Strategy

Verify:

### Insert

```text
New value → true
Duplicate → false
```

---

### Remove

```text
Existing value → true
Missing value → false
```

---

### Random

Validate:

```text
Returned value exists.
```

and distribution appears uniform.

---

# FAANG-Level Variations

Interviewers frequently extend this problem.

---

## Variation 1

### Duplicates Allowed

Problem:

```text
LeetCode 381
```

Modification:

```text
HashMap<Value, Set<Indices>>
```

---

## Variation 2

### Weighted Random Selection

Each element has a weight.

Need:

```text
Probability proportional to weight.
```

Possible solutions:

- Prefix Sum
- Fenwick Tree
- Segment Tree

---

## Variation 3

### Random Removal

Support:

```text
removeRandom()
```

Can be implemented in:

```text
O(1)
```

using the same structure.

---

## Variation 4

### Thread-Safe Randomized Set

Support:

```text
Multiple readers
Multiple writers
```

Topics:

- Locking
- Atomic operations
- Concurrent data structures

---

## Variation 5

### Distributed Randomized Set

Data stored across servers.

Discussion topics:

- Sharding
- Consistent hashing
- Distributed coordination
- Random sampling across shards

---

# Interview Summary

The most important insight is:

```text
HashMap
+
Array
+
Swap-and-Remove
=
Insert O(1)
Delete O(1)
GetRandom O(1)
```

Candidates who discover the swap-and-remove technique quickly typically perform very well on this problem.

The interviewer is less interested in syntax and more interested in whether you can systematically design a data structure that satisfies multiple competing requirements.