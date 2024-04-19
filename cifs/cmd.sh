
# DaemonSet https://raw.githubusercontent.com/kubernetes-csi/csi-driver-smb/master/deploy/csi-smb-node.yaml
# References: https://www.redhat.com/en/blog/cifs-and-openshift-using-the-container-storage-interface-1

oc create secret generic cifs-csi-credentials --from-literal=username=<username> --from-literal=password=<password>



oc adm policy add-scc-to-user -z csi-smb-node-sa privileged


oc -n cifs-csi-demo  create secret generic cifs-csi-creds --from-literal=username=WotchQASA\@mac.org.il --from-literal=password='5@#LIZ$PviY6'
oc -n cifs-csi-demo  create secret generic cifs-csi-creds --from-literal=username=WotchQASA --from-literal=password=5AHLIZDPviY6


# New csidrivers.storage.k8s.io has deployed
oc create -f csi-smb-driver.yaml

# Deploy CSI SMB Nodes DaemonSet
oc create -f csi-smb-node-privileged-affinity.yaml

# Deploy DaemonSet pods on kube-system
oc apply -f csi-smb-node-012024.yaml

# Create secret creds to access share in the cifs-csi-demo project
oc -n cifs-csi-demo  create secret generic cifs-csi-creds --from-literal=username=WotchQASA --from-literal=password=5AHLIZDPviY6

# Create PVC to store the files from the share

# Deploy the depolyment 
