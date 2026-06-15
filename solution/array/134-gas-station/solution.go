package main

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(cost) == 0 {
		return -1
	}

	totalBalance := 0
	currentTank := 0
	startStation := 0

	for i := 0; i < len(gas); i++ {
		fuelDifference := gas[i] - cost[i]

		totalBalance += fuelDifference
		currentTank += fuelDifference

		if currentTank < 0 {
			startStation = i + 1
			currentTank = 0
		}
	}

	if totalBalance < 0 {
		return -1
	}

	return startStation
}