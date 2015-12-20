[![Build Status](https://travis-ci.org/jacobhands/multiweb.svg?branch=master)](https://travis-ci.org/jacobhands/multiweb)

# multiweb
Serve multiple websites easily using Go.

Base URL: `sites.example.com`

Folder structure:
```
sites // The root folder
  - foo.com
    - www                   // foo.com website files go here.
    - ssl                   // SSL certs for foo.com go here
        ssl.bundle.crt      // SSL bundle for site
        ssl.key             // SSL key
    - a.foo.com             // a.foo.com
      - www               // a.foo.com site files.
      - ssl               // SSL certs for a.foo.com  
          ssl.bundle.crt
          ssl.key
    - b.foo.com
      - www
  - bar.com
    - www
    - c.bar.com
      - www
      - d.bar.com
        - www
        - ssl
            ssl.bundle.crt
            ssl.key
```
foo.com DNS CNAME: `foo.com.sites.example.com`

bar.com DNS CNAME: `bar.com.sites.example.com`

c.bar.com DNS CNAME: `c.bar.com.sites.example.com`