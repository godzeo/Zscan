package banner

import "fmt"

func toString(v interface{}) string {
	return fmt.Sprint(v)
}
func toInt(v interface{}) int {
	return int(v.(float64))
}
