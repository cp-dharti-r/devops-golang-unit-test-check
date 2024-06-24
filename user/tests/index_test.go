package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"test"
	"testing"
	"user"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type Repository struct {
	Db *sqlx.DB
}

var testDb *sqlx.DB
var err error
var userRepo *user.Repository

func TestInit(t *testing.T) {
	testDb = test.TestDB() // connection of your test database
	fmt.Println(testDb)

	userRepo = user.New(testDb)
}

func TestCreateUserBadRequest(t *testing.T) {

	// required user table operations
	DropUsersTable(testDb)
	CreateUsersTable(testDb)

	// create an API route
	router := gin.Default()
	router.POST("/api/users", userRepo.Create)

	engine := gin.New()

	// send request
	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer([]byte(`{"email": 123}`)))
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	// match an API response
	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}

func TestCreateUserSuccess(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)

	router := gin.Default()
	router.POST("/api/users", userRepo.Create)

	engine := gin.New()

	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer([]byte(`{"name":"John Doe","email":"john@example.com"}`)))
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
}

func TestGetUserBadRequest(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)

	router := gin.Default()
	router.GET("/api/users/:id", userRepo.Get)

	engine := gin.New()

	req, err := http.NewRequest("GET", "/api/users/", nil)
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}

func TestGetUserNotFound(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)

	router := gin.Default()
	router.GET("/api/users/:id", userRepo.Get)

	engine := gin.New()

	req, err := http.NewRequest("GET", "/api/users/1", nil)
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)
}

func TestGetUserSuccess(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)
	InsertIntoUsersTable(testDb)

	router := gin.Default()
	router.GET("/api/users/:id", userRepo.Get)

	engine := gin.New()

	req, err := http.NewRequest("GET", "/api/users/1", nil)
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	got := GotData(w, t)
	expected := `{"id":1, "name":"John Doe"}`

	assert.Equal(t, expected, got)
}

func TestUpdateUserBadRequest(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)

	router := gin.Default()
	router.PUT("/api/users/:id", userRepo.Update)

	engine := gin.New()

	req, err := http.NewRequest("PUT", "/api/users/1", bytes.NewBuffer([]byte(`{"email": 123}`)))
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)

	got := GotData(w, t)
	assert.Empty(t, got)
}

func TestUpdateUserNotFound(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)
	InsertIntoUsersTable(testDb)

	router := gin.Default()
	router.PUT("/api/users/:id", userRepo.Update)

	engine := gin.New()

	req, err := http.NewRequest("PUT", "/api/users/5", bytes.NewBuffer([]byte(`{"name":"John Doe","email":"john@example.com"}`)))
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)

	got := GotData(w, t)
	assert.Empty(t, got)
}

func TestUpdateUserSuccess(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)
	InsertIntoUsersTable(testDb)

	router := gin.Default()
	router.PUT("/api/users/:id", userRepo.Update)

	engine := gin.New()

	req, err := http.NewRequest("PUT", "/api/users/1", bytes.NewBuffer([]byte(`{"name":"John Doe","email":"john@example.com"}`)))
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	got := GotData(w, t)
	assert.Empty(t, got)
}

func TestDeleteUserNotFound(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)

	router := gin.Default()
	router.DELETE("/api/users/:id", userRepo.Delete)

	engine := gin.New()

	req, err := http.NewRequest("DELETE", "/api/users/5", nil)
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)
}

func TestDeleteUserSuccess(t *testing.T) {

	DropUsersTable(testDb)
	CreateUsersTable(testDb)
	InsertIntoUsersTable(testDb)

	router := gin.Default()
	router.DELETE("/api/users/:id", userRepo.Delete)

	engine := gin.New()

	req, err := http.NewRequest("DELETE", "/api/users/1", nil)
	if err != nil {
		t.Errorf("Error in creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	engine.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)
}

// create users table
func CreateUsersTable(Db *sqlx.DB) {
	Db.MustExec(`CREATE TABLE IF NOT EXISTS users 
		(id int(11) NOT NULL AUTO_INCREMENT,
		name varchar(195) default null,
		email varchar(195) default null
		primary key (id));`)
}

// insert user
func InsertIntoUsersTable(Db *sqlx.DB) {
	Db.MustExec("INSERT INTO users(name, email) VALUES('John Doe', 'john@example.com');")
}

// drop users table
func DropUsersTable(Db *sqlx.DB) {
	Db.MustExec(`DROP TABLE IF EXISTS users`)
}

// make json data
func GotData(w *httptest.ResponseRecorder, t *testing.T) map[string]interface{} {
	var got map[string]interface{}
	if len(w.Body.Bytes()) != 0 {
		err := json.Unmarshal(w.Body.Bytes(), &got)
		if err != nil {
			t.Fatal(err)
		}
	}
	return got
}
