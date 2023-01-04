# syntax=docker/dockerfile:1
FROM golang:1.18
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0 
ENV GOOS=linux

RUN cd cmd/agent;go build -o /runner -mod mod -a .

# # FROM alpine
# # RUN apk --no-cache add ca-certificates git npm chromium
# # # RUN npm install -g selenium-webdriver
# # RUN npm install -g mocha
# FROM node:17
# RUN apt-get update && apt-get install -y libgtk2.0-0 \
#     libgtk-3-0 \
#     libgbm-dev \
#     libnotify-dev \
#     libgconf-2-4 \
#     libnss3 \
#     libxss1 \
#     libasound2 \
#     libxtst6 \
#     xauth \
#     xvfb \
# 	git
# RUN apt-get install -y gconf-service libasound2 libatk1.0-0 libatk-bridge2.0-0 libc6 libcairo2 libcups2 libdbus-1-3 libexpat1 libfontconfig1 libgcc1 libgconf-2-4 libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 libnspr4 libpango-1.0-0 libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 libxss1 libxtst6 ca-certificates fonts-liberation libappindicator1 libnss3 lsb-release xdg-utils wget
# RUN apt-get install -y libgbm-dev  
# RUN npm install -g mocha
# RUN npm install -g chromedriver
# # RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
# # RUN apt-get -y install ./google-chrome-stable_current_amd64.deb
# RUN mkdir /home/runner
# WORKDIR /home/runner/
# COPY --from=0 /runner /bin/runner
# # COPY go /usr/local/bin/go
# ENTRYPOINT ["/bin/runner"]


# FROM selenium/standalone-chrome:latest 
# # RUN mkdir /home/runner
# RUN sudo mkdir /home/runner

# RUN sudo apt-get update && sudo apt-get install -y libgtk2.0-0 \
#     libgtk-3-0 \
#     libgbm-dev \
#     libnotify-dev \
#     libgconf-2-4 \
#     libnss3 \
#     libxss1 \
#     libasound2 \
#     libxtst6 \
#     xauth \
#     xvfb \
# 	git
FROM node:18-alpine3.14
# Set up glibc
# ENV LANG en_US.UTF-8
# ENV LANGUAGE en_US:en
# ENV LC_ALL en_US.UTF-8
# ENV GLIBC_REPO=https://github.com/sgerrand/alpine-pkg-glibc
# ENV GLIBC_VERSION=2.34-r0
# RUN set -ex && \
#     apk --update add libstdc++ curl ca-certificates 
    # for pkg in glibc-${GLIBC_VERSION} glibc-bin-${GLIBC_VERSION}; \
    #     do curl -sSL ${GLIBC_REPO}/releases/download/${GLIBC_VERSION}/${pkg}.apk -o /tmp/${pkg}.apk; done && \
    # apk add --allow-untrusted /tmp/*.apk && \
    # rm -v /tmp/*.apk && \
    # /usr/glibc-compat/sbin/ldconfig /lib /usr/glibc-compat/lib

# Install prerequisites and helper packages
# RUN apk add --no-cache \
# 	bash dpkg xeyes

# Download and unpack Chrome
# RUN set -ex && \
# 	curl -SL https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb -o /google-chrome-stable_current_amd64.deb && \
# 	dpkg -x /google-chrome-stable_current_amd64.deb google-chrome-stable && \
# 	mv /google-chrome-stable/usr/bin/* /usr/bin && \
# 	mv /google-chrome-stable/usr/share/* /usr/share && \
# 	mv /google-chrome-stable/etc/* /etc && \
# 	mv /google-chrome-stable/opt/* /opt && \
# 	rm -r /google-chrome-stable

# Install Chrome dependencies
RUN apk add --no-cache --update \
	# alsa-lib \
	# atk \
	# at-spi2-atk \
	# expat \
	# glib \
	# gtk+3.0 \
	# libdrm \
	# libx11 \
	# libxcomposite \
	# libxcursor \
	# libxdamage \
	# libxext \
	# libxi \
	# libxrandr \
	# libxscrnsaver \
	# libxshmfence \
	# libxtst \
	# mesa-gbm \
	# nss \
	# pango \
    npm \
    git 
    
# RUN npm install -g n
# RUN n lts
# RUN npm install -g chromedriver
RUN npm install -g mocha
WORKDIR /home/runner
COPY --from=0 /runner /bin/runner
# ENV DEPENDENCY_MANAGER=npm
ENTRYPOINT ["/bin/runner"]