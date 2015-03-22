
# installing csideanime.com

The steps to installing our application will be provided here.  _These will be considered our "deployment process"_


Example for proxying from nginx to golang for non-static api requests:

    upstream go_api_alpha {
        server 127.0.0.1:8080;
    }

    upstream go_api_beta {
        server 127.0.0.1:8081;
    }

    upstream go_api_1 {
        server 127.0.0.1:8082;
    }

    server {

        # serve api traffic /w various versions
        location /api/alpha {
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host $host;
            proxy_pass http://go_api_alpha;
        }
        location /api/beta {
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host $host;
            proxy_pass http://go_api_beta;
        }
        location /api/v1 {
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host $host;
            proxy_pass http://go_api_1;
        }

    }


