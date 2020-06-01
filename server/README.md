PCR Server
==========

The PCR server is used as an interface to create channels, upload packages and
provide any meta-data api for downstream users such as the web application.

By default, the application (registry) server runs on port 5060 and the repository runs 
on port 5050. The application server handles things such as channel creation, meta 
information api, package upload etc. The web interface and CLI tool mainly touches this
layer. The repository server is the one that is used whenever you do a `conda install`. 

## API

Here's a laundry list of the api

### `/healthcheck` [GET]

Runs a health check on the server. Returns status code **200** if everything
is okay

### `/meta` [GET]

Returns meta information about the application

```typescript
type Output = {
  image:      string // name of the image used to index the channels"
  registry:   string // url of the registry server (usually 5060)
  repository: string // url of the repository server (usually 5050)
}
``` 

### `/channel` [GET]

Returns a list of all the channels

```typescript
type Output = {
    channel:   string // channel name
    email:     string // email address 
    join_date: string // date channel was created
}[]
```

### `/channel/{name} [GET]

Gets information about the specified channel

```typescript
type Output = {
    channel:   string // channel name
    email:     string // email address 
    join_date: string // date channel was created
}
```

### `/channel [POST]

Creates a channel. The credentials specified are used for validation when 
uploading/removing packages

```typescript
type Input = {
    channel:  string // channel name
    password: string  // channel password
    email:    string // email address
}

type Output = {
    channel: string // channel name
    email: string // email address 
    join_date: string // date channel was created
}
```

### `/channel/check [POST]

Used to validate a sign in as channel owner. 

Returns status code:
  - **200** if the credentials are valid
  - **403** if credentials are invalid
  - **400** any other errors

```typescript
type Input = {
    channel:  string // channel name
    password: string  // channel password
}
```

### `/channel` [DELETE]

Removes a channel from the repository. This will also delete all packages that 
have been uploaded onto the channel.

Returns:
  - **200** if deleted successfully 
  - **400** if any errors 
  
```typescript
type Input = {
    channel: string // channel name
    password: string  // channel password
}
```

### `/p` [GET]

Get a list of all packages in the repository.

```typescript
type Output = {
    channel:        string
    platforms:      string
    version:        string | null
    description:    string | null
    dev_url:        string | null
    doc_url:        string | null
    home:           string | null
    license:        string | null
    summary:        string | null
    timestamp:      string
    name:           string
}[]
```

### `/p/{channel}` [GET]

Get a list of all packages in the specified channel

```typescript
type Output = {
    channel:        string
    platforms:      string
    version:        string | null
    description:    string | null
    dev_url:        string | null
    doc_url:        string | null
    home:           string | null
    license:        string | null
    summary:        string | null
    timestamp:      string
    name:           string
}[]
```

### `/p/{channel}/{package}` [GET]

Get meta information about the specific package in the channel

```typescript
type Output = {
	channel: string
	package: string
	details: {
        channel:        string
        package:        string
        build_string:   string
        build_number:   number
        version:        string
        platform:       string
        count:          number
        upload_date:    string
    }   
	latest: {
        channel:        string
        platforms:      string
        version:        string | null
        description:    string | null
        dev_url:        string | null
        doc_url:        string | null
        home:           string | null
        license:        string | null
        summary:        string | null
        timestamp:      string
        name:           string
    }
}
```

### `/p` [POST]

Uploads a package to the specified channel. This must be a form post 
request. The file must be a tar.bz2 that is created after the conda-build
process

```typescript
type Input = {
    channel:  string
    password: string
    file:     File
}

type Output = {
    name:           string
    version:        string
    build_string:   string
    build_number:   number
    platform:       string
}
```

### `/p` [DELETE]

Remove a single package specified by the input

Returns:
  - **200** if removed successfully
  - **400** if input data is malformed or invalid
  - **500** if server encountered any error during the process
  
```typescript
type Input = {
    channel:  string
    password: string
    package: {
        name:        string
        version:     string
        buildString: string
        buildNumber: number
        platform:    string
      }
}
```

### `/p/{package}` [DELETE]

Removes all the packages in the channel for the specified package. 
For example, if you have a `numpy` package with different versions and
for different platforms, this is used to remove all of them.

```typescript
type Input = {
    channel:  string // channel name
    password: string  // channel password
}
```
