package formatters

import (
	"encoding/json"

	"code/models"
)

func FormatJSON(diff []*models.Node) (string, error) {
	result, err := json.MarshalIndent(diff, "", "  ")
	if err != nil {
		return "", err
	}
	return string(result), nil
}
