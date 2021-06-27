package string

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	//[57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,57,99,50,113,102,114,99,110,53,103,101,114,107,113,97,100,104,50,101,48,48]

	str := "2021-06-24 10:28:41.888000+07:00 info|grpcMethod:node-agent.worker|traceId:|uid:|file:|func:|node_report_consumer_error: consumerGroup error: Error: Not a message set. Magic byte is 2\\n    at /data/server/node_modules/kafka-node/lib/protocol/protocol.js:280:25\\n    at /data/server/node_modules/kafka-node/node_modules/async/dist/async.js:3880:24\\n    at replenish (/data/server/node_modules/kafka-node/node_modules/async/dist/async.js:1011:17)\\n    at /data/server/node_modules/kafka-node/node_modules/async/dist/async.js:1016:9\\n    at eachOfLimit (/data/server/node_modules/kafka-node/node_modules/async/dist/async.js:1041:24)\\n    at /data/server/node_modules/kafka-node/node_modules/async/dist/async.js:1046:16\\n    at _parallel (/data/server/node_modules/kafka-node/node_modules/async/dist/async.js:3879:5)\\n    at Object.series (/data/server/node_modules/kafka-node/node_modules/async/dist/async.js:4735:5)\\n    at _decodeFetchResponse (/data/server/node_modules/kafka-node/lib/protocol/protocol.js:244:9)\\n    at /data/server/node_modules/kafka-node/lib/protocol/protocol.js:199:12 {\\n  topic: 'ads_apa_test_v100',\\n  partition: 3\\n} Error: Not a message set. Magic byte is 2\\n    at /data/server/node_modules/kafka-node/lib/protocol/protocol.js:280:25\\n    at /data/server/node_modules/kafka-node/node_modules/async/dist/async.js:3880:24\\n    at replenish (/data/server/node_modules/kafka-node/node_modules/async/dist/async.js:1011:17)\\n    at /data/server/node_modules/kafka-node/node_modules/async/dist/async.js:1016:9\\n    at eachOfLimit (/data/server/node_modules/kafka-node/node_modules/async/dist/async.js:1041:24)\\n    at /data/server/node_modules/kafka-node/node_modules/async/dist/async.js:1046:16\\n    at _parallel (/data/server/node_modules/kafka-node/node_modules/async/dist/async.js:3879:5)\\n    at Object.series (/data/server/node_modules/kafka-node/node_modules/async/dist/async.js:4735:5)\\n    at _decodeFetchResponse (/data/server/node_modules/kafka-node/lib/protocol/protocol.js:244:9)\\n    at /data/server/node_modules/kafka-node/lib/protocol/protocol.js:199:12\\n"
	str2 := "2021-06-20 00:00:03.402873+07:00 info|grpcMethod:|traceId:50b44b699f5b0678|uid:|file:kyc_uplift_config.go/166|func:FillCache"
	fmt.Println(len(str))
	fmt.Println(len(str2))

	//fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))
	//fmt.Println(base64.StdEncoding.EncodeToString([]byte(str2)))

	//fmt.Println([]byte(str))
	//
	//fmt.Println(len(str[:16]))
	//fmt.Println(len(str[16:]))
}
