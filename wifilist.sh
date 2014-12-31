#!/bin/ash
# Shell script running on raspberry pi.
# Returns JSON with names of available SSIDs. eg. {"network": [{"SSID":"RippleOS"},{"SSID":"homewiFi"}]}

SSID_list=`iw dev wlan0 scan | grep SSID | awk {'print $2'} | awk ' BEGIN { ORS = ""; print "["; } { print "@{\"SSID\":\""$0"\"}@" } END { print "]"; }' | sed "s^\"^\"^g;s^@@^, ^g;s^@^^g"`

printf '{"network": '
printf %s ${SSID_list}}
