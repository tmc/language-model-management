---
title: llm.v1
description: API Specification for the llm.v1 package.
---

<a name="output-proto"></a><p align="right"><a href="#top">Top</a></p>

<!-- begin services -->

<!-- begin services -->



<a name="llm-v1-Output"></a>

### Output

Output represents a single output from a generation.




| Field | Type | Description |
| ----- | ---- | ----------- |
| id |string| Unique identifier for the output.   |
| content |string| The actual content produced by the model.   |
| rating |[OutputRating](#llm-v1-OutputRating)| Associated rating for this output, if available.   |
| created_at |Timestamp| Timestamp indicating when the output was created.   |




 <!-- end nested messages -->

 <!-- end nested enums -->




<a name="llm-v1-OutputRating"></a>

### OutputRating

OutputRating represents the evaluation or assessment of an output.




| Field | Type | Description |
| ----- | ---- | ----------- |
| score |float| Numeric score representing the quality or other rating dimension.   |
| rater |[Rater](./prompt.md#llm-v1-Rater)| Details about the entity providing the rating.   |
| feedback |string| Additional comments or feedback provided during the rating.   |
| rated_at |Timestamp| Timestamp indicating when the rating was provided.   |




 <!-- end nested messages -->

 <!-- end nested enums -->


 <!-- end messages -->

<!-- begin file-level enums -->
 <!-- end file-level enums -->

<!-- begin file-level extensions -->
 <!-- end file-level extensions -->

