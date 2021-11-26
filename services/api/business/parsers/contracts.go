package parsers

import (
	"encoding/json"
	"regexp"

	pkgjson "github.com/consensys/orchestrate/pkg/encoding/json"
	ethabi "github.com/consensys/orchestrate/pkg/ethereum/abi"
	"github.com/consensys/orchestrate/pkg/types/entities"
)

var contractRegexp = `^(?P<contract>[a-zA-Z0-9]+)(?:\[(?P<tag>[0-9a-zA-Z-.]+)\])?(?::(?P<abi>\[.+\]))?(?::(?P<bytecode>0[xX][a-fA-F0-9]+))?(?::(?P<deployedBytecode>0[xX][a-fA-F0-9]+))?$`
var contractPattern = regexp.MustCompile(contractRegexp)

// ParseJSONABI returns a decoded ABI object
func ParseJSONABI(data string) (methods, events map[string]string, err error) {
	var parsedFields []entities.RawABI
	err = pkgjson.Unmarshal([]byte(data), &parsedFields)
	if err != nil {
		return nil, nil, err
	}

	// Retrieve raw JSONs
	normalizedJSON, err := pkgjson.Marshal(parsedFields)
	if err != nil {
		return nil, nil, err
	}
	var rawFields []json.RawMessage
	err = pkgjson.Unmarshal(normalizedJSON, &rawFields)
	if err != nil {
		return nil, nil, err
	}

	methods = make(map[string]string)
	events = make(map[string]string)
	for i := 0; i < len(rawFields) && i < len(parsedFields); i++ {
		fieldJSON, err := rawFields[i].MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
		switch parsedFields[i].Type {
		case "function", "":
			var m *ethabi.Method
			err := json.Unmarshal(fieldJSON, &m)
			if err != nil {
				return nil, nil, err
			}
			methods[m.Name+m.Sig()] = string(fieldJSON)
		case "event":
			var e *ethabi.Event
			err := json.Unmarshal(fieldJSON, &e)
			if err != nil {
				return nil, nil, err
			}
			events[e.Name+e.Sig()] = string(fieldJSON)
		}
	}
	return methods, events, nil
}

