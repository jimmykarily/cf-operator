## cf-operator util variable-interpolation

Interpolate variables

### Synopsis

Interpolate variables of a manifest:

This will interpolate all the variables from a folder and write an
interpolated manifest to STDOUT


```
cf-operator util variable-interpolation [flags]
```

### Options

```
  -m, --bosh-manifest-path string   (BOSH_MANIFEST_PATH) path to the bosh manifest file
  -h, --help                        help for variable-interpolation
      --output-file-path string     (OUTPUT_FILE_PATH) Path of the file to which json output is written.
  -v, --variables-dir string        (VARIABLES_DIR) path to the variables dir
```

### Options inherited from parent commands

```
      --apply-crd                                (APPLY_CRD) If true, apply CRDs on start (default true)
      --bosh-dns-docker-image string             (BOSH_DNS_DOCKER_IMAGE) The docker image used for emulating bosh DNS (a CoreDNS image) (default "coredns/coredns:1.6.3")
  -n, --cf-operator-namespace string             (CF_OPERATOR_NAMESPACE) The operator namespace (default "default")
      --cluster-domain string                    (CLUSTER_DOMAIN) The Kubernetes cluster domain (default "cluster.local")
      --ctx-timeout int                          (CTX_TIMEOUT) context timeout for each k8s API request in seconds (default 30)
  -o, --docker-image-org string                  (DOCKER_IMAGE_ORG) Dockerhub organization that provides the operator docker image (default "cfcontainerization")
      --docker-image-pull-policy string          (DOCKER_IMAGE_PULL_POLICY) Image pull policy (default "IfNotPresent")
  -r, --docker-image-repository string           (DOCKER_IMAGE_REPOSITORY) Dockerhub repository that provides the operator docker image (default "cf-operator")
  -t, --docker-image-tag string                  (DOCKER_IMAGE_TAG) Tag of the operator docker image (default "0.0.1")
  -c, --kubeconfig string                        (KUBECONFIG) Path to a kubeconfig, not required in-cluster
  -l, --log-level string                         (LOG_LEVEL) Only print log messages from this level onward (default "debug")
      --max-boshdeployment-workers int           (MAX_BOSHDEPLOYMENT_WORKERS) Maximum number of workers concurrently running BOSHDeployment controller (default 1)
      --max-quarks-secret-workers int            (MAX_QUARKS_SECRET_WORKERS) Maximum number of workers concurrently running QuarksSecret controller (default 5)
      --max-quarks-statefulset-workers int       (MAX_QUARKS_STATEFULSET_WORKERS) Maximum number of workers concurrently running QuarksStatefulSet controller (default 1)
  -w, --operator-webhook-service-host string     (CF_OPERATOR_WEBHOOK_SERVICE_HOST) Hostname/IP under which the webhook server can be reached from the cluster
  -p, --operator-webhook-service-port string     (CF_OPERATOR_WEBHOOK_SERVICE_PORT) Port the webhook server listens on (default "2999")
  -x, --operator-webhook-use-service-reference   (CF_OPERATOR_WEBHOOK_USE_SERVICE_REFERENCE) If true the webhook service is targeted using a service reference instead of a URL
      --watch-namespace string                   (WATCH_NAMESPACE) Namespace to watch for BOSH deployments
```

### SEE ALSO

* [cf-operator util](cf-operator_util.md)	 - Calls a utility subcommand

###### Auto generated by spf13/cobra on 8-Jan-2020
