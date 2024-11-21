package utils

import (
	"SecretCare/entity"
	"context"
)

type ContextKey string

const UserContextKey ContextKey = "user"

func SetUserInContext(ctx context.Context, user *entity.Users) context.Context {
	return context.WithValue(ctx, UserContextKey, user)
}

func GetUserFromContext(ctx context.Context) (*entity.Users, bool) {
	user, ok := ctx.Value(UserContextKey).(*entity.Users)
	return user, ok
}
