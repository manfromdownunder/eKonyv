# eKonyv Docker Image

> See [Github](https://github.com/manfromdownunder/eKonyv) for more information.

## Usage

### Basic

`docker run -d -p 8080:80 manfromdownunder/eKonyv --name my_irc_name`

### Persist eBook Files

`docker run -d -p 8080 -v ~/Downloads:/books manfromdownunder/eKonyv --name my_irc_name --persist`

### Host at a sub path behind a reverse proxy

`docker run -d -p 8080:80 -e BASE_PATH=/eKonyv/ manfromdownunder/eKonyv --name my_irc_name`

## Arguments

```
--name string
    Required name when connecting to irchighway
--persist
    Keep book files in the download dir. Default is to delete after sending.
```

## Docker Compose

```docker
version: '3.3'
services:
    eKonyv:
        ports:
            - '8080:80'
        volumes:
            - 'booksVolume:/books'
        restart: unless-stopped
        container_name: eKonyv
        command: --name my_irc_name --persist
        environment:
          - BASE_PATH=/eKonyv/
        image: manfromdownunder/eKonyv:latest

volumes:
    booksVolume:
```
