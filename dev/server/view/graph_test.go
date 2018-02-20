// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package view

import (
	"testing"

	"github.com/google/shenzhen-go/dev/model"
)

type nopWriter struct{}

func (nopWriter) Write(r []byte) (int, error) { return len(r), nil }

func TestGraphEditorTemplate(t *testing.T) {
	// Smoke testing the editor template.
	for name, g := range model.TestGraphs {
		ei := &editorInput{
			Graph:     g,
			GraphJSON: `{"json": true}`,
			PartTypes: model.PartTypes,
		}

		if err := graphEditorTemplate.Execute(nopWriter{}, ei); err != nil {
			t.Errorf("graphEditorTemplate.Execute(%v) = %v, want nil error", name, err)
		}
	}
}
