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
	qtyMatcher, _ := regexp.Compile("^rat_[0-9]+_qty")
	ratPairs := make(map[int][]string)
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		if apiMatcher.MatchString(pair[0]) {
			idx, _ := strconv.Atoi(strings.Split(pair[0], "_")[1])
			if _, ok := ratPairs[idx]; !ok {
				ratPairs[idx] = make([]string, 3)
			}
			ratPairs[idx][0] = pair[1]
			continue
		}
		if secretMatcher.MatchString(pair[0]) {
			idx, _ := strconv.Atoi(strings.Split(pair[0], "_")[1])
			if _, ok := ratPairs[idx]; !ok {
				ratPairs[idx] = make([]string, 3)
			}
			ratPairs[idx][1] = pair[1]
			continue
		}
		if qtyMatcher.MatchString(pair[0]) {
			idx, _ := strconv.Atoi(strings.Split(pair[0], "_")[1])
			if _, ok := ratPairs[idx]; !ok {
				ratPairs[idx] = make([]string, 3)
			}
			ratPairs[idx][2] = pair[1]
		}
	}
	return &ratPairs
}