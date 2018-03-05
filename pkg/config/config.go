package config

import (
	"strconv"
)

//Config for API
type Config struct {
	port             int
	cookieLifetime   int
	oauthCallbackURL string
	allowedHeaders   []string
	allowedMethods   []string
	allowedOrigins   []string
}

//GetConfig returns a populated config object
func GetConfig(env string) *Config {
	switch env {
	case "prod":
		config := Config{
			port:             8080,
			cookieLifetime:   1000,
			oauthCallbackURL: "http://www.adventure-plan.com/api/auth/callback",
			allowedHeaders:   []string{"X-Requested-With", "Content-Type"},
			allowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
			allowedOrigins:   []string{"http://localhost:8100", "http://127.0.0.1:8100", "http://www.adventure-plan.com", "http://adventure-plan.com"},
		}
		return &config
	default:
		config := Config{
			port:             8080,
			cookieLifetime:   1000,
			oauthCallbackURL: "http://127.0.1/api/auth/callback",
			allowedHeaders:   []string{"X-Requested-With", "Content-Type"},
			allowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
			allowedOrigins:   []string{"*"},
		}
		return &config
	}
}

// GetPort returns the port
func (c *Config) GetPort() int {
	return c.port
}

// GetPortListenerStr returns the port as a listener string
func (c *Config) GetPortListenerStr() string {
	return ":" + strconv.Itoa(c.port)
}

//GetOAuthCallbackURL returns the URL to provide as a callback to oauth providers
func (c *Config) GetOAuthCallbackURL() string {
	return c.oauthCallbackURL
}

//GetAllowedHeaders returns the permissible HTTP Headers
func (c *Config) GetAllowedHeaders() []string {
	return c.allowedHeaders
}

//GetAllowedMethods returns the permissible HTTP methods
func (c *Config) GetAllowedMethods() []string {
	return c.allowedMethods
}

//GetAllowedOrigins returns the permissible HTTP Origins
func (c *Config) GetAllowedOrigins() []string {
	return c.allowedOrigins
}

//GetCookieLifetime returns how long the cookies should be stored for
func (c *Config) GetCookieLifetime() int {
	return c.cookieLifetime
}
