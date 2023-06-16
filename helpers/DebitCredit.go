package helpers

import "flowable-cash-backend/models"

func DebitCreditDecider(accountName string, total uint) models.Balance {
	switch accountName {

	case "Kas":
		return models.Balance{
			Debit:  float64(total),
			Credit: 0,
		}

	case "Persediaan Barang Dagang":
		return models.Balance{
			Debit:  float64(total),
			Credit: 0,
		}

	case "Perlengkapan":
		return models.Balance{
			Debit:  float64(total),
			Credit: 0,
		}

	case "Akumulasi Penyusutan Perlengkapan":
		return models.Balance{
			Debit:  0,
			Credit: float64(total),
		}

	case "Hutang Dagang":
		return models.Balance{
			Debit:  0,
			Credit: float64(total),
		}

	case "Modal":
		return models.Balance{
			Debit:  0,
			Credit: float64(total),
		}

	case "Laba Disimpan":
		return models.Balance{
			Debit:  0,
			Credit: float64(total),
		}

	case "Mengambil Laba":
		return models.Balance{
			Debit:  0,
			Credit: float64(total),
		}

	case "Beban Penjualan":
		return models.Balance{
			Debit:  float64(total),
			Credit: 0,
		}

	case "Beban Pembelian":
		return models.Balance{
			Debit:  float64(total),
			Credit: 0,
		}
	}

	return models.Balance{}
}
