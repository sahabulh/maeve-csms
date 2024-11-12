#!/bin/sh

cp certs/ca/csms/CPO_SUB_CA2.pem cpo_sub_ca2.pem
cp certs/ca/csms/CPO_SUB_CA1.pem cpo_sub_ca1.pem

cp certs/client/csms_server/CSMS_SERVER.pem csmsLeaf.pem
cp certs/client/csms_server/CSMS_SERVER.key csms.key

cat csmsLeaf.pem cpo_sub_ca2.pem cpo_sub_ca1.pem > csms.pem
cat cpo_sub_ca2.pem cpo_sub_ca1.pem > trust.pem
cp certs/ca/v2g/V2G_ROOT_CA.pem root-V2G-cert.pem
cp certs/ca/mo/MO_ROOT_CA.pem root-MO-cert.pem
