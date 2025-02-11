install:
	python -m pip install -r requirements-dev.txt
	pre-commit install
	
run-docker:
	bash scripts/reset_redis.sh
	python -O server.py

run-docker-no-linewrap:
	setterm -linewrap off
	bash scripts/reset_redis.sh
	python -O server.py

install-miniconda:
	mkdir -p ~/miniconda3
	# wget https://repo.anaconda.com/miniconda/Miniconda3-latest-Linux-x86_64.sh -O ~/miniconda3/miniconda.sh
	#wget https://repo.anaconda.com/miniconda/Miniconda3-latest-MacOSX-x86_64.sh -O ~/miniconda3/miniconda.sh 
	wget https://repo.anaconda.com/miniconda/Miniconda3-latest-MacOSX-arm64.sh -O ~/miniconda3/miniconda.sh 

	bash ~/miniconda3/miniconda.sh -b -u -p ~/miniconda3
	rm ~/miniconda3/miniconda.sh
	~/miniconda3/bin/conda init --all
	source ~/.zshrc

format:
	pyink .
