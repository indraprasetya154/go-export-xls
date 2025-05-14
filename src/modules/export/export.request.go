package export

type HeaderItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ExportRequest struct {
	Header  []HeaderItem             `json:"header"`
	Content []map[string]interface{} `json:"content"`
}
