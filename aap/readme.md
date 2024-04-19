Automation Controller deployment using argocd:
    while deploying the resources will not deploy and no error is shown, this is dou to missing secrets:
        you can find the following line the in cr yaml asking for the secrets:

            ldap_cacert_secret: ca-bundle-custom-certs-maccabi
            ldap_password_secret: aap-ldap-password

    How to overcome this? 
     - deploy the secrets using agrocd (save the secrest in vault or as yaml)
     - OR create the secrest manually:

         $oc create secret generic ca-bundle-custom-certs-maccabi \
         --from-file=ldap-ca.crt=<PATH/TO/YOUR/CA/PEM/FILE> \
         --from-file=bundle-ca.crt= <PATH/TO/YOUR/CA/PEM/FILE>

        $oc create secret generic aap-ldap-password \
        --from-file=ldap-password=<your_ldap_dn_password>

        Incase you want to test your CA PEM FILE and validate its the right one use the following commnad:
        $openssl s_client -connect mac.org.il:636 -CAfile ldap.pem 
            (ldap.pem is the certificate you want to check)

            if everything is alright you should see the following output:
              SSL handshake has read 3474 bytes and written 474 bytes
              Verification: OK
            meaning its the right certificate!
 