package rat_forwarder

import (
	"fmt"
	"os"
	"testing"
	"net/http"
)


func BenchmarkRatForwarder(t *testing.B) {
	for i := 0; i < 30; i++ {
		apiKey, apiSecret := fmt.Sprintf("rat_%v_api_key", i), fmt.Sprintf("rat_%v_api_secret", i)
		os.Setenv(apiKey, fmt.Sprintf("key%v", i))
		os.Setenv(apiSecret, fmt.Sprintf("value%v", i))
	}
	
	r := RatForwarder()

	for _, token := range *r {
		fmt.Printf("%v: %v\n", token[0], token[1])
		http.Get("https://www.google.com.tw/?hl=zh_TW")
	}
}
