FROM buildpack-deps:bullseye-scm

# the setup procedure according to https://repo.powerdns.com/ (Debian 9 Stretch)
RUN echo "deb [signed-by=/etc/apt/keyrings/auth-48-pub.asc arch=amd64] http://repo.powerdns.com/debian bullseye-auth-48 main" > /etc/apt/sources.list.d/pdns.list \
	&& echo "Package: pdns-*\nPin: origin repo.powerdns.com\nPin-Priority: 600\n" >> /etc/apt/preferences.d/pdns \
	&& install -d /etc/apt/keyrings; curl https://repo.powerdns.com/FD380FBB-pub.asc | tee /etc/apt/keyrings/auth-48-pub.asc \
	&& apt-get -y update \
	&& apt-get install -y pdns-server pdns-backend-mysql  \
	&& rm -rf /var/lib/apt/lists/*

COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 5353

CMD ["pdns_server", "--daemon=no"]
