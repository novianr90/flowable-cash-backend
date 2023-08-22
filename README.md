# FinTrack Backend

FinTrack is an accounting information system that MSMEs can use to provide financial
information for their needs. This is the codebase for BE that used to be consumed on FE.

### Features
- **Record Transactions**: Easily input, update, and delete financial transactions. Each transaction can include details such as date, description, amount, and accounts involved.
- **View Transactions**: Display a list of recorded transactions with filters and sorting options. This feature helps users quickly find and review specific transactions.
- **Income Summary**: Generate an income summary report that compiles all the income-related transactions over a specific period. This report aids in understanding revenue trends.
- **Expense Summary**: Obtain an overview of expenses by generating an expense summary report. It provides insights into where the company's money is being spent.
- **Cost of Goods Sold (COGS) Analysis**: Access a Cost of Goods Sold (COGS) report to analyze the costs directly associated with producing goods or services. This information is vital for evaluating profitability.
- **Profit and Loss Statement**: Generate a profit and loss statement that outlines the company's revenues, costs, and expenses over a specified period. This report helps assess the financial health of the business.
- **Balance Sheet**: View a balance sheet report that presents a snapshot of the company's financial position, including assets, liabilities, and equity.
 
### Tech
- [Gin](https://github.com/gin-gonic/gin) - Libraries framework to provide REST API
- [Gorm](https://gorm.io/index.html) - Libraries to handle query database
- [gocron](https://github.com/go-co-op/gocron) - Libraries to handle cron job

### Usage
##### Expenses
| Method | Description | Endpoint |
| :---:  | :---: | :---: | 
| POST | Record Expenses Transaction | api/pengeluaran |
| GET | Get All Expenses Transactions | api/pengeluaran |
| GET | Get 1 Expenses Transaction by ID | api/pengeluaran/ |

##### Incomes
| Method | Description | Endpoint |
| :---:  | :---: | :---: | 
| POST | Record Income Transaction | api/pemasukkan |
| GET | Get All Income Transaction | api/pemasukkan |
| GET | Get 1 Income Transaction by ID | api/pemasukkan/ |

##### Transactions
| Method | Description | Endpoint |
| :---:  | :---: | :---: | 
| GET | Get All Transaction | api/transactions |
| PUT | Update 1 Transaction by ID | api/transactions |
| DELETE | Delete 1 Transaction by ID | api/transactions |

##### Accounts
| Method | Description | Endpoint |
| :---:  | :---: | :---: | 
| GET | Get All Accounts | /api/balance-sheets |
| GET | Get Accounts by Account Name | /api/balance-sheets/ |
| PUT | Update 1 Transaction by ID | api/balance-sheets |

## License
Distributed under the MIT License. See `LICENSE.txt` for more information.
