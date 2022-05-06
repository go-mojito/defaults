package defaults

import (
	"github.com/go-mojito/handlebars"
	zerolog "github.com/go-mojito/logger-zerolog"
	fasthttp "github.com/go-mojito/router-fasthttp"
)

func init() {
	fasthttp.AsDefault()
	handlebars.AsDefault()
	zerolog.AsDefault()
}
