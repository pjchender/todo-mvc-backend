package configs

type Auth struct {
	JWT
	Password
}

type JWT struct {
	Issuer      string `default:"pjchender"`
	ExpireHours int    `default:"4"`
	Secret      string `default:"pjchender"`
}

type Password struct {
	Strength int    `default:"10"`
	Salt     string `default:"pjchender"`
}
