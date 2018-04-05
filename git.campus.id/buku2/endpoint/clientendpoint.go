package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/git.campus.id/buku2/server"
)

func (be BukuEndpoint) AddBukuService(ctx context.Context, buku sv.Buku) error {
	_, err := be.AddBukuEndpoint(ctx, buku)
	return err
}

func (be BukuEndpoint) ReadBukuByKodeService(ctx context.Context, kode string) (sv.Buku, error) {
	req := sv.Buku{KodeBuku: kode}
	fmt.Println(req)
	resp, err := be.ReadBukuByKodeEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	bk := resp.(sv.Buku)
	return bk, err
}

func (be BukuEndpoint) ReadBukuService(ctx context.Context) (sv.Bukus, error) {
	resp, err := be.ReadBukuEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Bukus), err
}

func (be BukuEndpoint) UpdateBukuService(ctx context.Context, bk sv.Buku) error {
	_, err := be.UpdateBukuEndpoint(ctx, bk)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (be BukuEndpoint) ReadBukuByKeteranganService(ctx context.Context, keterangan string) (sv.Bukus, error) {
	req := sv.Buku{Keterangan: keterangan}
	resp, err := be.ReadBukuByKeteranganEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	bk := resp.(sv.Bukus)
	return bk, err
}
