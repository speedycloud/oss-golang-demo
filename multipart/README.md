# 分段管理类接口

待完善

## 接口示例导航

- [ListMultipartUploads](#ListMultipartUploads)

## ListMultipartUploads

### 运行示例

```bash
~# go run listMulitpartUploads.go
```

### 示例输出

```
{
  Bucket: "software",
  IsTruncated: false,
  MaxUploads: 1000,
  NextKeyMarker: "B0F0888A648E08E5B35A5E298A9D69A9.mp4",
  NextUploadIdMarker: "2~e3S5jN5sZWcE87ElFxVQOOvFghALCfF",
  Uploads: [{
      Initiated: 2018-10-16 11:11:55.393 +0000 UTC,
      Initiator: {
        DisplayName: "liul",
        ID: "liul"
      },
      Key: "B0F0888A648E08E5B35A5E298A9D69A9.mp4",
      Owner: {
        DisplayName: "liul",
        ID: "liul"
      },
      StorageClass: "STANDARD",
      UploadId: "2~e3S5jN5sZWcE87ElFxVQOOvFghALCfF"
    }]
}
```