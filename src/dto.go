package foggo

type Data struct {
	Id string `json:"id,omitempty"`

	Temperature float32 `json:"temperature,omitempty"`

	Timestamp int32 `json:"timestamp,omitempty"`
}
