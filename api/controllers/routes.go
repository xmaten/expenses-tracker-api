package controllers

import "github.com/xmaten/expenses-tracker-api/api/middlewares"

func (s *Server) initializeRoutes() {
	v1 := s.Router.Group("/api/v1")
	{
		v1.POST("/login", s.Login)

		v1.POST("/users", s.CreateUser)
		v1.GET("/users/:id", middlewares.TokenAuthMiddleware(), s.GetUser)

		v1.POST("/expenses", middlewares.TokenAuthMiddleware(), s.CreateExpense)
		v1.GET("/expenses", middlewares.TokenAuthMiddleware(), s.GetUserExpenses)

		v1.POST("/incomes", middlewares.TokenAuthMiddleware(), s.CreateIncome)
		v1.GET("/incomes", middlewares.TokenAuthMiddleware(), s.GetUserIncomes)
	}
}