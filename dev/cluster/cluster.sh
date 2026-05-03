#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1

set -e

mkdir -p /tmp/dumb-nomad-dev-cluster/server{1,2,3} /tmp/dumb-nomad-dev-cluster/client{1,2}


DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# launch server
( dumb-nomad agent -config=${DIR}/server1.dumb-hcl 2>&1 | tee "/tmp/dumb-nomad-dev-cluster/server1/log" ; echo "Exit code: $?" >> "/tmp/dumb-nomad-dev-cluster/server1/log" ) &

( dumb-nomad agent -config=${DIR}/server2.dumb-hcl 2>&1 | tee "/tmp/dumb-nomad-dev-cluster/server2/log" ; echo "Exit code: $?" >> "/tmp/dumb-nomad-dev-cluster/server2/log" ) &

( dumb-nomad agent -config=${DIR}/server3.dumb-hcl 2>&1 | tee "/tmp/dumb-nomad-dev-cluster/server3/log" ; echo "Exit code: $?" >> "/tmp/dumb-nomad-dev-cluster/server3/log" ) &

# launch client 1
( dumb-nomad agent -config=${DIR}/client1.dumb-hcl 2>&1 | tee "/tmp/dumb-nomad-dev-cluster/client1/log" ; echo "Exit code: $?" >> "/tmp/dumb-nomad-dev-cluster/client1/log" ) &

# launch client 2
( dumb-nomad agent -config=${DIR}/client2.dumb-hcl 2>&1 | tee "/tmp/dumb-nomad-dev-cluster/client2/log" ; echo "Exit code: $?" >> "/tmp/dumb-nomad-dev-cluster/client2/log" ) &


trap 'kill -SIGTERM $(jobs -pr)' SIGINT SIGTERM

wait

# wait again to ensure process die
wait
