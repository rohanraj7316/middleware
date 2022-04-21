package configs

import "time"

type Time struct {
	Format string
}

var TimeDefault = Time{
	Format: time.RFC3339,
}
