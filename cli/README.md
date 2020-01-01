PCR CLI
=======

The PCR cli tool is a tool used to help you manage the packages. 

## Building the tool

The Makefile shows an example of how to build the tool for your computer.
The command is for a Windows computer, you can adapt it to whatever you're
using. Note that you need to have 
[Go version >= 13](https://golang.org/dl/) to build the cli.

You can use Docker to build the CLI tool too. Just specify the GOOS and GOARCH
variables before the build and share the build folder volume with your host.
Here's a [reference](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
for all valid GOOS and GOARCH values.

## Using the CLI

Assuming your built binary is called `pcr`, you can get help by calling

```bash
pcr help
```

### Set the Repository

To get started, you'll need to set your registry. If your registry is running
at `http://localhost:5050`, then run the following command

```bash
pcr registry set http://localhost:5060
``` 

Now that the CLI knows where to communicate, you can either login or create
a account / channel. If you've already created an account via the web
interface, you can proceed to login straightaway.

### Signup or Login 

```bash
# if creating an account
pcr registry register -c <channel> -p <password> -e <email>

# if logging in
pcr registry login -c <channel> -p <password>
```

If you don't want to expose any of the values in shell, you can just run
`pcr registry login` and you'll be prompted for the channel and password.
In this instance, **the password will be masked**.

### Uploading package

Assuming you went through all the steps to build your pacakge via
`conda build ...`, you can upload your package by 

```bash
pcr upload path/to/pkg

# example
pcr upload dist/noarch/numpy-0.1.1-py_0.tar.bz
```
