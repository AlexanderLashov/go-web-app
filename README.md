# Go Web App

This repository contains a simple Go web server.

## Workflows

### CI Workflow (`ci.yml`)

This workflow runs automatically on each push or pull request to the `main` branch. It performs the following steps:

1. Lint the Go code using `golangci-lint`.
2. Perform Static Application Security Testing (SAST).
3. Build the Go binary.

### Release Workflow (`release.yml`)

This workflow is triggered manually upon creating a new release. It performs the following steps:

1. Lint the Go code using `golangci-lint`.
2. Perform Static Application Security Testing (SAST).
3. Build and test the Go binary.
4. Package the binary into a `tar.gz` file.
5. Use `wangyoucao577/go-release-action@v1` to add the `tar.gz` to the release assets.
6. Dispatch a GitHub event to trigger workflows in the second [repository](https://github.com/AlexanderLashov/go-web-app2).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
