package api_content

type ContentService struct {
	dataStore map[string]Content `field:"required"`
}

func (s *ContentService) CreateContent(request CreateContentRequest) (CreateContentResponse, error) {
	panic("implement me")
}
