
# hosting

Our plans are as follows:

- **Server:** [DigitalOcean](https://www.digitalocean.com)
- **Storage:** AWS S3 will be used to store static content and multimedia
- **CDN:** AWS CloudFront will be used to server static content efficiently

_The initial configuration will be a single server, but the architecture will be designed with scalable execution in mind for the future._


## notes

Some tricky business with AWS CloudFront and S3:

- the cli will be easier to manage large numbers of files at once, such as to set public accessibility
- CloudFront and S3 will need to be configured to work with HTTPS traffic and CORS compatibility
    - it would be ideal if we could use the sites ssl on aws instead of cloudfronts ssl

