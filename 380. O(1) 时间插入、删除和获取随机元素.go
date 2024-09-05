package main

import "math/rand"

type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:       []int{},
		valToIndex: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valToIndex[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.valToIndex[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.valToIndex[val]
	if !ok {
		return false
	}
	length := len(this.nums)
	this.valToIndex[this.nums[length-1]] = index
	this.nums[index], this.nums[length-1] = this.nums[length-1], this.nums[index]
	this.nums = this.nums[:length-1]
	delete(this.valToIndex, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
