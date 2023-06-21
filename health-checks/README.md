# Health Checks

This is a simple app to test Kubernetes container health check configurations.

## Getting started

```
oc new-project health-checks
oc apply -f https://raw.githubusercontent.com/frarodroc/kubernetes/master/health-checks/health-checks.yaml
```

Force the liveness probe to fail:

```
oc rsh $(oc get pod -o name) rm -f /tmp/healthy
```

Check for the container failure:

```
watch oc get po
```

Sample output:

```
NAME                             READY   STATUS    RESTARTS       AGE
health-checks-XXXXXXXXXX-XXXXX   1/1     Running   1 (102s ago)   38m
```

## Authors

* Francisco Rodríguez Rocha
