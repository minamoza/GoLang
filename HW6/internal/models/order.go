package models

type Order struct{
	ID 		int 	`json:"orderId" db:"orderID"`
	Name 	string `json:"userName" db:"userName"`
	Address string `json:"address" db:"address"`
	Order 	[]Dish `json:"order" db:"order"`
}

type Category struct{
	ID 		int 	`json:"categoryId" db:"categoryID"`
	Name 	string  `json:"categoryName" db:"categoryName"`
}

type Restaurant struct{
	ID 			int 	`json:"rest_id" db:"rest_id"`
	CategoryId  int 	`json:"category_id" db:"category_id"`
	Name 		string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Menu 		[]Dish `json:"menu" db:"menu"`
	Address 	string `json:"address" db:"address"`
	Rate 		float32 `json:"rate" db:"rate"`
	Commentary []Comment `json:"comments" db:"comments"`
}

type User struct{
	ID 		int 	`json:"user_id" db:"id"`
	Name 	string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
	Email 	string `json:"email" db:"email"`
	Phone 	string `json:"phone" db:"phone"`
	Address string `json:"address" db:"address"`
}

type Dish struct{
	Dish 		string  `json:"dish" db:"dish"`
	Cost 		float32 `json:"cost" db:"cost"`
	Ingredients []string `json:"ingredients" db:"ingredients"`
}

type Comment struct{
	Text string `json:"comment" db:"comment"`
	Rate float32 `json:"rate" db:"rate"`
	Name string `json:"username" db:"username"'`
}


