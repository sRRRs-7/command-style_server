package cfg

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	conf, err := LoadConfig("../")
	require.NoError(t, err)
	require.NotEmpty(t, conf)

	err = godotenv.Load("../app.env")
	require.NoError(t, err)

	v := os.Getenv("DB_DRIVER")
	require.Equal(t, conf.DBdriver, v)

	v = os.Getenv("DB_SOURCE")
	require.Equal(t, conf.DBsource, v)

	v = os.Getenv("HTTP_SERVER_ADDRESS")
	require.Equal(t, conf.HttpServerAddress, v)

	v = os.Getenv("GRPC_SERVER_ADDRESS")
	require.Equal(t, conf.GrpcServerAddress, v)

	v = os.Getenv("TOKEN_SYMMETRIC_KEY")
	require.Equal(t, conf.TokenSymmetricKey, v)

	v = os.Getenv("ACCESS_TOKEN_DURATION")
	m, err := time.ParseDuration(v)
	require.NoError(t, err)
	require.Equal(t, int64(conf.AccessTokenDuration), m.Nanoseconds())

	v = os.Getenv("REFRESH_TOKEN_DURATION")
	m, err = time.ParseDuration(v)
	require.NoError(t, err)
	require.Equal(t, int64(conf.RefreshTokenDuration), m.Nanoseconds())

	v = os.Getenv("GIN_CONTEXT_KEY")
	require.Equal(t, conf.GinContextKey, v)

	v = os.Getenv("REDIS_COOKIE_KEY")
	require.Equal(t, conf.RedisCookieKey, v)

	v = os.Getenv("REDIS_COOKIE_ADMIN_KEY")
	require.Equal(t, conf.AdminCookieKey, v)

	v = os.Getenv("ACCESS_COOKIE_DURATION")
	require.Equal(t, string(fmt.Sprint(conf.AccessCookieDuration)), v)

	v = os.Getenv("ACCESS_REDIS_DURATION")
	m, err = time.ParseDuration(v)
	require.NoError(t, err)
	require.Equal(t, int64(conf.AccessRedisDuration), m.Nanoseconds())

	v = os.Getenv("TRANSLATE_ENDPOINT")
	require.Equal(t, conf.TranslateEndpoint, v)
}
