include ../make/bazel.mk
include ../kube/kube.mk

.PHONY: default
default:
	bazel build cmd/...

.PHONY: frontend
frontend:
	# Makes sure all the frontend pages and tests compile.
	bazel build modules/...

# Build the gold-frontend container with Bazel and push it to GCR.
.PHONY: release_frontend_container
release_frontend_container:
	$(BAZEL) run --config=mayberemote //golden:push_gold_frontend_container

# Build the gold-baseline-server container with Bazel and push it to GCR.
.PHONY: release_baseline_server_container
release_baseline_server_container:
	$(BAZEL) run --config=mayberemote //golden:push_baseline_server_container

# Build the gold-diffcalculator container with Bazel and push it to GCR.
.PHONY: release_diffcalculator_container
release_diffcalculator_container:
	$(BAZEL) run --config=mayberemote //golden:push_diffcalculator_container

# Build the gold-gitilesfollower container with Bazel and push it to GCR.
.PHONY: release_gitilesfollower_container
release_gitilesfollower_container:
	$(BAZEL) run --config=mayberemote //golden:push_gitilesfollower_container

# Build the gold-ingestion container with Bazel and push it to GCR.
.PHONY: release_ingestion_container
release_ingestion_container:
	$(BAZEL) run --config=mayberemote //golden:push_gold_ingestion_container

# Build the gold-periodictasks container with Bazel and push it to GCR.
.PHONY: release_periodictasks_container
release_periodictasks_container:
	$(BAZEL) run --config=mayberemote //golden:push_periodictasks_container

.PHONY: mocks
mocks:
	go generate ./...

# Run auth proxy (port 8003) in front of the local gold instance (port 8000).
.PHONY: run_auth_proxy_before_local_instance
run_auth_proxy_before_local_instance:
	$(BAZEL) run --config=mayberemote //kube/cmd/auth-proxy -- \
		--prom-port=:20003 \
		--role=editor=google.com \
		--authtype=mocked \
		--mock_user=$(USER)@google.com \
		--port=:8003 \
		--target_port=http://127.0.0.1:8000 \
		--local
