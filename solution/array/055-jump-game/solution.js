/**
 * @param {number[]} nums
 * @return {boolean}
 */
function canJump(nums) {
    let farthestReach = 0;
    const lastIndex = nums.length - 1;

    for (let index = 0; index < nums.length; index++) {
        if (index > farthestReach) {
            return false;
        }

        farthestReach = Math.max(
            farthestReach,
            index + nums[index]
        );

        if (farthestReach >= lastIndex) {
            return true;
        }
    }

    return true;
}