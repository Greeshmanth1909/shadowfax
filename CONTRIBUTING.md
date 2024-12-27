# Contributing
Thank you for contributing to the developent of Shadowfax!
All contributions follow the traditional github pull request cycle. Please make sure you've done the following before submitting a PR:
- Format all written code by running `go fmt ./...` at the root of the project
- Run unit tests with `go test ./...` at the root of the project

Only submit a PR if both of the above mentioned steps are performed.

# testing
- All unit tests must make sure `util.InitAll()` is called before starting. the results will be inconsistent if this is not followed
- Uncomment `checkboard` to enter debug mode
