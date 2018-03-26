package notifications

import (
	"time"
)

type UserModel struct {
	ID        int
	userID    int
	name      string
	timestamp time.time
}
