// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package output

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/state"
	yaml "gopkg.in/yaml.v3"
)

// JSON outputs resources in JSON format.
type JSON struct {
	writer     io.Writer
	withEvents bool
}

// NewJSON initializes JSON resource output.
func NewJSON(writer io.Writer) *JSON {
	return &JSON{
		writer: writer,
	}
}

// WriteHeader implements output.Writer interface.
func (j *JSON) WriteHeader(_ *meta.ResourceDefinition, withEvents bool) error {
	j.withEvents = withEvents

	return nil
}

// prepareEncodableData prepares the data of a resource to be encoded as JSON and populates it with some extra information.
func (j *JSON) prepareEncodableData(r resource.Resource, event state.EventType) (map[string]any, error) {
	out, err := resource.MarshalYAML(r)
	if err != nil {
		return nil, err
	}

	yamlBytes, err := yaml.Marshal(out)
	if err != nil {
		return nil, err
	}

	var data map[string]any

	err = yaml.Unmarshal(yamlBytes, &data)
	if err != nil {
		return nil, err
	}

	if j.withEvents {
		data["event"] = strings.ToLower(event.String())
	}

	return data, nil
}

func writeAsIndentedJSON(wr io.Writer, data any) error {
	enc := json.NewEncoder(wr)
	enc.SetIndent("", "    ")

	return enc.Encode(data)
}

// WriteResource implements output.Writer interface.
func (j *JSON) WriteResource(r resource.Resource, event state.EventType) error {
	data, err := j.prepareEncodableData(r, event)
	if err != nil {
		return err
	}

	return writeAsIndentedJSON(j.writer, data)
}

// Flush implements output.Writer interface.
func (j *JSON) Flush() error {
	return nil
}
