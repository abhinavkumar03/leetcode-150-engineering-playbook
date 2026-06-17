/**
 * @param {number[]} height
 * @return {number}
 */
function trap(height) {
    if (height.length < 3) {
        return 0;
    }

    let left = 0;
    let right = height.length - 1;

    let leftMax = 0;
    let rightMax = 0;

    let trappedWater = 0;

    while (left < right) {
        leftMax = Math.max(leftMax, height[left]);
        rightMax = Math.max(rightMax, height[right]);

        if (leftMax < rightMax) {
            trappedWater += leftMax - height[left];
            left++;
        } else {
            trappedWater += rightMax - height[right];
            right--;
        }
    }

    return trappedWater;
}