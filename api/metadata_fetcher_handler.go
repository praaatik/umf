package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praaatik/umf/metadatafetcher"
)

// MetadataFetcherHandler returns the metadata of the URL back to the caller as a JSON response.
// This function returns a 200 code if the fetching was successful along with the fetched Metadata.
// If this function is unable to fetch the data, 500 code is returned along with empty struct.
// No errors are thrown to avoid any issues on the frontend. It is possible the URL is not accessible by this server but valid and need to save it.
func (server *Server) MetadataFetcherHandler(ctx *gin.Context) {
	inputUrl := ctx.Query("url")
	if inputUrl == "" {
		ctx.JSON(http.StatusInternalServerError, metadatafetcher.OpenGraphMetaData{})
	}

	Metadata, err := metadatafetcher.FetchHTML(inputUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, metadatafetcher.OpenGraphMetaData{})
	}

	ctx.JSON(http.StatusOK, Metadata)
}
