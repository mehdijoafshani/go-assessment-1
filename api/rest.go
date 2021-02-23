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

const (
	noId = -1
)

var manager balanceManager

func accountsPost(c *gin.Context) {
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

// #Desc
// To follow REST best practices the endpoints represent resources, that's why one func (accountsGet for instance) is declared to
// get one account's balance, all balances' sum or all balances in list
// I also could have separate them into sub functions to better segregation
func accountsGet(c *gin.Context) {
	res := c.Query("result")

	idStr := c.Query("id")

	var id int
	var err error
	id, err = strconv.Atoi(idStr)
	if err != nil {
		id = noId
	}

	// if there is an id in the request (getting only one balance)
	if id != noId {
		balance, err := manager.Get(id)
		if err != nil {
			logger.Zap().Error("failed to get balance", zap.Error(err))
			// TODO: return more explicit error code
			c.String(http.StatusInternalServerError, "failed to get balance")
			return
		}
		c.String(http.StatusOK, "%d", balance)
		return
	}

	// getting all balances
	switch res {
	// to return sum of all balances
	case "aggregate":
		balance, err := manager.GetAll()
		if err != nil {
			logger.Zap().Error("failed to get all balances", zap.Error(err))
			// TODO: return more explicit error code
			c.String(http.StatusInternalServerError, "failed to get all balances")
			return
		}
		c.String(http.StatusOK, "%d", balance)
	// to return the list all balances
	case "list":
		// TODO: declare business api to support returning a list all the accounts (ids and values)
		// Its difference in concurrent mode would be the way we protect the global slice in the memory (which all go routines will add an account to)
		// We need to use locking (mutex) around the piece of code that inserts into the slice in that case
	default:
		c.String(http.StatusUnprocessableEntity, "please specify the result mod (aggregate,...)")
	}
}

func accountsPut(c *gin.Context) {
	balanceS := c.Query("increase")
	idS := c.Query("id")

	balance, err := strconv.Atoi(balanceS)
	if err != nil {
		logger.Zap().Debug("non-numeric parameter for the amount to add balance",
			zap.Error(err))

		c.String(http.StatusUnprocessableEntity, "extra balance ought to be numeric")
		return
	}

	var id int
	id, err = strconv.Atoi(idS)
	if err != nil {
		id = noId
	}

	if id == noId {
		//add to all balances
		err = manager.AddToAll(balance)
		if err != nil {
			logger.Zap().Error("failed to add balance", zap.Error(err))
			// TODO: return more explicit error code
			c.String(http.StatusInternalServerError, "failed to add balance")
			return
		}
		c.String(http.StatusOK, "the extra balance %d is applied to all accounts", balance)
		return
	} else {
		// apply on one specific balance
		err = manager.Add(balance, id)
		if err != nil {
			logger.Zap().Error("failed to add balance", zap.Error(err))
			// TODO: return more explicit error code
			c.String(http.StatusInternalServerError, "failed to add balance")
			return
		}
		c.String(http.StatusOK, "the extra balance %d is applied to account with id %d", balance, id)
	}
}

func StartRestServer() error {
	manager = account.CreateManager(config.Data.IsConcurrent)

	gin.SetMode(gin.ReleaseMode)
	// TODO fix gin logger

	r := gin.New()
	r.POST("/accounts", accountsPost)
	r.GET("/accounts", accountsGet)
	r.PUT("/accounts", accountsPut)

	err := r.Run(":" + config.Data.RestPort)
	if err != nil {
		return err
	}

	return nil
}
