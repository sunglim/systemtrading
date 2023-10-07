# Korea investment APIs

## Development tip

* The server may not respond. Do not `panic()` when response is EOF.
* The server may not response because of network issue. Use [retry pattern](https://learn.microsoft.com/en-us/azure/architecture/patterns/retry)

## TODOs

* Instead of `panic()`, return an error.