module github.com/pulumi/pulumi-vsphere

go 1.12

require (
	github.com/apparentlymart/go-rundeck-api v0.0.0-20160826143032-f6af74d34d1e // indirect
	github.com/fsouza/go-dockerclient v0.0.0-20160427172547-1d4f4ae73768 // indirect
	github.com/hashicorp/terraform v0.12.0-rc1.0.20190509225429-28b2383eacae
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v0.17.6-0.20190410045519-ef5e148a73c0
	github.com/pulumi/pulumi-terraform v0.18.3-0.20190604214533-7ace3e9b5f2d
	github.com/rancher/go-rancher v0.0.0-20170407040943-ec24b7f12fca // indirect
	github.com/stretchr/testify v1.3.1-0.20190311161405-34c6fa2dc709
	github.com/terraform-providers/terraform-provider-vsphere v1.11.0
	k8s.io/kubernetes v1.6.1 // indirect
)

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.3.0
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
)
