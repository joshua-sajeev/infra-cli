FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update && \
    apt install -y --no-install-recommends \
        openssh-server \
        vim \
        git \
        curl \
    && rm -rf /var/lib/apt/lists/*

# create ssh runtime directory
RUN mkdir /var/run/sshd

# set root password
RUN echo "root:infra" | chpasswd

# enable root login and password authentication
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config && \
    sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config

# prepare ssh folder
RUN mkdir -p /root/.ssh && chmod 700 /root/.ssh

EXPOSE 22

CMD ["/usr/sbin/sshd","-D"]
