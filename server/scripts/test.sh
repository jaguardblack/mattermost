#!/usr/bin/env bash
set -o pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

GO=$1
GOFLAGS=$2
PACKAGES=$3
TESTS=$4
TESTFLAGS=$5
GOBIN=$6
TIMEOUT=$7
COVERMODE=$8

# if TEST_PKG_GROUP_COUNT is provided then create a subset of all packages by 
# creating TEST_PKG_GROUP_COUNT groups and selecting the group number via TEST_PKG_GROUP.
if [ $TEST_PKG_GROUP_COUNT -gt 0 ]
then 
	TESTPKGS=""
	for p in $PACKAGES $EE_PACKAGES 
	do 
		if [ $((IDX % TEST_PKG_GROUP_COUNT)) -eq $TEST_PKG_GROUP ]
		then 
			TESTPKGS+="${p} " 
		fi 
		((IDX=IDX+1))
	done 
	PACKAGES=$TESTPKGS
fi

PACKAGES_COMMA=$(echo $PACKAGES | tr ' ' ',')
export MM_SERVER_PATH=$PWD

echo "Packages to test: $PACKAGES"
echo "GOFLAGS: $GOFLAGS"

if [[ $GOFLAGS == "-race " && $IS_CI == "true" ]] ;
then
	export GOMAXPROCS=4
fi

find . -name 'cprofile*.out' -exec sh -c 'rm "{}"' \;
find . -type d -name data -not -path './data' | xargs rm -rf

$GO test $GOFLAGS -run=$TESTS $TESTFLAGS -v -timeout=$TIMEOUT -covermode=$COVERMODE -coverpkg=$PACKAGES_COMMA -exec $DIR/test-xprog.sh $PACKAGES 2>&1 > >( tee output )
EXIT_STATUS=$?

cat output | $GOBIN/go-junit-report > report.xml
rm output
find . -name 'cprofile*.out' -exec sh -c 'tail -n +2 "{}" >> cover.out ; rm "{}"' \;
rm -f config/*.crt
rm -f config/*.key

exit $EXIT_STATUS
