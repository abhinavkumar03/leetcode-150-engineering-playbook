/**
 * @param {number[]} nums
 * @return {number}
 */
function removeDuplicates(nums) {
    if (nums.length <= 2) {
        return nums.length;
    }

    let write = 2;

    for (let read = 2; read < nums.length; read++) {
        if (nums[read] !== nums[write - 2]) {
            nums[write] = nums[read];
            write++;
        }
    }

    return write;
}