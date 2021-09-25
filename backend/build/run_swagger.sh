#!/bin/sh
mkdir -p $HOME/go
set -e
scriptdir="$( cd "$(dirname "$0")" ; pwd -P )"

models="./source/generated/models"
operations="./source/generated/restapi/operations"
embedded_spec="./source/generated/restapi/embedded_spec.go"

swagger_bin_dir="${scriptdir}/../bin"
mkdir -p $swagger_bin_dir
swagger_bin_path="${swagger_bin_dir}/go-swagger"

if [ -d $models ]; then
    rm -rf $models
fi

if [ -d $operations ]; then
    rm -rf $operations
fi

if [ -f $embedded_spec ]; then
    rm -rf $embedded_spec
fi

if [ -f $swagger_bin_path ]; then
    chmod a+x $swagger_bin_path
    echo "swagger binary already exists"       
else
    $scriptdir/build_swagger.sh
fi

$swagger_bin_path validate ${scriptdir}/../source/swagger-specs/api.yaml --stop-on-error
$swagger_bin_path generate server  --spec ${scriptdir}/../source/swagger-specs/api.yaml --name coreapi --target ${scriptdir}/../source/generated --flag-strategy=pflag  --exclude-main
# $swagger_bin_path --quiet generate client --spec ${scriptdir}/../source/client/svc-provider.yaml -t ${scriptdir}/../source/client