package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mehdijoafshani/go-assessment-1/account"
	"github.com/mehdijoafshani/go-assessment-1/internal/config"
	"github.com/mehdijoafshani/go-assessment-1/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// #SOLID: S
// If the number of APIs increases (ex. a new set of APIs to handle a new business area) it sounds reasonable to separate them into different files
// To follow Single Responsibility Principle

// #Desc
// Maybe defining a separate struct, including the following manager var inside it and implement an interface would sound a good idea
// However, each API (REST, gRPC, ...) would need a separate set of arguments in their methods, I decided to keep it as it is.
var manager balanceManager

func create(c *gin.Context) {
	numberS := c.DefaultQuery("number", strconv.Itoa(config.Data.DefaultAccountNumbers))

	number, err := strconv.Atoi(numberS)
	if err != nil {
		logger.Zap().Debug("non-numeric parameter for the number of accounts to create",
			zap.String("numbers", numberS),
			zap.Error(err))

		c.String(http.StatusUnprocessableEntity, "number ought to be numeric")
		return
	}

	err = manager.Create(number)
	if err != nil {
		logger.Zap().Error("failed to create accounts", zap.Error(err))
		// TODO: return more explicit error code
		// Temporarily I will return the error message to the client
		c.String(http.StatusInternalServerError, "failed to create accounts, err: %s", err)
		return
	}

	c.String(http.StatusOK, "%d accounts are created", number)
}

func getBalance(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Zap().Debug("non-numeric parameter for an account's id to get balance",
			zap.String("id", idStr),
			zap.Error(err))

		c.String(http.StatusUnprocessableEntity, "id ought to be numeric")
		return
	}

	balance, err := manager.Get(id)
	if err != nil {
		logger.Zap().Error("failed to get balance", zap.Error(err))
		// TODO: return more explicit error code
		c.String(http.StatusInternalServerError, "failed to get balance")
		return
	}

	c.String(http.StatusOK, "%d", balance)
}

func getAllBalances(c *gin.Context) {
	balance, err := manager.GetAll()
	if err != nil {
		logger.Zap().Error("failed to get all balances", zap.Error(err))
		// TODO: return more explicit error code
		c.String(http.StatusInternalServerError, "failed to get all balances")
		return
	}

	c.String(http.StatusOK, "%d", balance)
}

func addBalance(c *gin.Context) {
	idS := c.Query("id")
	balanceS := c.Query("balance")

	id, err := strconv.Atoi(idS)
	if err != nil {
		logger.Zap().Debug("non-numeric parameter for an account's id to add balance",
			zap.String("id", idS),
			zap.Error(err))

		c.String(http.StatusUnprocessableEntity, "id ought to be numeric")
		return
	}

	balance, err := strconv.Atoi(balanceS)
	if err != nil {
		logger.Zap().Debug("non-numeric parameter for the amount to add balance",
			zap.String("balance", balanceS),
			zap.Error(err))

		c.String(http.StatusUnprocessableEntity, "extra balance ought to be numeric")
		return
	}

	if balance < 0 {
		c.String(http.StatusUnprocessableEntity, "extra balance ought to be positive")
		return
	}

	err = manager.Add(balance, id)
	if err != nil {
		logger.Zap().Error("failed to add balance", zap.Error(err))
		// TODO: return more explicit error code
		c.String(http.StatusInternalServerError, "failed to add balance")
		return
	}

	c.String(http.StatusOK, "the extra balance %d  is applied to id %d", balance, id)
}

func addToAllBalances(c *gin.Context) {
	balanceS := c.Query("balance")

	balance, err := strconv.Atoi(balanceS)
	if err != nil {
		logger.Zap().Debug("non-numeric parameter for the amount to add balance",
			zap.Error(err))

		c.String(http.StatusUnprocessableEntity, "extra balance ought to be numeric")
		return
	}

	if balance < 0 {
		c.String(http.StatusUnprocessableEntity, "extra balance ought to be positive")
		return
	}

	err = manager.AddToAll(balance)
	if err != nil {
		logger.Zap().Error("failed to add balance", zap.Error(err))
		// TODO: return more explicit error code
		c.String(http.StatusInternalServerError, "failed to add balance")
		return
	}

	c.String(http.StatusOK, "the extra balance %d is applied to all accounts", balance)
}

func StartRestServer() error {
	manager = account.CreateManager(config.Data.IsConcurrent)

	gin.SetMode(gin.ReleaseMode)
	// TODO fix gin logger

	r := gin.New()
	r.POST("/create", create)
	r.GET("/getBalance", getBalance)
	r.GET("/getAllBalances", getAllBalances)
	r.PUT("/addBalance", addBalance)
	r.PUT("/addToAllBalances", addToAllBalances)

	err := r.Run(":" + config.Data.RestPort)
	if err != nil {
		return err
	}

	return nil
}
