//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// databaseawscrossplaneio
package databaseawscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateRdsInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRdsInstance_ManifestParameters(props *RdsInstanceProps) error {
	return nil
}

func validateRdsInstance_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewRdsInstanceParameters(scope constructs.Construct, id *string, props *RdsInstanceProps) error {
	return nil
}

