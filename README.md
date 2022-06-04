# Github Connector

The [Github](https://github.com) connector allows for the ingestion of webhook events from Github. To setup the connector, use the following [guide](https://docs.github.com/en/developers/webhooks-and-events/webhooks/about-webhooks) and ensure that the URL used for the ingestion ends with `/integrations/github/webhooks`.

## Endpoints

|Url|Methods|Description|
|---|---|---|
|`/integrations/github/webhooks`|POST|Captures a given event into the system (assuming it passes validation and ingestion policies)|
|`/_system/health`|GET|The liveness health check endpoint|
|`/_system/health/ready`|GET|The readiness health check endpoint|

## Configuration

The Github connector looks for the following configuration objects within a Kubernetes cluster:

- ConfigMaps:
  - connector-github-cm
- Secrets:
  - connector-github-secret
  - ingestion-secret

### ConfigMap - `connector-github-cm`

|Key|Description|
|---|---|
|ingestion.uri|The path of the ingestion API, including the `/ingest` suffix. Default: `http://keas-ingestion.keas.svc.cluster.local/ingest`|
|log.level|The log level that should be written to the console. Default: `debug`|
|server.port|The port to listen on. It can be useful to change this for local development. Default: `5000`|

### Secret - `connector-github-secret`

_This secret is required. If the secret does not exist, the readiness checks will fail._

|Key|Description|
|---|---|
|github.webhook.token|The secret used to validate that the incoming requests are indeed coming from Github. This must match what's you set on the Github UI|

### Secret - `ingestion-secret`

_This secret is required. If the secret does not exist, the readiness checks will fail. This secret is often setup by the [Ingestion API](https://github.com/projectkeas/ingestion)._

|Key|Description|
|---|---|
|ingestion.auth.token|The API Key that's used for authentication against the ingestion API|
