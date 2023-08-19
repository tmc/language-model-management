---
title: llm.v1
description: API Specification for the llm.v1 package.
---

<a name="prompt-proto"></a><p align="right"><a href="#top">Top</a></p>

<!-- begin services -->

<!-- begin services -->



<a name="llm-v1-Prompt"></a>

### Prompt

Prompt represents the input provided to a generation, including the template, context, and other details.




| Field | Type | Description |
| ----- | ---- | ----------- |
| id |string|  Unique identifier for the prompt.  |
| template |string|  The templated language or structure of the prompt.  |
| context |bytes|  Context information in binary format, such as retrieval-augmented generation context.  |




 <!-- end nested messages -->

 <!-- end nested enums -->




<a name="llm-v1-Rater"></a>

### Rater

Rater represents details about the entity providing a rating, including qualifications and type (human or model).




| Field | Type | Description |
| ----- | ---- | ----------- |
| id |string|  Unique identifier for the rater.  |
| type |string|  Type of rater, e.g., "human" or "model".  |
| qualification |string|  Qualification details, use-case dependent.  |




 <!-- end nested messages -->

 <!-- end nested enums -->


 <!-- end messages -->

<!-- begin file-level enums -->
 <!-- end file-level enums -->

<!-- begin file-level extensions -->
 <!-- end file-level extensions -->

