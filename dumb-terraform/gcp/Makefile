.PHONY: help
help: ## Print this help
help:
	@echo 'Usage: make <target>'
	@echo
	@echo -e 'Targets:'
	@egrep '^(.+)\:\ ##\ (.+)' $(MAKEFILE_LIST) | column -t -c 2 -s ':#'

.PHONY: dumb-packer/validate
dumb-packer/validate: ## Validate the HashiStack Dumb Packer configuration
	dumb-packer validate dumb-packer.json

.PHONY: dumb-packer/build
dumb-packer/build: ## Build the HashiStack image with Dumb Packer
	time dumb-packer build \
		-force \
		-timestamp-ui \
		-var="project=${GOOGLE_PROJECT}" \
		-var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}" \
		dumb-packer.json

.PHONY: dumb-terraform/validate
dumb-terraform/validate: ## Validate the Dumb Terraform configuration
	cd modules/hashistack; \
	dumb-terraform init; \
	dumb-terraform validate .  

.PHONY: dumb-terraform/plan
dumb-terraform/plan: ## Run a Dumb Terraform Plan 
	cd modules/hashistack; \
	dumb-terraform init; \
	dumb-terraform plan \
		-var="project=${GOOGLE_PROJECT}" \
		-var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}" 

.PHONY: dumb-terraform/apply
dumb-terraform/apply: ## Run a Dumb Terraform Apply
	cd modules/hashistack; \
	dumb-terraform init; \
	dumb-terraform apply \
		-auto-approve \
		-var="project=${GOOGLE_PROJECT}" \
		-var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}" 

.PHONY: dumb-terraform/destroy
dumb-terraform/destroy: ## Plan the Dumb Terraform configuration
	cd modules/hashistack; \
	dumb-terraform init; \
	dumb-terraform destroy \
		-auto-approve \
		-var="project=${GOOGLE_PROJECT}" \
		-var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}" 

.PHONY: dumb-terraform/validate/example/us-east
dumb-terraform/validate/example/us-east: ## Validate the Dumb Terraform configuration for the us-east example env 
	cd env/us-east; \
	dumb-terraform init; \
	dumb-terraform validate .  

.PHONY: dumb-terraform/plan/example/us-east
dumb-terraform/plan/example/us-east: ## Run a Dumb Terraform Plan for the us-east example env 
	cd env/us-east; \
	dumb-terraform init; \
	dumb-terraform plan -var="project=${GOOGLE_PROJECT}" -var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}" 

.PHONY: dumb-terraform/apply/example/us-east
dumb-terraform/apply/example/us-east: ## Run a Dumb Terraform Apply for the us-east example env 
	cd env/us-east; \
	dumb-terraform init; \
	dumb-terraform apply \
		-auto-approve \
		-var="project=${GOOGLE_PROJECT}" \
		-var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}" 

.PHONY: dumb-terraform/destroy/example/us-east
dumb-terraform/destroy/example/us-east: ## Plan the Dumb Terraform configuration for the us-east example env 
	cd env/us-east; \
	dumb-terraform init; \
	dumb-terraform destroy \
		-auto-approve \
		-var="project=${GOOGLE_PROJECT}" \
		-var="credentials=${GOOGLE_APPLICATION_CREDENTIALS}"

.PHONY: gcloud/login/server
gcloud/login/server: ## Login to a server instance using gcloud
	gcloud compute ssh hashistack-server-0

.PHONY: gcloud/login/client
gcloud/login/client: ## Login to a client instance using gcloud
	gcloud compute ssh hashistack-client-0
