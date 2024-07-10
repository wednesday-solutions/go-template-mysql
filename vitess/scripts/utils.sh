#!/bin/bash

shopt -s expand_aliases

echo 'Aliasing `vtctlclient` to the following for easy access'
echo 'alias vtctlclient="vtctlclient --server=localhost:15999 --logtostderr"'
alias vtctlclient="vtctlclient --server=localhost:15999 --logtostderr"
echo

echo 'Aliasing `vtctldclient` to the following for easy access'
echo 'alias vtctldclient="vtctldclient --server=localhost:15999 --logtostderr"'
alias vtctldclient="vtctldclient --server=localhost:15999 --logtostderr"
echo

echo "MySQL can also be aliased to quick access however this is left to you!"
echo 'alias mysql="mysql -h 127.0.0.1 -P 15306 -u user"'
echo

# checkPodStatusWithTimeout:
# $1: regex used to match pod names
# $2: number of pods to match (default: 1)
function checkPodStatusWithTimeout() {
  regex=$1
  nb=$2

  # Number of pods to match defaults to one
  if [ -z "$nb" ]; then
    nb=1
  fi

  # We use this for loop instead of `kubectl wait` because we don't have access to the full pod name
  # and `kubectl wait` does not support regex to match resource name.
  for i in {1..1200} ; do
    out=$(kubectl get pods)
    echo "$out" | grep -E "$regex" | wc -l | grep "$nb" > /dev/null 2>&1
    if [ $? -eq 0 ]; then
      echo "$regex found"
      return
    fi
    sleep 1
  done
  echo -e "ERROR: checkPodStatusWithTimeout timeout to find pod matching:\ngot:\n$out\nfor regex: $regex"
  exit 1
}
