package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/thedodd/barge/core"
	"github.com/thedodd/barge/registry"

	"github.com/hashicorp/hcl"
	"github.com/mitchellh/cli"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v8"
)

var validate = validator.New(&validator.Config{TagName: "validate"})

func init() {
	validate.RegisterValidation("validDriver", validDriver)
}

// GetConfig taken from the local Bargefile. Will panic if Bargefile not found or if there was trouble reading the file.
func GetConfig(ui cli.Ui) (*core.Bargefile, error) {
	// TODO(TheDodd): make this optional at some point, like the Appfile.
	// Attempt to read from the Bargefile.
	bargeBytes, err := ioutil.ReadFile("Bargefile")
	if err != nil {
		return nil, errors.New("Bargefile not found in current directory.")
	}

	// Unmarshal bytes onto raw map.
	bargefile := &core.Bargefile{Development: &core.DevEnvConfig{}}
	rawBargefile := &map[string]map[string]interface{}{}
	if err := hcl.Unmarshal(bargeBytes, rawBargefile); err != nil {
		handleConfigErrors(ui, bargefile, err)
		return nil, errors.New("Error reading Bargefile.")
	}

	// Map raw Bargefile config onto Bargefile struct.
	if err := mapstructure.Decode(rawBargefile, bargefile); err != nil {
		handleConfigErrors(ui, bargefile, err)
		return nil, errors.New("Error(s) parsing Bargefile.")
	}

	// Validate Bargefile.
	if err := validate.Struct(bargefile); err != nil {
		handleConfigErrors(ui, bargefile, err)
		return nil, errors.New("Error(s) validating Bargefile.")
	}

	return bargefile, nil
}

// Handle all errors related to getting runtime configuration.
func handleConfigErrors(ui cli.Ui, bargefile *core.Bargefile, err error) {
	switch eType := err.(type) {
	case validator.ValidationErrors:
		handleValidationErrors(ui, bargefile, eType)

	case *mapstructure.Error:
		for _, er := range eType.WrappedErrors() {
			ui.Error(er.Error())
		}

	default:
		ui.Error(eType.Error())
	}
}

// Handle validator validation errors.
func handleValidationErrors(ui cli.Ui, bargefile *core.Bargefile, errs validator.ValidationErrors) {
	for _, fieldError := range errs {
		switch fieldError.Tag {

		case "validDriver":
			ui.Error(fmt.Sprintf("Driver must be one of: %s Given value: %s", registry.ValidDrivers, fieldError.Value))

		default:
			ui.Error(
				fmt.Sprintf(
					"Validation failed for field '%s' with validator '%s=%s' and value '%v'.",
					fieldError.NameNamespace,
					fieldError.Tag,
					fieldError.Param,
					fieldError.Value,
				),
			)
		}
	}
}

// A custom validator for validating the Development.Provider.
func validDriver(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	valid := false
	for _, val := range registry.ValidDrivers {
		if field.String() == val {
			valid = true
		}
	}
	return valid
}
