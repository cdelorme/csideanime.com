
# configuration

This document assumes you already have a web server in place, but if you are just setting one up we [can recommend some configuration instructions](https://github.com/cdelorme/system-setup).

**Ideally our production server will not contain development packages, since the site consists of an executable binary and static html and javascript.**

Key packages to install include:

- `nginx`
- `mongodb`


## server content

Ideally you should place server files in `/srv/`, my suggested pathing is:

- `/srv/www/csideanime.com/`

Simply clone the repository to that path, and then copy and modify the [supplied nginx configuration file](#).


## nginx

Ideally you should configure nginx to handle both 80 & 443 traffic.

You should at the very least use https for logging in, but it would be ideal to use it everywhere given that you can acquire and ssl certificate for free from [startssl](https://www.startssl.com/), and setting them up takes barely 5 minutes.

Best practices for SEO results indicate that you should redirect non-www prefix to `www.` for your domain.

It would be best that you configure nginx to cache the static files it delivers.

For best performance the latest javascript and css should be minified and uploaded ot S3 and delivered through CloudFront, also named by the git short-hash of the latest commit.  A decent alternative to prevent nginx caching problems would be to append a querystring with the short-hash.
