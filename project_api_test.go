package notabbble_test

import (
	"encoding/json"
	"github.com/jerryclinesmith/notabbble/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Project", func() {

	BeforeEach(func() {
		DeleteAllData()
	})

	Describe("Project Index", func() {

		Context("Without existing projects", func() {

			It("returns empty array", func() {
				response := Request("GET", "/api/projects", nil)

				Expect(response.Code).To(Equal(200))
				expectedJSON, _ := json.Marshal([]models.Project{})
				Expect(response.Body).To(MatchJSON(expectedJSON))
			})

		})

		Context("With existing projects", func() {

			var projects []models.Project

			BeforeEach(func() {
				TestDB.Save(models.Project{Name: "My Project", CreatedAt: time.Now(), UpdatedAt: time.Now()})
				TestDB.Save(models.Project{Name: "My Other Project", CreatedAt: time.Now(), UpdatedAt: time.Now()})
				TestDB.Find(&projects)
			})

			It("returns projects", func() {
				response := Request("GET", "/api/projects", nil)

				Expect(response.Code).To(Equal(200))
				expectedJSON, _ := json.Marshal(projects)
				Expect(response.Body).To(MatchJSON(expectedJSON))
			})

		})

	})

})
