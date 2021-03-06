package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	models "v1/article"
	articleHttp "v1/article/delivery"
	mocks "v1/article/usecase/mocks"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)
	mockUCase := new(mocks.MockArticleUsecase)
	mockListArticle := make([]*models.Article, 0)
	mockListArticle = append(mockListArticle, &mockArticle)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(mockListArticle, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := articleHttp.HTTPArticleHandler{
		AUsecase: mockUCase,
	}
	handler.FetchArticle(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestFetchError(t *testing.T) {
	mockUCase := new(mocks.MockArticleUsecase)
	num := 1
	cursor := "2"
	mockUCase.On("Fetch", cursor, int64(num)).Return(nil, "", models.INTERNAL_SERVER_ERROR)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := articleHttp.HTTPArticleHandler{
		AUsecase: mockUCase,
	}
	handler.FetchArticle(c)

	responseCursor := rec.Header().Get("X-Cursor")
	assert.Equal(t, "", responseCursor)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockUCase := new(mocks.MockArticleUsecase)

	num := int(mockArticle.ID)

	mockUCase.On("GetByID", int64(num)).Return(&mockArticle, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/article/"+strconv.Itoa(int(num)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("article/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := articleHttp.HTTPArticleHandler{
		AUsecase: mockUCase,
	}
	handler.GetByID(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockArticle := models.Article{
		Title:     "Title",
		Content:   "Content",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tempMockArticle := mockArticle
	tempMockArticle.ID = 0
	mockUCase := new(mocks.MockArticleUsecase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUCase.On("Store", mock.AnythingOfType("*article.Article")).Return(&mockArticle, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/article", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/article")

	handler := articleHttp.HTTPArticleHandler{
		AUsecase: mockUCase,
	}
	handler.Store(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	var mockArticle models.Article
	err := faker.FakeData(&mockArticle)
	assert.NoError(t, err)

	mockUCase := new(mocks.MockArticleUsecase)

	num := int(mockArticle.ID)

	mockUCase.On("Delete", int64(num)).Return(true, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.DELETE, "/article/"+strconv.Itoa(int(num)), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("article/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := articleHttp.HTTPArticleHandler{
		AUsecase: mockUCase,
	}
	handler.Delete(c)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockUCase.AssertExpectations(t)

}
