package main

import (
	"context"
	"fmt"
	"time"

	cli "MiniProject/git.campus.id/buku2/endpoint"
	//svc "MiniProject/git.campus.id/buku2/server"
	opt "MiniProject/git.campus.id/util/grpc"
	util "MiniProject/git.campus.id/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCBukuClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Buku
	//client.AddBukuService(context.Background(), svc.Buku{KodeBuku: "B009", NamaBuku: "The Eye", Keterangan: "Siap Pakai", CreateBy: "dzaky"})

	//Get Mahasiswa By Nim No
	//kode, _ := client.ReadBukuByKodeService(context.Background(), "A%")
	//fmt.Println("Buku based on KodeBuku:", kode)

	//List Buku
	//buku, _ := client.ReadBukuService(context.Background())
	//fmt.Println("allBuku:", buku)

	//Update Buku
	//client.UpdateBukuService(context.Background(), svc.Buku{KodeBuku: "B002", NamaBuku: "Komputer", Status: 1, UpdateBy: "dzaky"})

	//Get Mahasiswa By Nim No
	keterangan, _ := client.ReadBukuByKeteranganService(context.Background(), "A%")
	fmt.Println("Buku based on Keterangan:", keterangan)
	/*
		//Get Customer By Email
		cusEmail, _ := client.ReadCustomerByEmailService(context.Background(), "joko@gmail.com")
		fmt.Println("customer based on email:", cusEmail)
	*/
}
