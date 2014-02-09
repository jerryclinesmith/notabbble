package main_test

import (
	"encoding/json"
	"fmt"
	"github.com/jerryclinesmith/notabbble/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
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
				if err := TestDB.Save(models.Project{Name: "My Project", CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error; err != nil {
					log.Fatal(err)
				}
				if err := TestDB.Save(models.Project{Name: "My Other Project", CreatedAt: time.Now(), UpdatedAt: time.Now()}).Error; err != nil {
					log.Fatal(err)
				}
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

	Describe("Project Get", func() {

		Context("when project doesn't exist", func() {

			It("should return 404", func() {
				response := Request("GET", "/api/projects/1", nil)
				Expect(response.Code).To(Equal(404))
			})

		})

		Context("when project does exist", func() {
			var project models.Project

			BeforeEach(func() {
				project = models.Project{Name: "My Project", CreatedAt: time.Now(), UpdatedAt: time.Now()}
				if err := TestDB.Save(&project).Error; err != nil {
					log.Fatal(err)
				}
				TestDB.First(&project, project.Id) // Need to reload - times are truncated by the db
			})

			It("Should return the project", func() {
				url := fmt.Sprint("/api/projects/", project.Id)
				response := Request("GET", url, nil)
				Expect(response.Code).To(Equal(200))
				expectedJSON, _ := json.Marshal(project)
				Expect(response.Body).To(MatchJSON(expectedJSON))
			})
		})
	})

})
