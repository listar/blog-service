package config

import "time"

type JwtConfig struct {
	SECRET string
	EXP    time.Duration // 过期时间
	ALG    string        // 算法
}

func GetJwtConfig() *JwtConfig {
	return &JwtConfig{
		SECRET: GetEnv().AppSecret,
		EXP:    time.Hour * 24 * 7,
		ALG:    "HS256",
	}
}
