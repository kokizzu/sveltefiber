Just need to do `rsync -Rv svelte /var/www/html` and create a service (eg. with systemd or docker-compose for the golang service)

if the server used as shared host, just add a reverse proxy, so each backend can handle specific domain
```
/etc/nginx/
  nginx.conf
  sites-available
    default
        server {
            listen 80;
            server_name _;
            root /var/www/html;
            index index.html index.htm index.php;
            location /api/* {
              proxy_pass http://127.0.0.1:3001; # golang/fiber
            }
            try_files $uri $uri/; # static
        }
    whatever.com
      server_name whatever.com
  sites-enabled
    symlink ../whatever.com
```
