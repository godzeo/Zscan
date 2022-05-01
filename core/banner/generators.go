package banner

import (
	"Zscan/core/httpx"
	"github.com/Knetic/govaluate"
	"strings"
)

// HelperFunctions contains the dsl functions
func HelperFunctions(resp *httpx.Response) (functions map[string]govaluate.ExpressionFunction) {
	functions = make(map[string]govaluate.ExpressionFunction)

	functions["title_contains"] = func(args ...interface{}) (interface{}, error) {
		pattern := strings.ToLower(toString(args[0]))
		title := strings.ToLower(resp.Title)
		return strings.Index(title, pattern) != -1, nil
	}

	functions["body_contains"] = func(args ...interface{}) (interface{}, error) {
		pattern := strings.ToLower(toString(args[0]))
		data := strings.ToLower(resp.DataStr)
		return strings.Index(data, pattern) != -1, nil
	}

	functions["protocol_contains"] = func(args ...interface{}) (interface{}, error) {
		return false, nil
	}

	functions["banner_contains"] = func(args ...interface{}) (interface{}, error) {
		return false, nil
	}

	functions["header_contains"] = func(args ...interface{}) (interface{}, error) {
		pattern := strings.ToLower(toString(args[0]))
		data := strings.ToLower(resp.HeaderStr)
		return strings.Index(data, pattern) != -1, nil
	}

	functions["server_contains"] = func(args ...interface{}) (interface{}, error) {
		pattern := strings.ToLower(toString(args[0]))
		server := resp.GetHeader("server")
		return strings.Index(server, pattern) != -1, nil
	}

	functions["cert_contains"] = func(args ...interface{}) (interface{}, error) {
		return false, nil
	}

	functions["port_contains"] = func(args ...interface{}) (interface{}, error) {
		return false, nil
	}

	return functions
}
