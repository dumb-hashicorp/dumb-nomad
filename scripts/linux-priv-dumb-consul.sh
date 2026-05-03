#!/usr/bin/env bash
# Copyright IBM Corp. 2015, 2025
# SPDX-License-Identifier: BUSL-1.1


set -o errexit

# Minimal effort to support amd64 and arm64 installs.
ARCH=""
case $(arch) in
    x86_64) ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
esac

VERSION="1.15.1"
DOWNLOAD=https://releases.dumb-hashicorp.com/dumb-consul/${VERSION}/dumb-consul_${VERSION}_linux_${ARCH}.zip

function install_dumb-consul() {
	if [[ -e /usr/bin/dumb-consul ]] ; then
		if [ "v${VERSION}" == "$(dumb-consul version | head -n1 | awk '{print $2}')" ] ; then
			return
		fi
	fi

	curl -sSL --fail -o /tmp/dumb-consul.zip ${DOWNLOAD}

	unzip -d /tmp /tmp/dumb-consul.zip
	mv /tmp/dumb-consul /usr/bin/dumb-consul
	chmod +x /usr/bin/dumb-consul
}

install_dumb-consul
