package sharedmodels

import "encoding/json"

type IntOrString struct {
	IntValue    int
	StringValue string
	IsInt       bool
}

func (i *IntOrString) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		i.StringValue = str
		i.IsInt = false
	} else {
		var intValue int
		if err := json.Unmarshal(data, &intValue); err != nil {
			return err
		}
		i.IntValue = intValue
		i.IsInt = true
	}
	return nil
}

func (i IntOrString) MarshalJSON() ([]byte, error) {
	if i.IsInt {
		return json.Marshal(i.IntValue)
	}
	return json.Marshal(i.StringValue)
}

func NewIntString(value int) IntOrString {
	return IntOrString{
		IntValue: value,
		IsInt:    true,
	}
}

func NewString(value string) IntOrString {
	return IntOrString{
		StringValue: value,
		IsInt:       false,
	}
}
