package router

import (
	"log"
	"ultradex/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) error {
	r := gin.Default()

	r.LoadHTMLGlob("web/*.html")
	r.Static("/assets", "./web/assets")

	setUltramanController(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
		return err
	}

	return nil
}

func setUltramanController(router *gin.Engine, db *gorm.DB) {
	ultraman := controller.NewUltramanController(db)

	// Create a group for ultraman-related routes
	ultramanGroup := router.Group("/ultraman")
	{
		// Route for listing all ultramans
		ultramanGroup.GET("", func(c *gin.Context) {
			ultramanList, err := ultraman.GetUltramanList(c)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.HTML(200, "index.html", gin.H{
				"ultramanList": ultramanList,
			})
		})

		// Route for viewing details of a specific ultraman by slug
		ultramanGroup.GET("/details/:slug", func(c *gin.Context) {
			ultraman, err := ultraman.GetUltraman(c)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.HTML(200, "details.html", gin.H{
				"ultraman": ultraman,
			})
		})
	}

	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/ultraman")
	})
}
