package analytics

import (
	"net/http"

	"github.com/game-api/DataAccessLayer/repos"
	"github.com/game-api/utilities"
)

const (
	ERROR_ANALYTICS_CALCULATION = "Error while calculating analytics"
)

type Analytics struct {
	wallAttackEvent repos.WallAttackEventDB
	tribe           repos.TribeDB
	wall            repos.WallDB
	analyticsEngine AnalyticsEngine
	logger          utilities.Logger
}

func NewAnalytics(wallAttackEvent repos.WallAttackEventDB, tribe repos.TribeDB, wall repos.WallDB, logger utilities.Logger, analyticsEngine AnalyticsEngine) *Analytics {
	return &Analytics{
		wallAttackEvent: wallAttackEvent,
		tribe:           tribe,
		wall:            wall,
		analyticsEngine: analyticsEngine,
	}
}

func (analytics *Analytics) Process() any {
	events, err := analytics.wallAttackEvent.GetAll("active = true")

	if err != nil {
		analytics.logger.Error(ERROR_ANALYTICS_CALCULATION, err)

		return utilities.NewAPIError(http.StatusInternalServerError, ERROR_ANALYTICS_CALCULATION, err)
	}

	analytics.analyticsEngine.Calculate(events)

	return nil
}
