Allow create/edit/delete of a set of redirects rules containting
  + An implicit redirect from non-www to www domain and vice versa
  + A from domain: e.g. cosmicvent.org
  + A to domain: e.g.   cosmicvent.com
  + A redirect from $from_domain$path to $to_domain$path
  - A list of urls with their from/to paths with an optional status code


# setup

~~~~bash
createdb redirector_production
sudo -u postgres psql < db.sql
go build
./redirector
~~~~


