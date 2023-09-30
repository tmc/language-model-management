---
title: llm.v1
description: API Specification for the llm.v1 package.
---

<a name="model-proto"></a><p align="right"><a href="#top">Top</a></p>

<!-- begin services -->

<!-- begin services -->



<a name="llm-v1-Model"></a>

### Model

Model represents a high-level abstraction for a language model.




| Field | Type | Description |
| ----- | ---- | ----------- |
| id |string| Unique identifier for the model.   |
| name |string| Descriptive name or label of the model.   |
| versions[] |[ModelVersion](#llm-v1-ModelVersion)| List of versions associated with this model.   |




 <!-- end nested messages -->

 <!-- end nested enums -->




<a name="llm-v1-ModelVersion"></a>

### ModelVersion

ModelVersion represents a specific version of a model.




| Field | Type | Description |
| ----- | ---- | ----------- |
| version |string| Unique identifier for this version, e.g., "v1.0".   |
| description |string| Optional description providing more details about this version.   |
| metadata |[ModelVersion.MetadataEntry](#llm-v1-ModelVersion-MetadataEntry)| Additional metadata associated with this version represented as key-value pairs.   |
| created_at |Timestamp| Timestamp indicating when this model version was created.   |
| updated_at |Timestamp| Timestamp indicating the last time this model version was updated.   |






<a name="llm-v1-ModelVersion-MetadataEntry"></a>

### MetadataEntry





| Field | Type | Description |
| ----- | ---- | ----------- |
| key |string|   |
| value |string|   |




 <!-- end nested messages -->

 <!-- end nested enums -->


 <!-- end nested messages -->

 <!-- end nested enums -->


 <!-- end messages -->

<!-- begin file-level enums -->
 <!-- end file-level enums -->

<!-- begin file-level extensions -->
 <!-- end file-level extensions -->

