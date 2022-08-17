package config

// TODO should create the file in the test and delete it
// Otherwise this will not work in CI
// func TestConfigLoadFromFile(t *testing.T) {
// 	config, err := LoadConfig("../")
// 	if err != nil {
// 		t.Fatal("cannot load config", err)
// 	}
// 	require.Equal(t, config.PostgresDB, "simple-bank")
// 	require.Equal(t, config.PostgresHost, "localhost")
// 	require.Equal(t, config.PostgresUser, "simple-bank-user")
// 	require.Equal(t, config.PostgresPort, "5432")
// 	require.Equal(t, config.PostgresSSLMode, "disable")
// 	// TODO this does not actually work, I need to fix
// 	require.Equal(t, time.Minute*15, config.AccessTokenDuration)
// 	require.NotEmpty(t, config.PostgresPassword)  // TODO
// 	require.NotEmpty(t, config.TokenSymmetricKey) // TODO
// }

// TODO this doesn't work in ci
// func TestConfigLoadFromEnv(t *testing.T) {
// 	os.Setenv("POSTGRES_DB", "simple-bank-env")
// 	os.Setenv("POSTGRES_HOST", "simple-bank")
// 	os.Setenv("POSTGRES_USER", "simple-bank-user-env")
// 	os.Setenv("POSTGRES_PORT", "5433")
// 	os.Setenv("POSTGRES_SSL_MODE", "enable")
// 	os.Setenv("ACCESS_TOKEN_DURATION", "20m")
// 	os.Setenv("POSTGRES_PASSWORD", "password")
// 	os.Setenv("TOKEN_SYMMETRIC_KEY", "token")

// 	config, err := LoadConfig("../")
// 	if err != nil {
// 		t.Fatal("cannot load config", err)
// 	}
// 	require.Equal(t, config.PostgresDB, "simple-bank-env")
// 	require.Equal(t, config.PostgresHost, "simple-bank")
// 	require.Equal(t, config.PostgresUser, "simple-bank-user-env")
// 	require.Equal(t, config.PostgresPort, "5433")
// 	require.Equal(t, config.PostgresSSLMode, "enable")
// 	// TODO this does not actually work, I need to fix
// 	require.Equal(t, time.Minute*20, config.AccessTokenDuration)
// 	require.Equal(t, config.PostgresPassword, "password") // TODO
// 	require.Equal(t, config.TokenSymmetricKey, "token")   // TODO
// }
