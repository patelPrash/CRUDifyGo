package store

import (
	"context"
	"database/sql"
	"io"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	gofrLog "gofr.dev/pkg/log"

	"sample/model"
)

func newMock(t *testing.T) (*gofr.Context, sqlmock.Sqlmock) {
	mockLogger := gofrLog.NewMockLogger(io.Discard)

	db, mock, errMock := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if errMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errMock)
	}

	ctx := gofr.NewContext(nil, nil, &gofr.Gofr{DataStore: datastore.DataStore{ORM: db}, Logger: mockLogger})

	ctx.Context = context.Background()

	return ctx, mock
}

func Test_Create(t *testing.T) {
	ctx, mock := newMock(t)
	product := &model.Product{ID: 1, Name: "SAMPLE-NAME", MinAmount: 100, MaxAmount: 400, PurchaseAmount: 200, Enable: 1}

	testCases := []struct {
		desc        string
		dbMock      []interface{}
		input       *model.Product
		expectedRes *model.Product
		expectedErr error
	}{
		{desc: "failure case", dbMock: []interface{}{
			mock.ExpectExec(createQuery).
				WillReturnError(errors.Error("error from db"))}, input: product, expectedErr: errors.DB{Err: errors.Error("error from db")},
		},
		{desc: "success case", input: &model.Product{ID: 1, Name: "SAMPLE-NAME", MinAmount: 100, MaxAmount: 400, PurchaseAmount: 200, Enable: 1},
			dbMock: []interface{}{
				mock.ExpectExec(createQuery).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)},
			expectedRes: product, expectedErr: nil},
	}

	s := New()
	for i, tc := range testCases {
		res, err := s.Create(ctx, tc.input)

		assert.Equal(t, tc.expectedRes, res, "Test[%d] Failed,Expected : %v\nGot : %v\n", i, tc.expectedErr, err)
		assert.Equal(t, tc.expectedErr, err, "Test[%d] Failed,Expected : %v\nGot : %v\n", i, tc.expectedErr, err)
	}
}

func Test_GetByID(t *testing.T) {
	ctx, mock := newMock(t)

	testCases := []struct {
		desc        string
		input       int
		dbMock      []interface{}
		expOutput   *model.Product
		expectedErr error
	}{
		{desc: "success case", input: 1,
			dbMock: []interface{}{mock.ExpectQuery(getByIDQuery).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "minAmount", "maxAmount", "purchaseAmount", "enable"}).AddRow(1,
					"SAMPLE-NAME", 100, 200, 400, 1))},
			expOutput:   &model.Product{ID: 1, Name: "SAMPLE-NAME", MinAmount: 100, MaxAmount: 400, PurchaseAmount: 200, Enable: 1},
			expectedErr: nil},
		{desc: "failure case", input: 1,
			dbMock: []interface{}{mock.ExpectQuery(getByIDQuery).
				WillReturnError(errors.Error("error from db"))},
			expOutput:   nil,
			expectedErr: errors.Error("error from db")},
		{desc: "failure case", input: 1,
			dbMock: []interface{}{mock.ExpectQuery(getByIDQuery).
				WillReturnError(sql.ErrNoRows)},
			expectedErr: errors.EntityNotFound{Entity: "product", ID: "1"}},
	}

	s := New()

	for i, tc := range testCases {
		output, errRet := s.GetByID(ctx, tc.input)

		assert.Equal(t, tc.expOutput, output, "Test[%d] Failed ,Expected : %v\nGot : %v\n", i, tc.expOutput, output)
		assert.Equal(t, tc.expectedErr, errRet, "Test[%d] Failed , Expected : %v\nGot : %v\n", i, tc.expectedErr, errRet)
	}
}

func Test_Update(t *testing.T) {
	ctx, mock := newMock(t)

	testCases := []struct {
		desc        string
		input       *model.Product
		dbMock      []interface{}
		expOutput   *model.Product
		expectedErr error
	}{
		{desc: "failure case", input: &model.Product{ID: 1, Name: "sample-name"},
			dbMock: []interface{}{mock.ExpectExec(updateQuery).
				WillReturnError(errors.Error("error from db"))},
			expOutput:   nil,
			expectedErr: errors.DB{Err: errors.Error("error from db")}},
		{desc: "success case", input: &model.Product{ID: 1, Name: "sample-name"},
			dbMock: []interface{}{mock.ExpectExec(updateQuery).
				WillReturnResult(sqlmock.NewResult(1, 1)),
				mock.ExpectQuery(updateQuery).
					WillReturnRows(sqlmock.NewRows([]string{"id", "domain"}).AddRow("f0ddfd00-8554-4ccc-b6cf-eb7577b5dbbb",
						"auth.zopsmart.com"))},
			expOutput:   &model.Product{ID: 1, Name: "sample-name"},
			expectedErr: nil},
	}

	s := New()

	for i, tc := range testCases {
		_, errUpdate := s.Update(ctx, tc.input)

		assert.Equal(t, tc.expectedErr, errUpdate, "Test[%d] Failed,Expected : %v\nGot : %v\n", i, tc.expectedErr, errUpdate)
	}
}

func Test_Delete(t *testing.T) {
	ctx, mock := newMock(t)

	testCases := []struct {
		desc        string
		input       *model.Product
		dbMock      []interface{}
		expOutput   *model.Product
		expectedErr error
	}{
		{desc: "failure case", input: &model.Product{ID: 1, Name: "sample-name"},
			dbMock: []interface{}{mock.ExpectExec(deleteQuery).
				WillReturnError(errors.Error("error from db"))},
			expOutput:   nil,
			expectedErr: errors.DB{Err: errors.Error("error from db")}},
		{desc: "success case", input: &model.Product{ID: 1, Name: "sample-name"},
			dbMock: []interface{}{mock.ExpectExec(deleteQuery).
				WillReturnResult(sqlmock.NewResult(1, 1)),
				mock.ExpectExec(deleteQuery)},
			expOutput:   &model.Product{ID: 1, Name: "sample-name"},
			expectedErr: nil},
	}

	s := New()

	for i, tc := range testCases {
		errUpdate := s.Delete(ctx, tc.input.ID)

		assert.Equal(t, tc.expectedErr, errUpdate, "Test[%d] Failed,Expected : %v\nGot : %v\n", i, tc.expectedErr, errUpdate)
	}
}
