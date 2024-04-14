# RanHash

A simple random hash generator for go.
For any backend I end up making in go, I always end writing some sort of a
random hash generator, so this is to reduce that boiler code in my code bases

This is quite a simple yet useful implementation.
Here are the main functions that you can use:

- `Generate(length int, pool string) string`: This is the most generic function.
All you need to do is to send in a pool, with a given length
The package itself also gives you 3 predefined pools: NumberPool, AlphabetPool,
HashPool (combination of two).

- `GenerateRandomString(length int) string`: Generates a random string
- `GenerateRandomNumber(length int) string`: Generates a random number
- `RanHash(length int) string`: Generates a random hash with a given length. Has both alphabets and numbers

Just simple. Hope you find it useful

## License

This above package is licensed under the MIT License, the details of which can be found in the 
[LICENSE](LICENSE) file
