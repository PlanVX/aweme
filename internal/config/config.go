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

// Package config provides the configuration for the application
package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Config is the configuration for the application
type Config struct {
	// Release is the release mode of the application
	// if true, the application will run in release mode
	// zap logger will be used in production mode
	Release bool `yaml:"release"`
	JWT     struct {
		Secret    string   `yaml:"secret"`    // JwtSecret is the secret used to sign the JWT
		TTL       int64    `yaml:"ttl"`       // JwtTTL is the time to live for the JWT in seconds
		Whitelist []string `yaml:"whitelist"` // JwtWhitelist is the list of paths that can be accessed either with or without a JWT
	} `yaml:"jwt"` // JWT is the configuration for the JWT
	API struct {
		Prefix  string `yaml:"prefix"`  // ApiPrefix is the prefix for the API
		Address string `yaml:"address"` // ApiAddress is the address of the API server, like 0.0.0.0:8080
	} `yaml:"api"` // API is the configuration for the API
	MySQL struct {
		DSN      string   `yaml:"dsn"`      // MySQLDSN is the DSN of the MySQL database
		Replicas []string `yaml:"replicas"` // MySQLDSN is the replicas database dsn of the MySQL database
	} `yaml:"mysql"` // MySQL is the configuration for the MySQL database
	S3 struct {
		Endpoint  string `yaml:"endpoint"`  // S3Endpoint is the endpoint of the S3 bucket
		Bucket    string `yaml:"bucket"`    // S3Bucket is the bucket of the S3 bucket
		Region    string `yaml:"region"`    // S3Region is the region of the S3 bucket
		AccessKey string `yaml:"accessKey"` // S3AccessKey is the access key of the S3 bucket
		SecretKey string `yaml:"secretKey"` // S3SecretKey is the secret key of the S3 bucket
		Partition string `yaml:"partition"` // S3Partition is the partition of the S3 bucket
	} `yaml:"s3"` // S3 is the configuration for the S3 bucket
	Redis struct {
		Addr     []string `yaml:"address"`  // RedisAddr is the address of the Redis database
		Password string   `yaml:"password"` // RedisPassword is the password of the Redis database
		DB       int      `yaml:"db"`       // RedisDB is the database of the Redis database
	} `yaml:"redis"` // Redis is the configuration for the Redis database
	Otel struct {
		Enabled     bool   `yaml:"enabled"`     // OtelEnabled is the flag to enable OpenTelemetry
		Endpoint    string `yaml:"endpoint"`    // OtelEndpoint is the endpoint of the OpenTelemetry collector
		Service     string `yaml:"service"`     // OtelService is the service name
		Version     string `yaml:"version"`     // OtelVersion is the version of service
		Environment string `yaml:"environment"` // OtelEnvironment is the environment of service
	} `yaml:"otel"` // Otel is the configuration for the OpenTelemetry collector
}

// NewConfig returns a new Config
func NewConfig() (*Config, error) {
	path := "configs/config.yml" // default config file path
	if v := os.Getenv("CONFIG_FILE_PATH"); v != "" {
		path = v
	} // override config file path
	s, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c, err := parseConfig(s)
	return c, err
}

// parseConfig
// parses the binary config file data
// 1. expand environment variables in the config file
// 2. unmarshal the config file data to a Config struct
func parseConfig(s []byte) (*Config, error) {
	s = []byte(os.ExpandEnv(string(s))) // expand environment variables
	c := new(Config)
	err := yaml.Unmarshal(s, c) // unmarshal yaml to struct Config
	return c, err
}
