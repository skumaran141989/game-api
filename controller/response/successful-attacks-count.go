package response

type SimpleSuccefulAttacksCount struct {
	StartTime int64 `json:"start_time,omitempty"`
	Count     int64 `json:"count,omitempty"`
	EndTime   int64 `json:"end_time,omitempty"`
}
