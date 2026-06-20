# Interview Notes — 58. Length of Last Word

# What Interviewer Is Testing

Although this is classified as an Easy problem, interviewers are usually evaluating much more than whether you can count characters.

---

## 1. String Manipulation Fundamentals

The interviewer wants to verify that you can:

- Traverse a string efficiently
- Access characters safely
- Handle boundaries correctly
- Work with whitespace conditions

Expected questions:

> How would you find the last word?

> How would you handle extra spaces?

---

## 2. Edge Case Awareness

Many candidates solve the happy path but fail on edge cases.

Examples:

```text
"Hello World"
```

Easy case.

---

```text
"Hello World   "
```

Trailing spaces.

---

```text
"leetcode"
```

Single word.

---

```text
"     "
```

Only spaces.

---

Interviewers expect candidates to proactively discuss these cases.

---

## 3. Space Optimization

A common beginner solution is:

```text
split()
```

Example:

```text
s.split(" ")
```

This works but creates extra memory.

Interviewers often ask:

> Can you solve it without creating additional arrays?

The expected answer is:

```text
Reverse Traversal
```

with:

```text
Time  : O(n)
Space : O(1)
```

---

## 4. Communication Skills

This problem is often used to evaluate how clearly you explain simple logic.

Strong candidates can:

- Explain intuition first
- Walk through examples
- Justify complexity
- Compare alternatives

---

# Typical Follow-up Questions

## Follow-up 1

### Can you solve this without using split()?

Expected Answer:

Yes.

Traverse from the end:

1. Skip trailing spaces
2. Count characters
3. Stop at a space

---

## Follow-up 2

### Why is reverse traversal better?

Answer:

```text
split():
Time  O(n)
Space O(n)

Reverse Traversal:
Time  O(n)
Space O(1)
```

No temporary arrays are created.

---

## Follow-up 3

### What if there are multiple spaces between words?

Example:

```text
"a   b   c"
```

The solution still works because it only counts the final sequence of non-space characters.

Output:

```text
1
```

---

## Follow-up 4

### What if tabs and spaces are both separators?

Instead of checking:

```text
s[i] == ' '
```

Use a whitespace check.

Examples:

```java
Character.isWhitespace(c)
```

```go
unicode.IsSpace(r)
```

This generalizes the solution.

---

## Follow-up 5

### Can this be solved in one forward pass?

Yes.

Track:

```text
currentWordLength
lastWordLength
```

When a space is encountered:

- Save current length if valid
- Reset current length

The final stored value is the answer.

---

# Optimization Journey

Interviewers like to hear how a solution evolves.

---

## Solution 1 — Split String

### Idea

Separate words into an array.

```text
"Hello World"
      ↓
["Hello", "World"]
```

Return:

```text
len(lastWord)
```

### Complexity

```text
Time  : O(n)
Space : O(n)
```

### Drawback

Creates unnecessary memory allocations.

---

## Solution 2 — Trim + Scan

### Idea

Remove trailing spaces.

Then scan backward.

### Complexity

```text
Time  : O(n)
Space : O(n)
```

Depending on language implementation, trim may create a new string.

---

## Solution 3 — Reverse Traversal (Optimal)

### Idea

Start from the end.

```text
Skip spaces
Count letters
Stop at space
```

### Complexity

```text
Time  : O(n)
Space : O(1)
```

### Interview Choice

This is the preferred solution.

---

# Whiteboard Strategy

If solving on a whiteboard:

---

## Step 1

Write an example.

```text
"Hello World  "
```

---

## Step 2

Mark the pointer.

```text
Hello World__
           ^
```

---

## Step 3

Skip spaces.

```text
Hello World
          ^
```

---

## Step 4

Count characters.

```text
d → 1
l → 2
r → 3
o → 4
W → 5
```

---

## Step 5

Hit space.

```text
Stop
```

Return:

```text
5
```

This explanation is concise and interview-friendly.

---

# Communication Tips

## Good Explanation

> Since only the final word matters, I can scan from the end instead of processing the entire string structure. First I'll skip trailing spaces, then count consecutive non-space characters until I hit another space or the beginning of the string.

This demonstrates:

- Understanding
- Optimization thinking
- Clarity

---

## Avoid Saying

> I'll split the string because it's easier.

This may lead to follow-up questions about memory optimization.

---

## Mention Complexity Explicitly

Always state:

```text
Time Complexity: O(n)
Space Complexity: O(1)
```

before finishing.

---

# Senior-Level Discussion Points

Senior engineers often expand beyond merely solving the problem.

---

## Memory Efficiency

Avoid creating:

- Arrays
- Substrings
- Temporary buffers

This reduces allocation pressure.

---

## Large-Scale Text Processing

Imagine:

```text
Log processing systems
Search engines
Streaming text pipelines
```

Repeated string allocations become expensive.

The reverse traversal approach scales better.

---

## API Design Perspective

If this logic were implemented in a utility library:

```text
lengthOfLastWord(string)
```

You would consider:

- Unicode support
- Different whitespace types
- Empty input handling
- Performance guarantees

---

# FAANG-Level Variations

Interviewers sometimes extend simple problems.

---

## Variation 1

Return the last word itself.

Example:

```text
Input:
"Hello World"

Output:
"World"
```

---

## Variation 2

Return both:

```text
Last Word
Length
```

Example:

```text
("World", 5)
```

---

## Variation 3

Handle Unicode Whitespace

Example:

```text
tab
newline
multiple whitespace characters
```

Use language-specific whitespace utilities.

---

## Variation 4

Streaming Input

Characters arrive one-by-one.

Requirements:

```text
O(1) memory
```

Track current word length dynamically.

---

## Variation 5

Find Length of the Nth Word From the End

Example:

```text
Input:
"one two three four"

n = 2

Output:
5
```

Requires additional traversal logic.

---

# Recognition Pattern

When you see:

```text
Last element
Last word
Trailing spaces
Need only final segment
```

Think:

```text
Reverse Traversal
```

instead of:

```text
Split Everything
```

---

# Interview Takeaway

This problem appears simple, but it tests:

- String fundamentals
- Boundary handling
- Optimization awareness
- Communication quality

The strongest solution is:

```text
Reverse Traversal

Time  : O(n)
Space : O(1)
```

A candidate who clearly explains the optimization journey from `split()` to reverse traversal demonstrates strong interview readiness and engineering thinking.