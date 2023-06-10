package helpers

import "flowable-cash-backend/account"

func AccountNoDecider(value string) string {
	switch value {

	case "Kas":
		return account.Kas

	case "Persediaan Barang Dagang":
		return account.PersediaanBarangDagang

	case "Device":
		return account.Device

	case "Hutang Dagang":
		return account.HutangDagang

	case "Modal":
		return account.ModalOwner

	case "Laba Disimpan":
		return account.LabaDitahan

	case "Mengambil Laba":
		return account.Prive

	case "Penjualan":
		return account.Penjualan

	case "Pembelian":
		return account.Pembelian

	case "Beban Pembelian":
		return account.BebanPembelian

	case "Beban Penjualan":
		return account.BebanPenjualan

	case "Beban Operasional":
		return account.BebanOperasional
	}

	return ""
}
