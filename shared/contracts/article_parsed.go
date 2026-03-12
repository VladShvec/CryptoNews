package contracts

type ArticleParsed struct {
	SourceID    string `json:"source_id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	PublishedAt string `json:"published_at"`
	Content     string `json:"content"`
}
