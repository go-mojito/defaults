package defaults

import (
	"github.com/go-mojito/handlebars"
	fasthttp "github.com/go-mojito/router-fasthttp"
)

func init() {
	fasthttp.AsDefault()
	handlebars.AsDefault()
}
