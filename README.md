# New Relic Infrastructure Integration for Zookeeper
New Relic Infrastructure Integration for Kafka captures critical performance metrics and inventory reported by Kafka server.

## Requirements
* Apache Zookeeper 3.4 or above


## Installation
* Download an archive file for the Zookeeper Integration
* Place the executables under `bin` directory and the definition file `zookeeper-definition.yml` in `/var/db/newrelic-infra/newrelic-integrations`
* Set execution permissions for the binary files `nr-zookeeper` and `nrjmx` (if required)
* Place the integration configuration file `zookeeper-config.yml.sample` in `/etc/newrelic-infra/integrations.d` and update its values.

## Usage
This is the description about how to run the Zookeeper Integration with New Relic Infrastructure agent, so it is required to have the agent installed (see [agent installation](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux)).

In order to use the Kafka Integration it is required to configure `zookeeper-config.yml.sample` file. Firstly, rename the file to `zookeeper-config.yml`. Then, depending on your needs, specify all instances that you want to monitor. Once this is done, restart the Infrastructure agent.

You can view your data in Insights by creating your own custom NRQL queries. To
do so use **ZookeeperSample** event types.

## Integration development usage
Assuming that you have source code you can build and run the Zookeeper Integration locally.
* Go to directory of the Zookeeper Integration and build it
```bash
$ make
```
* The command above will execute tests for the Kafka Integration and build an executable file called `nr-zookeeper` in `bin` directory.
```bash
$ ./bin/nr-zookeeper -port <zookeeper port>
```
* If you want to know more about usage of `./bin/nr-zookeeper` check
```bash
$ ./bin/nr-zookeeper --help
```

For managing external dependencies [govendor tool](https://github.com/kardianos/govendor) is used. It is required to lock all external dependencies to specific version (if possible) into vendor directory.
