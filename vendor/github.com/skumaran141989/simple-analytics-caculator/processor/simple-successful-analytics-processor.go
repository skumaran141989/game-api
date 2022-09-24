package processor

import (
	"github.com/skumaran141989/simple-analytics-caculator/models"
)

type SimpleSuccessfulAttacksCalculator struct {
}

func (*SimpleSuccessfulAttacksCalculator) Calculate(input any) (any, error) {
	events := input.([]models.WallAttackEvent)
	count := 0
	currentWallSize := make(map[string]int64)
	for _, event := range events {

		if currentWallSize[event.GetWallId()] == 0 && event.GetStrength() > 0 {
			currentWallSize[event.GetWallId()] = event.GetStrength()
		}

		if currentWallSize[event.GetWallId()] < event.GetStrength() && event.GetStrength() > 0 {
			currentWallSize[event.GetWallId()] = event.GetStrength()
			count = count + 1
		}
	}

	return count, nil
}
