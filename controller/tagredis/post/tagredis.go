package post

import (
	"boilerplate/models/tagredis"
	"boilerplate/util/genericmodel"
	"boilerplate/util/localization"
	"errors"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Handler - Handle the Post requests at /v1/tags
// @Summary store data of Last Ticker request
// @Description Api for Last Ticker data.
// @Tags Tags
// @Accept  json
// @Produce  json
// @Param authorization header string true "Authorization header"
// @Param body body tags.LastTick true "Payload for Album"
// @Success 201 {object} tags.LastTickResponse
// @Failure 400 {object} genericmodel.ErrResponse "Invalid Payload."
// @Failure 401 {object} genericmodel.ErrResponse "Unauthorized."
// @Failure 406 {object} genericmodel.ErrResponse "Token Expired."
// @Failure 500 {object} genericmodel.ErrResponse "Internal Server Error."
// @Failure 409 {object} genericmodel.ErrResponse "Conflict."
// @Failure 503 {object} genericmodel.ErrResponse "service Unavailable."
// @Router /tags [POST]
func Handler(c *gin.Context) {
	//Language in header
	lang, _ := c.Get("language")

	//Common Error Response
	errResponse := genericmodel.ErrResponse{
		Message: localization.GetMessage(lang, "common.500", nil),
	}

	//Reqiest payload variable
	var RequstPayload tagredis.TagSegment
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
	//RequstPayload.TickDate = int64(time.Now().Unix())
	err := tagredis.AddTagSegment(RequstPayload)
	if err != nil {
		apiRes := genericmodel.ErrResponse{
			Message: localization.GetMessage(lang, "common.503", nil),
		}
		c.JSON(http.StatusServiceUnavailable, apiRes)
		return
	}
	apiRes := tagredis.TagSegementResponse{
		Message: localization.GetMessage(lang, "common.201", nil),
	}
	c.JSON(http.StatusOK, apiRes)
	return
}
