// Code generated for package assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/openshift-kube-conformance.yaml
// manifests/openshift-provider-cert-level-1.yaml
// manifests/openshift-provider-cert-level-2.yaml
// manifests/openshift-provider-cert-level-3.yaml
package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _manifestsOpenshiftKubeConformanceYaml = []byte(`podSpec:
  restartPolicy: Never
  serviceAccountName: sonobuoy-serviceaccount
  tolerations:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
    operator: Exists
  - key: CriticalAddonsOnly
    operator: Exists
  containers:
    - name: report-progress
      image: quay.io/mrbraga/openshift-tests-provider-cert:latest
      imagePullPolicy: Always
      priorityClassName: system-node-critical
      command: ["./report-progress.sh"]
      volumeMounts:
      - mountPath: /tmp/sonobuoy/results
        name: results
      env:
        - name: CERT_LEVEL
          value: "0"

sonobuoy-config:
  driver: Job
  plugin-name: openshift-kube-conformance
  result-format: junit
  description: The end-to-end tests maintained by Kubernetes to certify the platform.
  source-url: https://github.com/openshift/provider-certification-tool/blob/mvp/tools/plugins/openshift-kube-conformance.yaml
  skipCleanup: true
spec:
  name: plugin
  image: quay.io/mrbraga/openshift-tests-provider-cert:latest
  imagePullPolicy: Always
  priorityClassName: system-node-critical
  resources: {}
  volumeMounts:
  - mountPath: /tmp/sonobuoy/results
    name: results
  env:
    - name: CERT_LEVEL
      value: "0"
`)

func manifestsOpenshiftKubeConformanceYamlBytes() ([]byte, error) {
	return _manifestsOpenshiftKubeConformanceYaml, nil
}

