
# source code

Our project will be hosted on [github](https://github.com/cdelorme/csideanime.com) such that development progress is publicly visible, and contributions can be made early on in the project.

_Our objective is to reach a first-release state before we begin accepting pull-requests, but that will depend on the quality of the pull request._

Our planned project hierarchy:

- src/
- pkg/
- bin/
- docs/
    - release-plans/
        - first-release/
    - api/
    - application/
- readme.md
- public/
    - js/
    - css/
- www.csideanime.com
- csideanime-app
- build.sh

A gitignore file will filter out third party files in the gopath, minus our own project source.

The readme will contain a brief description of the project, examples of how to work on the code, examples of how to install it, and links to the more detailed documentation.

Our golang code will live inside `srv/github.com/csideanime.com`, and our static front-end code will live in `public/`.

We will include example nginx configuration files, and a service shell script, and potentially systemd config file, for launching it at boot and restarting it.

**These plans are still under construction.**

