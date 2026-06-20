# Longest Common Prefix

## Problem Statement

Given an array of strings `strs`, find the longest common prefix shared among all strings.

If there is no common prefix, return an empty string `""`.

### Example 1

```text
Input: strs = ["flower","flow","flight"]
Output: "fl"
```

### Example 2

```text
Input: strs = ["dog","racecar","car"]
Output: ""
```

### Constraints

```text
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of lowercase English letters.
```

---

## Difficulty

**Easy**

---

## Tags

* String
* Array

---

## Pattern

**Primary Pattern:** Horizontal Scanning

**Secondary Pattern:** String Processing

---

## Intuition

Every string must share the same prefix.

Instead of comparing every character of every string independently, we can start with the first string as a candidate prefix and continuously reduce it until all strings contain that prefix.

If a string does not start with the current prefix, remove the last character from the prefix and try again.

Eventually, the remaining prefix is the longest common prefix shared by all strings.

---

## Key Observation

The longest common prefix can never be longer than the shortest matching prefix encountered during comparisons.

Therefore:

1. Use the first string as the initial prefix.
2. Compare it against every other string.
3. Shrink the prefix whenever a mismatch occurs.
4. Stop early if the prefix becomes empty.

---

## Brute Force Approach

### Algorithm

1. Find the shortest string.
2. Check characters column by column.
3. Compare the same index across all strings.
4. Stop when a mismatch occurs.
5. Return the prefix collected so far.

### Complexity

#### Time Complexity

```text
O(N × M)
```

Where:

* N = number of strings
* M = length of shortest string

#### Space Complexity

```text
O(1)
```

### Limitations

* Requires checking every string for every character position.
* More comparisons than necessary in some cases.
* Less intuitive to explain during interviews.

---

## Optimized Approach

### Algorithm

1. Initialize `prefix = strs[0]`.
2. Iterate through remaining strings.
3. While the current string does not start with `prefix`:

   * Remove the last character from `prefix`.
4. If `prefix` becomes empty:

   * Return `""`.
5. After processing all strings, return `prefix`.

### Why It Works

The algorithm maintains an invariant:

```text
prefix = longest common prefix among processed strings
```

Whenever a mismatch occurs, the prefix is reduced until it becomes valid for the current string.

After all strings are processed, the remaining prefix is guaranteed to be shared by every string.

### Complexity

#### Time Complexity

```text
O(N × M)
```

Where:

* N = number of strings
* M = length of the longest common prefix candidate

#### Space Complexity

```text
O(1)
```

Only a few variables are used.

---

## Edge Cases

### Empty Input

```text
Input: []
Output: ""
```

Return an empty string immediately.

---

### Single Element

```text
Input: ["leetcode"]
Output: "leetcode"
```

The string itself is the common prefix.

---

### Duplicates

```text
Input: ["test","test","test"]
Output: "test"
```

Entire string remains the prefix.

---

### Negative Values

Not applicable because inputs contain strings only.

---

### Large Inputs

```text
200 strings
each length = 200
```

The algorithm remains efficient because the prefix only shrinks and never expands.

---

## Dry Run

### Example

```text
strs = ["flower","flow","flight"]
```

Initial:

| Step  | Prefix |
| ----- | ------ |
| Start | flower |

---

Compare with `"flow"`

| Action                    | Prefix |
| ------------------------- | ------ |
| flower not prefix of flow | flowe  |
| flowe not prefix of flow  | flow   |
| flow matches              | flow   |

---

Compare with `"flight"`

| Action        | Prefix |
| ------------- | ------ |
| flow mismatch | flo    |
| flo mismatch  | fl     |
| fl matches    | fl     |

---

Final Result

```text
"fl"
```

---

## Common Mistakes

### Forgetting Empty Array Handling

```text
strs = []
```

Can cause index out of range errors.

---

### Comparing Entire Strings Repeatedly

This increases unnecessary operations.

---

### Incorrect Prefix Shrinking

Removing characters from the wrong side:

```text
Correct:
prefix = prefix[:len(prefix)-1]
```

---

### Not Stopping When Prefix Becomes Empty

Once the prefix becomes empty, no further work is needed.

---

## Interview Discussion

### Why Use Horizontal Scanning?

* Simple to implement.
* Easy to explain.
* Efficient enough for constraints.

### Alternative Solutions

#### Vertical Scanning

Compare characters column-by-column.

#### Divide and Conquer

Split strings into groups and merge prefixes.

#### Binary Search on Prefix Length

Search for maximum valid prefix length.

These alternatives are useful for demonstrating deeper problem-solving skills.

---

## Follow-up Questions

1. How would you solve this using vertical scanning?
2. Can you implement a divide-and-conquer solution?
3. Can you use a Trie?
4. Which approach performs best when strings are extremely long?
5. How would Unicode characters affect the solution?

---

## Real World Applications

### Search Engines

Finding common query prefixes.

### Auto-Completion Systems

Generating suggestions from typed text.

### IDE Code Completion

Detecting shared symbol prefixes.

### File Systems

Grouping files by common path prefixes.

### DNA Sequence Analysis

Finding shared sequence beginnings.

---

## Related Problems

* Two Sum (#1)
* Implement Trie (Prefix Tree) (#208)
* Search Suggestions System (#1268)
* Longest Common Subsequence (#1143)
* Word Search II (#212)

These problems expand understanding of string matching and prefix-based algorithms.
