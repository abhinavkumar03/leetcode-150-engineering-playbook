# Insert Delete GetRandom O(1)

## Problem Statement

Design a data structure that supports the following operations in average O(1) time:

- insert(val): Inserts an item into the set if not present.
- remove(val): Removes an item from the set if present.
- getRandom(): Returns a random element from the current set. Each element must have the same probability of being returned.

Implement the `RandomizedSet` class:

```text
RandomizedSet()

bool insert(int val)

bool remove(int val)

int getRandom()
```

---

## Difficulty

Medium

---

## Tags

- Array
- Hash Table
- Design
- Randomized
- Data Structure

---

## Pattern

### Primary Pattern

Hash Map + Dynamic Array

### Secondary Pattern

Data Structure Design

---

## Intuition

At first glance, this looks simple.

We need:

| Operation | Required Complexity |
|------------|-------------------|
| Insert | O(1) |
| Delete | O(1) |
| Random Access | O(1) |

Let's evaluate common data structures.

### HashMap Only

Insertion → O(1)

Deletion → O(1)

Random Selection → O(n)

Not acceptable.

### Array Only

Insertion → O(1)

Random Selection → O(1)

Deletion → O(n)

Not acceptable.

### Combined Approach

Use:

- Array/List to support random access
- HashMap to track element indices

Now we can potentially achieve O(1) for all operations.

The only remaining challenge is deletion.

---

## Key Observation

Deleting an element from the middle of an array normally requires shifting elements.

Example:

```text
[10, 20, 30, 40]

Remove 20

[10, 30, 40]
```

This shift costs O(n).

Instead:

### Swap with Last Element

```text
[10, 20, 30, 40]

Remove 20

Swap 20 and 40

[10, 40, 30, 20]

Remove last element

[10, 40, 30]
```

Update the HashMap accordingly.

This avoids shifting and makes deletion O(1).

This is the core trick behind the entire solution.

---

## Brute Force Approach

Store elements in a list.

### Algorithm

Insert:

1. Check if value exists.
2. Append to list.

Remove:

1. Find value by scanning list.
2. Remove element.
3. Shift remaining elements.

GetRandom:

1. Generate random index.
2. Return element.

### Complexity

| Operation | Complexity |
|------------|------------|
| Insert | O(1) |
| Remove | O(n) |
| GetRandom | O(1) |

### Overall

```text
Time: O(n) worst case
Space: O(n)
```

### Limitations

- Deletion becomes expensive.
- Large datasets perform poorly.
- Does not satisfy interview constraints.

---

## Optimized Approach

Use two structures:

```text
Array/List
+
HashMap<Value, Index>
```

Example:

```text
nums = [10, 20, 30]

map = {
 10 -> 0
 20 -> 1
 30 -> 2
}
```

### Insert

If value already exists:

```text
Return false
```

Otherwise:

1. Append value to array.
2. Store index in HashMap.

### Remove

Suppose:

```text
nums = [10, 20, 30, 40]

map:
10 -> 0
20 -> 1
30 -> 2
40 -> 3
```

Remove 20.

Step 1:

Find index.

```text
idx = 1
```

Step 2:

Take last element.

```text
last = 40
```

Step 3:

Move last element into removed position.

```text
nums = [10, 40, 30, 40]
```

Step 4:

Update index.

```text
40 -> 1
```

Step 5:

Remove last element.

```text
nums = [10, 40, 30]
```

Step 6:

Delete 20 from HashMap.

### GetRandom

Generate:

```text
randomIndex = rand(0, size-1)
```

Return:

```text
nums[randomIndex]
```

Because every element occupies exactly one position in the array, each has equal probability.

---

## Why It Works

The array provides:

```text
O(1) random access
```

The HashMap provides:

```text
O(1) lookup
O(1) existence check
O(1) index retrieval
```

The swap-and-remove technique ensures:

```text
O(1) deletion
```

Together they satisfy all requirements.

---

