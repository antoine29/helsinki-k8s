package main

import (
	"strings"
)

var ExpectedParams = []string{"serverPort", "strLength", "secsInterval"}

func processDashSplitedParam(dashSplitedParam string) *string {
	spaceSplitedParams := strings.Split(strings.TrimSpace(dashSplitedParam), " ")
	if len(spaceSplitedParams) == 2 {
		return &spaceSplitedParams[1]
	}

	return nil
}

func BuildProgramParamsDict(params []string) map[string]string {
	join := strings.Join(params[:], " ")
	dashSplitedParams := strings.Split(join, "-")

	paramDict := make(map[string]string)

	for _, expectedParam := range ExpectedParams {
		for _, dashSplitedParam := range dashSplitedParams {
			_, present := paramDict[expectedParam]
			if strings.Contains(dashSplitedParam, expectedParam) && !present {
				spaceSplitedParamPointer := processDashSplitedParam(dashSplitedParam)
				if spaceSplitedParamPointer != nil {
					paramDict[expectedParam] = *spaceSplitedParamPointer
				}
			}
		}
	}

	return paramDict
}
