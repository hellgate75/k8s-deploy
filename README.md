# k8s-deploy
Kubernetes Chart/Files Repository, Scheduler, Executor components composing the Kubernetes Continuous Delivery Framework.

Framework Component(s):
* Registry and Registry Api Manager
* Scheduler and Scheduler Api Manager
* Executor backend service
* Remote Data Services: MongoDb Remote Instance


## Kubernetes Continuous Delivery Manager

Process flow has been splitted in interaction between provided components.

Registry allows insert of process flow components:
* Charts or Kubernetes Files
* Repositories (containing charts and/or Kubernetes Files)
* Projects (containing reference to charts / Kubernetes files)
* Instance (Instance of components of a job, based on a project version, variables list and rules)
* Job (Containing execution information and the instance of execution objects, variables, etc...)
* Deploy (Containing job list, defining the deployment list)
