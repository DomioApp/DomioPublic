#!/usr/bin/env bash
set -e

templates_folder=/usr/local/domio_public/templates

echo Copying templates folder...

rm -rf ${templates_folder}
yes | mv -f ~/domiopublic/src/domio_public/templates ${templates_folder}

echo Templates copied!