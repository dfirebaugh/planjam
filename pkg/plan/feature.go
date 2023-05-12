package plan

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/dfirebaugh/planjam/pkg/strfmt"
	"gopkg.in/yaml.v3"
)

type FeatureReference struct {
	ID    string
	Label string
}

func (r FeatureReference) Feature() Feature {
	return Feature{FeatureReference: r}
}

type Field struct {
	Label string `yaml:"label"`
	Value any    `yaml:"value"`
}

type Feature struct {
	FeatureReference `yaml:"ref"`
	Fields           []Field
}

func GetFeature(data []byte) (Feature, error) {
	var feature Feature
	err := yaml.Unmarshal(data, &feature)
	return feature, err
}

func (f Feature) Write(w io.Writer) error {
	raw, err := yaml.Marshal(f)
	if err != nil {
		return err
	}

	_, err = w.Write(raw)
	return err
}

func (f Feature) String() string {
	table := strfmt.NewTable()

	table.AppendToColumn(0, f.Label)

	if len(f.Fields) == 0 && len(f.Label) == 0 {
		table.AppendToColumn(0, "empty")
		return table.String()
	}
	for _, f := range f.Fields {
		table.AppendToColumn(0, fmt.Sprintf("    %s: %v", f.Label, f.Value))
	}

	return table.String()
}

func (f Feature) FileName(dir string) string {
	return filepath.Join(dir, f.ID+"_"+f.Label+".yaml")
}

func (f *Feature) AddField(label string, value any) {
	f.Fields = append(f.Fields, Field{
		Label: label,
		Value: value,
	})
}

func (f *Feature) RemoveField(label string) {
	var fields []Field
	for _, field := range f.Fields {
		if field.Label == label {
			continue
		}
		fields = append(fields, field)
	}
	f.Fields = fields
}
