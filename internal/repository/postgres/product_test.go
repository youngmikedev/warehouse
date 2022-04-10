package postgres

import (
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/youngmikedev/warehouse/internal/domain"
	"github.com/youngmikedev/warehouse/internal/repository/postgres/ent"
)

func initTestDataset(t *testing.T, rowCount int) (uid int) {
	uid, _, _ = createUserAndGetTokens(t, "", "")
	uid2, _, _ := createUserAndGetTokens(t, "seconduid@test.com", "Tom")
	var pcount = rowCount
	var batchSize = 10

	tmpData := make([]*ent.ProductCreate, batchSize)
	for i := 0; i < pcount; i++ {
		tmpData[i%batchSize] = db.Product.Create().
			SetArticle(strconv.Itoa(i + 1)).
			SetName("Product no: " + strconv.Itoa(i+1)).
			SetPrice(1000 + i).
			SetOwnerID(uid)

		if i != 0 && (i%batchSize == 9 || i+1 == pcount) {
			_, err := db.Product.CreateBulk(tmpData...).Save(context.Background())
			if err != nil {
				t.Fatalf("initTestDataset.CreateBulk() error = %v", err)
			}
		}
	}

	for i := 0; i < batchSize; i++ {
		tmpData[i%batchSize] = db.Product.Create().
			SetArticle(strconv.Itoa(i + 1)).
			SetName("Product no: " + strconv.Itoa(i+1)).
			SetPrice(1000 + i).
			SetOwnerID(uid2)

		if i != 0 && (i%batchSize == 9 || i+1 == pcount) {
			_, err := db.Product.CreateBulk(tmpData...).Save(context.Background())
			if err != nil {
				t.Fatalf("initTestDataset.CreateBulk() error = %v", err)
			}
		}
	}

	return uid
}

