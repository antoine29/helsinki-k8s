package src

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	seed := time.Now().UnixNano()
	randomSource := rand.NewSource(seed)
	generator := rand.New(randomSource)
	randomString := ""
	for i := 0; i < length; i++ {
		// lowercase letters: 97 - 122 => 0+97 = 97, 25+97 = 122
		randomLowerCaseAsciInt := generator.Intn(26) + 97
		randomLowerCaseAsci := string(rune(randomLowerCaseAsciInt))
		randomString += randomLowerCaseAsci
	}

	return randomString
}

func ParseStrToInt(str string) int {
	intVal, parseError := strconv.ParseInt(str, 0, 0)
	if parseError != nil {
		fmt.Printf("Error: Cannot parse '%s' to int", str)
		os.Exit(3)
	}

	return int(intVal)
}

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

func CheckPassedParams(paramsDict map[string]string) {
	for _, expectedParam := range ExpectedParams {
		_, exists := paramsDict[expectedParam]
		if !exists {
			fmt.Printf("Error: Expected '%s' parameter. \n", expectedParam)
			os.Exit(3)
		}
	}
}
