
# source code

The project source code will be hosted on [github](https://github.com/cdelorme/csideanime.com).  Development progress can be publicly tracked and contributed to.  We favor visibility.


## file hierarchy

_Tentative layout of files:_

- docs/
    - api/
    - design
    - setup/
- public/
    - js/
    - css/
- logs/.gitignore
- controllers/
- dao/
- models/
- contribute.md
- main.go
- readme.md

The `logs/` folder is for application logging on the server, and is specific to the configuration.  _It should be empty._

The `public/` folder is for the static client-side files, and _is its own application._

The API source begins with `main.go` and dives into `dao/`, `models/`, `controllers/`, and `services/`.  **All core components of the application will exist within this structure, any components that _can_ be swapped will be turned into external libraries.  You can consider these files to be the "Core" of csideanime.com.**

The `docs/` folder will be used to track all documentation for design, setup, and an api reference.

**It is expected that you will handle your deployment of the binary in your own way, and that you will compile on a separate server.**  However, compilation should be as simple as installing golang, then running `go build`.

_We are still missing example configuration files, though they may work best in `docs/setup/`._


## deployment

Our source combines two projects, the API and the static client source, so we have two deployment processes:

- static site deployment
- executable api deployment

For our static site deployment, we can simply setup a cronjob to pull the latest source from master on a regular schedule (say hourly), or on demand manually.  With nginx caching set to 5 minutes changes will trickle through shortly after.  If we use appropriate git-hashes in query-strings we can eliminate caching as well.

For the api it is recommended that you keep your production server light and compile from another location, such as a staging or development system.  Ideally you should be able to simply `scp` the file onto the server, but for first-install you will need to add a layer to launch it at boot time.  Alternative deployment methods include continuous integration and a deploy script that executes with the cronjob that pulls new masters, uploading a binary to S3 and copying it onto one or more machines.  Of course something like ansible could also be used to easily distribute a new binary to many servers.


## feature support

We will want to make sure the following features are viable for use in new ways for the project:

- [localStorage](http://caniuse.com/#feat=namevalue-storage)
- [pushState](http://caniuse.com/#feat=history)
- [requestAnimationFrame](http://caniuse.com/#feat=requestanimationframe)

The above features are available in all modern browsers, including (and especially) mobile devices.

The localStorage API and behavior is (mostly) consistent in almost every browser back to (and including) IE8, and is the best choice for client­side caching. Library wrappers exist, such as `lsCache` to make it even easier. Using this for cache storage would allow us to reduce calls to the API and improve overall performance, especially for multiple tabs.

Modifying the URL and history of the browser is possible in all but IE9­ and Opera Mini. It also has some quirky behavior in Safari. However, a majority support it. Using this would allow us to have a single actual static page and fully dynamic content, but also support all URL’s for link and bookmark support as well as the back button.

The requestAnimationFrame is a nice alternative to registering a callback for anything that modifies the application state. It is also supported in all modern browsers, except IE9­ and Opera Mini. Ideally this can be used to rebuild sections of the page anytime the state has been updated, which could be as simple as comparing timestamps to as complex as comparing virtual DOM.

_In the worst­case scenario our code should gracefully degrade, depriving users of these features, while still running standard ajax calls and operating like a normal site would._
