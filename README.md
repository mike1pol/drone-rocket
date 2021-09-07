# drone-[rocket](https://rocket.chat) [![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/mike1pol/drone-rocket/latest)](https://hub.docker.com/r/mike1pol/drone-rocket/) [![Docker Image Size (tag)](https://img.shields.io/docker/image-size/mike1pol/drone-rocket/latest)](https://hub.docker.com/r/mike1pol/drone-rocket/tags)

> [Drone](https://www.drone.io/) plugin for sending [Rocket.Chat](https://rocket.chat) messages to a channel using the latest [REST API v1](https://developer.rocket.chat/reference/api/rest-api).<br>

Sending [Rocket.Chat](https://rocket.chat) message using a binary, Docker or [Drone CI](http://docs.drone.io/).

## Usage

There are three ways to send notifications.

- [drone-rocket](#drone-rocket)
  - [Usage](#usage)
    - [Usage from binary](#usage-from-binary)
      - [Send Notification](#send-notification)
    - [Usage from Docker](#usage-from-docker)
      - [Send Notification](#send-notification-1)
    - [Usage from Drone CI](#usage-from-drone-ci)
      - [Send Notification](#send-notification-2)
  - [Testing](#testing)

<a name="usage-from-binary"></a>
### Usage from binary

#### Send Notification

```bash
drone-rocket \
  --url xxxx \
  --user-id xxxx \
  --token xxx \
  --channel xxx \
  --message "Test Message"
```

<a name="usage-from-docker"></a>
### Usage from Docker

#### Send Notification

```bash
docker run --rm \
  -e URL=xxxxxxx \
  -e USER_ID=xxxxxxx \
  -e TOKEN=xxxxxxx \
  -e CHANNEL='#general' \
  -e AVATAR_URL=http://example.com/xxxx.png \
  -e MESSAGE=test \
  mike1pol/drone-rocket
```

<a name="usage-from-drone-ci"></a>
### Usage from Drone CI

#### Send Notification

Execute from the working directory:

```bash
docker run --rm \
  -e URL=xxxxxxx \
  -e USER_ID=xxxxxxx \
  -e TOKEN=xxxxxxx \
  -e CHANNEL='#general' \
  -e AVATAR_URL='https://upload.wikimedia.org/wikipedia/commons/6/69/June_odd-eyed-cat_cropped.jpg' \
  -e MESSAGE=test \
  -e DRONE_REPO_OWNER=mike1pol \
  -e DRONE_REPO_NAME=drone-rocket \
  -e DRONE_COMMIT_SHA=e5e82b5eb3737205c25955dcc3dcacc839b7be52 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=mike1pol \
  -e DRONE_COMMIT_AUTHOR_EMAIL=mikle.sol@gmail.com \
  -e DRONE_COMMIT_MESSAGE=Test_Your_Commit \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/mike1pol/drone-rocket \
  -e DRONE_JOB_STARTED=1477550550 \
  -e DRONE_JOB_FINISHED=1477550750 \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  mike1pol/drone-rocket
```

## Testing

Test the package with the following command:

```
$ make test
```
