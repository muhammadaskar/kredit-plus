package routes

import (
	"github.com/gin-gonic/gin"
	consumerDeliver "github.com/muhammadaskar/kredit-plus/app/consumer/delivery/http"
	consumerRepo "github.com/muhammadaskar/kredit-plus/app/consumer/repository/mysql"
	consumerUseCase "github.com/muhammadaskar/kredit-plus/app/consumer/usecase"
	tenorRepo "github.com/muhammadaskar/kredit-plus/app/tenor/repository/mysql"
	transactionDeliver "github.com/muhammadaskar/kredit-plus/app/transaction/delivery/http"
	transactionRepo "github.com/muhammadaskar/kredit-plus/app/transaction/repository/mysql"
	transactionUseCase "github.com/muhammadaskar/kredit-plus/app/transaction/usecase"
	mysqldriver "github.com/muhammadaskar/kredit-plus/infrastructures/mysql_driver"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	db := mysqldriver.InitDatabase()

	transactionRepository := transactionRepo.NewRepository(db)
	tenorRepository := tenorRepo.NewRepository(db)
	consumerRepository := consumerRepo.NewRepository(db)

	transactionUseCase := transactionUseCase.NewUseCase(transactionRepository, tenorRepository, consumerRepository)
	consumerUseCase := consumerUseCase.NewUseCase(consumerRepository)

	transactionHandler := transactionDeliver.NewTransactionHandler(transactionUseCase)
	consumerHandler := consumerDeliver.NewConsumerHandler(consumerUseCase)

	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "welcome to api",
		})
	})

	api := router.Group("api/v1")
	{
		api.POST("transaction", transactionHandler.Create)

		api.GET("consumer/:id", consumerHandler.FindById)
	}

	return router
}
