#!/usr/bin/env bash
set -e

echo Copying Domio Service Config...

yes | cp -rf ~/domiopublic/deploy/config/domio_service.sh /etc/init.d/domio_public

systemctl daemon-reload
service domio_public stop
service domio_public start

echo Domio Service Config copied!