# Interview Notes — Best Time to Buy and Sell Stock

## What Interviewer Is Testing

Although LeetCode 121 is categorized as an Easy problem, interviewers rarely use it to test coding syntax.

Instead, they use it to evaluate how candidates think about optimization, state management, and algorithm design.

---

### 1. Brute Force Recognition

Most candidates immediately see:

```text
Buy on one day
Sell on a future day
Find maximum profit
```

The obvious solution is:

```text
Try every possible buy day
Try every possible sell day
```

Interviewers expect candidates to identify this approach first.

Expected complexity:

```text
Time  = O(n²)
Space = O(1)
```

A candidate who cannot derive the brute force solution usually struggles with optimization discussions.

---

### 2. Complexity Analysis

Interviewers want to hear:

```text
Nested loops are expensive.
```

Candidates should explain:

```text
For every price,
I repeatedly inspect future prices.
```

This creates redundant work.

Strong candidates naturally ask:

```text
Can we avoid checking all pairs?
```

---

### 3. Pattern Recognition

This is the core skill being evaluated.

Interviewers want candidates to notice:

```text
To calculate profit today,
I only need the lowest price seen previously.
```

That observation transforms the problem.

Instead of remembering every previous value:

```text
Store only the minimum.
```

This is a common optimization pattern.

---

### 4. State Tracking

A strong solution uses only two variables:

```text
minPrice
maxProfit
```

Interviewers are checking whether candidates can maintain evolving state efficiently.

State-driven solutions appear frequently in:

- Arrays
- Sliding Window
- Dynamic Programming
- Greedy Problems
- Streaming Systems

---

### 5. Greedy Thinking

The solution is fundamentally greedy.

At every step:

```text
Keep the best buy opportunity.
```

The interviewer wants to see whether candidates can justify why this local decision leads to a globally optimal answer.

---

### 6. Communication Skills

Strong candidates explain:

```text
What state is tracked
Why it is sufficient
Why correctness holds
```

Weak candidates often jump directly to code.

Interviewers care more about reasoning than syntax.

---

## Typical Follow-up Questions

### Follow-up 1

What if multiple transactions are allowed?

Expected Answer:

```text
LeetCode 122
Best Time to Buy and Sell Stock II
```

Idea:

```text
Capture every profitable increase.
```

---

### Follow-up 2

What if only two transactions are allowed?

Expected Answer:

```text
LeetCode 123
```

Requires:

```text
Dynamic Programming
```

---

### Follow-up 3

What if transaction fees exist?

Expected Answer:

```text
LeetCode 714
```

Profit calculation changes.

---

### Follow-up 4

What if there is a cooldown day?

Expected Answer:

```text
LeetCode 309
```

Requires state transitions.

---

### Follow-up 5

Can you solve it using Dynamic Programming?

Expected Answer:

Yes.

Track:

```text
Holding Stock
Not Holding Stock
```

The greedy solution is essentially a space-optimized DP.

---

### Follow-up 6

What if prices arrive as a stream?

Expected Answer:

Current solution already works.

Maintain:

```text
minPrice
maxProfit
```

As data arrives.

Memory usage remains O(1).

---

## Optimization Journey

### Stage 1 — Brute Force

Idea:

```text
Check every buy/sell combination.
```

Algorithm:

```text
for buy
    for sell
```

Complexity:

```text
O(n²)
```

---

### Stage 2 — Observe Redundancy

Notice:

```text
Profit =
sellPrice - buyPrice
```

For a given sell day:

```text
Only the smallest previous buy price matters.
```

All larger buy prices are inferior.

---

### Stage 3 — Track Running Minimum

Maintain:

```text
minPrice
```

While traversing.

Update:

```text
minPrice = min(minPrice, currentPrice)
```

---

### Stage 4 — Calculate Profit On The Fly

For every day:

```text
profit = currentPrice - minPrice
```

Update:

```text
maxProfit
```

Complexity becomes:

```text
Time  = O(n)
Space = O(1)
```

---

## Whiteboard Strategy

When solving on a whiteboard:

### Step 1

Restate the problem.

