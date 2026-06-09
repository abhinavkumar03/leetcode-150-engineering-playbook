# Interview Notes — Best Time to Buy and Sell Stock II

## What Interviewer Is Testing

This problem is less about stock trading and more about recognizing a Greedy optimization opportunity.

Interviewers typically evaluate whether you can:

### 1. Identify the Correct Pattern

Many candidates incorrectly classify this as:

* Dynamic Programming
* Sliding Window
* Two Pointers

Strong candidates recognize:

```text
Greedy
```

because every local profit contributes directly to the global optimum.

---

### 2. Simplify the Problem

Interviewers want to see whether you can transform:

```text
Find optimal buy/sell transactions
```

into:

```text
Sum all positive price increases
```

The ability to simplify a problem is often more valuable than writing code.

---

### 3. Optimization Thinking

Typical progression:

```text
Brute Force
    ↓
Dynamic Programming
    ↓
Greedy Observation
```

Interviewers often care more about this journey than the final code.

---

### 4. Complexity Awareness

Expected discussion:

| Approach    | Time  | Space |
| ----------- | ----- | ----- |
| Brute Force | O(2ⁿ) | O(n)  |
| DP          | O(n)  | O(n)  |
| Greedy      | O(n)  | O(1)  |

Candidates should justify why Greedy is optimal.

---

### 5. Communication Skills

Interviewers want to hear:

> "Every profitable trend can be decomposed into smaller profitable transactions. Therefore summing all positive daily gains achieves the same total profit."

A clear explanation is often the deciding factor between candidates.

---

# Typical Follow-up Questions

## Follow-up 1

### Why does the greedy approach work?

Expected answer:

For an increasing sequence:

```text
1 → 3 → 5 → 8
```

Single transaction:

```text
8 - 1 = 7
```

Greedy:

```text
(3-1)+(5-3)+(8-5)
=
7
```

Both produce identical profit.

---

## Follow-up 2

### Why can't we simply buy every day?

Because:

```text
Only one stock can be held at a time.
```

The greedy solution is mathematically equivalent to valid buy/sell operations.

---

## Follow-up 3

### What if only one transaction is allowed?

This becomes:

```text
LeetCode 121
Best Time to Buy and Sell Stock
```

Pattern:

```text
Running Minimum
```

---

## Follow-up 4

### What if there is a transaction fee?

This becomes:

```text
LeetCode 714
```

Requires state tracking and Dynamic Programming.

---

## Follow-up 5

### What if there is a cooldown day?

This becomes:

```text
LeetCode 309
```

Additional DP states are required.

---

## Follow-up 6

### What if only k transactions are allowed?

This becomes:

```text
LeetCode 188
```

Classic DP optimization problem.

---

# Optimization Journey

A strong candidate should present solutions in this order.

---

## Stage 1 — Brute Force

Idea:

```text
Try all buy/sell combinations.
```

Decision tree:

```text
Buy
Sell
Skip
```

Complexity:

```text
Time: O(2ⁿ)
Space: O(n)
```

Problem:

```text
Too many repeated states.
```

---

## Stage 2 — Dynamic Programming

State:

```text
hold
notHold
```

Transition:

```text
hold = max(hold, notHold - price)

notHold = max(notHold, hold + price)
```

Complexity:

```text
Time: O(n)
Space: O(1)
```

Observation:

```text
Still more complicated than necessary.
```

---

## Stage 3 — Greedy Insight

Key realization:

```text
Every upward movement contributes profit.
```

Therefore:

```text
Profit += max(0, prices[i] - prices[i-1])
```

Complexity:

```text
Time: O(n)
Space: O(1)
```

Optimal.

---

# Whiteboard Strategy

When solving on a whiteboard:

---

## Step 1

Write a small example:

```text
[1,5,3,6]
```

---

## Step 2

Show optimal transactions:

```text
Buy 1 Sell 5 = 4

Buy 3 Sell 6 = 3

Total = 7
```

---

## Step 3

Rewrite as differences:

```text
(5-1)
+
(6-3)
=
7
```

---

## Step 4

Generalize

```text
Whenever price increases,
add the difference.
```

---

## Step 5

Write final algorithm

```text
profit = 0

for i from 1 to n-1:
    if prices[i] > prices[i-1]:
        profit += prices[i] - prices[i-1]

return profit
```

---

# Communication Tips

## Good Interview Explanation

> We are allowed to make unlimited transactions while holding at most one stock at a time. Any increasing sequence can be broken into multiple smaller profitable transactions without changing total profit. Therefore, summing every positive day-to-day increase yields the maximum achievable profit.

---

## Avoid Saying

❌

```text
I memorized this solution.
```

❌

```text
This is just a greedy problem.
```

❌

```text
LeetCode says this works.
```

Interviewers want reasoning, not memorization.

---

## Better Explanation

✅

```text
The total gain across an increasing trend equals the sum of all intermediate gains, which allows us to accumulate every positive difference independently.
```

---

# Senior-Level Discussion Points

Senior engineers are expected to go beyond implementation.

---

## Business Interpretation

The algorithm captures:

```text
Every profitable market movement.
```

instead of attempting to predict:

```text
The single best trade.
```

---

## Scalability Discussion

For large datasets:

```text
n = 1,000,000+
```

The algorithm remains efficient because:

* Single pass
* Constant memory
* No recursion
* No auxiliary structures

---

## Production Considerations

Potential concerns:

### Data Validation

```text
Null input
Empty arrays
Corrupted pricing feeds
```

---

### Numerical Limits

Large financial datasets may require:

```text
long
long long
BigInteger
```

depending on constraints.

---

### Streaming Prices

If prices arrive continuously:

```text
price₁ → price₂ → price₃ ...
```

profit can be updated online:

```text
if current > previous:
    profit += current - previous
```

No historical storage is required.

---

# FAANG-Level Variations

Interviewers frequently extend Stock II into harder variants.

---

## Variant 1

### Best Time to Buy and Sell Stock I

LeetCode:

```text
121
```

Pattern:

```text
Running Minimum
```

---

## Variant 2

### Best Time to Buy and Sell Stock III

LeetCode:

```text
123
```

Constraint:

```text
Maximum two transactions
```

Pattern:

```text
Dynamic Programming
```

---

## Variant 3

### Best Time to Buy and Sell Stock IV

LeetCode:

```text
188
```

Constraint:

```text
Maximum k transactions
```

Pattern:

```text
DP + State Machine
```

---

## Variant 4

### Stock With Cooldown

LeetCode:

```text
309
```

Additional state:

```text
Cooldown Day
```

---

## Variant 5

### Stock With Transaction Fee

LeetCode:

```text
714
```

Modification:

```text
Profit = Sell - Buy - Fee
```

---

# Interview Cheat Answer

If asked:

### "What's the key insight?"

Answer:

> Every positive price increase contributes directly to the optimal solution. Since an increasing trend's total profit equals the sum of its daily gains, we can simply accumulate all positive differences in one pass.

---

# Hiring Manager Perspective

A strong solution demonstrates:

* Pattern recognition
* Optimization skills
* Complexity awareness
* Clear communication
* Ability to justify correctness

A candidate who explains *why* Greedy works is typically rated significantly higher than one who only writes the code.