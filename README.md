<p align="center">
    <img src="/.github/assets/gopher.png"
        height="300">
</p>

<h1 align="center"><strong>zerolog for Mojito</strong></h1>
<p align="center">
    <a href="https://goreportcard.com/report/github.com/go-mojito/logger-zerolog" alt="Go Report Card">
        <img src="https://goreportcard.com/badge/github.com/go-mojito/logger-zerolog" /></a>
	<a href="https://github.com/go-mojito/logger-zerolog" alt="Go Version">
        <img src="https://img.shields.io/github/go-mod/go-version/go-mojito/logger-zerolog.svg" /></a>
	<a href="https://godoc.org/github.com/go-mojito/logger-zerolog" alt="GoDoc reference">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"/></a>
	<a href="https://github.com/go-mojito/logger-zerolog/blob/main/LICENSE" alt="Licence">
        <img src="https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square" /></a>
	<a href="https://makeapullrequest.com">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome"></a>
</p>
<p align="center">
    <a href="https://go.dev/" alt="Made with Go">
        <img src="https://ForTheBadge.com/images/badges/made-with-go.svg" /></a>
		
</p>
<p align="center">
zerolog for Mojito provides a zerolog Logger implementation that was designed specifically for Mojito with full compatibility.
</p>

<p align="center"><strong>SonarCloud Report</strong></p>
<p align="center">
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_logger-zerolog" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_logger-zerolog&metric=alert_status" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_logger-zerolog" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_logger-zerolog&metric=sqale_rating" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_logger-zerolog" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_logger-zerolog&metric=reliability_rating" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_logger-zerolog" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_logger-zerolog&metric=security_rating" /></a>
	<br>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_logger-zerolog" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_logger-zerolog&metric=vulnerabilities" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_logger-zerolog" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_logger-zerolog&metric=code_smells" /></a>
    <a href="https://sonarcloud.io/summary/overall?id=go-mojito_logger-zerolog" alt="Quality Gate">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=go-mojito_logger-zerolog&metric=bugs" /></a>
</p>

<p align="center"><strong>Quickstart</strong></p>
<p align="center">To start using zerolog as the default logger, simply set it as default like this:</p>

```go
func init() {
    zerolog.AsDefault()
}
```

<p align="center">Or if you want it as a secondary logger you can register it as a named router like this:</p>

```go
func init() {
    zerolog.As("myLogger")
}
```

<p align="center">If you want a more readable, fancy looking log you can enable it by simply adding one line during init:</p>

```go
func init() {
    zerolog.Pretty()
}
```

<p align="center"><strong>Documentation</strong></p>
<p align="center">
	Read our
	<a href="https://go-mojito.infinytum.co/docs">Documentation</a>, check out the 
	<a href="https://go-mojito.infinytum.co/">Project Website</a>.
</p>

<p align="center"><sub>Icon made with <a href="https://gopherize.me">Gopherize</a> and <a href="https://www.flaticon.com/free-icon/mojito_920710">flaticon</a>.</sub></p>
