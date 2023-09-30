---
title: llm.v1
description: API Specification for the llm.v1 package.
---

<a name="prompt-proto"></a><p align="right"><a href="#top">Top</a></p>

<!-- begin services -->

<!-- begin services -->



<a name="llm-v1-Prompt"></a>

### Prompt

Prompt represents the input provided to a generation.




| Field | Type | Description |
| ----- | ---- | ----------- |
| id |string| Unique identifier for the prompt.   |
| template |string| Templated language or structure of the prompt.   |
| context |[Prompt.ContextEntry](#llm-v1-Prompt-ContextEntry)| Contextual information supplied with the prompt.   |
| created_at |Timestamp| Timestamp indicating when the prompt was created.   |
| metadata[] |Any| Metadata.   |






<a name="llm-v1-Prompt-ContextEntry"></a>

### ContextEntry





| Field | Type | Description |
| ----- | ---- | ----------- |
| key |string|   |
| value |string|   |




 <!-- end nested messages -->

 <!-- end nested enums -->


 <!-- end nested messages -->

 <!-- end nested enums -->




<a name="llm-v1-Rater"></a>

### Rater

Rater represents the entity providing a rating, either human or model-based.




| Field | Type | Description |
| ----- | ---- | ----------- |
| id |string| Unique identifier for the rater.   |
| type |[Rater.RaterType](#llm-v1-Rater-RaterType)| Type of the rater.   |
| qualification |string| Qualification details of the rater, which are use-case dependent.   |
| qualified_at |Timestamp| Timestamp indicating when the rater was qualified or registered.   |
| metadata[] |Any| Metadata.   |




 <!-- end nested messages -->



<a name="llm-v1-Rater-RaterType"></a>

### Rater.RaterType



| Name | Number | Description |
| ---- | ------ | ----------- |
| RATER_TYPE_UNSPECIFIED | 0 | Unspecified rater type.   |
| RATER_TYPE_HUMAN | 1 | Human rater.   |
| RATER_TYPE_MODEL | 2 | Model-based rater.   |


 <!-- end nested enums -->


 <!-- end messages -->

<!-- begin file-level enums -->
 <!-- end file-level enums -->

<!-- begin file-level extensions -->
 <!-- end file-level extensions -->

