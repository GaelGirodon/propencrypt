package env

const (
	// prefix is the prefix of all the application environment variables.
	prefix = "PROPENCRYPT_"

	// Key is the name of the env var providing the encryption key.
	Key = prefix + "KEY"

	// Pattern is the name of the env var providing the sensitive property pattern.
	Pattern = prefix + "PATTERN"

	// Ext is the name of the env var providing the filename extension to add / remove.
	Ext = prefix + "EXT"
)
