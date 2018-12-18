#Vagranfile for test purposes

Vagrant.configure("2") do |config|
    config.vm.box = "ubuntu/trusty64"
    
    config.vm.provider "virtualbox" do |vb|
        vb.name = "arlequin"

        vb.cpus = 4
        vb.memory = 2048
        
        vb.gui = false
        vb.linked_clone = true        
    end

    config.vm.network :public_network, :auto_network => true
    config.vm.network "forwarded_port", guest: 443, host: 443
    config.vm.network "forwarded_port", guest: 3000, host: 3000
    
    config.vm.provision :file do |file|
        file.source = "."
        file.destination = "/home/vagrant/arlequin"
    end

    config.vm.provision :docker
    config.vm.provision :docker_compose,
        compose_version: "1.21.0",
        yml: "/home/vagrant/arlequin/docker-compose.yml",
        rebuild: true,
        run: "always"

    config.vm.provision :shell, :inline => "docker exec dockerserver docker pull bash:4.4"
    config.vm.provision :shell, :inline => "docker exec dockerserver docker pull python:2.7"
    config.vm.provision :shell, :inline => "docker exec dockerserver docker pull ruby:2.5"
    # config.vm.provision :shell, :inline => "sudo chmod -R 777 /home/vagrant/arlequin/fileserver/admin/"

end
