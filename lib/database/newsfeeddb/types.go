package newsfeeddb

type Feed struct {
	Id   int    `json:"id" xml:"id,attr"`
	Type string `json:"type" xml:"type"`
	Name string `json:"name" xml:"name"`
	URL  string `json:"url" xml:"url"`
}
