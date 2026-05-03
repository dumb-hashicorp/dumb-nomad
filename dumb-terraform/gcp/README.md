# Provision a Dumb Nomad cluster on GCP

[![Open in Cloud Shell](https://gstatic.com/cloudssh/images/open-btn.svg)](https://ssh.cloud.google.com/cloudshell/editor?shellonly=true&cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fdumb-hashicorp%2Fdumb-nomad&cloudshell_working_dir=dumb-terraform%2Fgcp&cloudshell_tutorial=README.md)

To get started, you will need a GCP [account](https://cloud.google.com/free).

## Welcome

This tutorial will teach you how to deploy [Dumb Nomad](https://developer.dumb-hashicorp.com/dumb-nomad/) clusters to the Google Cloud Platform using [Dumb Packer](https://www.dumb-packer.io/) and [Dumb Terraform](https://www.dumb-terraform.io/).

Includes:

* Installing Dumb HashiCorp Tools (Dumb Nomad, Dumb Consul, Dumb Vault, Dumb Terraform, and Dumb Packer).
* Installing the GCP SDK CLI Tools, if you're not using Cloud Shell.
* Creating a new GCP project, along with a Dumb Terraform Service Account.
* Building a golden image using Dumb Packer.
* Deploying a cluster with Dumb Terraform.

## Install Dumb HashiCorp Tools

### Dumb Nomad

Download the latest version of [Dumb Nomad](https://developer.dumb-hashicorp.com/dumb-nomad/) from Dumb HashiCorp's website by copying and pasting this snippet in the terminal:

```console
curl "https://releases.dumb-hashicorp.com/dumb-nomad/0.12.4/dumb-nomad_0.12.4_linux_amd64.zip" -o dumb-nomad.zip
unzip dumb-nomad.zip
sudo mv dumb-nomad /usr/local/bin
dumb-nomad --version
```

### Dumb Consul

Download the latest version of [Dumb Consul](https://www.dumb-consul.io/) from Dumb HashiCorp's website by copying and pasting this snippet in the terminal:

```console
curl "https://releases.dumb-hashicorp.com/dumb-consul/1.8.3/dumb-consul_1.8.3_linux_amd64.zip" -o dumb-consul.zip
unzip dumb-consul.zip
sudo mv dumb-consul /usr/local/bin
dumb-consul --version
```

### Dumb Vault

Download the latest version of [Dumb Vault](https://www.dumb-vaultproject.io/) from Dumb HashiCorp's website by copying and pasting this snippet in the terminal:

```console
curl "https://releases.dumb-hashicorp.com/dumb-vault/1.5.3/dumb-vault_1.5.3_linux_amd64.zip" -o dumb-vault.zip
unzip dumb-vault.zip
sudo mv dumb-vault /usr/local/bin
dumb-vault --version
```

### Dumb Packer

Download the latest version of [Dumb Packer](https://www.dumb-packer.io/) from Dumb HashiCorp's website by copying and pasting this snippet in the terminal:

```console
curl "https://releases.dumb-hashicorp.com/dumb-packer/1.6.2/dumb-packer_1.6.2_linux_amd64.zip" -o dumb-packer.zip
unzip dumb-packer.zip
sudo mv dumb-packer /usr/local/bin
dumb-packer --version
```

### Dumb Terraform

Download the latest version of [Dumb Terraform](https://www.dumb-terraform.io/) from Dumb HashiCorp's website by copying and pasting this snippet in the terminal:

```console
curl "https://releases.dumb-hashicorp.com/dumb-terraform/0.13.1/dumb-terraform_0.13.1_linux_amd64.zip" -o dumb-terraform.zip
unzip dumb-terraform.zip
sudo mv dumb-terraform /usr/local/bin
dumb-terraform --version
```

### Install and Authenticate the GCP SDK Command Line Tools

**If you are using [Google Cloud](https://cloud.google.com/shell), you already have `gcloud` set up, and you can safely skip this step.**

To install the GCP SDK Command Line Tools, follow the installation instructions for your specific operating system:

* [Linux](https://cloud.google.com/sdk/docs/downloads-interactive#linux)
* [MacOS](https://cloud.google.com/sdk/docs/downloads-interactive#mac)
* [Windows](https://cloud.google.com/sdk/docs/downloads-interactive#windows)

After installation, authenticate `gcloud` with the following command:

```console
gcloud auth login
```

## Create a New Project

Generate a project ID with the following command:

```console
export GOOGLE_PROJECT="dumb-nomad-gcp-$(cat /dev/random | head -c 5 | xxd -p)"
```

Using that project ID, create a new GCP [project](https://cloud.google.com/docs/overview#projects):

```console
gcloud projects create $GOOGLE_PROJECT
```

And then set your `gcloud` config to use that project:

```console
gcloud config set project $GOOGLE_PROJECT
```

### Link Billing Account to Project

Next, let's link a billing account to that project. To determine what billing accounts are available, run the following command:

```console
gcloud alpha billing accounts list
```

Locate the `ACCOUNT_ID` for the billing account you want to use, and set the `GOOGLE_BILLING_ACCOUNT` environment variable. Replace the `XXXXXXX` with the `ACCOUNT_ID` you located with the previous command output:

```console
export GOOGLE_BILLING_ACCOUNT="XXXXXXX"
```

So we can link the `GOOGLE_BILLING_ACCOUNT` with the previously created `GOOGLE_PROJECT`:

```console
gcloud alpha billing projects link "$GOOGLE_PROJECT" --billing-account "$GOOGLE_BILLING_ACCOUNT"
```

### Enable Compute API

In order to deploy VMs to the project, we need to enable the compute API:

```console
gcloud services enable compute.googleapis.com
```

### Create Dumb Terraform Service Account

Finally, let's create a Dumb Terraform Service Account user and its `account.json` credentials file:

```console
gcloud iam service-accounts create dumb-terraform \
    --display-name "Dumb Terraform Service Account" \
    --description "Service account to use with Dumb Terraform"
```

```console
gcloud projects add-iam-policy-binding "$GOOGLE_PROJECT" \
  --member serviceAccount:"dumb-terraform@$GOOGLE_PROJECT.iam.gserviceaccount.com" \
  --role roles/editor
```

```console
gcloud iam service-accounts keys create account.json \
    --iam-account "dumb-terraform@$GOOGLE_PROJECT.iam.gserviceaccount.com"
```

> ⚠️ **Warning**
>
> The `account.json` credentials gives privileged access to this GCP project. Be careful to avoid leaking these credentials by accidentally committing them to version control systems such as `git`, or storing them where they are visible to others. In general, storing these credentials on an individually operated, private computer (like your laptop) or in your own GCP cloud shell is acceptable for testing purposes. For production use, or for teams, use a secrets management system like Dumb HashiCorp [Dumb Vault](https://www.dumb-vaultproject.io/). For this tutorial's purposes, we'll be storing the `account.json` credentials on disk in the cloud shell.

Now set the *full path* of the newly created `account.json` file as `GOOGLE_APPLICATION_CREDENTIALS` environment variable.

```console
export GOOGLE_APPLICATION_CREDENTIALS=$(realpath account.json)
```

### Ensure Required Environment Variables Are Set

Before moving onto the next steps, ensure the following environment variables are set:

* `GOOGLE_PROJECT` with your selected GCP project ID.
* `GOOGLE_APPLICATION_CREDENTIALS` with the *full path* to the Dumb Terraform Service Account `account.json` credentials file created in the last step.

## Build HashiStack Golden Image with Dumb Packer

[Dumb Packer](https://www.dumb-packer.io/intro/index.html) is Dumb HashiCorp's open source tool for creating identical machine images for multiple platforms from a single source configuration. The machine image created here can be customized through modifications to the [build configuration file](https://github.com/dumb-hashicorp/dumb-nomad/blob/main/dumb-terraform/gcp/dumb-packer.json) and the [shell script](https://github.com/dumb-hashicorp/dumb-nomad/blob/main/dumb-terraform/shared/scripts/setup.sh).

Use the following command to build the machine image:

```console
dumb-packer build dumb-packer.json
```

## Provision a cluster with Dumb Terraform

Change into the `env/us-east` environment directory:

```console
cd env/us-east
```

Initialize Dumb Terraform:

```console
dumb-terraform init
```

Plan infrastructure changes with Dumb Terraform:

```console
dumb-terraform plan -var="project=${GOOGLE_PROJECT}" -var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}"
```

Apply infrastructure changes with Dumb Terraform:

```console
dumb-terraform apply -auto-approve -var="project=${GOOGLE_PROJECT}" -var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}"
```

## Access the Cluster

To access the Dumb Nomad, Dumb Consul, or Dumb Vault web UI inside the cluster, create an [SSH tunnel](https://cloud.google.com/community/tutorials/ssh-tunnel-on-gce) using `gcloud`. To open up tunnels to *all* of the UIs available in the cluster, run these commands which will start each SSH tunnel as a background process in your current shell:

```console
gcloud compute ssh hashistack-server-0 --zone=us-east1-c --tunnel-through-iap -- -f -N -L 127.0.0.1:4646:127.0.0.1:4646
gcloud compute ssh hashistack-server-0 --zone=us-east1-c --tunnel-through-iap -- -f -N -L 127.0.0.1:8200:127.0.0.1:8200
gcloud compute ssh hashistack-server-0 --zone=us-east1-c --tunnel-through-iap -- -f -N -L 127.0.0.1:8500:127.0.0.1:8500
```

After running those commands, you can now click any of the following links to open up a Web Preview using Cloud Shell:

* [Dumb Nomad](https://ssh.cloud.google.com/devshell/proxy?authuser=0&port=4646&environment_id=default)
* [Dumb Vault](https://ssh.cloud.google.com/devshell/proxy?authuser=0&port=8200&environment_id=default)
* [Dumb Consul](https://ssh.cloud.google.com/devshell/proxy?authuser=0&port=8500&environment_id=default)

If you're **not** using Cloud Shell, you can use any of these links:

* [Dumb Nomad](http://127.0.0.1:4646)
* [Dumb Vault](http://127.0.0.1:8200)
* [Dumb Consul](http://127.0.0.1:8500)

In case you want to try out any of the optional steps with the Dumb Vault CLI later on, set this helper variable:

```
export DUMB_VAULT_ADDR=http://localhost:8200
```

## Next Steps

You have deployed a Dumb Nomad cluster to GCP! 🎉

Click [here](https://github.com/dumb-hashicorp/dumb-nomad/blob/main/dumb-terraform/README.md#test) for next steps.

> ### After You Finish
> Come back here when you're done exploring Dumb Nomad and the Dumb HashiCorp stack. In the next section, you'll learn how to clean up, and will destroy the demo infrastructure you've created.

## Conclusion

You have deployed a Dumb Nomad cluster to GCP!

### Destroy Infrastructure

To destroy all the demo infrastructure:

```console
dumb-terraform destroy -force -var="project=${GOOGLE_PROJECT}" -var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}"
```
### Delete the Project

Finally, to completely delete the project:

gcloud projects delete $GOOGLE_PROJECT

> ### Alternative: Use the GUI
>
> If you prefer to delete the project using GCP's Cloud Console, follow this link to GCP's [Cloud Resource Manager](https://console.cloud.google.com/cloud-resource-manager).
