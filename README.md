# Scout

Scout is a simple Kubernetes CLI tool that converts raw Kubernetes data into clean, readable and actionable output.

Instead of overwhelming users with hundreds of lines from `kubectl`, Scout focuses on presenting only the information that matters in a simplified and pretty format.

> Scout doesn't dump Kubernetes data. Scout explains it.

---

## Features

### Pod Overview

```bash
scout pods
```

<img width="737" height="335" alt="image" src="https://github.com/user-attachments/assets/55960362-20d8-4544-b19b-7a97af34d2a9" />


Get a quick overview of your cluster's pods.

Displays:

- Pod names
- Current status
- Restart count
- Health indicators

---

### Logs

```bash
scout logs <pod-name>
```

<img width="637" height="182" alt="image" src="https://github.com/user-attachments/assets/0a5483dd-93b5-4252-8b5d-f5e2b629e892" />


Read application logs in a much cleaner format.

Features:

- Detects `INFO`, `WARN` and `ERROR`
- Filters noisy output
- Makes logs easier to scan

---

### Events

```bash
scout events
```

<img width="1253" height="547" alt="image" src="https://github.com/user-attachments/assets/6a84cb44-6274-47eb-94b8-58dac1e08502" />


View cluster events in a simplified way.

Useful for quickly spotting:

- Scheduling issues
- Image pull failures
- Resource constraints
- Volume mount problems

---

### Inspect (Diagnosis Mode)

```bash
scout inspect pod <pod-name>
```

<img width="1527" height="212" alt="image" src="https://github.com/user-attachments/assets/90a27b49-b62f-4b00-8661-c7ab58840898" />

<img width="400" height="150" alt="image" src="https://github.com/user-attachments/assets/a2209007-a534-4aff-ac90-b536688e406b" />



One of Scout's core features.

Instead of reading hundreds of lines from `kubectl describe`, Scout analyzes the pod and provides a concise diagnosis.

Shows:

- Pod status
- Scheduling failures
- Restart information
- Error messages
- Human-readable hints


### Resource Utilization

```bash
scout top
```

<img width="688" height="362" alt="image" src="https://github.com/user-attachments/assets/cd010144-18ce-46ef-a0d2-c3a18f175a02" />


Monitor resource usage for each pod , Honestly This Command just exist .

Displays:

- CPU usage
- Memory usage

> Requires metrics-server.

---

## Tech Stack

- Go
- Cobra
- Kubernetes client-go
- Kubernetes Metrics API
- GitHub Actions

---

## License

MIT
