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

- [ ] endpoint to upload image to s3
- [x] - save image to s3
- [ ] - save metadata (image, user, s3 destination credentials)
- [x] lambda function to resize image on s3
- [x] lambda function save results to s3 destination storaged 
- [ ] control and register user requests (API keys)

### Reports
>Downloaded **1700700 bytes**
REPORT RequestId: 7ac49325-4732-11e8-a3d4-cb703b18e9f2	Duration: **6619.20 ms**	Billed Duration: 6700 ms 	Memory Size: 1024 MB	Max Memory Used: **131 MB**

>Downloaded **3300624 bytes**
REPORT RequestId: 4ca640dd-4736-11e8-afff-f50762fb573d	Duration: **9996.33 ms**	Billed Duration: 10000 ms 	Memory Size: 1024 MB	Max Memory Used: **224 MB**

>Downloaded **5254922 bytes**
REPORT RequestId: 7a7879bb-4733-11e8-abec-21ad946dacd3	Duration: **10561.86 ms**	Billed Duration: 10600 ms 	Memory Size: 1024 MB	Max Memory Used: **216 MB**

