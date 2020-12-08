package riskcontrolapidemo_go

import (
	"log"
	"testing"
)





func TestIsAllow(t *testing.T) {

	Start(HOST, MERCHANT_KEY, func() {

		IsAllow("1", "2", PROJECT_PREFIX+"15999999999", func(code int, result string) {

			log.Println(code, result)
		})

		IsAllow("1", "2", PROJECT_PREFIX+"13811343201", func(code int, result string) {

			log.Println(code, result)
		})

		IsAllow("1", "2", PROJECT_PREFIX+"15988497532", func(code int, result string) {

			log.Println(code, result)
		})

	})

	select {}

}
