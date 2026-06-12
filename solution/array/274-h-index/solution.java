import java.util.Arrays;

class Solution {
    public int hIndex(int[] citations) {
        Arrays.sort(citations);

        int n = citations.length;
        int h = 0;

        for (int i = n - 1; i >= 0; i--) {
            int papersWithAtLeastCurrentCitation = n - i;

            if (citations[i] >= papersWithAtLeastCurrentCitation) {
                h = papersWithAtLeastCurrentCitation;
            } else {
                break;
            }
        }

        return h;
    }
}