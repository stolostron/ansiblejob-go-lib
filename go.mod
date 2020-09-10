module github.com/open-cluster-management/ansiblejob-go-lib

go 1.15

require (
	cloud.google.com/go v0.49.0 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20200414173820-0848c9571904 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.0.0-20200403190813-44a64ad78b9b // indirect
	google.golang.org/appengine v1.6.5 // indirect
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v12.0.0+incompatible
	sigs.k8s.io/controller-runtime v0.6.0
)

replace k8s.io/client-go => k8s.io/client-go v0.18.2
