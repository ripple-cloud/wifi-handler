#!/bin/sh
# Shell script running on raspberry pi.
# Returns JSON array with names of available SSIDs. eg. ["RippleOS","homewiFi"]

#SSID_list=`ls | awk ' BEGIN { ORS = ""; print "["; } { print "@\""$0"\"@" } END { print "]"; }' | sed "s^\"^\"^g;s^@@^, ^g;s^@^^g"`

SSID_list=`iw dev wlan0 scan | grep SSID | awk ' BEGIN { ORS = ""; print "["; } { print "@\""$0"\"@" } END { print "]"; }' | sed "s^\"^\"^g;s^@@^, ^g;s^@^^g"`

echo ${SSID_list} 
