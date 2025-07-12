package replace

import (
	"regexp"
	"strings"
)

func Replace(str, old, new string) (string, error) {
	if len(old) >= 2 && old[0] == '/' && old[len(old)-1] == '/' {
		expr := old[1 : len(old)-1]
		re, err := regexp.Compile(expr)

		if err != nil {
			return "", err
		}

		return re.ReplaceAllString(str, new), nil
	} else {
		return strings.ReplaceAll(str, old, new), nil
	}
}

func MultiReplace(str string, newByOld map[string]string) (string, error) {
	for old, new := range newByOld {
		var err error

		if str, err = Replace(str, old, new); err != nil {
			return "", err
		}
	}

	return str, nil
}
