package main

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/gofr/request"
)

func Test_Integration(t *testing.T) {
	go main()

	time.Sleep(5 * time.Second)

	productCreateBody := []byte(`{"id": 1, "name": "test", "minAmount": 100, "maxAmount": 400, "purchaseAmount": 200}`)
	productUpdateBody := []byte(`{"id": 1, "name": "test", "minAmount": 100, "maxAmount": 400, "purchaseAmount": 200}`)

	successResp := `{"data":{"ID": 1, "name": "test", "minAmount": 100, "maxAmount": 400, "purchaseAmount": 200}}`
	successUpdateResp := `{"data":{"ID": 1, "name": "test", "minAmount": 100, "maxAmount": 400, "purchaseAmount": 200}}`

	testCases := []struct {
		desc          string
		method        string
		endpoint      string
		body          []byte
		expStatusCode int
		expResp       string
	}{
		{"Create product", http.MethodPost, "/product", productCreateBody, http.StatusCreated,
			successResp},
		{"Get product", http.MethodGet, "/product/1", nil, http.StatusOK, successResp},
		{"Update product", http.MethodPut, "/product/1", productUpdateBody, http.StatusOK,
			successUpdateResp},
		{"Delete product", http.MethodDelete, "/product/1", nil, http.StatusNoContent, ``},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			req, _ := request.NewMock(tc.method, "http://localhost:9000"+tc.endpoint, bytes.NewBuffer(tc.body))
			client := http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error occurred in calling api: %v", err)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Error while reading response: %v", err)
			}

			respBody := strings.TrimSpace(string(body))

			assert.Equal(t, tc.expStatusCode, resp.StatusCode, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expResp, respBody, "Test [%d] failed", i+1)

			resp.Body.Close()
		})
	}
}