func manifestsOpenshiftKubeConformanceYaml() (*asset, error) {
	bytes, err := manifestsOpenshiftKubeConformanceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/openshift-kube-conformance.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsOpenshiftProviderCertLevel1Yaml = []byte(`podSpec:
  restartPolicy: Never
  serviceAccountName: sonobuoy-serviceaccount
  tolerations:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
    operator: Exists
  - key: CriticalAddonsOnly
    operator: Exists
  containers:
    - name: report-progress
      image: quay.io/mrbraga/openshift-tests-provider-cert:latest
      imagePullPolicy: Always
      priorityClassName: system-node-critical
      command: ["./report-progress.sh"]
      volumeMounts:
      - mountPath: /tmp/sonobuoy/results
        name: results
      env:
        - name: CERT_LEVEL
          value: "1"

sonobuoy-config:
  driver: Job
  plugin-name: openshift-provider-cert-level1
  result-format: junit
  description: The end-to-end tests maintained by OpenShift to certify the Provider running the OpenShift Container Platform.
  source-url: https://github.com/openshift/provider-certification-tool/blob/mvp/tools/plugins/openshift-provider-cert-level-1.yaml
  skipCleanup: true
spec:
  name: plugin
  image: quay.io/mrbraga/openshift-tests-provider-cert:latest
  imagePullPolicy: Always
  priorityClassName: system-node-critical
  resources: {}
  volumeMounts:
  - mountPath: /tmp/sonobuoy/results
    name: results
  env:
    - name: CERT_LEVEL
      value: "1"
`)

func manifestsOpenshiftProviderCertLevel1YamlBytes() ([]byte, error) {
	return _manifestsOpenshiftProviderCertLevel1Yaml, nil
}

func manifestsOpenshiftProviderCertLevel1Yaml() (*asset, error) {
	bytes, err := manifestsOpenshiftProviderCertLevel1YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/openshift-provider-cert-level-1.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsOpenshiftProviderCertLevel2Yaml = []byte(`podSpec:
  restartPolicy: Never
  serviceAccountName: sonobuoy-serviceaccount
  tolerations:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
    operator: Exists
  - key: CriticalAddonsOnly
    operator: Exists
  containers:
    - name: report-progress
      image: quay.io/mrbraga/openshift-tests-provider-cert:latest
      imagePullPolicy: Always
      priorityClassName: system-node-critical
      command: ["./report-progress.sh"]
      volumeMounts:
      - mountPath: /tmp/sonobuoy/results
        name: results
      env:
        - name: CERT_LEVEL
          value: "2"

sonobuoy-config:
  driver: Job
  plugin-name: openshift-provider-cert-level2
  result-format: junit
  description: The end-to-end tests maintained by OpenShift to certify the Provider running the OpenShift Container Platform.
  source-url: https://github.com/openshift/provider-certification-tool/blob/mvp/tools/plugins/openshift-provider-cert-level-2.yaml
  skipCleanup: true
spec:
  name: plugin
  image: quay.io/mrbraga/openshift-tests-provider-cert:latest
  imagePullPolicy: Always
  priorityClassName: system-node-critical
  resources: {}
  volumeMounts:
  - mountPath: /tmp/sonobuoy/results
    name: results
  env:
    - name: CERT_LEVEL
      value: "2"
`)

func manifestsOpenshiftProviderCertLevel2YamlBytes() ([]byte, error) {
	return _manifestsOpenshiftProviderCertLevel2Yaml, nil
}

func manifestsOpenshiftProviderCertLevel2Yaml() (*asset, error) {
	bytes, err := manifestsOpenshiftProviderCertLevel2YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/openshift-provider-cert-level-2.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsOpenshiftProviderCertLevel3Yaml = []byte(`podSpec:
  restartPolicy: Never
  serviceAccountName: sonobuoy-serviceaccount
  tolerations:
  - effect: NoSchedule
    key: node-role.kubernetes.io/master
    operator: Exists
  - key: CriticalAddonsOnly
    operator: Exists
  containers:
    - name: report-progress
      image: quay.io/mrbraga/openshift-tests-provider-cert:latest
      imagePullPolicy: Always
      priorityClassName: system-node-critical
      command: ["./report-progress.sh"]
      volumeMounts:
      - mountPath: /tmp/sonobuoy/results
        name: results
      env:
        - name: CERT_LEVEL
          value: "3"

sonobuoy-config:
  driver: Job
  plugin-name: openshift-provider-cert-level3
  result-format: junit
  description: The end-to-end tests maintained by OpenShift to certify the Provider running the OpenShift Container Platform.
  source-url: https://github.com/openshift/provider-certification-tool/blob/mvp/tools/plugins/openshift-provider-cert-level-3.yaml
  skipCleanup: true
spec:
  name: plugin
  image: quay.io/mrbraga/openshift-tests-provider-cert:latest
  imagePullPolicy: Always
  priorityClassName: system-node-critical
  resources: {}
  volumeMounts:
  - mountPath: /tmp/sonobuoy/results
    name: results
  env:
    - name: CERT_LEVEL
      value: "3"
`)

func manifestsOpenshiftProviderCertLevel3YamlBytes() ([]byte, error) {
	return _manifestsOpenshiftProviderCertLevel3Yaml, nil
}

func manifestsOpenshiftProviderCertLevel3Yaml() (*asset, error) {
	bytes, err := manifestsOpenshiftProviderCertLevel3YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/openshift-provider-cert-level-3.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"manifests/openshift-kube-conformance.yaml":      manifestsOpenshiftKubeConformanceYaml,
	"manifests/openshift-provider-cert-level-1.yaml": manifestsOpenshiftProviderCertLevel1Yaml,
	"manifests/openshift-provider-cert-level-2.yaml": manifestsOpenshiftProviderCertLevel2Yaml,
	"manifests/openshift-provider-cert-level-3.yaml": manifestsOpenshiftProviderCertLevel3Yaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"manifests": &bintree{nil, map[string]*bintree{
		"openshift-kube-conformance.yaml":      &bintree{manifestsOpenshiftKubeConformanceYaml, map[string]*bintree{}},
		"openshift-provider-cert-level-1.yaml": &bintree{manifestsOpenshiftProviderCertLevel1Yaml, map[string]*bintree{}},
		"openshift-provider-cert-level-2.yaml": &bintree{manifestsOpenshiftProviderCertLevel2Yaml, map[string]*bintree{}},
		"openshift-provider-cert-level-3.yaml": &bintree{manifestsOpenshiftProviderCertLevel3Yaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}