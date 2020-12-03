package users

import (
	"context"
	"dfee/api/ent"
	"dfee/api/ent/enttest"
	"dfee/api/users/dtos"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

// NewTestingEnt uses an in-memory sqlite
func NewTestingEnt(t *testing.T) (client *ent.Client) {
	client = enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	t.Cleanup(func() {
		client.Close()
	})
	return
}

// NewTestingLogger returns a sugared zap (test) logger
func NewTestingLogger(t *testing.T) *zap.SugaredLogger {
	return zaptest.NewLogger(t).Sugar()
}

func NewTestingServiceParams(t *testing.T) (p NewServiceParams) {
	p = NewServiceParams{
		Ent:    NewTestingEnt(t),
		Logger: NewTestingLogger(t),
	}
	return
}

type testingToolbelt struct {
	context context.Context
	ent     *ent.Client
	logger  *zap.SugaredLogger
	service *Service
}

func getTestingToolbelt(t *testing.T) *testingToolbelt {
	params := NewTestingServiceParams(t)

	return &testingToolbelt{
		context: context.Background(),
		ent:     params.Ent,
		logger:  params.Logger,
		service: NewService(params),
	}
}

func diffUser(left *ent.User, right *ent.User) (err error) {
	if diff := cmp.Diff(left, right, cmp.Comparer(AreEqual)); diff != "" {
		err = fmt.Errorf("Create() mismatch (-left +right):\n%s", diff)
	}
	return
}

func TestService(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		tb := getTestingToolbelt(t)
		var (
			err      error
			expected = &ent.User{
				ID:   1,
				Age:  87,
				Name: "Willie Nelson",
			}
			result   *ent.User
			resultDB *ent.User
		)

		if result, err = tb.service.Create(tb.context, dtos.AddUser{
			Age:  87,
			Name: "Willie Nelson",
		}); err != nil {
			t.Fatal("Received unexpected error", err)
		}

		if err = diffUser(result, expected); err != nil {
			t.Fatal(err)
		}

		if resultDB, err = tb.ent.User.Get(tb.context, result.ID); err != nil {
			t.Fatal("Could not read from db!", err)
		}
		if err = diffUser(result, resultDB); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetOne", func(t *testing.T) {
		tb := getTestingToolbelt(t)
		var (
			err      error
			expected = &ent.User{
				ID:   1,
				Age:  87,
				Name: "Willie Nelson",
			}
			result   *ent.User
			resultDB *ent.User
		)

		if resultDB, err = tb.ent.User.
			Create().
			SetName(expected.Name).
			SetAge(expected.Age).
			Save(tb.context); err != nil {
			t.Fatal("Failed to create user:", err)
		}
		if err = diffUser(resultDB, expected); err != nil {
			t.Fatal(err)
		}

		if result, err = tb.service.GetOne(tb.context, expected.ID); err != nil {
			t.Fatal("Failed to get user:", err)
		}
		if err = diffUser(result, expected); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		tb := getTestingToolbelt(t)
		var (
			err      error
			expected = []*ent.User{
				{
					ID:   1,
					Age:  87,
					Name: "Willie Nelson",
				},
				{
					ID:   2,
					Age:  52,
					Name: "James Hetfield",
				},
			}
			result   []*ent.User
			resultDB *ent.User
		)

		for i, v := range expected {
			if resultDB, err = tb.ent.User.
				Create().
				SetName(v.Name).
				SetAge(v.Age).
				Save(tb.context); err != nil {
				t.Fatal("Failed to create user:", err)
			}
			if err = diffUser(resultDB, expected[i]); err != nil {
				t.Fatal(err)
			}
		}

		if result, err = tb.service.GetAll(tb.context); err != nil {
			t.Fatal("Failed to get users:", err)
		}
		for i, v := range result {
			if err = diffUser(v, expected[i]); err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Update", func(t *testing.T) {
		tb := getTestingToolbelt(t)
		var (
			err      error
			expected = &ent.User{
				ID:   1,
				Age:  87,
				Name: "Willie Nelson",
			}
			expectedUpdated = &ent.User{
				ID:   expected.ID,
				Age:  52,
				Name: "James Hetfield",
			}
			result   *ent.User
			resultDB *ent.User
		)

		if resultDB, err = tb.ent.User.
			Create().
			SetName(expected.Name).
			SetAge(expected.Age).
			Save(tb.context); err != nil {
			t.Fatal("Failed to create user:", err)
		}
		if err = diffUser(resultDB, expected); err != nil {
			t.Fatal(err)
		}

		dto := dtos.UpdateUser{
			Age:  expectedUpdated.Age,
			Name: expectedUpdated.Name,
		}
		if result, err = tb.service.Update(
			tb.context,
			expectedUpdated.ID,
			dto,
		); err != nil {
			t.Fatal("Failed to update user:", err)
		}
		if err = diffUser(result, expectedUpdated); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		tb := getTestingToolbelt(t)
		var (
			err      error
			expected = &ent.User{
				ID:   1,
				Age:  87,
				Name: "Willie Nelson",
			}
			resultDB        *ent.User
			resultDBDeleted *ent.User
		)

		if resultDB, err = tb.ent.User.
			Create().
			SetName(expected.Name).
			SetAge(expected.Age).
			Save(tb.context); err != nil {
			t.Fatal("Failed to create user:", err)
		}
		if err = diffUser(resultDB, expected); err != nil {
			t.Fatal(err)
		}

		if err = tb.service.Delete(
			tb.context,
			expected.ID,
		); err != nil {
			t.Fatal("Failed to delete user:", err)
		}

		if resultDBDeleted, err = tb.ent.User.Get(
			tb.context,
			resultDB.ID,
		); err == nil {
			t.Fatal("User still exists!", resultDBDeleted)
		}
	})
}
