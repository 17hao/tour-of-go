package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "\n*Summary*: Operation Lantency is abnormal(\u003e20s), current: 25.38265714285714s. `warning`\n*Labels*: \n• host: `ln75.para.bscc`\n\n\n\n*Summary*: Operation Lantency is abnormal(\u003e20s), current: 24.949194444444444s. `warning`\n*Labels*: \n• host: `ln76.para.bscc`\n\n\n\n*Summary*: Operation Lantency is abnormal(\u003e20s), current: 37.077969696969696s. `warning`\n*Labels*: \n• host: `monopoly-ln02`\n\n\n\n*Summary*: Operation Lantency is abnormal(\u003e20s), current: 40.303727272727265s. `warning`\n*Labels*: \n• host: `monopoly-ln03`\n\n\n\n"
	res := strconv.QuoteToASCII(str)
	fmt.Printf("%s", res)
}
