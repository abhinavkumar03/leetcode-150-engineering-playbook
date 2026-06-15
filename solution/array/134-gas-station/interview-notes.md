# Interview Notes - 134. Gas Station

---

# What Interviewer Is Testing

This problem is much more than an array traversal exercise.

Interviewers use it to evaluate whether a candidate can:

- Move from brute force to optimization
- Identify a Greedy pattern
- Reason about correctness
- Prove why a solution works
- Eliminate unnecessary computation
- Communicate algorithmic insights clearly

---

## Technical Skills Being Evaluated

### Array Traversal

Can you efficiently process all stations?

### Running State Management

Can you maintain:

```text
currentTank
totalBalance
candidateStart
```

without additional memory?

### Greedy Thinking

Can you recognize that some candidates can be permanently discarded?

### Complexity Optimization

Can you improve:

```text
O(n²)
```

to:

```text
O(n)
```

---

# Expected Interview Progression

A strong candidate typically follows this path:

```text
Understand Problem
        ↓
Brute Force
        ↓
Find Repeated Work
        ↓
Discover Elimination Property
        ↓
Greedy Solution
        ↓
Correctness Proof
        ↓
Complexity Analysis
```

---

# Typical Follow-up Questions

## Q1. Why Do We Need totalBalance?

### Answer

Even if a starting station appears valid locally, a solution is impossible when:

```text
sum(gas) < sum(cost)
```

Therefore:

```text
totalBalance < 0
```

means no valid starting station exists.

Return:

```text
-1
```

immediately.

---

## Q2. Why Can We Skip Stations?

Suppose:

```text
start = s
```

and we fail at:

```text
station i
```

Then every station between:

```text
[s ... i]
```

must also fail.

Reason:

They start with less accumulated fuel than station `s`.

Therefore none can become valid starting points.

---

## Q3. Why Reset to i + 1?

When failure occurs at station `i`:

```text
currentTank < 0
```

we know:

```text
s ... i
```

are invalid.

The next possible candidate is:

```text
i + 1
```

---

## Q4. Is the Answer Guaranteed Unique?

For this LeetCode problem:

```text
Yes
```

The problem statement guarantees uniqueness when a solution exists.

---

## Q5. Can We Solve It Using Prefix Sums?

Yes.

Possible approaches:

- Prefix sums
- Circular prefix sums
- Minimum prefix position

However:

```text
Greedy
```

is simpler and optimal.

---

## Q6. What Is the Space Complexity?

```text
O(1)
```

Only a few variables are stored.

---

# Optimization Journey

---

## Stage 1 — Brute Force

Try every station.

For each station:

```text
simulate complete trip
```

### Complexity

```text
Time  = O(n²)
Space = O(1)
```

---

### Why It Is Inefficient

Many route segments are repeatedly evaluated.

Example:

```text
Start at 0
Fail at 10

Start at 1
Fail at 10

Start at 2
Fail at 10
```

Large amounts of duplicated work occur.

---

## Stage 2 — Key Observation

Failure provides information.

If:

```text
Start = s
Fail = i
```

then:

```text
[s ... i]
```

cannot contain the answer.

This allows elimination.

---

## Stage 3 — Greedy Solution

Maintain:

```text
currentTank
totalBalance
startStation
```

Whenever:

```text
currentTank < 0
```

reset:

```text
startStation = i + 1
currentTank = 0
```

Continue scanning.

---

### Complexity

```text
Time  = O(n)
Space = O(1)
```

---

# Whiteboard Strategy

When solving on a whiteboard:

---

## Step 1

Draw a circle.

```text
0 → 1 → 2 → 3 → 4 → 0
```

Show that the route is circular.

---

## Step 2

Write:

```text
gas[i] - cost[i]
```

for every station.

Convert the problem into:

```text
net gain / loss
```

at each station.

---

## Step 3

Explain Failure Logic

Example:

```text
Start at 0
Fail at 4
```

Ask:

```text
Can 1 work?
Can 2 work?
Can 3 work?
```

No.

This naturally motivates the greedy elimination.

---

## Step 4

Derive Variables

```text
totalBalance
currentTank
startStation
```

Explain the purpose of each.

---

## Step 5

Walk Through Example

Use:

```text
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]
```

Demonstrate resets.

---

# Communication Tips

Interviewers care about explanation quality.

Avoid jumping directly into code.

---

## Good Explanation

> First, I would check whether the total gas is at least the total travel cost. If not, completing the circuit is impossible. Then I maintain a running tank balance. Whenever the tank becomes negative, I know the current starting position and all stations before the failure point are invalid, so I move the candidate start to the next station.

---

## Bad Explanation

> It's greedy. We reset whenever it becomes negative.

This states the rule but not the reason.

---

# Senior-Level Discussion Points

Experienced engineers should be able to discuss:

---

## Correctness Proof

Why does skipping stations remain safe?

This is the most important discussion point.

---

## Invariant

During traversal:

```text
currentTank >= 0
```

for the current candidate segment.

If violated:

```text
candidate becomes invalid
```

---

## Global vs Local Conditions

### Local Condition

```text
currentTank
```

tracks current route feasibility.

### Global Condition

```text
totalBalance
```

tracks overall feasibility.

Both are required.

---

## Failure Elimination

This is a classic optimization pattern.

A local failure allows us to eliminate an entire set of future candidates.

Similar ideas appear in:

- Jump Game
- Scheduling problems
- Interval problems
- Resource allocation systems

---

# FAANG-Level Variations

---

## Variation 1

Return all valid starting stations.

Discussion:

```text
Multiple solutions possible
```

May require additional analysis.

---

## Variation 2

Route Is Not Circular

Now:

```text
start → end
```

instead of:

```text
start → start
```

Greedy proof changes.

---

## Variation 3

Limited Tank Capacity

Additional constraint:

```text
tank <= capacity
```

Requires modified logic.

---

## Variation 4

Dynamic Fuel Updates

Gas values change during execution.

Possible discussion:

- Segment Trees
- Fenwick Trees
- Online Queries

---

## Variation 5

Find Maximum Remaining Fuel

Instead of finding a valid start:

```text
maximize ending fuel
```

Optimization objective changes.

---

# Red Flags Interviewers Notice

---

## Red Flag 1

Candidate never checks:

```text
sum(gas) < sum(cost)
```

---

## Red Flag 2

Candidate cannot explain why stations are skipped.

---

## Red Flag 3

Candidate memorized the solution but cannot prove correctness.

---

## Red Flag 4

Candidate jumps directly to code without discussing strategy.

---

# Key Interview Takeaway

The hardest part of this problem is not implementation.

The implementation is only a few lines.

The real challenge is recognizing and proving the Greedy elimination rule:

```text
If a start station fails,
every station before the failure point
can be safely discarded.
```

Candidates who clearly communicate this insight typically perform very well in technical interviews.