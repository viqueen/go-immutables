//go:generate go run ../cmd/builders.go
package api_content

import "strings"

type ContentService struct {
	dataStore map[string]Content `field:"required"`
}

func (s *ContentService) CreateContent(request CreateContentRequest) (CreateContentResponse, error) {
	key := strings.ToLower(request.title)
	content := NewContentBuilder().
		SetTitle(request.title).
		SetBody(request.body).
		SetSummary(request.summary).
		Build()

	s.dataStore[key] = content
	response := NewCreateContentResponseBuilder().
		SetContent(content).
		Build()
	return response, nil
}
