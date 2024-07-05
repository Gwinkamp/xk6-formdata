package formdata

import (
	"github.com/dop251/goja"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/formdata", new(FormData))
}

type FormData struct{}

// XBuilder creates Builder object
// following call is used in js: `const builder = new formdata.Builder();`
func (f *FormData) XBuilder(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {
	return rt.ToValue(NewBuilder()).ToObject(rt)
}
