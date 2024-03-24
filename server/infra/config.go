package infra

import "github.com/spf13/viper"

func LoadDefaultConfig() {
	viper.SetDefault("sqlite.filepath", "database.db")
	viper.SetDefault("auth.cookie_name", "sh_account_token")

	viper.SetDefault("auth.jwt.issuer", "sh-auth")
	viper.SetDefault("auth.jwt.audience", "sh-auth")
	viper.SetDefault("auth.jwt.expirySeconds", 259200)
	viper.SetDefault("auth.jwt.private_key_path", ".keys/ecdsa-public.pem")
	viper.SetDefault("auth.jwt.public_key_path", ".keys/ecdsa-private.pem")

	viper.AutomaticEnv()
}
