class Solution {
    public int jump(int[] nums) {
        int n = nums.length;

        if (n <= 1) {
            return 0;
        }

        int jumps = 0;
        int currentEnd = 0;
        int farthestReach = 0;

        for (int i = 0; i < n - 1; i++) {
            farthestReach = Math.max(farthestReach, i + nums[i]);

            if (i == currentEnd) {
                jumps++;
                currentEnd = farthestReach;

                if (currentEnd >= n - 1) {
                    break;
                }
            }
        }

        return jumps;
    }
}