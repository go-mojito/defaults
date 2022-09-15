<p align="center">
    <img src="/.github/assets/gopher.png"
        height="300">
</p>

<h1 align="center"><strong>Default Stack for Mojito</strong></h1>
<p align="center">
    <a href="https://goreportcard.com/report/github.com/go-mojito/defaults" alt="Go Report Card">
        <img src="https://goreportcard.com/badge/github.com/go-mojito/defaults" /></a>
	<a href="https://github.com/go-mojito/defaults" alt="Go Version">
        <img src="https://img.shields.io/github/go-mod/go-version/go-mojito/defaults.svg" /></a>
	<a href="https://godoc.org/github.com/go-mojito/defaults" alt="GoDoc reference">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"/></a>
	<a href="https://github.com/go-mojito/defaults/blob/main/LICENSE" alt="Licence">
        <img src="https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square" /></a>
	<a href="https://makeapullrequest.com">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome"></a>
</p>
<p align="center">
    <a href="https://go.dev/" alt="Made with Go">
        <img src="https://ForTheBadge.com/images/badges/made-with-go.svg" /></a>
		
</p>
<p align="center">
Defaults provides our recommended stack of implementations for Mojito with focus on performance and ease of use.
</p>

</p>

<h2 align="center"><strong>Implementations</strong></h2>
<p align="center"><a href="https://github.com/go-mojito/logger-zerolog">ZeroLog</a> &bullet; <a href="https://github.com/go-mojito/router-fasthttp">FastHTTP</a> &bullet; <a href="https://github.com/go-mojito/handlebars">Handlebars</a></p>

<h2 align="center"><strong>How to enable this Stack</strong></h2>
<p align="center">To enable the Infinytum-recommended Stack, import this repository in your mainfile like this.</p>

```go
package main

import(
    ...
    "github.com/go-mojito/mojito"
    _ "github.com/go-mojito/defaults"
    ...
)
```
