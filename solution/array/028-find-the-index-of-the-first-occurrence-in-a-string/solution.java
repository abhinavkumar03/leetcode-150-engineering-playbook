class Solution {
    public int strStr(String haystack, String needle) {
        int n = haystack.length();
        int m = needle.length();

        if (m == 0) {
            return 0;
        }

        for (int start = 0; start <= n - m; start++) {
            boolean match = true;

            for (int offset = 0; offset < m; offset++) {
                if (haystack.charAt(start + offset) != needle.charAt(offset)) {
                    match = false;
                    break;
                }
            }

            if (match) {
                return start;
            }
        }

        return -1;
    }
}