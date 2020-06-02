Private Conda Repository
========================

Private Conda Repository (PCR) is an application which you can run to host your conda packages on-prem.

## When to use it

You'll likely use it when you're

- working in an enterprise environment
- can't share your conda / python packages on pypi or on conda
- have to share your packages with others in your firm

## Getting started

For all instances, you'll find a Dockerfile which packages the build
and runs it. All you need is [docker](https://www.docker.com/). 

### If you just want a repository

If you just want to have a place to host your packages, then the
[server](https://github.com/DanielBok/private-conda-repo/tree/master/server) 
is all you'll need. Just build the go binary and run it as is.

### If you want a repository and a web interface

If you want to also have the client application, you'll have to
build the client (React + Typescript) application. Head over to
the [web](https://github.com/DanielBok/private-conda-repo/tree/master/web) 
application and build it up. 

Since you're running both of them, I have set up a 
[sample application](https://github.com/DanielBok/private-conda-repo/tree/master/_example)
using docker-compose. You can refer and edit it to your specification.

## Commands

To learn how to interact with the repository, check out the 
documentation at the [server](https://github.com/DanielBok/private-conda-repo/tree/master/server)
section. 

## Command Line Tool

If you'll like a command line tool to help you upload your packages,
you can check it out at the 
[cli](https://github.com/DanielBok/private-conda-repo/tree/master/cli) section. 

## Installing Packages

By default, the application server will run at port 5060 and the 
repository server will run at port 5050. 

The application server is used to handle things such as user creation, meta information api, package upload etc. So the web interface and the 
CLI tool mainly touches this server.

The repository server is the one that is used whenever you do a 
`conda install`. 

For example, if your servers are located at https://my-server.com on 
the default ports and you have created a channel named `pikachu`. To
install a package called `numpy` that you have uploaded, you will call
from you shell

```bash
conda install -c https://my-server.com:5050/pikachu numpy
```

## FAQ

Below are a list of errors I encountered while using it in my firm and the solutions for them.

### Built and published package correctly but keep getting "inflexible solves" during conda install

The likely reason this happens is due to your meta.yaml recipe setup. The simple solution is to run `pcr upload --no-abi path/to/package`.

If you build packages where you need to compile C code, you should know that you need to build this package for every platform you need 
to support. However, if you build a noarch package, technically, you don't have to. But that's technically, because there are instances
where you use a 3.8 feature that is not available in 3.7. Thus when conda indexes the channel, it'll add the `python_abi 3.x*` as one of
the package's dependencies. This may or may not happen based on how you build your package.

Since it's really tedious to solve, if you think you've done everything correctly and that your package does not need to be locked onto a
specific python version or ABI, then I've provided the `--no-abi` flag in the upload command. Basically, it'll read through the 
`current_repodata.json` and `repodata.json` file and remove all `python_abi` dependencies.

Note that it removes for all the packages in the channel and not the specific channel. I may come back in the future to make it a little
more specific when I've got the time.
