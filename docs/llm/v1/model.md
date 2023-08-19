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
| id |string|  Unique identifier for the model.  |
| name |string|  Name or description of the model.  |
| versions[] |[ModelVersion](#llm-v1-ModelVersion)|  List of versions for this model.  |




 <!-- end nested messages -->

 <!-- end nested enums -->




<a name="llm-v1-ModelVersion"></a>

### ModelVersion

ModelVersion represents a specific version of a model, including details like the version number and any associated metadata.




| Field | Type | Description |
| ----- | ---- | ----------- |
| version |string|  Version identifier, e.g., "v1.0".  |
| description |string|  Optional description of this version.  |
| metadata |[ModelVersion.MetadataEntry](#llm-v1-ModelVersion-MetadataEntry)|  Key-value pairs for any additional metadata.  |






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

