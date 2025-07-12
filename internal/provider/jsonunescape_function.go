package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/winebarrel/terraform-provider-multireplace/internal/replace"
)

var (
	_ function.Function = JsonUnescape{}

	jsonUnescapeMapping = map[string]string{
		`\u003c`: "<",
		`\u003e`: ">",
		`\u0026`: "&",
	}
)

func NewJsonUnescape() function.Function {
	return JsonUnescape{}
}

type JsonUnescape struct{}

func (r JsonUnescape) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "jsonunescape"
}

func (r JsonUnescape) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "jsonunescape function",
		MarkdownDescription: "Replacing `\\u003c`, `\\u003e`, and `\\u0026` with `<`, `>`, `&`.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "source-string",
				MarkdownDescription: "Source string.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r JsonUnescape) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var srcStr string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &srcStr))

	if resp.Error != nil {
		return
	}

	dstStr, err := replace.MultiReplace(srcStr, jsonUnescapeMapping)

	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, dstStr))
}
