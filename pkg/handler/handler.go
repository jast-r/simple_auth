package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	routers := gin.New()
	auth := routers.Group("/auth")
	{
		auth.POST("/sign_up", h.signUp)
		auth.POST("/sign_in", h.signIn)
	}

	api := routers.Group("/api")
	{
		companyList := api.Group("/company_list")
		{
			companyList.GET("/company_list", h.getAllCompany)
			companyList.PUT("/", h.updateManyCompany)
			companyList.DELETE("/", h.deleteAllCompany)
		}

		company := companyList.Group(":id/company")
		{
			company.POST("/", h.createCompany)
			company.GET("/:company_id", h.getCompanyById)
			company.PUT("/:company_id", h.updateCompany)
			company.DELETE("/:company_id", h.deleteCompany)
		}
	}

	return routers
}
