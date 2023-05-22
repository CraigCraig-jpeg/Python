"cd_content": [
  "/path/to/file1",
  "/path/to/file2",
  {
    "path": "/data/meta-data",
    "content": "instance-id: iid-local01\nlocal-hostname: ubuntu-server"
  },
  {
    "path": "/data/user-data",
    "content": "#cloud-config\n\n# THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE\n# WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR\n# COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR\n# OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.\n\n# Ubuntu Server 20.04 LTS\n\nautoinstall:\n  version: 1\n  apt:\n    geoip: true\n    preserve_sources_list: false\n    primary:\n      - arches: [amd64, i386]\n        uri: http://archive.ubuntu.com/ubuntu\n      - arches: [default]\n        uri: http://ports.ubuntu.com/ubuntu-ports\n  early-commands:\n    - sudo systemctl stop ssh\n  locale: ${vm_guest_os_language}\n  keyboard:\n    layout: ${vm_guest_os_keyboard}\n  storage:\n    config:\n      - ptable: gpt\n        path: /dev/sda\n        wipe: superblock\n        type: disk\n        id: disk-sda\n      - device: disk-sda\n        size: 1024M\n        wipe: superblock\n        flag: boot\n        number: 1\n        grub_device: true\n        type: partition\n        id: partition-0\n      - fstype: fat32\n        volume: partition-0\n        label: EFIFS\n        type: format\n        id: format-efi\n      - device: disk-sda\n        size: 1024M\n        wipe: superblock\n        number: 2\n        type: partition\n        id: partition-1\n      - fstype: xfs\n        volume: partition-1\n        label: BOOTFS\n        type: format\n        id: format-boot\n      - device: disk-sda\n        size: -1\n        wipe: superblock\n        number: 3\n        type: partition\n        id: partition-2\n      - name: sysvg\n        devices:\n          - partition-2\n        type: lvm_volgroup\n        id: lvm_volgroup-0\n      - name: root\n        volgroup: lvm_volgroup-0\n        size: 12288M\n        wipe: superblock\n        type: lvm_partition\n        id: lvm_partition-root\n      - fstype: xfs\n        volume: lvm_partition-root\n        type: format\n        label: ROOTFS\n        id: format-root\n      - name: home\n        volgroup: lvm_volgroup-0\n        size: 4096M\n        wipe: superblock\n        type: lvm_partition\n        id: lvm_partition-home\n      - fstype: xfs\n        volume: lvm_partition-home\n        type: format\n        label: HOMEFS\n        id: format-home\n      - name: opt\n        vol

    
    "cd_files": [
  {
    "path": "/preseed.cfg",
    "content": "#cloud-config\n\nautoinstall:\n  version: 1\n  apt:\n    geoip: true\n    preserve_sources_list: false\n    primary:\n      - arches: [amd64, i386]\n        uri: http://archive.ubuntu.com/ubuntu\n      - arches: [default]\n        uri: http://ports.ubuntu.com/ubuntu-ports\n  early-commands:\n    - sudo systemctl stop ssh\n  locale: ${vm_guest_os_language}\n  keyboard:\n    layout: ${vm_guest_os_keyboard}\n  identity:\n    hostname: ubuntu-server\n    username: ${build_username}\n    password: ${build_password_encrypted}\n  ssh:\n    install-server: true\n    allow-pw: true\n  packages:\n    - openssh-server\n    - open-vm-tools\n    - cloud-init\n  user-data:\n    disable_root: false\n    timezone: ${vm_guest_os_timezone}\n  late-commands:\n    - sed -i -e 's/^#\\?PasswordAuthentication.*/PasswordAuthentication yes/g' /target/etc/ssh/sshd_config\n    - echo '${build_username} ALL=(ALL) NOPASSWD:ALL' > /target/etc/sudoers.d/${build_username}\n    - curtin in-target --target=/target -- chmod 440 /etc/sudoers.d/${build_username}"
  }
]

    {
  "builders": [
    {
      "type": "vsphere-iso",
      "vcenter_server": "vcenter.example.com",
      "username": "username",
      "password": "password",
      "insecure_connection": true,
      "vm_name": "ubuntu-server",
      "datacenter": "datacenter-name",
      "cluster": "cluster-name",
      "network": "network-name",
      "datastore": "datastore-name",
      "convert_to_template": true,
      "disk_size": "20000",
      "communicator": "ssh",
      "ssh_username": "ubuntu",
      "ssh_password": "ubuntu",
      "ssh_wait_timeout": "30m",
      "ssh_port": 22,
      "shutdown_command": "sudo systemctl poweroff",
      "guest_os_type": "ubuntu-64",
      "iso_paths": [
        "[datastore-name] path/to/ubuntu-20.04.3-live-server-amd64.iso"
      ],
      "cd_files": [
        {
          "content": "#cloud-config\n\nautoinstall:\n  version: 1\n  apt:\n    geoip: true\n    preserve_sources_list: false\n    primary:\n      - arches: [amd64, i386]\n        uri: http://archive.ubuntu.com/ubuntu\n      - arches: [default]\n        uri: http://ports.ubuntu.com/ubuntu-ports\n  early-commands:\n    - sudo systemctl stop ssh\n  locale: ${vm_guest_os_language}\n  keyboard:\n    layout: ${vm_guest_os_keyboard}\n  identity:\n    hostname: ubuntu-server\n    username: ${build_username}\n    password: ${build_password_encrypted}\n  ssh:\n    install-server: true\n    allow-pw: true\n  packages:\n    - openssh-server\n    - open-vm-tools\n    - cloud-init\n  user-data:\n    disable_root: false\n    timezone: ${vm_guest_os_timezone}\n  late-commands:\n    - sed -i -e 's/^#\\?PasswordAuthentication.*/PasswordAuthentication yes/g' /target/etc/ssh/sshd_config\n    - echo '${build_username} ALL=(ALL) NOPASSWD:ALL' > /target/etc/sudoers.d/${build_username}\n    - curtin in-target --target=/target -- chmod 440 /etc/sudoers.d/${build_username}",
          "destination": "/root/user-data"
        }
      ]
    }
  ],
  "provisioners": [
    {
      "type": "shell",
      "execute_command": "echo 'password' | sudo -S sh '{{ .Path }}'",
      "scripts": [
        "scripts/01-base.sh",
        "scripts/02-extras.sh"
      ]
    }
  ]
}
    
    
    
------------
    
{
  "builders": [
    {
      "type": "vsphere-iso",
      "vcenter_server": "vcenter.example.com",
      "username": "username",
      "password": "password",
      "insecure_connection": true,
      "vm_name": "ubuntu-server",
      "datacenter": "datacenter-name",
      "cluster": "cluster-name",
      "network": "network-name",
      "datastore": "datastore-name",
      "convert_to_template": true,
      "disk_size": "20000",
      "communicator": "ssh",
      "ssh_username": "ubuntu",
      "ssh_password": "ubuntu",
      "ssh_wait_timeout": "30m",
      "ssh_port": 22,
      "shutdown_command": "sudo systemctl poweroff",
      "guest_os_type": "ubuntu-64",
      "iso_paths": [
        "[datastore-name] path/to/ubuntu-20.04.3-live-server-amd64.iso"
      ],
      "cd_files": [
        {
          "content": "#cloud-config\n\nautoinstall:\n  version: 1\n  apt:\n    geoip: true\n    preserve_sources_list: false\n    primary:\n      - arches: [amd64, i386]\n        uri: http://archive.ubuntu.com/ubuntu\n      - arches: [default]\n        uri: http://ports.ubuntu.com/ubuntu-ports\n  early-commands:\n    - sudo systemctl stop ssh\n  locale: ${vm_guest_os_language}\n  keyboard:\n    layout: ${vm_guest_os_keyboard}\n  identity:\n    hostname: ubuntu-server\n    username: ${build_username}\n    password: ${build_password_encrypted}\n  ssh:\n    install-server: true\n    allow-pw: true\n  packages:\n    - openssh-server\n    - open-vm-tools\n    - cloud-init\n  user-data:\n    disable_root: false\n    timezone: ${vm_guest_os_timezone}\n  late-commands:\n    - sed -i -e 's/^#\\?PasswordAuthentication.*/PasswordAuthentication yes/g' /target/etc/ssh/sshd_config\n    - echo '${build_username} ALL=(ALL) NOPASSWD:ALL' > /target/etc/sudoers.d/${build_username}\n    - curtin in-target --target=/target -- chmod 440 /etc/sudoers.d/${build_username}",
          "destination": "/root/user-data"
        }
      ],
      "boot_command": [
        "<esc><wait>",
        "<esc><wait>",
        "<enter><wait>",
        "/install/vmlinuz<wait>",
        " auto<wait>",
        " console-setup/ask_detect=false<wait>",
        " console-setup/layoutcode=us<wait>",
        " console-setup/modelcode=pc105<wait>",
        " debconf/frontend=noninteractive<wait>",
        " debian-installer=en_US<wait>",
        " fb=false<wait>",
        " initrd=/install/initrd.gz<wait>",
        " kbd-chooser/method=us<wait>",
        " keyboard-configuration/layout=USA<wait>",
        " keyboard-configuration/variant=USA<wait>",
        " locale=en_US<wait>",
        " netcfg/get_domain=vm<wait>",
        " netcfg/get_hostname=ubuntu<wait>",
        " grub-installer/bootdev=/dev/sda<wait>",
        " noapic<wait>",
        " preseed/url=http://{{ .HTTPIP }}:{{ .HTTPPort }}/preseed.cfg<wait>",
        " preseed/url/checksum={{user `preseed_checksum`}}<wait>",
        " preseed/url/tls={{user `preseed_tls`}}<wait>",
        " preseed/url/cert={{user `preseed_cert`}}<wait>",
        " -- <wait>",
        "<enter>"
      ]
    }
  ],
  "provisioners": [
    {
      "type": "shell",
      "execute_command": "echo 'password' | sudo -S sh '{{ .Path }}'",
      "scripts": [
        "scripts/01-base.sh",
        "scripts/02-extras.sh"
      ]
    }
  ]
}
0-----------
    
    "boot_command": [
        "<esc><wait>",
        "<esc><wait>",
        "<enter><wait>",
        "/install/vmlinuz<wait>",
        " auto<wait>",
        " console-setup/ask_detect=false<wait>",
        " console-setup/layoutcode=us<wait>",
        " console-setup/modelcode=pc105<wait>",
        " debconf/frontend=noninteractive<wait>",
        " debian-installer=en_US<wait>",
        " fb=false<wait>",
        " initrd=/install/initrd.gz<wait>",
        " kbd-chooser/method=us<wait>",
        " keyboard-configuration/layout=USA<wait>",
        " keyboard-configuration/variant=USA<wait>",
        " locale=en_US<wait>",
        " netcfg/get_domain=vm<wait>",
        " netcfg/get_hostname=ubuntu<wait>",
        " grub-installer/bootdev=/dev/sda<wait>",
        " noapic<wait>",
        " preseed/file=/preseed.cfg<wait>",
        " -- <wait>",
        "<enter>"
      ],
      "floppy_files": [
        "data/user-data.pkrtpl.hcl",
        "data/network-config.pkrtpl.hcl"
      ]
    }
  ],



