# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [armory/v1/armory.proto](#armory_v1_armory-proto)
    - [BaseResponse](#armory-v1-BaseResponse)
    - [Device](#armory-v1-Device)
    - [GetPrintStatusRequest](#armory-v1-GetPrintStatusRequest)
    - [GetPrintStatusResponse](#armory-v1-GetPrintStatusResponse)
    - [GetSlicerResourcesRequest](#armory-v1-GetSlicerResourcesRequest)
    - [GetSlicerResourcesResponse](#armory-v1-GetSlicerResourcesResponse)
    - [GetTTCodeRequest](#armory-v1-GetTTCodeRequest)
    - [GetTTCodeResponse](#armory-v1-GetTTCodeResponse)
    - [ListMyMessagesRequest](#armory-v1-ListMyMessagesRequest)
    - [ListMyMessagesResponse](#armory-v1-ListMyMessagesResponse)
    - [ListMyTasksRequest](#armory-v1-ListMyTasksRequest)
    - [ListMyTasksResponse](#armory-v1-ListMyTasksResponse)
    - [ListUserDevicesRequest](#armory-v1-ListUserDevicesRequest)
    - [ListUserDevicesResponse](#armory-v1-ListUserDevicesResponse)
    - [LoginRequest](#armory-v1-LoginRequest)
    - [LoginResponse](#armory-v1-LoginResponse)
    - [MessageUser](#armory-v1-MessageUser)
    - [PrintDeviceStatus](#armory-v1-PrintDeviceStatus)
    - [PrintTask](#armory-v1-PrintTask)
    - [RefreshTokenRequest](#armory-v1-RefreshTokenRequest)
    - [RefreshTokenResponse](#armory-v1-RefreshTokenResponse)
    - [Resource](#armory-v1-Resource)
    - [Software](#armory-v1-Software)
    - [TaskMessage](#armory-v1-TaskMessage)
    - [UserMessage](#armory-v1-UserMessage)
  
    - [Armory](#armory-v1-Armory)
  
- [Scalar Value Types](#scalar-value-types)



<a name="armory_v1_armory-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## armory/v1/armory.proto



<a name="armory-v1-BaseResponse"></a>

### BaseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |
| code | [int32](#int32) |  |  |
| error | [string](#string) |  |  |






<a name="armory-v1-Device"></a>

### Device



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dev_id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| online | [bool](#bool) |  |  |
| dev_model_name | [string](#string) |  |  |
| dev_product_name | [string](#string) |  |  |






<a name="armory-v1-GetPrintStatusRequest"></a>

### GetPrintStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| force | [bool](#bool) |  |  |






<a name="armory-v1-GetPrintStatusResponse"></a>

### GetPrintStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [BaseResponse](#armory-v1-BaseResponse) |  |  |
| devices | [PrintDeviceStatus](#armory-v1-PrintDeviceStatus) | repeated |  |






<a name="armory-v1-GetSlicerResourcesRequest"></a>

### GetSlicerResourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |






<a name="armory-v1-GetSlicerResourcesResponse"></a>

### GetSlicerResourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [BaseResponse](#armory-v1-BaseResponse) |  |  |
| software | [Software](#armory-v1-Software) |  |  |
| resources | [Resource](#armory-v1-Resource) | repeated |  |






<a name="armory-v1-GetTTCodeRequest"></a>

### GetTTCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dev_id | [string](#string) |  |  |






<a name="armory-v1-GetTTCodeResponse"></a>

### GetTTCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [BaseResponse](#armory-v1-BaseResponse) |  |  |
| ttcode | [string](#string) |  |  |
| passwd | [string](#string) |  |  |
| authkey | [string](#string) |  |  |






<a name="armory-v1-ListMyMessagesRequest"></a>

### ListMyMessagesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [int32](#int32) |  |  |
| after | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |






<a name="armory-v1-ListMyMessagesResponse"></a>

### ListMyMessagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hits | [UserMessage](#armory-v1-UserMessage) | repeated |  |






<a name="armory-v1-ListMyTasksRequest"></a>

### ListMyTasksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device_id | [string](#string) |  |  |
| after | [string](#string) |  |  |
| limit | [int32](#int32) |  |  |






<a name="armory-v1-ListMyTasksResponse"></a>

### ListMyTasksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int32](#int32) |  |  |
| hits | [PrintTask](#armory-v1-PrintTask) | repeated |  |






<a name="armory-v1-ListUserDevicesRequest"></a>

### ListUserDevicesRequest







<a name="armory-v1-ListUserDevicesResponse"></a>

### ListUserDevicesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base | [BaseResponse](#armory-v1-BaseResponse) |  |  |
| devices | [Device](#armory-v1-Device) | repeated |  |






<a name="armory-v1-LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| password | [string](#string) |  |  |
| code | [string](#string) |  |  |






<a name="armory-v1-LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |
| login_type | [string](#string) |  |  |
| expires_in | [int64](#int64) |  |  |






<a name="armory-v1-MessageUser"></a>

### MessageUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| avatar | [string](#string) |  |  |
| is_followed | [bool](#bool) |  |  |






<a name="armory-v1-PrintDeviceStatus"></a>

### PrintDeviceStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dev_id | [string](#string) |  |  |
| dev_name | [string](#string) |  |  |
| dev_online | [bool](#bool) |  |  |
| progress | [int32](#int32) |  |  |
| task_id | [string](#string) |  |  |






<a name="armory-v1-PrintTask"></a>

### PrintTask



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| model_id | [string](#string) |  |  |
| title | [string](#string) |  |  |
| cover | [string](#string) |  |  |
| status | [int32](#int32) |  |  |
| start_time | [string](#string) |  |  |
| end_time | [string](#string) |  |  |
| weight | [double](#double) |  |  |
| cost_time | [int64](#int64) |  |  |
| device_id | [string](#string) |  |  |
| mode | [string](#string) |  |  |






<a name="armory-v1-RefreshTokenRequest"></a>

### RefreshTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| refresh_token | [string](#string) |  |  |






<a name="armory-v1-RefreshTokenResponse"></a>

### RefreshTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| access_token | [string](#string) |  |  |
| refresh_token | [string](#string) |  |  |
| expires_in | [int64](#int64) |  |  |
| refresh_expires_in | [int64](#int64) |  |  |






<a name="armory-v1-Resource"></a>

### Resource



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| version | [string](#string) |  |  |
| description | [string](#string) |  |  |
| url | [string](#string) |  |  |
| force_update | [bool](#bool) |  |  |






<a name="armory-v1-Software"></a>

### Software



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [string](#string) |  |  |
| description | [string](#string) |  |  |
| url | [string](#string) |  |  |
| force_update | [bool](#bool) |  |  |






<a name="armory-v1-TaskMessage"></a>

### TaskMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| cover | [string](#string) |  |  |
| status | [int32](#int32) |  |  |
| device_id | [string](#string) |  |  |






<a name="armory-v1-UserMessage"></a>

### UserMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| type | [int32](#int32) |  |  |
| task_message | [TaskMessage](#armory-v1-TaskMessage) |  |  |
| from | [MessageUser](#armory-v1-MessageUser) |  |  |
| create_time | [string](#string) |  |  |





 

 

 


<a name="armory-v1-Armory"></a>

### Armory
---------- Auth ----------

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#armory-v1-LoginRequest) | [LoginResponse](#armory-v1-LoginResponse) |  |
| RefreshToken | [RefreshTokenRequest](#armory-v1-RefreshTokenRequest) | [RefreshTokenResponse](#armory-v1-RefreshTokenResponse) |  |
| ListMyMessages | [ListMyMessagesRequest](#armory-v1-ListMyMessagesRequest) | [ListMyMessagesResponse](#armory-v1-ListMyMessagesResponse) |  |
| ListMyTasks | [ListMyTasksRequest](#armory-v1-ListMyTasksRequest) | [ListMyTasksResponse](#armory-v1-ListMyTasksResponse) |  |
| GetSlicerResources | [GetSlicerResourcesRequest](#armory-v1-GetSlicerResourcesRequest) | [GetSlicerResourcesResponse](#armory-v1-GetSlicerResourcesResponse) |  |
| ListUserDevices | [ListUserDevicesRequest](#armory-v1-ListUserDevicesRequest) | [ListUserDevicesResponse](#armory-v1-ListUserDevicesResponse) |  |
| GetPrintStatus | [GetPrintStatusRequest](#armory-v1-GetPrintStatusRequest) | [GetPrintStatusResponse](#armory-v1-GetPrintStatusResponse) |  |
| GetTTCode | [GetTTCodeRequest](#armory-v1-GetTTCodeRequest) | [GetTTCodeResponse](#armory-v1-GetTTCodeResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

