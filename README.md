# infra-cli

## Overview
Infra CLI is a command-line tool and API system that automates deploying infrastructure configurations to multiple Linux servers using Ansible.
It provides job tracking, rollback capabilities, and an audit trail for all deployments.

## The Problem
Infrastructure teams often manually SSH into many servers and run commands individually.

Example:

- Deploy nginx to server-1  
  `ssh server-1 "apt install nginx"`

- Deploy nginx to server-2  
  `ssh server-2 "apt install nginx"`

- Repeat this process across dozens of servers.

This approach is:
- slow
- error-prone
- difficult to audit
- difficult to roll back

## The Solution

Infra CLI automates the entire deployment process.

```
infra deploy webserver-stack --env production
```

This command:

- Deploys configuration to all target servers automatically
- Executes deployments consistently across environments
- Tracks deployment status
- Maintains an audit trail
- Allows safe rollbacks
