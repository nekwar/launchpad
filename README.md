# Mirantis Launchpad CLI Tool

> The Next Generation UCP Installer & Lifecycle Management Tool

The purpose of `launchpad` is to provide amazing new user experience for anyone interested in getting started with UCP product. It will simplify the complex installation process and provides "from zero to hero" experience in less than 5mins for IT admin / DevOps people who are experienced with various command line tools and cloud technologies. In addition, it'll provide functionality to upgrade existing UCP clusters to new versions with no downtime or service interruptions (high availability clusters). In the future, more functionality may be added.

See the public Github repo for getting started instructions, documentation and more at https://github.com/mirantis/launchpad.

## Background

Based on brainstorming session, we decided to spike the tool with an idea to prove it can be done and to gather some more ideas from the field & sales organizations after there is something to show. Very soon after having the first prototypes working, the scope was changed to make it available for Barracuda release.

## Design goals

* Infrastructure agnostic (works on any infra; on-prem, public cloud, private cloud, hybrid, baremetal)
* No mandatory dependencies to other tools or prerequisites to install anything beforehand to cluster machines
* Ultra Fast (as fast as possible given the stack we are dealing with)
* Multi cluster management (see all clusters, their health, running versions)
* Built-in telemetry (Installation & Upgrade related; errors too)
* Will provide meaningful output for diagnostics (e.g. error.log / install.log)
* Support Ubuntu 18.04 & RHEL 7/8 & CentOS 7/8 (+easily add new host OS support)
* Support for infrastructure management via Terraform
* Bonus: Possibility to use as an API/library for 3rd party tools & services & products such as MCM

We will draw inspiration from existing tooling such as [Docker Cluster](https://github.com/Mirantis/cluster) and [testkit](https://github.com/Mirantis/testkit)

## Built-In Telemetry & Improved Insights

1. When tool is used, it'll send data of every action performed with relevant payload. We don't collect any sensitive data or info about the workloads running in clusters itself. That's the job for the UCP built-in telemetry (which may be disabled too). The telemetry coming out from this tool will augment the telemetry data coming from UCP (and DTR and Engine in context of UCP).
2. Tool will require registration that we can hopefully use for sales & marketing purposes. We can see which users are actually actively interacting with our product (evaluation). In addition, it'll provide product management (or some other function) to get in touch with users to learn more about their needs. In my previous work this was super important. We contacted basically all people evaluation our product on a personal level. As outcome, we got very valuable feedback that we could apply into our products. I hope we can do this with Mirantis too; at least on some capacity.
3. We will start seeing funnel of users from their first interaction to our product --> successfully creating a cluster (for evaluation) --> subscribing to a license. We want to pay very close attention to this % since we want most people succeed and have positive experience with our product.

We try to find answer to questions like:

* How many people are interested in our product (did download the tool) but fail to create a working cluster? What are the common reasons for failures? Can we enhance our product or docs to improve the conversion rate?
* How many people are successful with our product? Did they get it up and running at first try or did they go through some hoops? How many failed attempts before working install? How long it will take from zero to hero experience?
* How people deal with updated version of our product? Do they try to upgrade their clusters right away or is there some significant delay? Is there anything in our product to improve frequency people update?
* What are the usage patterns; how many clusters people create? How often there is a need for new clusters?

The implementation will be made using Segment + Snowflake + Looker (similar to most of our other products). Detailed telemetry events & payload (TBD).

## Test Plan

* **PR Tests** - There are CI tests that are run on Jenkins for each pull request. Today, the tests will cover basic testing such as linter, unit tests and elementary smoke tests (simple integration tests). The coverage for unit tests is still rather limited due to massive pressure related to initial release schedule + stage of the product. **The Plan:** More unit tests will be made once the product is getting more mature and new features are added.
* **Integration Tests** - At this stage only elementary smoke tests are included to test the product on various host OS environments such as CentOS7/8 and Ubuntu 18.04.  **The Plan:** Add more smoke tests to cover more host OS options, and add k8s/conformance + k8s/sig-windows suites part of the smoke tests. In the future, automate tests utilizing built-in terraform integration on AWS/Azure/GCP/OpenStack/VMWare.
* **Manual Tests** - We hope QA team would run some of their existing test plans manually on clusters created with `launchpad`.

## Release Process

Releases are made from git tags by CICD system. The release builds must be triggered manually. The release process is the following:

1. Create new or update the existing release branch
2. Create new tag, for example `TAG=0.11.3 git tag $TAG && git push origin $TAG`
3. Go to [Jenkins](https://ci.docker.com/teams-orchestration/job/mcc/job/mcc/view/tags/) and select `Build now` from the dropdown menu of the corresponding tag to trigger the release build.
4. After the release build is ready, go to [Launchpad releases](https://github.com/Mirantis/launchpad/releases) in GitHub. Edit the draft release, write the changelog in the description field and publish the release.

## Comparison to Existing Tools

Overall, none of the pre-existing tools provide the convenience and flexibility we'd like to have for a tool that is designed for self service with an amazing new user experience.

* **Convenience** means user experience; minimizing prerequisites, must read docs, domain specific knowledge and actual steps required for performing the install (or upgrade).
* **Flexibility** means this tool should work in any environment, system or infrastructure the user might have already available for the purpose (=evaluation, running in prod or other). Note: with _any environment_ we should mean literally any. AWS, Azure, GCP, VMware, RH satellite managed systems, OpenStack, Mirantis private cloud where machines are ordered via IT tickets, any other private datacenter or whatever public cloud provider, rack of machines in your own home closet... etc.
