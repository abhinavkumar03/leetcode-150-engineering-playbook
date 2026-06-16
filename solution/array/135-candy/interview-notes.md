# Candy — Interview Notes

## What Interviewer Is Testing

This problem is rarely about coding difficulty alone.

Interviewers use it to evaluate whether a candidate can identify hidden constraints and transform them into an optimal greedy solution.

---

### 1. Constraint Analysis

Can you correctly identify the rules?

Requirements:

```text
Every child >= 1 candy
```

and

```text
Higher rating than neighbor
→ More candies than neighbor
```

Strong candidates quickly recognize that the problem is fundamentally a constraint satisfaction problem.

---

### 2. Greedy Thinking

The interviewer wants to see whether you can answer:

> What is the minimum amount of candy needed?

This naturally suggests a greedy strategy.

A good candidate avoids unnecessary increases and only assigns candies when required.

---

### 3. Bidirectional Dependencies

Most candidates initially think only left-to-right.

Example:

```text
[5,4,3,2,1]
```

A single forward pass fails.

Interviewers want candidates to recognize:

```text
Constraints come from BOTH directions.
```

This is the key breakthrough.

---

### 4. Optimization Skills

Typical progression:

```text
Brute Force
    ↓
Repeated Adjustments
    ↓
Two-Pass Greedy
    ↓
O(1) Space Greedy (Follow-up)
```

Interviewers often evaluate whether the candidate naturally moves toward more efficient solutions.

---

### 5. Correctness Reasoning

A strong candidate can explain:

- Why left pass is needed.
- Why right pass is needed.
- Why taking max() is correct.
- Why the result is minimal.

---

## Typical Follow-up Questions

### Follow-up #1

Why can't one pass solve this problem?

Example:

```text
ratings = [5,4,3,2,1]
```

A left-to-right traversal never learns about future decreasing slopes.

Therefore:

```text
One direction
≠
Enough information
```

---

### Follow-up #2

Why do we use max()?

During the second pass:

```text
candies[i] =
max(
    candies[i],
    candies[i + 1] + 1
)
```

Reason:

A larger value may already be required by the left-side constraint.

Overwriting it could violate a previously satisfied condition.

---

### Follow-up #3

Can we reduce space complexity?

Yes.

There is an O(1) space solution.

It tracks:

- Increasing slopes
- Decreasing slopes
- Peak heights

instead of storing an entire candy array.

Complexities:

```text
Time: O(n)
Space: O(1)
```

---

### Follow-up #4

How would you prove optimality?

Each child initially receives:

```text
1 candy
```

Additional candies are assigned only when a constraint demands it.

Therefore:

```text
No unnecessary candies are added.
```

This guarantees minimality.

---

### Follow-up #5

Would sorting help?

No.

The relative order of children matters.

Sorting destroys adjacency information.

---

## Optimization Journey

### Stage 1 — Brute Force

Idea:

Keep adjusting candy counts until all constraints are satisfied.

Example:

```text
Repeat until stable:
    Fix violations
```

Complexity:

```text
O(n²)
```

or worse.

---

### Stage 2 — Observation

Each constraint originates from:

```text
Left Neighbor
or
Right Neighbor
```

These can be handled independently.

---

### Stage 3 — Two-Pass Greedy

Pass 1:

```text
Left → Right
```

Handles:

```text
ratings[i] > ratings[i-1]
```

Pass 2:

```text
Right → Left
```

Handles:

```text
ratings[i] > ratings[i+1]
```

Result:

```text
O(n)
```

time.

---

### Stage 4 — O(1) Space Greedy

Track slope lengths instead of storing candy counts.

Useful for advanced interview discussion.

Not usually required unless specifically requested.

---

## Whiteboard Strategy

### Step 1

Restate requirements.

Example:

```text
At least one candy each.
Higher rating than neighbor → more candies.
Minimize total candies.
```

---

### Step 2

Draw examples.

Increasing sequence:

```text
1 2 3 4
```

Decreasing sequence:

```text
4 3 2 1
```

Valley:

```text
2 1 2
```

Peak:

```text
1 3 5 4 2
```

Visualizing patterns often reveals the two-direction dependency.

---

### Step 3

Propose brute force.

Interviewers expect this.

Demonstrating progression is valuable.

---

### Step 4

Identify inefficiency.

Repeated updates are expensive.

Ask:

```text
Can we satisfy constraints in one pass?
```

---

### Step 5

Discover bidirectional processing.

This is the interview breakthrough.

---

### Step 6

Explain correctness before coding.

A concise explanation:

> The left pass satisfies all increasing relationships.
> The right pass satisfies all decreasing relationships.
> Taking the maximum preserves both constraints with the minimum valid allocation.

---

## Communication Tips

### Good Interview Language

Instead of:

> I think this works.

Say:

> The left pass guarantees every increasing relationship is satisfied.

---

Instead of:

> Let's try this.

Say:

> Since constraints originate from both sides, I'll process the array in both directions independently.

---

Instead of:

> We update the value.

Say:

> We assign the minimum candy count necessary to satisfy the local constraint.

---

### Explain Before Coding

Recommended structure:

```text
Problem
→ Constraints
→ Observation
→ Approach
→ Complexity
→ Code
```

This creates confidence in your solution.

---

## Senior-Level Discussion Points

Senior engineers are often evaluated beyond implementation.

---

### 1. Proof of Correctness

Discuss:

- Constraint satisfaction
- Minimality
- Preservation of previous assignments
- Why max() is required

---

### 2. Tradeoff Analysis

Compare:

#### Two-Pass Solution

```text
Time:  O(n)
Space: O(n)
```

Pros:

- Readable
- Easy to verify
- Low bug risk

---

#### O(1) Space Solution

```text
Time:  O(n)
Space: O(1)
```

Pros:

- More efficient

Cons:

- Harder to reason about
- Easier to introduce bugs

---

### 3. Maintainability

Production code often favors:

```text
Readability
over
micro-optimizations
```

The two-pass solution is typically preferred.

---

### 4. Testing Strategy

Test categories:

```text
Single element
```

```text
All equal
```

```text
Increasing
```

```text
Decreasing
```

```text
Valley
```

```text
Peak
```

```text
Random mixed ratings
```

A senior engineer should proactively discuss validation.

---

## FAANG-Level Variations

### Variation 1

Return the candy distribution array instead of the total.

Example:

```text
[2,1,2]
```

instead of:

```text
5
```

---

### Variation 2

Support circular adjacency.

First and last children become neighbors.

This significantly changes the constraint model.

---

### Variation 3

Weighted candy costs.

Each candy type has a different cost.

Objective becomes:

```text
Minimize total cost
```

rather than candy count.

---

### Variation 4

K-Neighbor Constraints

Instead of immediate neighbors:

```text
Distance <= K
```

must satisfy ordering requirements.

This becomes a more complex graph constraint problem.

---

### Variation 5

Streaming Ratings

Ratings arrive continuously.

Need incremental updates without recomputing the entire solution.

Discussion topics:

- Online algorithms
- Incremental constraint maintenance
- Dynamic updates

---

## Interview Success Checklist

Before finishing the interview, confirm that you can explain:

- Why one pass is insufficient.
- Why two passes work.
- Why max() is necessary.
- Why the solution is minimal.
- Time complexity O(n).
- Space complexity O(n).
- How O(1) space can be achieved as a follow-up.

If you can clearly articulate all six points, you are demonstrating the level of understanding expected in strong mid-level and senior engineering interviews.