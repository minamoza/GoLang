package models


type (
	Dish struct{
		ID			int		`json:"id" db:"id"`
		Dish 		string  `json:"dish" db:"dish"`
		Cost 		float32 `json:"cost" db:"cost"`
		Ingredients []string `json:"ingredients" db:"ingredients"`
	}

	DishFilter struct{
		Query *string `json:"query"`
	}
)

//func (a DishFilter) Value() (driver.Value, error) {
//	return json.Marshal(a)
//}
//
//func (a Dish) Scan(value interface{}) error {
//	b, ok := value.([]byte)
//	if !ok {
//		return errors.New("type assertion to []byte failed")
//	}
//
//	return json.Unmarshal(b, &a)
//}

