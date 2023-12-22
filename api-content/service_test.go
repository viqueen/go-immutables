package api_content_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	apicontent "github.com/viqueen/go-immutables/api-content"
)

type ContentPartial struct {
	title string                              `field:"required"`
	body  map[apicontent.ContentFormat]string `field:"required"`
}

func TestContentService(t *testing.T) {
	tests := []struct {
		name     string
		partials []ContentPartial
		err      error
	}{
		{
			name: "success: creates contents with unique titles",
			partials: []ContentPartial{
				{},
				{},
			},
		},
		{
			name: "error: fails to create contents with duplicate titles",
			partials: []ContentPartial{
				{},
				{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := apicontent.NewContentServiceBuilder().
				SetDataStore(make(map[string]apicontent.Content)).
				Build()
			var errs []error

			for _, input := range test.partials {
				_, err := service.CreateContent(
					*apicontent.NewCreateContentRequestBuilder().
						SetTitle(input.title).
						SetBody(input.body).
						Build(),
				)
				if err != nil {
					errs = append(errs, err)
				}
			}

			if test.err != nil {
				assert.Len(t, errs, 1)
				assert.Equal(t, test.err, errs[0])
			} else {
				assert.Len(t, errs, 0)
			}
		})
	}
}
