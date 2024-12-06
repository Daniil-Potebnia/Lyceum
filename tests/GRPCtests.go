package service_test

import (
	"context"
	"fmt"
	pb "service/pkg/api/order"
	"testing"

	"google.golang.org/grpc"
)

func TestCreateOrderGRPC(t *testing.T) {
	conn, err := grpc.NewClient("localhost:80")
	if err != nil {
		t.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderServiceClient(conn)

	req := &pb.CreateOrderRequest{
		Item: "Петя",
		Quantity: 2,
	}
	resp, err := client.CreateOrder(context.Background(), req)
	if err != nil {
        t.Fatalf("CreateOrder failed: %v", err)
    }
	if resp.Id == 0 {
		t.Fatalf("CreateOrder failed: %v", fmt.Errorf("invalid id"))
	}
}
