var rotate = function (nums, k) {
    const n = nums.length;

    if (n <= 1) {
        return;
    }

    k %= n;

    reverse(nums, 0, n - 1);
    reverse(nums, 0, k - 1);
    reverse(nums, k, n - 1);
};

function reverse(nums, left, right) {
    while (left < right) {
        [nums[left], nums[right]] = [nums[right], nums[left]];

        left++;
        right--;
    }
}