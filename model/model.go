package model

type Product struct {
	ID             int    `json:"ID"`             //productId
	Name           string `json:"name"`           //productName
	MinAmount      int    `json:"minAmount"`      //minimumAmount of the product
	MaxAmount      int    `json:"maxAmount"`      //maximumAmount of the product
	PurchaseAmount int    `json:"purchaseAmount"` //purchaseAmount of the product
	Enable         int    `json:"enable"`         //if the product is available or not
}
