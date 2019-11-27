package debug

import "os"
import "io/ioutil"
import "encoding/json"
import "log"

// Debug represents the data of the configuration file or a subobject of it
type Debug struct {
	data map[string]interface{}
}

func (d *Debug) get(key string) (bool, interface{}) {
	v, ok := d.data[key]
	return ok, v
}

// GetInt returns whether a specific key exists and its int value if applicable
func (d *Debug) GetInt(key string) (bool, int) {
	if ok, v := d.get(key); ok {
		if vn, tcok := v.(int); tcok {
			return true, vn
		}
	}
	return false, 0
}

// GetIntDefault returns whether a specific key exists and its int value if applicable and if not the given fallback value
func (d *Debug) GetIntDefault(key string, fallback int) int {
	if ok, v := d.GetInt(key); ok {
		return v
	} else {
		return fallback
	}
}

// GetBool returns whether a specific key exists and its bool value if applicable
func (d *Debug) GetBool(key string) (bool, bool) {
	if ok, v := d.get(key); ok {
		if vn, tcok := v.(bool); tcok {
			return true, vn
		}
	}
	return false, false
}

// GetBoolDefault returns whether a specific key exists and its bool value if applicable and if not the given fallback value
func (d *Debug) GetBoolDefault(key string, fallback bool) bool {
	if ok, v := d.GetBool(key); ok {
		return v
	} else {
		return fallback
	}
}

// GetString returns whether a specific key exists and its string value if applicable
func (d *Debug) GetString(key string) (bool, string) {
	if ok, v := d.get(key); ok {
		if vn, tcok := v.(string); tcok {
			return true, vn
		}
	}
	return false, ""
}

// GetStringDefault returns whether a specific key exists and its string value if applicable and if not the given fallback value
func (d *Debug) GetStringDefault(key string, fallback string) string {
	if ok, v := d.GetString(key); ok {
		return v
	} else {
		return fallback
	}
}

// GetSub returns whether a specific key exists and its Debug value if applicable
func (d *Debug) GetSub(key string) (bool, *Debug) {
	if ok, v := d.get(key); ok {
		if vn, tcok := v.(map[string]interface{}); tcok {
			return true, &Debug{data: vn}
		}
	}
	return false, nil
}

// GetSubDefault returns whether a specific key exists and its bool value if applicable and if not an empty Debug object
func (d *Debug) GetSubDefault(key string) *Debug {
	if ok, v := d.GetSub(key); ok {
		return v
	} else {
		return &Debug{}
	}
}

// NewDebug imports a json file to a Debug object
func NewDebug(path string) (*Debug, error) {
	d := new(Debug)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	if err := json.Unmarshal(byteValue, &d.data); err != nil {
		log.Fatal(err)
	}
	return d, nil
}
