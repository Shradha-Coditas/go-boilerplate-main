package post

import (
	"boilerplate/models/albums"
	"boilerplate/util/genericmodel"
	"boilerplate/util/localization"
	"errors"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handler - Handle the Post requests at /v1/album
// @Summary store data of album request
// @Description Api for insert album data.
// @Tags Albums
// @Accept  json
// @Produce  json
// @Param authorization header string true "Authorization header"
// @Param body body albums.AlbumRequest true "Payload for Album"
// @Success 200 {object} albums.AlbumResponse
// @Failure 400 {object} genericmodel.ErrResponse "Invalid Payload."
// @Failure 401 {object} genericmodel.ErrResponse "Unauthorized."
// @Failure 406 {object} genericmodel.ErrResponse "Token Expired."
// @Failure 500 {object} genericmodel.ErrResponse "Internal Server Error."
// @Failure 409 {object} genericmodel.ErrResponse "Conflict."
// @Failure 503 {object} genericmodel.ErrResponse "service Unavailable."
// @Router /album [POST]
func Handler(c *gin.Context) {
	//Language in header
	lang, _ := c.Get("language")

	//Common Error Response
	errResponse := genericmodel.ErrResponse{
		Message: localization.GetMessage(lang, "common.500", nil),
	}

	//Reqiest payload variable
	var RequstPayload albums.AlbumRequest

	//Custom Validator
	if err := c.ShouldBindJSON(&RequstPayload); err != nil {
		var verr validator.ValidationErrors
		fields := []string{}
		if errors.As(err, &verr) {
			for _, f := range verr {
				fields = append(fields, f.Field())
			}
		}
		errResponse.Message = localization.GetMessage(lang, "common.400", map[string]interface{}{
			"Fields": strings.Join(fields, ", "),
		})
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}
	rowsInserted, err := albums.InsertAlbum(RequstPayload)
	if err != nil {
		apiRes := genericmodel.ErrResponse{
			Message: localization.GetMessage(lang, "common.503", nil),
		}
		c.JSON(http.StatusServiceUnavailable, apiRes)
		return
	}
	if rowsInserted {
		apiRes := albums.AlbumResponse{
			Message: localization.GetMessage(lang, "common.201", nil),
		}
		c.JSON(http.StatusOK, apiRes)
		return
	} else {
		c.JSON(http.StatusInternalServerError, errResponse)
		return
	}
}
