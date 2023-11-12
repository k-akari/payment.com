package fixture

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
)

func User(u *domain.User) *domain.User {
	defaultID := rand.Int() //nolint:gosec
	defaultName := "name" + strconv.Itoa(defaultID)
	now := time.Now()
	result := &domain.User{
		ID:        domain.UserID(defaultID),
		Name:      defaultName,
		Email:     defaultName + "@example.com",
		Password:  "password",
		CreatedAt: now,
		UpdatedAt: now,
	}
	if u == nil {
		return result
	}
	if u.ID != 0 {
		result.ID = u.ID
	}
	if u.Name != "" {
		result.Name = u.Name
	}
	if u.Email != "" {
		result.Email = u.Email
	}
	if u.Password != "" {
		result.Password = u.Password
	}
	if !u.CreatedAt.IsZero() {
		result.CreatedAt = u.CreatedAt
	}
	if !u.UpdatedAt.IsZero() {
		result.UpdatedAt = u.UpdatedAt
	}

	return result
}
