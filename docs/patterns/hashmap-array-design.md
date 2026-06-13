# Hash Map + Array Design Pattern

## Pattern Definition

The Hash Map + Array Design Pattern combines:

```text
HashMap<Key, Index>
+
Dynamic Array/List
```

to achieve operations that require both:

- Fast lookup
- Fast insertion
- Fast deletion
- Fast random access

This pattern is commonly used in data structure design problems where a single data structure cannot satisfy all complexity requirements.

---

## Motivation

Different data structures excel at different operations.

### HashMap Strengths

| Operation | Complexity |
|------------|------------|
| Insert | O(1) |
| Delete | O(1) |
| Lookup | O(1) |
| Random Access | O(n) |

---

### Dynamic Array Strengths

| Operation | Complexity |
|------------|------------|
| Append | O(1) |
| Random Access | O(1) |
| Lookup | O(n) |
| Middle Delete | O(n) |

---

Neither structure alone satisfies all requirements.

Combining them allows us to leverage the strengths of both.

---

## Core Structure

```text
Array

Index:   0   1   2
Value:  10  20  30
```

```text
HashMap

10 -> 0
20 -> 1
30 -> 2
```

### Responsibilities

#### Array

Provides:

```text
O(1) random access
O(1) append
```

---

#### HashMap

Provides:

```text
O(1) lookup
O(1) existence check
O(1) index retrieval
```

---

## Key Technique: Swap-and-Remove

The major challenge is deletion.

### Problem

Removing from the middle of an array causes shifting.

Example:

```text
[10,20,30,40]

remove(20)

[10,30,40]
```

Complexity:

```text
O(n)
```

---

### Solution

Swap target with the last element.

Before:

```text
[10,20,30,40]
```

Swap:

```text
[10,40,30,20]
```

Pop last:

```text
[10,40,30]
```

Update HashMap:

```text
40 -> 1
```

Deletion becomes:

```text
O(1)
```

---

## Recognition Signals

Use this pattern when you see:

### Signal #1

Requirements include:

```text
Insert
Delete
Lookup
Random Access
```

in constant time.

---

### Signal #2

Problem statement contains:

```text
Design a data structure
```

---

### Signal #3

Need:

```text
Fast retrieval
+
Fast mutation
```

simultaneously.

---

### Signal #4

Element ordering is not important.

This is often the strongest indicator that swap-and-remove is allowed.

---

### Signal #5

Operations must be:

```text
Average O(1)
```

---

## Generic Template

### Insert

```text
if value exists:
    return false

append value to array

store:
value -> lastIndex

return true
```

---

### Remove

```text
find target index

find last value

move last value
to target index

update mapping

remove last element

delete target mapping
```

---

### Random

```text
randomIndex = rand(0, size-1)

return array[randomIndex]
```

---

## Complexity

| Operation | Complexity |
|------------|------------|
| Insert | O(1) |
| Delete | O(1) |
| Lookup | O(1) |
| Random Access | O(1) |
| Space | O(n) |

---

## Common Pitfalls

### Pitfall #1

Forgetting to update the moved element's index.

Wrong:

```text
Swap
Pop
```

Correct:

```text
Swap
Update Index
Pop
```

---

### Pitfall #2

Trying to preserve order.

Preserving order generally forces:

```text
O(n)
```

deletion.

---

### Pitfall #3

Using only a HashMap.

Cannot support:

```text
O(1) random selection
```

---

### Pitfall #4

Using only an Array.

Cannot support:

```text
O(1) lookup
```

---

### Pitfall #5

Removing directly from the middle of the array.

Causes element shifting.

---

## Engineering Applications

This pattern appears in:

### In-Memory Caches

Store values and maintain fast lookup indexes.

---

### Search Engines

Document identifiers mapped to storage positions.

---

### Recommendation Systems

Random candidate sampling from indexed collections.

---

### Online Gaming

Random player or item selection.

---

### Distributed Systems

Efficient membership tracking with random node selection.

---

## Related Problems

### Easy

- 217. Contains Duplicate
- 706. Design HashMap
- 705. Design HashSet

---

### Medium

- 380. Insert Delete GetRandom O(1)
- 381. Insert Delete GetRandom O(1) — Duplicates Allowed
- 146. LRU Cache

---

### Hard

- 432. All O`one Data Structure
- 460. LFU Cache

---

# Problem Catalog

## LeetCode 380 — Insert Delete GetRandom O(1)

### Pattern Usage

Uses:

```text
HashMap<Value, Index>
+
Dynamic Array
```

to support:

```text
Insert      O(1)
Delete      O(1)
GetRandom   O(1)
```

### Key Insight

```text
Swap-and-Remove
```

eliminates the shifting cost of array deletion.

### Complexity

| Operation | Complexity |
|------------|------------|
| Insert | O(1) |
| Delete | O(1) |
| GetRandom | O(1) |
| Space | O(n) |

### Interview Importance

One of the most frequently asked data structure design problems.

Tests:

- Data structure selection
- Complexity analysis
- Design tradeoffs
- Optimization thinking

### Takeaway

```text
HashMap
+
Array
+
Swap-and-Remove
=
O(1) Insert
O(1) Delete
O(1) Random
```

This problem is considered the canonical example of the Hash Map + Array Design Pattern.