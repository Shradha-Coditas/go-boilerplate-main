package delete

import (
	"boilerplate/models/albums"
	"boilerplate/util/genericmodel"
	"boilerplate/util/localization"
	"errors"

	"boilerplate/util/logwrapper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

// AlbumDeleteRequest - for Request Query param of Album Delete
type AlbumDeleteRequest struct {
	AlbumId int `form:"albumid" json:"albumid" example:"1" binding:"required"` // Id for album request to delete
}

// Handler - Handle the Delete requests at /v1/album
// @Summary Delete a album details
// @Description Delete a album by particular Id.
// @Tags Albums
// @Accept  json
// @Produce  json
// @Param authorization header string true "Authorization header"
// @Param body body AlbumDeleteRequest true "Params for delete Album by id."
// @Success 200 {object} albums.AlbumResponse
// @Failure 400 {object} genericmodel.ErrResponse "Invalid Payload."
// @Failure 401 {object} genericmodel.ErrResponse "Unauthorized."
// @Failure 404 {object} genericmodel.ErrResponse "Album Not Found."
// @Failure 500 {object} genericmodel.ErrResponse "Internal Server Error."
// @Router /album [DELETE]
func Handler(c *gin.Context) {
	lang, _ := c.Get("language")
	errResponse := genericmodel.ErrResponse{
		Message: localization.GetMessage(lang, "common.500", nil),
	}

	var queryData AlbumDeleteRequest

	logwrapper.Logger.Debugln(&queryData)
	if err := c.ShouldBindJSON(&queryData); err != nil {
		logwrapper.Logger.Errorln("error in RequestData to struct : ", err)
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

	deleteResponse, err := albums.DeletAlbum(int64(queryData.AlbumId))
	if err != nil {
		apiRes := genericmodel.ErrResponse{
			Message: localization.GetMessage(lang, "common.503", nil),
		}
		c.JSON(http.StatusServiceUnavailable, apiRes)
		return
	}
	if deleteResponse > 0 {
		// send success response
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
