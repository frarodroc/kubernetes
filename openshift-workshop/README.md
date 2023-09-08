# OpenShift Workshop

This YAML file creates an "admin" user, 50 regular users and 50 namespaces.

## Getting Started

The passwords are:

* admin: openshift
* user: P4ssw0rd

```
oc apply -f https://raw.githubusercontent.com/frarodroc/kubernetes/master/openshift-workshop/openshift-workshop.yaml
```

To create it using the web console you must first delete the OAuth object that is created by default:

```
oc delete oauth cluster
oc create -f https://raw.githubusercontent.com/frarodroc/kubernetes/master/ope>
```



## Authors

* Francisco Rodríguez Rocha
