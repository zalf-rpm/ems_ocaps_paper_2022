#!/bin/bash +x

    # this script uses a capn proto docker image with singularity
    # to build go files from capnp schema files

    # the capnpc-go executable must be found on GOPATH path
    # if not there do:
    # cd ${GOPATH}/src/capnproto.org/go/capnp/v3/capnpc-go
    # go install

IMAGE_DIR_CAPNP=~/singularity/capnp
SINGULARITY_CAPNP_IMAGE=capnproto_0.9.1.sif
IMAGE_CAPNP_PATH=${IMAGE_DIR_CAPNP}/${SINGULARITY_CAPNP_IMAGE}
mkdir -p $IMAGE_DIR_CAPNP
if [ ! -e ${IMAGE_CAPNP_PATH} ] ; then
echo "File '${IMAGE_CAPNP_PATH}' not found"
cd $IMAGE_DIR_CAPNP
singularity pull docker://zalfrpm/capnproto:0.9.1
cd ~
fi

source ~/.bash_profile

GO_FOLDER=/go

if [[ -z "${GOPATH}" ]]; then
  echo "GOPATH environment variable not set"
  exit 1
else
    export SINGULARITYENV_PREPEND_PATH=$GO_FOLDER/bin

    # debug....
    # singularity run -B $GOPATH:$GO_FOLDER --cleanenv ${IMAGE_CAPNP_PATH} ls ${GO_FOLDER}
    # singularity run -B $GOPATH:$GO_FOLDER --cleanenv ${IMAGE_CAPNP_PATH} ls ${GO_FOLDER}/src/capnproto.org/go/capnp/v3/std
    # singularity run -B $GOPATH:$GO_FOLDER --cleanenv ${IMAGE_CAPNP_PATH} env
    cd ..
    singularity run -B $GOPATH:$GO_FOLDER --cleanenv ${IMAGE_CAPNP_PATH} capnp compile -I${GO_FOLDER}/src/capnproto.org/go/capnp/v3/std -ogo:go capnp/revokable_forwarder.capnp 
fi


