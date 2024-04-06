package kafka

import (
	"api-gateway/api/handlers/models"
	"api-gateway/config"
	pbc "api-gateway/genproto/comment"
	pbp "api-gateway/genproto/post"
	pbu "api-gateway/genproto/user"
	"api-gateway/pkg/logger"
	"context"

	kafka "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Produce struct {
	Log          logger.Logger
	User         *kafka.Writer
	CategoryIcon *kafka.Writer
	Category     *kafka.Writer
	Product      *kafka.Writer
	Orderproduct *kafka.Writer
	Comment      *kafka.Writer
}

func NewProducerInit(conf config.Config, log logger.Logger) *Produce {
	return &Produce{
		Log: log,
		User: &kafka.Writer{
			Addr:                   kafka.TCP(conf.Kafka.Address),
			Topic:                  conf.Kafka.UserCreateTopic,
			AllowAutoTopicCreation: true,
		},
		CategoryIcon: &kafka.Writer{
			Addr:                   kafka.TCP(conf.Kafka.Address),
			Topic:                  conf.Kafka.CategoryIconCreateTopic,
			AllowAutoTopicCreation: true,
		},
		Category: &kafka.Writer{
			Addr:                   kafka.TCP(conf.Kafka.Address),
			Topic:                  conf.Kafka.CategoryCreateTopic,
			AllowAutoTopicCreation: true,
		},
		Product: &kafka.Writer{
			Addr:                   kafka.TCP(conf.Kafka.Address),
			Topic:                  conf.Kafka.ProductCreateTopic,
			AllowAutoTopicCreation: true,
		},
		Orderproduct: &kafka.Writer{
			Addr:                   kafka.TCP(conf.Kafka.Address),
			Topic:                  conf.Kafka.OrderproductCreateTopic,
			AllowAutoTopicCreation: true,
		},
		Comment: &kafka.Writer{
			Addr:                   kafka.TCP(conf.Kafka.Address),
			Topic:                  conf.Kafka.CommentCreateTopic,
			AllowAutoTopicCreation: true,
		},
	}
}

func (p *Produce) ProduceUser(ctx context.Context, key string, proto models.User) error {
	event := pbu.User{
		Id:           proto.Id,
		FirstName:    proto.FirstName,
		LastName:     proto.LastName,
		Username:     proto.Username,
		PhoneNumber:  proto.PhoneNumber,
		Bio:          proto.Bio,
		BirthDay:     proto.BirthDay,
		Email:        proto.Email,
		Avatar:       proto.Avatar,
		Password:     proto.Password,
		RefreshToken: proto.RefreshToken,
		CreatedAt:    proto.CreatedAt,
		UpdatedAt:    proto.UpdatedAt,
	}
	byteData, err := event.Marshal()
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(key),
		Value: byteData,
	}
	return p.User.WriteMessages(ctx, message)
}

func (p *Produce) ProduceCategoryIcon(ctx context.Context, key string, proto models.CategoryIcon) error {
	event := pbp.CategoryIcon{
		Id:      proto.Id,
		Name:    proto.Name,
		Picture: proto.Picture,
	}

	data, err := event.Marshal()
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	return p.Product.WriteMessages(ctx, message)
}

func (p *Produce) ProduceCategory(ctx context.Context, key string, proto models.Category) error {
	event := pbp.Category{
		Id:        proto.Id,
		Name:      proto.Name,
		IconId:    proto.IconId,
		CreatedAt: proto.CreatedAt,
		UpdatedAt: proto.UpdatedAt,
	}

	data, err := event.Marshal()
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	return p.Product.WriteMessages(ctx, message)
}

func (p *Produce) ProduceProduct(ctx context.Context, key string, proto models.Product) error {
	event := pbp.Product{
		Id:          proto.Id,
		Title:       proto.Title,
		Description: proto.Description,
		Price:       proto.Price,
		Picture:     proto.Picture,
		CategoryId:  proto.CategoryId,
		CreatedAt:   proto.CreatedAt,
		UpdatedAt:   proto.UpdatedAt,
	}

	data, err := event.Marshal()
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	return p.Product.WriteMessages(ctx, message)
}

func (p *Produce) ProduceOrderproduct(ctx context.Context, key string, proto models.Orderproduct) error {
	event := pbp.Orderproduct{
		Id:          proto.Id,
		UserId:       proto.UserId,
		ProductId: proto.ProductId,
		CreatedAt:   proto.CreatedAt,
	}

	data, err := event.Marshal()
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	return p.Product.WriteMessages(ctx, message)
}

func (p *Produce) ProduceComment(ctx context.Context, key string, proto models.Comment) error {
	event := pbc.Comment{
		Id:        proto.Id,
		Content:   proto.Content,
		UserId:    proto.UserId,
		ProductId: proto.ProductId,
		CreatedAt: proto.CreatedAt,
		UpdatedAt: proto.UpdatedAt,
	}

	data, err := event.Marshal()
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}

	return p.Comment.WriteMessages(ctx, message)
}

func (p *Produce) Close() {
	if err := p.User.Close(); err != nil {
		p.Log.Error("error while closing User create", zap.Error(err))
	}
	if err := p.CategoryIcon.Close(); err != nil {
		p.Log.Error("error while closing CategoryIcon create", zap.Error(err))
	}
	if err := p.Category.Close(); err != nil {
		p.Log.Error("error while closing Category create", zap.Error(err))
	}
	if err := p.Product.Close(); err != nil {
		p.Log.Error("error while closing Product create", zap.Error(err))
	}
	if err := p.Orderproduct.Close(); err != nil {
		p.Log.Error("error while closing Orderproduct create", zap.Error(err))
	}
	if err := p.Comment.Close(); err != nil {
		p.Log.Error("error while closing Comment create", zap.Error(err))
	}
}