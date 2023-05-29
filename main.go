#cloud-config
users:
  - name: root
    passwd: your_root_password
    lock-passwd: false
    ssh-authorized-keys:
      - ssh-rsa your_public_ssh_key
    sudo: ALL=(ALL) NOPASSWD:ALL
    shell: /bin/bash

ssh_pwauth: true
