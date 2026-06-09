/**
 * @param {number[]} prices
 * @return {number}
 */
function maxProfit(prices) {
    if (!prices || prices.length < 2) {
        return 0;
    }

    let totalProfit = 0;

    for (let day = 1; day < prices.length; day++) {
        const priceDifference = prices[day] - prices[day - 1];

        if (priceDifference > 0) {
            totalProfit += priceDifference;
        }
    }

    return totalProfit;
}