package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `{
				"actionName": "Delay.A.Delay",
				"inputs": {
					"type": "object",
					"data": {
						"duration": {
							"type": "number",
							"data": 60000
						},
						"interval": {
							"type": "number",
							"data": 1
						},
						"unit": {
							"type": "singleSelector",
							"data": {
								"value": "",
								"label": "分"
							}
						}
					}
				}
			}`

	m := map[string]interface{}{}
	_ = json.Unmarshal([]byte(str), &m)
	fmt.Printf("%+v", m)
}
