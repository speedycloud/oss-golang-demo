# 对象管理类接口

待完善

## 接口示例导航

- [ListObjects](#ListObjects)

## ListObjects

### 运行示例

```bash
~# go run listObject.go
```

### 示例输出

```
{
  Contents: [{
      ETag: "\"be704c94ef1c492a843315378e644b0a-3\"",
      Key: "0.4.2.c4.zip",
      LastModified: 2018-08-15 03:20:30.454 +0000 UTC,
      Owner: {
        DisplayName: "liul",
        ID: "liul"
      },
      Size: 39450596,
      StorageClass: "STANDARD"
    },{
      ETag: "\"2a51f16427d31cf5d39d2f6298d10a65-14\"",
      Key: "Doxygen-1.8.14.dmg",
      LastModified: 2018-10-11 10:27:01.957 +0000 UTC,
      Owner: {
        DisplayName: "liul",
        ID: "liul"
      },
      Size: 120299622,
      StorageClass: "STANDARD"
    }],
  IsTruncated: true,
  Marker: "",
  MaxKeys: 2,
  Name: "software",
  NextMarker: "Doxygen-1.8.14.dmg",
  Prefix: ""
}
```