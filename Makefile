

dockerbuild_user_service:
	cd user_service && docker build -t abjnet_user_service . 

dockerbuild_user_service_proxy:
	cd user_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/ --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_user_service . 

dockerbuild_email_service:
	cd email_service && docker build -t abjnet_email_service . 

dockerbuild_email_service_proxy:
	cd email_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_email_service . 

dockerbuild_product_service:
	cd product_service && docker build -t abjnet_product_service . 

dockerbuild_product_service_proxy:
	cd product_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_product_service . 

dockerbuild_souscription_service:
	cd souscription_service && docker build -t abjnet_souscription_service . 

dockerbuild_souscription_service_proxy:
	cd souscription_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_souscription_service . 

dockerbuild_taskrunner_service:
	cd taskrunner_service && docker build -t abjnet_taskrunner_service . 

dockerbuild_taskrunner_service_proxy:
	cd taskrunner_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_taskrunner_service . 

dockerbuild_prestation_service:
	cd prestation_service && docker build -t abjnet_prestation_service . 

dockerbuild_prestation_service_proxy:
	cd prestation_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_prestation_service . 

dockerbuild_payment_service:
	cd payment_service && docker build -t abjnet_payment_service . 

dockerbuild_payment_service_proxy:
	cd payment_service && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_payment_service . 

dockerbuild_api_user:
	cd restapi && docker build -t abjnet_api_rest . 

dockerbuild_api_user_proxy:
	 && cd restapi && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_api_rest . 

dockerbuild_api_gateway:
	cd api_gateway && docker build -t abjnet_api_gateway . 

dockerbuild_api_gateway_proxy:
	cd api_gateway && docker build --build-arg HTTP_PROXY=http://172.17.0.1:3128/  --build-arg HTTPS_PROXY=http://172.17.0.1:3128/ -t abjnet_api_gateway . 

dockerbuild_all:
	make dockerbuild_user_service dockerbuild_email_service dockerbuild_product_service dockerbuild_souscription_service dockerbuild_taskrunner_service dockerbuild_prestation_service dockerbuild_payment_service dockerbuild_api_gateway dockerbuild_api_user
dockersave_all:
	cd .. && cd .. && cd shared/containerimages && sudo docker save abjnet_user_service > abjnet_user_service.tar && sudo docker save abjnet_product_service > abjnet_product_service.tar && sudo docker save abjnet_souscription_service > abjnet_souscription_service.tar && sudo docker save abjnet_payment_service > abjnet_payment_service.tar && sudo docker save abjnet_prestation_service > abjnet_prestation_service.tar && sudo docker save abjnet_email_service > abjnet_email_service.tar && sudo docker save abjnet_taskrunner_service > abjnet_taskrunner_service.tar && sudo docker save abjnet_api_rest > abjnet_api_rest.tar    

dockerload_all:
	 cd ~/shared/containerimages && sudo docker load --input abjnet_user_service.tar && sudo docker load --input abjnet_product_service.tar && sudo docker load --input abjnet_souscription_service.tar && sudo docker load --input abjnet_payment_service.tar && sudo docker load --input abjnet_prestation_service.tar && sudo docker load --input abjnet_email_service.tar && sudo docker load --input abjnet_taskrunner_service.tar && sudo docker load --input abjnet_api_rest.tar    
dockerrun_all:
	sudo make dockerload_all && sudo docker-compose down --remove-orphans && sudo docker-compose up

