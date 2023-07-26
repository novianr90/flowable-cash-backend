package helpers

import "flowable-cash-backend/account"

func AccountNoDecider(value string) string {
	switch value {

	case "Kas":
		return account.Kas

	case "Persediaan Barang Dagang":
		return account.PersediaanBarangDagang

	case "Perlengkapan":
		return account.Perlengkapan

	case "Akumulasi Penyusutan Perlengkapan":
		return account.AkumulasiPenyusutanPerlengkapan

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

	case "Biaya Pengemasan":
		return account.BebanPengemasan

	case "Biaya Ongkos Kirim":
		return account.BebanOngkosKirim

	case "Beban Penyusutan":
		return account.BebanPenyusutan

	case "Beban Lainnya":
		return account.BebanLainnya

	case "Pembelian":
		return account.Pembelian

	case "Piutang Dagang":
		return account.PiutangDagang
	}

	return ""
}
