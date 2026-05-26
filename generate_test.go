package everything

import (
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
	c2go.Generate(t, []c2go.BindgenConfig{{
		HeadersDir:  "clone/include",
		OutputDir:   ".",
		PackageName: "everything",
		HeaderOrder: []string{"Everything.h"},
		BindDll:     true,
		DllName:     "Everything64.dll",
		Predefined: `
#define EVERYTHINGAPI
#define EVERYTHINGUSERAPI
#define UNICODE
#define _UNICODE

`,
		DllFuncFilter: func(name string) bool {
			return strings.HasPrefix(name, "Everything") || strings.HasPrefix(name, "EVERYTHING")
		},
	}})
}
