package mocktest

// import (
// 	pbp "api-gateway/genproto/book"
// 	"context"

// 	"google.golang.org/grpc"
// )

// type BookServiceClient interface {
// 	CreateBook(ctx context.Context, in *pbp.Book, opts ...grpc.CallOption) (*pbp.Book, error)
// 	GetBookById(ctx context.Context, in *pbp.BookId, opts ...grpc.CallOption) (*pbp.Book, error)
// 	UpdateBook(ctx context.Context, in *pbp.Book, opts ...grpc.CallOption) (*pbp.Book, error)
// 	DeleteBook(ctx context.Context, in *pbp.BookId, opts ...grpc.CallOption) (*pbp.Status, error)
// 	ListBooks(ctx context.Context, in *pbp.GetAllBookRequest, opts ...grpc.CallOption) (*pbp.GetAllBookResponse, error)
// }

// type bookServiceClient struct {
// }

// func NewBookServiceClient() BookServiceClient {
// 	return &bookServiceClient{}
// }

// func (c *bookServiceClient) CreateBook(ctx context.Context, in *pbp.Book, opts ...grpc.CallOption) (*pbp.Book, error) {
// 	return in, nil
// }

// func (c *bookServiceClient) GetBookById(ctx context.Context, in *pbp.BookId, opts ...grpc.CallOption) (*pbp.Book, error) {

// 	return &pbp.Book{
// 		Id:          "1",
// 		Name:        "Book name",
// 		Description: "Book description",
// 		Price:       9.99,
// 		Amount:      19,
// 	}, nil
// }

// func (c *bookServiceClient) UpdateBook(ctx context.Context, in *pbp.Book, opts ...grpc.CallOption) (*pbp.Book, error) {
// 	return &pbp.Book{
// 		Id:          "1",
// 		Name:        "Book name",
// 		Description: "Book description",
// 		Price:       9.99,
// 		Amount:      51,
// 	}, nil
// }

// func (c *bookServiceClient) DeleteBook(ctx context.Context, in *pbp.BookId, opts ...grpc.CallOption) (*pbp.Status, error) {
// 	return &pbp.Status{
// 		Success: true,
// 	}, nil
// }

// func (c *bookServiceClient) ListBooks(ctx context.Context, in *pbp.GetAllBookRequest, opts ...grpc.CallOption) (*pbp.GetAllBookResponse, error) {
// 	pr := pbp.Book{
// 		Id:          "1",
// 		Name:        "Book name",
// 		Description: "Book description",
// 		Price:       9.99,
// 		Amount:      15,
// 	}
// 	return &pbp.GetAllBookResponse{
// 		Count: 7,
// 		Books: []*pbp.Book{
// 			&pr,
// 			&pr,
// 			&pr,
// 			&pr,
// 			&pr,
// 			&pr,
// 			&pr,
// 		},
// 	}, nil
// }
