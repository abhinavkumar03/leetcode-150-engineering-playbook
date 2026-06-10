class Solution {
    public boolean canJump(int[] nums) {
        int farthestReach = 0;
        int lastIndex = nums.length - 1;

        for (int index = 0; index < nums.length; index++) {
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
}