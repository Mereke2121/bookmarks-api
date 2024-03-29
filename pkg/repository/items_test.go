package repository

import (
	"github.com/bookmarks-api/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestItemsPostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewItemsRepository(db)

	type mockBehavior func(item models.Item)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		item         models.Item
		id           int
		wantErr      bool
	}{
		{
			name: "OK",
			item: models.Item{
				UserId: 1,
				Url:    "test url",
				Title:  "test title",
			},
			mockBehavior: func(item models.Item) {
				mock.ExpectBegin()

				rows := mock.NewRows([]string{"id"}).AddRow(item.Id)
				mock.ExpectQuery(`insert into bookmarks_items`).WithArgs(item.UserId, item.Url, item.Title).WillReturnRows(rows)

				mock.ExpectCommit()
			},
			id: 0,
		},
		{
			name: "Empty Fields",
			item: models.Item{
				UserId: 1,
				Title:  "test title",
			},
			mockBehavior: func(item models.Item) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id"}).AddRow(item.Id).RowError(1, errors.New("some error"))
				mock.ExpectQuery(`insert into bookmarks_items`).WithArgs(item.UserId, item.Url, item.Title).WillReturnRows(rows)

				mock.ExpectRollback()
			},
			id:      0,
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.item)

			got, err := r.AddItem(&testCase.item)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
			}
		})
	}
}
