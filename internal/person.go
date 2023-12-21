package internal

import "time"

//go:generate reform

//reform:people
type Person struct {
	ID        int32      `reform:"id,pk"`
	Name      string     `reform:"name"`
	Email     *string    `reform:"email"`
	CreatedAt time.Time  `reform:"created_at"`
	UpdatedAt *time.Time `reform:"updated_at"`
}
