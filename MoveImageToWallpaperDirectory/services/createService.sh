#!/bin/bash

sudo cp  moveImages.service /etc/systemd/system/

systemctl enable moveImages.service

systemctl start moveImages.service
