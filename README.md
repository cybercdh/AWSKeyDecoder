# AWSKeyDecoder

AWSKeyDecoder converts AWS Key Ids to AWS Account Ids en masse.

## Installation

Assuming you have [Go](https://go.dev/doc/install) installed:

```bash
go install github.com/cybercdh/AWSKeyDecoder@latest
```

## Usage

```bash
echo AKIAIOSFODNN7EXAMPLE | AWSKeyDecoder
123456789012
```

or 
```bash
cat keys.txt | AWSKeyDecoder -v
AWS Key ID: AKIAIOSFODNN7EXAMPLE -> Account ID: 123456789012
AWS Key ID: AKIAI44QH8DHBEXAMPLE -> Account ID: 234567890123
AWS Key ID: AKIAI7ZAFDUSN7EXAMPLE -> Account ID: 345678901234
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Thanks

This code is based on the Python version which was blogged about [here](https://medium.com/@TalBeerySec/a-short-note-on-aws-key-id-f88cc4317489). This Go version uses concurrency to process large volumes of these faster.

## License

[MIT](https://choosealicense.com/licenses/mit/)