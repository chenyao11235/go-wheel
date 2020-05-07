package router

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "wheel/gin-demo/handler/sd"
    "wheel/gin-demo/handler/user"
    "wheel/gin-demo/router/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
    g.Use(gin.Recovery())
    g.Use(middleware.NoCache)
    g.Use(middleware.Options)
    g.Use(middleware.Secure)
    g.Use(mw...)
    g.NoRoute(func(c *gin.Context) {
        c.String(http.StatusNotFound, "The incorrect api route.")
    })

    u := g.Group("/v1/user")
    {
        u.POST("", user.Create)
        u.DELETE("/:username", user.Delete)
        u.PUT("/:id", user.Update)
        u.GET("", user.List)
        u.GET("/:username", user.Get)
    }

    svcd := g.Group("/sd")
    {
        svcd.GET("/health", sd.HealthCheck)
        svcd.GET("/disk", sd.DiskCheck)
        svcd.GET("/cpu", sd.CPUCheck)
        svcd.GET("/mem", sd.RAMCheck)
    }
    return g
}
