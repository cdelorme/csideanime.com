
# infrastructure

This document is a breakdown of the hosting, technologies, and architecture we plan to use for our service.


## hosting

- **Server:** [DigitalOcean](https://www.digitalocean.com)
- **Storage/CDN:** AWS S3 + AWS CloudFront
- **Repository:** [github](https://github.com/cdelorme/csideanime.com)
- **Operating System:** [debian/linux](http://www.debian.org/)
- **Proxy:** [nginx](http://wiki.nginx.org/Main)

We will use S3 to upload and store static content and multimedia, then deliver it using AWS CloudFront as the CDN.

We will host the source openly on github, which will make deployment a breeze.

The decision to use debian is for stability and familiarity, but it should have little to no impact on whether the code can be deployed to another distribution.  _However, the instructions will be tailored to the distro._

We will be using nginx for static content, caching, and as a proxy to backend services.

_The initial configuration will be a single server, but the architecture will be designed with scalable execution in mind for the future._


### notes

Some tricky business with AWS CloudFront and S3:

- the cli will be easier to manage large numbers of files at once, such as to set public accessibility
- CloudFront and S3 will need to be configured to work with HTTPS traffic and CORS compatibility
    - it would be ideal if we could use the sites ssl on aws instead of cloudfronts ssl


## technologies

Our backend code will primarily consist of [golang](http://golang.org/).

We plan to use [mongodb](https://www.mongodb.org/) as the database for persistent storage.

We may, at a later time, explore [elastic search](http://www.elasticsearch.org/) as for search indexing and potentially caching.

We plan to _try_ the following javascript libraries mixed with custom code for the frontend:

- [lscache](https://github.com/pamelafox/lscache)
- [mithriljs](http://lhorie.github.io/mithril/)

We will likely be using [bootstrap](http://getbootstrap.com/) for its grid and then custom styles for the rest (further integration with bootstrap will depend on plans for theming in future releases).


## architecture

Our plan is to build two entirely independent resources:

- an api
- a static frontend

The static pages will contain javscript which requests information through the API.

We will use nginx to serve the static files, and to forward requests to the api.

This separation also means we can create a 100% mocked backend to test our frontend, and that breaks in one area do not necessarily also cause the other area to break in development.

We also plan to test client side caching using localstorage, and real-time rendering using requestAnimationFrame (as opposed to the traditional DOM manipulation model).

_We may explore the use of websockets, but browser support is still very limited._


## uploads

We have not made plans yet to handle file uploads, but if we did here are some best-practices we would follow:

- filter media through related ffmpeg/libav/gm tools
- organize files by sha256 hashes to prevent hosting of duplicates
- store files in S3, and make accessible via CloudFront

_This will likely be a part of the beta, if not a later feature._


### scalability

While we anticipate having plenty of time before we need to look at scalability, we have a good action plan assembled to embrace growth.

The first step to take is to make sure all large files, binaries, images, etc are hosted on S3 and served via CloudFront or a similar Caching DNS service.  This will greatly improve loading speeds while reducing the load on your server.

It often goes unsaid, but in order to write a good scalable server software, it must be written in a way that is stateless.

**Expected Scenarios & Responses:**

- golang api requires a lot of resources
    - move the api to its own server
    - consider creating a pool of API servers with nginx as a proxy
- a single nginx cannot handle the number of requests
    - look into regional DNS routing to different boxes (AWS is good at this)
- database cannot handle the number of requests or needs more cpu or memory
    - consider sharding
- our database io is too high for the disk
    - consider RAID
    - consider sharding or intelligent replication

**Notes:**

Single Server, RAID, and Sharding all suffer in design from single-point-of-failure.  If you are concerned with data reliability then you should use replication.  Sadly, because `mongodb` uses the master-slave architecture replication does not solve IO problems.  So you'll need to combine replication and RAID or Sharding to get the best results.

It is good practice to write applications to look to the environment (`/etc/environment`) for credentials, because those can be entered once and shared among all running systems.  However, this becomes less useful as you separate services to individual machines.  _It is no more of a security risk than hard-coding your usernames and passwords into configuration files, and far more flexible with regard to deploying new code._

_While auto-scaling can be acheived in digital-ocean, I believe that by the time we reach the need for it we will move to AWS._
