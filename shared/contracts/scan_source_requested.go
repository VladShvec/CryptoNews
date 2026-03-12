package contracts

type ScanSourceRequested struct {
	SourceID    string `json:"source_id"`
	SourceType  string `json:"source_type"`
	URL         string `json:"url"`
	RequestedAt string `json:"requested_at"`
}
