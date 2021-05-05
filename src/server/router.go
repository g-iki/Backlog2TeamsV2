package server

import (
	"Backlog2Teams/src/controller"

	"github.com/gin-gonic/gin"
)

// GetRouter ...
func GetRouter() *gin.Engine { // *gin.Engineの表記は返り値の型

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		backlog := v1.Group("/backlog")
		{
			ctrl := controller.BacklogController{}
			// 新規
			backlog.POST("", ctrl.Create)
			// 一覧
			backlog.GET("", ctrl.Index)
			// 詳細
			backlog.GET("/:id", ctrl.Show)
			// 更新
			backlog.PUT("/:id/add", ctrl.Add)
			backlog.PUT("/:id/remove", ctrl.Remove)
			// データ削除
			backlog.DELETE("/:id", ctrl.Delete)
		}
		notice := v1.Group("/notice")
		{
			ctrl := controller.NoticeController{}
			notice.POST("/do", ctrl.DoNotice)
		}
	}

	return router
}
