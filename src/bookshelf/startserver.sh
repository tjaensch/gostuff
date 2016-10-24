#!/bin/sh

ipaddress=${hostname -I}

hugo server --bind=$ipaddress --baseUrl=http://$ipaddress:1313 --buildDrafts
