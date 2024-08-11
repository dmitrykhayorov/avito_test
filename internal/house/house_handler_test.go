package house_test

import (
	"avito/internal/house"
	"avito/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type MockHouseRepository struct {
	mock.Mock
}

func (m *MockHouseRepository) Create(flat models.House) (models.House, error) {
	args := m.Called(flat)
	return args.Get(0).(models.House), args.Error(1)
}

func (m *MockHouseRepository) GetFlatsByHouseID(userRole models.UserRole, houseId int) ([]models.Flat, error) {
	args := m.Called(userRole, houseId)
	return args.Get(0).([]models.Flat), args.Error(1)
}

func TestHouseHandler_GetFlatsByHouseID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockHouseRepository)
	handler := house.NewHouseHandler(mockRepo)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		userRole := models.Moderator
		houseID := 1
		flats := []models.Flat{
			{Id: 1, HouseId: houseID},
			{Id: 2, HouseId: houseID},
		}

		ctx.Request, _ = http.NewRequest("GET", "/house/1", nil)
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = gin.Params{{Key: "id", Value: strconv.Itoa(houseID)}}
		ctx.Set("userRole", string(userRole))

		mockRepo.On("GetFlatsByHouseID", userRole, houseID).Return(flats, nil)

		handler.GetFlatsByHouseID(ctx)

		mockRepo.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Invalid House ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		userRole := models.Moderator

		ctx.Request, _ = http.NewRequest("GET", "/house/abc", nil)
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Set("userRole", userRole)

		handler.GetFlatsByHouseID(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Repository Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		userRole := "admin"
		houseID := 1

		ctx.Request, _ = http.NewRequest("GET", "/house/1", nil)
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Params = gin.Params{{Key: "id", Value: strconv.Itoa(houseID)}}
		ctx.Set("userRole", userRole)

		mockRepo.On("GetFlatsByHouseID", models.UserRole(userRole), houseID).Return([]models.Flat{}, errors.New("repository error"))

		handler.GetFlatsByHouseID(ctx)

		mockRepo.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockHouseRepository)
	handler := house.NewHouseHandler(mockRepo)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		houseInput := models.House{Address: "ulitsa rayana goslinga 10", Year: 2015, Developer: "PIK"}
		houseOutput := models.House{Id: 1, Address: "ulitsa rayana goslinga 10", Year: 2015, Developer: "PIK"}
		jsonData, _ := json.Marshal(houseInput)

		ctx.Request, _ = http.NewRequest("POST", "/house/create", bytes.NewBuffer(jsonData))
		ctx.Request.Header.Set("Content-Type", "application/json")

		mockRepo.On("Create", houseInput).Return(houseOutput, nil)

		handler.Create(ctx)

		mockRepo.AssertExpectations(t)
		assert.Equal(t, http.StatusOK, w.Code)
		var response models.House
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, houseOutput, response)
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request, _ = http.NewRequest("POST", "/houses", bytes.NewBuffer([]byte("{invalid json")))
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Вызов хэндлера
		handler.Create(ctx)

		// Проверка
		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockRepo.AssertNotCalled(t, "Create")
	})

	t.Run("Validation Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		houseInput := models.House{Address: ""}
		jsonData, _ := json.Marshal(houseInput)

		ctx.Request, _ = http.NewRequest("POST", "/house/create", bytes.NewBuffer(jsonData))
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Mocking validateHouseData to simulate a validation error
		// Assuming validateHouseData returns an error for invalid data
		err := errors.New("empty address")

		// Вызов хэндлера
		handler.Create(ctx)

		// Проверка
		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response gin.H
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, err.Error(), response["error"])
	})

	t.Run("Repository Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		houseInput := models.House{Address: "Test House"}
		jsonData, _ := json.Marshal(houseInput)

		ctx.Request, _ = http.NewRequest("POST", "/houses", bytes.NewBuffer(jsonData))
		ctx.Request.Header.Set("Content-Type", "application/json")

		mockRepo.On("Create", houseInput).Return(models.House{}, errors.New("repository error"))

		handler.Create(ctx)

		mockRepo.AssertExpectations(t)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var response models.Response500
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, "repository error", response.Message)
	})
}
