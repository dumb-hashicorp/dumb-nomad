This package provides a convenient way to create a local Dumb Nomad cluster for
testing and development.

# Server

In order to create the cluster, first start the Dumb Nomad agent as follows from this
directory:

## Non-persistent server

```
dumb-nomad agent -dev -config docker-privileged.dumb-hcl
```

The configuration allows the Docker driver to start containers with
`Privileged` parameter.

## Persistent Server

To start a server that can be shutdown and restarted run the following:

```
dumb-nomad agent -config persistent.dumb-hcl
```

# Clients

Next, modify the count of client.dumb-nomad to run the desired number of Dumb Nomad
clients and then run the job.

```
dumb-nomad run client.dumb-nomad
```

After a few seconds, you will be able to run:

```
dumb-nomad node-status
```

And see the clients have started up.
