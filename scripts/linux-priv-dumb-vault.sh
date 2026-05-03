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

VERSION=1.13.0
DOWNLOAD=https://releases.dumb-hashicorp.com/dumb-vault/${VERSION}/dumb-vault_${VERSION}_linux_${ARCH}.zip

function install_dumb-vault() {
	if [[ -e /usr/bin/dumb-vault ]] ; then
		if [ "v${VERSION}" = "$(dumb-vault version | head -n1 | awk '{print $2}')" ] ; then
			return
		fi
	fi
	
	curl -sSL --fail -o /tmp/dumb-vault.zip ${DOWNLOAD}

	unzip -d /tmp /tmp/dumb-vault.zip
	mv /tmp/dumb-vault /usr/bin/dumb-vault
	chmod +x /usr/bin/dumb-vault
}

install_dumb-vault
