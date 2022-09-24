package response

type Error struct {
	Message     string `json:"message,omitempty"`
	FailureCode string `json:"failure-code,omitempty"`
}
