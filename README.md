[![Community Project header](https://github.com/newrelic/open-source-office/raw/master/examples/categories/images/Community_Project.png)](https://github.com/newrelic/open-source-office/blob/master/examples/categories/index.md#community-project)

# New Relic infrastructure integration for Zookeeper

The New Relic infrastructure integration for Zookeeper captures critical performance metrics reported by Zookeeper nodes.

## Requirements

* Apache Zookeeper 3.4 or above

## Installation 

* Download the Zookeeper Integration from [Releases](https://github.com/newrelic/nri-zookeeper/releases).
* Place the executable `nr-zookeeper` in `/var/db/newrelic-infra/newrelic-integrations/bin/`.
* Place the definition file `zookeeper-definition.yml` in `/var/db/newrelic-infra/newrelic-integrations`.
* Set execution permissions for the binary files `nr-zookeeper`
* Place the integration configuration file `zookeeper-config.yml.sample` in `/etc/newrelic-infra/integrations.d` and update its values.
* Verify the path to  `nc` utility in the  `zookeeper-config.yml`  

  ```
      arguments:
        host: localhost
        port: 2181
        cmd: /bin/nc  <-- this PATH varies per OS
  ```

## Usage

The Zookeeper Integration runs with New Relic Infrastructure agent (see [agent installation](https://docs.newrelic.com/docs/infrastructure/new-relic-infrastructure/installation/install-infrastructure-linux)).

To use the Zookeeper Integration you must edit the `zookeeper-config.yml.sample` file. Rename the file to `zookeeper-config.yml`; then, depending on your needs, specify all instances that you want to monitor. Once this is done, restart the Infrastructure agent.

Data should start flowing into your New Relic account. See [Understand and use data from Infrastructure integrations](https://docs.newrelic.com/docs/integrations/infrastructure-integrations/get-started/understand-use-data-infrastructure-integrations).

To validate the metrics collected from Zookeeper, you can run the command below.  

```
$ echo mntr | nc localhost 2181
zk_version	3.4.10-39d3a4f269333c922ed3db283be479f9deacaa0f, built on 03/23/2017 10:13 GMT
zk_avg_latency	0
zk_max_latency	0
zk_min_latency	0
zk_packets_received	14
zk_packets_sent	13
zk_num_alive_connections	1
zk_outstanding_requests	0
zk_server_state	standalone
zk_znode_count	4
zk_watch_count	0
zk_ephemerals_count	0
zk_approximate_data_size	27
zk_open_file_descriptor_count	110
zk_max_file_descriptor_count	1048576
```

## Building

Golang is required to build the integration. We recommend Golang 1.11 or higher.

After cloning this repository, go to the directory of the Zookeeper integration and build it:

```bash
$ make
```

The command above executes the tests for the Zookeeper integration and builds an executable file called `nr-zookeeper` under the `bin` directory. 

To start the integration, run `nr-zookeeper`:

```bash
$ ./bin/nr-zookeeper
```

If you want to know more about usage of `./bin/nr-zookeeper`, pass the `-help` parameter:

```bash
$ ./bin/nr-zookeeper -help
```

External dependencies are managed through the [govendor tool](https://github.com/kardianos/govendor). Locking all external dependencies to a specific version (if possible) into the vendor directory is required.

## Testing

To run the tests execute:

```bash
$ make test
```

## Support

### Disclaimer

New Relic has open-sourced this integration to enable monitoring of this technology. This integration is provided AS-IS WITHOUT WARRANTY OR SUPPORT, although you can report issues and contribute to this integration via GitHub. Support for this integration is available with an [Expert Services subscription](https://newrelic.com/expertservices).

## Privacy

At New Relic we take your privacy and the security of your information seriously, and are committed to protecting your information. We must emphasize the importance of not sharing personal data in public forums, and ask all users to scrub logs and diagnostic information for sensitive information, whether personal, proprietary, or otherwise.

We define “Personal Data” as any information relating to an identified or identifiable individual, including, for example, your name, phone number, post code or zip code, Device ID, IP address and email address.

Review [New Relic’s General Data Privacy Notice](https://newrelic.com/termsandconditions/privacy) for more information.

## Contributing

We encourage your contributions to improve the Zookeeper integration! Keep in mind when you submit your pull request, you'll need to sign the CLA via the click-through using CLA-Assistant. You only have to sign the CLA one time per project.

If you have any questions, or to execute our corporate CLA, required if your contribution is on behalf of a company,  please drop us an email at opensource@newrelic.com.

**A note about vulnerabilities**

As noted in our [security policy](/SECURITY.md), New Relic is committed to the privacy and security of our customers and their data. We believe that providing coordinated disclosure by security researchers and engaging with the security community are important means to achieve our security goals.

If you believe you have found a security vulnerability in this project or any of New Relic's products or websites, we welcome and greatly appreciate you reporting it to New Relic through [HackerOne](https://hackerone.com/newrelic).

If you would like to contribute to this project, please review [these guidelines](./CONTRIBUTING.md).

To all contributors, we thank you!  Without your contribution, this project would not be what it is today.

## License
nri-zookeeper is licensed under the [Apache 2.0](http://apache.org/licenses/LICENSE-2.0.txt) License.
