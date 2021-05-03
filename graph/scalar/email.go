// EXAMPLE OF DEFINED TYPE
package scalar

import (
	"errors"
	"fmt"
	"io"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
)

var emailRegex = regexp.MustCompile(".{1,50}@.{1,30}\\..{2,7}") //nolint

// MarshalEmail marshals string into gql Email type.
func MarshalEmail(email string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(`"` + email + `"`))
	})
}

// UnmarshalEmail unmarshal string from gql Email type.
func UnmarshalEmail(v interface{}) (result string, e error) {
	switch v := v.(type) {
	case string:
		match := emailRegex.MatchString(v)
		if !match {
			e = errors.New(fmt.Sprintf("%T is not a valid email, should be *{1,50}@*{1,30}.*{2,7} formatted string", v)) //nolint
		}
		result = v
	default:
		e = errors.New(fmt.Sprintf("%T is not a valid email string", v)) //nolint
	}

	return result, e
}
