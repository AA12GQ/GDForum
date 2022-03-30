package requests

type Check404Input struct {
	U []string `json:"u,omitempty"`
	K []string `json:"k"`
}

type ResponseCheck struct {
	Check404Input
}

