# dbt Cloud Resource Provider

The dbt Cloud Resource Provider lets you manage dbt Cloud resources.

## Installing

This package is available for several languages/platforms:

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
pip install aschot_pulumi_dbtcloud
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

## Configuration

The following configuration points are available for the dbt Cloud provider:

- (required) `dbtcloud:accountId` - The Account ID for your dbt Cloud (can be set using environment variable: `DBT_CLOUD_ACCOUNT_ID`)
- (required) `dbtcloud:token` - The API Key for your dbt Cloud (can be set using environment variable: `DBT_CLOUD_TOKEN`)
- (optional) `dbtcloud:hostUrl` - The host URL for your dbt Cloud (can be set using environment variable: `DBT_CLOUD_HOST_URL`). Defaults to: [https://cloud.getdbt.com/api](https://cloud.getdbt.com/api)

<!-- ## Reference -->
<!--  -->
<!-- For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/foo/api-docs/). -->
