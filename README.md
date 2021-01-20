![stona-banner](./assets/stona-banner-alt.png)

# Stona Storage

Stona is an API service that serves your files

## 📦 Installation

Stona has two way for installation. You can install it by follow the steps below when you decide to install.

- [Development](#installation--development-)
- [Locally](#installation--locally-)
- Docker ( Not stable yet )

### Installation ( Development )

- First of all, check your machine whether it has golang, If it doesn't,you can find it [here](https://golang.org/)

- Clone this repository with git 

```sh
    git clone https://github.com/ali-furkqn/stona
    cd stona
```

- Download Dependencies

```sh
    make install
```

- Copy `.env.example` and paste it as `.env`, fill the fields

- Download [**Service account key**](https://console.cloud.google.com/apis/credentials/serviceaccountkey) as JSON and fill the env (special configuration file is going to adding soon. You won't need to fill env)

- Start the application
    - To start as development mode with hot reloading Use `make watch`
    - To start as production mode:
        - Build the application `make`
        - Start the generated binary file. For example `"./main"`

### Installation ( Locally )

- Download latest builded application in [release](./releases)

- Downloads your [**Service account key**](https://console.cloud.google.com/apis/credentials/serviceaccountkey) as JSON and fill the env (like a [env.example](./env.example))

- Start the application with terminal. `./stona`

## API Reference

### Base URL

You can change Base Path of Stona API with `ROOT_PATH` field of env as example 

**Example Base URL**
```
    https://your-awesome-domain/ROOT_PATH
```

### Authentication

Authentication is performed with the `Authorization` HTTP header in the format `Authorization Bearer YOUR_ACCESS_TOKEN`

**Example Authorization Header**
```
    Authorization: Bearer QUs97eBB7an8q3AEFg8qSPRguq524cgKHzvs
```

### Get the file

**GET** `/{bucket.folder}/{file.id}`

Returns a file data for given `bucker.folder` and `file.id` 

#### Query String Params

| Name  | Type          | Description   | Required  | Default           |
|-------|---------------|---------------|:---------:|-------------------|
| size  | integer (px)  | image width   | false     | Orginally size    |

### Get Current Bucket Folder Files

**GET** `/{bucket.folder}`

Returns a `bucket.folder` list of partial files. Requires the **authentication**.

#### Query String Params

| Name  | Type          | Description                               | Required  | Default   |
|-------|---------------|-------------------------------------------|:---------:|-----------|
| size  | integer       | max number of files to returns (1-100)    | false     | 10        |
| begin | integer       | list's starting point                     | false     | 0         |

#### Example Usage

**Request:**

```
GET https://s.alifurkan.co/test-folder?size=1&begin=0 (with bearer token)
```

**Response:**

```json
[
    {
        "type": "image/png",
        "url": "https://s.alifurkan.co/test-folder/85498988458459.png",
        "size": 350545,
        "createdAt": "2020-12-17T00:00:00.000Z" 
    }
]
```

### Put New File

**PUT** `/{bucket.folder}`

**Note:** `{image.id}` can add optional on path like `/{bucket.folder}/{image.id}`

Returns url and metadata for given `{bucket.folder}`, if request finalized success. Requires the authentication.

#### Form Data Params

| Name  | Type          | Description       | Required  | Default   |
|-------|---------------|-------------------|:---------:|-----------|
| file  | file          | file to attach    | true      | none      |


### Delete File

**DELETE** `/{bucket.folder}/{file.id}`

Returns deleted `{file.id}` metadata for given `{bucket.folder}` and `{file.id}`. Requires the authentication.

## License

Stona is [MIT licensed](LICENSE)
