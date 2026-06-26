# Interview Notes

## Problem

**LeetCode 3 — Longest Substring Without Repeating Characters**

**Difficulty:** Medium

**Primary Pattern:** Sliding Window

---

# What Interviewer Is Testing

This problem is much more than a string manipulation exercise. Interviewers use it to evaluate your ability to recognize patterns, optimize algorithms, and clearly communicate your reasoning.

## 1. Pattern Recognition

Can you identify that this is a **Sliding Window** problem?

Typical signals include:

- Longest/shortest **substring**
- Contiguous sequence
- Maintain a valid window
- Expand and shrink dynamically
- Optimize from brute force

A strong candidate quickly identifies these cues and proposes a window-based approach.

---

## 2. Optimization Skills

Interviewers expect you to progress naturally through the optimization journey.

### Stage 1 — Brute Force

Generate every substring and check for duplicates.

```
Time: O(n²)–O(n³)
```

Discuss why this approach is inefficient due to repeated work on overlapping substrings.

---

### Stage 2 — Sliding Window

Maintain only the current valid substring.

```
Time: O(n)
```

Explain that each character is processed at most twice—once when entering the window and once when leaving it.

---

## 3. Data Structure Selection

Explain why a **HashSet** is appropriate:

- Fast membership checks (`O(1)` average).
- Easy insertion and removal.
- Represents the current window naturally.

Then mention the **HashMap** optimization:

- Stores the last seen index of each character.
- Allows the `left` pointer to jump directly.
- Improves practical performance while keeping `O(n)` complexity.

Showing awareness of both solutions demonstrates depth.

---

## 4. Maintaining the Window Invariant

The key invariant is:

> **Every character inside the current window is unique.**

Whenever a duplicate appears:

1. Shrink the window.
2. Remove characters from the left.
3. Restore the invariant.
4. Continue expanding.

Interviewers want to hear this invariant stated explicitly because it proves correctness.

---

## 5. Complexity Analysis

Be prepared to justify why the solution is linear.

### Time Complexity

```
O(n)
```

Reason:

- The `right` pointer moves from left to right once.
- The `left` pointer also moves only forward.
- No pointer ever moves backward.

Therefore, each character is processed at most twice.

### Space Complexity

```
O(min(n, charset))
```

The window stores only distinct characters currently in the substring.

---

# Typical Follow-up Questions

## Q1. Can you return the substring instead of its length?

Yes.

Track:

- Start index of the best window.
- Maximum length.

After traversal:

```
substring = s[start : start + maxLength]
```

No additional complexity is introduced.

---

## Q2. Can this be solved using a HashMap?

Yes.

Instead of removing characters one by one:

```
left = max(left, lastSeen[currentChar] + 1)
```

This skips directly past the previous occurrence.

Benefits:

- Fewer iterations in practice.
- Common production implementation.

---

## Q3. What if the input contains Unicode characters?

The algorithm remains unchanged.

Only the storage changes:

- Go: `map[rune]int`
- Java: `Map<Character, Integer>` (or Unicode-aware handling as needed)
- JavaScript: `Map`

---

## Q4. Can you solve it without extra space?

In general, no.

You need a way to determine whether a character already exists in the current window.

Without auxiliary storage, duplicate detection becomes expensive, leading back toward quadratic time.

---

## Q5. Can this work on an infinite stream of characters?

Yes.

Maintain the same sliding window and state as new characters arrive.

The algorithm naturally supports streaming because it processes input incrementally.

---

# Optimization Journey

A senior engineer should articulate the evolution of the solution:

```
Brute Force
      │
      ▼
Observe overlapping substrings
      │
      ▼
Use two pointers
      │
      ▼
Maintain a valid window
      │
      ▼
Use HashSet for uniqueness
      │
      ▼
Optimize further with HashMap (last seen index)
```

Interviewers value the reasoning process as much as the final code.

---

# Whiteboard Strategy

When solving on a whiteboard:

