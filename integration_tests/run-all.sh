#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

DETAILED_OUTPUT=true

LOCALNET_DIR=".localnet"

LOG_PREFIX="[run all] "
SED_EXT=
if [ "$(uname)" == "Darwin" ]; then
    # Mac OS X sed needs the file extension when -i flag is used. Keeping it empty as we don't need backupfile
    SED_EXT="''"
fi

if ${DETAILED_OUTPUT}; then
  DETAILED_OUTPUT_TARGET=/dev/stdout
else
  DETAILED_OUTPUT_TARGET=/dev/null
fi

cleanup() {
    make localnet_clean
}
trap cleanup EXIT

source integration_tests/cli/common.sh

log() {
  echo "${LOG_PREFIX}$1"
}

  # patch configs properly by having all values >= 1 sec, otherwise headers may start having time from the future and light client verification will fail
  # if we patch config to have new blocks created in less than 1 sec, the min time in a time header is still 1 sec.
  # So, new blocks started to be from the future.
patch_consensus_config() {
  local NODE_CONFIGS="$(find "$LOCALNET_DIR" -type f -name "config.toml" -wholename "*node*")"

  for NODE_CONFIG in ${NODE_CONFIGS}; do
    sed -i $SED_EXT 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "${NODE_CONFIG}"
    #sed -i $SED_EXT 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' "${NODE_CONFIG}"
    #sed -i $SED_EXT 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' "${NODE_CONFIG}"
    sed -i $SED_EXT 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "${NODE_CONFIG}"
  done
}

init_pool() {
  local _patch_config="${1:-yes}";
  log "Setting up pool"

  log "-> Generating network configuration" >${DETAILED_OUTPUT_TARGET}
  make localnet_init &>${DETAILED_OUTPUT_TARGET}

  if [ "$_patch_config" = "yes" ];
  then
    patch_consensus_config
  fi;

  log "-> Running pool" >${DETAILED_OUTPUT_TARGET}
  make localnet_start &>${DETAILED_OUTPUT_TARGET}

  log "-> Waiting for the second block (needed to request proofs)" >${DETAILED_OUTPUT_TARGET}
  wait_for_height 2 20
}

cleanup_pool() {
  log "Cleaning up pool"
  log "-> Stopping pool & Removing configurations" >${DETAILED_OUTPUT_TARGET}
  make localnet_clean
}

stop_rest_server() {
  log "Stopping cli in rest-server mode"
  killall dcld
}

# Global init
set -euo pipefail

log "Verifying the environment"
check_env

log "Compiling local binaries"
make install &>${DETAILED_OUTPUT_TARGET}

log "Building docker image"
make image &>${DETAILED_OUTPUT_TARGET}

cleanup_pool

# Cli shell tests
CLI_SHELL_TESTS=$(find integration_tests/cli -type f -name '*.sh' -not -name "common.sh")

for CLI_SHELL_TEST in ${CLI_SHELL_TESTS}; do
  init_pool

   log "*****************************************************************************************"
   log "Running $CLI_SHELL_TEST"
   log "*****************************************************************************************"

   if bash "$CLI_SHELL_TEST" &>${DETAILED_OUTPUT_TARGET}; then
     log "$CLI_SHELL_TEST finished successfully"
   else
     log "$CLI_SHELL_TEST failed"
     exit 1
   fi

   cleanup_pool
done

# Light Client Proxy Cli shell tests
CLI_SHELL_TESTS=$(find integration_tests/light_client_proxy -type f -name '*.sh' -not -name "common.sh")

for CLI_SHELL_TEST in ${CLI_SHELL_TESTS}; do
  init_pool 

  log "*****************************************************************************************"
  log "Running $CLI_SHELL_TEST"
  log "*****************************************************************************************"
  
  if bash "$CLI_SHELL_TEST" &>${DETAILED_OUTPUT_TARGET}; then
    log "$CLI_SHELL_TEST finished successfully"
  else
    log "$CLI_SHELL_TEST failed"
    exit 1
  fi

  cleanup_pool
done

# Go rest tests
GO_REST_TESTS="$(find integration_tests/grpc_rest -type f -name '*_test.go')"

for GO_REST_TEST in ${GO_REST_TESTS}; do
  init_pool

  log "*****************************************************************************************"
  log "Running $GO_REST_TEST"
  log "*****************************************************************************************"

  # TODO issue 99: improve, that await helps with the cases of not ready connections to Cosmos endpoints
  sleep 5

  dcld config keyring-backend test
  if go test "$GO_REST_TEST" &>${DETAILED_OUTPUT_TARGET}; then
    log "$GO_REST_TEST finished successfully"
  else
    log "$GO_REST_TEST failed"
    exit 1
  fi

  cleanup_pool
done