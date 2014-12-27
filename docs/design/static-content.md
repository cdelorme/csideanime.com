
# static content

I am expecting we will have a single static index page, which dynamically loads all the meat of the site.

We want a minimalistic set of statically delivered files:

- index.html
- css/main.css
- js/main/js

For development and testing we may want to have multiple javascript and css files. For production we should use a utility to merge and minify all of our css and javascript.

The basic index.html page may contain a skeletal layout, plus maybe some hard­coded elements like the default menu, login form, static footer, and header with dynamic image rotation.

**Any and all media should be loaded from our CDN and not hosted by Digital Ocean.**


**Javascript Systems:**

- centralized service registration and API callback system
    - tied to lscache
- immediate render system
- login & registration (credentials not profiles)
- member access (profiles & names vs credentials)
- threads service (should encapsulate threads, posts, and tags)

All services must be built with hmac­sha in mind, to validate data over the wire in both directions.

Centralized callbacks means incoming data can be validated globally from the service registration component, and then shuffled to the appropriate named/registered service.

