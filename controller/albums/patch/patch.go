package patch

import (
	"boilerplate/models/albums"
	"boilerplate/util/genericmodel"
	"boilerplate/util/localization"
	"boilerplate/util/logwrapper"
	"errors"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handler - Handle the PATCH requests at /v1/album
// @Summary Update data of album request
// @Description Api for update album data.
// @Tags Albums
// @Accept  json
// @Produce  json
// @Param authorization header string true "Authorization header"
// @Param body body albums.AlbumUpdateRequest true "Payload for Album"
// @Success 200 {object} albums.AlbumResponse
// @Failure 400 {object} genericmodel.ErrResponse "Invalid Payload."
// @Failure 401 {object} genericmodel.ErrResponse "Unauthorized."
// @Failure 406 {object} genericmodel.ErrResponse "Token Expired."
// @Failure 500 {object} genericmodel.ErrResponse "Internal Server Error."
// @Failure 409 {object} genericmodel.ErrResponse "Conflict."
// @Failure 404 {object} genericmodel.ErrResponse "Not Found."
// @Failure 503 {object} genericmodel.ErrResponse "service Unavailable."
// @Router /album [PATCH]
func Handler(c *gin.Context) {
	//Language in header
	lang, _ := c.Get("language")

	//Common Error Response
	errResponse := genericmodel.ErrResponse{
		Message: localization.GetMessage(lang, "common.500", nil),
	}

	//Reqiest payload variable
	var RequstPayload albums.AlbumUpdateRequest

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
	rowsInserted, err := albums.UpdateAlbum(RequstPayload)
	if err != nil {
		logwrapper.Logger.Debugln(err)
		apiRes := genericmodel.ErrResponse{
			Message: localization.GetMessage(lang, "common.503", nil),
		}
		c.JSON(http.StatusServiceUnavailable, apiRes)
		return
	}
	if rowsInserted > 0 {
		apiRes := albums.AlbumResponse{
			Message: localization.GetMessage(lang, "common.200", nil),
		}
		c.JSON(http.StatusOK, apiRes)
		return
	} else {
		apiRes := albums.AlbumResponse{
			Message: localization.GetMessage(lang, "common.404", nil),
		}
		c.JSON(http.StatusNotFound, apiRes)
		return
	}
}
