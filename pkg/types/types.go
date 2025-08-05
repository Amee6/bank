package types 

// Money предсттавляет собой денежную единицу
type Money int64

// Category представляет собой категорию, в которой совершен платеж
type Category string

// Payment представляет информацию о платеже
type Payment struct {
	ID int64	
	Amount Money
	Category Category
}
