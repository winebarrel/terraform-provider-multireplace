package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/winebarrel/terraform-provider-multireplace/internal/replace"
)

var (
	_ function.Function = MultiRepaceFunction{}
)

func NewMultiRepaceFunction() function.Function {
	return MultiRepaceFunction{}
}

type MultiRepaceFunction struct{}

func (r MultiRepaceFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "multireplace"
}

func (r MultiRepaceFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "multireplace function",
		MarkdownDescription: "Replace multiple substrings of a string.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "source-string",
				MarkdownDescription: "Source string.",
			},
		},
		VariadicParameter: function.MapParameter{
			Name:        "replacement-by-substring",
			ElementType: types.StringType,
			Description: "Mapping of replacement substrings. If the key is wrapped in forward slashes, it is treated as a regular expression.",
		},
		Return: function.StringReturn{},
	}
}

func (r MultiRepaceFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var srcStr string
	var replBySubstrs []map[string]string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &srcStr, &replBySubstrs))

	if len(replBySubstrs) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "wrong number of arguments (given 0, expected 1..)"))
	}

	if resp.Error != nil {
		return
	}

	for _, newByOld := range replBySubstrs {
		var err error

		if srcStr, err = replace.MultiReplace(srcStr, newByOld); err != nil {
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
			return
		}
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, srcStr))
}
