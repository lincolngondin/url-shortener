package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func HomePageHandler(ctx *gin.Context) {
	ctx.HTML(200, "home.html", gin.H{
        "state": false,
    })
}

func ShortURLHandler(ctx *gin.Context) {
	URL := ctx.PostForm("url")
	url, err := url.Parse(URL)
    
	if err != nil || !url.IsAbs() {
		ctx.HTML(200, "home.html", gin.H{
            "error_msg": "Invalid URL!",
			"state": true,
		})
		return
	}
	newURL := NewURL(URL)
	err = PutURL(newURL)
	if err != nil {
		ctx.HTML(200, "home.html", gin.H{
            "error_msg": err.Error(),
			"state": true,
		})
		return
	}
    ctx.HTML(http.StatusOK, "short.html", gin.H{
        "full_url": fmt.Sprintf("%s/%s", ctx.Request.Host, newURL.ShortenderURL),
        "track_url": fmt.Sprintf("%s/track/%s", ctx.Request.Host, newURL.ShortenderURL),
        "short_url": newURL.ShortenderURL,
    })
}


type ShortenedURL struct {
	Url string `uri:"short_url"`
}

func RedirectToURL(ctx *gin.Context) {
	var shortURL ShortenedURL = ShortenedURL{}
	ctx.ShouldBindUri(&shortURL)
	url, err := GetOriginalURLAndInclementClicks(shortURL.Url)
	if err != nil {
		ctx.Status(404)
		ctx.Writer.Write([]byte("Error: URL Not Found!"))
		return
	}
	ctx.Redirect(http.StatusPermanentRedirect, url)
}

func TrackURL(ctx *gin.Context) {
	shortURL := ShortenedURL{}
	ctx.ShouldBindUri(&shortURL)
	url, err := GetURLInfo(shortURL.Url)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		ctx.Writer.Write([]byte("Error: URL Not Found!"))
		return
	}
	ctx.HTML(http.StatusOK, "track.html", gin.H{
        "original": url.OriginalURL,
        "short": fmt.Sprintf("%s/%s", ctx.Request.Host, url.ShortenderURL),
        "clicks": url.TotalClicks,
    })
}
