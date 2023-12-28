package router

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"zadanie/internal/app"
	"zadanie/internal/storage/storage_mock"

	"github.com/stretchr/testify/assert"
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
	a := app.NewApp(s)
	h := NewHandler(a)
	r := NewRouter(h, "dev")

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
