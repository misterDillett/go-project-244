package formatters

import (
	"fmt"

	"code/models"
)

// Format выбирает нужный форматтер и применяет его к дереву различий
func Format(diff []*models.Node, format string) (string, error) {
	switch format {
	case "stylish":
		return FormatStylish(diff), nil
	case "plain":
		return FormatPlain(diff, ""), nil
	case "json":
		return FormatJSON(diff)
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}

// formatValue используется только для plain и json форматов
func formatValue(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case bool:
		if v {
			return "true"
		}
		return "false"
	case float64:
		return fmt.Sprintf("%v", v)
	case string:
		return fmt.Sprintf("'%s'", v)
	case map[string]interface{}, []interface{}:
		return "[complex value]"
	default:
		return fmt.Sprintf("%v", v)
	}
}
