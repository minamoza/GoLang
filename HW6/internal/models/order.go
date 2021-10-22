package models

type Order struct{
	ID 		int 	`json: order_id`
	Name 	string `json: name`
	Adress 	string `json: adress`
	Order 	string `json: order`
}
