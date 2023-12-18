package store

const (
	createQuery  = `INSERT INTO product (id,name,minAmount,maxAmount,purchaseAmount,enable)values(?,?,?,?,?,?)`
	updateQuery  = `UPDATE product SET name=?, minAmount=? maxAmount=? purchaseAmount=? enable=? WHERE id=?`
	getByIDQuery = `SELECT id,name,minAmount,maxAmount,purchaseAmount,enable FROM product WHERE id=?`
	deleteQuery  = `Delete from product where id =?`
	getAllQuery  = `SELECT id,name,minAmount,maxAmount,purchaseAmount,enable FROM product`
)
