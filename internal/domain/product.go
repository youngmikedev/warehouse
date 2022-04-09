package domain

import "time"

type Product struct {
	ID        int
	Price     int
	Article   string
	Name      string
	CreatedAt time.Time
}

type GetManyProductsFilter struct {
	UID int

	ArticleLike string
	NameLike    string

	PriceFrom int
	PriceTo   int

	Page  int
	Limit int

	// One of:
	//  id
	//  article
	//  name
	//  price
	//  created
	SortField string
	// ASC or DESC
	SortOrder string
}

type GetManyProductsResponse struct {
	Products []Product
	Page     int
	Limit    int

	Count int
}
