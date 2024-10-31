This repository is source code for [Have a good the DNS in the Cloud / openSUSE.Asia Summit 2024](https://events.opensuse.org/conferences/oSAS24/program/proposals/4872).

This repository list up DNS Service implementation patterns.

1. Traditional zone trasnfer
    - dns_zone_transfer/
2. Flexible DNS Server(PowerDNS)
    - powerdns_with_lua/
3. Flexible DNS Server(PowerDNS) + Extension(MySQL Backend)
    - powerdns_with_mysql/
4. Flexible DNS Server(PowerDNS) + Extension(Remote Backend) + Backend Implementation(Golang Program)
    - powerdns_with_rest_api/
5. Implement DNS Zone Manager(Golang Program) + Simple DNS Server
    - dns_zone_transfer_implement/
6. Implement DNS Server(Golang program)
    - dns_server_implement/


1, 3, 5 are Key-value store DNS service pattern.

2, 4, 6 are Service Discovery(GSLB) DNS service pattern.