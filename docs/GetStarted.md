# Getting started

- [Getting started](#getting-started)
  - [Create a Test](#create-a-test)
    - [Input Contract](#input-contract)
    - [Export Contract](#export-contract)
    - [Shared Folder](#shared-folder)
    - [Images](#images)
  - [Create a Testrun](#create-a-testrun)
  - [Configuration](#configuration)
    - [Types](#types)
    - [Sources](#sources)
    - [Location](#location)

## Create a Test

A prerequisites for a test to be executed by the TestMachinery is to be available as a GitHub repository.
The repository does not necessarily need to be public.
But if it is private you need to grant the TestMachinery read access to this repository.

Tests in the TestMachinery are configuration files called TestDefinitions which are located in the `repoRoot/.test-defs` folder.
The TestMachinery automatically searches for TestDefinitions files in all specified testlocations (see [testrun description](#create-a-testrun) for more information).
A TestDefinition consists of a name and a description of how the test should be executed by the TestMachinery.
Therefore, a basic test consists of a command which will be executed inside the specified container image.

```yaml
kind: TestDefinition
metadata:
  name: TestDefName
spec:
  owner: gardener@example.com # test owner and contact person in case of a test failure
  recipientsOnFailure: developer1@example.com, developer2@example.com # optional, list of emails to be notified if a step fails
  description: test # optional; description of the test.

  activeDeadlineSeconds: 600 # optional; maximum seconds to wait for the test to finish.
  labels: ["default"] # optional;

  # optional, specify specific behavior of a test.
  # By default steps are executed in parallel.
  # By specifying "serial behavior", tests can be forced to be executed in serial.
  behavior: ["serial"]

  # required; Entrypoint array. Not executed within a shell.
  # The docker image's ENTRYPOINT is used if this is not provided.
  command: [bash, -c]
  # Arguments to the entrypoint. The docker image's CMD is used if this is not provided.
  args: ["test.sh"]

  image: golang:1.11.2 # optional, default image is "eu.gcr.io/gardener-project/gardener/testmachinery/base-step:0.27.0"

  # optional; Configuration of a TestDefinition.
  # Environment Variables can be configured per TestDefinition
  # by specifying the varibale's name and a value, secret or configmap.
  # Files can be mounted into the test by specifying a base64 encoded value, secret or configmap.
  config:
  - type: env
    name: TESTENV1
    value: "Env content"
  - type: env
    name: TESTENV2
    valueFrom:
      secretKeyRef:
        name: secretName
        key: secretKey
  - type: file
    name: file1 # name for description
    path: /tmp/tm/file1.txt
    value: "aGVsbG8gd29ybGQK" # base64 encoded file content: "hello world"
  - type: file
    name: file2
    path: /tmp/tm/file2.txt
    valueFrom:
      configMapKeyRef:
        name: configmapName
        key: configmapKey
```
> Note that the working directory is set to the root of your repository.

### Input Contract

| Environment Variable Name        | Description           |
| ------------- |-------------|
| TM_KUBECONFIG_PATH      | points to a directory containing all kubeconfig files (defaults to `/tmp/env/kubeconfig`). </br> The files contained in this dir depend on the concrete TestRun and can contain up to 3 files: <ul><li>_gardener.config_: the kubeconfig pointing to the gardener cluster created by TM (or predefined/given by a TestRun)</li><li>_seed.config_: the kubeconfig pointing to the seed cluster configured by TM (or predefined/given by a TestRun)</li><li>_shoot.config_: the kubeconfig pointing to the shoot cluster created by TM (or predefined/given by a TestRun)</li></ul>|
| TM_EXPORT_PATH | points to a directory where the test can place arbitrary test-run data which will be archived at the end. Useful if some postprocessing needs to be done on that data. Further information can be found [here](#export-contract) |


### Export Contract

Some installations of the TestMachinery contain a connection to an elasticsearch installation for persistence and evaluation of test results.

The TestMachinery writes some metadata into elasticsearch upon each TestRun completion. It conconsists of the following attributes:
```
tm_meta.landscape: Gardener Landscape of th shoot, e.g. dev, staging,... .
tm_meta.cloudprovider: Cloudprovider of the shoot, e.g. gcp, aws, azure or openstack.
tm_meta.kubernetes_version: Kubernetes version of the shoot.
tm_meta.testrun_id: ID of the overall testrun.
```

TestDefinition can additionally place json files in `TM_EXPORT_PATH` to have them picked up by the TestMachinery and forwarded into elasticsearch.

Such additional data (written by a single test) has to be in one of the 3 formats below. The TestMachinery automatically derives which of the 3 formats is used.
It then automatically uploads these documents to an index named like the TestDefinition. A TestDefinition called `CreateShoot` will be uploaded to the index `createshoot`.

1. Valid JSON document
2. Newline-delimited JSON (multiple json documents in one file, separated by newlines)
    ```
      { "key": "value" }
      { "key2": true }
    ```
3. ElasticSearch bulk format with a specific index (see https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html).
    The documents are then uploaded to the specified index prefixed by `tm-`.
    ```
      { "index": { "_index": "mySpecificIndex", "_type": "_doc" } }
      { "key": "value" }
      { "index": { "_index": "mySpecificIndex", "_type": "_doc" } }
      { "key2": true }
      { "index": { "_index": "mySecondSpecificIndex", "_type": "_doc" } }
      { "key3": 5 }
    ```

### Shared Folder

Data that is stored in `TM_SHARED_PATH` location, can be accessed from within any testflow step of a the workflow. This is essential if e.g. a test flow step needs to evaluate the output of the previously finished test flow step. This folder is also available as an artifact in the Argo UI.

### Images

The Testmachinery provides some images to run your Integration Tests (Dockerfiles can be found in hack/images).

#### Base image (eu.gcr.io/gardener-project/gardener/testmachinery/base-step:0.32.0) <!-- omit in toc -->
- Kubectl
- Helm
- coreutils
- python3
- [cc-utils](https://github.com/gardener/cc-utils) at `/cc/utils` and cli.py added to $PATH
- SAP Root CA

#### Golang image (eu.gcr.io/gardener-project/gardener/testmachinery/golang:0.32.0) <!-- omit in toc -->
- FROM base image
- Golang v1.11.5
- ginkgo test suite at $GOPATH/src/github.com/onsi/ginkgo
- Go project setup script
  - automatically setup test repository at the provided gopath and cd's into it
  - RUN ```/tm/setup github.com/org repo``` e.g. ``` /tm/setup github.com/gardener/test-infra ```

## Create a Testrun
Before a TestDefinition is executed by the TestMachinery, it must be added to a Testrun.
The Testrun can then be deployed to the Kubernetes cluster running a TestMachinery.
It is picked up and the Testflow with the TestDefinitions is executed.

```yaml
apiVersion: testmachinery.sapcloud.io/v1beta1
kind: Testrun
metadata:
  generateName: integration-
  namespace: default
spec:
  ttlSecondsAfterFinished: # optional; define when the Testrun should cleanup itself.

  # TestLocations define where to search for TestDefinitions.
  testLocations:
  - type: git
    repo: https://github.com/gardener/test-infra.git
    revision: master

  # Specific kubeconfigs can be defined for the garden, seed or shoot cluster.
  # These kubeconfigs are available in "TM_KUBECONFIG_PATH/xxx.config" inside every TestDefinition container.
  # Serial steps can also generate new or update current kubeconfigs in this path.
  # Usefull for testing a TestDefinition with a specific shoot.
  kubeconfigs:
    gardener: # base64 encoded gardener cluster kubeconfig
    seed: # base64 encoded seed cluster kubeconfig
    shoot: # base64 encoded shoot cluster kubeconfig


  # Global config available to every test task in all phases (testFlow and onExit)
  config:
    - name: PROJECT_NAMESPACE
      type: env
      value: "garden-it"
    - name: SHOOT_NAME
      type: env
      value: "xxx"

  # The testFlow describes the execution order of the Testrun.
  # It defines which TestDefinition (that can be found in the TestLocations)
  # are executed in which specified order.
  # If a label is specified then all TestDefinitions labeled with the specific key are executed.
  testFlow:
  - - name: CreateShoot
  - - label: default
  - - name: DeleteShoot

  # OnExit specifies the same execution flow as the testFlow.
  # This flow is run after the testFlow and every step can specify the condition
  # under which it should run depending on the outcome of the testFlow.
  onExit:
  - - name: DeleteShoot
      condition: error|success|always # optional; default is always;
```

## Configuration

Test can be configured by passing environment variables to the test or mounting files.
The testmachinery offers 2 types of configuration (Environment Variable and File) and 3 value sources (raw value, secret, configmap).

### Types
Test configuration can be of type "env" and of type "file".

"env" configuration is available as environment variable with the specified `name` to the test.
  ```yaml
  config:
  - type: env
    name: ENV_NAME
  ```

"file" configration is available as mounted file at the specified `path`.
  ```yaml
  config:
  - type: file
    name: file # only for description; no effect on the file itself.
    path: /file/path
  ```

### Sources
The value of a configuration type can be defined by 3 different sources

1. *Value*:<br> The value is directly defined in the yaml. :warning: Vale has to be base64 encoded for config type "file"
  ```yaml
  config:
  - type: env | file
    name: config
    value: "Env content" # or base64 encoded content for files
  ```
2. *Secret*:<br> Value from a secret that is available on the cluster.
  ```yaml
  config:
  - type: env | file
    name: config
    valueFrom:
      configMapKeyRef:
        name: configmapName
        key: configmapKey
  ```
3. *ConfigMap*: <br> Value from a configmap that is available on the cluster
  ```yaml
  config:
  - type: env | file
    name: config
    valueFrom:
      configMapKeyRef:
        name: configmapName
        key: configmapKey
  ```

### Location
This configuration can be defined in 3 possible section:

1. *TestDefinition:* <br>Configurations are testdefinition scoped which means that all configuration is only available for these testdefinitions.
  ```yaml
  kind: TestDefinition
  metadata:
    name: TestDefName
  spec:
    config: # Specify configuration here
  ```
2. *Testrun Step:*<br> Configuration will be available to all tests defined in the specific step.
  ```yaml
  apiVersion: testmachinery.sapcloud.io/v1beta1
  kind: Testrun
  metadata:
    generateName: integration-
    namespace: default
  spec:
    testFlow:
     - - label: default
         config: # Specify configuration here
  ```
3. *Testrun Global:*<br> Configuration will be available to all tests.
  ```yaml
  apiVersion: testmachinery.sapcloud.io/v1beta1
  kind: Testrun
  metadata:
    generateName: integration-
    namespace: default
  spec:
    config: # Specify configuration here
  ```
