
package api_content

type CreateContentResponseBuilder interface {
    Build() *CreateContentResponse
    
    SetContent(value Content) CreateContentResponseBuilder
    
}

type createcontentresponseBuilder struct {
    target *CreateContentResponse
}

func NewCreateContentResponseBuilder() CreateContentResponseBuilder {
    return &createcontentresponseBuilder{
        target: &CreateContentResponse{},
    }
}


func (b *createcontentresponseBuilder) SetContent(value Content) CreateContentResponseBuilder {
    b.target.content = value
    return b
}


func (b *createcontentresponseBuilder) Build() *CreateContentResponse {
    return b.target
}
