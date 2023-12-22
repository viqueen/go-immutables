
package api_content

type ContentBuilder interface {
    Build() *Content
    
    SetId(value string) ContentBuilder
    
    SetTitle(value string) ContentBuilder
    
    SetBody(value map[ContentFormat]string) ContentBuilder
    
    SetSummary(value string) ContentBuilder
    
}

type contentBuilder struct {
    target *Content
}

func NewContentBuilder() ContentBuilder {
    return &contentBuilder{
        target: &Content{},
    }
}


func (b *contentBuilder) SetId(value string) ContentBuilder {
    b.target.id = value
    return b
}

func (b *contentBuilder) SetTitle(value string) ContentBuilder {
    b.target.title = value
    return b
}

func (b *contentBuilder) SetBody(value map[ContentFormat]string) ContentBuilder {
    b.target.body = value
    return b
}

func (b *contentBuilder) SetSummary(value string) ContentBuilder {
    b.target.summary = value
    return b
}


func (b *contentBuilder) Build() *Content {
    return b.target
}
