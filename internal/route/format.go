package route

import (
	"io"

	"github.com/danielgtaylor/huma/v2"
	"gopkg.in/yaml.v3"
)

// Custom formatter for application/yaml content type

var DefaultYAMLFormat = huma.Format{
	Marshal: func(w io.Writer, v any) error {
		return yaml.NewEncoder(w).Encode(v)
	},
	Unmarshal: yaml.Unmarshal,
}

func init() {
	huma.DefaultFormats["application/yaml"] = DefaultYAMLFormat
}
