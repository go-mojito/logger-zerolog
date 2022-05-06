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

<p align="center"><strong>How to register</strong></p>
<p align="center">To register the zerolog as the default logger, do the following in your main file</p>

```go
func init() {
    zerolog.AsDefault()
}
```
<p align="center">To register zerolog as a named logger, do the following in your main file</p>

```go
func init() {
    zerolog.As("myLogger")
}
```
<p align="center">You can use fancy logging by enabling it via:</p>

```go
func init() {
    zerolog.Pretty()
}
```
