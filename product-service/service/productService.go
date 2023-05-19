package service

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/bernardn38/ebuy/product-service/models"
	"github.com/bernardn38/ebuy/product-service/sql/products"
	"github.com/redis/go-redis/v9"
)

type ProductService struct {
	productQuries *products.Queries
	productDb     *sql.DB
	redisClient   *redis.Client
}

func NewProductService(p *products.Queries, db *sql.DB, rdb *redis.Client) *ProductService {
	return &ProductService{
		productQuries: p,
		productDb:     db,
		redisClient:   rdb,
	}
}
func (p *ProductService) RunAsync() {
	p.UpdateCacheAsync()
}
func (p *ProductService) CreateProduct(ctx context.Context, product models.CreateProductParams) (int32, error) {
	productId, err := p.productQuries.CreateProduct(ctx, products.CreateProductParams{
		UserID:      product.UserId,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	})
	if err != nil {
		return 0, err
	}
	return productId, nil
}
func (p *ProductService) GetAllProducts(ctx context.Context) ([]products.Product, error) {
	products, err := p.productQuries.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductService) GetRecentUploadedProducts(ctx context.Context) ([]products.Product, error) {
	var recentProducts []products.Product
	val, err := p.redisClient.Get(ctx, "recentProductsHome").Result()
	if err != nil {
		recentProducts, err = p.productQuries.GetRecentProducts(ctx)
		if err != nil {
			return nil, err
		}
	}
	err = json.Unmarshal([]byte(val), &recentProducts)
	if err != nil {
		return nil, err
	}
	return recentProducts, nil
}
