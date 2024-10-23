// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

type Mutation struct {
}

type Query struct {
}

type Subscription struct {
}

type UpdateProfileInput struct {
	Nickname *string `json:"nickname,omitempty"`
	Bio      *string `json:"bio,omitempty"`
}

type UploadResponse struct {
	Success bool    `json:"success"`
	URL     *string `json:"url,omitempty"`
	Error   *string `json:"error,omitempty"`
}
