#!/bin/bash
echo "Starting to build the arozos"
cd arozos/src/
go build

echo "Restarting arozos service with systemctl"
sudo systemctl start arozos

echo "Build completed"