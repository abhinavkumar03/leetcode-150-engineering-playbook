# Hash Map Pattern

## Pattern Definition

The Hash Map pattern uses a key-value data structure to achieve constant-time lookups during problem solving.

A hash map is useful when:

- Values need to be retrieved quickly.
- Repeated searches would otherwise require linear scans.
- Keys can uniquely identify data.
- Symbolic representations need conversion into numeric or structured values.

Typical operation costs:

| Operation | Complexity |
|------------|------------|
| Insert | O(1) |
| Lookup | O(1) |
| Update | O(1) |
| Delete | O(1) |

---

# Recognition Signals

Look for this pattern when the problem contains:

✅ Character-to-value mapping

✅ Number-to-frequency tracking

✅ Fast lookups

✅ Duplicate detection

✅ Key-based retrieval

✅ Symbol conversion

Common interview phrases:

```text
Find duplicates

Lookup value

Count frequency

Map characters

Transform symbols

Store relationships
```

---

# Generic Template

```go
lookup := map[KeyType]ValueType{}

for _, item := range data {
    value := lookup[item]
}
```

---

# Complexity

| Operation | Time |
|------------|------|
| Lookup | O(1) |
| Insert | O(1) |
| Update | O(1) |

Space Complexity:

```text
O(k)
```

Where:

```text
k = number of stored keys
```

---

# Common Pitfalls

### Pitfall 1

Using arrays when keys are not sequential.

---

### Pitfall 2

Forgetting default values.

---

### Pitfall 3

Using expensive searches instead of direct lookup.

---

### Pitfall 4

Overcomplicating logic when a map provides direct access.

---

# Related Problems

| Problem | Difficulty | Usage |
|----------|------------|--------|
| 1. Two Sum | Easy | Value lookup |
| 13. Roman to Integer | Easy | Symbol lookup |
| 205. Isomorphic Strings | Easy | Character mapping |
| 217. Contains Duplicate | Easy | Membership checking |
| 242. Valid Anagram | Easy | Frequency counting |
| 290. Word Pattern | Easy | Bidirectional mapping |

---

# Added Problem

## 13. Roman to Integer

### Why It Uses Hash Maps

Roman symbols naturally map to fixed values.

```text
I → 1
V → 5
X → 10
L → 50
C → 100
D → 500
M → 1000
```

A hash map enables:

```text
O(1) lookup
```

for each character during traversal.

### Key Insight

```text
current < next
    => subtract

otherwise
    => add
```

### Complexity

```text
Time  : O(n)
Space : O(1)
```