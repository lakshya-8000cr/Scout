# Scout

Scout is a Kubernetes CLI tool that transforms raw Kubernetes data into clean, readable and actionable insights.

Instead of overwhelming users with hundreds of lines from `kubectl`, Scout focuses on surfacing only the information that matters in a simplified and human-friendly format.

> Scout doesn't dump Kubernetes data. Scout explains it.

---

## Features

### Pod Overview

```bash
scout pods
```

<img width="737" height="335" alt="image" src="https://github.com/user-attachments/assets/55960362-20d8-4544-b19b-7a97af34d2a9" />

Get a quick overview of your cluster.

Displays:

* Pod names
* Current status
* Restart count
* Health indicators

Supports:

```bash
scout pods --namespace forge-apps
```

---

### Logs

```bash
scout logs <pod-name>
```

<img width="637" height="182" alt="image" src="https://github.com/user-attachments/assets/0a5483dd-93b5-4252-8b5d-f5e2b629e892" />

Read application logs in a cleaner and easier-to-scan format.

Features:

* Detects `INFO`, `WARN` and `ERROR`
* Reduces noisy output
* Makes debugging easier

Supports:

```bash
scout logs <pod-name> --namespace <app-name>
```

---

### Events

```bash
scout events
```

<img width="1253" height="547" alt="image" src="https://github.com/user-attachments/assets/6a84cb44-6274-47eb-94b8-58dac1e08502" />

View Kubernetes events in a simplified way.

Useful for quickly spotting:

* Scheduling issues
* Image pull failures
* Resource constraints
* Volume mount problems

Supports:

```bash
scout events --namespace <app-name>
```

---

### Inspect (Diagnosis Mode)

```bash
scout inspect pod <pod-name>
```

<img width="1527" height="212" alt="image" src="https://github.com/user-attachments/assets/90a27b49-b62f-4b00-8661-c7ab58840898" />

<img width="400" height="150" alt="image" src="https://github.com/user-attachments/assets/a2209007-a534-4aff-ac90-b536688e406b" />

One of Scout's core features.

Instead of reading hundreds of lines from `kubectl describe`, Scout analyzes the pod and provides a concise diagnosis.

Displays:

* Pod status
* Scheduling failures
* Restart information
* Error messages
* Human-readable hints

  Supports:

```bash
scout inspect pod <pod-name> --namespace <app-name>
```

---

### Resource Utilization

```bash
scout top
```

<img width="688" height="362" alt="image" src="https://github.com/user-attachments/assets/cd010144-18ce-46ef-a0d2-c3a18f175a02" />

Monitor CPU and memory usage for each pod.

Displays:

* CPU usage
* Memory usage

> Requires metrics-server.

Supports:

```bash
scout top --namespace <app-name>
```


---

## Installation

### Build locally

```bash
go build -o scout.exe
```

### Run

```bash
scout pods

scout logs <pod-name>

scout events

scout inspect pod <pod-name>

scout top
```

---

## Tech Stack

* Go
* Cobra
* Kubernetes client-go
* Kubernetes Metrics API
* GitHub Actions

### Prerequisites

- A Kubernetes cluster (Minikube, EKS, Kind, etc.)
- A valid `~/.kube/config`
- `metrics-server` enabled for `scout top`

---

## License

MIT