```text
Need one buy
Need one sell
Buy must happen first
Maximize profit
```

---

### Step 2

Present brute force.

Draw:

```text
prices = [7,1,5,3,6,4]
```

Show:

```text
(7,1)
(7,5)
(7,3)
...
```

---

### Step 3

Identify inefficiency.

Ask:

```text
Do I really need all previous prices?
```

Answer:

```text
No.
Only the minimum matters.
```

---

### Step 4

Introduce state.

Write:

```text
minPrice
maxProfit
```

---

### Step 5

Walk through an example.

This demonstrates correctness.

---

### Step 6

Analyze complexity.

State clearly:

```text
Time  = O(n)
Space = O(1)
```

Interviewers expect this discussion.

---

## Communication Tips

### Good Explanation

> "For each day, I compute the profit if I sold today. To maximize that profit, I only need the cheapest stock price seen before today. Therefore I maintain a running minimum and update the maximum profit as I traverse the array."

This communicates:

- Insight
- Correctness
- Optimization

---

### Avoid Saying

```text
I memorized this solution.
```

Instead explain:

```text
How you derived it.
```

Interviewers value reasoning.

---

### Use These Keywords

- Running Minimum
- Greedy
- State Tracking
- Single Pass
- Optimization
- Streaming Friendly
- Constant Space

These terms signal algorithmic maturity.

---

## Senior-Level Discussion Points

Senior engineers are often expected to go beyond the accepted solution.

---

### Production Considerations

Real trading systems may include:

- Transaction fees
- Market slippage
- Taxes
- Liquidity constraints

The LeetCode model simplifies these concerns.

---

### Streaming Data

Prices may arrive continuously:

```text
price1
price2
price3
...
```

Current solution naturally supports streaming.

Maintain:

```text
minPrice
maxProfit
```

No historical storage required.

---

### Memory Efficiency

The solution is:

```text
O(1)
```

Additional memory.

This is optimal.

---

### Scalability

Works efficiently for:

```text
Millions of records
```

because only one pass is required.

---

### Relation To Dynamic Programming

Many senior candidates mention:

```text
This can be viewed as a DP problem.
```

State:

```text
Best buy opportunity so far
Best profit so far
```

The accepted greedy solution is effectively a compressed DP formulation.

---

## FAANG-Level Variations

Interviewers often extend this problem.

---

### Variation 1

Unlimited Transactions

Related:

```text
LeetCode 122
```

Pattern:

```text
Greedy
```

---

### Variation 2

At Most Two Transactions

Related:

```text
LeetCode 123
```

Pattern:

```text
Dynamic Programming
```

---

### Variation 3

At Most K Transactions

Related:

```text
LeetCode 188
```

Pattern:

```text
DP + State Machine
```

---

### Variation 4

Cooldown After Selling

Related:

```text
LeetCode 309
```

Pattern:

```text
DP State Transitions
```

---

### Variation 5

Transaction Fee

Related:

```text
LeetCode 714
```

Pattern:

```text
DP + Greedy
```

---

### Variation 6

Return Buy and Sell Days

Instead of returning profit:

```text
Return:
buyIndex
sellIndex
```

Additional tracking variables are required.

---

### Variation 7

Real-Time Data Feed

Prices arrive continuously.

Goal:

```text
Update profit instantly.
```

Current algorithm adapts naturally.

---

## Recruiter Perspective

When this problem appears in a portfolio repository, recruiters notice:

### Positive Signals

- Complexity analysis included
- Optimization journey documented
- Multiple language implementations
- Edge cases covered
- Clear communication

---

### Strong Engineering Signal

A candidate who explains:

```text
Why O(n) works
```

creates a stronger impression than someone who only provides code.

Documentation quality often differentiates engineers with similar coding ability.

---

## Key Interview Takeaway

The crucial insight is:

```text
For any sell day,
only the lowest previous price matters.
```

Track:

```text
minPrice
maxProfit
```

and update them during a single traversal.

Final Complexity:

```text
Time  = O(n)
Space = O(1)
```

This pattern appears repeatedly in array optimization, greedy algorithms, and dynamic programming interviews.