## Complexity

| Operation | Time |
|------------|------|
| Insert | O(1) |
| Remove | O(1) |
| GetRandom | O(1) |

### Space Complexity

```text
O(n)
```

for storing elements in:

- Array
- HashMap

---

## Edge Cases

### Empty Structure

```text
[]
```

Calling getRandom is not allowed by problem constraints.

---

### Single Element

```text
insert(5)

[5]
```

Random always returns:

```text
5
```

---

### Duplicate Insertions

```text
insert(10)
insert(10)
```

Second insertion:

```text
false
```

---

### Removing Missing Values

```text
remove(99)
```

Value does not exist.

```text
false
```

---

### Negative Values

```text
insert(-5)
insert(-10)
```

HashMap and Array handle them normally.

---

### Large Inputs

```text
100000+ operations
```

Efficient because every operation remains O(1) average.

---

## Dry Run

### Operations

```text
insert(1)
insert(2)
remove(1)
insert(3)
getRandom()
```

### State Table

| Step | Operation | Array | Map |
|--------|----------|--------|------|
| 1 | insert(1) | [1] | {1:0} |
| 2 | insert(2) | [1,2] | {1:0,2:1} |
| 3 | remove(1) | [2] | {2:0} |
| 4 | insert(3) | [2,3] | {2:0,3:1} |
| 5 | getRandom() | [2,3] | {2:0,3:1} |

Possible output:

```text
2
```

or

```text
3
```

Both are equally likely.

---

## Common Mistakes

### Forgetting to Update Last Element Index

Wrong:

```text
Swap
Remove
```

Without updating HashMap.

This causes stale indices.

---

### Removing Before Swapping

Wrong order can lose data.

Always:

```text
Find index
Swap
Update index
Remove last
Delete mapping
```

---

### Using ArrayList.remove(index)

In Java:

```java
list.remove(index);
```

causes shifting.

Complexity becomes O(n).

---

### Using Random on HashMap Keys

HashMaps do not support O(1) random access.

You still need an array.

---

## Interview Discussion

A strong candidate should explain:

1. Why HashMap alone fails.
2. Why Array alone fails.
3. Why deletion is the main challenge.
4. How swap-and-remove solves deletion.
5. Why random access remains uniform.

Expected progression:

```text
Brute Force
→ HashMap
→ HashMap + Array
→ Swap-and-Remove
→ O(1) Design
```

---

## Follow-up Questions

### Can duplicates be allowed?

Yes.

Use:

```text
HashMap<Value, Set<Indices>>
```

Related to:

```text
LeetCode 381
Insert Delete GetRandom O(1) - Duplicates Allowed
```

---

### Can weighted random selection be supported?

Yes.

Use:

- Prefix sums
- Segment Tree
- Binary Indexed Tree

---

### Can this be made thread-safe?

Yes.

Possible approaches:

- Mutex
- Read-Write Lock
- Concurrent Data Structures

---

### What happens if operations become distributed?

Need:

- Sharding
- Consistent Hashing
- Distributed Coordination

---

## Real World Applications

### In-Memory Caches

Random eviction strategies.

---

### Online Gaming

Random player selection.

---

### Recommendation Systems

Random candidate sampling.

---

### Load Testing

Random request generation.

---

### Distributed Systems

Random node selection.

---

## Related Problems

### Easy

- 217. Contains Duplicate
- 706. Design HashMap

### Medium

- 380. Insert Delete GetRandom O(1)
- 381. Insert Delete GetRandom O(1) - Duplicates Allowed
- 146. LRU Cache

### Hard

- 432. All O`one Data Structure

---

## Takeaway

This problem is a classic example of combining multiple data structures to satisfy strict performance requirements.

Key insight:

```text
Array + HashMap + Swap-and-Remove
=
Insert O(1)
Delete O(1)
GetRandom O(1)
```

Understanding this technique is extremely valuable because it appears repeatedly in advanced system design and data structure interview questions.