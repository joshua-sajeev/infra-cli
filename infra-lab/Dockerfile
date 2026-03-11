FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        openssh-server \
        vim \
        git \
        curl \
    && rm -rf /var/lib/apt/lists/*

# create ssh runtime directory
RUN mkdir /var/run/sshd

# create dev user
RUN useradd -m -s /bin/bash dev

RUN echo "root:infra-cli" | chpasswd && \
    echo "dev:infra" | chpasswd

# ssh configuration
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config && \
    sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config && \
    echo "PubkeyAuthentication yes" >> /etc/ssh/sshd_config

# prepare ssh folder
RUN mkdir -p /home/dev/.ssh && \
    chown -R dev:dev /home/dev/.ssh && \
    chmod 700 /home/dev/.ssh

EXPOSE 22

CMD ["/usr/sbin/sshd","-D"]
