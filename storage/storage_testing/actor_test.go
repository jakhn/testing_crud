package storage_test

import (
	"context"
	"github.com/jakhn/film_crud/models"
	"testing"

	faker "github.com/bxcodec/faker/v3"
	"github.com/google/go-cmp/cmp"
)

func TestActorCreate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateActor
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.CreateActor{
				First_name: faker.Name(),
				Last_name:  faker.Name(),
			},
			WantErr: false,
		},
		{
			Name: "case 2",
			Input: &models.CreateActor{
				First_name: faker.Name(),
				Last_name:  faker.ChineseLastName(),
			},
			WantErr: true,
		},
		{
			Name: "case 3",
			Input: &models.CreateActor{
				First_name: faker.Name(),
				Last_name:  faker.Name(),
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := actorRepo.Create(
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

func TestActorGetById(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.ActorPrimarKey
		Output  *models.Actor
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.ActorPrimarKey{
				Id: "1",
			},
			Output: &models.Actor{
				Id:         "1",
				First_name: "Penelope",
				Last_name:  "Guiness",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			got, err := actorRepo.GetByPKey(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			comparer := cmp.Comparer(func(x, y models.Actor) bool {
				return x.First_name == y.First_name && x.Last_name == y.Last_name
			})

			if diff := cmp.Diff(tc.Output, got, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestActorUpdate(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.UpdateActor
		Output  *models.Actor
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.UpdateActor{
				Id:         "1",
				First_name: "Penelope",
				Last_name:  "Guine",
			},
			Output: &models.Actor{
				Id:         "1",
				First_name: "Penelope",
				Last_name:  "Guine",
			},
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			_, err := actorRepo.Update(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			res, err := actorRepo.GetByPKey(
				context.Background(),
				&models.actorPrimarKey{
					Id: tc.Input.Id,
				},
			)

			comparer := cmp.Comparer(func(x, y models.Actor) bool {
				return x.First_name == y.First_name
			})

			if diff := cmp.Diff(tc.Output, res, comparer); diff != "" {
				t.Errorf(diff)
				return
			}
		})
	}

}

func TestActorDelete(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.ActorPrimarKey
		Want    string
		WantErr bool
	}{
		{
			Name: "case 1",
			Input: &models.FilmPrimarKey{
				Id: "1",
			},
			Want:    "no rows in result set",
			WantErr: false,
		},
	}

	for _, tc := range tests {

		t.Run(tc.Name, func(t *testing.T) {

			err := actorRepo.Delete(
				context.Background(),
				tc.Input,
			)

			if err != nil {
				t.Errorf("%s: expected: %v, got: %v", tc.Name, tc.WantErr, err)
				return
			}

			_, err = actorRepo.GetByPKey(
				context.Background(),
				&models.actorPrimarKey{
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
