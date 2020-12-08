package riskcontrolapidemo_go

import (
	"encoding/json"
	"net/http"

	"github.com/luaxlou/wsclient"
)

var wsc *wsclient.WsClient

func Start(addr string, merchantKey string, onReady func()) {

	wsc = wsclient.New(addr, http.Header{
		"MERCHANT_KEY": []string{
			merchantKey,
		},
	})

	wsc.Ready = onReady
}

type IsAllowReq struct {
	CallId string `json:"call_id"`
	Caller string `json:"caller"`
	Callee string `json:"callee"`
}

func IsAllow(callId string, caller string, callee string, cb func(code int, result string)) {

	wsc.SendRequest("allow", IsAllowReq{
		CallId: callId,
		Caller: caller,
		Callee: callee,
	}, func(code int, data []byte) {

		if code != 200 {

			cb(code, "time out")
			return
		}

		var r wsclient.WsResult

		json.Unmarshal(data, &r)

		cb(r.Status, r.Msg)

	})

}
