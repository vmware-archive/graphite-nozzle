# graphite-nozzle CI

This directory contains configuration for a [concourse](http://concourse.ci) pipeline.

## Useful commands

```
# Configuring the pipeline

fly -t vagrant sp -c ci/pipeline.yml --load-vars-from ci/secrets.yml -p graphite-nozzle

# Running a one-off task

fly -t vagrant execute -c ci/unit.yml -i graphite-nozzle-src=.

# Hijacking a recent container

fly -t vagrant hijack
```
