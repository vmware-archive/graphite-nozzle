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
git clone git@github.com:teddyking/graphite-nozzle.git
cd graphite-nozzle
godep restore
```

You may need to update the consts defined on lines [17-20](https://github.com/teddyking/graphite-nozzle/blob/master/sample/main.go#L17-L20) so that they are configured for your environment. The current values work for a bosh-lite install of CF and Graphite using the 'standard' manifests. Once that's done, build the sample app:

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

At the moment graphite-nozzle only supports the following metrics, however many more are expected to be added soon :D.

| Name | Description | StatsD Metric Type |
| ---- | ----------- | ------------------ |
| HttpStartStopResponseTime | HTTP response times in milliseconds | Timer |
| HttpStartStopStatusCodeCount | A count of each HTTP status code | Counter |
