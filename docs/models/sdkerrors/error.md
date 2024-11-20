# Error


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Category`                                                | [*sdkerrors.Category](../../models/sdkerrors/category.md) | :heavy_minus_sign:                                        | N/A                                                       | user_error                                                |
| `Code`                                                    | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       | MA_00001                                                  |
| `Message`                                                 | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       | missing parameters                                        |
| `Details`                                                 | [*sdkerrors.Details](../../models/sdkerrors/details.md)   | :heavy_minus_sign:                                        | N/A                                                       | {<br/>"parameters": [<br/>"uri",<br/>"analysis_type"<br/>]<br/>} |
| `RequestID`                                               | **string*                                                 | :heavy_minus_sign:                                        | N/A                                                       | 17c3b70c5096df0e77e838323abb7029                          |