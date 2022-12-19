package storage_test

import (
	"context"
	"github.com/jakhn/film_crud/models"
	"testing"

	faker "github.com/bxcodec/faker/v3"
	"github.com/google/go-cmp/cmp"
)

func TestFilmCreate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateFilm
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CreateFilm{
				Title:       faker.Name(),
				Description: faker.Paragraph(),
				ReleaseYear: faker.Date(),
				Duration:    int32(faker.UnixTime()),
			},
			WantErr: false,
		},
		{
			Name: "case 2",
			Input: &models.CreateFilm{
				Title:       faker.Name(),
				Description: faker.Paragraph(),
				ReleaseYear: faker.ChineseLastName(),
				Duration:    int32(faker.UnixTime()),
			},
			WantErr: true,
		},
		{
			Name: "case 3",
			Input: &models.CreateFilm{
				Title:       faker.Name(),
				Description: faker.Paragraph(),
				ReleaseYear: faker.Date(),
				Duration:    int32(faker.UnixTime()),
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := filmRepo.Create(
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

func TestFilmGetById(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.FilmPrimarKey
		Output  *models.Film
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.FilmPrimarKey{
				Id: "fc73c0b0-91de-48e3-be78-4485dd829704",
			},
			Output: &models.Film{
				Id:          "fc73c0b0-91de-48e3-be78-4485dd829704",
				Title:       "King Garrett Casper",
				Description: "Voluptatem sit accusantium aut consequatur perferendis. Accusantium sit voluptatem consequatur aut perferendis. Aut perferendis accusantium voluptatem consequatur sit. Accusantium aut sit voluptatem perferendis consequatur. Consequatur sit perferendis accusantium voluptatem aut. Voluptatem perferendis aut consequatur accusantium sit. Accusantium voluptatem aut sit perferendis consequatur. Sit voluptatem aut consequatur perferendis accusantium. Perferendis sit accusantium consequatur aut voluptatem.",
				ReleaseYear: "2000-05-11",
				Duration:    152420074,
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := filmRepo.GetByPKey(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Film) bool {
				return x.Title == y.Title && x.Description == y.Description
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestFilmUpdate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UpdateFilm
		Output  *models.Film
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.UpdateFilm{
				Id:          "fc73c0b0-91de-48e3-be78-4485dd829704",
				Title:       "King Garrett Caspe",
				ReleaseYear: "2022-10-10",
			},
			Output: &models.Film{
				Id:          "fc73c0b0-91de-48e3-be78-4485dd829704",
				Title:       "King Garrett Caspe",
				ReleaseYear: "2022-10-10",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			_, err := filmRepo.Update(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			res, err := filmRepo.GetByPKey(
				context.Background(),
				&models.FilmPrimarKey{
					Id: tc.Input.Id,
				},
			)

			comparer := cmp.Comparer(func(x, y models.Film) bool {
				return x.Title == y.Title
			})

			if diff := cmp.Diff(tc.Output, res, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestFilmDelete(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.FilmPrimarKey
		Want    string
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.FilmPrimarKey{
				Id: "98561f78-b5ce-47e0-906f-c72828fcfef1",
			},
			Want:    "no rows in result set",
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			err := filmRepo.Delete(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			_, err = filmRepo.GetByPKey(
				context.Background(),
				&models.FilmPrimarKey{
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
