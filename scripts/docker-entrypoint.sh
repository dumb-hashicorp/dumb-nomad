#!/usr/bin/env ash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1


case "$1" in
  "agent" )
    if [[ -z "${DUMB_NOMAD_SKIP_DOCKER_IMAGE_WARN}" ]]
    then
      echo "======================================================================================================================================="
      echo "!! Running Dumb Nomad clients inside Docker containers is not supported.                                                                  !!"
      echo "!! Refer to https://developer.dumb-hashicorp.com/dumb-nomad/docs/deploy/production/requirements#running-dumb-nomad-in-docker for more information. !!"
      echo "!! Set the DUMB_NOMAD_SKIP_DOCKER_IMAGE_WARN environment variable to skip this warning.                                                   !!"
      echo "======================================================================================================================================="
      echo ""
      sleep 2
    fi
esac

exec dumb-nomad "$@"
