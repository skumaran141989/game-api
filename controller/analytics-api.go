package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skumaran141989/game-api/BusinesslogicLayer/analytics"
	"github.com/skumaran141989/game-api/DataAccessLayer/repos"
	"github.com/skumaran141989/game-api/controller/response"
	"github.com/skumaran141989/game-api/utilities"
	"github.com/skumaran141989/simple-analytics-caculator/processor"
)

var creds *utilities.DBCreds

func Analytics(router *gin.Engine, creds *utilities.DBCreds) {
	router.GET("v1/successful-attacks", getSuccessfulAttacks)
}

func getSuccessfulAttacks(context *gin.Context) {
	simpleprocessor := &processor.SimpleSuccessfulAttacksCalculator{}

	logger := utilities.GetConsoleLoggerInstance()
	database, connerr := utilities.GetDBInstance(creds)

	errorresponse := response.Error{}

	if connerr != nil {
		errorresponse.FailureCode = "0001"
		errorresponse.Message = "Database Connection Failure" + connerr.Error()

		context.JSON(http.StatusInternalServerError, errorresponse)

		return
	}

	wallAttackEvent := repos.NewWallAttackEvent(database)
	analytics := analytics.NewAnalytics(wallAttackEvent, logger)

	count, starttime, endtime, err := analytics.ProcessSuccessfulAttacks(simpleprocessor)

	if err != nil {
		errorresponse.FailureCode = "0002-app-err"
		errorresponse.Message = "Database Connection Failure" + connerr.Error()

		context.JSON(int(err.GetCode()), errorresponse)

		return
	}

	response := &response.SimpleSuccefulAttacksCount{}
	response.Count = count
	response.StartTime = starttime
	response.EndTime = endtime

	context.JSON(http.StatusOK, response)

	return
}
