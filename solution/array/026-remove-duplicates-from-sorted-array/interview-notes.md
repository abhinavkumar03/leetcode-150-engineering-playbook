# Interview Notes — Remove Duplicates from Sorted Array

## What Interviewer Is Testing

Although LeetCode 26 is classified as an Easy problem, interviewers use it to evaluate several fundamental engineering skills.

### 1. Pattern Recognition

The strongest signal is whether the candidate immediately notices:

```text
The array is already sorted.
```

Because the array is sorted:

```text
Duplicates are adjacent.
```

This observation eliminates the need for:

- HashSet
- HashMap
- Extra arrays
- Sorting again

Interviewers want to see whether you can exploit problem constraints.

---

### 2. In-Place Array Manipulation

Many candidates solve the problem using additional memory.

However, the actual challenge is:

```text
Modify the existing array
without allocating extra storage.
```

This tests understanding of:

- Memory efficiency
- Data movement
- Index management

---

### 3. Two Pointer Technique

The expected solution uses:

```text
Fast Pointer
+
Slow Pointer
```

Interviewers evaluate whether you understand:

- Read pointer vs write pointer
- Maintaining invariants
- Single-pass processing

---

### 4. Complexity Awareness

A strong candidate should explicitly state:

| Metric | Expected |
|----------|----------|
| Time | O(n) |
| Space | O(1) |

Interviewers want evidence that you can reason about tradeoffs rather than simply writing working code.

---

## Typical Follow-up Questions

### Follow-up 1

**Why does the sorted property matter?**

Expected Answer:

```text
Because duplicates appear consecutively.
```

Without sorting:

```text
[1,3,1,2]
```

duplicate detection requires additional storage.

---

### Follow-up 2

**How would you solve this if the array were unsorted?**

Possible Solution:

```text
HashSet
```

Complexity:

```text
Time: O(n)
Space: O(n)
```

---

### Follow-up 3

**Can you preserve order in an unsorted array?**

Yes.

Use:

```text
HashSet + Write Index
```

while processing from left to right.

---

### Follow-up 4

**What if duplicates can appear at most twice?**

This leads directly to:

LeetCode 80 — Remove Duplicates from Sorted Array II

Candidates should recognize that the same pattern extends naturally.

---

### Follow-up 5

**Can you solve it recursively?**

Yes.

However:

```text
Recursive solution
=
Extra stack space
```

Iterative is preferred.

---

## Optimization Journey

Interviewers often want to hear how you arrived at the optimal solution.

### Stage 1 — Brute Force

Use a set.

```text
Store unique elements
Copy back to array
```

Complexity:

```text
Time: O(n)
Space: O(n)
```

---

### Stage 2 — Observe Sorting

Because the array is sorted:

```text
Equal values are adjacent.
```

This means:

```text
Only compare neighboring unique values.
```

---

### Stage 3 — Introduce Two Pointers

One pointer:

```text
Reads data
```

Another pointer:

```text
Writes valid data
```

Result:

```text
Time: O(n)
Space: O(1)
```

Optimal.

---

## Whiteboard Strategy

When solving in an interview:

### Step 1

State observations first.

Example:

```text
The array is sorted, which means duplicates are grouped together.
```

---

### Step 2

Define pointer roles.

```text
slow → last unique element

fast → scanning pointer
```

Interviewers appreciate clear terminology.

---

### Step 3

Explain the invariant.

A useful statement:

```text
Everything from index 0 to slow
contains unique values.
```

This demonstrates algorithmic maturity.

---

### Step 4

Walk through a small example.

Example:

```text
[1,1,2,3,3]
```

Show:

```text
slow movement
fast movement
writes
```

before coding.

---

### Step 5

Write code last.

Strong candidates spend more time reasoning than typing.

---

## Communication Tips

### Good Explanation

```text
I maintain a write pointer that marks
the end of the unique section.

As I scan with another pointer,
every new unique value is written
into the next available position.
```

This clearly communicates intent.

---

### Avoid Saying

```text
I just move pointers around.
```

or

```text
I memorized this solution.
```

Interviewers want reasoning, not recollection.

---

### Mention the Invariant

A high-signal statement:

```text
At every step,
nums[0...slow]
contains only unique elements.
```

This demonstrates confidence in correctness.

---

## Senior-Level Discussion Points

A senior engineer should connect the algorithm to broader system design concepts.

### Data Compaction

This algorithm behaves like:

```text
Array compaction
```

used in:

- Memory managers
- Storage engines
- Data pipelines

---

### Stream Processing

The pattern resembles:

```text
Read Stream
→ Filter
→ Write Stream
```

which appears in:

- Kafka consumers
- Event processing systems
- ETL pipelines

---

### Write Optimization

The algorithm minimizes writes.

Only unique values trigger writes.

This matters when:

- Data is large
- Storage is expensive
- Cache efficiency matters

---

### Memory-Constrained Systems

The O(1) space requirement mirrors real-world environments:

- Embedded systems
- Mobile devices
- High-throughput services

---

## FAANG-Level Variations

### Variation 1

Allow each element at most twice.

Example:

```text
[1,1,1,2,2,3]
```

Result:

```text
[1,1,2,2,3]
```

Related Problem:

```text
LeetCode 80
```

---

### Variation 2

Keep at most K occurrences.

Generalize:

```text
Allowed Frequency = K
```

Common Google-style extension.

---

### Variation 3

Remove all duplicates completely.

Example:

```text
[1,1,2,3,3]
```

Result:

```text
[2]
```

Now the problem becomes more complex.

---

### Variation 4

Deduplicate an unsorted stream.

Constraints:

```text
Cannot sort
Must preserve order
```

Likely solution:

```text
HashSet
```

---

### Variation 5

Deduplicate billions of records.

Discussion topics:

- External sorting
- Bloom filters
- Distributed processing
- MapReduce
- Memory limitations

This transforms a simple array problem into a large-scale systems discussion.

---

## Interview Summary

### Recognition Signal

```text
Sorted Array
+
Remove Duplicates
+
O(1) Space
```

Immediately suggests:

```text
Two Pointers
```

---

### Expected Optimal Complexity

```text
Time  : O(n)
Space : O(1)
```

---

### Key Interview Soundbite

"The sorted property guarantees that duplicates are adjacent, so I can maintain a compacted unique region using a slow pointer while a fast pointer scans the remaining elements."