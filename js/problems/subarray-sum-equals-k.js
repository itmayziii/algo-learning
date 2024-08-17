/**
 * Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
 *
 * A subarray is a contiguous non-empty sequence of elements within an array.
 *
 * Example 1:
 * Input: nums = [1,1,1], k = 2
 * Output: 2
 *
 * Example 2:
 * Input: nums = [1,2,3], k = 3
 * Output: 2
 *
 * ----------------------------------------
 * This is by far the best explanation of the problem I could find to achieve O(n) time and space complexity.
 * https://algo.monster/liteproblems/560
 *
 * Every other video / article seem to fumble through the algorithm.
 */

/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var subarraySum = function (nums, k) {
  const sumFrequency = new Map([[0, 1]])
  let prefixSum = 0
  let answer = 0
  for (const num of nums) {
    prefixSum += num
    if (sumFrequency.has(prefixSum - k)) {
      answer += sumFrequency.get(prefixSum - k)
    }
    sumFrequency.set(prefixSum, (sumFrequency.get(prefixSum) ?? 0) + 1)
  }

  return answer
}
