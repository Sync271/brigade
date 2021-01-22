package authn

import (
	"encoding/json"
	"reflect"
	"testing"
	"unsafe"

	"github.com/brigadecore/brigade/v2/apiserver/internal/meta"
	"github.com/stretchr/testify/require"
)

// TODO: This isn't very DRY. It would be nice to figure out how to reuse these
// bits across a few different packages. The only way I (krancour) know of
// is to move these into their own package and NOT have them in files suffixed
// by _test.go. But were we to do that, Go would not recognize the functions as
// code used exclusively for testing and would therefore end up dinging us on
// coverage... for not testing the tests. :sigh:

func requireAPIVersionAndType(
	t *testing.T,
	obj interface{},
	expectedType string,
) {
	objJSON, err := json.Marshal(obj)
	require.NoError(t, err)
	objMap := map[string]interface{}{}
	err = json.Unmarshal(objJSON, &objMap)
	require.NoError(t, err)
	require.Equal(t, meta.APIVersion, objMap["apiVersion"])
	require.Equal(t, expectedType, objMap["kind"])
}

func setUnexportedField(
	objPtr interface{},
	fieldName string,
	fieldValue interface{},
) {
	field := reflect.ValueOf(objPtr).Elem().FieldByName(fieldName)
	reflect.NewAt(
		field.Type(),
		unsafe.Pointer(field.UnsafeAddr()),
	).Elem().Set(reflect.ValueOf(fieldValue))
}