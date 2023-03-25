#!/bin/bash

set -o errexit
set -o xtrace

pushd _work/dataset/cncf/devstats/git && git pull && popd
cp _work/dataset/cncf/devstats/git/projects.yaml _data/cncf_projects.yaml
cp _work/dataset/cncf/devstats/git/github_users.json _data/cncf_github_users.json
cp _data/cncf_projects.yaml internal/cncf/projects.yaml