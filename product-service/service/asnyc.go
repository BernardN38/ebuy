package service

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

func (p *ProductService) UpdateCacheAsync() {
	ticker := time.NewTicker(time.Second * 30)
	for range ticker.C {
		go p.UpdateRecentProducts()
	}
}
func (p *ProductService) UpdateRecentProducts() {
	ctx := context.Background()
	recentProducts, err := p.productQuries.GetRecentProducts(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	productsBytes, err := json.Marshal(recentProducts)
	if err != nil {
		log.Println(err)
		return
	}
	result, err := p.redisClient.Set(ctx, "recentProductsHome", string(productsBytes), time.Hour).Result()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("cache updated: ", result)
}
