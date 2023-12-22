
package api_content

type CreateContentRequestBuilder interface {
    Build() *CreateContentRequest
    
    SetTitle(value string) CreateContentRequestBuilder
    
    SetBody(value map[ContentFormat]string) CreateContentRequestBuilder
    
    SetSummary(value string) CreateContentRequestBuilder
    
}

type createcontentrequestBuilder struct {
    target *CreateContentRequest
}

func NewCreateContentRequestBuilder() CreateContentRequestBuilder {
    return &createcontentrequestBuilder{
        target: &CreateContentRequest{},
    }
}


func (b *createcontentrequestBuilder) SetTitle(value string) CreateContentRequestBuilder {
    b.target.title = value
    return b
}

func (b *createcontentrequestBuilder) SetBody(value map[ContentFormat]string) CreateContentRequestBuilder {
    b.target.body = value
    return b
}

func (b *createcontentrequestBuilder) SetSummary(value string) CreateContentRequestBuilder {
    b.target.summary = value
    return b
}


func (b *createcontentrequestBuilder) Build() *CreateContentRequest {
    return b.target
}
