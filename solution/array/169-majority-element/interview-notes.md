# Majority Element — Interview Notes

## What Interviewer Is Testing

LeetCode 169 appears simple, but interviewers use it to evaluate much more than frequency counting.

### Core Skills Being Evaluated

#### 1. Problem Decomposition

Can you identify:

```text
Find element occurring > n/2 times
```

and translate that into a measurable condition?

```text
frequency > floor(n/2)
```

---

#### 2. Optimization Mindset

Expected progression:

```text
Brute Force
    ↓
Hash Map
    ↓
Boyer-Moore Voting
```

Interviewers want to see whether you naturally search for better solutions.

---

#### 3. Complexity Analysis

You should be able to discuss:

| Approach | Time | Space |
|-----------|--------|--------|
| Brute Force | O(n²) | O(1) |
| Hash Map | O(n) | O(n) |
| Sorting | O(n log n) | O(1) |
| Boyer-Moore | O(n) | O(1) |

---

#### 4. Algorithmic Reasoning

Many candidates memorize Boyer-Moore.

Strong candidates can explain:

```text
WHY it works
```

without relying on memorization.

---

#### 5. Communication

Interviewers often care more about:

```text
How you explain
```

than:

```text
Whether you already know the solution
```

---

# Typical Follow-up Questions

## Follow-up 1

### What if the majority element is NOT guaranteed?

Expected Answer:

Boyer-Moore still finds a candidate.

However:

```text
Candidate ≠ Guaranteed Majority
```

Perform a second pass:

```text
Count occurrences of candidate.
Verify count > n/2.
```

Complexity:

```text
O(n) time
O(1) space
```

---

## Follow-up 2

### Can you solve it using sorting?

Yes.

After sorting:

```text
nums[n/2]
```

must be the majority element.

Example:

```text
[2,2,2,3,3]
```

Middle element:

```text
2
```

Complexity:

```text
O(n log n)
```

---

## Follow-up 3

### Can you solve it with a HashMap?

Yes.

Store frequencies.

Return the element whose count exceeds:

```text
n/2
```

Complexity:

```text
O(n)
```

Space:

```text
O(n)
```

---

## Follow-up 4

### What if elements occur more than n/3 times?

This becomes:

```text
LeetCode 229
Majority Element II
```

Important insight:

```text
At most 2 elements
can occur more than n/3 times.
```

Extended Boyer-Moore keeps:

```text
candidate1
candidate2
count1
count2
```

---

## Follow-up 5

### Can this work in a data stream?

Partially.

Boyer-Moore can continuously maintain a candidate.

Verification may require:

```text
Additional frequency tracking
```

depending on requirements.

---

# Optimization Journey

Interviewers like seeing the thought process.

---

## Stage 1 — Brute Force

### Idea

Count frequency for every element.

```text
for each element
    count occurrences
```

Complexity:

```text
O(n²)
```

---

### Why Improve?

Repeated work.

Same values counted multiple times.

---

## Stage 2 — Hash Map

### Idea

Store frequency once.

```text
frequency[num]++
```

Complexity:

```text
O(n)
```

Space:

```text
O(n)
```

---

### Why Improve?

Extra memory usage.

Problem guarantees majority existence.

Maybe we can exploit that.

---

## Stage 3 — Boyer-Moore Voting

### Key Insight

Majority element appears:

```text
More than all other elements combined
```

Therefore:

```text
Different elements can cancel each other.
```

Remaining candidate must be majority.

Complexity:

```text
O(n)
O(1)
```

Optimal.

---

# Whiteboard Strategy

When solving on a whiteboard:

---

## Step 1

Write the guarantee.

```text
Majority Element > n/2
```

This hint is extremely important.

---

## Step 2

Discuss brute force quickly.

```text
Count everything
```

Show:

```text
O(n²)
```

---

## Step 3

Upgrade to HashMap.

```text
frequency map
```

Show:

```text
O(n)
O(n)
```

---

## Step 4

Observe special property.

```text
Majority > all others combined
```

Introduce cancellation.

---

## Step 5

Write Boyer-Moore.

```text
candidate
count
```

Explain each rule before coding.

---

## Step 6

Run a small example.

Example:

```text
[2,2,1,1,1,2,2]
```

This demonstrates cancellation clearly.

---

# Communication Tips

## Good Explanation

> "Since the majority element appears more than half the time, every occurrence of a different number can cancel out at most one occurrence of the majority element. Because the majority count exceeds all other counts combined, it must survive every cancellation process."

---

## Strong Candidate Behavior

Explain:

```text
Why count decreases
```

instead of only:

```text
What count does
```

---

## Avoid This

Bad explanation:

> "I remember Boyer-Moore from LeetCode."

Interviewers learn nothing from this.

---

## Better Explanation

> "I'm exploiting the mathematical guarantee that one value appears more than half the time."

---

# Senior-Level Discussion Points

These topics help differentiate experienced engineers.

---

## Correctness Argument

Prove:

```text
Majority Count > Other Counts
```

Let:

```text
M = majority count
O = all other counts
```

Given:

```text
M > O
```

After all cancellations:

```text
M - O > 0
```

Therefore majority survives.

---

## Memory Efficiency

Compare:

### HashMap

```text
O(n) space
```

### Boyer-Moore

```text
O(1) space
```

Important for:

- Large datasets
- Embedded systems
- Streaming systems

---

## Streaming Applications

Possible use cases:

- Event analytics
- Telemetry systems
- Log aggregation
- Real-time monitoring

---

## Distributed Systems Discussion

For distributed datasets:

```text
Partition Data
→ Local Candidates
→ Merge Candidates
→ Verify
```

Advanced but useful discussion.

---

# FAANG-Level Variations

## Variation 1

### Majority Element II

Find all elements appearing:

```text
> n/3 times
```

LeetCode:

```text
229
```

---

## Variation 2

Find all elements appearing:

```text
> n/k times
```

Generalized Boyer-Moore.

Keep:

```text
k - 1 candidates
```

---

## Variation 3

No Majority Guarantee

Need:

```text
Candidate Selection
+
Verification Pass
```

---

## Variation 4

Massive Dataset

Data cannot fit into memory.

Possible solutions:

- Streaming processing
- Distributed counting
- Candidate aggregation

---

## Variation 5

Weighted Majority

Each element has a weight.

Now frequency becomes:

```text
weighted frequency
```

Common in voting systems.

---

# Interview Cheat Answers

### Why does Boyer-Moore work?

```text
Majority element appears more than all other elements combined,
so pairwise cancellation cannot eliminate it completely.
```

---

### Why is the space O(1)?

```text
Only candidate and count are stored.
```

---

### Why is the time O(n)?

```text
Single pass through the array.
```

---

### When is verification required?

```text
When majority existence is not guaranteed.
```

---

### Most Important Interview Takeaway

Whenever a problem guarantees:

```text
An element appears more than half the time
```

immediately consider:

```text
Boyer-Moore Voting Algorithm
```

This pattern appears frequently in interviews and is one of the highest-value array optimization techniques to master.