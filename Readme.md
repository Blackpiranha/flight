## 获取航班起始点位置

### 1) 请求地址

>http://127.0.0.1:8080/v1/tracker

### 2) 调用方式：HTTP POST

### 3) 接口描述：

* 获取航班起始点位置

### 4) 请求参数:

#### BODY请求体参数:
| 字段名称    | 字段说明 |类型            | 必填 |备注     |
|---------|:----:|:--------------:|:--:| ------:|
| address |{"start":"IND", "end":"EWR"}| list |Y| 不能为空  |

```
{
    "address": [{"start":"IND", "end":"EWR"}, {"start":"SFO", "end":"ATL"}, {"start":"GSO", "end":"IND"}, {"start":"ATL", "end":"GSO"}]
}
```

### 5) 请求返回正确结果:

```
{
    "code": 200,
    "data": {
        "start": "SFO",
        "end": "EWR"
    }
}
```

### 5) 请求返回错误结果:

```
{
    "code": 400,
    "message": "invalid parameter;"
}
```


### END  