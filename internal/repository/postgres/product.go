package postgres

import (
	"context"
	"errors"

	"github.com/imranzahaev/warehouse/internal/domain"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent/predicate"

	entproduct "github.com/imranzahaev/warehouse/internal/repository/postgres/ent/product"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent/user"
)

type ProductRepo struct {
	client *ent.Client
}

func NewProductRepo(client *ent.Client) *ProductRepo {
	return &ProductRepo{
		client: client,
	}
}

// Create new product
func (r *ProductRepo) Create(ctx context.Context, uid int, product domain.Product) (int, error) {
	res, err := r.client.Product.
		Create().
		SetArticle(product.Article).
		SetName(product.Name).
		SetPrice(product.Price).
		SetOwnerID(uid).
		Save(ctx)
	if err != nil {
		switch {
		case ent.IsValidationError(err):
			e := err.(*ent.ValidationError)
			return 0, domain.NewValidationError(e.Name, errors.Unwrap(e.Unwrap()))
		}
		return 0, err
	}

	return res.ID, nil
}

// Update product by id
func (r *ProductRepo) Update(ctx context.Context, uid int, product domain.Product) error {
	if product.ID == 0 {
		return domain.AppError{Err: errors.New("empty product id")}
	}

	if uid == 0 {
		return domain.AppError{Err: errors.New("empty user id")}
	}

	u := r.client.Product.Update().
		Where(entproduct.And(
			entproduct.ID(product.ID),
			entproduct.HasOwnerWith(user.ID(uid)),
		))
	if product.Article != "" {
		u = u.SetArticle(product.Article)
	}
	if product.Name != "" {
		u = u.SetName(product.Name)
	}
	if product.Price != 0 {
		u = u.SetPrice(product.Price)
	}

	c, err := u.Save(ctx)
	if err != nil {
		switch {
		case ent.IsValidationError(err):
			e := err.(*ent.ValidationError)
			return domain.NewValidationError(e.Name, errors.Unwrap(e.Unwrap()))
		case ent.IsNotFound(err):
			return domain.ErrProductNotFound
		}

		return err
	}

	if c == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}

// Get product by id
func (r *ProductRepo) Get(ctx context.Context, uid, id int) (domain.Product, error) {
	p, err := r.client.Product.Query().
		Where(entproduct.ID(id)).
		Where(entproduct.HasOwnerWith(user.ID(uid))).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return domain.Product{}, domain.ErrProductNotFound
		}
		return domain.Product{}, err
	}

	return convertProductToDomain(p), nil
}

// product?article_eq=abc&article_like=abc&
// 			name_eq=abc&name_like=abc&
// 			price_from=0&price_to=100
// 			page=1&count=10&
// 			sort_field=abc&sort_order=asc

func (r *ProductRepo) GetManyByFilter(ctx context.Context, filter domain.GetManyProductsFilter) (domain.GetManyProductsResponse, error) {
	q := r.client.Product.Query().
		Where(entproduct.HasOwnerWith(user.ID(filter.UID)))

	p := make([]predicate.Product, 0, 4)
	if filter.ArticleLike != "" {
		p = append(p, entproduct.ArticleContains(filter.ArticleLike))
	}
	if filter.NameLike != "" {
		p = append(p, entproduct.NameContains(filter.NameLike))
	}
	if filter.PriceFrom != 0 {
		p = append(p, entproduct.PriceGTE(filter.PriceFrom))
	}
	if filter.PriceTo != 0 {
		p = append(p, entproduct.PriceLTE(filter.PriceTo))
	}

	if len(p) > 0 {
		q = q.Where(entproduct.And(p...))
	}

	count, err := q.Count(ctx)
	if err != nil {
		return domain.GetManyProductsResponse{}, err
	}

	if count == 0 {
		return domain.GetManyProductsResponse{
			Page:  filter.Page,
			Limit: filter.Limit,
			Count: count,
		}, nil
	}

	if filter.Limit == 0 {
		filter.Limit = 20
	}
	q = q.Limit(filter.Limit)

	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Page != 1 {
		offset := (filter.Page - 1) * filter.Limit
		q = q.Offset(offset)
	}

	var order func(fields ...string) ent.OrderFunc
	switch filter.SortOrder {
	case "asc":
		order = ent.Asc
	default:
		order = ent.Desc
	}

	var sortField string
	switch filter.SortField {
	case "article":
		sortField = entproduct.FieldArticle
	case "name":
		sortField = entproduct.FieldName
	case "price":
		sortField = entproduct.FieldPrice
	case "created":
		sortField = entproduct.FieldCreatedAt
	default:
		sortField = entproduct.FieldID
	}
	q = q.Order(order(sortField))

	ps, err := q.All(ctx)
	if err != nil || len(ps) == 0 {
		return domain.GetManyProductsResponse{}, err
	}

	products := make([]domain.Product, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		products = append(products, convertProductToDomain(ps[i]))
	}

	return domain.GetManyProductsResponse{
		Products: products,
		Page:     filter.Page,
		Limit:    filter.Limit,
		Count:    count,
	}, nil
}

func convertProductToDomain(product *ent.Product) domain.Product {
	return domain.Product{
		ID:        product.ID,
		Name:      product.Name,
		Article:   product.Article,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
	}
}
