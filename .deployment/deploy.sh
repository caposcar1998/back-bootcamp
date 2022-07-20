aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 036437189980.dkr.ecr.us-east-1.amazonaws.com/gobootcamp
docker pull 036437189980.dkr.ecr.us-east-1.amazonaws.com/gobootcamp:production
docker pull 036437189980.dkr.ecr.us-east-1.amazonaws.com/jsbootcamp:production
BACK_IMAGE_TAG=036437189980.dkr.ecr.us-east-1.amazonaws.com/gobootcamp:production FRONT_IMAGE_TAG=036437189980.dkr.ecr.us-east-1.amazonaws.com/jsbootcamp:production docker-compose -f docker-compose.yml down --remove-orphans
BACK_IMAGE_TAG=036437189980.dkr.ecr.us-east-1.amazonaws.com/gobootcamp:production FRONT_IMAGE_TAG=036437189980.dkr.ecr.us-east-1.amazonaws.com/jsbootcamp:production docker-compose -f docker-compose.yml up -d --force-recreate