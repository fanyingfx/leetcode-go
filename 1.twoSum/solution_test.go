package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m:=make(map[int]int)
	for idx,num := range nums{
		if i,ok:=m[target-num];ok{
			return []int{i,idx}
		}
		m[num]=idx
	}
	return []int{}


}

// @lc code=end

func Test(t *testing.T) {
	assert.Equal(t, twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	assert.Equal(t, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
}
