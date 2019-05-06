<img src="https://rocket.chat/images/default/logo--dark.svg">

# drone-[rocket](https://rocket.chat)

Drone plugin for sending message to Rocket.Chant channel using API.

[![Build Status](https://ci.piterjs.org/api/badges/mike1pol/drone-rocket/status.svg)](https://ci.piterjs.org/mike1pol/drone-rocket)
[![Docker Pulls](https://img.shields.io/docker/pulls/mike1pol/drone-rocket.svg)](https://hub.docker.com/r/mike1pol/drone-rocket/)
[![](https://images.microbadger.com/badges/image/mike1pol/drone-rocket.svg)](https://microbadger.com/images/mike1pol/drone-rocket "Get your own image badge on microbadger.com")


Sending Rocket.Chat message using a binary, docker or [Drone CI](http://docs.drone.io/).


## Usage

There are three ways to send notification.

* [usage from binary](#usage-from-binary)
* [usage from docker](#usage-from-docker)
* [usage from drone ci](#usage-from-drone-ci)

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
### Usage from docker

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
### Usage from drone ci

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
