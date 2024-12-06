package service_test

import (
	"encoding/json"
	"net/http"
	"testing"
)

type Order struct {
	ID       int32 `json:"id"`
	Item     string `json:"item"`
	Quantity int32  `json:"quantity"`
}

type OrdersResponse struct {
	Orders []Order `json:"orders"`
}

func TestGetOrderListREST(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost/rest", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("GetOrderList failed: %v", err)
	}
	var ordersResponse OrdersResponse
	err = json.NewDecoder(resp.Body).Decode(&ordersResponse)
	if err != nil {
		t.Fatalf("GetOrderList failed: %v", err)
	}
}
