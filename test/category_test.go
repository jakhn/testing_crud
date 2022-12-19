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

func TestCategory(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createCategory(t)
			deleteCategory(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func createCategory(t *testing.T) string {
	response := &models.Caregory{}

	request := &models.CreateCategory{
		First_name: faker.Name(),
		Last_name:  faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPost, "/category", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func updateCategory(t *testing.T, id string) string {
	response := &models.Caregory{}
	request := &models.UpdateCategory{
		First_name: faker.Name(),
		Last_name:  faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPut, "/category/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.Id
}

func deleteCategory(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/category/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
