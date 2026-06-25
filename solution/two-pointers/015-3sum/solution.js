/**
 * @param {number[]} nums
 * @return {number[][]}
 */
function threeSum(nums) {
    nums.sort((a, b) => a - b);

    const result = [];
    const n = nums.length;

    for (let i = 0; i < n - 2; i++) {

        // Skip duplicate fixed elements.
        if (i > 0 && nums[i] === nums[i - 1]) {
            continue;
        }

        let left = i + 1;
        let right = n - 1;

        while (left < right) {

            const currentSum = nums[i] + nums[left] + nums[right];

            if (currentSum < 0) {
                left++;
            } else if (currentSum > 0) {
                right--;
            } else {

                result.push([
                    nums[i],
                    nums[left],
                    nums[right]
                ]);

                left++;
                right--;

                // Skip duplicate left values.
                while (left < right && nums[left] === nums[left - 1]) {
                    left++;
                }

                // Skip duplicate right values.
                while (left < right && nums[right] === nums[right + 1]) {
                    right--;
                }
            }
        }
    }

    return result;
}