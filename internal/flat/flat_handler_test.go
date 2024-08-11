package flat_test

import (
	"avito/internal/flat"
	"avito/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the repository
type MockFlatRepository struct {
	mock.Mock
}

func (m *MockFlatRepository) Create(flat models.Flat) (models.Flat, error) {
	args := m.Called(flat)
	return args.Get(0).(models.Flat), args.Error(1)
}

func (m *MockFlatRepository) GetFlatStatus(flatId int) (models.Status, error) {
	args := m.Called(flatId)
	return args.Get(0).(models.Status), args.Error(1)
}

func (m *MockFlatRepository) Update(flatId int, houseId int, status models.Status) (models.Flat, error) {
	args := m.Called(flatId, status)
	return args.Get(0).(models.Flat), args.Error(1)
}

func TestCreateFlat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockFlatRepository)
	handler := flat.NewFlatHandler(mockRepo)

	t.Run("Success", func(t *testing.T) {
		var price, rooms uint32 = 100, 1
		newFlat := models.Flat{
			HouseId: 1,
			Price:   &price,
			Rooms:   &rooms,
		}
		mockRepo.On("Create", newFlat).Return(newFlat, nil)

		flatJSON, _ := json.Marshal(newFlat)
		req, err := http.NewRequest(http.MethodPost, "/flat/create", bytes.NewBuffer(flatJSON))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		handler.Create(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, string(flatJSON), w.Body.String())
		mockRepo.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/flat/create", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req

		handler.Create(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUpdateFlat(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockFlatRepository)
	handler := flat.NewFlatHandler(mockRepo)

	t.Run("Success", func(t *testing.T) {
		updatedFlat := models.Flat{
			Id:      10,
			HouseId: 1,
			Status:  models.StatusCreated,
		}
		updateBody := models.FlatUpdateRequestBody{FlatId: 1, Status: models.StatusApproved}

		mockRepo.On("Update", uint32(1), updateBody).Return(updatedFlat, nil)

		bodyJSON, _ := json.Marshal(updateBody)
		req, err := http.NewRequest(http.MethodPut, "/flats/1", bytes.NewBuffer(bodyJSON))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Params = gin.Params{{Key: "id", Value: "1"}}

		handler.Update(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"ID":1,"Name":"Updated Flat"}`, w.Body.String())
		mockRepo.AssertExpectations(t)
	})

	t.Run("BadRequest", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/flats/1", bytes.NewBuffer([]byte("invalid json")))
		assert.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = req
		c.Params = gin.Params{{Key: "id", Value: "1"}}

		handler.Update(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}
