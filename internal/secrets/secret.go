package secrets

type Secret struct {
	Key   string
	Value string
}

type SecretList struct {
	Secrets []Secret
}
