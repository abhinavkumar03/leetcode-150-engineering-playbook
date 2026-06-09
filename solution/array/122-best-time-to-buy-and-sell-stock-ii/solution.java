class Solution {
    public int maxProfit(int[] prices) {
        if (prices == null || prices.length < 2) {
            return 0;
        }

        int totalProfit = 0;

        for (int day = 1; day < prices.length; day++) {
            int priceDifference = prices[day] - prices[day - 1];

            if (priceDifference > 0) {
                totalProfit += priceDifference;
            }
        }

        return totalProfit;
    }
}