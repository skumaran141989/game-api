package analytics

import (
	"net/http"

	"github.com/skumaran141989/game-api/DataAccessLayer/repos"
	"github.com/skumaran141989/game-api/utilities"
)

const (
	ERROR_FETCHING_WALL_EVENT   = "Error while fetching wall events"
	ERROR_ANALYTICS_CALCULATION = "Error while calculating analytics"
)

type Analytics struct {
	wallAttackEvent repos.WallAttackEventDB
	logger          utilities.Logger
}

func NewAnalytics(wallAttackEvent repos.WallAttackEventDB, logger utilities.Logger) *Analytics {
	return &Analytics{
		wallAttackEvent: wallAttackEvent,
	}
}

func (analytics *Analytics) ProcessSuccessfulAttacks(analyticsEngine AnalyticsEngine) (int64, int64, int64, *utilities.APIError) {
	events, err := analytics.wallAttackEvent.GetAll("active = true")

	if err != nil {
		analytics.logger.Error(ERROR_ANALYTICS_CALCULATION, err)
		return -1, -1, -1, utilities.NewAPIError(http.StatusInternalServerError, ERROR_ANALYTICS_CALCULATION, err)
	}

	startTime := events[0].CreatedDate
	endTime := events[len(events)-1].CreatedDate

	result, err := analyticsEngine.Calculate(events)

	if err != nil {
		analytics.logger.Error(ERROR_ANALYTICS_CALCULATION, err)
		return -1, -1, -1, utilities.NewAPIError(http.StatusInternalServerError, ERROR_ANALYTICS_CALCULATION, err)
	}

	return result.(int64), startTime, endTime, nil
}
