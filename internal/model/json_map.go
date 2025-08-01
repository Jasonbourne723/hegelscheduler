package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONMap map[string]interface{}

func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONMap: %v", value)
	}
	return json.Unmarshal(bytes, &j)
}
