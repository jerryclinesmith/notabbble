package notabbble_test

import (
	"bytes"
	"github.com/codegangsta/martini"
	"github.com/jerryclinesmith/notabbble"
	"github.com/jerryclinesmith/notabbble/db"
	"github.com/jerryclinesmith/notabbble/models"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	TestDB gorm.DB
	Server *martini.ClassicMartini
)

func TestNotabbble(t *testing.T) {
	RegisterFailHandler(Fail)

	Server = notabbble.InitServer()
	TestDB = db.Connect()

	RunSpecs(t, "Notabbble Suite")

	DeleteAllData()
}

func DeleteAllData() {
	TestDB.Where("id > 0").Delete(&models.Project{})
}

func Request(method string, url string, body []byte) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, url, bytes.NewReader(body))
	response := httptest.NewRecorder()
	Server.ServeHTTP(response, request)
	return response
}
