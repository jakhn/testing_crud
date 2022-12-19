package order_service

import (
	"fmt"
	"github.com/jakhn/testing_crud/models"
	"net/http"
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

var s int64

func TestFilm(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createFilm(t)
			deleteFilm(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createFilm(t *testing.T) string {
	response := &models.Film{}

	request := &models.CreateFilm{
		Title:       faker.Name(),
		Description: faker.Paragraph(),
		ReleaseYear: faker.Date(),
		Duration:    int32(faker.UnixTime()),
	}

	resp, err := PerformRequest(http.MethodPost, "/film", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func updateFilm(t *testing.T, id string) string {
	response := &models.Film{}
	request := &models.UpdateFilm{
		Title:       faker.Name(),
		Description: faker.Paragraph(),
		ReleaseYear: faker.Date(),
		Duration:    int32(faker.UnixTime()),
	}

	resp, err := PerformRequest(http.MethodPut, "/film/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.Id
}

func deleteFilm(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/film/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
