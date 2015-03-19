# graphite-nozzle

This library consumes events off the Cloud Foundry Firehose, processes them, and then sends them off to a StatsD endpoint.

## Getting Started

An example app is included under the sample directory. To run the app you'll need:

* A user who has access to the Cloud Foundry Firehose (see [here](http://cloudcredo.com/cloud-foundry-firehose-and-friends/) for a tutorial on how to create one).
* A Graphite and StatsD server (see [here](https://github.com/teddyking/graphite-boshrelease) for a Graphite/StatsD BOSH release).
* Golang installed and configured (see [here](https://golang.org/doc/install) for a tutorial on how to do this).
* godep (see [here](https://github.com/tools/godep) for installation instructions).
* The cf cli > 6.7.0 (optional, but useful for retrieving an oauth token that's required by the sample app. It can be downloaded [here](https://github.com/cloudfoundry/cli/releases)).

Once you've met all the prerequisites, you'll need to download the library and install the dependencies:

```
mkdir -p $GOPATH/src/github.com/CloudCredo
cd $GOPATH/src/github.com/CloudCredo
git clone git@github.com:CloudCredo/graphite-nozzle.git
cd graphite-nozzle
godep restore
```

You may need to update the consts defined on lines [18-21](https://github.com/CloudCredo/graphite-nozzle/blob/master/sample/main.go#L18-L21) so that they are configured for your environment. The current values work for a bosh-lite install of CF and Graphite using the 'standard' manifests. Once that's done, build the sample app:

```
go build -o bin/graphite-nozzle sample/main.go
```

Before running the app you'll need to export an environment variabled named `CF_ACCESS_TOKEN`. This should contain an oauth token from the user who can access the Firehose (replace 'admin admin' below with the username and password of your admin user):

```
cf auth admin admin && export CF_ACCESS_TOKEN="$(cf oauth-token | tail -n 1)"
```

Finally, run the app:

```
bin/graphite-nozzle
```

10 seconds later you should be able to see the metrics appearing in your Graphite web UI.

## Metrics Overview

Following is a brief overview of the metrics that graphite-nozzle will extract from the Firehose and send off to Graphite.

### ContainerMetric

CPU, RAM and disk usage metrics for app containers will be sent through to StatsD as a Gauge metric. Note that ContainerMetric Events will not appear on the Firehose by default (at the moment) so you'll need to run a separate app to generate these. There is a sample ContainerMetric-generating app included in the noaa repository [here](https://github.com/cloudfoundry/noaa/tree/master/container_metrics_sample). These metrics appear in the Graphite Web UI under `Graphite.stats.gauges.<statsdPrefix>.apps.<appID>.<containerMetric>.<instanceIndex>`.

### HTTPStartStop

HTTP requests passing through the Cloud Foundry routers get recorded as HTTPStartStop Events. graphite-nozzle takes these events and extracts useful information, such as the response time and status code. These metrics are then sent through to StatsD. The following table gives an overview of the HTTP metrics graphite-nozzle handles:

| Name | Description | StatsD Metric Type |
| ---- | ----------- | ------------------ |
| HttpStartStopResponseTime | HTTP response times in milliseconds | Timer |
| HttpStartStopStatusCodeCount | A count of each HTTP status code | Counter |

For all HTTPStartStop Events, the hostname is extracted from the URI and used in the Metric name. `.` characters are also replaced with `_` characters. This means that, for example, HTTP requests to `http://api.mycf.com/v2/info` will be recorded under `http://api_mycf_com` in the Graphite web UI. This is to avoid polluting the UI with hundreds of endpoints.

### ValueMetric

Any ValueMetric Event that appears on the Firehose will be sent through to StatsD as a Gauge metric. This includes metrics such as numCPUS, numGoRoutines, memoryStats, etc. These metrics appear in the Graphite web UI under `Graphite.stats.gauges.<statsdPrefix>.ops.<Origin>`. Note that the values get sent as int64s so there may be a small loss of precision if the original values are floats.

## Testing

To run the test suite, first make sure you have ginkgo and gomega installed:

```
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
```

Then run `ginkgo -r` from root of this repository.
