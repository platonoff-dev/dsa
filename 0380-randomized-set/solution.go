package _0380_randomized_set

import (
	"math/rand"
)

type RandomizedSet struct {
	arr []int
	set map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: []int{},
		set: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}

	pos := len(this.arr)
	this.arr = append(this.arr, val)
	this.set[val] = pos

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	pos, ok := this.set[val]
	if !ok {
		return false
	}

	last := this.arr[len(this.arr)-1]
	this.arr[pos] = last
	this.set[last] = pos
	delete(this.set, val)
	this.arr = this.arr[:len(this.arr)-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
