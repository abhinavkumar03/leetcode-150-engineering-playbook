package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	values    []int
	valueToIndex map[int]int
	randomizer *rand.Rand
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		values:       make([]int, 0),
		valueToIndex: make(map[int]int),
		randomizer:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.valueToIndex[val]; exists {
		return false
	}

	rs.values = append(rs.values, val)
	rs.valueToIndex[val] = len(rs.values) - 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	index, exists := rs.valueToIndex[val]
	if !exists {
		return false
	}

	lastIndex := len(rs.values) - 1
	lastValue := rs.values[lastIndex]

	rs.values[index] = lastValue
	rs.valueToIndex[lastValue] = index

	rs.values = rs.values[:lastIndex]
	delete(rs.valueToIndex, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	randomIndex := rs.randomizer.Intn(len(rs.values))
	return rs.values[randomIndex]
}