/*
 * Copyright (c) 2023 The PlanVX Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logic

import (
	"context"
	"github.com/PlanVX/aweme/internal/config"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"time"
)

// ContextKey is the key of user id in context
const ContextKey = "owner_id"

// customClaims are custom claims extending default ones.
type customClaims struct {
	Name   string `json:"name"`
	UserID int64  `json:"user_id"`
	jwt.RegisteredClaims
}

func genToken(username string, id int64, duration time.Duration) *jwt.Token {
	// 构造 claims
	claims := &customClaims{
		username,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	// Create token with claims
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

// JWTSigner jwt 签名器
type JWTSigner struct {
	secret    []byte
	duration  time.Duration
	whiteList map[string]struct{}
}

// NewJWTSigner 生成 jwt 签名器
func NewJWTSigner(config *config.Config) *JWTSigner {
	j := &JWTSigner{
		secret:    []byte(config.JWT.Secret),
		duration:  time.Duration(config.JWT.TTL) * time.Second,
		whiteList: make(map[string]struct{}),
	}
	for _, path := range config.JWT.Whitelist {
		j.whiteList[path] = struct{}{}
	}
	return j
}

// NewJWTMiddleware 生成 jwt 中间件
func (j *JWTSigner) NewJWTMiddleware() echo.MiddlewareFunc {
	conf := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(customClaims)
		},
		SuccessHandler: func(c echo.Context) {
			// write owner id to context, so that we can use it
			claims := c.Get("user").(*jwt.Token).Claims.(*customClaims)
			ctx := context.WithValue(c.Request().Context(), ContextKey, claims.UserID)
			c.SetRequest(c.Request().WithContext(ctx))
		},
		SigningKey: j.secret,
		ErrorHandler: func(c echo.Context, err error) error {
			if _, ok := j.whiteList[c.Path()]; !ok {
				return echo.ErrUnauthorized
			} // 如果不在白名单中，返回错误
			t, err := j.genSignedToken("guest", 0) // 生成默认的 token
			if err != nil {
				return err
			}
			c.Set("user", t) // 设置默认的 token
			return nil
		},
		ContinueOnIgnoredError: true,
		TokenLookup:            "query:token,form:token,header:Authorization:Bearer",
	}
	return echojwt.WithConfig(conf)
}

// AddWhiteList 添加白名单列表
func (j *JWTSigner) AddWhiteList(paths []string) {
	for _, path := range paths {
		j.whiteList[path] = struct{}{}
	}
}

// 签发 jwt token
func (j *JWTSigner) genSignedToken(username string, id int64) (string, error) {
	// 构造 jwt token
	token := genToken(username, id, j.duration)
	// Generate encoded token
	return token.SignedString(j.secret)
}
