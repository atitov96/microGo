package resolver

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	graphql1 "microGo/gen/graphql"
	"microGo/pkg/models"
)

type Resolver struct{}

// UpdateProfile is the resolver for the updateProfile field.
func (r *mutationResolver) UpdateProfile(ctx context.Context, input graphql1.UpdateProfileInput) (*models.User, error) {
	panic("not implemented")
}

// UploadAvatar is the resolver for the uploadAvatar field.
func (r *mutationResolver) UploadAvatar(ctx context.Context, file string) (*graphql1.UploadResponse, error) {
	panic("not implemented")
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	panic("not implemented")
}

// SearchUsers is the resolver for the searchUsers field.
func (r *queryResolver) SearchUsers(ctx context.Context, query string, limit *int, offset *int) ([]*models.User, error) {
	panic("not implemented")
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*models.User, error) {
	panic("not implemented")
}

// UserUpdated is the resolver for the userUpdated field.
func (r *subscriptionResolver) UserUpdated(ctx context.Context, id string) (<-chan *models.User, error) {
	panic("not implemented")
}

// Friends is the resolver for the friends field.
func (r *userResolver) Friends(ctx context.Context, obj *models.User) ([]*models.User, error) {
	panic("not implemented")
}

// FriendCount is the resolver for the friendCount field.
func (r *userResolver) FriendCount(ctx context.Context, obj *models.User) (int, error) {
	panic("not implemented")
}

// CreatedAt is the resolver for the createdAt field.
func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	panic("not implemented")
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *userResolver) UpdatedAt(ctx context.Context, obj *models.User) (string, error) {
	panic("not implemented")
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

// Subscription returns graphql1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graphql1.SubscriptionResolver { return &subscriptionResolver{r} }

// User returns graphql1.UserResolver implementation.
func (r *Resolver) User() graphql1.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
