# birthday

_ a small utility for exploring the birthday paradox with hash functions_

## Installation
```shell
$ go install github.com/stripedpajamas/birthday/cmd/...
$ birthday -h
```

## Use
Provide a truncation bit-length and a common hashing algorithm (currently supported: md5, sha1, sha256, sha512):

```shell
$ birthday -bits 40 -alg sha256
Searching for collisions in 40-bit truncated sha256 output
Expect to find a collision in <= 2^40/2 ~= 1048576.00 hashes
Collision found after 1323519 hashes generated.
	b7c564487a27614f963e3c59596b107dd98389b6: [8aafd4263c]316d0d90357f396dc27f3da77a22aa924d7a8dd3b6ceb5b28be175
	57018ae5c0c15ceca3c385d3a4534a9acd9a5a9f: [8aafd4263c]62cd7192ba788fdb31c4fc94f9eaba5348cba941d6753815f98265
```

## License
MIT