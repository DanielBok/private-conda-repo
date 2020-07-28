# Overview

The Private Conda Repository is a place to host your fantastic conda packages.

## Register an account

To get started, register an account with the GUI. The account name will also be your
channel name. Once you're done, you can start uploading and sharing your packages.

## CLI tool

A CLI tool will be available for each release. You can download the CLI
[here](https://github.com/DanielBok/private-conda-repo/releases).

Assuming you saved your tool as `PCR` in your path

### Using the CLI tool

To start using the CLI tool, you must first set the registry and login.
To register your console with the API server:

```sh
pcr registry set @registry
```

Subsequently, login to your console via

```bash
pcr registry login  # you will be prompted for your username and password here
```

### Uploading packages

```sh
pcr upload path/to/your/package.tar.bz2
```

Please note that your package must be in the `.tar.bz2` format. If you built your
package via `conda-build`, this will be the default format.

## More documentation

Underneath the hood, the CLI tool makes a couple of API calls to the server.

The full list of the available API is documented [here](https://github.com/DanielBok/private-conda-repo/tree/master/server).
These APIs can be used in the event that you want to build your own CI/CD process
for package publishing.
