package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/khusainnov/lamoda/internal/model"
	"go.uber.org/zap"
)

var l, _ = zap.NewProduction()

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:9001")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	uploadProduct(client)
	//updateProduct(client)
	//getProduct(client)
	//listProduct(client)
	//deleteProduct(client)
	//reserveProduct(client)
	//deleteReservation(client)
}

func uploadProduct(client *rpc.Client) {
	req := model.UploadProductRequest{
		Name:        "Zara shirt",
		Size:        "L",
		Code:        "71241291",
		Quantity:    51,
		WarehouseID: "433772c7-f85c-468c-9b59-a912d9231e47",
	}

	var resp model.UploadProductResponse
	err := client.Call("ProductImpl.UploadProduct", req, &resp)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Printf("Item: %v \tID: %s\n", req.Name, resp.Message)
}

func updateProduct(client *rpc.Client) {
	req := model.UpdateProductRequest{
		Name:        "Zara pants",
		Size:        "52/54",
		Code:        "71241290",
		Quantity:    101,
		WarehouseID: "d2129f75-8368-4269-82f0-58cf0d7305f8",
	}

	err := client.Call("ProductImpl.UpdateProduct", req, nil)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Printf("Product: %v\n", req.Name)

}

func getProduct(client *rpc.Client) {
	req := model.GetProductRequest{
		ProductCode: "71241290",
		WarehouseID: "d2129f75-8368-4269-82f0-58cf0d7305f8",
	}

	var resp model.Product
	err := client.Call("ProductImpl.GetProduct", req, &resp)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Printf("\v%v\n", resp)

}

func listProduct(client *rpc.Client) {
	req := model.ListProductRequest{
		WarehouseID: "c795cf7e-fa28-4946-a2cf-3a41a687815f",
	}

	var resp []model.Product
	err := client.Call("ProductImpl.ListProduct", req, &resp)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Printf("\v%v\n", resp)
}

func deleteProduct(client *rpc.Client) {
	req := model.DeleteProductRequest{
		ProductCode: "544724293",
		WarehouseID: "8e640e03-4540-4d07-a6ba-c0b62fa75335",
	}

	err := client.Call("ProductImpl.DeleteProduct", req, nil)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Printf("Req: %v\n", req)
}

func reserveProduct(client *rpc.Client) {
	req := model.ReserveProductRequest{
		ProductCode: []string{
			"162705324",
			"542242324",
		},
		WarehouseID: "c795cf7e-fa28-4946-a2cf-3a41a687815f",
		Quantity:    20,
	}

	err := client.Call("ProductImpl.ReserveProduct", req, nil)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

}

func deleteReservation(client *rpc.Client) {
	req := model.CancelReservationRequest{
		ProductCode: []string{
			"162705324",
		},
		WarehouseID: "c795cf7e-fa28-4946-a2cf-3a41a687815f",
	}

	err := client.Call("ProductImpl.CancelReservation", req, nil)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}
}
