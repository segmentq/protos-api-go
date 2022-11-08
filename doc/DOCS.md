# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [segment.proto](#segment-proto)
    - [DeleteSegmentRequest](#api-DeleteSegmentRequest)
    - [DeleteSegmentResponse](#api-DeleteSegmentResponse)
    - [GetSegmentRequest](#api-GetSegmentRequest)
    - [GetSegmentResponse](#api-GetSegmentResponse)
    - [PutSegmentRequest](#api-PutSegmentRequest)
    - [PutSegmentResponse](#api-PutSegmentResponse)
    - [Segment](#api-Segment)
    - [SegmentField](#api-SegmentField)
    - [SegmentFieldBlob](#api-SegmentFieldBlob)
    - [SegmentFieldBool](#api-SegmentFieldBool)
    - [SegmentFieldFloat](#api-SegmentFieldFloat)
    - [SegmentFieldGeoPoint](#api-SegmentFieldGeoPoint)
    - [SegmentFieldGeoRect](#api-SegmentFieldGeoRect)
    - [SegmentFieldInt](#api-SegmentFieldInt)
    - [SegmentFieldRangeFloat](#api-SegmentFieldRangeFloat)
    - [SegmentFieldRangeInt](#api-SegmentFieldRangeInt)
    - [SegmentFieldRepeatedBlob](#api-SegmentFieldRepeatedBlob)
    - [SegmentFieldRepeatedBool](#api-SegmentFieldRepeatedBool)
    - [SegmentFieldRepeatedFloat](#api-SegmentFieldRepeatedFloat)
    - [SegmentFieldRepeatedGeoPoint](#api-SegmentFieldRepeatedGeoPoint)
    - [SegmentFieldRepeatedGeoRect](#api-SegmentFieldRepeatedGeoRect)
    - [SegmentFieldRepeatedInt](#api-SegmentFieldRepeatedInt)
    - [SegmentFieldRepeatedRangeFloat](#api-SegmentFieldRepeatedRangeFloat)
    - [SegmentFieldRepeatedRangeInt](#api-SegmentFieldRepeatedRangeInt)
    - [SegmentFieldRepeatedString](#api-SegmentFieldRepeatedString)
    - [SegmentFieldRepeatedUInt](#api-SegmentFieldRepeatedUInt)
    - [SegmentFieldString](#api-SegmentFieldString)
    - [SegmentFieldUInt](#api-SegmentFieldUInt)
  
- [api.proto](#api-proto)
    - [IndexService](#api-IndexService)
    - [SegmentService](#api-SegmentService)
  
- [field.proto](#field-proto)
    - [FieldDefinition](#api-FieldDefinition)
  
    - [GeoType](#api-GeoType)
    - [ScalarType](#api-ScalarType)
  
- [index.proto](#index-proto)
    - [AddIndexRequest](#api-AddIndexRequest)
    - [AddIndexResponse](#api-AddIndexResponse)
    - [DeleteIndexRequest](#api-DeleteIndexRequest)
    - [DeleteIndexResponse](#api-DeleteIndexResponse)
    - [DescribeIndexRequest](#api-DescribeIndexRequest)
    - [DescribeIndexResponse](#api-DescribeIndexResponse)
    - [IndexDefinition](#api-IndexDefinition)
    - [ListIndexesRequest](#api-ListIndexesRequest)
    - [ListIndexesResponse](#api-ListIndexesResponse)
  
    - [Status](#api-Status)
  
- [lookup.proto](#lookup-proto)
    - [Lookup](#api-Lookup)
    - [LookupField](#api-LookupField)
    - [LookupSegmentKeyRequest](#api-LookupSegmentKeyRequest)
    - [LookupSegmentKeyResponse](#api-LookupSegmentKeyResponse)
    - [LookupSegmentRequest](#api-LookupSegmentRequest)
    - [LookupSegmentResponse](#api-LookupSegmentResponse)
  
- [Scalar Value Types](#scalar-value-types)



<a name="segment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## segment.proto



<a name="api-DeleteSegmentRequest"></a>

### DeleteSegmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [string](#string) |  |  |
| key | [SegmentField](#api-SegmentField) |  | Lookup on index primary key |






<a name="api-DeleteSegmentResponse"></a>

### DeleteSegmentResponse
Empty






<a name="api-GetSegmentRequest"></a>

### GetSegmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [string](#string) |  |  |
| key | [SegmentField](#api-SegmentField) |  | Lookup on index primary key |






<a name="api-GetSegmentResponse"></a>

### GetSegmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| segment | [Segment](#api-Segment) |  |  |






<a name="api-PutSegmentRequest"></a>

### PutSegmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [string](#string) |  |  |
| segment | [Segment](#api-Segment) |  |  |






<a name="api-PutSegmentResponse"></a>

### PutSegmentResponse
Empty






<a name="api-Segment"></a>

### Segment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [SegmentField](#api-SegmentField) | repeated |  |






<a name="api-SegmentField"></a>

### SegmentField



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| string_value | [SegmentFieldString](#api-SegmentFieldString) |  |  |
| repeated_string_value | [SegmentFieldRepeatedString](#api-SegmentFieldRepeatedString) |  |  |
| int_value | [SegmentFieldInt](#api-SegmentFieldInt) |  |  |
| repeated_int_value | [SegmentFieldRepeatedInt](#api-SegmentFieldRepeatedInt) |  |  |
| uint_value | [SegmentFieldUInt](#api-SegmentFieldUInt) |  |  |
| repeated_uint_value | [SegmentFieldRepeatedUInt](#api-SegmentFieldRepeatedUInt) |  |  |
| float_value | [SegmentFieldFloat](#api-SegmentFieldFloat) |  |  |
| repeated_float_value | [SegmentFieldRepeatedFloat](#api-SegmentFieldRepeatedFloat) |  |  |
| bool_value | [SegmentFieldBool](#api-SegmentFieldBool) |  |  |
| repeated_bool_value | [SegmentFieldRepeatedBool](#api-SegmentFieldRepeatedBool) |  |  |
| blob_value | [SegmentFieldBlob](#api-SegmentFieldBlob) |  |  |
| repeated_blob_value | [SegmentFieldRepeatedBlob](#api-SegmentFieldRepeatedBlob) |  |  |
| range_int_value | [SegmentFieldRangeInt](#api-SegmentFieldRangeInt) |  |  |
| repeated_range_int_value | [SegmentFieldRepeatedRangeInt](#api-SegmentFieldRepeatedRangeInt) |  |  |
| range_float_value | [SegmentFieldRangeFloat](#api-SegmentFieldRangeFloat) |  |  |
| repeated_range_float_value | [SegmentFieldRepeatedRangeFloat](#api-SegmentFieldRepeatedRangeFloat) |  |  |
| geo_point_value | [SegmentFieldGeoPoint](#api-SegmentFieldGeoPoint) |  |  |
| repeated_geo_point_value | [SegmentFieldRepeatedGeoPoint](#api-SegmentFieldRepeatedGeoPoint) |  |  |
| geo_rect_value | [SegmentFieldGeoRect](#api-SegmentFieldGeoRect) |  |  |
| repeated_geo_rect_value | [SegmentFieldRepeatedGeoRect](#api-SegmentFieldRepeatedGeoRect) |  |  |






<a name="api-SegmentFieldBlob"></a>

### SegmentFieldBlob



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="api-SegmentFieldBool"></a>

### SegmentFieldBool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bool](#bool) |  |  |






<a name="api-SegmentFieldFloat"></a>

### SegmentFieldFloat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  |  |






<a name="api-SegmentFieldGeoPoint"></a>

### SegmentFieldGeoPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| x | [double](#double) |  |  |
| y | [double](#double) |  |  |






<a name="api-SegmentFieldGeoRect"></a>

### SegmentFieldGeoRect



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| top_left | [SegmentFieldGeoPoint](#api-SegmentFieldGeoPoint) |  |  |
| bottom_right | [SegmentFieldGeoPoint](#api-SegmentFieldGeoPoint) |  |  |






<a name="api-SegmentFieldInt"></a>

### SegmentFieldInt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  |  |






<a name="api-SegmentFieldRangeFloat"></a>

### SegmentFieldRangeFloat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min | [double](#double) |  |  |
| max | [double](#double) |  |  |






<a name="api-SegmentFieldRangeInt"></a>

### SegmentFieldRangeInt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min | [int64](#int64) |  |  |
| max | [int64](#int64) |  |  |






<a name="api-SegmentFieldRepeatedBlob"></a>

### SegmentFieldRepeatedBlob



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) | repeated |  |






<a name="api-SegmentFieldRepeatedBool"></a>

### SegmentFieldRepeatedBool



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bool](#bool) | repeated |  |






<a name="api-SegmentFieldRepeatedFloat"></a>

### SegmentFieldRepeatedFloat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) | repeated |  |






<a name="api-SegmentFieldRepeatedGeoPoint"></a>

### SegmentFieldRepeatedGeoPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [SegmentFieldGeoPoint](#api-SegmentFieldGeoPoint) | repeated |  |






<a name="api-SegmentFieldRepeatedGeoRect"></a>

### SegmentFieldRepeatedGeoRect



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [SegmentFieldGeoRect](#api-SegmentFieldGeoRect) | repeated |  |






<a name="api-SegmentFieldRepeatedInt"></a>

### SegmentFieldRepeatedInt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) | repeated |  |






<a name="api-SegmentFieldRepeatedRangeFloat"></a>

### SegmentFieldRepeatedRangeFloat



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [SegmentFieldRangeFloat](#api-SegmentFieldRangeFloat) | repeated |  |






<a name="api-SegmentFieldRepeatedRangeInt"></a>

### SegmentFieldRepeatedRangeInt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [SegmentFieldRangeInt](#api-SegmentFieldRangeInt) | repeated |  |






<a name="api-SegmentFieldRepeatedString"></a>

### SegmentFieldRepeatedString



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) | repeated |  |






<a name="api-SegmentFieldRepeatedUInt"></a>

### SegmentFieldRepeatedUInt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [uint64](#uint64) | repeated |  |






<a name="api-SegmentFieldString"></a>

### SegmentFieldString



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  |  |






<a name="api-SegmentFieldUInt"></a>

### SegmentFieldUInt



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [uint64](#uint64) |  |  |





 

 

 

 



<a name="api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto


 

 

 


<a name="api-IndexService"></a>

### IndexService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| AddIndex | [AddIndexRequest](#api-AddIndexRequest) | [AddIndexResponse](#api-AddIndexResponse) |  |
| DescribeIndex | [DescribeIndexRequest](#api-DescribeIndexRequest) | [DescribeIndexResponse](#api-DescribeIndexResponse) |  |
| DeleteIndex | [DeleteIndexRequest](#api-DeleteIndexRequest) | [DeleteIndexResponse](#api-DeleteIndexResponse) | rpc AlterIndex(AlterIndexRequest) returns (AlterIndexResponse); |
| ListIndexes | [ListIndexesRequest](#api-ListIndexesRequest) | [ListIndexesResponse](#api-ListIndexesResponse) |  |


<a name="api-SegmentService"></a>

### SegmentService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PutSegment | [PutSegmentRequest](#api-PutSegmentRequest) | [PutSegmentResponse](#api-PutSegmentResponse) |  |
| GetSegment | [GetSegmentRequest](#api-GetSegmentRequest) | [GetSegmentResponse](#api-GetSegmentResponse) |  |
| DeleteSegment | [DeleteSegmentRequest](#api-DeleteSegmentRequest) | [DeleteSegmentResponse](#api-DeleteSegmentResponse) |  |
| LookupSegment | [LookupSegmentRequest](#api-LookupSegmentRequest) | [LookupSegmentResponse](#api-LookupSegmentResponse) |  |
| LookupSegmentKey | [LookupSegmentKeyRequest](#api-LookupSegmentKeyRequest) | [LookupSegmentKeyResponse](#api-LookupSegmentKeyResponse) |  |

 



<a name="field-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## field.proto



<a name="api-FieldDefinition"></a>

### FieldDefinition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| scalar | [ScalarType](#api-ScalarType) |  |  |
| geo | [GeoType](#api-GeoType) |  |  |
| is_primary | [bool](#bool) |  |  |
| repeated | [bool](#bool) |  | Allow arrays |
| fields | [FieldDefinition](#api-FieldDefinition) | repeated | Allow nested structure |





 


<a name="api-GeoType"></a>

### GeoType


| Name | Number | Description |
| ---- | ------ | ----------- |
| DATA_TYPE_RANGE | 0 | RANGEs |
| DATA_TYPE_RANGE_INT | 1 |  |
| DATA_TYPE_RANGE_FLOAT | 2 |  |
| DATA_TYPE_GEO | 10 | GEOs |
| DATA_TYPE_GEO_RECT | 11 |  |
| DATA_TYPE_GEO_POINT | 12 |  |



<a name="api-ScalarType"></a>

### ScalarType


| Name | Number | Description |
| ---- | ------ | ----------- |
| DATA_TYPE_UNDEFINED | 0 |  |
| DATA_TYPE_STRING | 10 | STRINGs |
| DATA_TYPE_INT | 20 | INTs |
| DATA_TYPE_INT8 | 21 |  |
| DATA_TYPE_INT16 | 22 |  |
| DATA_TYPE_INT32 | 23 |  |
| DATA_TYPE_INT64 | 24 |  |
| DATA_TYPE_UINT | 30 | UINTs |
| DATA_TYPE_UINT8 | 31 |  |
| DATA_TYPE_UINT16 | 32 |  |
| DATA_TYPE_UINT32 | 33 |  |
| DATA_TYPE_UINT64 | 34 |  |
| DATA_TYPE_FLOAT | 40 | FLOATs |
| DATA_TYPE_FLOAT32 | 41 |  |
| DATA_TYPE_FLOAT64 | 42 |  |
| DATA_TYPE_BOOL | 50 | BOOL |
| DATA_TYPE_BLOB | 60 | BLOB |


 

 

 



<a name="index-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## index.proto



<a name="api-AddIndexRequest"></a>

### AddIndexRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [IndexDefinition](#api-IndexDefinition) |  |  |






<a name="api-AddIndexResponse"></a>

### AddIndexResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  |  |






<a name="api-DeleteIndexRequest"></a>

### DeleteIndexRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="api-DeleteIndexResponse"></a>

### DeleteIndexResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [IndexDefinition](#api-IndexDefinition) |  |  |






<a name="api-DescribeIndexRequest"></a>

### DescribeIndexRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="api-DescribeIndexResponse"></a>

### DescribeIndexResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [IndexDefinition](#api-IndexDefinition) |  |  |






<a name="api-IndexDefinition"></a>

### IndexDefinition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| status | [Status](#api-Status) |  |  |
| fields | [FieldDefinition](#api-FieldDefinition) | repeated |  |
| created | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-ListIndexesRequest"></a>

### ListIndexesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pattern | [string](#string) |  |  |
| created_after | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| created_before | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_after | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_before | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-ListIndexesResponse"></a>

### ListIndexesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| results | [DescribeIndexResponse](#api-DescribeIndexResponse) | repeated |  |





 


<a name="api-Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| INDEX_STATUS_UNKNOWN | 0 |  |
| INDEX_STATUS_CREATING | 5 |  |
| INDEX_STATUS_ACTIVE | 10 |  |
| INDEX_STATUS_UPDATING | 15 |  |
| INDEX_STATUS_DELETING | 20 |  |


 

 

 



<a name="lookup-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## lookup.proto



<a name="api-Lookup"></a>

### Lookup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [LookupField](#api-LookupField) | repeated |  |






<a name="api-LookupField"></a>

### LookupField



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| string_value | [SegmentFieldString](#api-SegmentFieldString) |  |  |
| repeated_string_value | [SegmentFieldRepeatedString](#api-SegmentFieldRepeatedString) |  |  |
| int_value | [SegmentFieldInt](#api-SegmentFieldInt) |  |  |
| repeated_int_value | [SegmentFieldRepeatedInt](#api-SegmentFieldRepeatedInt) |  |  |
| uint_value | [SegmentFieldUInt](#api-SegmentFieldUInt) |  |  |
| repeated_uint_value | [SegmentFieldRepeatedUInt](#api-SegmentFieldRepeatedUInt) |  |  |
| float_value | [SegmentFieldFloat](#api-SegmentFieldFloat) |  |  |
| repeated_float_value | [SegmentFieldRepeatedFloat](#api-SegmentFieldRepeatedFloat) |  |  |
| bool_value | [SegmentFieldBool](#api-SegmentFieldBool) |  |  |
| repeated_bool_value | [SegmentFieldRepeatedBool](#api-SegmentFieldRepeatedBool) |  |  |
| range_int_value | [SegmentFieldRangeInt](#api-SegmentFieldRangeInt) |  |  |
| repeated_range_int_value | [SegmentFieldRepeatedRangeInt](#api-SegmentFieldRepeatedRangeInt) |  |  |
| range_float_value | [SegmentFieldRangeFloat](#api-SegmentFieldRangeFloat) |  |  |
| repeated_range_float_value | [SegmentFieldRepeatedRangeFloat](#api-SegmentFieldRepeatedRangeFloat) |  |  |
| geo_point_value | [SegmentFieldGeoPoint](#api-SegmentFieldGeoPoint) |  |  |
| repeated_geo_point_value | [SegmentFieldRepeatedGeoPoint](#api-SegmentFieldRepeatedGeoPoint) |  |  |
| geo_rect_value | [SegmentFieldGeoRect](#api-SegmentFieldGeoRect) |  |  |
| repeated_geo_rect_value | [SegmentFieldRepeatedGeoRect](#api-SegmentFieldRepeatedGeoRect) |  |  |






<a name="api-LookupSegmentKeyRequest"></a>

### LookupSegmentKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [string](#string) |  |  |
| lookup | [Lookup](#api-Lookup) |  |  |






<a name="api-LookupSegmentKeyResponse"></a>

### LookupSegmentKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| results | [SegmentField](#api-SegmentField) | repeated |  |






<a name="api-LookupSegmentRequest"></a>

### LookupSegmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [string](#string) |  |  |
| lookup | [Lookup](#api-Lookup) |  |  |






<a name="api-LookupSegmentResponse"></a>

### LookupSegmentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| results | [Segment](#api-Segment) | repeated |  |





 

 

 

 



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

