Locker
===

Locker is a container runtime example to show that containers are just processes with isolation. We can provide this isolation simply, in a way that's easy to demonstrate.

The key things that you need for a container is a process to run, process namespacing (A private /proc), Hostname isolation (or a private Unix Timesharing System), filesystem namespacing (a private namespace), network isolation (ommitted for brevity).

If you're coming from Docker, the process you run is either explicitly defined in the Dockerfile as ENTRYPOINT, or you provide it when the container is run with something akin to:

```sh
▶ docker run busybox echo "Hello World"
Hello World
```

This is your process entrypoint; the process you want to run. This process 'echo "Hello World"' isolated from your host because it has private versions of those facilities mentioned above.
 