
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

_In the worst­case scenario our code should gracefully degrade, depriving members of these features, while still running standard ajax calls and operating like a normal site would._


## static content

Ideally we will have a limited number of static pages, which dynamically load all the contents of the site.

Using features such as `pushState` to change the URL dynamically while loading new contents from the API.  Merging assets into single, minified files which can be delivered via nginx or via our content delivery network.

- index.html
- css/main.css
- js/main/js

_Ideally we would want a source map for production errors, but for the start we may just use the code without minification._  For development and testing we want the separate and fully spaced versions of these files, and a utility to merge them.

The basic index.html page may contain a skeletal layout, plus some hard-coded elements like the default menu, login form, static footer, and header.  Most of this content should be able to be rebuilt using javascript once the requests begin.

**Any and all media should be loaded from our CDN and not hosted by Digital Ocean.**

**Javascript Systems:**

- centralized service registration and API callback system
    - tied to lscache
- immediate render system
- login & registration (credentials not profiles)
- member access (profiles & names vs credentials)
- threads service (should encapsulate threads, posts, and tags)

All services must be built with hmac­-sha in mind, to validate data over the wire in both directions.  _This must be designed in a way that does not restrict the number of open tabs or concurrent requests._

Centralized callbacks means incoming data can be validated globally from the service registration component, and then shuffled to the appropriate named/registered service.  This allows us to create modular code that is, by design, used only-once (efficiently).


## user experience

At the moment we do not have a UX engineer. However, when we begin beta development I would be open to having someone fill this position to help us improve the overall design of the site for our members.

Since we are building an API the good news is we can easily swap out one front­end for another.  With a well ­built front-end employing the immediate­-render technique, we can possibly also reuse all the client­-side services that register data-­callbacks and really only need to rebuild the rendering logic.


# api core

**The core of the application will exist above our major code layers, and is responsible for establishing all connections with external resources, creating each of them and supplying them with their dependencies.  It will also start the proxy service layer that connects with nginx.**

The api will load server settings from a "standard path" configuration file:

- `/etc/csideanime.com.json`
- `~/.csideanime.com.json`
- `./csideanime.com.json`

Expect format will be json, and it should contain any `application-level-security-config` options, which are things that are best kept out of "mortal hands" (eg. administrators/moderators).

_We may (optionally) also look at environment variables for this information, and others such as database credentials, but a configuration file would supercede / override environment variables._

Our core should instantiate, configure, and connect (in order):

- data access layers
- service layers
- controllers
- router

By creating all our core components up front we greatly simplify code complexity and improve readability.

Dependencies are created and supplied up front, without any abstraction layers, lazy loading, factors, inversion of control, or service managers.

_This may add to the initial load-time by a small amount, but should take care of the initial warmup that would be required in lazy-loading, and prevents our code from becoming difficult to trace and debug.  Further, we may improve member performance by creating objects sequentially before space fills up and becomes fragmented._

After connecting all of our core, we register all controllers with the router.

Controllers contain the routes, and the router uses wildcard­left routing to reach them, allowing complete support for any number of `clean­-url` parameters and additional constructs without complicating the server or proxy configuration.

Finally, we supply the router to a new server, and have the server begin listening for incoming traffic.


### settings

@TODO

I would like to identify all application-level settings and document them here.


## code layers

Our api will be made up of four major layers of components, giving us a solid pattern to work with and good separation of concerns:

<table>
    <tr>
        <th>Layer</th>
        <th>Purpose</th>
    </tr>
    <tr>
        <td>Controller</td>
        <td>Handles routed connections, parses supplied arguments and handles controller-level-conditions and first-pass-acl.  Generally depends on two or more services, and responsible for replying with final results to the request object.</td>
    </tr>
    <tr>
        <td>Service</td>
        <td>Handles business logic and second-pass-acl.  Depends on the data layer and may depend directly on some models.</td>
    </tr>
    <tr>
        <td>Data Access</td>
        <td>Depends on models, connets to database, manipulates data, and handles third-pass-acl.</td>
    </tr>
    <tr>
        <td>Model</td>
        <td>Definition of the data (eg. structs), should have no dependencies.</td>
    </tr>
</table>

All valid requests will be delivered to a controller, which will determine the required action to perform.

Each action will depend on one or more services.  _Services should not depend on other services directly, the controller should coordinate operations between them._  Services may, independently, require specific models.

The data access layer interfaces with our database, and depends on models.

Models are stand-alone representations of data, which exist to house and manipulate data locally.  _No model-functions should deal with external resources._

Keeping out code as simple as possible, and following these rules, will greatly reduce the potential for security exploits and bugs.


## design concepts

The following concepts need to be tested, but are considerations for our applications construction:

- client side caching (lscache?)
- immediate-rendering (mithriljs?)
- pushState history/url augmentation
- utc-only time storage and client-side conversion using client-os
- open authentication system /w google & facebook connection
- centralized callback handler, _requires a well-designed standard message format._
- nginx proxy cache invalidations, and alternative caching layers that may be more flexible?
