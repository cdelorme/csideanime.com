
# feature support

- [localStorage](http://caniuse.com/#feat=namevalue-storage)
- [pushState](http://caniuse.com/#feat=history)
- [requestAnimationFrame](http://caniuse.com/#feat=requestanimationframe)

The above features are available in all modern browsers, including (and especially) mobile devices.

The localStorage API and behavior is (mostly) consistent in almost every browser back to (and including) IE8, and is the best choice for client­side caching. Library wrappers exist, such as `lsCache` to make it even easier. Using this for cache storage would allow us to reduce calls to the API and improve overall performance, especially for multiple tabs.

Modifying the URL and history of the browser is possible in all but IE9­ and Opera Mini. It also has some quirky behavior in Safari. However, a majority support it. Using this would allow us to have a single actual static page and fully dynamic content, but also support all URL’s for link and bookmark support as well as the back button.

The requestAnimationFrame is a nice alternative to registering a callback for anything that modifies the application state. It is also supported in all modern browsers, except IE9­ and Opera Mini. Ideally this can be used to rebuild sections of the page anytime the state has been updated, which could be as simple as comparing timestamps to as complex as comparing virtual DOM.

_In the worst­case scenario our code should gracefully degrade, depriving users of these features, while still running standard ajax calls and operating like a normal site would._

