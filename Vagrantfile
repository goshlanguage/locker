Vagrant.configure("2") do |config|
  config.vm.define "vagrant.locker"
  config.vm.box = "ubuntu/trusty64"
  config.vm.hostname = "locker"
  config.vm.network :private_network, ip: "192.168.100.42"
  config.vm.synced_folder ".", "/home/vagrant/go/src/github.com/ryanhartje/locker/"

  # Add gopher repo so we have access to Go 1.3+ (seriously..)
  config.vm.provision "shell",
    inline: "sudo add-apt-repository ppa:gophers/archive -y;"

  # Install a go that's not horribly outdated 
  config.vm.provision "shell",
    inline: "apt update && apt install -y golang-1.10 git"

  # Move go binary to something in $PATH
  config.vm.provision "shell",
    inline: "mv /usr/lib/go-1.10/bin/go /usr/bin/go;"

  # set $GOPATH
  config.vm.provision "shell",
    inline: "echo 'export GOPATH=/go' >> /root/.bashrc;"

end
