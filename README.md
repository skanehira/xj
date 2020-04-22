# xj
`xj` is convert xml to JSON.

## Installtion
```sh
$ git clone https://github.com/skanehira/xj
$ cd xj && go install
```

## Usage
```sh
$ xj testdata/test.xml | jq
{
  "CATALOG": {
    "CD": [
      {
        "PRICE": "10.90",
        "YEAR": "1985",
        "TITLE": "Empire Burlesque",
        "ARTIST": "Bob Dylan",
        "COUNTRY": "USA",
        "COMPANY": "Columbia"
      },
      {
        "PRICE": "9.90",
        "YEAR": "1988",
        "TITLE": "Hide your heart",
        "ARTIST": "Bonnie Tyler",
        "COUNTRY": "UK",
        "COMPANY": "CBS Records"
      }
    ]
  }
}
```

## Author
skanehira
