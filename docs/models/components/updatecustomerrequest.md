# UpdateCustomerRequest


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Name`                                                         | **string*                                                      | :heavy_minus_sign:                                             | The name of the customer.                                      | Victor Solv                                                    |
| `Email`                                                        | **string*                                                      | :heavy_minus_sign:                                             | The customer's email.                                          | vsolv@scalarsolutions.com                                      |
| `Phone`                                                        | **string*                                                      | :heavy_minus_sign:                                             | The customer's phone number.                                   | (160) 217-6634                                                 |
| `Metadata`                                                     | **string*                                                      | :heavy_minus_sign:                                             | JSON object of metadata related to the customer being updated. | {<br/>"Language Preference": "English"<br/>}                   |