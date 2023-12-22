//go:generate go run ../cmd/builders.go
package api_content

type Service interface {
	CreateContent(request CreateContentRequest) (*CreateContentResponse, error)
}

type ContentFormat string

const (
	EditorV1 ContentFormat = "editor_v1"
	EditorV2 ContentFormat = "editor_v2"
)

type Content struct {
	title   string                   `field:"required"`
	body    map[ContentFormat]string `field:"required"`
	summary string
}

type CreateContentRequest struct {
	title   string                   `field:"required"`
	body    map[ContentFormat]string `field:"required"`
	summary string
}

type CreateContentResponse struct {
	content Content
}
