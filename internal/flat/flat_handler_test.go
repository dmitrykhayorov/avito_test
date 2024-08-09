package flat_test

import (
	"avito/internal/api"
	"avito/internal/auth"
	"avito/internal/flat"
	"avito/internal/house"
	"avito/internal/models"
	"avito/internal/repository"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var (
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
)

func loadEnvVariables() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
}

func prepareServer() *api.Server {
	loadEnvVariables()
	gin.SetMode(gin.TestMode)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	retries := 5
	var db *sql.DB
	var err error

	for i := 0; i < retries; i++ {
		db, err = sql.Open("postgres", connStr)

		if err == nil {
			break
		}

		time.Sleep(time.Second * 5)
	}
	if err != nil {
		log.Fatalln("cannot connect db: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("cannot connect db: ", err)
	}

	houseRepository := repository.NewHouseRepository(db)
	flatRepository := repository.NewFlatRepository(db)

	authHandler := auth.NewAuthHandler()
	flatHandler := flat.NewFlatHandler(flatRepository)
	houseHandler := house.NewHouseHandler(houseRepository)
	server := api.NewServer(authHandler, flatHandler, houseHandler)
	return server
}

var server = prepareServer()
var tokenClient = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyUm9sZSI6ImNsaWVudCIsImV4cCI6MTcyMzI3Njk3N30.wMeyv_z6OS-sJ9ljWPj6p4fY6J3wb4jNmqtMzSS6X6o"
var tokenModerator = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyUm9sZSI6Im1vZGVyYXRvciIsImV4cCI6MTcyMzI3Njc5MX0.nzY7fPqrOqvJaIXUtEN7BCCS4so_YPZ2Dk3srwW1lU0"
var url = "http://localhost:8080/flat/create"

func TestFlatHandler_CreateValidInput(t *testing.T) {
	flatToCreateValid := models.FlatCreateRequestBody{
		HouseId: 1,
		Price:   100,
		Rooms:   1,
	}
	body, _ := json.Marshal(flatToCreateValid)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/flat/create", bytes.NewBuffer(body))
	req.Header.Set("Authorization", tokenModerator)

	server.Router.ServeHTTP(w, req)
	//d, _ := io.ReadAll(w.Body)
	assert.Equal(t, 200, w.Code)

	server.Router.ServeHTTP(w, req)
	//d, _ = io.ReadAll(w.Body)
	assert.Equal(t, 200, w.Code)
}

func TestFlatHandler_CreateInvalidInput(t *testing.T) {
	flatInvalid := []models.FlatCreateRequestBody{
		{
			HouseId: 1_000_000_000,
			Price:   100,
			Rooms:   4,
		},
		{
			HouseId: 1,
			Rooms:   4,
		},
		{
			HouseId: 1,
			Price:   100,
			Rooms:   0,
		},
		{
			HouseId: 0,
			Price:   10101,
			Rooms:   5,
		},
	}

	w := httptest.NewRecorder()
	for _, flat := range flatInvalid {
		t.Run("", func(t *testing.T) {
			body, _ := json.Marshal(flat)
			req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			server.Router.ServeHTTP(w, req)
			assert.NotEqual(t, 200, w.Code)
		})
	}
}

func TestFlatHandler_CreateUnauthorized(t *testing.T) {
	flatToCreateValid := models.FlatCreateRequestBody{
		HouseId: 1,
		Price:   100,
		Rooms:   1,
	}

	body, _ := json.Marshal(flatToCreateValid)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/flat/create", bytes.NewBuffer(body))

	server.Router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestFlatHandler_UpdateUnauthorized(t *testing.T) {
	body := models.FlatUpdateRequestBody{
		FlatId: 4,
		Status: "approved",
	}

	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/flat/create", bytes.NewBuffer(jsonBody))
	server.Router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	req.Header.Set("Authorization", tokenClient)
	assert.Equal(t, 401, w.Code)
}

func TestFlatHandler_UpdateValid(t *testing.T) {
	body := models.FlatUpdateRequestBody{
		FlatId: 4,
		Status: "approved",
	}

	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/flat/update", bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", tokenModerator)
	server.Router.ServeHTTP(w, req)

	//d, _ := io.ReadAll(w.Body)
	//fmt.Println(string(d))
	assert.Equal(t, 200, w.Code)
}

func TestFlatHandler_UpdateInvalidRequest(t *testing.T) {
	body := models.FlatUpdateRequestBody{
		FlatId: 4,
		Status: "pending",
	}

	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/flat/update", bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", tokenModerator)
	server.Router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}
