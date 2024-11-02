package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"go-graphql-api/graph/generated"
	"go-graphql-api/graph/model"
	"time"
)

// CreatePost is the resolver for the CreatePost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	Addpost := model.Post{
		Title:       input.Title,
		Content:     input.Content,
		Author:      *input.Author,
		Hero:        *input.Hero,
		PublishedAt: time.Now().Format("01-02-2006"),
		UpdatedAt:   time.Now().Format("01-02-2006"),
	}

	if err := r.Database.Create(&Addpost).Error; err != nil {
		fmt.Println(err)
		return nil, err

	}

	return &Addpost, nil
}

// GetAllPosts is the resolver for the GetAllPosts field.
func (r *queryResolver) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	posts := []*model.Post{}

	GetPosts := r.Database.Model(&posts).Find(&posts)

	if GetPosts.Error != nil {
		fmt.Println(GetPosts.Error)
		return nil, GetPosts.Error
	}
	return posts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }