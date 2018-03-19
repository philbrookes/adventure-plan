package config

import (
	"fmt"
	"strconv"
)

//MySQLConfig stores values for connecting to a MySQL server
type MySQLConfig struct {
	port   int
	host   string
	user   string
	pass   string
	dbname string
}

//Config for API
type Config struct {
	port             int
	cookieLifetime   int
	oauthCallbackURL string
	allowedHeaders   []string
	allowedMethods   []string
	allowedOrigins   []string
	mysqlConfig      *MySQLConfig
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
			mysqlConfig: &MySQLConfig{
				port:   3306,
				host:   "localhost",
				user:   "root",
				pass:   "!}@C.G{Py?tu97nN",
				dbname: "adventureplan",
			},
		}
		return &config
	default:
		config := Config{
			port:             8080,
			cookieLifetime:   1000,
			oauthCallbackURL: "http://127.0.0.1/api/auth/callback",
			allowedHeaders:   []string{"X-Requested-With", "Content-Type"},
			allowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
			allowedOrigins:   []string{"*"},
			mysqlConfig: &MySQLConfig{
				port:   3306,
				host:   "localhost",
				user:   "root",
				pass:   "!}@C.G{Py?tu97nN",
				dbname: "adventureplan",
			},
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

//GetMySQLConfig returns a handle to the MySQL config object
func (c *Config) GetMySQLConfig() *MySQLConfig {
	return c.mysqlConfig
}

//GetPort returns the port MySQL is running on
func (m *MySQLConfig) GetPort() int {
	return m.port
}

//GetHost returns the host MySQL is running on
func (m *MySQLConfig) GetHost() string {
	return m.host
}

//GetUser returns the user to connect to MySQL with
func (m *MySQLConfig) GetUser() string {
	return m.user
}

//GetPass returns the password to connect to MySQL with
func (m *MySQLConfig) GetPass() string {
	return m.pass
}

//GetDBName returns the dbname for the MySQL connection
func (m *MySQLConfig) GetDBName() string {
	return m.dbname
}

//GetConnectionString returns a MySQL connection string
func (m *MySQLConfig) GetConnectionString() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/", m.user, m.pass, m.host, strconv.Itoa(m.port))
}
