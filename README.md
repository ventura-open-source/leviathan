# resize image

### dependencies


- [go](https://golang.org/)

- [deb](https://github.com/golang/dep)

- [aws cli]() setup with credentials

- [serverless](https://serverless.com/blog/framework-example-golang-lambda-support/)

### Start

- update variables in config.json

### build
> Make sure you're in your ```${GOPATH}/src``` directory, then run:

```bash
make
sls deploy
```

### Todo

[ ] - endpoint to upload image to s3
[x] - - save image to s3
[ ] - - save metadata (image, user, s3 destination credentials)
[x] - lambda function to resize image on s3
[x] - lambda function save results to s3 destination storaged 
[ ] - control and register user requests (API keys)

