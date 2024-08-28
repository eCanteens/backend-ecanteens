package config

import "time"

type cors struct {
	AllowOrigin     []string
	AllowCredential bool
	AllowHeaders    []string
	AllowMethod     []string
}

type limiter struct {
	Rate  float64
	Burst int
}

type auth struct {
	AccessTokenExpiresIn  time.Duration
	RefreshTokenExpiresIn time.Duration
}

type config struct {
	MaxMultipartMemory int64
	Cors               *cors
	Limiter            *limiter
	Auth               *auth
}

var App = config{
	MaxMultipartMemory: 8 << 20, // 8 MiB
	Cors: &cors{
		AllowOrigin:     []string{"*"},
		AllowCredential: true,
		AllowHeaders:    []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		AllowMethod:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	},
	Limiter: &limiter{
		Rate:  20,
		Burst: 30,
	},
	Auth: &auth{
		AccessTokenExpiresIn:  time.Hour * 5,      // 5 hour
		RefreshTokenExpiresIn: time.Hour * 24 * 5, // 5 days
	},
}
