Manual systemReserved for OCP calaulation:

https://access.redhat.com/solutions/5843241


Node    CPU  RAM
Master  4    16GB (14.9012Gib)
Worker  8    32GB (29.8023Gib)
Infra   16   64GB (59.6046Gib)


CPU
Master 60m + 10m + 10m = 80m (0.08)
Worker 60m + 10m + 10m + $(2.5m X 4) = 90m (0.09)
Infra 60m + 10m + 10m + $(2.5m X 12) = 110m (0.11)

RAM
Master 1GiB + 0.8GiB + 0.6GiB = 2.4GiB
Worker 1GiB + 0.8GiB + 0.8GiB + 0.8GiB = 3.4GiB
Infra 1GiB + 0.8GiB + 0.8GiB + 2.6GiB = 5.2GiB



Implementation: https://docs.openshift.com/container-platform/4.7/nodes/nodes/nodes-nodes-managing.html

What is the default system reservation: https://access.redhat.com/solutions/5215901

systemReserved:
  cpu: 500m
  memory: 1Gi
  ephemeral-storage: 1Gi

How To ?

# Create label for the relevant MCP

oc label machineconfigpool worker custom-kubelet=enabled
oc label machineconfigpool master custom-kubelet=enabled

#show label
oc get machineconfigpool  --show-labels

# apply the MCP

oc create -f system-reserved-worker.yaml