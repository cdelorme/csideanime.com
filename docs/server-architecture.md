
# server architecture

_A portion of the initial configuration of our server, as well as the automated scripts to create it can be found [here](https://github.com/cdelorme/system-setup)._

**Ideally we will not need to install development packages and build tools on our production system, because development will occur on local machines, and deployment will involve simply copying an executable to the production machine.**

Key Packages:

- [nginx](http://wiki.nginx.org/Main)
- [mongodb](https://www.mongodb.org/)

Enabling Scalability:

- we can separate the database to one or more servers and shard the data as needed
- we can supply connection information via `/etc/environment` to any number of services on any server
- both imaging and automation scripts could be used to prepare a new box faster than a fresh rebuild

Expected Scenarios:

- golang api may require a lot of resources for large amounts of traffic, and need to be pooled separately
- a single nginx cannot handle the number of requests and we create a pool of proxy servers
- our database io is too high for the disk, and we need to investigate RAID solutions or sharding/replication
- our database cannot handle the amount of requests, cpu required, or needs more memory for efficient caching

_It is also expected that we will have plenty of time to prepare and respond to this problems, as we do not expect to have a surge of permanent traffic happen unexpectedly._

Configuring nginx:

- have nginx redirect 80 to 443 (prefer-https)
- use self-signed ssl certificate
    - wild card registration pending (not necessary for first-release)
- redirect non-www to www prefix for preferred SEO best-practices
- use nginx caching for delivery of static files, with github hashes as query-strings on css/js to prevent caching problems
- proxy to golang app at `api/` route (or sub-domain)
    - slightly more depth, may involve versioning (eg. `/api/v1/`)

uploads:

- filter media through related ffmpeg/libav/gm tools
- organize files by sha256 hashes to prevent hosting of duplicates
- store files in S3, and make accessible via CloudFront

Multimedia handling is **not** planned as part of first-release.  _A lot more detailed planning is needed around this area still._


