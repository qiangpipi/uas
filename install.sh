#!/bin/bash
#Check if running as root
id=`id -u`
if [ 0 != ${id}  ];then
  echo "Please run install as root"
  exit 1
fi
#Check if sldap installed
ldap=`service slapd status|grep -c "is running"`
if [ 1 != ${ldap} ];then
  echo "Please install/start slapd service"
  exit 1
fi
#Check if ldapscripts installed
#Check if the basic ldif loaded
#Create folder
mkdir -p /usr/local/uas
#Copy binary to /usr/local/uas
cp uas /usr/local/uas/.
cp uas.sh /usr/local/uas/.
cp -r ./data /usr/local/uas/.
cp -r ./conf /usr/local/uas/.
#Create link in /usr/bin
ln -sf /usr/local/uas/uas.sh /usr/bin/uas
