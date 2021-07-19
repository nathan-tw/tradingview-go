package rat_forwarder

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func RatForwarder() *map[int][]string{
	apiMatcher, _ := regexp.Compile("^rat_[0-9]+_api_key")
	secretMatcher, _ := regexp.Compile("^rat_[0-9]+_api_secret")
	ratPairs := make(map[int][]string)
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		if apiMatcher.MatchString(pair[0]) {
			idx, _ := strconv.Atoi(strings.Split(pair[0], "_")[1])
			if _, ok := ratPairs[idx]; !ok {
				ratPairs[idx] = make([]string, 2)
			}
			ratPairs[idx][0] = pair[1]
		}
		if secretMatcher.MatchString(pair[0]) {
			idx, _ := strconv.Atoi(strings.Split(pair[0], "_")[1])
			if _, ok := ratPairs[idx]; !ok {
				ratPairs[idx] = make([]string, 2)
			}
			ratPairs[idx][1] = pair[1]
		}
	}
	return &ratPairs
}