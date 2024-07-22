package api

import "github.com/gin-gonic/gin"

func (server *Server) generateRoutes() {
	router := gin.Default()

	router.GET("/v1/healthcheck", server.HealthCheckHandler)
	/*
			*
			* fetch the metadata for a url
		  * this should save it to the database AND return back the metadata as JSON
			* router.GET("/v1/metadata/fetch", {url})
			*
			* refresh the metadata for a url
		  * this should save it to the database AND return back the metadata as JSON
			* router.GET("/v1/metadata/refresh", {url})
			*
			* refresh the metadata for all urls for a user
		  * this should save it to the database AND return back the metadata as JSON
			* router.GET("/v1/metadata/refresh/all")
			*
	*/
	server.router = router
}
