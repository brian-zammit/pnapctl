# pnap-cli

To generate mocks, get [`mockgen`](https://github.com/golang/mock) and follow its instructions.

Clone this repository in `~/go/src/phoenixnap.com/`

To test everything, run `go test ./tests/...`. To colorize the output, download `gotest` from `go get -u github.com/rakyll/gotest`.

## Configuration
Details can be passed using a config file. This file can be passed as an argument, or can be read if placed in `~/pnap.yaml`. An example of this file is in `sample-config.yaml`. In order to currently test the application, this `yaml` file can be used by using the following command: `pnapctl bmc --config=sample-config.yaml ...` or simply copying/symlinking the file to your home directory.

## Current folder structure

Every command is its own folder, having a `.go` file that represents it. So, to check `pnapctl bmc get servers`, the directory structure would be `./pnapctl/bmc/get/servers`.


