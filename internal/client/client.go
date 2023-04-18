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
	client, err := jsonrpc.Dial("tcp", "localhost:80")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	//uploadProduct(client)
	//updateProduct(client)
	//getProduct(client)
	//listProduct(client)
	//deleteProduct(client)
	//reserveProduct(client)
	releaseReserve(client)
}

func uploadProduct(client *rpc.Client) {
	req := model.UploadProductRequest{
		Name:        "Hoodie",
		Size:        "50/52",
		Code:        "33278291",
		Quantity:    131,
		WarehouseID: "1e8d248b-eca8-423d-b0f7-cdd11f959775",
	}

	var resp model.UploadProductResponse
	err := client.Call("ProductImpl.UploadProduct", req, &resp)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Println(resp)
}

func updateProduct(client *rpc.Client) {
	req := model.UpdateProductRequest{
		Name:        "Pants",
		Size:        "52/54",
		Code:        "33278291",
		Quantity:    130,
		WarehouseID: "1e8d248b-eca8-423d-b0f7-cdd11f959775",
	}

	var resp model.Response
	err := client.Call("ProductImpl.UpdateProduct", req, &resp)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Println(resp)

}

func getProduct(client *rpc.Client) {
	req := model.GetProductRequest{
		ProductCode: "73271291",
		WarehouseID: "ce6ba671-e63c-4bd9-99af-f2f0cd110498",
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
		WarehouseID: "ce6ba671-e63c-4bd9-99af-f2f0cd110498",
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
		ProductCode: "33278291",
		WarehouseID: "1e8d248b-eca8-423d-b0f7-cdd11f959775",
	}

	var resp model.Response
	err := client.Call("ProductImpl.DeleteProduct", req, &resp)
	if err != nil {
		l.Fatal("error due call rpc", zap.Any("response", resp), zap.Error(err))
	}

	fmt.Println(resp)
}

func reserveProduct(client *rpc.Client) {
	req := model.ReserveProductRequest{
		ProductCode: []string{
			"73278291",
		},
		WarehouseID: "ce6ba671-e63c-4bd9-99af-f2f0cd110498",
		Quantity:    17,
	}

	var resp model.Response
	err := client.Call("ProductImpl.ReserveProduct", req, &resp)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Println(resp)
}

func releaseReserve(client *rpc.Client) {
	req := model.ReleaseReserveRequest{
		ProductCode: []string{
			"33278291",
		},
		WarehouseID: "72d49a8d-4d6d-4371-8c94-83e5d5dc0e99",
	}

	var resp model.Response
	err := client.Call("ProductImpl.ReleaseReserve", req, &resp)
	if err != nil {
		l.Fatal("error due call rpc", zap.Error(err))
	}

	fmt.Println(resp)
}
