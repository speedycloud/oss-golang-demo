## 获取桶列表接口示例 (GetService / ListBuckets)

This implementation of the GET operation returns a list of all buckets owned by the authenticated sender of the request.

To authenticate a request, you must use a valid AWS Access Key ID that is registered with Amazon S3. Anonymous requests cannot list buckets, and you cannot list buckets that you did not create.

### 参考

[https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/API/RESTServiceGET.html](https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/API/RESTServiceGET.html)

### 运行示例

```bash
~# go build listBuckets.go
~# ./listBuckets
```

### 示例输出

```json
{
  Buckets: [{
      CreationDate: 2018-08-10 13:34:39.537 +0000 UTC,
      Name: "software"
    }],
  Owner: {
    DisplayName: "liul",
    ID: "liul"
  }
}
```
