dev:
  hugo server

build:
  hugo -t hugo-notepadium

deploy_oss: build
  ossutil cp -r docs oss://rsocketbyexampleus --update

deploy_server: build
  scp -r docs/* root@rsocketbyexample.com:/home/virtual_hosts/rsocketbyexample.com

