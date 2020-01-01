PCR Server
==========

The PCR server is used as an interface to create channels, upload packages and
provide any meta-data api for downstream users such as the web application.

By default, the application server runs on port 5060 and the repository runs on
port 5050. 

The application server is used to handle things such as user creation, meta information api, package upload etc. So the web interface and the 
CLI tool mainly touches this server.

The repository server is the one that is used whenever you do a `conda install`. 

## Nomenclature

Users and channels are used interchangeably in this application. The username
is also the channel name from which you install or upload your packages

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

### `/user` [GET]

Returns a list of all the users / channels

```typescript
type Output = {
    channel:   string // channel name
    email:     string // email address 
    join_date: string // date channel was created
}[]
```

### `/user/{name} [GET]

Gets information about the single channel / user

```typescript
type Output = {
    channel:   string // channel name
    email:     string // email address 
    join_date: string // date channel was created
}
```

### `/user [POST]

Creates a channel. The credentials are used later to upload/remove packages

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

### `/user/check [POST]

Used to validate a user. 

Returns status code:
  - **200** if the user is valid
  - **403** if user / password do not match
  - **400** any other errors

```typescript
type Input = {
    channel:  string // channel name
    password: string  // channel password
}
```

### `/user` [DELETE]

Removes a user and channel from the repository. This will also delete
all packages that the user uploaded previously.

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

### `/p/{user}` [GET]

Get a list of all packages that the specified user has uploaded

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

### `/p/{user}/{package}` [GET]

Get meta information about the specific package that the user has uploaded.

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
