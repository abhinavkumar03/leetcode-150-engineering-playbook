# Majority Element

## Problem Statement

Given an integer array `nums` of size `n`, return the majority element.

The majority element is the element that appears more than `⌊n / 2⌋` times.

You may assume that the majority element always exists in the array.

### Example 1

```text
Input: nums = [3,2,3]
Output: 3
```

### Example 2

```text
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

### Constraints

```text
n == nums.length
1 <= n <= 5 * 10^4
-10^9 <= nums[i] <= 10^9
```

---

## Difficulty

Easy

---

## Tags

- Array
- Hash Table
- Divide and Conquer
- Sorting
- Counting
- Boyer-Moore Voting Algorithm

---

## Pattern

### Primary Pattern

Boyer-Moore Voting Algorithm

### Secondary Patterns

- Frequency Counting
- Greedy Candidate Elimination
- Array Traversal

---

## Intuition

A majority element appears more than half of the time in the array.

This means:

```text
majority_count > n / 2
```

Even if we repeatedly remove one occurrence of the majority element and one occurrence of a different element, the majority element will still remain.

This observation leads to the Boyer-Moore Voting Algorithm.

The algorithm continuously cancels out different elements and keeps track of a potential majority candidate.

Since the problem guarantees a majority element exists, the final candidate must be the answer.

---

## Key Observation

Consider:

```text
[2,2,1,1,1,2,2]
```

Pair cancellation:

```text
2 cancels 1
2 cancels 1
2 cancels 1
```

Remaining:

```text
2
2
```

The majority element survives all cancellations.

This is the fundamental idea behind Boyer-Moore Voting.

---

## Brute Force Approach

### Idea

For every element:

1. Count its occurrences.
2. Check whether count > n/2.
3. Return the element if condition is satisfied.

### Algorithm

1. Iterate through every element.
2. For each element, scan the entire array.
3. Count occurrences.
4. Return the first element whose count exceeds n/2.

### Complexity

```text
Time Complexity: O(n²)
Space Complexity: O(1)
```

### Limitations

- Extremely slow for large arrays.
- Repeated counting of the same values.
- Not scalable.

---

## Optimized Approach

### Approach 1: Hash Map Counting

Store frequency of each element.

#### Algorithm

1. Create a frequency map.
2. Traverse the array.
3. Increment count of current number.
4. Return element once frequency exceeds n/2.

#### Why It Works

The majority element occurs more than half the array size.

Tracking frequencies guarantees we eventually identify it.

#### Complexity

```text
Time Complexity: O(n)
Space Complexity: O(n)
```

---

### Approach 2: Boyer-Moore Voting Algorithm (Optimal)

#### Algorithm

1. Initialize:

```text
candidate = null
count = 0
```

2. Traverse array:

- If count == 0:
  - Set current element as candidate.
- If current element equals candidate:
  - Increment count.
- Otherwise:
  - Decrement count.

3. Return candidate.

---

#### Why It Works

Whenever we encounter:

```text
candidate != current_element
```

we perform a cancellation.

Since the majority element appears more than all other elements combined:

```text
majority_count > n / 2
```

it can never be completely cancelled out.

Thus, the final surviving candidate is the majority element.

---

#### Complexity

```text
Time Complexity: O(n)

Space Complexity: O(1)
```

This is the optimal solution.

---

## Edge Cases

### Empty Input

Not possible according to constraints.

```text
[]
```

---

### Single Element

```text
[7]
```

Output:

```text
7
```

The only element is the majority.

---

### All Elements Same

```text
[5,5,5,5]
```

Output:

```text
5
```

---

### Negative Values

```text
[-1,-1,-1,2,3]
```

Output:

```text
-1
```

Algorithm works identically.

---

### Large Inputs

```text
50,000 elements
```

Boyer-Moore still runs efficiently:

```text
O(n) time
O(1) space
```

---

## Dry Run

### Input

```text
nums = [2,2,1,1,1,2,2]
```

| Index | Value | Candidate | Count |
|---------|---------|---------|---------|
| 0 | 2 | 2 | 1 |
| 1 | 2 | 2 | 2 |
| 2 | 1 | 2 | 1 |
| 3 | 1 | 2 | 0 |
| 4 | 1 | 1 | 1 |
| 5 | 2 | 1 | 0 |
| 6 | 2 | 2 | 1 |

Final Answer:

```text
2
```

---

## Common Mistakes

### Mistake 1

Using sorting without understanding complexity.

```text
Time: O(n log n)
```

Works, but not optimal.

---

### Mistake 2

Forgetting the majority element guarantee.

Many candidates unnecessarily perform verification.

For this problem, verification is not required.

---

### Mistake 3

Incorrect candidate reset logic.

Wrong:

```text
if count < 0
```

Correct:

```text
if count == 0
```

---

### Mistake 4

Using extra memory when O(1) solution exists.

HashMap works but is not optimal.

---

## Interview Discussion

### Expected Progression

Level 1:

```text
Brute Force
O(n²)
```

Level 2:

```text
Hash Map
O(n)
O(n)
```

Level 3:

```text
Boyer-Moore
O(n)
O(1)
```

Interviewers usually expect candidates to reach Level 3.

---

### What Interviewers Evaluate

- Pattern recognition
- Optimization skills
- Complexity analysis
- Ability to justify correctness
- Understanding of cancellation logic

---

## Follow-up Questions

### Follow-up 1

What if the majority element is not guaranteed?

Answer:

After Boyer-Moore, perform a second pass to verify frequency.

---

### Follow-up 2

Can you solve using sorting?

Yes.

```text
Sort array
Return nums[n/2]
```

Complexity:

```text
O(n log n)
```

---

### Follow-up 3

Can this be solved in parallel systems?

Yes.

Frequency counting can be distributed across partitions and merged later.

---

### Follow-up 4

Can there be multiple majority elements?

Not under the given definition.

However, a variation exists:

```text
Elements occurring more than n/3 times
```

which uses an extended Boyer-Moore algorithm.

---

## Real World Applications

### Vote Counting Systems

Determine dominant candidate in election-like data.

---

### Event Stream Analysis

Identify the most common event in a stream.

---

### Log Processing

Find dominant log category.

---

### Network Traffic Monitoring

Identify the most frequent request type.

---

### Recommendation Systems

Detect overwhelmingly popular user actions.

---

## Related Problems

### Easy

- 217. Contains Duplicate
- 219. Contains Duplicate II

### Medium

- 229. Majority Element II
- 347. Top K Frequent Elements

### Hard

- 295. Find Median from Data Stream

---

## Key Takeaway

The Boyer-Moore Voting Algorithm is one of the most important interview algorithms because it transforms a frequency counting problem into a constant-space solution using candidate elimination and pair cancellation.

Whenever a problem guarantees a dominant element, Boyer-Moore should immediately be considered.