### Step 1

Clarify the problem:

- Substring (contiguous)?
- Return length or substring?
- Character set assumptions?
- Input constraints?

---

### Step 2

State the brute-force approach first.

This demonstrates structured thinking before optimization.

---

### Step 3

Identify the repeated work.

Explain that adjacent substrings overlap heavily.

---

### Step 4

Introduce the Sliding Window.

Draw the window:

```
a b c a b c b b
^     ^
L     R
```

Explain:

- `R` expands the window.
- `L` shrinks it only when necessary.

---

### Step 5

Describe the invariant.

> The current window always contains unique characters.

---

### Step 6

Walk through a sample input.

Example:

```
abcabcbb
```

Show pointer movement and window updates.

---

### Step 7

Analyze complexity.

Always conclude with:

```
Time: O(n)
Space: O(min(n, charset))
```

---

# Communication Tips

During the interview:

- Think aloud while solving.
- Explain *why* you move each pointer.
- Avoid jumping directly into code.
- Use meaningful variable names (`left`, `right`, `window`, `maxLength`).
- Mention the maintained invariant frequently.
- Verify the solution with a small example before finishing.
- Discuss edge cases proactively.

Strong communication often differentiates candidates with similar coding ability.

---

# Senior-Level Discussion Points

A senior candidate should go beyond the basic implementation.

## Compare HashSet vs HashMap

### HashSet

Pros:

- Simple.
- Easy to reason about.
- Great for interviews.

Cons:

- May remove several characters one by one.

---

### HashMap

Pros:

- Jump directly to the next valid position.
- Fewer pointer updates.
- Often preferred in production.

Cons:

- Slightly more complex logic.

---

## Memory Considerations

For ASCII:

```
Space ≈ O(1)
```

For Unicode:

```
Space ≈ O(number of unique characters)
```

Understanding how the character set affects memory usage shows attention to implementation details.

---

## Correctness Argument

To prove correctness:

1. The window always satisfies the uniqueness invariant.
2. Any invalid window is repaired immediately.
3. Every possible valid window is considered.
4. The maximum length is updated whenever a larger valid window is found.

This forms a concise correctness proof.

---

# FAANG-Level Variations

Interviewers may extend the problem in several directions.

## Variation 1

Return the **actual longest substring** instead of its length.

---

## Variation 2

Find the longest substring containing **at most K distinct characters**.

Pattern:

Sliding Window + Frequency Map.

---

## Variation 3

Find the longest substring after replacing at most **K** characters.

Related problem:

**LeetCode 424 — Longest Repeating Character Replacement**

---

## Variation 4

Find the **minimum window** satisfying a condition.

Related problem:

**LeetCode 76 — Minimum Window Substring**

---

## Variation 5

Count all substrings with unique characters.

Requires adapting the window contribution rather than tracking only the maximum.

---

## Variation 6

Support a real-time character stream.

Maintain:

- Sliding window.
- Last seen positions.
- Incremental updates.

---

## Related Interview Problems

| Problem | Pattern | Difficulty |
|---------|---------|------------|
| 209. Minimum Size Subarray Sum | Sliding Window | Medium |
| 76. Minimum Window Substring | Sliding Window | Hard |
| 567. Permutation in String | Sliding Window | Medium |
| 438. Find All Anagrams in a String | Sliding Window | Medium |
| 424. Longest Repeating Character Replacement | Sliding Window | Medium |
| 904. Fruit Into Baskets | Sliding Window | Medium |
| 1004. Max Consecutive Ones III | Sliding Window | Medium |

---

# Key Takeaways

- Recognize Sliding Window from "longest/shortest contiguous sequence" problems.
- Always maintain the window invariant.
- Optimize from brute force before presenting the final solution.
- Understand both the HashSet and HashMap implementations.
- Be ready to explain correctness, complexity, and trade-offs—not just write code.
- Clear communication and structured reasoning are as important as the implementation itself.