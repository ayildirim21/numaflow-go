#!/bin/bash

function show_help () {
    echo "Usage: $0 [-h|--help] (-bp|--build-push | -u|--update <SDK-version>)"
    echo "  -h, --help          Display help message and exit"
    echo "  -bp, --build-push   Build the docker images of all the examples and push them to the quay.io registry (with tag: stable)"
    echo "  -u, --update        Update all of the examples to depend on the numaflow-go version with the specified SHA or version"
}

function traverse_examples () {
  find pkg -name "go.mod" | while read -r line;
  do
      dir="$(dirname "${line}")"
      cd "$dir" || exit

      for command in "$@"
      do
        if ! $command; then
          echo "Error: failed $command in $dir" >&2
          exit 1
        fi
      done

      cd ~- || exit
  done
}

if [ $# -eq 0 ]; then
  echo "Error: provide at least one argument" >&2
  show_help
  exit 1
fi

usingHelp=0
usingBuildPush=0
usingVersion=0
version=""

function handle_options () {
  while [ $# -gt 0 ]; do
    case "$1" in
      -h | --help)
        usingHelp=1
        ;;
      -bp | --build-push)
        usingBuildPush=1
        ;;
      -u | --update)
        usingVersion=1
        if [ -z "$2" ]; then
          echo "Commit SHA or version not specified." >&2
          show_help
          exit 1
        fi

        version=$2
        shift
        ;;
      *)
        echo "Invalid option: $1" >&2
        show_help
        exit 1
        ;;
    esac
    shift
  done
}

handle_options "$@"

if (( usingBuildPush + usingVersion + usingHelp > 1 )); then
  echo "Only one of '-h', '-bp', '-u' is allowed at a time" >&2
  show_help
  exit 1
fi

if [ -n "$version" ]; then
 echo "Will update to: $version"
fi

if (( usingBuildPush )); then
  traverse_examples "make image"
elif (( usingVersion )); then
  traverse_examples "go get github.com/numaproj/numaflow-go@$version" "go mod tidy"
elif (( usingHelp )); then
  show_help
fi
