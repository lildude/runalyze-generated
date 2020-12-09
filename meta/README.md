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