func TestProductRepo_GetManyByFilter(t *testing.T) {
	uid := initTestDataset(t, 100)
	type args struct {
		ctx    context.Context
		filter domain.GetManyProductsFilter
	}
	tests := []struct {
		name    string
		args    args
		want    domain.GetManyProductsResponse
		wantErr bool
	}{
		{
			name: "1. Valid",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:       uid,
					SortField: "id",
					Limit:     2,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      100,
						Article: "100",
						Name:    "Product no: 100",
						Price:   1099,
					},
					{
						ID:      99,
						Article: "99",
						Name:    "Product no: 99",
						Price:   1098,
					},
				},
				Page:  1,
				Limit: 2,
				Count: 100,
			},
		},
		{
			name: "2. Valid page 10",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:   uid,
					Page:  10,
					Limit: 2,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      82,
						Article: "82",
						Name:    "Product no: 82",
						Price:   1081,
					},
					{
						ID:      81,
						Article: "81",
						Name:    "Product no: 81",
						Price:   1080,
					},
				},
				Page:  10,
				Limit: 2,
				Count: 100,
			},
		},
		{
			name: "3. Valid sort article asc",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:       uid,
					SortField: "article",
					SortOrder: "asc",
					Limit:     2,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      1,
						Article: "1",
						Name:    "Product no: 1",
						Price:   1000,
					},
					{
						ID:      10,
						Article: "10",
						Name:    "Product no: 10",
						Price:   1009,
					},
				},
				Page:  1,
				Limit: 2,
				Count: 100,
			},
		},
		{
			name: "4. Valid name like 65",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:      uid,
					NameLike: "65",
					Limit:    2,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      65,
						Article: "65",
						Name:    "Product no: 65",
						Price:   1064,
					},
				},
				Page:  1,
				Limit: 2,
				Count: 1,
			},
		},
		{
			name: "5. Valid article like 7",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:         uid,
					ArticleLike: "7",
					Limit:       2,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      97,
						Article: "97",
						Name:    "Product no: 97",
						Price:   1096,
					},
					{
						ID:      87,
						Article: "87",
						Name:    "Product no: 87",
						Price:   1086,
					},
				},
				Page:  1,
				Limit: 2,
				Count: 19,
			},
		},
		{
			name: "6. Valid price from 1050",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:       uid,
					PriceFrom: 1050,
					Limit:     2,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      100,
						Article: "100",
						Name:    "Product no: 100",
						Price:   1099,
					},
					{
						ID:      99,
						Article: "99",
						Name:    "Product no: 99",
						Price:   1098,
					},
				},
				Page:  1,
				Limit: 2,
				Count: 50,
			},
		},
		{
			name: "7. Valid price to 1060",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:     uid,
					PriceTo: 1060,
					Limit:   2,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      61,
						Article: "61",
						Name:    "Product no: 61",
						Price:   1060,
					},
					{
						ID:      60,
						Article: "60",
						Name:    "Product no: 60",
						Price:   1059,
					},
				},
				Page:  1,
				Limit: 2,
				Count: 61,
			},
		},
		{
			name: "8. Price from 1010 to 1011",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:       uid,
					PriceFrom: 1010,
					PriceTo:   1011,
					Limit:     2,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      12,
						Article: "12",
						Name:    "Product no: 12",
						Price:   1011,
					},
					{
						ID:      11,
						Article: "11",
						Name:    "Product no: 11",
						Price:   1010,
					},
				},
				Page:  1,
				Limit: 2,
				Count: 2,
			},
		},
		{
			name: "9. All filters",
			args: args{
				ctx: context.Background(),
				filter: domain.GetManyProductsFilter{
					UID:         uid,
					ArticleLike: "2",
					NameLike:    "1",
					PriceFrom:   1000,
					PriceTo:     1050,
				},
			},
			want: domain.GetManyProductsResponse{
				Products: []domain.Product{
					{
						ID:      21,
						Article: "21",
						Name:    "Product no: 21",
						Price:   1020,
					},
					{
						ID:      12,
						Article: "12",
						Name:    "Product no: 12",
						Price:   1011,
					},
				},
				Page:  1,
				Limit: 20,
				Count: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProductRepo{
				client: db,
			}
			got, err := r.GetManyByFilter(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepo.GetManyByFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := 0; i < len(got.Products); i++ {
				got.Products[i].CreatedAt = time.Time{}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepo.GetManyByFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepo_Create(t *testing.T) {
	uid, _, _ := createUserAndGetTokens(t, "", "")
	type args struct {
		ctx     context.Context
		uid     int
		product domain.Product
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "1. Valid",
			args: args{
				ctx: context.TODO(),
				uid: uid,
				product: domain.Product{
					Article: "0001",
					Name:    "t-short",
					Price:   100,
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProductRepo{
				client: db,
			}
			got, err := r.Create(tt.args.ctx, tt.args.uid, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ProductRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepo_Get(t *testing.T) {
	uid := initTestDataset(t, 10)
	type args struct {
		ctx context.Context
		uid int
		id  int
	}
	tests := []struct {
		name        string
		args        args
		wantProduct domain.Product
		wantErr     bool
		errType     error
	}{
		{
			name: "1. Valid",
			args: args{
				ctx: context.Background(),
				uid: uid,
				id:  1,
			},
			wantProduct: domain.Product{
				ID:      1,
				Article: "1",
				Name:    "Product no: 1",
				Price:   1000,
			},
		},
		{
			name: "2. Invalid user id",
			args: args{
				ctx: context.Background(),
				uid: uid,
				id:  15,
			},
			wantErr: true,
			errType: domain.ErrProductNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProductRepo{
				client: db,
			}
			gotProduct, err := r.Get(tt.args.ctx, tt.args.uid, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.errType != nil && tt.errType != err {
				t.Errorf("ProductRepo.Get() error = %v, wantErrType %v", err, tt.errType)
				return
			}
			gotProduct.CreatedAt = time.Time{}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("ProductRepo.Get() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func TestProductRepo_Update(t *testing.T) {
	uid := initTestDataset(t, 10)
	type args struct {
		ctx     context.Context
		uid     int
		product domain.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1. Valid",
			args: args{
				ctx: context.Background(),
				uid: uid,
				product: domain.Product{
					ID:      1,
					Article: "test",
					Name:    "test_name",
					Price:   99,
				},
			},
		},
		{
			name: "2. Invalid user id",
			args: args{
				ctx: context.Background(),
				uid: uid + 1000,
				product: domain.Product{
					ID:      1,
					Article: "test",
					Name:    "test_name",
					Price:   99,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ProductRepo{
				client: db,
			}
			if err := r.Update(tt.args.ctx, tt.args.uid, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("ProductRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			p, err := r.client.Product.Get(context.Background(), tt.args.product.ID)
			if err != nil {
				t.Errorf("Product.Get() failed get product by id %v, error = %v", tt.args.product.ID, err)
			}

			p.CreatedAt = time.Time{}
			got := convertProductToDomain(p)
			if !reflect.DeepEqual(tt.args.product, got) {
				t.Errorf("ProductRepo.Update() = %v, want %v", got, tt.args.product)
			}
		})
	}
}
