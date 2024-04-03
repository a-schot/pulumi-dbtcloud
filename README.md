# dbt Cloud Resource Provider

The dbt Cloud Resource Provider lets you manage dbt Cloud resources.

## Installation

This package is available for several languages/platforms:

- JavaScript/TypeScript: [`@aschot/pulumi-dbtcloud`](https://www.npmjs.com/package/@aschot/pulumi-dbtcloud)
- Python: [`aschot-pulumi-dbtcloud`](https://pypi.org/project/aschot-pulumi-dbtcloud/)
- Go: [`github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud`](https://pkg.go.dev/github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud)
- .NET: [`ASchot.Pulumi.Dbtcloud`](https://www.nuget.org/packages/ASchot.Pulumi.Dbtcloud)

## Provider Binary

The dbt Cloud provider binary is a third party binary. It can be installed using the pulumi plugin command.

```bash
pulumi plugin install resource dbtcloud <version> --server github://api.github.com/a-schot/pulumi-dbtcloud
```

Replace `<version>` with your desired version.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @achot/pulumi_dbtcloud
```

or `yarn`:

```bash
yarn add @achot/pulumi_dbtcloud
```

### Python

To use from Python, install using `pip`:

```bash
pip install aschot-pulumi-dbtcloud
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/a-schot/pulumi-dbtcloud/sdk/go/dbtcloud
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package ASchot.Pulumi.Dbtcloud
```

## Configuration Options

Use `pulumi config set dbtcloud:<option> (--secret)`.

| Option     | Environment Variable   | Required/Optional | Default                                                      | Description                             | 
|------------|------------------------|-------------------|--------------------------------------------------------------|-----------------------------------------|
| `token`    | `DBT_CLOUD_TOKEN`      | Required          |                                                              | The API token for your dbt Cloud user   |
| `accountId`| `DBT_CLOUD_ACCOUNT_ID` | Required          |                                                              | The ID for your dbt Cloud account       |
| `hostUrl`  | `DBT_CLOUD_HOST_URL`   | Optional          | [https://cloud.getdbt.com/api](https://cloud.getdbt.com/api) | The host URL for your dbt Cloud account |
