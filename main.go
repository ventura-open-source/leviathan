package main

import (
    "bytes"
    "fmt"
    "os"
    "context"
    "image/jpeg"
    "strings"
    "strconv"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    // "github.com/nfnt/resize"
    //"gopkg.in/h2non/bimg.v1"

    //"github.com/aws/aws-sdk-go/aws"
    //"github.com/aws/aws-sdk-go/aws/awsutil"
    //"github.com/aws/aws-sdk-go/aws/credentials"
    //"github.com/aws/aws-sdk-go/service/s3"
    //"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sizes = [17]string {
    "2048x1536",
    "640x480",
    "480x360",
    "226x170",
    "192x144",
    "181x136",
    "173x130",
    "160x120",
    "108x81",
    "96x72",
    "60x45",
    // new sizes for responsive web site
    "1920x1440",
    "1200x900",
    "992x744",
    "768x576",
    "544x408",
    "136x102",
}

func Handler(ctx context.Context, s3Event events.S3Event) {

    destBucket := "leviathan-det"
    bucket     := s3Event.Records[0].S3.Bucket.Name
    item       := s3Event.Records[0].S3.Object.Key

    fmt.Println("%s - %s => %s", bucket, item, destBucket)

    sess, err := session.NewSession()

    if err != nil {
        exitErrorf("Error session %v", err)
    }

    downloader := s3manager.NewDownloader(sess)

    buff := &aws.WriteAtBuffer{}
    numBytes, err := downloader.Download(buff,
        &s3.GetObjectInput{
            Bucket: aws.String(bucket),
            Key:    aws.String(item),
        })

    if err != nil {
        exitErrorf("Unable to download item %q, %v", item, err)
    }

    fmt.Println("Downloaded", numBytes, "bytes")

    tmp := bytes.NewReader(buff.Bytes());
    img, err := jpeg.Decode(tmp)
    if err != nil {
        exitErrorf("Unable to decode", err)
    }

    uploader := s3manager.NewUploader(sess)

    for i, v := range sizes {
        //get the width and height 
        size := strings.Split(v, "x")
        w, _ := strconv.ParseUint(size[0], 10, 64)
        //h, _ := strconv.ParseUint(size[1], 10, 64)
        fmt.Println("%d. resizing to %s => %T, %T", i, v, w, img)

        // resize to especific size using Nearest-neighbor interpolation
        // and preserve aspect ratio
        m := resize.Resize(uint(w), 0, img, resize.NearestNeighbor)

        buffer := make([]byte, 0)
        file := bytes.NewBuffer(buffer);
        jpeg.Encode(file, m, nil)

        filename := item + "resized_" + v + ".jpg"

        _, err = uploader.Upload(&s3manager.UploadInput{
            Bucket: aws.String(destBucket),
            Key: aws.String(filename),
            Body: file,
        })

        if err != nil {
            // Print the error and exit.
            exitErrorf("Unable to upload %q to %q, %v", filename, destBucket, err)
        }

        fmt.Printf("Successfully uploaded %q to %q\n", filename, destBucket)
    }


    /*
     *aws_access_key_id := "Insert Key ID here"
     *aws_secret_access_key := "Insert Secret Here"
     *token := ""
     *creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
     *_, err := creds.Get()
     *if err != nil {
     *    fmt.Printf("bad credentials: %s", err)
     *}
     *cfg := aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds)
     *svc := s3.New(session.New(), cfg)
     */


/*
 *
 *    // open "test.jpg"
 *    file, err := os.Open("test.jpg")
 *    if err != nil {
 *        log.Fatal(err)
 *    }
 *
 *    // decode jpeg into image.Image
 *    img, err := jpeg.Decode(file)
 *    if err != nil {
 *        log.Fatal(err)
 *    }
 *    file.Close()
 *
 *    for i, v := range sizes {
 *        //get the width and height 
 *        size := strings.Split(v, "x")
 *        w, _ := strconv.ParseUint(size[0], 10, 64)
 *        //h, _ := strconv.ParseUint(size[1], 10, 64)
 *        fmt.Printf("%d. resizing to %s \n", i, v)
 *
 *        // resize to especific size using Nearest-neighbor interpolation
 *        // and preserve aspect ratio
 *        m := resize.Resize(uint(w), 0, img, resize.NearestNeighbor)
 *
 *        out, err := os.Create("test_resized_" + v + ".jpg")
 *        if err != nil {
 *            log.Fatal(err)
 *        }
 *        defer out.Close()
 *
 *        // write new image to file
 *        jpeg.Encode(out, m, nil)
 *    }
 */
}

func main() {
    lambda.Start(Handler)
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}
