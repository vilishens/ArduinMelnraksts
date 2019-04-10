package pointConfig

import (
	"fmt"
	"strings"
	"time"
)

func intervalCfgSecs(str string) (t time.Duration, err error) {

	parts := strings.Split(str, ":")

	fmt.Printf("Bortich \"%s\" ---> %v\n", str, parts)

	return
}
