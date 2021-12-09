package get

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

// AlbumGetRequest - for Request data of Album details
type AlbumGetRequest struct {
	Offset int `form:"offset" json:"offset" example:"0"`         // Skip records for pagination start from 0
	Limit  int `form:"limit" json:"limit" xml:"to" example:"10"` // limit records for pagination Max 20
}

// AlbumGetResponse - response message data Album
type AlbumGetResponse struct {
	Message    string                      `json:"message" example:"Albums data found successFully"`
	Data       []albums.AlbumsResponseData `json:"Albums"`
	TotalCount int64                       `json:"totalAlbums" example:"10"`
}

// Handler - Handle the Get requests at /v1/album
// @Summary To fetch Album list with pagination
// @Description Get all Album list with pagination.
// @Tags Albums
// @Accept  json
// @Produce  json
// @Param authorization header string true "Authorization header"
// @Param filter query AlbumGetRequest true "Params for get Albums."
// @Success 200 {object} AlbumGetResponse
// @Failure 400 {object} genericmodel.ErrResponse "Invalid Request."
// @Failure 401 {object} genericmodel.ErrResponse "Unauthorized."
// @Failure 404 {object} genericmodel.ErrResponse "Albums Not Found."
// @Failure 204
// @Failure 500 {object} genericmodel.ErrResponse "Internal Server Error"
// @Router /album [get]
func Handler(c *gin.Context) {
	lang, _ := c.Get("language")
	errResponse := genericmodel.ErrResponse{
		Message: localization.GetMessage(lang, "common.500", nil),
	}

	//Request Validator
	var queryData AlbumGetRequest

	if err := c.ShouldBind(&queryData); err != nil {
		logwrapper.Logger.Errorln("error in queryData to struct : ", err)
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

	if queryData.Offset < 0 {
		errResponse.Message = localization.GetMessage(lang, "common.400", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}

	if queryData.Limit == 0 {
		queryData.Limit = 20
	}

	records, err := albums.GetAlbum(int64(queryData.Offset), int64(queryData.Limit))
	if err != nil {
		logwrapper.Logger.Errorln("error get : ", err)
		errResponse.Message = localization.GetMessage(lang, "common.500", nil)
		c.AbortWithStatusJSON(http.StatusInternalServerError, errResponse)
		return
	}
	if records == nil {
		errResponse.Message = localization.GetMessage(lang, "common.404", nil)
		c.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	if len(records) <= 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, errResponse)
		return
	}
	AlbumResponse := AlbumGetResponse{
		Message:    localization.GetMessage(lang, "common.200", nil),
		Data:       records,
		TotalCount: 10,
	}
	c.JSON(http.StatusOK, AlbumResponse)

}
