# Creating an application with a Spring Boot code sample

**Note:** The Spring Boot code sample uses the **8085** HTTP port.

Before you begin creating an application with this `devfile` code sample, it's helpful to understand the relationship between the `devfile` and `Dockerfile` and how they contribute to your build. You can find these files at the following URLs:

* [Spring Boot `devfile.yaml`](https://github.com/devfile-samples/devfile-sample-java-springboot-basic/blob/main/devfile.yaml)
* [Spring Boot `Dockerfile`](https://github.com/devfile-samples/devfile-sample-java-springboot-basic/blob/main/docker/Dockerfile)

  
**Example**

`oc new-build https://github.com/amazingkj/okd-spring-build-test.git`

**Start a new build from the update-message branch**

`oc new-build https://github.com/amazingkj/okd-spring-build-test.git#main`

**Use --context-dir to build from a subdirectory**

`oc new-build https://github.com/amazingkj/okd-spring-build-test.git --context-dir example`


### Use S2I in a build

The syntax is the same as normal Builds. OpenShift uses S2I when there is no Dockerfile

**oc new-app works with S2I**

`oc new-app <Git URL with no Dockerfile>`

**oc new-build works with S2I**

`oc new-build <Git URL with no Dockerfile>`

**Example**

`oc new-app https://github.com/amazingkj/okd-spring-build-test.git#nodockerfile`
