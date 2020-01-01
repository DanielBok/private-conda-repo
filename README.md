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
