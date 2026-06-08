# Running Minimum Pattern

## Pattern Definition

The Running Minimum pattern is an optimization technique where we continuously maintain the smallest value encountered so far while traversing a sequence.

Instead of repeatedly searching previous elements, we store the minimum state and update it incrementally.

This transforms many brute-force O(n²) solutions into O(n) solutions.

---

## Core Idea

While traversing:

```text
Maintain:

minimum value seen so far
```

For every new element:

```text
Update minimum if needed
Use current minimum to solve the problem
```

General form:

```text
runningMin = min(runningMin, currentValue)
```

---

## Why This Pattern Exists

Many problems ask:

```text
Find maximum gain
Find largest difference
Find best improvement
Find optimal value after a previous event
```

The brute-force approach compares every pair.

Example:

```text
for i
    for j > i
```

This results in:

```text
O(n²)
```

However, if only the minimum previous value matters:

```text
Store it.
```

Then:

```text
O(n)
```

becomes possible.

---

# Recognition Signals

Use the Running Minimum pattern when you see:

---

## Signal 1

Problems involving:

```text
Maximum Difference
```

Examples:

```text
currentValue - previousValue
```

---

## Signal 2

Ordering Matters

Examples:

```text
Buy before sell
Earlier index before later index
Previous event before current event
```

---

## Signal 3

Questions asking:

```text
Best opportunity so far
```

or

```text
Lowest value before current position
```

---

## Signal 4

Array traversal from:

```text
Left → Right
```

with incremental decisions.

---

## Signal 5

The solution only depends on:

```text
Current value
+
Minimum previous value
```

---

# Generic Template

## Brute Force

```text
for i
    for j > i
```

Complexity:

```text
O(n²)
```

---

## Optimized Template

```text
runningMin = firstElement

answer = initialValue

for each value:

    runningMin = min(runningMin, value)

    update answer using:

    value
    runningMin

return answer
```

---

# Standard Pseudocode

```text
runningMin = nums[0]

result = 0

for value in nums:

    runningMin = min(runningMin, value)

    result = update(result, value, runningMin)

return result
```

---

# Complexity

| Metric | Value |
|----------|----------|
| Time | O(n) |
| Space | O(1) |

---

# Correctness Intuition

At position:

```text
i
```

the only useful information from the past is:

```text
minimum value encountered before i
```

Any larger value is dominated by the smaller one.

Therefore:

```text
Tracking only the minimum
```

is sufficient.

This reduces memory and computation.

---

# Common Pitfalls

## Pitfall 1

Using global minimum.

Incorrect:

```text
globalMax - globalMin
```

because ordering may be violated.

Always respect:

```text
Earlier element before later element.
```

---

## Pitfall 2

Updating answer before minimum.

Carefully consider:

```text
When should minimum be updated?
```

Order matters.

---

## Pitfall 3

Returning negative values.

Many problems require:

```text
Return 0 if no improvement exists.
```

Read constraints carefully.

---

## Pitfall 4

Not handling single-element input.

Example:

```text
[5]
```

No valid comparison exists.

---

# Worked Example

Input:

```text
[7,1,5,3,6,4]
```

Initialize:

```text
runningMin = 7
bestProfit = 0
```

---

Process 1:

```text
runningMin = 1
```

Profit:

```text
1 - 1 = 0
```

---

Process 5:

```text
profit = 5 - 1 = 4
```

Update:

```text
bestProfit = 4
```

---

Process 6:

```text
profit = 6 - 1 = 5
```

Update:

```text
bestProfit = 5
```

Answer:

```text
5
```

---

# Related Patterns

## Running Maximum

Track:

```text
Largest value seen so far
```

instead of minimum.

---

## Prefix Minimum

Store minimum for every index.

Useful when future queries are required.

---

## Kadane's Algorithm

Tracks optimal state while traversing.

Closely related optimization mindset.

---

## Greedy State Tracking

Maintain best state encountered so far.

Common interview pattern.

---

# Problems Using This Pattern

## LeetCode 121

### Best Time to Buy and Sell Stock

Difficulty:

```text
Easy
```

Key Insight:

```text
Track lowest stock price seen so far.
```

Formula:

```text
profit = currentPrice - minPrice
```

Complexity:

```text
O(n)
```

---

## LeetCode 2016

### Maximum Difference Between Increasing Elements

Key Insight:

```text
Track smallest previous value.
```

---

## LeetCode 53 (Related Thinking)

### Maximum Subarray

Not identical, but demonstrates:

```text
Maintain optimal state while traversing.
```

---

# Interview Notes

Interviewers commonly use this pattern to evaluate:

- Optimization skills
- State management
- Greedy reasoning
- Complexity reduction
- Ability to derive O(n) solutions

Strong candidates:

```text
Start with brute force
Identify redundant work
Track only necessary state
```

---

# Quick Recognition Checklist

Before coding, ask:

```text
Do I need every previous value?
```

If answer is:

```text
No
```

Ask:

```text
Do I only need the minimum previous value?
```

If yes:

```text
Running Minimum Pattern
```

is likely the correct approach.

---

# Problems Added To This Pattern

| Problem # | Problem Name | Difficulty |
|------------|--------------|------------|
| 121 | Best Time to Buy and Sell Stock | Easy |

---

# One-Line Memory Trick

```text
If only the smallest previous value matters, store it and keep moving.
```