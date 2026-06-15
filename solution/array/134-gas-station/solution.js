/**
 * @param {number[]} gas
 * @param {number[]} cost
 * @return {number}
 */
function canCompleteCircuit(gas, cost) {
    if (!gas || !cost || gas.length === 0) {
        return -1;
    }

    let totalBalance = 0;
    let currentTank = 0;
    let startStation = 0;

    for (let i = 0; i < gas.length; i++) {
        const fuelDifference = gas[i] - cost[i];

        totalBalance += fuelDifference;
        currentTank += fuelDifference;

        if (currentTank < 0) {
            startStation = i + 1;
            currentTank = 0;
        }
    }

    return totalBalance >= 0 ? startStation : -1;
}