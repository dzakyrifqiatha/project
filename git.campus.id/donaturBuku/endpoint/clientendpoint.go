package endpoint

import (
	"context"
	"fmt"

	sv "MiniProject/git.campus.id/donaturBuku/server"
)

func (dbe DonaturBukuEndpoint) AddDonaturBukuService(ctx context.Context, donaturBk sv.donaturBk) error {
	_, err := ce.AddDonaturBukuEndpoint(ctx, donaturBk)
	return err
}

func (dbe DonaturBukuEndpoint) ReadDonaturBukuByKd_DonaturBukuService(ctx context.Context, kode string) (sv.donaturBk, error) {
	req := sv.donaturBk{Kode: kode}
	fmt.Println(req)
	resp, err := dbe.ReadDonaturBukuByKd_DonaturBukuEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	dbk := resp.(sv.donaturBk)
	return dbk, err
}

func (dbe DonaturBukuEndpoint) ReadDonaturBukuService(ctx context.Context) (sv.donaturBk, error) {
	resp, err := ce.ReadDonaturBukuEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.donaturBk), err
}

func (dbe DonaturBukuEndpoint) UpdateDonaturBukuService(ctx context.Context, dbk sv.donaturBk) error {
	_, err := dbe.UpdateDonaturBukuEndpoint(ctx, dbk)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}

func (dbe DonaturBukuEndpoint) ReadDonaturBukuByNama_DonaturBukuService(ctx context.Context, nama string) (sv.donaturBk, error) {
	req := sv.donaturBk{Nama: nama}
	resp, err := dbe.ReadDonaturBukuByNama_DonaturBukuEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.donaturBk)
	return dbk, err
}