---------------


{
  "builders": [
    {
      "type": "vsphere-iso",
      "vcenter_server": "vcenter.example.com",
      "username": "username",
      "password": "password",
      "insecure_connection": true,
      "vm_name": "ubuntu-server",
      "datacenter": "datacenter-name",
      "cluster": "cluster-name",
      "network": "network-name",
      "datastore": "datastore-name",
      "convert_to_template": true,
      "disk_size": "20000",
      "communicator": "ssh",
      "ssh_username": "ubuntu",
      "ssh_password": "ubuntu",
      "ssh_wait_timeout": "30m",
      "ssh_port": 22,
      "shutdown_command": "sudo systemctl poweroff",
      "guest_os_type": "ubuntu-64",
      "iso_paths": [
        "[datastore-name] path/to/ubuntu-20.04.3-live-server-amd64.iso"
      ],
      "boot_command": [
        "<esc><wait>",
        "<esc><wait>",
        "<enter><wait>",
        "/install/vmlinuz<wait>",
        " auto<wait>",
        " console-setup/ask_detect=false<wait>",
        " console-setup/layoutcode=us<wait>",
        " console-setup/modelcode=pc105<wait>",
        " debconf/frontend=noninteractive<wait>",
        " debian-installer=en_US<wait>",
        " fb=false<wait>",
        " initrd=/install/initrd.gz<wait>",
        " kbd-chooser/method=us<wait>",
        " keyboard-configuration/layout=USA<wait>",
        " keyboard-configuration/variant=USA<wait>",
        " locale=en_US<wait>",
        " netcfg/get_domain=vm<wait>",
        " netcfg/get_hostname=ubuntu<wait>",
        " grub-installer/bootdev=/dev/sda<wait>",
        " noapic<wait>",
        " preseed/file=/preseed.cfg<wait>",
        " -- <wait>",
        "<enter>"
      ],
      "cd_files": [
        {
          "content": "#cloud-config\n\nautoinstall:\n  version: 1\n  apt:\n    geoip: true\n    preserve_sources_list: false\n    primary:\n      - arches: [amd64, i386]\n        uri: http://archive.ubuntu.com/ubuntu\n      - arches: [default]\n        uri: http://ports.ubuntu.com/ubuntu-ports\n  early-commands:\n    - sudo systemctl stop ssh\n  locale: en_US\n  keyboard:\n    layout: USA\n  identity:\n    hostname: ubuntu-server\n    username: ubuntu\n    password: ubuntu\n  ssh:\n    install-server: true\n    allow-pw: true\n  packages:\n    - openssh-server\n    - open-vm-tools\n    - cloud-init\n  user-data:\n    disable_root: false\n    timezone: UTC\n  late-commands:\n    - sed -i -e 's/^#\\?PasswordAuthentication.*/PasswordAuthentication yes/g' /target/etc/ssh/sshd_config\n    - echo 'ubuntu ALL=(ALL) NOPASSWD:ALL' > /target/etc/sudoers.d/ubuntu\n    - curtin in-target --target=/target -- chmod 440 /etc/sudoers.d/ubuntu",
          "destination": "/root/user-data"
        }
      ]
    }
  ],
  "provisioners": [
    {
      "type": "shell",
      "execute_command": "echo 'password' | sudo -S sh '{{ .Path }}'",
      "scripts": [
        "scripts/01-base.sh",
        "scripts/02-extras.sh"
      ]
    }
  ],
  "post-processors": [
    {
      "type": "vsphere",
      "vcenter_server": "vcenter.example.com",
      "username": "username",
      "password": "password",
      "insecure_connection": true,
      "datacenter": "datacenter-name",
      "cluster": "cluster-name",
      "template_name": "ubuntu-server",
      "vm_name": "ubuntu-server",
      "convert_to_template": false
    }
  ]
}

