// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package htmltests

import (
	"testing"

	"github.com/googlecodelabs/tools/claat/proto/constructor"
	"github.com/googlecodelabs/tools/claat/proto/renderer/html"
	"github.com/googlecodelabs/tools/claat/proto/renderer/testing-utils"
	"github.com/googlecodelabs/tools/third_party"
)

func TestRenderTestBlockTemplate(t *testing.T) {
	tests := []*testingutils.CanonicalRenderingBatch{
		{
			InProto: &tutorial.Paragraph{},
			Out:     "",
			Ok:      false,
		},
		{
			InProto: protoconstructors.NewParagraph(),
			Out:     "",
			Ok:      false,
		},
	}
	testingutils.TestCanonicalRendererBatch(html.Render, tests, t)
}

func TestRenderTestBlockTemplateFromFile(t *testing.T) {
	paragraphProto := protoconstructors.NewParagraph(
		protoconstructors.NewInlineContentTextPlain(`hello, `),
		protoconstructors.NewInlineContentTextStrong(`world!`),
		protoconstructors.NewInlineContentTextEmphasized(` goodbye`),
		protoconstructors.NewInlineContentTextPlain(` `),
		protoconstructors.NewInlineContentTextStrongAndEmphasized(`cruel `),
		protoconstructors.NewInlineContentCode(`world!`),
	)

	tests := []*testingutils.CanonicalFileRenderingBatch{
		{
			InProto: paragraphProto,
			OutPath: "google_weather.txt",
			Ok:      true,
		},
	}
	testingutils.TestCanonicalFileRenderBatch(html.Render, tests, t)
}
