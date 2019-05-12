package resolvers

import (
	"context"

	"github.com/shobhitsharma/go-gql-starter/handler"
	"github.com/shobhitsharma/go-gql-starter/model"
)

// ChangePassword mutation change password
func (r *Resolvers) ChangePassword(ctx context.Context, args changePasswordMutationArgs) (*ChangePasswordResponse, error) {
	userID := ctx.Value(handler.ContextKey("userID"))

	if userID == nil {
		msg := "Not Authorized"
		return &ChangePasswordResponse{Status: false, Msg: &msg, User: nil}, nil
	}
	user := model.User{}

	if err := r.DB.First(&user, userID).Error; err != nil {
		msg := "Not existing user"
		return &ChangePasswordResponse{Status: false, Msg: &msg, User: nil}, nil
	}

	user.Password = args.Password
	user.HashPassword()

	r.DB.Save(&user)
	return &ChangePasswordResponse{Status: true, Msg: nil, User: &UserResponse{u: &user}}, nil
}

type changePasswordMutationArgs struct {
	Password string
}

// ChangePasswordResponse is the response type
type ChangePasswordResponse struct {
	Status bool
	Msg    *string
	User   *UserResponse
}

// Ok for ChangePasswordResponse
func (r *ChangePasswordResponse) Ok() bool {
	return r.Status
}

// Error for ChangePasswordResponse
func (r *ChangePasswordResponse) Error() *string {
	return r.Msg
}
