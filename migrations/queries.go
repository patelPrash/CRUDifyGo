package migrations

const (
	createProductTable = "CREATE TABLE IF NOT EXISTS product (ID INT, Name VARCHAR(255), MinAmount INT, MaxAmount INT, PurchaseAmount INT, Enable INT);"
	dropProductTable   = "DROP TABLE IF EXISTS product"
)
