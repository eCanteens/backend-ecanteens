package config

type Cors struct {
	AllowOrigin     []string
	AllowCredential bool
	AllowHeaders    []string
	AllowMethod     []string
}

type Limiter struct {
	Rate  float64
	Burst int
}

type Config struct {
	MaxMultipartMemory int64
	Cors               *Cors
	Limiter            *Limiter
}

var App = Config{
	MaxMultipartMemory: 8 << 20, // 8 MiB
	Cors: &Cors{
		AllowOrigin:     []string{"*"},
		AllowCredential: true,
		AllowHeaders:    []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		AllowMethod:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	},
	Limiter: &Limiter{
		Rate:  20,
		Burst: 30,
	},
}
