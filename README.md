# Hivebox - DevOpsRoadmap Project

Hivebox is a comprehensive DevOps learning project that implements modern cloud-native practices and technologies. This project serves as a practical guide through the DevOps roadmap, demonstrating real-world implementation of infrastructure as code, containerization, CI/CD pipelines, monitoring, and more.

## DevOps Roadmap

This project follows the comprehensive DevOps roadmap available at [https://devopsroadmap.io/](https://devopsroadmap.io/).

[![Dynamic DevOps Roadmap](https://devopshive.net/badges/dynamic-devops-roadmap.svg)](https://github.com/DevOpsHiveHQ/dynamic-devops-roadmap)

---

# Project Progress

## Phase 1

- [x] Understand your role in this project and how you work with other teams.

- [x] Brush up on your knowledge about Software Project management (Hint: What is agile project management? And Introduction to Software Product Management).

- [x] Decide which Agile methodology you will use. (Scrum, Kanban, Scrumban, etc.).

- [x] Document as you go. Always assume that someone else will read your project at any phase.

- [x] Avoid Scope Creep! Make it work, then make it right, then make it fast!

- [ ] Each change should be done in order and using pull requests (don't push directly to the main branch!).

---

## Phase 2

- [x] Create GitHub repository for the project.
- [x] Implement the code requirements.
- [x] Create a function that print current app version. It should print the version then exit the application.

```go
mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"version": APP_VERSION})
})
```

- [x] Use Semantic Versioning for the app version starting with v0.0.1.

```go
const APP_VERSION = "v0.0.1"
```

- [x] Create Dockerfile for the project.
- [x] Build the Docker image and run it locally.
- [x] Locally, run the app container and ensure that it returns the correct value.

## Phase 3

- [ ] Use Conventional Commits for Git commits
- [x] Familiarize yourself with openSenseMap API
- [x] Implement the code requirements (We Used net/http base package [there is change of migrating to gorilla/mux])

  - [x] Path: /version
    - [x] Returns the version of the deployed app
  - [x] Temperature endpoint:
    - [x] Path: /temperature
    - [x] Return current average temperature based on all senseBox data
    - [x] Ensure data is no older than 1 hour

- [x] Apply Best Practices for containers [Docker Best Practices! - Abouzaid](https://tech.aabouzaid.com/2021/09/docker-best-practices-workshop-presentation.html)

- [x] Create a GitHub Actions workflow for CI:

  - [x] Add step to lint code and Dockerfile
  - [x] Add step to build the Docker image
  - [x] Add step to run unit tests
  - [x] Setup OpenSSF Scorecard GitHub Action and fix reported issues
  - [x] In the CI pipeline, call the /version endpoint and verify correct response

- [x] Write unit tests for all endpoints
