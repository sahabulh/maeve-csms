#!/bin/sh

# ===============================================================================================================
# This shell script can be used to create all necessary certificates needed to
# - successfully perform a TLS handshake between the EVCC (TLS client) and the
#   SECC (TLS server) and
# - install a contract certificate in the EVCC.
#
# This file contains all information needed to create your own certificate chains.
#
# Helpful information about using OpenSSL is provided by Ivan Ristic's book
# "Bulletproof SSL and TLS". Furthermore, you should have OpenSSL >= 1.0.2
# installed to comply with all security requirements of ISO 15118.
# Some MacOS X installations unfortunately still use openssl < v1.0.2. You could use Homebrew to install openssl.
#
# Author: Dr. Marc Mültin (marc@switch-ev.com)
# ===============================================================================================================


# Change the validity periods (given in number of days) to test
# - valid certificates (e.g. contract certificate or Sub-CA certificate)
# - expired certificates (e.g. contract certificate or Sub-CA certificates)
#   -> you need to reset your system time to the past to create expired certificates
# - a to-be-updated contract certificate

VALIDITY_CSMS_SERVER_CERT=365

ISO_2="iso-2"
ISO_20="iso-20"

usage() {
  echo "
  Usage: "$0" [-h] [-v <iso-2|iso-20>] [-p password] [-k]

  Options:
   -h --help          Returns this helper
   -v --version       ISO version to run the script for: 'iso-2' refers to ISO 15118-2,
                      whereas 'iso-20' refers to ISO 15118-20
   -p --password      The password to encrypt and decrypt the private keys
   -k --keysight      Generate certificates to be used while pairing with Keysight test system,
                      alongside this iso15118 project.
   -i --install       Copy the required files inside everest-core/config/certs 


  Description:
    You can use this script to create all the private keys and public key certificates
    necessary to run an ISO 15118 Plug & Charge session (for testing purposes only).
    You need to provide the ISO 15118 version with the '-v' flag, choosing between the
    two admissible options stated above.

    This script uses by default a password value of 12345, which can be modified
    if -p option is used as exemplified above.

    NOTE: This script will create the following folder structure, if not already
    existing, depending on the protocol version you choose and place the corresponding
    private keys, certificate signing requests (csrs), and certificates (certs) in the
    right folder, overwriting any files with the same name:

    |__ iso15118_2 (or iso15118_20)
      |__ certs
      |__ csrs
      |__ private_keys
  "
  exit 0;
}

validate_option() {
    # Check if the version provided is valid, if not it returns
    if [ "$1" != $ISO_2 ] && [ "$1" != $ISO_20 ]; then
        echo "The version provided is invalid"
        usage
    fi
}


if [ -z $1 ]; then echo "No options were provided"; usage; fi

while [ -n "$1" ]; do
    case "$1" in
        -h|--help)
            usage
            ;;
        -v|--version)
            validate_option $2
            version=$2
            shift  # params with args need an extra shift
            ;;
        -p|--password)
            password=$2
            shift
            ;;
        -k|--keysight)
            keysight_certs="1"
            ;;
        -i|--install)
            install_everest_core="1"
            everest_core_path=$2
            shift
            ;;
         *)
            echo "Unknown option $1"
            usage
            ;;
    esac
    shift
done


# Set the cryptographic parameters, depending on whether to create certificates and key
# material for ISO 15118-2 or ISO 15118-20

if [ $version = $ISO_2 ];
then
    ISO_FOLDER=iso15118_2
    SYMMETRIC_CIPHER=-aes-128-cbc
    SYMMETRIC_CIPHER_PKCS12=-aes128
    SHA=-sha256
    # Note: OpenSSL does not use the named curve 'secp256r1' (as stated in
    # ISO 15118-2) but the equivalent 'prime256v1'
    EC_CURVE=prime256v1
else
    ISO_FOLDER=iso15118_20
    SYMMETRIC_CIPHER=-aes-128-cbc  # TODO Check correct version for ISO 15118-20
    SYMMETRIC_CIPHER_PKCS12=-aes128  # TODO Check correct version for ISO 15118-20
    SHA=-sha256  # TODO Check correct version for ISO 15118-20
    EC_CURVE=prime256v1  # TODO Check correct version for ISO 15118-20
    # TODO: Also enable cipher suite TLS_CHACHA20_POLY1305_SHA256
fi

# The password used to encrypt (and decrypt) private keys
# Security note: this is for testing purposes only!
if [ -z $password ]; then
    password=123456
fi

echo "Password used is: '$password'"

openssl ecparam -genkey -name $EC_CURVE | openssl ec $SYMMETRIC_CIPHER -passout pass:$password -out csms.key
openssl req -new -key csms.key -passin pass:$password -config csmsLeafCert.cnf -out csms.csr
openssl x509 -req -in csms.csr -extfile csmsLeafCert.cnf -extensions ext -CA cpo_sub_ca2.pem -CAkey cpo_sub_ca2.key -passin pass:$password -set_serial 12380 -days $VALIDITY_CSMS_SERVER_CERT -out csmsLeaf.pem
cat csmsLeaf.pem cpo_sub_ca2.pem cpo_sub_ca1.pem > csms.pem

openssl ec -in csms.key -passin pass:$password -out csms.key

rm csms.csr
