class Solution {
    public int canCompleteCircuit(int[] gas, int[] cost) {
        if (gas == null || cost == null || gas.length == 0) {
            return -1;
        }

        int totalBalance = 0;
        int currentTank = 0;
        int startStation = 0;

        for (int i = 0; i < gas.length; i++) {
            int fuelDifference = gas[i] - cost[i];

            totalBalance += fuelDifference;
            currentTank += fuelDifference;

            if (currentTank < 0) {
                startStation = i + 1;
                currentTank = 0;
            }
        }

        return totalBalance >= 0 ? startStation : -1;
    }
}