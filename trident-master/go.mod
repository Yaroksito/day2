module github.com/netapp/trident

go 1.21

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.8.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v4 v4.1.0
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph v0.7.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures v1.1.0
	github.com/RoaringBitmap/roaring v1.5.0
	github.com/aws/aws-sdk-go-v2 v1.21.1
	github.com/aws/aws-sdk-go-v2/config v1.18.44
	github.com/aws/aws-sdk-go-v2/credentials v1.13.42
	github.com/aws/aws-sdk-go-v2/service/fsx v1.33.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.21.4
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/container-storage-interface/spec v1.8.0
	github.com/docker/go-plugins-helpers v0.0.0-20211224144127-6eecb7beb651
	github.com/dustin/go-humanize v1.0.2-0.20230808044839-52939b84e7e1
	github.com/elastic/go-sysinfo v1.11.1
	github.com/evanphx/json-patch/v5 v5.7.0
	github.com/ghodss/yaml v1.0.1-0.20220118164431-d8423dcdf344 // 1/18/2022
	github.com/go-openapi/errors v0.20.4
	github.com/go-openapi/runtime v0.26.0
	github.com/go-openapi/strfmt v0.21.7
	github.com/go-openapi/swag v0.22.4
	github.com/go-openapi/validate v0.22.1
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/google/go-cmp v0.5.9
	github.com/google/uuid v1.3.1
	github.com/gorilla/mux v1.8.0
	github.com/jarcoal/httpmock v1.3.1
	github.com/kr/secureheader v0.2.0
	github.com/kubernetes-csi/csi-lib-utils v0.15.0
	github.com/kubernetes-csi/csi-proxy/client v1.1.3
	github.com/kubernetes-csi/external-snapshotter/client/v6 v6.3.0
	github.com/mattermost/xml-roundtrip-validator v0.1.1-0.20230502164821-3079e7b80fca
	github.com/mitchellh/copystructure v1.2.0
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/olekukonko/tablewriter v0.0.6-0.20230925090304-df64c4bbad77
	github.com/openshift/api v0.0.0-20230925155749-4c26c14d9ebb
	github.com/prometheus/client_golang v1.16.0
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/afero v1.10.0
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.4
	github.com/vishvananda/netlink v1.1.0
	github.com/zcalusic/sysinfo v1.0.1
	go.uber.org/multierr v1.11.0 // github.com/uber-go/multierr
	golang.org/x/crypto v0.14.0 // github.com/golang/crypto
	golang.org/x/net v0.17.0 // github.com/golang/net
	golang.org/x/oauth2 v0.12.0 // github.com/golang/oauth2
	golang.org/x/sys v0.13.0 // github.com/golang/sys
	golang.org/x/text v0.13.0 // github.com/golang/text
	golang.org/x/time v0.3.0 // github.com/golang/time
	google.golang.org/grpc v1.58.3 // github.com/grpc/grpc-go
	k8s.io/api v0.28.2 // github.com/kubernetes/api
	k8s.io/apiextensions-apiserver v0.28.2 // github.com/kubernetes/apiextensions-apiserver
	k8s.io/apimachinery v0.28.2 // github.com/kubernetes/apimachinery
	k8s.io/client-go v0.28.2 // github.com/kubernetes/client-go
	k8s.io/mount-utils v0.28.2 // github.com/kubernetes/mount-utils
	sigs.k8s.io/cloud-provider-azure/pkg/azclient v0.0.0-20231013064142-1a359a0d4d3b
)

require (
	cloud.google.com/go/compute v1.21.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v1.4.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5 v5.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4 v4.3.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4 v4.2.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns v1.1.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources v1.1.1 // indirect
	github.com/AzureAD/microsoft-authentication-library-for-go v1.1.1 // indirect
	github.com/Microsoft/go-winio v0.6.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.42 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.36 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.44 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.36 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.15.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.17.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.23.1 // indirect
	github.com/aws/smithy-go v1.15.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/elastic/go-windows v1.0.0 // indirect
	github.com/emicklei/go-restful/v3 v3.10.1 // indirect
	github.com/evanphx/json-patch v4.12.0+incompatible // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/analysis v0.21.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/loads v0.21.2 // indirect
	github.com/go-openapi/spec v0.20.8 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/rivo/uniseg v0.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/vishvananda/netns v0.0.0-20191106174202-0a2b9b5464df // indirect
	go.mongodb.org/mongo-driver v1.11.3 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/tools v0.12.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
	k8s.io/klog/v2 v2.100.1 // indirect
	k8s.io/kube-openapi v0.0.0-20230717233707-2695361300d9 // indirect
	k8s.io/utils v0.0.0-20230505201702-9f6742963106 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
