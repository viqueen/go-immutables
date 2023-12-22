
package api_content

type ContentServiceBuilder interface {
    Build() *ContentService
    
    SetDataStore(value map[string]Content) ContentServiceBuilder
    
}

type contentserviceBuilder struct {
    target *ContentService
}

func NewContentServiceBuilder() ContentServiceBuilder {
    return &contentserviceBuilder{
        target: &ContentService{},
    }
}


func (b *contentserviceBuilder) SetDataStore(value map[string]Content) ContentServiceBuilder {
    b.target.dataStore = value
    return b
}


func (b *contentserviceBuilder) Build() *ContentService {
    return b.target
}
