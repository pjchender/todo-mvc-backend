package configs

type Auth struct {
	JWT
	Password
	Facebook
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

type Facebook struct {
	ClientID     string `default:"904503003705537"`
	ClientSecret string `default:"foobar"`
	AppToken     string `default:"foobar"`
}
