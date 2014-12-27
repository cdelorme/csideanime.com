
# technologies

This broadly describes any of the software involved with the construction of the system, including the operating system, any delivery softare, and of course languages or techniques used to construct the system.

- **Operating System:** [debian/linux](http://www.debian.org/)
- **Proxy:** [nginx](http://wiki.nginx.org/Main)
- **UI:**
    - Static HTML
    - Javascript:
        - custom
        - [lscache](https://github.com/pamelafox/lscache)
        - [mithriljs](http://lhorie.github.io/mithril/)
- **API:** REST
- **API Language:** [golang](http://golang.org/)
- **Repository:** [github](https://github.com/cdelorme/csideanime.com)
- **CSS Library:**
    - [bootstrap](http://getbootstrap.com/)
    - custom

The goal is to build an entirely independent api, and static client side for the UI which uses javascript and the API.  _This eliminates a lot of area where development bugs come into play, but does also restrict the systems accessibility._ The API will be kept clean and simple, and follow standard (or sane) REST practices. The client side will use immediate-mode rendering, combined with localStorage caching and preferably a polling mechanism.  _The option for websockets is still under consideration, but websocket support is still limited._  A centralized callback system for all API responses will be used, which will make up the "custom" portion of the javascript source.  We will use bootstrap primarily for its grid system, and possibly for layout design such that we can possibly provide theming _in the future_.  Timezones will be stored in UTC, and javascript will be used to translate those according to the users locale (eliminating all timezone concerns).

