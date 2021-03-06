# Go API client for runalyze

Go client for the [Runalyze Personal API](https://runalyze.com/help/article/personal-api).

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./runalyze"
```

## Documentation for API Endpoints

All URIs are relative to *https://runalyze.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActivityUploadApi* | [**GetActivityByID**](docs/ActivityUploadApi.md#getactivitybyid) | **Get** /api/v1/activities/uploads/{id} | 
*ActivityUploadApi* | [**UploadActivity**](docs/ActivityUploadApi.md#uploadactivity) | **Post** /api/v1/activities/uploads | 
*HealthMetricsApi* | [**CreateBloodPressure**](docs/HealthMetricsApi.md#createbloodpressure) | **Post** /api/v1/metrics/bloodPressure | Creates a new blood pressure entry
*HealthMetricsApi* | [**CreateBodyComposition**](docs/HealthMetricsApi.md#createbodycomposition) | **Post** /api/v1/metrics/bodyComposition | Creates a new body composition entry
*HealthMetricsApi* | [**CreateHeartRateMax**](docs/HealthMetricsApi.md#createheartratemax) | **Post** /api/v1/metrics/heartRateMax | Creates a new maximum heart rate entry
*HealthMetricsApi* | [**CreateHeartRateRest**](docs/HealthMetricsApi.md#createheartraterest) | **Post** /api/v1/metrics/heartRateRest | Creates a new resting heart rate entry
*HealthMetricsApi* | [**CreateMental**](docs/HealthMetricsApi.md#createmental) | **Post** /api/v1/metrics/mental | Creates a new mental state entry
*HealthMetricsApi* | [**CreateMetrics**](docs/HealthMetricsApi.md#createmetrics) | **Post** /api/v1/metrics | Creates bulk entries of all existing metrics
*HealthMetricsApi* | [**CreateSleep**](docs/HealthMetricsApi.md#createsleep) | **Post** /api/v1/metrics/sleep | Creates a new sleeping entry

## Documentation For Models

 - [Apiv1metricsBodyComposition](docs/Apiv1metricsBodyComposition.md)
 - [Apiv1metricsSleep](docs/Apiv1metricsSleep.md)
 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [Body2](docs/Body2.md)
 - [Body3](docs/Body3.md)
 - [Body4](docs/Body4.md)
 - [Body5](docs/Body5.md)
 - [Body6](docs/Body6.md)
 - [Body7](docs/Body7.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


---

## Generated using the following process:

1. Download the swagger spec from the docs:

```console
$ mkdir meta
$ pup -f <(curl -sL https://runalyze.com/doc/personal) script#swagger-data text{} | jq -r .spec > meta/runalyze.json
```
I couldn't find the path to the file but saw the spec in the HTML 😉.

1. As the spec isn't complete, we need to manually edit the `runalyze.json` file and...
  1. Add a description to the `info` section,
  1. Add `"servers": [{"url": "https://runalyze.com"}],` before `"paths"`
  1. give all operations meaningful names, ie replace the `null` in all `operationId` props with something useful like `CreateMetrics`

1. Create a Go specific config file:

```console
$ echo '{"packageName":"runalyze"}' > meta/config.json
```

1. Generate the Go library code:

```console
$ swagger-codegen generate --input-spec meta/runalyze.json --config meta/config.json --lang go --output .
```

1. Append this README to the end of the generated README:

```console
$ cat meta/README.md >> README.md
```
---

## Generated using the following process:

1. Download the swagger spec from the docs:

```console
$ mkdir meta
$ pup -f <(curl -sL https://runalyze.com/doc/personal) script#swagger-data text{} | jq -r .spec > meta/runalyze.json
```
I couldn't find the path to the file but saw the spec in the HTML 😉.

1. As the spec isn't complete, we need to manually edit the `runalyze.json` file and...
  1. Add a description to the `info` section,
  1. Add `"servers": [{"url": "https://runalyze.com"}],` before `"paths"`
  1. give all operations meaningful names, ie replace the `null` in all `operationId` props with something useful like `CreateMetrics`

1. Create a Go specific config file:

```console
$ echo '{"packageName":"runalyze"}' > meta/config.json
```

1. Generate the Go library code:

```console
$ swagger-codegen generate --input-spec meta/runalyze.json --config meta/config.json --lang go --output .
```

1. Append this README to the end of the generated README:

```console
$ cat meta/README.md >> README.md
```

Note: Additional tweaks will be needed to the README.
