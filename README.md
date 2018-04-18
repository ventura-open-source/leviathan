# resize image

### dependencies


- [go](https://golang.org/)

- [deb](https://github.com/golang/dep)

- [aws cli]() setup with credentials

- [serverless](https://serverless.com/blog/framework-example-golang-lambda-support/)

### build
> Make sure you're in your ```${GOPATH}/src``` directory, then run:

```bash
make
sls deploy
```

### Use

Path: /upload
Method: POST
Paramethers:

```php
'file' => $file,
's3_store' => [
    'key' => config('filesystems.disks.s3.key'),
    'secret' => config('filesystems.disks.s3.secret'),
    'bucket' => config('filesystems.disks.s3.bucket'),
    'region' => config('filesystems.disks.s3.region'),
    'path' => /img/,
    'acl' => 'public_read',
    'headers' => [
        "Cache-Control" => "max-age=630700000000",
        "Expires" => 2535321600, //year 2050
    ],
],
```

### Todo

[ ] - endpoint to upload image to s3
[ ] - - save image to s3
[ ] - - save metadata (image, user, s3 destination credentials)
[x] - lambda function to resize image on s3
[ ] - lambda function save results to s3 destination storaged 
[ ] - control and register user requests (API keys)

