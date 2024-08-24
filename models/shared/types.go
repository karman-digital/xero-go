package sharedmodels

import "encoding/json"

type FloatOrString struct {
	FloatValue  float64
	StringValue string
	IsFloat     bool
}

func (f *FloatOrString) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		f.StringValue = str
		f.IsFloat = false
	} else {
		var floatValue float64
		if err := json.Unmarshal(data, &floatValue); err != nil {
			return err
		}
		f.FloatValue = floatValue
		f.IsFloat = true
	}
	return nil
}

func (f FloatOrString) MarshalJSON() ([]byte, error) {
	if f.IsFloat {
		return json.Marshal(f.FloatValue)
	}
	return json.Marshal(f.StringValue)
}

func NewFloatOrStringFloat(value float64) FloatOrString {
	return FloatOrString{
		FloatValue: value,
		IsFloat:    true,
	}
}

func NewFloatOrStringString(value string) FloatOrString {
	return FloatOrString{
		StringValue: value,
		IsFloat:     false,
	}
}
