package helpers

import "flowable-cash-backend/account"

func AccountNoDecider(value string) string {
	switch value {

	case "Kas":
		return account.Kas

	case "Bahan Baku":
		return account.BahanBaku

	case "Barang Dagang":
		return account.BarangDagang

	case "Bahan Tambahan":
		return account.BahanTambahan

	case "Piutang Dagang":
		return account.PiutangDagang

	case "Peralatan":
		return account.Peralatan

	case "Akumulasi Penyusutan Peralatan":
		return account.AkumulasiPenyusutanPeralatan

	case "Kendaraan":
		return account.Kendaraan

	case "Akumulasi Penyusutan Kendaraan":
		return account.AkumulasiPenyusutanKendaraan

	case "Hutang Dagang":
		return account.HutangDagang

	case "Modal":
		return account.ModalOwner

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
	}

	return ""
}
