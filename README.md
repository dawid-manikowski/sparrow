# Sparrow

Sparrow is a humble tool for basic website monitoring.\
It does two things and it does them as simple as possible:
- sparrow monitors
- sparrow notifies

## Monitoring
Sparrow does basic http get request to a specified website and checks if the response had chosen status code.

## Notification
Simplest form of notification - email. Using the specified credentials sparrow logs in to smtp server and sends a simple message whenever the status code is different than the one specified in config.

Important thing: NOTIFICATION IS SENT ONLY IF RETURNED STATUS CODE IS DIFFERENT THAN EXPECTED ONE

## Configuration
Sparrow is entirely controlled by environment variables.
Template .sparrowrc file is located in this repo under `.sparrowrc.example`.

Content is located below:
```
export SPARROW_MAIL_HOSTNAME=
export SPARROW_MAIL_USERNAME=
export SPARROW_MAIL_PORT=
export SPARROW_MAIL_PASSWORD=
export SPARROW_MAIL_DESTINATION=
export SPARROW_WEBSITE=
export SPARROW_STATUS_CODE=
```

- `SPARROW_MAIL_HOSTNAME` - smtp server hostname (ex. smtp.gmail.com)
- `SPARROW_MAIL_USERNAME` - smtp server username (ex. me@gmail.com)
- `SPARROW_MAIL_PORT`- smtp server port (ex. 465)
- `SPARROW_MAIL_PASSWORD` - smtp server password (base64 encoded)
- `SPARROW_MAIL_DESTINATION` - destination email address to send notification to (ex. bla@gmail.com)
- `SPARROW_WEBSITE` - webiste to check status on (ex. https://google.com/)
- `SPARROW_STATUS_CODE` - expected status code (ex. 200)

_Author: dd.manikowski@gmail.com_
