package main

import (
	pb "github.com/thiendsu2303/go-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlog(document *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       document.ID.Hex(),
		AuthorId: document.AuthorId,
		Title:    document.Title,
		Content:  document.Content,
	}
}
