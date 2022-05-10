#!/bin/bash

rm ginode-mongo
rz
export GIN_MODE=release
chmod u+x ginode-mongo
pm2 restart ginode-mongo