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
	"github.com/google/shenzhen-go/dev/jsutil"
)

// Diagram wraps the top SVG element and holds references to
// important parts of it.
type Diagram struct {
	jsutil.Element // the SVG element
	*View

	dragItem     draggable  // nil if nothing is being dragged
	selectedItem selectable // nil if nothing is selected
	errLabel     *textBox
}

func (d *Diagram) cursorPos(e jsutil.Object) (x, y float64) {
	bcr := d.Call("getBoundingClientRect")
	x = e.Get("clientX").Float() - bcr.Get("left").Float()
	y = e.Get("clientY").Float() - bcr.Get("top").Float()
	return
}

func (d *Diagram) mouseDown(e jsutil.Object) {
	if d.selectedItem == nil {
		return
	}
	d.selectedItem.loseFocus(e)
	d.View.ShowRHSPanel(d.View.GraphPropertiesPanel)
	e.Call("stopPropagation")
}

func (d *Diagram) mouseMove(e jsutil.Object) {
	if d.dragItem == nil {
		return
	}
	d.dragItem.drag(e)
	e.Call("stopPropagation")
}

func (d *Diagram) mouseUp(e jsutil.Object) {
	if d.dragItem == nil {
		return
	}
	d.dragItem.drag(e)
	d.dragItem.drop(e)
	d.dragItem = nil
	e.Call("stopPropagation")
}

// selecter makes an onclick handler for a selectable.
func (d *Diagram) selecter(s selectable) func(jsutil.Object) {
	return func(e jsutil.Object) {
		if d.selectedItem != nil {
			d.selectedItem.loseFocus(e)
		}
		d.selectedItem = s
		s.gainFocus(e)
		e.Call("stopPropagation")
	}
}

func (d *Diagram) saveSelected(e jsutil.Object) {
	if d.selectedItem == nil {
		return
	}
	d.selectedItem.save(e)
}

func (d *Diagram) deleteSelected(e jsutil.Object) {
	if d.selectedItem == nil {
		return
	}
	d.selectedItem.delete(e)
}

func (d *Diagram) setError(err string, x, y float64) {
	if err == "" {
		d.clearError()
		return
	}
	d.AddChildren(d.errLabel.moveTo(x, y).setText(err).show()) // Re-add = Bring to front
}

func (d *Diagram) clearError() {
	d.errLabel.hide()
}

// Point is anything that has a position on the canvas.
type Point interface {
	Pt() (x, y float64)
}

type ephemeral struct{ x, y float64 }

func (e ephemeral) Pt() (x, y float64) { return e.x, e.y }

type draggable interface {
	drag(jsutil.Object)
	drop(jsutil.Object)
}

type selectable interface {
	gainFocus(jsutil.Object)
	loseFocus(jsutil.Object)
	delete(jsutil.Object)
	save(jsutil.Object)
}
