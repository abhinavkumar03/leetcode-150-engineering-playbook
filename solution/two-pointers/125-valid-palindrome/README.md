# Valid Palindrome

## Problem Statement

Given a string `s`, determine whether it is a palindrome after:

1. Converting all uppercase letters to lowercase.
2. Removing all non-alphanumeric characters.

Return `true` if the resulting string is a palindrome, otherwise return `false`.

A palindrome reads the same forward and backward.

### Examples

#### Example 1

Input:

s = "A man, a plan, a canal: Panama"

Output:

true

Explanation:

Processed string:

"amanaplanacanalpanama"

This reads the same from left to right and right to left.

---

#### Example 2

Input:

s = "race a car"

Output:

false

Explanation:

Processed string:

"raceacar"

Forward and backward strings are different.

---

#### Example 3

Input:

s = " "

Output:

true

Explanation:

After removing non-alphanumeric characters, the string becomes empty.

An empty string is considered a valid palindrome.

---

## Difficulty

Easy

---

## Tags

- String
- Two Pointers

---

## Pattern

**Two Pointers**

This problem is a classic example of the Two Pointers pattern where:

- One pointer starts from the beginning.
- One pointer starts from the end.
- Both pointers move toward the center.
- Invalid characters are skipped.
- Valid characters are compared.

---

## Intuition

A palindrome has symmetry.

Instead of:

1. Creating a cleaned string.
2. Reversing it.
3. Comparing both strings.

We can compare characters directly from both ends.

While moving inward:

- Ignore spaces.
- Ignore punctuation.
- Ignore symbols.
- Compare lowercase versions of valid characters.

The moment a mismatch occurs, the string cannot be a palindrome.

---

## Key Observation

The problem only cares about:

- Letters (`a-z`, `A-Z`)
- Digits (`0-9`)

Everything else can be ignored.

Therefore:

- Skip invalid characters.
- Normalize case.
- Compare remaining characters.

Using two pointers allows us to do this in a single pass.

---

# Brute Force Approach

## Idea

Create a filtered string containing only lowercase alphanumeric characters.

Then:

1. Reverse the filtered string.
2. Compare original filtered string with reversed string.

---

## Algorithm

1. Initialize empty string builder.
2. Traverse entire input.
3. Add lowercase alphanumeric characters.
4. Create reversed version.
5. Compare both strings.
6. Return result.

---

## Complexity

### Time Complexity

O(n)

- Build filtered string → O(n)
- Reverse string → O(n)

Total:

O(n)

### Space Complexity

O(n)

Additional cleaned string storage is required.

---

## Limitations

- Requires extra memory.
- Creates intermediate strings.
- Less efficient than necessary.

---

# Optimized Approach

## Idea

Use two pointers.

- Left pointer starts at beginning.
- Right pointer starts at end.

While pointers have not crossed:

1. Skip non-alphanumeric characters.
2. Convert both characters to lowercase.
3. Compare them.
4. If mismatch occurs → return false.
5. Otherwise continue inward.

If all comparisons succeed, return true.

---

## Algorithm

1. Set:

   left = 0

   right = n - 1

2. While left < right:

   - Skip invalid characters from left.
   - Skip invalid characters from right.

3. Compare lowercase characters.

4. If different:

   return false

5. Move both pointers inward.

6. Return true.

---

## Why It Works

A palindrome must satisfy:

Character(i) = Character(n - 1 - i)

for every valid position.

The two-pointer approach checks exactly this property.

Skipping non-alphanumeric characters does not affect palindrome validity because the problem explicitly ignores them.

Since every character is visited at most once, correctness and efficiency are guaranteed.

---

## Complexity

### Time Complexity

O(n)

Each character is processed at most once.

---

### Space Complexity

O(1)

No extra data structures are created.

---

# Edge Cases

## Empty Input

Input:

```text
""
```

Output:

```text
true
```

Reason:

Empty string is a valid palindrome.

---

## Single Character

Input:

```text
"a"
```

Output:

```text
true
```

Reason:

Single character is always a palindrome.

---

## Only Symbols

Input:

```text
"!!!"
```

Output:

```text
true
```

Reason:

After filtering:

```text
""
```

Empty string is valid.

---

## Duplicates

Input:

```text
"aa"
```

Output:

```text
true
```

---

## Negative-Looking Characters

Input:

```text
"-121-"
```

Processed:

```text
121
```

Output:

```text
true
```

---

## Mixed Case

Input:

```text
"Aa"
```

Output:

```text
true
```

Case normalization makes them equal.

---

## Large Inputs

Input:

Very large palindrome string.

Output:

Still works efficiently because:

- O(n) time
- O(1) extra memory

---

# Dry Run

Input:

```text
"A man, a plan, a canal: Panama"
```

Processed logically during traversal.

| Left | Right | Left Char | Right Char | Action |
|--------|--------|------------|------------|----------|
| 0 | 29 | A | a | Match |
| 2 | 28 | m | m | Match |
| 3 | 27 | a | a | Match |
| 4 | 26 | n | n | Match |
| ... | ... | ... | ... | Continue |
| Center Reached | | | | Return true |

Result:

```text
true
```

---

# Common Mistakes

## 1. Forgetting Case Conversion

Wrong:

```text
'A' != 'a'
```

Correct:

Convert both to lowercase first.

---

## 2. Not Skipping Symbols

Wrong:

```text
"A,"
```

Comparing comma directly causes failure.

---

## 3. Using Extra Space Unnecessarily

Creating filtered strings works but is not optimal.

Interviewers usually expect O(1) space.

---

## 4. Incorrect Pointer Movement

Always move pointers only after valid comparisons.

Improper movement can skip important characters.

---

## 5. Ignoring Digits

Digits are valid alphanumeric characters.

Example:

```text
"1a1"
```

Should return:

```text
true
```

---

# Interview Discussion

### Why Two Pointers?

Because palindrome checking naturally compares symmetric positions.

---

### Can We Solve It Without Extra Space?

Yes.

Two pointers provide:

```text
Time: O(n)
Space: O(1)
```

---

### What If Unicode Characters Exist?

Current solution generally assumes ASCII alphanumeric checks.

For international text, language-specific libraries may be required.

---

### Why Is Every Character Visited Once?

Each pointer only moves inward.

No pointer ever moves backward.

Therefore total work remains linear.

---

# Follow-up Questions

### 1. Can you return the longest palindromic substring?

Related to:

- Expand Around Center
- Dynamic Programming

---

### 2. Can you ignore only spaces but keep punctuation?

Modify skip conditions.

---

### 3. Can you support Unicode letters?

Use Unicode-aware character functions.

---

### 4. Can you perform the comparison recursively?

Yes, but iterative is more efficient.

---

# Real World Applications

## Data Validation

Checking normalized user inputs.

---

## Text Processing Systems

Comparing cleaned text representations.

---

## Search Engines

Preprocessing textual data.

---

## NLP Pipelines

String normalization before analysis.

---

## Data Cleaning

Ignoring formatting symbols while comparing content.

---

# Related Problems

| LeetCode | Problem | Pattern |
|-----------|-----------|-----------|
| 344 | Reverse String | Two Pointers |
| 680 | Valid Palindrome II | Two Pointers |
| 5 | Longest Palindromic Substring | Palindrome |
| 9 | Palindrome Number | Two Pointers |
| 151 | Reverse Words in a String | Two Pointers |
| 557 | Reverse Words in a String III | Two Pointers |

---