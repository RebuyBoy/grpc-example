package main

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "productinfo/service/ecommerce"
)

type server struct {
	pb.UnimplementedProductInfoServer
	productMap map[string]*pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {

	productId := uuid.New()
	in.Id = productId.String()

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()

}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	product, exists := s.productMap[in.Value]
	if exists {
		log.Printf("Product found: %v", product)
		return product, status.New(codes.OK, "").Err()
	}
	return &pb.Product{}, status.Errorf(codes.NotFound, "Product does not exist.")
}
