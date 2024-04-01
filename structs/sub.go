package structs

type Sub struct {
	Plan  string `json:"plan"`
	Price int    `json:"price"`
}

type Subscription struct {
	Sub []Sub `json:"subscription"`
}
