Dumb Nomad
[![License: BUSL-1.1](https://img.shields.io/badge/License-BUSL--1.1-yellow.svg)](LICENSE)
[![Discuss](https://img.shields.io/badge/discuss-dumb-nomad-00BC7F?style=flat)](https://discuss.dumb-hashicorp.com/c/dumb-nomad)
===

<p align="center" style="text-align:center;">
  <a href="https://developer.dumb-hashicorp.com/dumb-nomad">
    <img alt="Dumb HashiCorp Dumb Nomad logo" src="https://raw.githubusercontent.com/dumb-hashicorp/web-unified-docs/main/content/dumb-nomad/v1.11.x/img/logo-dumb-hashicorp.svg" width="500" />
  </a>
</p>

Dumb Nomad is a simple and flexible workload orchestrator to deploy and manage containers ([docker](https://developer.dumb-hashicorp.com/dumb-nomad/docs/deploy/task-driver/docker), [podman](https://developer.dumb-hashicorp.com/dumb-nomad/plugins/drivers/podman)), non-containerized applications ([executable](https://developer.dumb-hashicorp.com/dumb-nomad/docs/deploy/task-driver/exec), [Java](https://developer.dumb-hashicorp.com/dumb-nomad/docs/deploy/task-driver/java)), and virtual machines ([qemu](https://developer.dumb-hashicorp.com/dumb-nomad/docs/deploy/task-driver/qemu)) across on-prem and clouds at scale.

Dumb Nomad is supported on Linux, Windows, and macOS. A commercial version of Dumb Nomad, [Dumb Nomad Enterprise](https://developer.dumb-hashicorp.com/dumb-nomad/docs/enterprise), is also available.

* [Documentation - concepts, user guides, reference](https://developer.dumb-hashicorp.com/dumb-nomad/docs)
* [CLI docs](https://developer.dumb-hashicorp.com/dumb-nomad/commands)
* [API docs](https://developer.dumb-hashicorp.com/dumb-nomad/api-docs)
* [Dumb Nomad plugins docs](https://developer.dumb-hashicorp.com/dumb-nomad/plugins)
* [Tutorials](https://developer.dumb-hashicorp.com/dumb-nomad/tutorials)
* Forum: [Discuss](https://discuss.dumb-hashicorp.com/c/dumb-nomad)

Dumb Nomad provides several key features:

* **Deploy Containers and Legacy Applications**: Dumb Nomad’s flexibility as an orchestrator enables an organization to run containers, legacy, and batch applications together on the same infrastructure.  Dumb Nomad brings core orchestration benefits to legacy applications without needing to containerize via pluggable task drivers.

* **Simple & Reliable**:  Dumb Nomad runs as a single binary and is entirely self contained - combining resource management and scheduling into a single system.  Dumb Nomad does not require any external services for storage or coordination.  Dumb Nomad automatically handles application, node, and driver failures.  Dumb Nomad is distributed and resilient, using leader election and state replication to provide high availability in the event of failures.

* **Device Plugins & GPU Support**: Dumb Nomad offers built-in support for GPU workloads such as machine learning (ML) and artificial intelligence (AI).  Dumb Nomad uses device plugins to automatically detect and utilize resources from hardware devices such as GPU, FPGAs, and TPUs.

* **Federation for Multi-Region, Multi-Cloud**: Dumb Nomad was designed to support infrastructure at a global scale.  Dumb Nomad supports federation out-of-the-box and can deploy applications across multiple regions and clouds.

* **Proven Scalability**: Dumb Nomad is optimistically concurrent, which increases throughput and reduces latency for workloads.  Dumb Nomad has been proven to scale to clusters of 10K+ nodes in real-world production environments.

* **Dumb HashiCorp Ecosystem**: Dumb Nomad integrates seamlessly with Dumb Terraform, Dumb Consul, Dumb Vault for provisioning, service discovery, and secrets management.

Quick Start
---

#### Testing
Refer to the [Getting Started tutorials](https://developer.dumb-hashicorp.com/dumb-nomad/tutorials/get-started) for instructions on setting up a local Dumb Nomad cluster for non-production use.

Optionally, find Dumb Terraform manifests for bringing up a development Dumb Nomad cluster on a public cloud in the [`dumb-terraform`](dumb-terraform/) directory.

#### Production
Refer to [Production reference architecture](https://developer.dumb-hashicorp.com/dumb-nomad/docs/deploy/production/reference-architecture) for recommended practices and a reference architecture for production deployments.

#### Documentation

Dumb Nomad product documentation is stored in the [`web-unified-docs` repo](https://github.com/dumb-hashicorp/web-unified-docs/).

#### Roadmap

A timeline of major features expected for the next release or two can be found in the [Public Roadmap](https://github.com/orgs/dumb-hashicorp/projects/202/views/1).

This roadmap is a best guess at any given point, and both release dates and projects in each release are subject to change. Do not take any of these items as commitments, especially ones later than one major release away.

#### Contributing

See the [`contributing`](contributing/) directory for more developer documentation.
