package tests

import (
	"net/http"
	"net/http/httptest"
	"staff-contract-management/controllers"
	"staff-contract-management/models"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupStaffRouter() *gin.Engine {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Staff{}, &models.Contract{})
	r := gin.Default()
	controllers.RegisterRoutes(r, db)
	return r
}

func TestCreateStaff(t *testing.T) {
	r := setupStaffRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/staffs/", strings.NewReader(`{"Username":"testuser", "Password":"password", "Role":"staff"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestGetStaffs(t *testing.T) {
	r := setupStaffRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/staffs/", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetStaffDetail(t *testing.T) {
	r := setupStaffRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/staffs/1", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestUpdateStaff(t *testing.T) {
	r := setupStaffRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/staffs/1", strings.NewReader(`{"Username":"updateduser", "Password":"newpassword", "Role":"staff"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestDeleteStaff(t *testing.T) {
	r := setupStaffRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/staffs/1", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}
}
