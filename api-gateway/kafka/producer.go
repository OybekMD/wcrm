package kafka

import (
	"api-gateway/api/handlers/models"
	"context"
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

var inst ProduceMessages

func Init(store ProduceMessages) {
	inst = store
}

func ProduceUser(ctx context.Context, key string, proto models.User) error {
	return inst.ProduceUser(ctx, key, proto)
}

func ProduceCategoryIcon(ctx context.Context, key string, proto models.CategoryIcon) error {
	return inst.ProduceCategoryIcon(ctx, key, proto)
}

func ProduceCategory(ctx context.Context, key string, proto models.Category) error {
	return inst.ProduceCategory(ctx, key, proto)
}

func ProduceProduct(ctx context.Context, key string, proto models.Product) error {
	return inst.ProduceProduct(ctx, key, proto)
}

func ProduceOrderproduct(ctx context.Context, key string, proto models.Orderproduct) error {
	return inst.ProduceOrderproduct(ctx, key, proto)
}

func ProduceComment(ctx context.Context, key string, proto models.Comment) error {
	return inst.ProduceComment(ctx, key, proto)
}

func Close() {
	inst.Close()
}
