#https://docs.openshift.com/container-platform/4.10/security/certificates/replacing-default-ingress-certificate.html

#by default ingress using self-sign ceritifcate, need to replace it with corporate certificate.
#wildcard....yaml holding the certificate (private) that why is mandatory to enrypt it.
#before replaccing wildcard....yaml need to generate ceritifcate and convert it from pfx to crt+key.
#https://www.ibm.com/docs/en/arl/9.7?topic=certification-extracting-certificate-keys-from-pfx-file

#to encrypt the secret file please use sealed-secrets.