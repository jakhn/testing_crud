package storage_test

import (
	"context"
	"github.com/jakhn/film_crud/models"
	"testing"

	faker "github.com/bxcodec/faker/v3"
	"github.com/google/go-cmp/cmp"
)

func TestCategoryCreate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateCategory
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CreateCategory{
				Name: faker.Name(),
			},
			WantErr: false,
		},
		{
			Name: "case 2",
			Input: &models.CreateCategory{
				Name: faker.ChineseLastName(),
			},
			WantErr: true,
		},
		{
			Name: "case 3",
			Input: &models.CreateCategory{
				Name: faker.Name(),
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := categoryRepo.Create(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			if got == "" {
				t.Errorf("%s: got: %v", tc.Name, err)
				return
			}
		})
	}

}

func TestCategoryGetById(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CategoryPrimarKey
		Output  *models.Category
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CategoryPrimarKey{
				Id: "1",
			},
			Output: &models.Caregory{
				Id:   "1",
				Name: "Action",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := categoryRepo.GetByPKey(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Category) bool {
				return x.Name == y.Name
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestCategoryUpdate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UpdateCategory
		Output  *models.Category
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.UpdateCategory{
				Id:   "1",
				Name: "Actio",
			},
			Output: &models.Actor{
				Id:         "1",
				First_name: "Actio",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			_, err := categoryRepo.Update(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			res, err := categoryRepo.GetByPKey(
				context.Background(),
				&models.categoryPrimarKey{
					Id: tc.Input.Id,
				},
			)

			comparer := cmp.Comparer(func(x, y models.Category) bool {
				return x.Name == y.Name
			})

			if diff := cmp.Diff(tc.Output, res, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestCategoryDelete(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CategoryPrimarKey
		Want    string
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CategoryPrimarKey{
				Id: "1",
			},
			Want:    "no rows in result set",
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			err := categoryRepo.Delete(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			_, err = categoryRepo.GetByPKey(
				context.Background(),
				&models.categoryPrimarKey{
					Id: tc.Input.Id,
				},
			)

			comparer := cmp.Comparer(func(x, y string) bool {
				return x == y
			})

			if diff := cmp.Diff(tc.Want, err.Error(), comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}
