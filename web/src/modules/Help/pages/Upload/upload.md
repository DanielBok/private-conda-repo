# Uploading Packages

You can upload your packages using a "multipart/form" POST request via curl, postman or
using the tools built in PCR.

<Divider />

## The basic request

Read this section if you want to know how it works or if you want to upload "raw".
A sample post request using javascript and **axios** is listed below.

<BasicRequest />
<Divider />

## Using the CLI

If you're using the cli, follow the steps written
[here](https://github.com/DanielBok/private-conda-repo/tree/master/cli). Remember to
set the `--no-abi` flag if you need it as it is not there by default.

## Using the web interface

Assuming you've already created a channel and is logged onto the channel, you can fill
up the form at this [link](@link) to upload a package. Remember that the package has
to have the `.tar.bz2` suffix (which shouldn't be a problem if you did your
`conda build` properly).
