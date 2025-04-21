package constants

import "time"

const (
    DefaultTimeout     = 10 * time.Second
    LongTimeout        = 30 * time.Second
    ShortTimeout       = 5 * time.Second
    DatabaseTimeout    = 10 * time.Second
    CacheTimeout      = 5 * time.Second
)