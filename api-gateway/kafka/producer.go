package kafka

import (
	"context"
	"api-gateway/api/handlers/models"
)

type ProduceMessages interface {
	ProduceUser(ctx context.Context, key string, proto models.User) error
	ProduceCategoryIcon(ctx context.Context, key string, proto models.CategoryIcon) error
	ProduceCategory(ctx context.Context, key string, proto models.Category) error
	ProduceProduct(ctx context.Context, key string, proto models.Product) error
	ProduceOrderproduct(ctx context.Context, key string, proto models.Orderproduct) error
	ProduceComment(ctx context.Context, key string, proto models.Comment) error
	Close()
}
