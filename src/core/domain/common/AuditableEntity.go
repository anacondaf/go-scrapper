package domain

import "time"

type AuditableEntity struct {
	CreatedAt  time.Time
	ModifiedAt time.Time
}
