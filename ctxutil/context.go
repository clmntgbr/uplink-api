package ctxutil

import (
	"context"

	"github.com/google/uuid"
)

const (
	ActiveProjectIDKey = "activeProjectID"
	UserIDKey          = "userID"
)

func WithActiveProjectID(ctx context.Context, activeProjectID *uuid.UUID) context.Context {
	return context.WithValue(ctx, ActiveProjectIDKey, activeProjectID)
}

func GetActiveProjectID(ctx context.Context) (*uuid.UUID, bool) {
	activeProjectID, ok := ctx.Value(ActiveProjectIDKey).(*uuid.UUID)
	return activeProjectID, ok
}

func WithUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(UserIDKey).(uuid.UUID)
	return userID, ok
}
