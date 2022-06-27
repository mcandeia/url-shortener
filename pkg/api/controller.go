package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mcandeia/url-shortener/pkg/shortener"
)

// ShortenRequest is the request to shorten a URL.
type ShortenRequest struct {
	// URL is the string representation of a URL.
	URL string `json:"url"`
}

// ShortenResponse is the response to shorten a URL.
type ShortenResponse struct {
	// URL is the string representation of a URL.
	URL string `json:"url"`
}

const (
	// engineQ is the engine query string specified.
	engineQ = "e"
	// aliasQ is the engine query string for alising.
	aliasQ = "alias"
	// defaultEngine is the default engine to be used.
	defaultEngine = shortener.Aliasing
	// locationHeader is the location header.
	location = "location"
)

const (
	errCouldNotBindJSONRequest = "could not handle shorten request."
)

// httpErr builds a gin response with the given error.
func httpErr(errStr string) gin.H {
	return gin.H{"error": errStr}
}

// engineFromString returns a engine from a given engine string.
func engineFromString(engineStr string) (shortener.Engine, shortener.EngineID, error) {
	engineInt, err := strconv.Atoi(engineStr)

	if err != nil {
		return nil, defaultEngine, err
	}

	engID := shortener.EngineID(engineInt)
	engine, err := shortener.Get(engID)

	return engine, engID, err
}

// Shorten returns a handler to handle the URL request.
func Shorten() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ShortenRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, httpErr(errCouldNotBindJSONRequest))
			return
		}

		validURL, parseErr := url.Parse(req.URL)

		if parseErr != nil {
			if err := ctx.ShouldBindJSON(&req); err != nil {
				ctx.JSON(http.StatusBadRequest, httpErr(errCouldNotBindJSONRequest))
				return
			}
		}

		var (
			engineID = defaultEngine
			engine   shortener.Engine
			err      error
		)

		engineStr, ok := ctx.GetQuery(engineQ)

		if ok {
			engine, engineID, err = engineFromString(engineStr)
		} else {
			engine, err = shortener.Get(defaultEngine)
		}

		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpErr(err.Error()))
			return
		}

		short, err := engine.Short(
			context.WithValue(ctx.Request.Context(), shortener.AliasKey, ctx.Query(aliasQ)), validURL.String(),
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpErr(err.Error()))
			return
		}

		relURL := fmt.Sprintf("/%d/%s", engineID, short)
		fullURL := fmt.Sprintf("http://%s%s", ctx.Request.Host, relURL)

		ctx.Header(location, relURL)
		ctx.JSON(http.StatusCreated, ShortenResponse{
			URL: fullURL,
		})
	}
}

// Redirect redirects to the long original URL.
func Redirect() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		engine, _, err := engineFromString(ctx.Param("engine"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpErr(err.Error()))
			return
		}

		long, err := engine.Long(ctx.Request.Context(), ctx.Param("short"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, httpErr(err.Error()))
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, long)
	}
}
