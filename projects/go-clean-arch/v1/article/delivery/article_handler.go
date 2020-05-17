package http

import (
	"net/http"
	"strconv"

	models "v1/article"
	articleUcase "v1/article/usecase"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// HTTPArticleHandler ...
type HTTPArticleHandler struct {
	AUsecase articleUcase.ArticleUsecase
}

// FetchArticle ...
func (a *HTTPArticleHandler) FetchArticle(c echo.Context) error {
	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := a.AUsecase.Fetch(cursor, int64(num))

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

// GetByID ...
func (a *HTTPArticleHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	art, err := a.AUsecase.GetByID(id)

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *models.Article) (bool, error) {
	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store ...
func (a *HTTPArticleHandler) Store(c echo.Context) error {
	var article models.Article
	err := c.Bind(&article)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&article); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ar, err := a.AUsecase.Store(&article)

	if err != nil {
		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.JSON(http.StatusCreated, ar)
}

// Delete ...
func (a *HTTPArticleHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	_, err = a.AUsecase.Delete(id)

	if err != nil {

		return c.JSON(getStatusCode(err), err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:

		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// NewArticleHTTPHandler ...
func NewArticleHTTPHandler(e *echo.Echo, us articleUcase.ArticleUsecase) {
	handler := &HTTPArticleHandler{
		AUsecase: us,
	}

	e.GET("/article", handler.FetchArticle)
	e.POST("/article", handler.Store)
	e.GET("/article/:id", handler.GetByID)
	e.DELETE("/article/:id", handler.Delete)
}
