# drone-helper

A simple way to accomplish two common tasks in [Drone CI](https://www.drone.io/): caching and notifications.

## Notifications

`drone-helper notify` is essentially a convenience wrapper over [shoutrrr](https://github.com/containrrr/shoutrrr).
All it does is gather the relevant build info from Drone, arrange it into a rich text message in a format supported by the respective platform and send it.

### Supported services

#### Discord

```yaml
  - name: notify
    image: zoickx/drone-helper
    commands:
      - drone-helper notify --discord
    settings:
      discord_webhook_token:
        # ex. value: hcXPnPGg_w6ZCPF-4OkLMn7nyTxEbZKex2R2suPNyUTWVl89ij9Qd46DosbREhnykUm4
        from_secret: discord_token 
      discord_webhook_id:
        # ex value: 1125469839924488938
        from_secret: discord_id
```

#### Slack

```yaml
  - name: notify
    image: zoickx/drone-helper
    commands:
      - drone-helper notify --slack
    settings:
      slack_webhook_token:
        # ex. value: T15S51XNJSJ-B05T0LOP280-n8zDEin7EmijY0iCuaISBDay
        from_secret: discord_token
```

## Caching

`drone-helper cache` implements full-system caching via Docker (a practice [endorsed](https://web.archive.org/web/20200617204324/https://discourse.drone.io/t/build-docker-image-and-re-use-in-the-next-step/6190) by the developer).
A cache is uniquely identified by its build dependencies (`--deps`), and will be rebuilt if any one changes.

### Usage

0. `drone-helper cache` must be run before any steps using the cache (even if cache exists and doesn't need rebuilding).
1. The repository must be "Trusted" in Drone (Settings -> General -> Project Settings -> Trusted)
2. The step invoking `drone-helper cache` must mount the host's `docker.sock`.
3. Subsequent steps that need to use the cache must be run in exactly:
``` yaml
    image: cache--${DRONE_REPO}:${DRONE_COMMIT_AFTER}
    pull: never
```


## Example pipeline

``` yaml
---
kind: pipeline
type: docker
name: default

volumes:
  - name: dockersock
    host:
      path: /var/run/docker.sock

steps:
  - name: rebuild-cache
    image: zoickx/drone-helper
    commands:
      - drone-helper cache --deps="Dockerfile dependencies.json"
    volumes:
      - name: dockersock
        path: /var/run/docker.sock

  - name: run-in-cache-1
    image: cache--${DRONE_REPO}:${DRONE_COMMIT_AFTER}
    pull: never
    commands:
      - echo "This command will be run in cache. [1]"

  - name: run-in-cache-2
    image: cache--${DRONE_REPO}:${DRONE_COMMIT_AFTER}
    pull: never
    commands:
      - echo "This command will be run in cache. [2]"

  - name: notify
    image: zoickx/drone-helper
    commands:
      - drone-helper notify --discord
    settings:
      discord_webhook_token:
        from_secret: discord_token
      discord_webhook_id:
        from_secret: discord_id
```
