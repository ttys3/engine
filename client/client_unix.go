// +build linux freebsd openbsd darwin

package client

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///run/podman/podman.sock"
