# LeetCode

1. Two Sum

```javascript
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  const map = new Map()
  for (let i = 0; i < nums.length; i++) {
    const matching = map.get(target - nums[i])
    if (matching != null) {
      return [matching, i]
    }

    map.set(nums[i], i)
  }

  return []
}
```
