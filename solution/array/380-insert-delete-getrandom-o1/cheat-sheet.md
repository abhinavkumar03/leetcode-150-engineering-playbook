# Cheat Sheet — 380. Insert Delete GetRandom O(1)

---

# Pattern Summary

## Pattern Name

```text
Hash Map + Dynamic Array
```

Also Known As:

```text
Indexed Hash Set
Randomized Data Structure
Swap-and-Remove Pattern
```

---

## Core Idea

Combine:

```text
Array/List
+
HashMap<Value, Index>
```

to achieve:

```text
Insert      O(1)
Delete      O(1)
GetRandom   O(1)
```

The key trick:

```text
Swap target element with last element
Remove last element
Update index mapping
```

instead of shifting elements.

---

# Recognition Signals

Look for problems requiring:

### Signal #1

Multiple operations must be:

```text
O(1)
```

---

### Signal #2

Requirements include:

```text
Insert
Delete
Lookup
Random Access
```

simultaneously.

---

### Signal #3

Need:

```text
Fast lookup
+
Fast random retrieval
```

---

### Signal #4

Order preservation is NOT required.

This is usually the clue that:

```text
Swap-and-Remove
```

is allowed.

---

### Signal #5

Problem contains wording like:

```text
Design a data structure
Average O(1)
```

---

# Data Structure Layout

```text
values

Index:   0   1   2
Value:  10  20  30
```

HashMap:

```text
10 -> 0
20 -> 1
30 -> 2
```

---

# Insert Template

### Goal

Add value if it doesn't already exist.

### Steps

```text
if value exists:
    return false

append value to array

store:
value -> last index

return true
```

### Complexity

```text
Time  : O(1)
Space : O(1)
```

---

# Remove Template

### Goal

Delete value in O(1).

### Steps

```text
find index of target

get last element

move last element
into target position

update last element index

remove last array element

delete target from map
```

---

### Visual

Before:

```text
[10,20,30]
```

Remove:

```text
20
```

Swap:

```text
[10,30,20]
```

Pop:

```text
[10,30]
```

---

### Complexity

```text
Time  : O(1)
Space : O(1)
```

---

# GetRandom Template

### Steps

```text
randomIndex = random(0, size-1)

return values[randomIndex]
```

---

### Complexity

```text
Time  : O(1)
Space : O(1)
```

---

# Key Formula

For an array of size:

```text
n
```

Random index:

```text
rand(0, n-1)
```

Probability of selecting any element:

```text
1 / n
```

Uniform distribution:

```text
✓ Guaranteed
```

because every element occupies exactly one slot.

---

# Complexity Cheatsheet

| Operation | Complexity |
|------------|------------|
| Insert | O(1) |
| Remove | O(1) |
| GetRandom | O(1) |
| Lookup | O(1) |
| Space | O(n) |

---

# Why O(1) Delete Works

Normal array deletion:

```text
Shift elements
```

Cost:

```text
O(n)
```

---

Swap-and-remove deletion:

```text
Swap with last
Pop last
```

Cost:

```text
O(1)
```

---

# Common Pitfalls

## Pitfall #1

Forgetting to update index map.

Wrong:

```text
Swap
Pop
```

Correct:

```text
Swap
Update map
Pop
Delete map entry
```

---

## Pitfall #2

Using:

```java
list.remove(index)
```

directly.

May cause:

```text
O(n)
```

shifting.

---

## Pitfall #3

Trying to preserve insertion order.

Order preservation conflicts with:

```text
O(1) deletion
```

---

## Pitfall #4

Using only a HashMap.

Cannot support:

```text
O(1) random retrieval
```

---

## Pitfall #5

Using only an Array.

Cannot support:

```text
O(1) deletion
```

---

# Mental Model

Think:

```text
Array
=
Storage Layer

HashMap
=
Index Layer
```

The HashMap tells us:

```text
Where is the value?
```

The Array gives us:

```text
Return any value quickly.
```

---

# Optimization Journey

### Brute Force

```text
Array Only

Insert      O(1)
Delete      O(n)
Random      O(1)
```

---

### Better

```text
HashMap Only

Insert      O(1)
Delete      O(1)
Random      O(n)
```

---

### Optimal

```text
Array
+
HashMap

Insert      O(1)
Delete      O(1)
Random      O(1)
```

---

# Similar Problems

## Directly Related

### 381. Insert Delete GetRandom O(1) — Duplicates Allowed

Extension:

```text
HashMap<Value, Set<Indices>>
```

---

### 146. LRU Cache

Combines:

```text
HashMap
+
Doubly Linked List
```

---

### 432. All O`one Data Structure

Advanced design problem involving:

```text
HashMap
+
Linked Structure
```

---

## Same Design Philosophy

### 706. Design HashMap

### 705. Design HashSet

### 355. Design Twitter

### 460. LFU Cache

---

# Interview Sound Bites

Use these concise explanations during interviews.

### Sound Bite #1

> "The challenge isn't insertion or lookup; it's achieving O(1) deletion while keeping O(1) random access."

---

### Sound Bite #2

> "A HashMap gives me the element's position, and the array gives me random access."

---

### Sound Bite #3

> "Swap-and-remove eliminates the O(n) shifting cost of array deletion."

---

### Sound Bite #4

> "Since order is not required, swapping is a legal optimization."

---

### Sound Bite #5

> "Every element occupies exactly one slot in the array, which guarantees uniform random selection."

---

# One-Minute Revision

```text
Need:
Insert O(1)
Delete O(1)
Random O(1)

Use:

Array/List
+
HashMap<Value, Index>

Insert:
Append + Save Index

Remove:
Find Index
Swap With Last
Update Index
Pop Last
Delete Mapping

GetRandom:
Random Index
Return Value

Complexities:

Insert      O(1)
Delete      O(1)
Random      O(1)

Space       O(n)
```

---

# Final Takeaway

```text
HashMap
+
Array
+
Swap-and-Remove
=
RandomizedSet
```

This is one of the most important data structure design patterns for coding interviews because it demonstrates how combining complementary data structures can satisfy strict performance constraints.