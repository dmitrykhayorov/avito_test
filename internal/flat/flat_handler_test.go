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
	dbHost      string
	dbPort      string
	dbUser      string
	dbPassword  string
	dbName      string
	servicePort string
)

func loadEnvVariables() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
	servicePort = os.Getenv("SERVICE_PORT")
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

func TestFlatHandler_Create(t *testing.T) {
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
	assert.Equal(t, 200, w.Code)

	req.Header.Set("Authorization", tokenClient)
	server.Router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
