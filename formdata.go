package formdata

import (
	"github.com/grafana/sobek"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/formdata", new(FormData))
}

type FormData struct{}

// XBuilder creates Builder object
// following call is used in js: `const builder = new formdata.Builder();`
func (f *FormData) XBuilder(call sobek.ConstructorCall, rt *sobek.Runtime) *sobek.Object {
	return rt.ToValue(NewBuilder()).ToObject(rt)
}
