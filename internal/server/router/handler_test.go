package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"zadanie/internal/app"
	"zadanie/internal/domain"
	"zadanie/internal/storage/storage_mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestHandler_reserveGoods(t *testing.T) {

	type testCase struct {
		in           []byte
		resCode      int
		resReserveID string
	}

	testSet := []testCase{
		{
			in:           []byte("[1,2]"),
			resCode:      200,
			resReserveID: "test_reserve_id_1",
		},
		{
			in:           []byte("[0]"),
			resCode:      500,
			resReserveID: "",
		},
		{
			in:           []byte("{[1,2]}"),
			resCode:      400,
			resReserveID: "",
		},
	}

	s := storage_mock.NewStorageMock()
	a := app.NewApp(s, zap.NewNop())
	h := NewHandler(a)
	r := NewRouter(h, "prod")

	for _, testcase := range testSet {
		w := httptest.NewRecorder()
		bodyreader := bytes.NewReader(testcase.in)
		req, _ := http.NewRequest("POST", "/goods/reserve", bodyreader)
		r.ServeHTTP(w, req)

		assert.Equal(t, testcase.resCode, w.Code)

		var responseBody []byte
		responseBody, err := io.ReadAll(w.Body)
		if err != nil {
			t.Errorf("cant read response body")
		}
		reserve := struct {
			ReserveID string `json:"reserveID"`
		}{ReserveID: ""}

		json.Unmarshal(responseBody, &reserve)

		assert.Equal(t, testcase.resReserveID, reserve.ReserveID)
	}

}

func TestHandler_releaseGoods(t *testing.T) {
	type testCase struct {
		in      []byte
		resCode int
	}

	testSet := []testCase{
		{
			in:      []byte("[1,2]"),
			resCode: 200,
		},
		{
			in:      []byte("[0]"),
			resCode: 500,
		},
		{
			in:      []byte("[]"),
			resCode: 400,
		},
		{
			in:      []byte("{[1,2]}"),
			resCode: 400,
		},
	}

	s := storage_mock.NewStorageMock()
	a := app.NewApp(s, zap.NewNop())
	h := NewHandler(a)
	r := NewRouter(h, "prod")

	for _, testcase := range testSet {
		w := httptest.NewRecorder()
		bodyreader := bytes.NewReader(testcase.in)
		req, _ := http.NewRequest("POST", "/goods/release", bodyreader)
		r.ServeHTTP(w, req)

		assert.Equal(t, testcase.resCode, w.Code)

	}
}

func TestHandler_getRemainGoods(t *testing.T) {

	type testCase struct {
		in            int
		resCode       int
		expectedLen   int
		expectedValue []domain.Good
	}

	testSet := []testCase{
		{
			in:          1,
			resCode:     200,
			expectedLen: 1,
			expectedValue: []domain.Good{
				{
					ID:       1,
					Name:     "name1",
					Size:     "1x1x1",
					Quantity: 10,
				},
			},
		},
		{
			in:            0,
			resCode:       400,
			expectedLen:   0,
			expectedValue: []domain.Good{},
		},
		{
			in:            2,
			resCode:       500,
			expectedLen:   0,
			expectedValue: []domain.Good{},
		},
	}

	s := storage_mock.NewStorageMock()
	a := app.NewApp(s, zap.NewNop())
	h := NewHandler(a)
	r := NewRouter(h, "prod")

	for _, testcase := range testSet {
		w := httptest.NewRecorder()

		path := fmt.Sprintf("/warehouse/%d/goods", testcase.in)
		req, _ := http.NewRequest("GET", path, nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, testcase.resCode, w.Code)

		var responseBody []byte
		responseBody, err := io.ReadAll(w.Body)
		if err != nil {
			t.Errorf("cant read response body")
		}

		answer := make([]domain.Good, 0, 1)

		goods := struct {
			Goods []domain.Good `json:"goods"`
		}{Goods: answer}

		json.Unmarshal(responseBody, &goods)

		if testcase.expectedLen > 0 {
			assert.Equal(t, testcase.expectedValue[0].ID, goods.Goods[0].ID)
			assert.Equal(t, testcase.expectedValue[0].Name, goods.Goods[0].Name)
			assert.Equal(t, testcase.expectedValue[0].Size, goods.Goods[0].Size)
			assert.Equal(t, testcase.expectedValue[0].Quantity, goods.Goods[0].Quantity)
		}
	}
}
